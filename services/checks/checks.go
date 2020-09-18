// pmm-managed
// Copyright (C) 2017 Percona LLC
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

// Package checks provides security checks functionality.
package checks

import (
	"bytes"
	"context"
	"io/ioutil"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	api "github.com/percona-platform/saas/gen/check/retrieval"
	"github.com/percona-platform/saas/pkg/check"
	"github.com/percona-platform/saas/pkg/starlark"
	"github.com/percona/pmm/api/agentpb"
	"github.com/percona/pmm/utils/tlsconfig"
	"github.com/percona/pmm/version"
	"github.com/pkg/errors"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services"
	"github.com/percona/pmm-managed/utils/envvars"
)

const (
	defaultRestartInterval = 24 * time.Hour
	defaultStartDelay      = time.Minute

	// Environment variables that affect checks service; only for testing.
	envHost            = "PERCONA_TEST_CHECKS_HOST"
	envPublicKey       = "PERCONA_TEST_CHECKS_PUBLIC_KEY"
	envRestartInterval = "PERCONA_TEST_CHECKS_INTERVAL" // not "restart" in the value - name is fixed
	envCheckFile       = "PERCONA_TEST_CHECKS_FILE"
	envResendInterval  = "PERCONA_TEST_CHECKS_RESEND_INTERVAL"

	checksTimeout       = time.Hour
	downloadTimeout     = 10 * time.Second
	resultTimeout       = 20 * time.Second // should greater than agents.defaultQueryActionTimeout
	resultCheckInterval = time.Second

	// sync with API tests
	resolveTimeoutFactor  = 3
	defaultResendInterval = 2 * time.Second

	prometheusNamespace = "pmm_managed"
	prometheusSubsystem = "checks"

	alertsPrefix        = "/stt/"
	maxSupportedVersion = 1
)

// pmm-agent versions with known changes in Query Actions.
var (
	pmmAgent260     = version.MustParse("2.6.0")
	pmmAgent270     = version.MustParse("2.7.0")
	pmmAgentInvalid = version.MustParse("3.0.0-invalid")
)

var defaultPublicKeys = []string{
	"RWTfyQTP3R7VzZggYY7dzuCbuCQWqTiGCqOvWRRAMVEiw0eSxHMVBBE5", // PMM 2.6
}

// Service is responsible for interactions with Percona Check service.
type Service struct {
	agentsRegistry      agentsRegistry
	alertmanagerService alertmanagerService
	db                  *reform.DB
	alertsRegistry      *registry

	l               *logrus.Entry
	host            string
	publicKeys      []string
	restartInterval time.Duration
	startDelay      time.Duration
	resendInterval  time.Duration

	cm               sync.Mutex
	mySQLChecks      []check.Check
	postgreSQLChecks []check.Check
	mongoDBChecks    []check.Check

	mScriptsExecuted *prom.CounterVec
	mAlertsGenerated *prom.CounterVec
}

