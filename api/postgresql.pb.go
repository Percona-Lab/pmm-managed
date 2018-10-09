// Code generated by protoc-gen-go. DO NOT EDIT.
// source: postgresql.proto

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

type PostgreSQLNode struct {
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Name                 string   `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLNode) Reset()         { *m = PostgreSQLNode{} }
func (m *PostgreSQLNode) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLNode) ProtoMessage()    {}
func (*PostgreSQLNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{0}
}
func (m *PostgreSQLNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLNode.Unmarshal(m, b)
}
func (m *PostgreSQLNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLNode.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLNode.Merge(dst, src)
}
func (m *PostgreSQLNode) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLNode.Size(m)
}
func (m *PostgreSQLNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLNode.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLNode proto.InternalMessageInfo

func (m *PostgreSQLNode) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *PostgreSQLNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostgreSQLService struct {
	Address              string   `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	Port                 uint32   `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	Engine               string   `protobuf:"bytes,6,opt,name=engine,proto3" json:"engine,omitempty"`
	EngineVersion        string   `protobuf:"bytes,7,opt,name=engine_version,json=engineVersion,proto3" json:"engine_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLService) Reset()         { *m = PostgreSQLService{} }
func (m *PostgreSQLService) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLService) ProtoMessage()    {}
func (*PostgreSQLService) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{1}
}
func (m *PostgreSQLService) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLService.Unmarshal(m, b)
}
func (m *PostgreSQLService) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLService.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLService) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLService.Merge(dst, src)
}
func (m *PostgreSQLService) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLService.Size(m)
}
func (m *PostgreSQLService) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLService.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLService proto.InternalMessageInfo

func (m *PostgreSQLService) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *PostgreSQLService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *PostgreSQLService) GetEngine() string {
	if m != nil {
		return m.Engine
	}
	return ""
}

func (m *PostgreSQLService) GetEngineVersion() string {
	if m != nil {
		return m.EngineVersion
	}
	return ""
}

type PostgreSQLInstanceID struct {
	Region               string   `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLInstanceID) Reset()         { *m = PostgreSQLInstanceID{} }
func (m *PostgreSQLInstanceID) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLInstanceID) ProtoMessage()    {}
func (*PostgreSQLInstanceID) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{2}
}
func (m *PostgreSQLInstanceID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLInstanceID.Unmarshal(m, b)
}
func (m *PostgreSQLInstanceID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLInstanceID.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLInstanceID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLInstanceID.Merge(dst, src)
}
func (m *PostgreSQLInstanceID) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLInstanceID.Size(m)
}
func (m *PostgreSQLInstanceID) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLInstanceID.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLInstanceID proto.InternalMessageInfo

func (m *PostgreSQLInstanceID) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *PostgreSQLInstanceID) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PostgreSQLInstance struct {
	Node                 *PostgreSQLNode    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	Service              *PostgreSQLService `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *PostgreSQLInstance) Reset()         { *m = PostgreSQLInstance{} }
func (m *PostgreSQLInstance) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLInstance) ProtoMessage()    {}
func (*PostgreSQLInstance) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{3}
}
func (m *PostgreSQLInstance) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLInstance.Unmarshal(m, b)
}
func (m *PostgreSQLInstance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLInstance.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLInstance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLInstance.Merge(dst, src)
}
func (m *PostgreSQLInstance) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLInstance.Size(m)
}
func (m *PostgreSQLInstance) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLInstance.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLInstance proto.InternalMessageInfo

func (m *PostgreSQLInstance) GetNode() *PostgreSQLNode {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *PostgreSQLInstance) GetService() *PostgreSQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

type PostgreSQLListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLListRequest) Reset()         { *m = PostgreSQLListRequest{} }
func (m *PostgreSQLListRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLListRequest) ProtoMessage()    {}
func (*PostgreSQLListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{4}
}
func (m *PostgreSQLListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLListRequest.Unmarshal(m, b)
}
func (m *PostgreSQLListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLListRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLListRequest.Merge(dst, src)
}
func (m *PostgreSQLListRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLListRequest.Size(m)
}
func (m *PostgreSQLListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLListRequest proto.InternalMessageInfo

type PostgreSQLListResponse struct {
	Instances            []*PostgreSQLInstance `protobuf:"bytes,1,rep,name=instances,proto3" json:"instances,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PostgreSQLListResponse) Reset()         { *m = PostgreSQLListResponse{} }
func (m *PostgreSQLListResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLListResponse) ProtoMessage()    {}
func (*PostgreSQLListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{5}
}
func (m *PostgreSQLListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLListResponse.Unmarshal(m, b)
}
func (m *PostgreSQLListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLListResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLListResponse.Merge(dst, src)
}
func (m *PostgreSQLListResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLListResponse.Size(m)
}
func (m *PostgreSQLListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLListResponse proto.InternalMessageInfo

func (m *PostgreSQLListResponse) GetInstances() []*PostgreSQLInstance {
	if m != nil {
		return m.Instances
	}
	return nil
}

type PostgreSQLAddRequest struct {
	Id                   *PostgreSQLInstanceID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string                `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password             string                `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PostgreSQLAddRequest) Reset()         { *m = PostgreSQLAddRequest{} }
func (m *PostgreSQLAddRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLAddRequest) ProtoMessage()    {}
func (*PostgreSQLAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{6}
}
func (m *PostgreSQLAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLAddRequest.Unmarshal(m, b)
}
func (m *PostgreSQLAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLAddRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLAddRequest.Merge(dst, src)
}
func (m *PostgreSQLAddRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLAddRequest.Size(m)
}
func (m *PostgreSQLAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLAddRequest proto.InternalMessageInfo

func (m *PostgreSQLAddRequest) GetId() *PostgreSQLInstanceID {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *PostgreSQLAddRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *PostgreSQLAddRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type PostgreSQLAddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLAddResponse) Reset()         { *m = PostgreSQLAddResponse{} }
func (m *PostgreSQLAddResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLAddResponse) ProtoMessage()    {}
func (*PostgreSQLAddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{7}
}
func (m *PostgreSQLAddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLAddResponse.Unmarshal(m, b)
}
func (m *PostgreSQLAddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLAddResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLAddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLAddResponse.Merge(dst, src)
}
func (m *PostgreSQLAddResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLAddResponse.Size(m)
}
func (m *PostgreSQLAddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLAddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLAddResponse proto.InternalMessageInfo

type PostgreSQLRemoveRequest struct {
	Id                   *PostgreSQLInstanceID `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PostgreSQLRemoveRequest) Reset()         { *m = PostgreSQLRemoveRequest{} }
func (m *PostgreSQLRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLRemoveRequest) ProtoMessage()    {}
func (*PostgreSQLRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{8}
}
func (m *PostgreSQLRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Unmarshal(m, b)
}
func (m *PostgreSQLRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLRemoveRequest.Merge(dst, src)
}
func (m *PostgreSQLRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLRemoveRequest.Size(m)
}
func (m *PostgreSQLRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLRemoveRequest proto.InternalMessageInfo

func (m *PostgreSQLRemoveRequest) GetId() *PostgreSQLInstanceID {
	if m != nil {
		return m.Id
	}
	return nil
}

type PostgreSQLRemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PostgreSQLRemoveResponse) Reset()         { *m = PostgreSQLRemoveResponse{} }
func (m *PostgreSQLRemoveResponse) String() string { return proto.CompactTextString(m) }
func (*PostgreSQLRemoveResponse) ProtoMessage()    {}
func (*PostgreSQLRemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_postgresql_3e60736af9303a1b, []int{9}
}
func (m *PostgreSQLRemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Unmarshal(m, b)
}
func (m *PostgreSQLRemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Marshal(b, m, deterministic)
}
func (dst *PostgreSQLRemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PostgreSQLRemoveResponse.Merge(dst, src)
}
func (m *PostgreSQLRemoveResponse) XXX_Size() int {
	return xxx_messageInfo_PostgreSQLRemoveResponse.Size(m)
}
func (m *PostgreSQLRemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PostgreSQLRemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PostgreSQLRemoveResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PostgreSQLNode)(nil), "api.PostgreSQLNode")
	proto.RegisterType((*PostgreSQLService)(nil), "api.PostgreSQLService")
	proto.RegisterType((*PostgreSQLInstanceID)(nil), "api.PostgreSQLInstanceID")
	proto.RegisterType((*PostgreSQLInstance)(nil), "api.PostgreSQLInstance")
	proto.RegisterType((*PostgreSQLListRequest)(nil), "api.PostgreSQLListRequest")
	proto.RegisterType((*PostgreSQLListResponse)(nil), "api.PostgreSQLListResponse")
	proto.RegisterType((*PostgreSQLAddRequest)(nil), "api.PostgreSQLAddRequest")
	proto.RegisterType((*PostgreSQLAddResponse)(nil), "api.PostgreSQLAddResponse")
	proto.RegisterType((*PostgreSQLRemoveRequest)(nil), "api.PostgreSQLRemoveRequest")
	proto.RegisterType((*PostgreSQLRemoveResponse)(nil), "api.PostgreSQLRemoveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PostgreSQLClient is the client API for PostgreSQL service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PostgreSQLClient interface {
	List(ctx context.Context, in *PostgreSQLListRequest, opts ...grpc.CallOption) (*PostgreSQLListResponse, error)
	Add(ctx context.Context, in *PostgreSQLAddRequest, opts ...grpc.CallOption) (*PostgreSQLAddResponse, error)
	Remove(ctx context.Context, in *PostgreSQLRemoveRequest, opts ...grpc.CallOption) (*PostgreSQLRemoveResponse, error)
}

