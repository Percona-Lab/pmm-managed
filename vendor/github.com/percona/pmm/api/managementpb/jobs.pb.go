// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/jobs.proto

package managementpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// JobType represents Job type.
type JobType int32

const (
	JobType_JOB_TYPE_INVALID JobType = 0
	JobType_ECHO             JobType = 1
)

// Enum value maps for JobType.
var (
	JobType_name = map[int32]string{
		0: "JOB_TYPE_INVALID",
		1: "ECHO",
	}
	JobType_value = map[string]int32{
		"JOB_TYPE_INVALID": 0,
		"ECHO":             1,
	}
)

func (x JobType) Enum() *JobType {
	p := new(JobType)
	*p = x
	return p
}

func (x JobType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JobType) Descriptor() protoreflect.EnumDescriptor {
	return file_managementpb_jobs_proto_enumTypes[0].Descriptor()
}

func (JobType) Type() protoreflect.EnumType {
	return &file_managementpb_jobs_proto_enumTypes[0]
}

func (x JobType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JobType.Descriptor instead.
func (JobType) EnumDescriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{0}
}

type GetJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Job ID.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *GetJobRequest) Reset() {
	*x = GetJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobRequest) ProtoMessage() {}

func (x *GetJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobRequest.ProtoReflect.Descriptor instead.
func (*GetJobRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{0}
}

func (x *GetJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type GetJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetJobResponse) Reset() {
	*x = GetJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetJobResponse) ProtoMessage() {}

func (x *GetJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetJobResponse.ProtoReflect.Descriptor instead.
func (*GetJobResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{1}
}

type StartEchoJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pmm-agent ID where to run this Action.
	PmmAgentId string `protobuf:"bytes,1,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Service ID for this Action. Required.
	ServiceId string             `protobuf:"bytes,2,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	Message   string             `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Delay     *duration.Duration `protobuf:"bytes,4,opt,name=delay,proto3" json:"delay,omitempty"`
}

func (x *StartEchoJobRequest) Reset() {
	*x = StartEchoJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEchoJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEchoJobRequest) ProtoMessage() {}

func (x *StartEchoJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEchoJobRequest.ProtoReflect.Descriptor instead.
func (*StartEchoJobRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{2}
}

func (x *StartEchoJobRequest) GetPmmAgentId() string {
	if x != nil {
		return x.PmmAgentId
	}
	return ""
}

func (x *StartEchoJobRequest) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *StartEchoJobRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StartEchoJobRequest) GetDelay() *duration.Duration {
	if x != nil {
		return x.Delay
	}
	return nil
}

type StartEchoJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Job ID.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// pmm-agent ID where to this Action was started.
	PmmAgentId string `protobuf:"bytes,2,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
}

func (x *StartEchoJobResponse) Reset() {
	*x = StartEchoJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartEchoJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartEchoJobResponse) ProtoMessage() {}

func (x *StartEchoJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartEchoJobResponse.ProtoReflect.Descriptor instead.
func (*StartEchoJobResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{3}
}

func (x *StartEchoJobResponse) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *StartEchoJobResponse) GetPmmAgentId() string {
	if x != nil {
		return x.PmmAgentId
	}
	return ""
}

type CancelJobRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique Job ID. Required.
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
}

func (x *CancelJobRequest) Reset() {
	*x = CancelJobRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobRequest) ProtoMessage() {}

func (x *CancelJobRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobRequest.ProtoReflect.Descriptor instead.
func (*CancelJobRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{4}
}

func (x *CancelJobRequest) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

type CancelJobResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CancelJobResponse) Reset() {
	*x = CancelJobResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_jobs_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CancelJobResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CancelJobResponse) ProtoMessage() {}

func (x *CancelJobResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_jobs_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CancelJobResponse.ProtoReflect.Descriptor instead.
func (*CancelJobResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_jobs_proto_rawDescGZIP(), []int{5}
}

var File_managementpb_jobs_proto protoreflect.FileDescriptor

var file_managementpb_jobs_proto_rawDesc = []byte{
	0x0a, 0x17, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x6a,
	0x6f, 0x62, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x06,
	0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf,
	0x1f, 0x02, 0x58, 0x01, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa9, 0x01,
	0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x6f, 0x62, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x6d, 0x6d, 0x5f, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6d, 0x6d,
	0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x05, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61, 0x79, 0x22, 0x4f, 0x0a, 0x14, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x6d, 0x6d, 0x5f,
	0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x70, 0x6d, 0x6d, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x43, 0x61,
	0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x22, 0x13, 0x0a,
	0x11, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2a, 0x29, 0x0a, 0x07, 0x4a, 0x6f, 0x62, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a,
	0x10, 0x4a, 0x4f, 0x42, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x45, 0x43, 0x48, 0x4f, 0x10, 0x01, 0x32, 0xdf, 0x02,
	0x0a, 0x04, 0x4a, 0x6f, 0x62, 0x73, 0x12, 0x66, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x41, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x19, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x1c, 0x22, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x47, 0x65, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x7b,
	0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x6f, 0x62, 0x12, 0x1f,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x20, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x4a, 0x6f, 0x62, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x45, 0x63, 0x68, 0x6f, 0x3a, 0x01, 0x2a, 0x12, 0x72, 0x0a, 0x0c, 0x43,
	0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a,
	0x6f, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x4a, 0x6f, 0x62,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f,
	0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2f, 0x4a, 0x6f, 0x62, 0x73, 0x2f, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x3a, 0x01, 0x2a, 0x42,
	0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_jobs_proto_rawDescOnce sync.Once
	file_managementpb_jobs_proto_rawDescData = file_managementpb_jobs_proto_rawDesc
)

func file_managementpb_jobs_proto_rawDescGZIP() []byte {
	file_managementpb_jobs_proto_rawDescOnce.Do(func() {
		file_managementpb_jobs_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_jobs_proto_rawDescData)
	})
	return file_managementpb_jobs_proto_rawDescData
}

var file_managementpb_jobs_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_managementpb_jobs_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_managementpb_jobs_proto_goTypes = []interface{}{
	(JobType)(0),                 // 0: management.JobType
	(*GetJobRequest)(nil),        // 1: management.GetJobRequest
	(*GetJobResponse)(nil),       // 2: management.GetJobResponse
	(*StartEchoJobRequest)(nil),  // 3: management.StartEchoJobRequest
	(*StartEchoJobResponse)(nil), // 4: management.StartEchoJobResponse
	(*CancelJobRequest)(nil),     // 5: management.CancelJobRequest
	(*CancelJobResponse)(nil),    // 6: management.CancelJobResponse
	(*duration.Duration)(nil),    // 7: google.protobuf.Duration
}
var file_managementpb_jobs_proto_depIdxs = []int32{
	7, // 0: management.StartEchoJobRequest.delay:type_name -> google.protobuf.Duration
	1, // 1: management.Jobs.GetAction:input_type -> management.GetJobRequest
	3, // 2: management.Jobs.StartEchoJob:input_type -> management.StartEchoJobRequest
	5, // 3: management.Jobs.CancelAction:input_type -> management.CancelJobRequest
	2, // 4: management.Jobs.GetAction:output_type -> management.GetJobResponse
	4, // 5: management.Jobs.StartEchoJob:output_type -> management.StartEchoJobResponse
	6, // 6: management.Jobs.CancelAction:output_type -> management.CancelJobResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_managementpb_jobs_proto_init() }
func file_managementpb_jobs_proto_init() {
	if File_managementpb_jobs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_jobs_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_jobs_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_jobs_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEchoJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_jobs_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartEchoJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_jobs_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelJobRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_managementpb_jobs_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CancelJobResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_jobs_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_jobs_proto_goTypes,
		DependencyIndexes: file_managementpb_jobs_proto_depIdxs,
		EnumInfos:         file_managementpb_jobs_proto_enumTypes,
		MessageInfos:      file_managementpb_jobs_proto_msgTypes,
	}.Build()
	File_managementpb_jobs_proto = out.File
	file_managementpb_jobs_proto_rawDesc = nil
	file_managementpb_jobs_proto_goTypes = nil
	file_managementpb_jobs_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// JobsClient is the client API for Jobs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type JobsClient interface {
	// GetAction gets an result of given Action.
	GetAction(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error)
	// StartEchoJob starts echo job.
	StartEchoJob(ctx context.Context, in *StartEchoJobRequest, opts ...grpc.CallOption) (*StartEchoJobResponse, error)
	// CancelAction stops a Job.
	CancelAction(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error)
}

type jobsClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsClient(cc grpc.ClientConnInterface) JobsClient {
	return &jobsClient{cc}
}

func (c *jobsClient) GetAction(ctx context.Context, in *GetJobRequest, opts ...grpc.CallOption) (*GetJobResponse, error) {
	out := new(GetJobResponse)
	err := c.cc.Invoke(ctx, "/management.Jobs/GetAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) StartEchoJob(ctx context.Context, in *StartEchoJobRequest, opts ...grpc.CallOption) (*StartEchoJobResponse, error) {
	out := new(StartEchoJobResponse)
	err := c.cc.Invoke(ctx, "/management.Jobs/StartEchoJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsClient) CancelAction(ctx context.Context, in *CancelJobRequest, opts ...grpc.CallOption) (*CancelJobResponse, error) {
	out := new(CancelJobResponse)
	err := c.cc.Invoke(ctx, "/management.Jobs/CancelAction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServer is the server API for Jobs service.
type JobsServer interface {
	// GetAction gets an result of given Action.
	GetAction(context.Context, *GetJobRequest) (*GetJobResponse, error)
	// StartEchoJob starts echo job.
	StartEchoJob(context.Context, *StartEchoJobRequest) (*StartEchoJobResponse, error)
	// CancelAction stops a Job.
	CancelAction(context.Context, *CancelJobRequest) (*CancelJobResponse, error)
}

// UnimplementedJobsServer can be embedded to have forward compatible implementations.
type UnimplementedJobsServer struct {
}

func (*UnimplementedJobsServer) GetAction(context.Context, *GetJobRequest) (*GetJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAction not implemented")
}
func (*UnimplementedJobsServer) StartEchoJob(context.Context, *StartEchoJobRequest) (*StartEchoJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartEchoJob not implemented")
}
func (*UnimplementedJobsServer) CancelAction(context.Context, *CancelJobRequest) (*CancelJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelAction not implemented")
}

func RegisterJobsServer(s *grpc.Server, srv JobsServer) {
	s.RegisterService(&_Jobs_serviceDesc, srv)
}

func _Jobs_GetAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).GetAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Jobs/GetAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).GetAction(ctx, req.(*GetJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_StartEchoJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartEchoJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).StartEchoJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Jobs/StartEchoJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).StartEchoJob(ctx, req.(*StartEchoJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jobs_CancelAction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServer).CancelAction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.Jobs/CancelAction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServer).CancelAction(ctx, req.(*CancelJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Jobs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.Jobs",
	HandlerType: (*JobsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAction",
			Handler:    _Jobs_GetAction_Handler,
		},
		{
			MethodName: "StartEchoJob",
			Handler:    _Jobs_StartEchoJob_Handler,
		},
		{
			MethodName: "CancelAction",
			Handler:    _Jobs_CancelAction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/jobs.proto",
}
