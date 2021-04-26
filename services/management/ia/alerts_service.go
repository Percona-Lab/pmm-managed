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

// Package ia contains Integrated Alerting APIs implementations.
package ia

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/percona-platform/saas/pkg/common"
	"github.com/percona/pmm/api/alertmanager/ammodels"
	"github.com/percona/pmm/api/managementpb"
	iav1beta1 "github.com/percona/pmm/api/managementpb/ia"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// AlertsService represents integrated alerting alerts API.
type AlertsService struct {
	db               *reform.DB
	l                *logrus.Entry
	alertManager     alertManager
	templatesService *TemplatesService
}

// NewAlertsService creates new alerts API service.
func NewAlertsService(db *reform.DB, alertManager alertManager, templatesService *TemplatesService) *AlertsService {
	return &AlertsService{
		l:                logrus.WithField("component", "management/ia/alerts"),
		db:               db,
		alertManager:     alertManager,
		templatesService: templatesService,
	}
}

// Enabled returns if service is enabled and can be used.
func (s *AlertsService) Enabled() bool {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		s.l.WithError(err).Error("can't get settings")
		return false
	}
	return settings.IntegratedAlerting.Enabled
}

// ListAlerts returns list of existing alerts.
func (s *AlertsService) ListAlerts(ctx context.Context, req *iav1beta1.ListAlertsRequest) (*iav1beta1.ListAlertsResponse, error) {
	var pageIndex int
	var pageSize int
	if req.PageParams != nil {
		pageIndex = int(req.PageParams.Index)
		pageSize = int(req.PageParams.PageSize)
	}
	var err error

	skip := pageIndex * pageSize

	alerts, err := s.alertManager.GetAlerts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get alerts form alertmanager")
	}

	pageTotals := &iav1beta1.PageTotals{
		TotalPages: 1,
	}

	res := make([]*iav1beta1.Alert, 0, len(alerts))
	for _, alert := range alerts {

		if _, ok := alert.Labels["ia"]; !ok { // Skip non-IA alerts
			continue
		}

		updatedAt, err := ptypes.TimestampProto(time.Time(*alert.UpdatedAt))
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert timestamp")
		}

		createdAt, err := ptypes.TimestampProto(time.Time(*alert.StartsAt))
		if err != nil {
			return nil, errors.Wrap(err, "failed to convert timestamp")
		}

		st := iav1beta1.Status_STATUS_INVALID
		if *alert.Status.State == "active" {
			st = iav1beta1.Status_TRIGGERING
		}

		if len(alert.Status.SilencedBy) != 0 {
			st = iav1beta1.Status_SILENCED
		}

		var rule *iav1beta1.Rule
		// Rules files created by user in directory /srv/prometheus/rules/ doesn't have associated rules in DB.
		// So alertname field will be empty or will keep invalid value. Don't fill rule field in that case.
		ruleID, ok := alert.Labels["alertname"]
		if ok && strings.HasPrefix(ruleID, "/rule_id/") {
			var r *models.Rule
			var channels []*models.Channel
			e := s.db.InTransaction(func(tx *reform.TX) error {
				var err error
				r, err = models.FindRuleByID(tx.Querier, ruleID)
				if err != nil {
					return err
				}

				channels, err = models.FindChannelsByIDs(tx.Querier, r.ChannelIDs)
				return err
			})
			if e != nil {
				// The codes.NotFound code can be returned just only by the FindRulesByID func
				// from the transaction above.
				if st, ok := status.FromError(e); ok && st.Code() == codes.NotFound {
					s.l.Warnf("The related alert rule was most likely removed: %s", st.Message())
					continue
				}

				return nil, e
			}

			template, ok := s.templatesService.getTemplates()[r.TemplateName]
			if !ok {
				s.l.Warnf("Failed to find template with name: %s", r.TemplateName)
				continue
			}

			rule, err = convertRule(s.l, r, template, channels)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to convert alert rule")
			}
		}
		pass, err := satisfiesFilters(alert, rule.Filters)
		if err != nil {
			return nil, err
		}

		if !pass {
			continue
		}

		pageTotals.TotalItems++
		if skip > 0 {
			skip--
			continue
		}

		if pageSize > 0 && len(res) >= pageSize {
			continue
		}

		res = append(res, &iav1beta1.Alert{
			AlertId:   getAlertID(alert),
			Summary:   alert.Annotations["summary"],
			Severity:  managementpb.Severity(common.ParseSeverity(alert.Labels["severity"])),
			Status:    st,
			Labels:    alert.Labels,
			Rule:      rule,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})

	}

	if pageSize > 0 {
		pageTotals.TotalPages = int32(len(res) / pageSize)
		if len(res)%pageSize > 0 {
			pageTotals.TotalPages++
		}
	}

	return &iav1beta1.ListAlertsResponse{Alerts: res, Totals: pageTotals}, nil
}

// satisfiesFilters checks that alert passes filters, returns true in case of success.
func satisfiesFilters(alert *ammodels.GettableAlert, filters []*iav1beta1.Filter) (bool, error) {
	for _, filter := range filters {
		value, ok := alert.Labels[filter.Key]
		if !ok {
			return false, nil
		}

		switch filter.Type {
		case iav1beta1.FilterType_EQUAL:
			if filter.Value != value {
				return false, nil
			}
		case iav1beta1.FilterType_REGEX:
			match, err := regexp.Match(filter.Value, []byte(value))
			if err != nil {
				return false, status.Errorf(codes.InvalidArgument, "bad regular expression: +%v", err)
			}

			if !match {
				return false, nil
			}
		case iav1beta1.FilterType_FILTER_TYPE_INVALID:
			fallthrough
		default:
			return false, status.Error(codes.Internal, "Unexpected filter type.")
		}
	}

	return true, nil
}

func getAlertID(alert *ammodels.GettableAlert) string {
	return *alert.Fingerprint
}

// ToggleAlert allows to silence/unsilence specified alerts.
func (s *AlertsService) ToggleAlert(ctx context.Context, req *iav1beta1.ToggleAlertRequest) (*iav1beta1.ToggleAlertResponse, error) {
	switch req.Silenced {
	case iav1beta1.BooleanFlag_DO_NOT_CHANGE:
		// nothing
	case iav1beta1.BooleanFlag_TRUE:
		err := s.alertManager.Silence(ctx, req.AlertId)
		if err != nil {
			return nil, err
		}
	case iav1beta1.BooleanFlag_FALSE:
		err := s.alertManager.Unsilence(ctx, req.AlertId)
		if err != nil {
			return nil, err
		}
	}

	return &iav1beta1.ToggleAlertResponse{}, nil
}

// Check interfaces.
var (
	_ iav1beta1.AlertsServer = (*AlertsService)(nil)
)
