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
	"context"

	iav1beta1 "github.com/percona/pmm/api/managementpb/ia"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/reform.v1"

	"github.com/percona/pmm-managed/models"
)

// ChannelsService represents integrated alerting channels API.
type ChannelsService struct {
	db *reform.DB
}

// NewChannelsService creates new channels API service.
func NewChannelsService(db *reform.DB) *ChannelsService {
	return &ChannelsService{
		db: db,
	}
}

// ListChannels returns list of available channels.
func (s *ChannelsService) ListChannels(ctx context.Context, request *iav1beta1.ListChannelsRequest) (*iav1beta1.ListChannelsResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
	}

	var channels []models.Channel
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		channels, err = models.FindChannels(tx.Querier)
		return err
	})
	if e != nil {
		return nil, e
	}

	res := make([]*iav1beta1.Channel, len(channels))
	for i, channel := range channels {
		c := &iav1beta1.Channel{
			ChannelId: channel.ID,
			Summary:   channel.Summary,
			Disabled:  channel.Disabled,
		}

		switch channel.Type {
		case models.Email:
			config := channel.EmailConfig
			c.Channel = &iav1beta1.Channel_EmailConfig{
				EmailConfig: &iav1beta1.EmailConfig{
					SendResolved: config.SendResolved,
					To:           config.To,
				},
			}
		case models.PagerDuty:
			config := channel.PagerDutyConfig
			c.Channel = &iav1beta1.Channel_PagerdutyConfig{
				PagerdutyConfig: &iav1beta1.PagerDutyConfig{
					SendResolved: config.SendResolved,
					RoutingKey:   config.RoutingKey,
					ServiceKey:   config.ServiceKey,
				},
			}
		case models.Slack:
			config := channel.SlackConfig
			c.Channel = &iav1beta1.Channel_SlackConfig{
				SlackConfig: &iav1beta1.SlackConfig{
					SendResolved: config.SendResolved,
					Channel:      config.Channel,
				},
			}
		case models.WebHook:
			config := channel.WebHookConfig
			c.Channel = &iav1beta1.Channel_WebhookConfig{
				WebhookConfig: &iav1beta1.WebhookConfig{
					SendResolved: config.SendResolved,
					Url:          config.URL,
					HttpConfig:   convertModelToHTTPConfig(config.HTTPConfig),
					MaxAlerts:    config.MaxAlerts,
				},
			}
		default:
			return nil, errors.Errorf("Unknown notification channel type %s", channel.Type)
		}

		res[i] = c
	}

	return &iav1beta1.ListChannelsResponse{Channels: res}, nil
}

// AddChannel adds new notification channel.
func (s *ChannelsService) AddChannel(ctx context.Context, req *iav1beta1.AddChannelRequest) (*iav1beta1.AddChannelResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
	}

	params := &models.CreateChannelParams{
		Summary:  req.Summary,
		Disabled: req.Disabled,
	}

	if req.EmailConfig != nil {
		params.EmailConfig = &models.EmailConfig{
			SendResolved: req.EmailConfig.SendResolved,
			To:           req.EmailConfig.To,
		}
	}
	if req.PagerdutyConfig != nil {
		params.PagerDutyConfig = &models.PagerDutyConfig{
			SendResolved: req.PagerdutyConfig.SendResolved,
			RoutingKey:   req.PagerdutyConfig.RoutingKey,
			ServiceKey:   req.PagerdutyConfig.ServiceKey,
		}
	}
	if req.SlackConfig != nil {
		params.SlackConfig = &models.SlackConfig{
			SendResolved: req.SlackConfig.SendResolved,
			Channel:      req.SlackConfig.Channel,
		}
	}
	if req.WebhookConfig != nil {
		params.WebHookConfig = &models.WebHookConfig{
			SendResolved: req.WebhookConfig.SendResolved,
			URL:          req.WebhookConfig.Url,
			MaxAlerts:    req.WebhookConfig.MaxAlerts,
			HTTPConfig:   convertHTTPConfigToModel(req.WebhookConfig.HttpConfig),
		}
	}

	var channel *models.Channel
	e := s.db.InTransaction(func(tx *reform.TX) error {
		var err error
		channel, err = models.CreateChannel(tx.Querier, params)
		return err
	})
	if e != nil {
		return nil, e
	}
	return &iav1beta1.AddChannelResponse{ChannelId: channel.ID}, nil
}

