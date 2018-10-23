// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package api

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

type RemoteNode struct {
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Region               string   `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteNode) Reset()         { *m = RemoteNode{} }
func (m *RemoteNode) String() string { return proto.CompactTextString(m) }
func (*RemoteNode) ProtoMessage()    {}
func (*RemoteNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_15aed4f7202f5474, []int{0}
}
func (m *RemoteNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteNode.Unmarshal(m, b)
}
func (m *RemoteNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteNode.Marshal(b, m, deterministic)
}
func (dst *RemoteNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteNode.Merge(dst, src)
}
func (m *RemoteNode) XXX_Size() int {
	return xxx_messageInfo_RemoteNode.Size(m)
}
func (m *RemoteNode) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteNode.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteNode proto.InternalMessageInfo

func (m *RemoteNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RemoteNode) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type RemoteService struct {
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Engine               string   `protobuf:"bytes,6,opt,name=engine,proto3" json:"engine,omitempty"`
	EngineVersion        string   `protobuf:"bytes,7,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteService) Reset()         { *m = RemoteService{} }
func (m *RemoteService) String() string { return proto.CompactTextString(m) }
func (*RemoteService) ProtoMessage()    {}
func (*RemoteService) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_15aed4f7202f5474, []int{1}
}
func (m *RemoteService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteService.Unmarshal(m, b)
}
func (m *RemoteService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteService.Marshal(b, m, deterministic)
}
func (dst *RemoteService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteService.Merge(dst, src)
}
func (m *RemoteService) XXX_Size() int {
	return xxx_messageInfo_RemoteService.Size(m)
}
func (m *RemoteService) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteService.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteService proto.InternalMessageInfo

func (m *RemoteService) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *RemoteService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *RemoteService) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *RemoteService) GetEngineVersion() string {
	if m != nil {
		return m.EngineVersion
	}
	return ""
}