// New returns Service with given PMM version.
func New(agentsRegistry agentsRegistry, alertmanagerService alertmanagerService, db *reform.DB) *Service {
	l := logrus.WithField("component", "checks")

	var resendInterval time.Duration
	if d, err := time.ParseDuration(os.Getenv(envResendInterval)); err == nil && d > 0 {
		l.Warnf("Interval changed to %s.", d)
		resendInterval = d
	} else {
		resendInterval = defaultResendInterval
	}

	s := &Service{
		agentsRegistry:      agentsRegistry,
		alertmanagerService: alertmanagerService,
		db:                  db,
		alertsRegistry:      newRegistry(resolveTimeoutFactor * resendInterval),

		l:               l,
		host:            envvars.GetSAASHost(envHost),
		publicKeys:      defaultPublicKeys,
		restartInterval: defaultRestartInterval,
		startDelay:      defaultStartDelay,
		resendInterval:  resendInterval,

		mScriptsExecuted: prom.NewCounterVec(prom.CounterOpts{
			Namespace: prometheusNamespace,
			Subsystem: prometheusSubsystem,
			Name:      "scripts_executed_total",
			Help:      "Counter of check scripts executed per service type",
		}, []string{"service_type"}),

		mAlertsGenerated: prom.NewCounterVec(prom.CounterOpts{
			Namespace: prometheusNamespace,
			Subsystem: prometheusSubsystem,
			Name:      "alerts_generated_total",
			Help:      "Counter of alerts generated per service type per check type",
		}, []string{"service_type", "check_type"}),
	}

	if k := os.Getenv(envPublicKey); k != "" {
		s.publicKeys = strings.Split(k, ",")
		l.Warnf("Public keys changed to %q.", k)
	}
	if d, err := time.ParseDuration(os.Getenv(envRestartInterval)); err == nil && d > 0 {
		l.Warnf("Interval changed to %s; start delay disabled.", d)
		s.restartInterval = d
		s.startDelay = 0
	}

	s.mScriptsExecuted.WithLabelValues(string(models.MySQLServiceType))
	s.mScriptsExecuted.WithLabelValues(string(models.PostgreSQLServiceType))
	s.mScriptsExecuted.WithLabelValues(string(models.MongoDBServiceType))

	s.mAlertsGenerated.WithLabelValues(string(models.MySQLServiceType), string(check.MySQLShow))
	s.mAlertsGenerated.WithLabelValues(string(models.MySQLServiceType), string(check.MySQLSelect))
	s.mAlertsGenerated.WithLabelValues(string(models.PostgreSQLServiceType), string(check.PostgreSQLShow))
	s.mAlertsGenerated.WithLabelValues(string(models.PostgreSQLServiceType), string(check.PostgreSQLSelect))
	s.mAlertsGenerated.WithLabelValues(string(models.MongoDBServiceType), string(check.MongoDBBuildInfo))
	s.mAlertsGenerated.WithLabelValues(string(models.MongoDBServiceType), string(check.MongoDBGetCmdLineOpts))
	s.mAlertsGenerated.WithLabelValues(string(models.MongoDBServiceType), string(check.MongoDBGetParameter))

	return s
}

// Run runs main service loops.
func (s *Service) Run(ctx context.Context) {
	s.l.Info("Starting...")
	defer s.l.Info("Done.")

	// delay for the first run to allow all agents to connect
	startCtx, startCancel := context.WithTimeout(ctx, s.startDelay)
	<-startCtx.Done()
	startCancel()
	if ctx.Err() != nil { // check main context, not startCtx
		return
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.resendAlerts(ctx)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		s.restartChecks(ctx)
	}()

	wg.Wait()
}

// resendAlerts resends collected alerts until ctx is canceled.
func (s *Service) resendAlerts(ctx context.Context) {
	t := time.NewTicker(s.resendInterval)
	defer t.Stop()

	for {
		s.alertmanagerService.SendAlerts(ctx, s.alertsRegistry.collect())

		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// nothing, continue for loop
		}
	}
}

// restartChecks restarts checks until ctx is canceled.
func (s *Service) restartChecks(ctx context.Context) {
	t := time.NewTicker(s.restartInterval)
	defer t.Stop()

	for {
		err := s.StartChecks(ctx)
		switch err {
		case nil:
			// nothing, continue
		case services.ErrSTTDisabled:
			s.l.Info("STT is not enabled, doing nothing.")
		default:
			s.l.Error(err)
		}

		select {
		case <-ctx.Done():
			return
		case <-t.C:
			// nothing, continue for loop
		}
	}
}

// GetSecurityCheckResults returns the results of the STT checks that were run. It returns services.ErrSTTDisabled if STT is disabled.
func (s *Service) GetSecurityCheckResults() ([]check.Result, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.SaaS.STTEnabled {
		return nil, services.ErrSTTDisabled
	}

	results := s.alertsRegistry.getCheckResults()
	checkResults := make([]check.Result, 0, len(results))
	for _, result := range results {
		checkResults = append(checkResults, result.result)
	}

	return checkResults, nil
}

