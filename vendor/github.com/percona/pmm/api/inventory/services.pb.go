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

// MySQLService represents MySQL-compatible Service configuration.
type MySQLService struct {
	// Unique Service identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique user-defined Service name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Node identifier where this Service runs.
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// MySQL access address (DNS name or IP address).
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// MySQL access port.
	Port uint32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	// MySQL access UNIX socket path.
	UnixSocket           string   `protobuf:"bytes,6,opt,name=unix_socket,json=unixSocket,proto3" json:"unix_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLService) Reset()         { *m = MySQLService{} }
func (m *MySQLService) String() string { return proto.CompactTextString(m) }
func (*MySQLService) ProtoMessage()    {}
func (*MySQLService) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{0}
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

func (m *MySQLService) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MySQLService) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MySQLService) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *MySQLService) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MySQLService) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *MySQLService) GetUnixSocket() string {
	if m != nil {
		return m.UnixSocket
	}
	return ""
}

type ListServicesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListServicesRequest) Reset()         { *m = ListServicesRequest{} }
func (m *ListServicesRequest) String() string { return proto.CompactTextString(m) }
func (*ListServicesRequest) ProtoMessage()    {}
func (*ListServicesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{1}
}
func (m *ListServicesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesRequest.Unmarshal(m, b)
}
func (m *ListServicesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesRequest.Marshal(b, m, deterministic)
}
func (dst *ListServicesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesRequest.Merge(dst, src)
}
func (m *ListServicesRequest) XXX_Size() int {
	return xxx_messageInfo_ListServicesRequest.Size(m)
}
func (m *ListServicesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesRequest proto.InternalMessageInfo

type ListServicesResponse struct {
	Mysql                []*MySQLService `protobuf:"bytes,1,rep,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListServicesResponse) Reset()         { *m = ListServicesResponse{} }
func (m *ListServicesResponse) String() string { return proto.CompactTextString(m) }
func (*ListServicesResponse) ProtoMessage()    {}
func (*ListServicesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{2}
}
func (m *ListServicesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListServicesResponse.Unmarshal(m, b)
}
func (m *ListServicesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListServicesResponse.Marshal(b, m, deterministic)
}
func (dst *ListServicesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListServicesResponse.Merge(dst, src)
}
func (m *ListServicesResponse) XXX_Size() int {
	return xxx_messageInfo_ListServicesResponse.Size(m)
}
func (m *ListServicesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListServicesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListServicesResponse proto.InternalMessageInfo

func (m *ListServicesResponse) GetMysql() []*MySQLService {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type GetServiceRequest struct {
	// Unique Service identifier.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetServiceRequest) Reset()         { *m = GetServiceRequest{} }
func (m *GetServiceRequest) String() string { return proto.CompactTextString(m) }
func (*GetServiceRequest) ProtoMessage()    {}
func (*GetServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{3}
}
func (m *GetServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceRequest.Unmarshal(m, b)
}
func (m *GetServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceRequest.Marshal(b, m, deterministic)
}
func (dst *GetServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceRequest.Merge(dst, src)
}
func (m *GetServiceRequest) XXX_Size() int {
	return xxx_messageInfo_GetServiceRequest.Size(m)
}
func (m *GetServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceRequest proto.InternalMessageInfo

func (m *GetServiceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetServiceResponse struct {
	// Types that are valid to be assigned to Service:
	//	*GetServiceResponse_Mysql
	Service              isGetServiceResponse_Service `protobuf_oneof:"service"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *GetServiceResponse) Reset()         { *m = GetServiceResponse{} }
func (m *GetServiceResponse) String() string { return proto.CompactTextString(m) }
func (*GetServiceResponse) ProtoMessage()    {}
func (*GetServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{4}
}
func (m *GetServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetServiceResponse.Unmarshal(m, b)
}
func (m *GetServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetServiceResponse.Marshal(b, m, deterministic)
}
func (dst *GetServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetServiceResponse.Merge(dst, src)
}
func (m *GetServiceResponse) XXX_Size() int {
	return xxx_messageInfo_GetServiceResponse.Size(m)
}
func (m *GetServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetServiceResponse proto.InternalMessageInfo

type isGetServiceResponse_Service interface {
	isGetServiceResponse_Service()
}

type GetServiceResponse_Mysql struct {
	Mysql *MySQLService `protobuf:"bytes,1,opt,name=mysql,proto3,oneof"`
}

func (*GetServiceResponse_Mysql) isGetServiceResponse_Service() {}

func (m *GetServiceResponse) GetService() isGetServiceResponse_Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *GetServiceResponse) GetMysql() *MySQLService {
	if x, ok := m.GetService().(*GetServiceResponse_Mysql); ok {
		return x.Mysql
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GetServiceResponse) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GetServiceResponse_OneofMarshaler, _GetServiceResponse_OneofUnmarshaler, _GetServiceResponse_OneofSizer, []interface{}{
		(*GetServiceResponse_Mysql)(nil),
	}
}

func _GetServiceResponse_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GetServiceResponse)
	// service
	switch x := m.Service.(type) {
	case *GetServiceResponse_Mysql:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Mysql); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GetServiceResponse.Service has unexpected type %T", x)
	}
	return nil
}

func _GetServiceResponse_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GetServiceResponse)
	switch tag {
	case 1: // service.mysql
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MySQLService)
		err := b.DecodeMessage(msg)
		m.Service = &GetServiceResponse_Mysql{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GetServiceResponse_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GetServiceResponse)
	// service
	switch x := m.Service.(type) {
	case *GetServiceResponse_Mysql:
		s := proto.Size(x.Mysql)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AddMySQLServiceRequest struct {
	// Unique user-defined Service name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Node identifier where this Service runs.
	NodeId string `protobuf:"bytes,3,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// MySQL access address (DNS name or IP address).
	Address string `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	// MySQL access port.
	Port uint32 `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	// MySQL access UNIX socket path.
	UnixSocket           string   `protobuf:"bytes,6,opt,name=unix_socket,json=unixSocket,proto3" json:"unix_socket,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLServiceRequest) Reset()         { *m = AddMySQLServiceRequest{} }
func (m *AddMySQLServiceRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceRequest) ProtoMessage()    {}
func (*AddMySQLServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{5}
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

func (m *AddMySQLServiceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddMySQLServiceRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddMySQLServiceRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddMySQLServiceRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddMySQLServiceRequest) GetUnixSocket() string {
	if m != nil {
		return m.UnixSocket
	}
	return ""
}

type AddMySQLServiceResponse struct {
	Mysql                *MySQLService `protobuf:"bytes,1,opt,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddMySQLServiceResponse) Reset()         { *m = AddMySQLServiceResponse{} }
func (m *AddMySQLServiceResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLServiceResponse) ProtoMessage()    {}
func (*AddMySQLServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{6}
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

func (m *AddMySQLServiceResponse) GetMysql() *MySQLService {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type ChangeMySQLServiceRequest struct {
	// Unique Service identifier.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique user-defined Service name.
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangeMySQLServiceRequest) Reset()         { *m = ChangeMySQLServiceRequest{} }
func (m *ChangeMySQLServiceRequest) String() string { return proto.CompactTextString(m) }
func (*ChangeMySQLServiceRequest) ProtoMessage()    {}
func (*ChangeMySQLServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{7}
}
func (m *ChangeMySQLServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeMySQLServiceRequest.Unmarshal(m, b)
}
func (m *ChangeMySQLServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeMySQLServiceRequest.Marshal(b, m, deterministic)
}
func (dst *ChangeMySQLServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeMySQLServiceRequest.Merge(dst, src)
}
func (m *ChangeMySQLServiceRequest) XXX_Size() int {
	return xxx_messageInfo_ChangeMySQLServiceRequest.Size(m)
}
func (m *ChangeMySQLServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeMySQLServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeMySQLServiceRequest proto.InternalMessageInfo

func (m *ChangeMySQLServiceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ChangeMySQLServiceRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ChangeMySQLServiceResponse struct {
	Mysql                *MySQLService `protobuf:"bytes,1,opt,name=mysql,proto3" json:"mysql,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *ChangeMySQLServiceResponse) Reset()         { *m = ChangeMySQLServiceResponse{} }
func (m *ChangeMySQLServiceResponse) String() string { return proto.CompactTextString(m) }
func (*ChangeMySQLServiceResponse) ProtoMessage()    {}
func (*ChangeMySQLServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{8}
}
func (m *ChangeMySQLServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangeMySQLServiceResponse.Unmarshal(m, b)
}
func (m *ChangeMySQLServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangeMySQLServiceResponse.Marshal(b, m, deterministic)
}
func (dst *ChangeMySQLServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeMySQLServiceResponse.Merge(dst, src)
}
func (m *ChangeMySQLServiceResponse) XXX_Size() int {
	return xxx_messageInfo_ChangeMySQLServiceResponse.Size(m)
}
func (m *ChangeMySQLServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeMySQLServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeMySQLServiceResponse proto.InternalMessageInfo

func (m *ChangeMySQLServiceResponse) GetMysql() *MySQLService {
	if m != nil {
		return m.Mysql
	}
	return nil
}

type RemoveServiceRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceRequest) Reset()         { *m = RemoveServiceRequest{} }
func (m *RemoveServiceRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceRequest) ProtoMessage()    {}
func (*RemoveServiceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{9}
}
func (m *RemoveServiceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServiceRequest.Unmarshal(m, b)
}
func (m *RemoveServiceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServiceRequest.Marshal(b, m, deterministic)
}
func (dst *RemoveServiceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServiceRequest.Merge(dst, src)
}
func (m *RemoveServiceRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveServiceRequest.Size(m)
}
func (m *RemoveServiceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServiceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServiceRequest proto.InternalMessageInfo

func (m *RemoveServiceRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type RemoveServiceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveServiceResponse) Reset()         { *m = RemoveServiceResponse{} }
func (m *RemoveServiceResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveServiceResponse) ProtoMessage()    {}
func (*RemoveServiceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_services_59e71f475cb58fcb, []int{10}
}
func (m *RemoveServiceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveServiceResponse.Unmarshal(m, b)
}
func (m *RemoveServiceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveServiceResponse.Marshal(b, m, deterministic)
}
func (dst *RemoveServiceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveServiceResponse.Merge(dst, src)
}
func (m *RemoveServiceResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveServiceResponse.Size(m)
}
func (m *RemoveServiceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveServiceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveServiceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MySQLService)(nil), "inventory.MySQLService")
	proto.RegisterType((*ListServicesRequest)(nil), "inventory.ListServicesRequest")
	proto.RegisterType((*ListServicesResponse)(nil), "inventory.ListServicesResponse")
	proto.RegisterType((*GetServiceRequest)(nil), "inventory.GetServiceRequest")
	proto.RegisterType((*GetServiceResponse)(nil), "inventory.GetServiceResponse")
	proto.RegisterType((*AddMySQLServiceRequest)(nil), "inventory.AddMySQLServiceRequest")
	proto.RegisterType((*AddMySQLServiceResponse)(nil), "inventory.AddMySQLServiceResponse")
	proto.RegisterType((*ChangeMySQLServiceRequest)(nil), "inventory.ChangeMySQLServiceRequest")
	proto.RegisterType((*ChangeMySQLServiceResponse)(nil), "inventory.ChangeMySQLServiceResponse")
	proto.RegisterType((*RemoveServiceRequest)(nil), "inventory.RemoveServiceRequest")
	proto.RegisterType((*RemoveServiceResponse)(nil), "inventory.RemoveServiceResponse")
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
	// ListServices returns a list of all Services.
	ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error)
	// ChangeMySQLService changes MySQL Service.
	ChangeMySQLService(ctx context.Context, in *ChangeMySQLServiceRequest, opts ...grpc.CallOption) (*ChangeMySQLServiceResponse, error)
	// RemoveService removes Service without any Agents.
	RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error)
}

type servicesClient struct {
	cc *grpc.ClientConn
}

func NewServicesClient(cc *grpc.ClientConn) ServicesClient {
	return &servicesClient{cc}
}

func (c *servicesClient) ListServices(ctx context.Context, in *ListServicesRequest, opts ...grpc.CallOption) (*ListServicesResponse, error) {
	out := new(ListServicesResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/ListServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) GetService(ctx context.Context, in *GetServiceRequest, opts ...grpc.CallOption) (*GetServiceResponse, error) {
	out := new(GetServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/GetService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) AddMySQLService(ctx context.Context, in *AddMySQLServiceRequest, opts ...grpc.CallOption) (*AddMySQLServiceResponse, error) {
	out := new(AddMySQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/AddMySQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) ChangeMySQLService(ctx context.Context, in *ChangeMySQLServiceRequest, opts ...grpc.CallOption) (*ChangeMySQLServiceResponse, error) {
	out := new(ChangeMySQLServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/ChangeMySQLService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *servicesClient) RemoveService(ctx context.Context, in *RemoveServiceRequest, opts ...grpc.CallOption) (*RemoveServiceResponse, error) {
	out := new(RemoveServiceResponse)
	err := c.cc.Invoke(ctx, "/inventory.Services/RemoveService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServicesServer is the server API for Services service.
type ServicesServer interface {
	// ListServices returns a list of all Services.
	ListServices(context.Context, *ListServicesRequest) (*ListServicesResponse, error)
	// GetService returns a single Service by ID.
	GetService(context.Context, *GetServiceRequest) (*GetServiceResponse, error)
	// AddMySQLService adds MySQL Service.
	AddMySQLService(context.Context, *AddMySQLServiceRequest) (*AddMySQLServiceResponse, error)
	// ChangeMySQLService changes MySQL Service.
	ChangeMySQLService(context.Context, *ChangeMySQLServiceRequest) (*ChangeMySQLServiceResponse, error)
	// RemoveService removes Service without any Agents.
	RemoveService(context.Context, *RemoveServiceRequest) (*RemoveServiceResponse, error)
}

func RegisterServicesServer(s *grpc.Server, srv ServicesServer) {
	s.RegisterService(&_Services_serviceDesc, srv)
}

func _Services_ListServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ListServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/ListServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ListServices(ctx, req.(*ListServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_GetService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).GetService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/GetService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).GetService(ctx, req.(*GetServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_AddMySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).AddMySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/AddMySQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).AddMySQLService(ctx, req.(*AddMySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_ChangeMySQLService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeMySQLServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).ChangeMySQLService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/ChangeMySQLService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).ChangeMySQLService(ctx, req.(*ChangeMySQLServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Services_RemoveService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServicesServer).RemoveService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Services/RemoveService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServicesServer).RemoveService(ctx, req.(*RemoveServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Services_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Services",
	HandlerType: (*ServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListServices",
			Handler:    _Services_ListServices_Handler,
		},
		{
			MethodName: "GetService",
			Handler:    _Services_GetService_Handler,
		},
		{
			MethodName: "AddMySQLService",
			Handler:    _Services_AddMySQLService_Handler,
		},
		{
			MethodName: "ChangeMySQLService",
			Handler:    _Services_ChangeMySQLService_Handler,
		},
		{
			MethodName: "RemoveService",
			Handler:    _Services_RemoveService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/services.proto",
}

func init() { proto.RegisterFile("inventory/services.proto", fileDescriptor_services_59e71f475cb58fcb) }

var fileDescriptor_services_59e71f475cb58fcb = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xfd, 0xc6, 0xcd, 0x4f, 0x73, 0xd3, 0x7e, 0xc0, 0xd0, 0x12, 0x63, 0x5a, 0x62, 0x86, 0xb6,
	0x44, 0x91, 0x1a, 0xa3, 0xb2, 0xeb, 0x06, 0x51, 0x84, 0x5a, 0xa0, 0x48, 0xe0, 0xec, 0xd8, 0x54,
	0xa6, 0x1e, 0x85, 0x11, 0xcd, 0x8c, 0xeb, 0x71, 0x03, 0x91, 0x58, 0x21, 0xb6, 0x48, 0x48, 0xec,
	0x78, 0x2d, 0x5e, 0x81, 0x07, 0x41, 0x1e, 0x8f, 0x9b, 0x49, 0x62, 0xb7, 0x15, 0x1b, 0x76, 0xf6,
	0xbd, 0xe7, 0x9e, 0x73, 0xe6, 0xce, 0xb1, 0xc1, 0x66, 0x7c, 0x44, 0x79, 0x22, 0xe2, 0xb1, 0x27,
	0x69, 0x3c, 0x62, 0xc7, 0x54, 0xf6, 0xa2, 0x58, 0x24, 0x02, 0x37, 0xce, 0x3b, 0xce, 0xda, 0x40,
	0x88, 0xc1, 0x09, 0xf5, 0x82, 0x88, 0x79, 0x01, 0xe7, 0x22, 0x09, 0x12, 0x26, 0xb8, 0x06, 0x92,
	0x9f, 0x08, 0x96, 0x5e, 0x8d, 0xfb, 0x6f, 0x0e, 0xfb, 0x19, 0x01, 0xfe, 0x1f, 0x2c, 0x16, 0xda,
	0xc8, 0x45, 0x9d, 0x86, 0x6f, 0xb1, 0x10, 0x63, 0xa8, 0xf0, 0x60, 0x48, 0x6d, 0x4b, 0x55, 0xd4,
	0x33, 0x6e, 0x41, 0x9d, 0x8b, 0x90, 0x1e, 0xb1, 0xd0, 0x5e, 0x50, 0xe5, 0x5a, 0xfa, 0xfa, 0x3c,
	0xc4, 0x36, 0xd4, 0x83, 0x30, 0x8c, 0xa9, 0x94, 0x76, 0x45, 0x35, 0xf2, 0xd7, 0x94, 0x26, 0x12,
	0x71, 0x62, 0x57, 0x5d, 0xd4, 0x59, 0xf6, 0xd5, 0x33, 0x6e, 0x43, 0xf3, 0x8c, 0xb3, 0x4f, 0x47,
	0x52, 0x1c, 0x7f, 0xa0, 0x89, 0x5d, 0x53, 0x13, 0x90, 0x96, 0xfa, 0xaa, 0x42, 0x56, 0xe1, 0xe6,
	0x21, 0x93, 0x89, 0xb6, 0x26, 0x7d, 0x7a, 0x7a, 0x46, 0x65, 0x42, 0x9e, 0xc1, 0xca, 0x74, 0x59,
	0x46, 0x82, 0x4b, 0x8a, 0xb7, 0xa1, 0x3a, 0x1c, 0xcb, 0xd3, 0x13, 0x1b, 0xb9, 0x0b, 0x9d, 0xe6,
	0x4e, 0xab, 0x77, 0xbe, 0x84, 0x9e, 0x79, 0x44, 0x3f, 0x43, 0x91, 0xfb, 0x70, 0x63, 0x9f, 0xe6,
	0x2c, 0x9a, 0x7b, 0xf6, 0xf8, 0xe4, 0x35, 0x60, 0x13, 0xa4, 0x95, 0xbc, 0x89, 0x12, 0xba, 0x40,
	0xe9, 0xe0, 0x3f, 0xad, 0xb5, 0xd7, 0x80, 0xba, 0xbe, 0xa1, 0x74, 0xe3, 0xb7, 0x9e, 0x84, 0xe1,
	0x94, 0x23, 0x2d, 0xfe, 0x0f, 0x77, 0xfd, 0xa2, 0xb2, 0x88, 0xae, 0x5b, 0xe4, 0x00, 0x5a, 0x73,
	0xde, 0xe6, 0xb7, 0x8b, 0xae, 0xb0, 0xdd, 0xc7, 0x70, 0xfb, 0xe9, 0xfb, 0x80, 0x0f, 0x68, 0xd1,
	0x41, 0xaf, 0x10, 0x32, 0xf2, 0x12, 0x9c, 0x22, 0x82, 0xbf, 0x73, 0xb3, 0x05, 0x2b, 0x3e, 0x1d,
	0x8a, 0x11, 0xbd, 0xe4, 0xba, 0x5b, 0xb0, 0x3a, 0x83, 0xcb, 0xf4, 0x76, 0xbe, 0x55, 0x61, 0x31,
	0x0f, 0x1c, 0xfe, 0x08, 0x4b, 0x66, 0x00, 0xf1, 0x5d, 0x43, 0xbd, 0x20, 0xb0, 0x4e, 0xbb, 0xb4,
	0x9f, 0xb1, 0x93, 0xad, 0x2f, 0xbf, 0x7e, 0xff, 0xb0, 0x5c, 0x72, 0xc7, 0x1b, 0x3d, 0xf4, 0x26,
	0x1f, 0x75, 0x8e, 0xf3, 0xd2, 0xa1, 0x5d, 0xd4, 0xc5, 0x11, 0xc0, 0x24, 0x8d, 0x78, 0xcd, 0xa0,
	0x9d, 0x4b, 0xb2, 0xb3, 0x5e, 0xd2, 0xd5, 0x92, 0x9b, 0x4a, 0xb2, 0x4d, 0x9c, 0x12, 0xc9, 0x7d,
	0xaa, 0x14, 0xbf, 0x22, 0xb8, 0x36, 0x93, 0x08, 0x7c, 0xcf, 0x60, 0x2e, 0x4e, 0xb2, 0x43, 0x2e,
	0x82, 0x68, 0x07, 0x5d, 0xe5, 0x60, 0x83, 0xb4, 0x4b, 0x1c, 0xe4, 0x73, 0xa9, 0x8d, 0xef, 0x08,
	0xf0, 0x7c, 0x1a, 0xf0, 0x86, 0x21, 0x53, 0x9a, 0x36, 0x67, 0xf3, 0x12, 0x94, 0xf6, 0xb3, 0xad,
	0xfc, 0x3c, 0x20, 0xa4, 0xc4, 0x8f, 0x31, 0x9a, 0x5a, 0xfa, 0x0c, 0xcb, 0x53, 0x51, 0xc1, 0xe6,
	0x2d, 0x17, 0x85, 0xcd, 0x71, 0xcb, 0x01, 0xda, 0x42, 0x47, 0x59, 0x20, 0x64, 0xbd, 0xc4, 0x42,
	0x36, 0xb5, 0x8b, 0xba, 0x7b, 0xcd, 0xb7, 0x93, 0x5f, 0xfc, 0xbb, 0x9a, 0xfa, 0x97, 0x3f, 0xfa,
	0x13, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x6c, 0xbd, 0xb3, 0x10, 0x06, 0x00, 0x00,
}
