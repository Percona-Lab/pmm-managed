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

package ia

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/percona-platform/saas/pkg/common"
	iav1beta1 "github.com/percona/pmm/api/managementpb/ia"
	"github.com/percona/promconfig"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"
	"gopkg.in/yaml.v3"

	"github.com/percona/pmm-managed/models"
	"github.com/percona/pmm-managed/services"
	"github.com/percona/pmm-managed/utils/dir"
)

const rulesDir = "/etc/ia/rules"

// RulesService represents API for Integrated Alerting Rules.
type RulesService struct {
	db           *reform.DB
	l            *logrus.Entry
	templates    *TemplatesService
	vmalert      vmAlert
	alertManager alertManager
	rulesPath    string // used for testing
}

// NewRulesService creates an API for Integrated Alerting Rules.
func NewRulesService(db *reform.DB, templates *TemplatesService, vmalert vmAlert, alertManager alertManager) *RulesService {
	l := logrus.WithField("component", "management/ia/rules")

	err := dir.CreateDataDir(rulesDir, "pmm", "pmm", dirPerm)
	if err != nil {
		l.Error(err)
	}

	return &RulesService{
		db:           db,
		l:            l,
		templates:    templates,
		vmalert:      vmalert,
		alertManager: alertManager,
		rulesPath:    rulesDir,
	}
}

// TODO Move this and related types to https://github.com/percona/promconfig
// https://jira.percona.com/browse/PMM-7069
type ruleFile struct {
	Group []ruleGroup `yaml:"groups"`
}

type ruleGroup struct {
	Name  string `yaml:"name"`
	Rules []rule `yaml:"rules"`
}

type rule struct {
	Alert       string              `yaml:"alert"` // Rule ID.
	Expr        string              `yaml:"expr"`
	Duration    promconfig.Duration `yaml:"for"`
	Labels      map[string]string   `yaml:"labels,omitempty"`
	Annotations map[string]string   `yaml:"annotations,omitempty"`
}

// writeVMAlertRulesFiles converts all available rules to VMAlert rule files.
func (s *RulesService) writeVMAlertRulesFiles() {
	rules, err := s.getAlertRules()
	if err != nil {
		s.l.Errorf("Failed to get available alert rules: %+v", err)
		return
	}

	ruleFiles, err := s.prepareRulesFiles(rules)
	if err != nil {
		s.l.Errorf("Failed to prepare alert rule files: %+v", err)
		return
	}

	matches, err := filepath.Glob(s.rulesPath + "/*.yml")
	if err != nil {
		s.l.Errorf("Failed to clean old alert rule files: %+v", err)
		return
	}

	for _, match := range matches {
		if err = os.RemoveAll(match); err != nil {
			s.l.Errorf("Failed to remove old rule file: %+v", err)
		}
	}

	for _, file := range ruleFiles {
		err = s.writeRuleFile(&file) //nolint:gosec
		if err != nil {
			s.l.Errorf("Failed to write alert rule file: %+v", err)
		}
	}
}