// StartChecks triggers STT checks downloading and execution. It returns services.ErrSTTDisabled if STT is disabled.
func (s *Service) StartChecks(ctx context.Context) error {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return err
	}

	if !settings.SaaS.STTEnabled {
		return services.ErrSTTDisabled
	}

	nCtx, cancel := context.WithTimeout(ctx, checksTimeout)
	defer cancel()

	s.collectChecks(nCtx)
	s.executeChecks(nCtx)
	s.alertmanagerService.SendAlerts(ctx, s.alertsRegistry.collect())

	return nil
}

// getMySQLChecks returns available MySQL checks.
func (s *Service) getMySQLChecks() []check.Check {
	s.cm.Lock()
	defer s.cm.Unlock()

	r := make([]check.Check, 0, len(s.mySQLChecks))
	return append(r, s.mySQLChecks...)
}

// getPostgreSQLChecks returns available PostgreSQL checks.
func (s *Service) getPostgreSQLChecks() []check.Check {
	s.cm.Lock()
	defer s.cm.Unlock()

	r := make([]check.Check, 0, len(s.postgreSQLChecks))
	return append(r, s.postgreSQLChecks...)
}

// getMongoDBChecks returns available MongoDB checks.
func (s *Service) getMongoDBChecks() []check.Check {
	s.cm.Lock()
	defer s.cm.Unlock()

	r := make([]check.Check, 0, len(s.mongoDBChecks))
	return append(r, s.mongoDBChecks...)
}

// waitForResult periodically checks result state and returns it when complete.
func (s *Service) waitForResult(ctx context.Context, resultID string) ([]map[string]interface{}, error) {
	ticker := time.NewTicker(resultCheckInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return nil, errors.WithStack(ctx.Err())
		}

		res, err := models.FindActionResultByID(s.db.Querier, resultID)
		if err != nil {
			return nil, err
		}

		if !res.Done {
			continue
		}

		if err = s.db.Delete(res); err != nil {
			s.l.Warnf("Failed to delete action result %s: %s.", resultID, err)
		}

		if res.Error != "" {
			return nil, errors.Errorf("action %s failed: %s", resultID, res.Error)
		}

		out, err := agentpb.UnmarshalActionQueryResult([]byte(res.Output))
		if err != nil {
			return nil, errors.Errorf("failed to parse action result: %s", err)
		}

		return out, nil
	}
}

// minPMMAgentVersion returns the minimal version of pmm-agent that can handle the given check type.
func (s *Service) minPMMAgentVersion(t check.Type) *version.Parsed {
	switch t {
	case check.MySQLSelect:
		fallthrough
	case check.MySQLShow:
		fallthrough
	case check.PostgreSQLSelect:
		fallthrough
	case check.PostgreSQLShow:
		fallthrough
	case check.MongoDBBuildInfo:
		fallthrough
	case check.MongoDBGetParameter:
		return pmmAgent260

	case check.MongoDBGetCmdLineOpts:
		return pmmAgent270

	default:
		s.l.Warnf("minPMMAgentVersion: unhandled check type %q.", t)
		return pmmAgentInvalid
	}
}

// executeChecks runs all available checks for all reachable services.
func (s *Service) executeChecks(ctx context.Context) {
	s.l.Info("Executing checks...")

	var checkResults []sttCheckResult

	mySQLCheckResults := s.executeMySQLChecks(ctx)
	checkResults = append(checkResults, mySQLCheckResults...)

	postgreSQLCheckResults := s.executePostgreSQLChecks(ctx)
	checkResults = append(checkResults, postgreSQLCheckResults...)

	mongoDBCheckResults := s.executeMongoDBChecks(ctx)
	checkResults = append(checkResults, mongoDBCheckResults...)

	s.alertsRegistry.set(checkResults)
}