type postgreSQLClient struct {
	cc *grpc.ClientConn
}

func NewPostgreSQLClient(cc *grpc.ClientConn) PostgreSQLClient {
	return &postgreSQLClient{cc}
}

func (c *postgreSQLClient) List(ctx context.Context, in *PostgreSQLListRequest, opts ...grpc.CallOption) (*PostgreSQLListResponse, error) {
	out := new(PostgreSQLListResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgreSQLClient) Add(ctx context.Context, in *PostgreSQLAddRequest, opts ...grpc.CallOption) (*PostgreSQLAddResponse, error) {
	out := new(PostgreSQLAddResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postgreSQLClient) Remove(ctx context.Context, in *PostgreSQLRemoveRequest, opts ...grpc.CallOption) (*PostgreSQLRemoveResponse, error) {
	out := new(PostgreSQLRemoveResponse)
	err := c.cc.Invoke(ctx, "/api.PostgreSQL/Remove", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgreSQLServer is the server API for PostgreSQL service.
type PostgreSQLServer interface {
	List(context.Context, *PostgreSQLListRequest) (*PostgreSQLListResponse, error)
	Add(context.Context, *PostgreSQLAddRequest) (*PostgreSQLAddResponse, error)
	Remove(context.Context, *PostgreSQLRemoveRequest) (*PostgreSQLRemoveResponse, error)
}

func RegisterPostgreSQLServer(s *grpc.Server, srv PostgreSQLServer) {
	s.RegisterService(&_PostgreSQL_serviceDesc, srv)
}

func _PostgreSQL_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).List(ctx, req.(*PostgreSQLListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgreSQL_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLAddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).Add(ctx, req.(*PostgreSQLAddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PostgreSQL_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostgreSQLRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.PostgreSQL/Remove",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).Remove(ctx, req.(*PostgreSQLRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostgreSQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.PostgreSQL",
	HandlerType: (*PostgreSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _PostgreSQL_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _PostgreSQL_Add_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _PostgreSQL_Remove_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "postgresql.proto",
}

func init() { proto.RegisterFile("postgresql.proto", fileDescriptor_postgresql_3e60736af9303a1b) }

var fileDescriptor_postgresql_3e60736af9303a1b = []byte{
	// 494 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6e, 0xd3, 0x40,
	0x14, 0x96, 0x7f, 0x48, 0xd2, 0x17, 0x35, 0x0a, 0xaf, 0x90, 0xb8, 0xa6, 0x48, 0xd1, 0x48, 0x88,
	0xd0, 0x45, 0x52, 0x05, 0xb1, 0x61, 0x57, 0x94, 0x05, 0x89, 0x2a, 0x7e, 0x5c, 0x89, 0x2e, 0xc1,
	0x74, 0x9e, 0xac, 0x91, 0xda, 0x19, 0xd7, 0xe3, 0x06, 0xb1, 0xe5, 0x08, 0x70, 0x01, 0xee, 0xc4,
	0x15, 0x38, 0x08, 0xf2, 0x8c, 0x53, 0x3b, 0x4e, 0xba, 0x60, 0x37, 0x6f, 0xbe, 0x37, 0xdf, 0x9f,
	0x12, 0x43, 0x3f, 0x55, 0x3a, 0x4f, 0x32, 0xd2, 0x37, 0x57, 0x93, 0x34, 0x53, 0xb9, 0x42, 0x2f,
	0x4e, 0x45, 0x78, 0x94, 0x28, 0x95, 0x5c, 0xd1, 0x34, 0x4e, 0xc5, 0x34, 0x96, 0x52, 0xe5, 0x71,
	0x2e, 0x94, 0xd4, 0x76, 0x85, 0xbd, 0x85, 0xde, 0x07, 0xfb, 0xec, 0xfc, 0xe3, 0xd9, 0x3b, 0xc5,
	0x09, 0x07, 0xd0, 0xca, 0x28, 0x11, 0x4a, 0x06, 0xde, 0xc8, 0x19, 0xef, 0x45, 0xe5, 0x84, 0x08,
	0xbe, 0x8c, 0xaf, 0x29, 0xf0, 0xcd, 0xad, 0x39, 0x2f, 0xfd, 0x8e, 0xd3, 0x77, 0x97, 0x7e, 0xc7,
	0xed, 0x7b, 0xec, 0xa7, 0x03, 0x0f, 0x2b, 0xaa, 0x73, 0xca, 0x56, 0xe2, 0x92, 0x30, 0x80, 0x76,
	0xcc, 0x79, 0x46, 0x5a, 0x97, 0x0f, 0xd7, 0x63, 0xc1, 0x97, 0xaa, 0x2c, 0x0f, 0x1e, 0x8c, 0x9c,
	0xf1, 0x7e, 0x64, 0xce, 0x85, 0x36, 0xc9, 0x44, 0x48, 0x0a, 0x5a, 0x56, 0xdb, 0x4e, 0xf8, 0x0c,
	0x7a, 0xf6, 0xf4, 0x79, 0x45, 0x99, 0x2e, 0xbc, 0xb5, 0x0d, 0xbe, 0x6f, 0x6f, 0x3f, 0xd9, 0xcb,
	0xba, 0x9d, 0xa5, 0xdf, 0xf1, 0xfa, 0x3e, 0x7b, 0x03, 0x8f, 0x2a, 0x4f, 0x0b, 0xa9, 0xf3, 0x58,
	0x5e, 0xd2, 0x62, 0x5e, 0x0b, 0xe9, 0xec, 0x0c, 0xe9, 0x56, 0x21, 0x99, 0x02, 0xdc, 0xe6, 0xc0,
	0xe7, 0xe0, 0x4b, 0xc5, 0xc9, 0xbc, 0xef, 0xce, 0x0e, 0x26, 0x71, 0x2a, 0x26, 0x9b, 0x4d, 0x46,
	0x66, 0x01, 0x4f, 0xa0, 0xad, 0x6d, 0x19, 0x86, 0xb5, 0x3b, 0x1b, 0x34, 0x76, 0xcb, 0xaa, 0xa2,
	0xf5, 0x1a, 0x1b, 0xc2, 0xe3, 0x0a, 0x3d, 0x13, 0x3a, 0x8f, 0xe8, 0xe6, 0x96, 0x74, 0xce, 0xde,
	0xc3, 0xa0, 0x09, 0xe8, 0x54, 0x49, 0x4d, 0xf8, 0x0a, 0xf6, 0x44, 0xe9, 0x4c, 0x07, 0xce, 0xc8,
	0x1b, 0x77, 0x67, 0xc3, 0x86, 0xcc, 0xda, 0x79, 0x54, 0x6d, 0xb2, 0xef, 0xf5, 0x7a, 0x4e, 0x39,
	0x2f, 0x85, 0xf0, 0x05, 0xb8, 0x82, 0x97, 0xd1, 0x0e, 0xef, 0xe1, 0x59, 0xcc, 0x23, 0x57, 0x70,
	0x0c, 0xa1, 0x73, 0xab, 0x29, 0xab, 0xb5, 0x76, 0x37, 0x17, 0x58, 0x1a, 0x6b, 0xfd, 0x4d, 0x65,
	0xbc, 0xfc, 0x31, 0xdd, 0xcd, 0x9b, 0x21, 0x8d, 0xb4, 0x8d, 0xc2, 0xe6, 0x30, 0xac, 0x80, 0x88,
	0xae, 0xd5, 0x8a, 0xfe, 0xdf, 0x16, 0x0b, 0x21, 0xd8, 0x66, 0xb1, 0x0a, 0xb3, 0xdf, 0x2e, 0x40,
	0x05, 0xe2, 0x05, 0xf8, 0x45, 0x97, 0x18, 0x36, 0x18, 0x6b, 0xcd, 0x87, 0x4f, 0x76, 0x62, 0xa5,
	0xe3, 0xc1, 0x8f, 0x3f, 0x7f, 0x7f, 0xb9, 0x7d, 0xec, 0x4d, 0x57, 0x27, 0xd3, 0xea, 0x4f, 0x88,
	0x17, 0xe0, 0x9d, 0x72, 0x8e, 0x4d, 0xa7, 0x55, 0xcf, 0x61, 0xb8, 0x0b, 0x2a, 0x59, 0x0f, 0x0d,
	0xeb, 0x01, 0x6b, 0xb0, 0xbe, 0x76, 0x8e, 0xf1, 0x0b, 0xb4, 0x6c, 0x24, 0x3c, 0x6a, 0x10, 0x6c,
	0xf4, 0x15, 0x3e, 0xbd, 0x07, 0xdd, 0x54, 0x38, 0xde, 0x56, 0xf8, 0xda, 0x32, 0x5f, 0x87, 0x97,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x56, 0x80, 0x6d, 0x54, 0x04, 0x00, 0x00,
}
