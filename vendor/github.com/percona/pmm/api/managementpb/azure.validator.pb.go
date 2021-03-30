// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/azure.proto

package managementpb

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

func (this *DiscoverAzureDatabaseRequest) Validate() error {
	if this.AzureClientId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureClientId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureClientId))
	}
	if this.AzureClientSecret == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureClientSecret", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureClientSecret))
	}
	if this.AzureTenantId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureTenantId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureTenantId))
	}
	if this.AzureSubscriptionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureSubscriptionId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureSubscriptionId))
	}
	return nil
}
func (this *DiscoverAzureDatabaseInstance) Validate() error {
	return nil
}
func (this *DiscoverAzureDatabaseResponse) Validate() error {
	for _, item := range this.AzureDatabaseInstance {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AzureDatabaseInstance", err)
			}
		}
	}
	return nil
}
func (this *AddAzureDatabaseRequest) Validate() error {
	if this.Region == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Region", fmt.Errorf(`value '%v' must not be an empty string`, this.Region))
	}
	if this.InstanceId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("InstanceId", fmt.Errorf(`value '%v' must not be an empty string`, this.InstanceId))
	}
	if this.Address == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Address", fmt.Errorf(`value '%v' must not be an empty string`, this.Address))
	}
	if !(this.Port > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Port", fmt.Errorf(`value '%v' must be greater than '0'`, this.Port))
	}
	if this.Username == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Username", fmt.Errorf(`value '%v' must not be an empty string`, this.Username))
	}
	if this.AzureClientId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureClientId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureClientId))
	}
	if this.AzureClientSecret == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureClientSecret", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureClientSecret))
	}
	if this.AzureTenantId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureTenantId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureTenantId))
	}
	if this.AzureSubscriptionId == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureSubscriptionId", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureSubscriptionId))
	}
	if this.AzureResourceGroup == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("AzureResourceGroup", fmt.Errorf(`value '%v' must not be an empty string`, this.AzureResourceGroup))
	}
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *AddAzureDatabaseResponse) Validate() error {
	return nil
}