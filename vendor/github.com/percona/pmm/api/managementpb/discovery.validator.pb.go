// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/discovery.proto

package managementpb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *DiscoverRDSRequest) Validate() error {
	return nil
}
func (this *DiscoverRDSInstance) Validate() error {
	return nil
}
func (this *DiscoverRDSResponse) Validate() error {
	for _, item := range this.RdsInstances {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RdsInstances", err)
			}
		}
	}
	return nil
}