// prepareRulesFiles converts collected IA rules to Alertmanager rule files content.
func (s *RulesService) prepareRulesFiles(rules []*iav1beta1.Rule) ([]ruleFile, error) {
	res := make([]ruleFile, 0, len(rules))
	for _, ruleM := range rules {
		if ruleM.Disabled {
			s.l.Debugf("Skipping rule %s as it is disabled.", ruleM.RuleId)
			continue
		}

		r := rule{
			Alert:       ruleM.RuleId,
			Duration:    promconfig.Duration(ruleM.For.AsDuration()),
			Labels:      make(map[string]string),
			Annotations: make(map[string]string),
		}

		params := make(map[string]string, len(ruleM.Params))
		for _, p := range ruleM.Params {
			var value string
			switch p.Type {
			case iav1beta1.ParamType_FLOAT:
				value = fmt.Sprint(p.GetFloat())
			case iav1beta1.ParamType_BOOL:
				value = fmt.Sprint(p.GetBool())
			case iav1beta1.ParamType_STRING:
				value = fmt.Sprint(p.GetString_())
			case iav1beta1.ParamType_PARAM_TYPE_INVALID:
				s.l.Warnf("Invalid parameter type %s", p.Type)
				continue
			}

			params[p.Name] = value
		}

		var buf bytes.Buffer
		t, err := newParamTemplate().Parse(ruleM.Template.Expr)
		if err != nil {
			return nil, errors.Wrap(err, "Failed to parse rule expression")
		}
		if err = t.Execute(&buf, params); err != nil {
			return nil, errors.Wrap(err, "Failed to fill expression placeholders")
		}
		r.Expr = buf.String()

		// Copy annotations form template
		if err = transformMaps(ruleM.Template.Annotations, r.Annotations, params); err != nil {
			return nil, errors.Wrap(err, "Failed to fill template annotations placeholders")
		}

		r.Annotations["rule"] = ruleM.Summary

		// Copy labels form template
		if err = transformMaps(ruleM.Template.Labels, r.Labels, params); err != nil {
			return nil, errors.Wrap(err, "Failed to fill template labels placeholders")
		}

		// Add rule labels
		if err = transformMaps(ruleM.CustomLabels, r.Labels, params); err != nil {
			return nil, errors.Wrap(err, "Failed to fill rule labels placeholders")
		}

		// Do not add volatile values like `{{ $value }}` to labels as it will break alerts identity.
		r.Labels["ia"] = "1"
		r.Labels["severity"] = ruleM.Severity.String() // FIXME https://jira.percona.com/browse/PMM-7228
		r.Labels["rule_id"] = ruleM.RuleId
		r.Labels["template_name"] = ruleM.Template.Name

		res = append(res, ruleFile{
			Group: []ruleGroup{{
				Name:  "PMM Integrated Alerting",
				Rules: []rule{r},
			}},
		})
	}

	return res, nil
}

// fills templates found in labels and annotaitons with values.
func transformMaps(src map[string]string, dest map[string]string, data map[string]string) error {
	var buf bytes.Buffer

	for k, v := range src {
		buf.Reset()
		t, err := newParamTemplate().Parse(v)
		if err != nil {
			return err
		}
		if err = t.Execute(&buf, data); err != nil {
			return err
		}
		dest[k] = buf.String()
	}
	return nil
}

// dump the transformed IA templates to a file.
func (s *RulesService) writeRuleFile(rule *ruleFile) error {
	b, err := yaml.Marshal(rule)
	if err != nil {
		return errors.Errorf("failed to marshal rule %s", err)
	}
	b = append([]byte("---\n"), b...)

	alertRule := rule.Group[0].Rules[0]
	if alertRule.Alert == "" {
		return errors.New("alert rule not initialized")
	}

	fileName := strings.TrimPrefix(alertRule.Alert, "/rule_id/")
	path := s.rulesPath + "/" + fileName + ".yml"
	if err = ioutil.WriteFile(path, b, 0o644); err != nil {
		return errors.Errorf("failed to dump rule to file %s: %s", s.rulesPath, err)
	}

	return nil
}

// ListAlertRules returns a list of all Integrated Alerting rules.
func (s *RulesService) ListAlertRules(ctx context.Context, req *iav1beta1.ListAlertRulesRequest) (*iav1beta1.ListAlertRulesResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", services.ErrAlertingDisabled)
	}

	res, err := s.getAlertRules()
	if err != nil {
		return nil, err
	}
	return &iav1beta1.ListAlertRulesResponse{Rules: res}, nil
}

// getAlertRules returns list of available alert rules.
func (s *RulesService) getAlertRules() ([]*iav1beta1.Rule, error) {
	var rules []*models.Rule
	var channels []*models.Channel
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		rules, err = models.FindRules(tx.Querier)
		if err != nil {
			return err
		}

		channels, err = models.FindChannels(tx.Querier)
		if err != nil {
			return err
		}
		return nil
	})

	if e != nil {
		return nil, e
	}

	templates := s.templates.getTemplates()

	res := make([]*iav1beta1.Rule, 0, len(rules))
	for _, rule := range rules {
		template, ok := templates[rule.TemplateName]
		if !ok {
			s.l.Warnf("Template %s used by rule %s doesn't exist, skipping that rule", template.Name, rule.ID)
			continue
		}

		r, err := convertRule(s.l, rule, template, channels)
		if err != nil {
			return nil, err
		}
		res = append(res, r)
	}

	return res, nil
}

