// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/backup/locations.proto

package backupv1beta1

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

// PMMServerLocationConfig represents file system config inside pmm-server.
type PMMServerLocationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PMMServerLocationConfig) Reset() {
	*x = PMMServerLocationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMMServerLocationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMMServerLocationConfig) ProtoMessage() {}

func (x *PMMServerLocationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMMServerLocationConfig.ProtoReflect.Descriptor instead.
func (*PMMServerLocationConfig) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{0}
}

func (x *PMMServerLocationConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// PMMClientLocationConfig represents file system config inside pmm-client.
type PMMClientLocationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *PMMClientLocationConfig) Reset() {
	*x = PMMClientLocationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PMMClientLocationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PMMClientLocationConfig) ProtoMessage() {}

func (x *PMMClientLocationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PMMClientLocationConfig.ProtoReflect.Descriptor instead.
func (*PMMClientLocationConfig) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{1}
}

func (x *PMMClientLocationConfig) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// S3LocationConfig represents S3 bucket configuration.
type S3LocationConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endpoint  string `protobuf:"bytes,1,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	AccessKey string `protobuf:"bytes,2,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey string `protobuf:"bytes,3,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
}

func (x *S3LocationConfig) Reset() {
	*x = S3LocationConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S3LocationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S3LocationConfig) ProtoMessage() {}

func (x *S3LocationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S3LocationConfig.ProtoReflect.Descriptor instead.
func (*S3LocationConfig) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{2}
}

func (x *S3LocationConfig) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *S3LocationConfig) GetAccessKey() string {
	if x != nil {
		return x.AccessKey
	}
	return ""
}

func (x *S3LocationConfig) GetSecretKey() string {
	if x != nil {
		return x.SecretKey
	}
	return ""
}

