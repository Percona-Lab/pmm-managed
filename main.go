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

package main

import (
	"bytes"
	"context"
	_ "expvar"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	agentAPI "github.com/percona/pmm/api/agent"
	inventoryAPI "github.com/percona/pmm/api/inventory"
	serverAPI "github.com/percona/pmm/api/server"
	"github.com/percona/pmm/version"
	prom "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"

	"github.com/percona/pmm-managed/handlers"
	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services/agents"
	"github.com/percona/pmm-managed/services/inventory"
	"github.com/percona/pmm-managed/services/logs"
	"github.com/percona/pmm-managed/services/prometheus"
	"github.com/percona/pmm-managed/services/telemetry"
	"github.com/percona/pmm-managed/utils/interceptors"
	"github.com/percona/pmm-managed/utils/logger"
	"github.com/percona/pmm-managed/utils/ports"
)

const (
	shutdownTimeout = 3 * time.Second
)

var (
	// TODO Switch to kingpin for flags parsing: https://jira.percona.com/browse/PMM-3259

	gRPCAddrF  = flag.String("listen-grpc-addr", "127.0.0.1:7771", "gRPC APIs server listen address")
	jsonAddrF  = flag.String("listen-json-addr", "127.0.0.1:7772", "JSON APIs server listen address")
	debugAddrF = flag.String("listen-debug-addr", "127.0.0.1:7773", "Debug server listen address")

	prometheusConfigF = flag.String("prometheus-config", "", "Prometheus configuration file path")
	prometheusURLF    = flag.String("prometheus-url", "http://127.0.0.1:9090/", "Prometheus base URL")
	promtoolF         = flag.String("promtool", "promtool", "promtool path")

	grafanaAddrF = flag.String("grafana-addr", "127.0.0.1:3000", "Grafana HTTP API address")

	dbNameF     = flag.String("db-name", "", "Database name")
	dbUsernameF = flag.String("db-username", "pmm-managed", "Database username")
	dbPasswordF = flag.String("db-password", "pmm-managed", "Database password")

	postgresDBNameF     = flag.String("postgres-name", "", "PostgreSQL database name")
	postgresDBUsernameF = flag.String("postgres-username", "pmm-managed", "PostgreSQL database username")
	postgresDBPasswordF = flag.String("postgres-password", "pmm-managed", "PostgreSQL database password")

	agentMySQLdExporterF    = flag.String("agent-mysqld-exporter", "/usr/local/percona/pmm-client/mysqld_exporter", "mysqld_exporter path")
	agentPostgresExporterF  = flag.String("agent-postgres-exporter", "/usr/local/percona/pmm-client/postgres_exporter", "postgres_exporter path")
	agentRDSExporterF       = flag.String("agent-rds-exporter", "/usr/sbin/rds_exporter", "rds_exporter path")
	agentRDSExporterConfigF = flag.String("agent-rds-exporter-config", "/etc/percona-rds-exporter.yml", "rds_exporter configuration file path")

	debugF = flag.Bool("debug", false, "Enable debug logging")
	traceF = flag.Bool("trace", false, "Enable trace logging")
)

func addLogsHandler(mux *http.ServeMux, logs *logs.Logs) {
	l := logrus.WithField("component", "logs.zip")

	mux.HandleFunc("/logs.zip", func(rw http.ResponseWriter, req *http.Request) {
		// fail-safe
		ctx, cancel := context.WithTimeout(req.Context(), 10*time.Second)
		defer cancel()

		filename := fmt.Sprintf("pmm-server_%s", time.Now().UTC().Format("2006-01-02_15-04"))
		rw.Header().Set(`Access-Control-Allow-Origin`, `*`)
		rw.Header().Set(`Content-Type`, `application/zip`)
		rw.Header().Set(`Content-Disposition`, `attachment; filename="`+filename+`"`)
		ctx = logger.Set(ctx, "logs")
		if err := logs.Zip(ctx, rw); err != nil {
			l.Error(err)
		}
	})
}