// CreateAlertRule creates Integrated Alerting rule.
func (s *RulesService) CreateAlertRule(ctx context.Context, req *iav1beta1.CreateAlertRuleRequest) (*iav1beta1.CreateAlertRuleResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", services.ErrAlertingDisabled)
	}

	params := &models.CreateRuleParams{
		TemplateName: req.TemplateName,
		Summary:      req.Summary,
		Disabled:     req.Disabled,
		For:          req.For.AsDuration(),
		Severity:     common.Severity(req.Severity),
		CustomLabels: req.CustomLabels,
		ChannelIDs:   req.ChannelIds,
		Filters:      convertFiltersToModel(req.Filters),
	}

	params.RuleParams, err = convertRuleParamsToModel(req.Params)
	if err != nil {
		return nil, err
	}

	if _, ok := s.templates.getTemplates()[params.TemplateName]; !ok {
		return nil, status.Errorf(codes.NotFound, "Unknown template %s.", params.TemplateName)
	}

	var rule *models.Rule
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		rule, err = models.CreateRule(tx.Querier, params)
		return err
	})
	if e != nil {
		return nil, e
	}

	s.writeVMAlertRulesFiles()
	s.vmalert.RequestConfigurationUpdate()
	s.alertManager.RequestConfigurationUpdate()

	return &iav1beta1.CreateAlertRuleResponse{RuleId: rule.ID}, nil
}

// UpdateAlertRule updates Integrated Alerting rule.
func (s *RulesService) UpdateAlertRule(ctx context.Context, req *iav1beta1.UpdateAlertRuleRequest) (*iav1beta1.UpdateAlertRuleResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", services.ErrAlertingDisabled)
	}

	params := &models.ChangeRuleParams{
		Disabled:     req.Disabled,
		For:          req.For.AsDuration(),
		Severity:     common.Severity(req.Severity),
		CustomLabels: req.CustomLabels,
		ChannelIDs:   req.ChannelIds,
	}

	ruleParams, err := convertRuleParamsToModel(req.Params)
	if err != nil {
		return nil, err
	}
	params.RuleParams = ruleParams
	params.Filters = convertFiltersToModel(req.Filters)

	e := s.db.InTransaction(func(tx *reform.TX) error {
		_, err := models.ChangeRule(tx.Querier, req.RuleId, params)
		return err
	})
	if e != nil {
		return nil, e
	}

	s.writeVMAlertRulesFiles()
	s.vmalert.RequestConfigurationUpdate()
	s.alertManager.RequestConfigurationUpdate()

	return &iav1beta1.UpdateAlertRuleResponse{}, nil
}

// ToggleAlertRule allows to switch between disabled and enabled states of an Alert Rule.
func (s *RulesService) ToggleAlertRule(ctx context.Context, req *iav1beta1.ToggleAlertRuleRequest) (*iav1beta1.ToggleAlertRuleResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", services.ErrAlertingDisabled)
	}

	var params models.ChangeRuleParams
	switch req.Disabled {
	case iav1beta1.BooleanFlag_DO_NOT_CHANGE:
		return &iav1beta1.ToggleAlertRuleResponse{}, nil
	case iav1beta1.BooleanFlag_TRUE:
		params.Disabled = true
	case iav1beta1.BooleanFlag_FALSE:
		// nothing
	}

	e := s.db.InTransaction(func(tx *reform.TX) error {
		_, err := models.ChangeRule(tx.Querier, req.RuleId, &params)
		return err
	})
	if e != nil {
		return nil, e
	}

	s.writeVMAlertRulesFiles()
	s.vmalert.RequestConfigurationUpdate()
	s.alertManager.RequestConfigurationUpdate()

	return &iav1beta1.ToggleAlertRuleResponse{}, nil
}

