// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/ia/channels.proto

package iav1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *BasicAuth) Validate() error {
	return nil
}
func (this *TLSConfig) Validate() error {
	return nil
}
func (this *HTTPConfig) Validate() error {
	if this.BasicAuth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.BasicAuth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("BasicAuth", err)
		}
	}
	if this.TlsConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.TlsConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("TlsConfig", err)
		}
	}
	return nil
}
func (this *EmailConfig) Validate() error {
	if len(this.To) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("To", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.To))
	}
	return nil
}
func (this *PagerDutyConfig) Validate() error {
	return nil
}
func (this *SlackConfig) Validate() error {
	if this.Channel == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Channel", fmt.Errorf(`value '%v' must not be an empty string`, this.Channel))
	}
	return nil
}
func (this *WebhookConfig) Validate() error {
	if this.Url == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Url", fmt.Errorf(`value '%v' must not be an empty string`, this.Url))
	}
	if this.HttpConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.HttpConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("HttpConfig", err)
		}
	}
	return nil
}
func (this *Channel) Validate() error {
	if oneOfNester, ok := this.GetChannel().(*Channel_EmailConfig); ok {
		if oneOfNester.EmailConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.EmailConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("EmailConfig", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChannel().(*Channel_PagerdutyConfig); ok {
		if oneOfNester.PagerdutyConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PagerdutyConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PagerdutyConfig", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChannel().(*Channel_SlackConfig); ok {
		if oneOfNester.SlackConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SlackConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SlackConfig", err)
			}
		}
	}
	if oneOfNester, ok := this.GetChannel().(*Channel_WebhookConfig); ok {
		if oneOfNester.WebhookConfig != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.WebhookConfig); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("WebhookConfig", err)
			}
		}
	}
	return nil
}
func (this *ListChannelsRequest) Validate() error {
	return nil
}
func (this *ListChannelsResponse) Validate() error {
	for _, item := range this.Channels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Channels", err)
			}
		}
	}
	return nil
}
func (this *AddChannelRequest) Validate() error {
	if this.EmailConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EmailConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EmailConfig", err)
		}
	}
	if this.PagerdutyConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PagerdutyConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PagerdutyConfig", err)
		}
	}
	if this.SlackConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SlackConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SlackConfig", err)
		}
	}
	if this.WebhookConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WebhookConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WebhookConfig", err)
		}
	}
	return nil
}
func (this *AddChannelResponse) Validate() error {
	return nil
}
func (this *ChangeChannelRequest) Validate() error {
	if this.ChannelId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ChannelId", fmt.Errorf(`value '%v' must not be an empty string`, this.ChannelId))
	}
	if this.EmailConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.EmailConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("EmailConfig", err)
		}
	}
	if this.PagerdutyConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PagerdutyConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PagerdutyConfig", err)
		}
	}
	if this.SlackConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SlackConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SlackConfig", err)
		}
	}
	if this.WebhookConfig != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.WebhookConfig); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("WebhookConfig", err)
		}
	}
	return nil
}
func (this *ChangeChannelResponse) Validate() error {
	return nil
}
func (this *RemoveChannelRequest) Validate() error {
	if this.ChannelId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ChannelId", fmt.Errorf(`value '%v' must not be an empty string`, this.ChannelId))
	}
	return nil
}
func (this *RemoveChannelResponse) Validate() error {
	return nil
}