// executeMySQLChecks runs MySQL checks for available MySQL services.
func (s *Service) executeMySQLChecks(ctx context.Context) []sttCheckResult {
	checks := s.getMySQLChecks()

	var res []sttCheckResult
	for _, c := range checks {
		pmmAgentVersion := s.minPMMAgentVersion(c.Type)
		targets, err := s.findTargets(models.MySQLServiceType, pmmAgentVersion)
		if err != nil {
			s.l.Warnf("Failed to find proper agents and services for check type: %s and "+
				"min version: %s, reason: %s.", c.Type, pmmAgentVersion, err)
			continue
		}

		for _, target := range targets {
			r, err := models.CreateActionResult(s.db.Querier, target.agentID)
			if err != nil {
				s.l.Warnf("Failed to prepare action result for agent %s: %s.", target.agentID, err)
				continue
			}

			switch c.Type {
			case check.MySQLShow:
				if err := s.agentsRegistry.StartMySQLQueryShowAction(ctx, r.ID, target.agentID, target.dsn, c.Query); err != nil {
					s.l.Warnf("Failed to start MySQL show query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			case check.MySQLSelect:
				if err := s.agentsRegistry.StartMySQLQuerySelectAction(ctx, r.ID, target.agentID, target.dsn, c.Query); err != nil {
					s.l.Warnf("Failed to start MySQL select query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			default:
				s.l.Warnf("Unknown MySQL check type: %s.", c.Type)
				continue
			}

			checkResults, err := s.processResults(ctx, c, target, r.ID)
			if err != nil {
				s.l.Warnf("Failed to process action result: %s.", err)
				continue
			}

			s.mScriptsExecuted.WithLabelValues(string(models.MySQLServiceType)).Inc()
			s.mAlertsGenerated.WithLabelValues(string(models.MySQLServiceType), string(c.Type)).Add(float64(len(checkResults)))
			res = append(res, checkResults...)
		}
	}

	return res
}

// executePostgreSQLChecks runs PostgreSQL checks for available PostgreSQL services.
func (s *Service) executePostgreSQLChecks(ctx context.Context) []sttCheckResult {
	checks := s.getPostgreSQLChecks()

	var res []sttCheckResult
	for _, c := range checks {
		pmmAgentVersion := s.minPMMAgentVersion(c.Type)
		targets, err := s.findTargets(models.PostgreSQLServiceType, pmmAgentVersion)
		if err != nil {
			s.l.Warnf("Failed to find proper agents and services for check type: %s and "+
				"min version: %s, reason: %s.", c.Type, pmmAgentVersion, err)
			continue
		}

		for _, target := range targets {
			r, err := models.CreateActionResult(s.db.Querier, target.agentID)
			if err != nil {
				s.l.Warnf("Failed to prepare action result for agent %s: %s.", target.agentID, err)
				continue
			}

			switch c.Type {
			case check.PostgreSQLShow:
				if err := s.agentsRegistry.StartPostgreSQLQueryShowAction(ctx, r.ID, target.agentID, target.dsn); err != nil {
					s.l.Warnf("Failed to start PostgreSQL show query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			case check.PostgreSQLSelect:
				if err := s.agentsRegistry.StartPostgreSQLQuerySelectAction(ctx, r.ID, target.agentID, target.dsn, c.Query); err != nil {
					s.l.Warnf("Failed to start PostgreSQL select query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			default:
				s.l.Warnf("Unknown PostgresSQL check type: %s.", c.Type)
				continue
			}

			checkResults, err := s.processResults(ctx, c, target, r.ID)
			if err != nil {
				s.l.Warnf("Failed to process action result: %s", err)
				continue
			}

			s.mScriptsExecuted.WithLabelValues(string(models.PostgreSQLServiceType)).Inc()
			s.mAlertsGenerated.WithLabelValues(string(models.PostgreSQLServiceType), string(c.Type)).Add(float64(len(checkResults)))
			res = append(res, checkResults...)
		}
	}

	return res
}

// executeMongoDBChecks runs MongoDB checks for available MongoDB services.
func (s *Service) executeMongoDBChecks(ctx context.Context) []sttCheckResult {
	checks := s.getMongoDBChecks()

	var res []sttCheckResult
	for _, c := range checks {
		pmmAgentVersion := s.minPMMAgentVersion(c.Type)
		targets, err := s.findTargets(models.MongoDBServiceType, pmmAgentVersion)
		if err != nil {
			s.l.Warnf("Failed to find proper agents and services for check type: %s and "+
				"min version: %s, reason: %s.", c.Type, pmmAgentVersion, err)
			continue
		}

		for _, target := range targets {
			r, err := models.CreateActionResult(s.db.Querier, target.agentID)
			if err != nil {
				s.l.Warnf("Failed to prepare action result for agent %s: %s.", target.agentID, err)
				continue
			}

			switch c.Type {
			case check.MongoDBGetParameter:
				if err := s.agentsRegistry.StartMongoDBQueryGetParameterAction(ctx, r.ID, target.agentID, target.dsn); err != nil {
					s.l.Warnf("Failed to start MongoDB get parameter query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			case check.MongoDBBuildInfo:
				if err := s.agentsRegistry.StartMongoDBQueryBuildInfoAction(ctx, r.ID, target.agentID, target.dsn); err != nil {
					s.l.Warnf("Failed to start MongoDB build info query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}
			case check.MongoDBGetCmdLineOpts:
				if err := s.agentsRegistry.StartMongoDBQueryGetCmdLineOptsAction(ctx, r.ID, target.agentID, target.dsn); err != nil {
					s.l.Warnf("Failed to start MongoDB getCmdLineOpts query action for agent %s, reason: %s.", target.agentID, err)
					continue
				}

			default:
				s.l.Warnf("Unknown MongoDB check type: %s.", c.Type)
				continue
			}

			checkResults, err := s.processResults(ctx, c, target, r.ID)
			if err != nil {
				s.l.Warnf("Failed to process action result: %s", err)
				continue
			}

			s.mScriptsExecuted.WithLabelValues(string(models.MongoDBServiceType)).Inc()
			s.mAlertsGenerated.WithLabelValues(string(models.MongoDBServiceType), string(c.Type)).Add(float64(len(checkResults)))
			res = append(res, checkResults...)
		}
	}

	return res
}

type sttCheckResult struct {
	checkName string
	target    target
	result    check.Result
}

func (s *Service) processResults(ctx context.Context, check check.Check, target target, resID string) ([]sttCheckResult, error) {
	nCtx, cancel := context.WithTimeout(ctx, resultTimeout)
	r, err := s.waitForResult(nCtx, resID)
	cancel()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get action result")
	}

	funcs, err := getFuncsForVersion(check.Version)
	if err != nil {
		return nil, err
	}

	env, err := starlark.NewEnv(check.Name, check.Script, funcs)
	if err != nil {
		return nil, errors.Wrap(err, "failed to prepare starlark environment")
	}

	l := s.l.WithFields(logrus.Fields{
		"name":       check.Name,
		"id":         resID,
		"service_id": target.serviceID,
	})
	l.Debugf("Running check script with: %+v.", r)
	results, err := env.Run(check.Name, r, l.Debugln)
	if err != nil {
		return nil, errors.Wrap(err, "failed to execute script")
	}
	l.Infof("Check script returned %d results.", len(results))
	l.Debugf("Results: %+v.", results)

	checkResults := make([]sttCheckResult, len(results))
	for i, result := range results {
		checkResults[i] = sttCheckResult{
			checkName: check.Name,
			target:    target,
			result:    result,
		}
	}

	return checkResults, nil
}

// target contains required info about check target.
type target struct {
	agentID   string
	serviceID string
	labels    map[string]string
	dsn       string
}

// findTargets returns slice of available targets for specified service type.
func (s *Service) findTargets(serviceType models.ServiceType, minPMMAgentVersion *version.Parsed) ([]target, error) {
	var targets []target
	services, err := models.FindServices(s.db.Querier, models.ServiceFilters{ServiceType: &serviceType})
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		// skip pmm own services
		if service.NodeID == models.PMMServerNodeID {
			s.l.Debugf("Skip PMM service, name: %s, type: %s.", service.ServiceName, service.ServiceType)
			continue
		}

		e := s.db.InTransaction(func(tx *reform.TX) error {
			agents, err := models.FindPMMAgentsForService(s.db.Querier, service.ServiceID)
			if err != nil {
				return err
			}
			if len(agents) == 0 {
				return errors.New("no available pmm agents")
			}

			agents = models.FindPMMAgentsForVersion(s.l, agents, minPMMAgentVersion)
			if len(agents) == 0 {
				return errors.New("all available agents are outdated")
			}
			agent := agents[0]

			dsn, err := models.FindDSNByServiceIDandPMMAgentID(s.db.Querier, service.ServiceID, agents[0].AgentID, "")
			if err != nil {
				return err
			}

			node, err := models.FindNodeByID(s.db.Querier, service.NodeID)
			if err != nil {
				return err
			}

			labels, err := models.MergeLabels(node, service, agent)
			if err != nil {
				return err
			}

			targets = append(targets, target{
				agentID:   agent.AgentID,
				serviceID: service.ServiceID,
				labels:    labels,
				dsn:       dsn,
			})
			return nil
		})
		if e != nil {
			s.l.Errorf("Failed to find agents for service %s, reason: %s.", service.ServiceID, e)
		}
	}

	return targets, nil
}

// groupChecksByDB splits provided checks by database and returns three slices: for MySQL, for PostgreSQL and for MongoDB.
func (s *Service) groupChecksByDB(checks []check.Check) (mySQLChecks, postgreSQLChecks, mongoDBChecks []check.Check) {
	for _, c := range checks {
		switch c.Type {
		case check.MySQLSelect:
			fallthrough
		case check.MySQLShow:
			mySQLChecks = append(mySQLChecks, c)

		case check.PostgreSQLSelect:
			fallthrough
		case check.PostgreSQLShow:
			postgreSQLChecks = append(postgreSQLChecks, c)

		case check.MongoDBGetParameter:
			fallthrough
		case check.MongoDBBuildInfo:
			fallthrough
		case check.MongoDBGetCmdLineOpts:
			mongoDBChecks = append(mongoDBChecks, c)

		default:
			s.l.Warnf("Unknown check type %s, skip it.", c.Type)
		}
	}

	return
}

// collectChecks loads checks from file or SaaS, and stores versions this pmm-managed can handle.
func (s *Service) collectChecks(ctx context.Context) {
	var checks []check.Check
	var err error
	if f := os.Getenv(envCheckFile); f != "" {
		s.l.Warnf("Using local test checks file: %s.", f)
		checks, err = s.loadLocalChecks(f)
		if err != nil {
			s.l.Errorf("Failed to load local checks file: %s.", err)
			return // keep previously loaded checks
		}
	} else {
		checks, err = s.downloadChecks(ctx)
		if err != nil {
			s.l.Errorf("Failed to download checks: %s.", err)
			return // keep previously downloaded checks
		}
	}

	checks = s.filterSupportedChecks(checks)
	mySQLChecks, postgreSQLChecks, mongoDBChecks := s.groupChecksByDB(checks)

	s.updateChecks(mySQLChecks, postgreSQLChecks, mongoDBChecks)
}

// loadLocalCheck loads checks form local file.
func (s *Service) loadLocalChecks(file string) ([]check.Check, error) {
	data, err := ioutil.ReadFile(file) //nolint:gosec
	if err != nil {
		return nil, errors.Wrap(err, "failed to read test checks file")
	}

	// be strict about local files
	params := &check.ParseParams{
		DisallowUnknownFields: true,
		DisallowInvalidChecks: true,
	}
	checks, err := check.Parse(bytes.NewReader(data), params)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse test checks file")
	}

	return checks, nil
}

// downloadChecks downloads checks form percona service endpoint.
func (s *Service) downloadChecks(ctx context.Context) ([]check.Check, error) {
	s.l.Infof("Downloading checks from %s ...", s.host)

	host, _, err := net.SplitHostPort(s.host)
	if err != nil {
		return nil, errors.Wrap(err, "failed to set checks host")
	}
	tlsConfig := tlsconfig.Get()
	tlsConfig.ServerName = host

	opts := []grpc.DialOption{
		// replacement is marked as experimental
		grpc.WithBackoffMaxDelay(downloadTimeout), //nolint:staticcheck

		grpc.WithBlock(),
		grpc.WithUserAgent("pmm-managed/" + version.Version),
		grpc.WithTransportCredentials(credentials.NewTLS(tlsConfig)),
	}

	ctx, cancel := context.WithTimeout(ctx, downloadTimeout)
	defer cancel()
	cc, err := grpc.DialContext(ctx, s.host, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to dial")
	}
	defer cc.Close() //nolint:errcheck

	resp, err := api.NewRetrievalAPIClient(cc).GetAllChecks(ctx, &api.GetAllChecksRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to request checks service")
	}

	if err = s.verifySignatures(resp); err != nil {
		return nil, err
	}

	// be liberal about files from SaaS for smooth transition to future versions
	params := &check.ParseParams{
		DisallowUnknownFields: false,
		DisallowInvalidChecks: false,
	}
	checks, err := check.Parse(strings.NewReader(resp.File), params)
	if err != nil {
		return nil, err
	}

	return checks, nil
}

// filterSupportedChecks returns supported checks and prints warning log messages about unsupported.
func (s *Service) filterSupportedChecks(checks []check.Check) []check.Check {
	res := make([]check.Check, 0, len(checks))
	for _, c := range checks {
		if c.Version > maxSupportedVersion {
			s.l.Warnf("Unsupported checks version: %d, max supported version: %d.", c.Version, maxSupportedVersion)
			continue
		}

		switch c.Type {
		case check.MySQLShow:
		case check.MySQLSelect:
		case check.PostgreSQLShow:
		case check.PostgreSQLSelect:
		case check.MongoDBGetParameter:
		case check.MongoDBBuildInfo:
		case check.MongoDBGetCmdLineOpts:
		default:
			s.l.Warnf("Unsupported check type: %s.", c.Type)
			continue
		}

		res = append(res, c)
	}

	return res
}

// updateChecks update service checks filed value under mutex.
func (s *Service) updateChecks(mySQLChecks, postgreSQLChecks, mongoDBChecks []check.Check) {
	s.cm.Lock()
	defer s.cm.Unlock()

	s.mySQLChecks = mySQLChecks
	s.postgreSQLChecks = postgreSQLChecks
	s.mongoDBChecks = mongoDBChecks
}

// verifySignatures verifies checks signatures and returns error in case of verification problem.
func (s *Service) verifySignatures(resp *api.GetAllChecksResponse) error {
	if len(resp.Signatures) == 0 {
		return errors.New("zero signatures received")
	}

	var err error
	for _, sign := range resp.Signatures {
		for _, key := range s.publicKeys {
			if err = check.Verify([]byte(resp.File), key, sign); err == nil {
				return nil
			}
			s.l.Debugf("Key %q doesn't match signature %q: %s.", key, sign, err)
		}
	}

	return errors.New("no verified signatures")
}

// Describe implements prom.Collector.
func (s *Service) Describe(ch chan<- *prom.Desc) {
	s.mScriptsExecuted.Describe(ch)
	s.mAlertsGenerated.Describe(ch)
}

// Collect implements prom.Collector.
func (s *Service) Collect(ch chan<- prom.Metric) {
	s.mScriptsExecuted.Collect(ch)
	s.mAlertsGenerated.Collect(ch)
}

// check interfaces
var (
	_ prom.Collector = (*Service)(nil)
)