// DeleteAlertRule deletes Integrated Alerting rule.
func (s *RulesService) DeleteAlertRule(ctx context.Context, req *iav1beta1.DeleteAlertRuleRequest) (*iav1beta1.DeleteAlertRuleResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", services.ErrAlertingDisabled)
	}

	e := s.db.InTransaction(func(tx *reform.TX) error {
		return models.RemoveRule(tx.Querier, req.RuleId)
	})
	if e != nil {
		return nil, e
	}

	s.writeVMAlertRulesFiles()
	s.vmalert.RequestConfigurationUpdate()
	s.alertManager.RequestConfigurationUpdate()

	return &iav1beta1.DeleteAlertRuleResponse{}, nil
}

func convertModelToRuleParams(params models.RuleParams) ([]*iav1beta1.RuleParam, error) {
	res := make([]*iav1beta1.RuleParam, len(params))
	for i, param := range params {
		p := &iav1beta1.RuleParam{Name: param.Name}

		switch param.Type {
		case models.Bool:
			p.Type = iav1beta1.ParamType_BOOL
			p.Value = &iav1beta1.RuleParam_Bool{Bool: param.BoolValue}
		case models.Float:
			p.Type = iav1beta1.ParamType_FLOAT
			p.Value = &iav1beta1.RuleParam_Float{Float: param.FloatValue}
		case models.String:
			p.Type = iav1beta1.ParamType_STRING
			p.Value = &iav1beta1.RuleParam_String_{String_: param.StringValue}
		default:
			return nil, errors.New("invalid rule param value type")
		}
		res[i] = p
	}
	return res, nil
}

func convertRuleParamsToModel(params []*iav1beta1.RuleParam) (models.RuleParams, error) {
	ruleParams := make(models.RuleParams, len(params))
	for i, param := range params {
		p := models.RuleParam{Name: param.Name}

		switch param.Type {
		case iav1beta1.ParamType_PARAM_TYPE_INVALID:
			return nil, errors.New("invalid model rule param value type")
		case iav1beta1.ParamType_BOOL:
			p.Type = models.Bool
			p.BoolValue = param.GetBool()
		case iav1beta1.ParamType_FLOAT:
			p.Type = models.Float
			p.FloatValue = param.GetFloat()
		case iav1beta1.ParamType_STRING:
			p.Type = models.Float
			p.StringValue = param.GetString_()
		default:
			return nil, errors.New("invalid model rule param value type")
		}

		ruleParams[i] = p
	}
	return ruleParams, nil
}

func convertModelToFilterType(filterType models.FilterType) iav1beta1.FilterType {
	switch filterType {
	case models.Invalid:
		return iav1beta1.FilterType_FILTER_TYPE_INVALID
	case models.Equal:
		return iav1beta1.FilterType_EQUAL
	case models.NotEqual:
		return iav1beta1.FilterType_NOT_EQUAL
	case models.Regex:
		return iav1beta1.FilterType_REGEX
	case models.NotRegex:
		return iav1beta1.FilterType_NOT_REGEX
	default:
		return iav1beta1.FilterType_FILTER_TYPE_INVALID
	}
}

func convertFiltersToModel(filters []*iav1beta1.Filter) models.Filters {
	res := make(models.Filters, len(filters))
	for i, filter := range filters {
		f := models.Filter{
			Key: filter.Key,
		}

		// Unquote the first encountered quote.
		// Do it only for filters as only they can be set in PMM 2.13.
		f.Val = filter.Value
		for _, q := range []string{`"`, `'`} {
			if strings.HasPrefix(f.Val, q) && strings.HasSuffix(f.Val, q) {
				f.Val = strings.TrimPrefix(f.Val, q)
				f.Val = strings.TrimSuffix(f.Val, q)
				break
			}
		}

		switch filter.Type {
		case iav1beta1.FilterType_FILTER_TYPE_INVALID:
			f.Type = models.Invalid
		case iav1beta1.FilterType_EQUAL:
			f.Type = models.Equal
		case iav1beta1.FilterType_NOT_EQUAL:
			f.Type = models.NotEqual
		case iav1beta1.FilterType_REGEX:
			f.Type = models.Regex
		case iav1beta1.FilterType_NOT_REGEX:
			f.Type = models.NotRegex
		default:
			f.Type = models.Invalid
		}
		res[i] = f
	}
	return res
}

// Check interfaces.
var (
	_ iav1beta1.RulesServer = (*RulesService)(nil)
)