type RemoteInstance struct {
	Node                 *RemoteNode    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Service              *RemoteService `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RemoteInstance) Reset()         { *m = RemoteInstance{} }
func (m *RemoteInstance) String() string { return proto.CompactTextString(m) }
func (*RemoteInstance) ProtoMessage()    {}
func (*RemoteInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_15aed4f7202f5474, []int{2}
}
func (m *RemoteInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteInstance.Unmarshal(m, b)
}
func (m *RemoteInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteInstance.Marshal(b, m, deterministic)
}
func (dst *RemoteInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteInstance.Merge(dst, src)
}
func (m *RemoteInstance) XXX_Size() int {
	return xxx_messageInfo_RemoteInstance.Size(m)
}
func (m *RemoteInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteInstance.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteInstance proto.InternalMessageInfo

func (m *RemoteInstance) GetNode() *RemoteNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *RemoteInstance) GetService() *RemoteService {
	if m != nil {
		return m.Service
	}
	return nil
}

type RemoteListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoteListRequest) Reset()         { *m = RemoteListRequest{} }
func (m *RemoteListRequest) String() string { return proto.CompactTextString(m) }
func (*RemoteListRequest) ProtoMessage()    {}
func (*RemoteListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_15aed4f7202f5474, []int{3}
}
func (m *RemoteListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteListRequest.Unmarshal(m, b)
}
func (m *RemoteListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteListRequest.Marshal(b, m, deterministic)
}
func (dst *RemoteListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteListRequest.Merge(dst, src)
}
func (m *RemoteListRequest) XXX_Size() int {
	return xxx_messageInfo_RemoteListRequest.Size(m)
}
func (m *RemoteListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteListRequest proto.InternalMessageInfo

type RemoteListResponse struct {
	Instances            []*RemoteInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RemoteListResponse) Reset()         { *m = RemoteListResponse{} }
func (m *RemoteListResponse) String() string { return proto.CompactTextString(m) }
func (*RemoteListResponse) ProtoMessage()    {}
func (*RemoteListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_15aed4f7202f5474, []int{4}
}
func (m *RemoteListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoteListResponse.Unmarshal(m, b)
}
func (m *RemoteListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoteListResponse.Marshal(b, m, deterministic)
}
func (dst *RemoteListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoteListResponse.Merge(dst, src)
}
func (m *RemoteListResponse) XXX_Size() int {
	return xxx_messageInfo_RemoteListResponse.Size(m)
}
func (m *RemoteListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoteListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoteListResponse proto.InternalMessageInfo

func (m *RemoteListResponse) GetInstances() []*RemoteInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

func init() {
	proto.RegisterType((*RemoteNode)(nil), "api.RemoteNode")
	proto.RegisterType((*RemoteService)(nil), "api.RemoteService")
	proto.RegisterType((*RemoteInstance)(nil), "api.RemoteInstance")
	proto.RegisterType((*RemoteListRequest)(nil), "api.RemoteListRequest")
	proto.RegisterType((*RemoteListResponse)(nil), "api.RemoteListResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteClient is the client API for Remote service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteClient interface {
	List(ctx context.Context, in *RemoteListRequest, opts ...grpc.CallOption) (*RemoteListResponse, error)
}

type remoteClient struct {
	cc *grpc.ClientConn
}

func NewRemoteClient(cc *grpc.ClientConn) RemoteClient {
	return &remoteClient{cc}
}

func (c *remoteClient) List(ctx context.Context, in *RemoteListRequest, opts ...grpc.CallOption) (*RemoteListResponse, error) {
	out := new(RemoteListResponse)
	err := c.cc.Invoke(ctx, "/api.Remote/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteServer is the server API for Remote service.
type RemoteServer interface {
	List(context.Context, *RemoteListRequest) (*RemoteListResponse, error)
}

func RegisterRemoteServer(s *grpc.Server, srv RemoteServer) {
	s.RegisterService(&_Remote_serviceDesc, srv)
}

func _Remote_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoteListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Remote/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServer).List(ctx, req.(*RemoteListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Remote_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Remote",
	HandlerType: (*RemoteServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Remote_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote.proto",
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor_remote_15aed4f7202f5474) }

var fileDescriptor_remote_15aed4f7202f5474 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x5d, 0x6a, 0xdb, 0x40,
	0x10, 0x46, 0xd2, 0x56, 0xb6, 0xc7, 0x3f, 0x75, 0xc7, 0xe0, 0x2e, 0xa6, 0x0f, 0x42, 0xa5, 0xa0,
	0x87, 0x62, 0xb5, 0xee, 0x15, 0x0a, 0x21, 0x4e, 0xc8, 0x83, 0x42, 0xf2, 0x1a, 0x36, 0xd6, 0x20,
	0x16, 0xe2, 0x5d, 0x45, 0xbb, 0xf1, 0x01, 0x72, 0x80, 0xbc, 0xe4, 0x68, 0xb9, 0x42, 0x0e, 0x12,
	0xb4, 0x2b, 0x63, 0x41, 0xde, 0x66, 0xbe, 0xef, 0x9b, 0xf9, 0xbe, 0x61, 0x17, 0x26, 0x0d, 0xed,
	0xb5, 0xa5, 0x75, 0xdd, 0x68, 0xab, 0x31, 0x12, 0xb5, 0x5c, 0xfd, 0xa8, 0xb4, 0xae, 0x1e, 0x28,
	0x17, 0xb5, 0xcc, 0x85, 0x52, 0xda, 0x0a, 0x2b, 0xb5, 0x32, 0x5e, 0x92, 0xfe, 0x07, 0x28, 0xdc,
	0xc8, 0x95, 0x2e, 0x09, 0x11, 0x98, 0x12, 0x7b, 0xe2, 0x51, 0x12, 0x64, 0xa3, 0xc2, 0xd5, 0xb8,
	0x84, 0xb8, 0xa1, 0x4a, 0x6a, 0xc5, 0x99, 0x43, 0xbb, 0x6e, 0xcb, 0x86, 0xc1, 0x3c, 0xdc, 0xb2,
	0x61, 0x38, 0x8f, 0xd2, 0x97, 0x00, 0xa6, 0x7e, 0xcd, 0x35, 0x35, 0x07, 0xb9, 0x23, 0xe4, 0x30,
	0x10, 0x65, 0xd9, 0x90, 0x31, 0xdd, 0xd8, 0xb1, 0x6d, 0x3d, 0x6a, 0xdd, 0x58, 0xfe, 0x25, 0x09,
	0xb2, 0x69, 0xe1, 0xea, 0xd6, 0x83, 0x54, 0x25, 0x15, 0xf1, 0xd8, 0x7b, 0xf8, 0x0e, 0x7f, 0xc1,
	0xcc, 0x57, 0x77, 0x07, 0x6a, 0x4c, 0x9b, 0x61, 0xe0, 0xf8, 0xa9, 0x47, 0x6f, 0x3d, 0xd8, 0x8f,
	0xb2, 0x65, 0xc3, 0x68, 0xce, 0xd2, 0x1d, 0xcc, 0x7c, 0x9e, 0x73, 0x65, 0xac, 0x50, 0x3b, 0xc2,
	0x9f, 0xc0, 0x94, 0x2e, 0x89, 0x07, 0x49, 0x90, 0x8d, 0x37, 0x5f, 0xd7, 0xa2, 0x96, 0xeb, 0xd3,
	0xe5, 0x85, 0x23, 0xf1, 0x37, 0x0c, 0x8c, 0x3f, 0x80, 0x87, 0x4e, 0x87, 0x3d, 0x5d, 0x77, 0x5a,
	0x71, 0x94, 0xa4, 0x0b, 0xf8, 0xe6, 0x99, 0x4b, 0x69, 0x6c, 0x41, 0x8f, 0x4f, 0x64, 0x6c, 0x7a,
	0x06, 0xd8, 0x07, 0x4d, 0xad, 0x95, 0x21, 0xfc, 0x0b, 0x23, 0xd9, 0x25, 0x31, 0x3c, 0x48, 0xa2,
	0x6c, 0xbc, 0x59, 0xf4, 0x56, 0x1f, 0x53, 0x16, 0x27, 0xd5, 0xe6, 0x06, 0x62, 0x4f, 0xe2, 0x05,
	0xb0, 0x76, 0x19, 0x2e, 0x7b, 0x13, 0x3d, 0xcb, 0xd5, 0xf7, 0x4f, 0xb8, 0x77, 0x4d, 0xf1, 0xf9,
	0xed, 0xfd, 0x35, 0x9c, 0x20, 0xe4, 0x87, 0x3f, 0xb9, 0xff, 0x19, 0xf7, 0xb1, 0x7b, 0xf7, 0x7f,
	0x1f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbb, 0xea, 0xbd, 0x3a, 0x2a, 0x02, 0x00, 0x00,
}