type serviceDependencies struct {
	prometheus     *prometheus.Service
	db             *reform.DB
	portsRegistry  *ports.Registry
	agentsRegistry *agents.Registry
	logs           *logs.Logs
}

// runGRPCServer runs gRPC server until context is canceled, then gracefully stops it.
func runGRPCServer(ctx context.Context, deps *serviceDependencies) {
	l := logrus.WithField("component", "gRPC")
	l.Infof("Starting server on http://%s/ ...", *gRPCAddrF)

	gRPCServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.Unary),
		grpc.StreamInterceptor(interceptors.Stream),
	)
	serverAPI.RegisterServerServer(gRPCServer, handlers.NewServerServer(version.Version))
	agentAPI.RegisterAgentServer(gRPCServer, &handlers.AgentServer{
		Registry: deps.agentsRegistry,
	})
	inventoryAPI.RegisterNodesServer(gRPCServer, handlers.NewNodesServer(
		inventory.NewNodesService(deps.db.Querier, deps.agentsRegistry),
	))
	inventoryAPI.RegisterServicesServer(gRPCServer, handlers.NewServicesServer(
		inventory.NewServicesService(deps.db.Querier, deps.agentsRegistry),
	))
	inventoryAPI.RegisterAgentsServer(gRPCServer, handlers.NewAgentsServer(
		inventory.NewAgentsService(deps.agentsRegistry),
		deps.db,
	))

	if *debugF {
		l.Debug("Reflection enabled.")
		reflection.Register(gRPCServer)
	}

	grpc_prometheus.EnableHandlingTimeHistogram()
	grpc_prometheus.Register(gRPCServer)

	listener, err := net.Listen("tcp", *gRPCAddrF)
	if err != nil {
		l.Panic(err)
	}
	go func() {
		for {
			err = gRPCServer.Serve(listener)
			if err == nil || err == grpc.ErrServerStopped {
				break
			}
			l.Errorf("Failed to serve: %s", err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	go func() {
		<-ctx.Done()
		gRPCServer.Stop()
	}()
	gRPCServer.GracefulStop()
	cancel()
}

// runJSONServer runs JSON proxy server (grpc-gateway) until context is canceled, then gracefully stops it.
func runJSONServer(ctx context.Context, logs *logs.Logs) {
	l := logrus.WithField("component", "JSON")
	l.Infof("Starting server on http://%s/ ...", *jsonAddrF)

	proxyMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	type registrar func(context.Context, *runtime.ServeMux, string, []grpc.DialOption) error
	for _, r := range []registrar{
		serverAPI.RegisterServerHandlerFromEndpoint,
		inventoryAPI.RegisterNodesHandlerFromEndpoint,
		inventoryAPI.RegisterServicesHandlerFromEndpoint,
		inventoryAPI.RegisterAgentsHandlerFromEndpoint,
	} {
		if err := r(ctx, proxyMux, *gRPCAddrF, opts); err != nil {
			l.Panic(err)
		}
	}

	mux := http.NewServeMux()
	addLogsHandler(mux, logs)
	mux.Handle("/", proxyMux)

	server := &http.Server{
		Addr:     *jsonAddrF,
		ErrorLog: log.New(os.Stderr, "runJSONServer: ", 0),
		Handler:  mux,
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

// runDebugServer runs debug server until context is canceled, then gracefully stops it.
func runDebugServer(ctx context.Context, collectors ...prom.Collector) {
	prom.MustRegister(collectors...)
	handler := promhttp.HandlerFor(prom.DefaultGatherer, promhttp.HandlerOpts{
		ErrorLog:      logrus.WithField("component", "metrics"),
		ErrorHandling: promhttp.ContinueOnError,
	})
	http.Handle("/debug/metrics", promhttp.InstrumentMetricHandler(prom.DefaultRegisterer, handler))

	l := logrus.WithField("component", "debug")

	handlers := []string{"/debug/metrics", "/debug/vars", "/debug/requests", "/debug/events", "/debug/pprof"}
	for i, h := range handlers {
		handlers[i] = "http://" + *debugAddrF + h
	}

	var buf bytes.Buffer
	err := template.Must(template.New("debug").Parse(`
	<html>
	<body>
	<ul>
	{{ range . }}
		<li><a href="{{ . }}">{{ . }}</a></li>
	{{ end }}
	</ul>
	</body>
	</html>
	`)).Execute(&buf, handlers)
	if err != nil {
		l.Panic(err)
	}
	http.HandleFunc("/debug", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(buf.Bytes())
	})
	l.Infof("Starting server on http://%s/debug\nRegistered handlers:\n\t%s", *debugAddrF, strings.Join(handlers, "\n\t"))

	server := &http.Server{
		Addr:     *debugAddrF,
		ErrorLog: log.New(os.Stderr, "runDebugServer: ", 0),
	}
	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			l.Panic(err)
		}
		l.Info("Server stopped.")
	}()

	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	if err := server.Shutdown(ctx); err != nil {
		l.Errorf("Failed to shutdown gracefully: %s", err)
	}
	cancel()
}

func runTelemetryService(ctx context.Context, db *reform.DB) {
	l := logrus.WithField("component", "telemetry")

	uuid, err := telemetry.GetTelemetryUUID(db)
	if err != nil {
		l.Panicf("cannot get/set telemetry UUID in DB: %+v", err)
	}

	svc := telemetry.NewService(uuid, version.Version)
	svc.Run(ctx)
}

func main() {
	log.SetFlags(0)
	log.Printf("%s.", version.ShortInfo())
	log.SetPrefix("stdlog: ")
	flag.Parse()

	if *dbNameF == "" {
		log.Fatal("-db-name flag must be given explicitly.")
	}
	if *postgresDBNameF == "" {
		log.Fatal("-postgres-name flag must be given explicitly.")
	}

	if *debugF {
		logrus.SetLevel(logrus.DebugLevel)
	}
	if *traceF {
		logrus.SetLevel(logrus.TraceLevel)
		logrus.SetReportCaller(true)
		grpclog.SetLoggerV2(&logger.GRPC{Entry: logrus.WithField("component", "grpclog")})
	}

	logrus.Infof("Log level: %s.", logrus.GetLevel())

	l := logrus.WithField("component", "main")
	ctx, cancel := context.WithCancel(context.Background())
	ctx = logger.Set(ctx, "main")
	defer l.Info("Done.")

	// handle termination signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT)
	go func() {
		s := <-signals
		signal.Stop(signals)
		logrus.Warnf("Got %s, shutting down...", unix.SignalName(s.(unix.Signal)))
		cancel()
	}()

	postgresDB, err := models.OpenDB(*postgresDBNameF, *postgresDBUsernameF, *postgresDBPasswordF, l.Debugf)
	if err != nil {
		l.Panicf("Failed to connect to database: %+v", err)
	}
	defer postgresDB.Close()
	db := reform.NewDB(postgresDB, postgresql.Dialect, nil)

	prometheus, err := prometheus.NewService(*prometheusConfigF, *promtoolF, db, *prometheusURLF)
	if err == nil {
		err = prometheus.Check(ctx)
	}
	if err != nil {
		l.Panicf("Prometheus service problem: %+v", err)
	}

	agentsRegistry := agents.NewRegistry(db, prometheus)
	logs := logs.New(version.Version)

	deps := &serviceDependencies{
		prometheus:     prometheus,
		db:             db,
		portsRegistry:  ports.NewRegistry(10000, 10999, nil),
		agentsRegistry: agentsRegistry,
		logs:           logs,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		runGRPCServer(ctx, deps)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runJSONServer(ctx, logs)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runDebugServer(ctx, agentsRegistry)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		runTelemetryService(ctx, db)
	}()

	wg.Wait()
}