// Location represents single Backup Location.
type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine-readable ID.
	LocationId string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
	// Location name
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Short description
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Types that are assignable to Config:
	//	*Location_PmmClientConfig
	//	*Location_PmmServerConfig
	//	*Location_S3Config
	Config isLocation_Config `protobuf_oneof:"config"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{3}
}

func (x *Location) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Location) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (m *Location) GetConfig() isLocation_Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (x *Location) GetPmmClientConfig() *PMMClientLocationConfig {
	if x, ok := x.GetConfig().(*Location_PmmClientConfig); ok {
		return x.PmmClientConfig
	}
	return nil
}

func (x *Location) GetPmmServerConfig() *PMMServerLocationConfig {
	if x, ok := x.GetConfig().(*Location_PmmServerConfig); ok {
		return x.PmmServerConfig
	}
	return nil
}

func (x *Location) GetS3Config() *S3LocationConfig {
	if x, ok := x.GetConfig().(*Location_S3Config); ok {
		return x.S3Config
	}
	return nil
}

type isLocation_Config interface {
	isLocation_Config()
}

type Location_PmmClientConfig struct {
	PmmClientConfig *PMMClientLocationConfig `protobuf:"bytes,4,opt,name=pmm_client_config,json=pmmClientConfig,proto3,oneof"`
}

type Location_PmmServerConfig struct {
	PmmServerConfig *PMMServerLocationConfig `protobuf:"bytes,5,opt,name=pmm_server_config,json=pmmServerConfig,proto3,oneof"`
}

type Location_S3Config struct {
	S3Config *S3LocationConfig `protobuf:"bytes,6,opt,name=s3_config,json=s3Config,proto3,oneof"`
}

func (*Location_PmmClientConfig) isLocation_Config() {}

func (*Location_PmmServerConfig) isLocation_Config() {}

func (*Location_S3Config) isLocation_Config() {}

type ListLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListLocationsRequest) Reset() {
	*x = ListLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsRequest) ProtoMessage() {}

func (x *ListLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{4}
}

type ListLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []*Location `protobuf:"bytes,1,rep,name=locations,proto3" json:"locations,omitempty"`
}

func (x *ListLocationsResponse) Reset() {
	*x = ListLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsResponse) ProtoMessage() {}

func (x *ListLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsResponse.ProtoReflect.Descriptor instead.
func (*ListLocationsResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{5}
}

func (x *ListLocationsResponse) GetLocations() []*Location {
	if x != nil {
		return x.Locations
	}
	return nil
}

type AddLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Location name
	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// PMM-client file system configuration. Exactly one config should be set.
	PmmClientConfig *PMMClientLocationConfig `protobuf:"bytes,3,opt,name=pmm_client_config,json=pmmClientConfig,proto3" json:"pmm_client_config,omitempty"`
	// PMM-server file system configuration. Exactly one config should be set.
	PmmServerConfig *PMMServerLocationConfig `protobuf:"bytes,4,opt,name=pmm_server_config,json=pmmServerConfig,proto3" json:"pmm_server_config,omitempty"`
	// S3 Bucket configuration. Exactly one config should be set.
	S3Config *S3LocationConfig `protobuf:"bytes,5,opt,name=s3_config,json=s3Config,proto3" json:"s3_config,omitempty"`
}

func (x *AddLocationRequest) Reset() {
	*x = AddLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLocationRequest) ProtoMessage() {}

func (x *AddLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLocationRequest.ProtoReflect.Descriptor instead.
func (*AddLocationRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{6}
}

func (x *AddLocationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddLocationRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddLocationRequest) GetPmmClientConfig() *PMMClientLocationConfig {
	if x != nil {
		return x.PmmClientConfig
	}
	return nil
}

func (x *AddLocationRequest) GetPmmServerConfig() *PMMServerLocationConfig {
	if x != nil {
		return x.PmmServerConfig
	}
	return nil
}

func (x *AddLocationRequest) GetS3Config() *S3LocationConfig {
	if x != nil {
		return x.S3Config
	}
	return nil
}

type AddLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Machine-readable ID.
	LocationId string `protobuf:"bytes,1,opt,name=location_id,json=locationId,proto3" json:"location_id,omitempty"`
}

func (x *AddLocationResponse) Reset() {
	*x = AddLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_backup_locations_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddLocationResponse) ProtoMessage() {}

func (x *AddLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_backup_locations_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddLocationResponse.ProtoReflect.Descriptor instead.
func (*AddLocationResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_backup_locations_proto_rawDescGZIP(), []int{7}
}

func (x *AddLocationResponse) GetLocationId() string {
	if x != nil {
		return x.LocationId
	}
	return ""
}

var File_managementpb_backup_locations_proto protoreflect.FileDescriptor

var file_managementpb_backup_locations_proto_rawDesc = []byte{
	0x0a, 0x23, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x35, 0x0a, 0x17, 0x50,
	0x4d, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0x35, 0x0a, 0x17, 0x50, 0x4d, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1a, 0x0a,
	0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x84, 0x01, 0x0a, 0x10, 0x53, 0x33,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22,
	0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69,
	0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x25, 0x0a, 0x0a, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2,
	0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x09, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79,
	0x22, 0xda, 0x02, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a,
	0x0b, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x55, 0x0a, 0x11, 0x70, 0x6d, 0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x50, 0x4d, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x6d, 0x6d, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x55, 0x0a, 0x11, 0x70,
	0x6d, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x4d, 0x4d, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48,
	0x00, 0x52, 0x0f, 0x70, 0x6d, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3f, 0x0a, 0x09, 0x73, 0x33, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x33, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x08, 0x73, 0x33, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x08, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x16, 0x0a,
	0x14, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x4f, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36,
	0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x18, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x09, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xbb, 0x02, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xe2, 0xdf, 0x1f,
	0x02, 0x58, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x53, 0x0a, 0x11, 0x70,
	0x6d, 0x6d, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x4d, 0x4d, 0x43, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x0f, 0x70, 0x6d, 0x6d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x53, 0x0a, 0x11, 0x70, 0x6d, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x62, 0x61,
	0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x50, 0x4d, 0x4d,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0f, 0x70, 0x6d, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x3d, 0x0a, 0x09, 0x73, 0x33, 0x5f, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75,
	0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x33, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x73, 0x33, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x22, 0x36, 0x0a, 0x13, 0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x32, 0xa4, 0x02, 0x0a,
	0x09, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x8d, 0x01, 0x0a, 0x0d, 0x4c,
	0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x24, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x2f, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x29, 0x22, 0x24, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x86, 0x01, 0x0a, 0x0b, 0x41,
	0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x62, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x4c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x41, 0x64, 0x64, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x2e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x22, 0x23, 0x2f, 0x76, 0x31,
	0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x2f, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x41, 0x64, 0x64,
	0x3a, 0x01, 0x2a, 0x42, 0x27, 0x5a, 0x25, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x3b, 0x62,
	0x61, 0x63, 0x6b, 0x75, 0x70, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_backup_locations_proto_rawDescOnce sync.Once
	file_managementpb_backup_locations_proto_rawDescData = file_managementpb_backup_locations_proto_rawDesc
)

func file_managementpb_backup_locations_proto_rawDescGZIP() []byte {
	file_managementpb_backup_locations_proto_rawDescOnce.Do(func() {
		file_managementpb_backup_locations_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_backup_locations_proto_rawDescData)
	})
	return file_managementpb_backup_locations_proto_rawDescData
}

var file_managementpb_backup_locations_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_managementpb_backup_locations_proto_goTypes = []interface{}{
	(*PMMServerLocationConfig)(nil), // 0: backup.v1beta1.PMMServerLocationConfig
	(*PMMClientLocationConfig)(nil), // 1: backup.v1beta1.PMMClientLocationConfig
	(*S3LocationConfig)(nil),        // 2: backup.v1beta1.S3LocationConfig
	(*Location)(nil),                // 3: backup.v1beta1.Location
	(*ListLocationsRequest)(nil),    // 4: backup.v1beta1.ListLocationsRequest
	(*ListLocationsResponse)(nil),   // 5: backup.v1beta1.ListLocationsResponse
	(*AddLocationRequest)(nil),      // 6: backup.v1beta1.AddLocationRequest
	(*AddLocationResponse)(nil),     // 7: backup.v1beta1.AddLocationResponse
}
var file_managementpb_backup_locations_proto_depIdxs = []int32{
	1, // 0: backup.v1beta1.Location.pmm_client_config:type_name -> backup.v1beta1.PMMClientLocationConfig
	0, // 1: backup.v1beta1.Location.pmm_server_config:type_name -> backup.v1beta1.PMMServerLocationConfig
	2, // 2: backup.v1beta1.Location.s3_config:type_name -> backup.v1beta1.S3LocationConfig
	3, // 3: backup.v1beta1.ListLocationsResponse.locations:type_name -> backup.v1beta1.Location
	1, // 4: backup.v1beta1.AddLocationRequest.pmm_client_config:type_name -> backup.v1beta1.PMMClientLocationConfig
	0, // 5: backup.v1beta1.AddLocationRequest.pmm_server_config:type_name -> backup.v1beta1.PMMServerLocationConfig
	2, // 6: backup.v1beta1.AddLocationRequest.s3_config:type_name -> backup.v1beta1.S3LocationConfig
	4, // 7: backup.v1beta1.Locations.ListLocations:input_type -> backup.v1beta1.ListLocationsRequest
	6, // 8: backup.v1beta1.Locations.AddLocation:input_type -> backup.v1beta1.AddLocationRequest
	5, // 9: backup.v1beta1.Locations.ListLocations:output_type -> backup.v1beta1.ListLocationsResponse
	7, // 10: backup.v1beta1.Locations.AddLocation:output_type -> backup.v1beta1.AddLocationResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_managementpb_backup_locations_proto_init() }
func file_managementpb_backup_locations_proto_init() {
	if File_managementpb_backup_locations_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_managementpb_backup_locations_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMMServerLocationConfig); i {
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
		file_managementpb_backup_locations_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PMMClientLocationConfig); i {
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
		file_managementpb_backup_locations_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S3LocationConfig); i {
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
		file_managementpb_backup_locations_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_managementpb_backup_locations_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsRequest); i {
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
		file_managementpb_backup_locations_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsResponse); i {
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
		file_managementpb_backup_locations_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLocationRequest); i {
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
		file_managementpb_backup_locations_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddLocationResponse); i {
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
	file_managementpb_backup_locations_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Location_PmmClientConfig)(nil),
		(*Location_PmmServerConfig)(nil),
		(*Location_S3Config)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_managementpb_backup_locations_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_backup_locations_proto_goTypes,
		DependencyIndexes: file_managementpb_backup_locations_proto_depIdxs,
		MessageInfos:      file_managementpb_backup_locations_proto_msgTypes,
	}.Build()
	File_managementpb_backup_locations_proto = out.File
	file_managementpb_backup_locations_proto_rawDesc = nil
	file_managementpb_backup_locations_proto_goTypes = nil
	file_managementpb_backup_locations_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LocationsClient is the client API for Locations service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocationsClient interface {
	// ListLocations returns a list of all backup locations.
	ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error)
	// AddLocation adds backup location.
	AddLocation(ctx context.Context, in *AddLocationRequest, opts ...grpc.CallOption) (*AddLocationResponse, error)
}

type locationsClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationsClient(cc grpc.ClientConnInterface) LocationsClient {
	return &locationsClient{cc}
}

func (c *locationsClient) ListLocations(ctx context.Context, in *ListLocationsRequest, opts ...grpc.CallOption) (*ListLocationsResponse, error) {
	out := new(ListLocationsResponse)
	err := c.cc.Invoke(ctx, "/backup.v1beta1.Locations/ListLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationsClient) AddLocation(ctx context.Context, in *AddLocationRequest, opts ...grpc.CallOption) (*AddLocationResponse, error) {
	out := new(AddLocationResponse)
	err := c.cc.Invoke(ctx, "/backup.v1beta1.Locations/AddLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationsServer is the server API for Locations service.
type LocationsServer interface {
	// ListLocations returns a list of all backup locations.
	ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error)
	// AddLocation adds backup location.
	AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error)
}

// UnimplementedLocationsServer can be embedded to have forward compatible implementations.
type UnimplementedLocationsServer struct {
}

func (*UnimplementedLocationsServer) ListLocations(context.Context, *ListLocationsRequest) (*ListLocationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocations not implemented")
}
func (*UnimplementedLocationsServer) AddLocation(context.Context, *AddLocationRequest) (*AddLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddLocation not implemented")
}

func RegisterLocationsServer(s *grpc.Server, srv LocationsServer) {
	s.RegisterService(&_Locations_serviceDesc, srv)
}

func _Locations_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLocationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Locations/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).ListLocations(ctx, req.(*ListLocationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Locations_AddLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationsServer).AddLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/backup.v1beta1.Locations/AddLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationsServer).AddLocation(ctx, req.(*AddLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Locations_serviceDesc = grpc.ServiceDesc{
	ServiceName: "backup.v1beta1.Locations",
	HandlerType: (*LocationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLocations",
			Handler:    _Locations_ListLocations_Handler,
		},
		{
			MethodName: "AddLocation",
			Handler:    _Locations_AddLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/backup/locations.proto",
}