// ChangeChannel changes existing notification channel.
func (s *ChannelsService) ChangeChannel(ctx context.Context, req *iav1beta1.ChangeChannelRequest) (*iav1beta1.ChangeChannelResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
	}

	params := &models.ChangeChannelParams{
		Disabled: req.Disabled,
	}

	if c := req.EmailConfig; c != nil {
		params.EmailConfig = &models.EmailConfig{
			SendResolved: c.SendResolved,
			To:           c.To,
		}
	}
	if c := req.PagerdutyConfig; c != nil {
		params.PagerDutyConfig = &models.PagerDutyConfig{
			SendResolved: c.SendResolved,
			RoutingKey:   c.RoutingKey,
			ServiceKey:   c.ServiceKey,
		}
	}
	if c := req.SlackConfig; c != nil {
		params.SlackConfig = &models.SlackConfig{
			SendResolved: c.SendResolved,
			Channel:      c.Channel,
		}
	}
	if c := req.WebhookConfig; c != nil {
		params.WebHookConfig = &models.WebHookConfig{
			SendResolved: c.SendResolved,
			URL:          c.Url,
			MaxAlerts:    c.MaxAlerts,
			HTTPConfig:   convertHTTPConfigToModel(c.HttpConfig),
		}
	}

	e := s.db.InTransaction(func(tx *reform.TX) error {
		_, err := models.ChangeChannel(tx.Querier, req.ChannelId, params)
		return err
	})
	if e != nil {
		return nil, e
	}
	return &iav1beta1.ChangeChannelResponse{}, nil
}

// RemoveChannel removes notification channel.
func (s *ChannelsService) RemoveChannel(ctx context.Context, req *iav1beta1.RemoveChannelRequest) (*iav1beta1.RemoveChannelResponse, error) {
	settings, err := models.GetSettings(s.db)
	if err != nil {
		return nil, err
	}

	if !settings.IntegratedAlerting.Enabled {
		return nil, status.Errorf(codes.FailedPrecondition, "%v.", err)
	}

	e := s.db.InTransaction(func(tx *reform.TX) error {
		return models.RemoveChannel(tx.Querier, req.ChannelId)
	})
	if e != nil {
		return nil, e
	}
	return &iav1beta1.RemoveChannelResponse{}, nil
}

func convertHTTPConfigToModel(config *iav1beta1.HTTPConfig) *models.HTTPConfig {
	if config == nil {
		return nil
	}

	res := &models.HTTPConfig{
		BearerToken:     config.BearerToken,
		BearerTokenFile: config.BearerTokenFile,
		ProxyURL:        config.ProxyUrl,
	}

	if basicAuthConf := config.BasicAuth; basicAuthConf != nil {
		res.BasicAuth = &models.HTTPBasicAuth{
			Username:     basicAuthConf.Username,
			Password:     basicAuthConf.Password,
			PasswordFile: basicAuthConf.PasswordFile,
		}
	}

	if tlsConfig := config.TlsConfig; tlsConfig != nil {
		res.TLSConfig = &models.TLSConfig{
			CaFile:             tlsConfig.CaFile,
			CertFile:           tlsConfig.CertFile,
			KeyFile:            tlsConfig.KeyFile,
			ServerName:         tlsConfig.ServerName,
			InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
		}
	}

	return res
}

func convertModelToHTTPConfig(config *models.HTTPConfig) *iav1beta1.HTTPConfig {
	if config == nil {
		return nil
	}

	res := &iav1beta1.HTTPConfig{
		BearerToken:     config.BearerToken,
		BearerTokenFile: config.BearerTokenFile,
		ProxyUrl:        config.ProxyURL,
	}

	if basicAuthConf := config.BasicAuth; basicAuthConf != nil {
		res.BasicAuth = &iav1beta1.BasicAuth{
			Username:     basicAuthConf.Username,
			Password:     basicAuthConf.Password,
			PasswordFile: basicAuthConf.PasswordFile,
		}
	}

	if tlsConfig := config.TLSConfig; tlsConfig != nil {
		res.TlsConfig = &iav1beta1.TLSConfig{
			CaFile:             tlsConfig.CaFile,
			CertFile:           tlsConfig.CertFile,
			KeyFile:            tlsConfig.KeyFile,
			ServerName:         tlsConfig.ServerName,
			InsecureSkipVerify: tlsConfig.InsecureSkipVerify,
		}
	}

	return res
}

// Check interfaces.
var (
	_ iav1beta1.ChannelsServer = (*ChannelsService)(nil)
)
