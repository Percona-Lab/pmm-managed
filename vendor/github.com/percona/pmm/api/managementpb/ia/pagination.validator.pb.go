// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: managementpb/ia/pagination.proto

package iav1beta1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PageParams) Validate() error {
	if !(this.PageSize > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("PageSize", fmt.Errorf(`value '%v' must be greater than '-1'`, this.PageSize))
	}
	if !(this.Index > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("Index", fmt.Errorf(`value '%v' must be greater than '-1'`, this.Index))
	}
	return nil
}
func (this *PageTotals) Validate() error {
	return nil
}
