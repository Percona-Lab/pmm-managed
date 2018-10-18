// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/services.proto

package inventory

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MySQLService struct {
	// Unique service identifier.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique service name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLService) Reset()         { *m = MySQLService{} }
func (m *MySQLService) String() string { return proto.CompactTextString(m) }
func (*MySQLService) ProtoMessage()    {}
func (*MySQLService) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_73dccc0aaa302009, []int{0}
}
func (m *MySQLService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLService.Unmarshal(m, b)
}
func (m *MySQLService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLService.Marshal(b, m, deterministic)
}
func (dst *MySQLService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLService.Merge(dst, src)
}
func (m *MySQLService) XXX_Size() int {
	return xxx_messageInfo_MySQLService.Size(m)
}
func (m *MySQLService) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLService.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLService proto.InternalMessageInfo

func (m *MySQLService) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MySQLService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type AddMySQLServiceRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLServiceRequest) Reset()         { *m = AddMySQLServiceRequest{} }
func (m *AddMySQLServiceRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceRequest) ProtoMessage()    {}
func (*AddMySQLServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_73dccc0aaa302009, []int{1}
}
func (m *AddMySQLServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLServiceRequest.Unmarshal(m, b)
}
func (m *AddMySQLServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLServiceRequest.Marshal(b, m, deterministic)
}
func (dst *AddMySQLServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLServiceRequest.Merge(dst, src)
}
func (m *AddMySQLServiceRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLServiceRequest.Size(m)
}
func (m *AddMySQLServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLServiceRequest proto.InternalMessageInfo

type AddMySQLServiceResponse struct {
	Service              *MySQLService `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddMySQLServiceResponse) Reset()         { *m = AddMySQLServiceResponse{} }
func (m *AddMySQLServiceResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceResponse) ProtoMessage()    {}
func (*AddMySQLServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_73dccc0aaa302009, []int{2}
}
func (m *AddMySQLServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLServiceResponse.Unmarshal(m, b)
}
func (m *AddMySQLServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLServiceResponse.Marshal(b, m, deterministic)
}
func (dst *AddMySQLServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLServiceResponse.Merge(dst, src)
}
func (m *AddMySQLServiceResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLServiceResponse.Size(m)
}
func (m *AddMySQLServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLServiceResponse proto.InternalMessageInfo

func (m *AddMySQLServiceResponse) GetService() *MySQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

func init() {
	proto.RegisterType((*MySQLService)(nil), "inventory.MySQLService")
	proto.RegisterType((*AddMySQLServiceRequest)(nil), "inventory.AddMySQLServiceRequest")
	proto.RegisterType((*AddMySQLServiceResponse)(nil), "inventory.AddMySQLServiceResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ServicesClient is the client API for Services service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServicesClient interface {
}

type servicesClient struct {
	cc *grpc.ClientConn
}

func NewServicesClient(cc *grpc.ClientConn) ServicesClient {
	return &servicesClient{cc}
}

// ServicesServer is the server API for Services service.
type ServicesServer interface {
}

func RegisterServicesServer(s *grpc.Server, srv ServicesServer) {
	s.RegisterService(&_Services_serviceDesc, srv)
}

var _Services_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "inventory/services.proto",
}

func init() { proto.RegisterFile("inventory/services.proto", fileDescriptor_services_73dccc0aaa302009) }

var fileDescriptor_services_73dccc0aaa302009 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc8, 0xcc, 0x2b, 0x4b,
	0xcd, 0x2b, 0xc9, 0x2f, 0xaa, 0xd4, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x2d, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0xcb, 0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4,
	0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x41,
	0x15, 0x2a, 0x19, 0x71, 0xf1, 0xf8, 0x56, 0x06, 0x07, 0xfa, 0x04, 0x43, 0xf4, 0x0b, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0x12,
	0x5c, 0x62, 0x8e, 0x29, 0x29, 0xc8, 0xda, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0x7c,
	0xb8, 0xc4, 0x31, 0x64, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x0c, 0xb9, 0xd8, 0xa1, 0x6e,
	0x04, 0x9b, 0xce, 0x6d, 0x24, 0xae, 0x07, 0x77, 0xa3, 0x1e, 0x8a, 0x0e, 0x98, 0x3a, 0x23, 0x2e,
	0x2e, 0x0e, 0xa8, 0x58, 0x71, 0x12, 0x1b, 0xd8, 0xb9, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xf6, 0x0f, 0xe4, 0xfe, 0xf3, 0x00, 0x00, 0x00,
}
