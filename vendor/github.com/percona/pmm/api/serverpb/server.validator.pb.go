// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: serverpb/server.proto

package serverpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *VersionRequest) Validate() error {
	return nil
}
func (this *VersionResponse) Validate() error {
	if this.Managed != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Managed); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Managed", err)
		}
	}
	return nil
}
func (this *VersionResponse_Managed) Validate() error {
	if this.Timestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Timestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Timestamp", err)
		}
	}
	return nil
}
func (this *ReadinessRequest) Validate() error {
	return nil
}
func (this *ReadinessResponse) Validate() error {
	return nil
}
func (this *CheckUpdatesRequest) Validate() error {
	return nil
}
func (this *CheckUpdatesResponse) Validate() error {
	if this.InstalledTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.InstalledTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("InstalledTimestamp", err)
		}
	}
	if this.LatestTimestamp != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LatestTimestamp); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LatestTimestamp", err)
		}
	}
	return nil
}
func (this *StartUpdateRequest) Validate() error {
	return nil
}
func (this *StartUpdateResponse) Validate() error {
	return nil
}
func (this *UpdateStatusRequest) Validate() error {
	return nil
}
func (this *UpdateStatusResponse) Validate() error {
	return nil
}
func (this *MetricsResolutions) Validate() error {
	if this.Hr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Hr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Hr", err)
		}
	}
	if this.Mr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Mr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Mr", err)
		}
	}
	if this.Lr != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Lr); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Lr", err)
		}
	}
	return nil
}
func (this *Settings) Validate() error {
	if this.MetricsResolutions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MetricsResolutions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MetricsResolutions", err)
		}
	}
	return nil
}
func (this *GetSettingsRequest) Validate() error {
	return nil
}
func (this *GetSettingsResponse) Validate() error {
	if this.Settings != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Settings); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Settings", err)
		}
	}
	return nil
}
func (this *ChangeSettingsRequest) Validate() error {
	if this.MetricsResolutions != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MetricsResolutions); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MetricsResolutions", err)
		}
	}
	return nil
}
func (this *ChangeSettingsResponse) Validate() error {
	if this.Settings != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Settings); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Settings", err)
		}
	}
	return nil
}
