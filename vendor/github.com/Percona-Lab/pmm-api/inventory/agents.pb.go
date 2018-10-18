// Code generated by protoc-gen-go. DO NOT EDIT.
// source: inventory/agents.proto

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

type MySQLdExporter struct {
	// Unique agent identifier.
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Node identifier where agent runs.
	RunsOnNodeId uint32 `protobuf:"varint,2,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	// MySQL username for extracting metrics.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for extracting metrics.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// HTTP listen port for exposing metrics.
	ListenPort           uint32   `protobuf:"varint,5,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MySQLdExporter) Reset()         { *m = MySQLdExporter{} }
func (m *MySQLdExporter) String() string { return proto.CompactTextString(m) }
func (*MySQLdExporter) ProtoMessage()    {}
func (*MySQLdExporter) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_af741dd464ceedaa, []int{0}
}
func (m *MySQLdExporter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MySQLdExporter.Unmarshal(m, b)
}
func (m *MySQLdExporter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MySQLdExporter.Marshal(b, m, deterministic)
}
func (dst *MySQLdExporter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MySQLdExporter.Merge(dst, src)
}
func (m *MySQLdExporter) XXX_Size() int {
	return xxx_messageInfo_MySQLdExporter.Size(m)
}
func (m *MySQLdExporter) XXX_DiscardUnknown() {
	xxx_messageInfo_MySQLdExporter.DiscardUnknown(m)
}

var xxx_messageInfo_MySQLdExporter proto.InternalMessageInfo

func (m *MySQLdExporter) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MySQLdExporter) GetRunsOnNodeId() uint32 {
	if m != nil {
		return m.RunsOnNodeId
	}
	return 0
}

func (m *MySQLdExporter) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *MySQLdExporter) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *MySQLdExporter) GetListenPort() uint32 {
	if m != nil {
		return m.ListenPort
	}
	return 0
}

type AddMySQLdExporterRequest struct {
	// Node identifier where agent should run.
	RunsOnNodeId uint32 `protobuf:"varint,1,opt,name=runs_on_node_id,json=runsOnNodeId,proto3" json:"runs_on_node_id,omitempty"`
	// MySQL username for extracting metrics.
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	// MySQL password for extracting metrics.
	Password             string   `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMySQLdExporterRequest) Reset()         { *m = AddMySQLdExporterRequest{} }
func (m *AddMySQLdExporterRequest) String() string { return proto.CompactTextString(m) }
func (*AddMySQLdExporterRequest) ProtoMessage()    {}
func (*AddMySQLdExporterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_af741dd464ceedaa, []int{1}
}
func (m *AddMySQLdExporterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLdExporterRequest.Unmarshal(m, b)
}
func (m *AddMySQLdExporterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLdExporterRequest.Marshal(b, m, deterministic)
}
func (dst *AddMySQLdExporterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLdExporterRequest.Merge(dst, src)
}
func (m *AddMySQLdExporterRequest) XXX_Size() int {
	return xxx_messageInfo_AddMySQLdExporterRequest.Size(m)
}
func (m *AddMySQLdExporterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLdExporterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLdExporterRequest proto.InternalMessageInfo

func (m *AddMySQLdExporterRequest) GetRunsOnNodeId() uint32 {
	if m != nil {
		return m.RunsOnNodeId
	}
	return 0
}

func (m *AddMySQLdExporterRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMySQLdExporterRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AddMySQLdExporterResponse struct {
	Agent                *MySQLdExporter `protobuf:"bytes,1,opt,name=agent,proto3" json:"agent,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AddMySQLdExporterResponse) Reset()         { *m = AddMySQLdExporterResponse{} }
func (m *AddMySQLdExporterResponse) String() string { return proto.CompactTextString(m) }
func (*AddMySQLdExporterResponse) ProtoMessage()    {}
func (*AddMySQLdExporterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_agents_af741dd464ceedaa, []int{2}
}
func (m *AddMySQLdExporterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMySQLdExporterResponse.Unmarshal(m, b)
}
func (m *AddMySQLdExporterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMySQLdExporterResponse.Marshal(b, m, deterministic)
}
func (dst *AddMySQLdExporterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMySQLdExporterResponse.Merge(dst, src)
}
func (m *AddMySQLdExporterResponse) XXX_Size() int {
	return xxx_messageInfo_AddMySQLdExporterResponse.Size(m)
}
func (m *AddMySQLdExporterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMySQLdExporterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMySQLdExporterResponse proto.InternalMessageInfo

func (m *AddMySQLdExporterResponse) GetAgent() *MySQLdExporter {
	if m != nil {
		return m.Agent
	}
	return nil
}

func init() {
	proto.RegisterType((*MySQLdExporter)(nil), "inventory.MySQLdExporter")
	proto.RegisterType((*AddMySQLdExporterRequest)(nil), "inventory.AddMySQLdExporterRequest")
	proto.RegisterType((*AddMySQLdExporterResponse)(nil), "inventory.AddMySQLdExporterResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AgentsClient is the client API for Agents service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentsClient interface {
	// Add mysqld_exporter agent.
	AddMySQLdExporter(ctx context.Context, in *AddMySQLdExporterRequest, opts ...grpc.CallOption) (*AddMySQLdExporterResponse, error)
}

type agentsClient struct {
	cc *grpc.ClientConn
}

func NewAgentsClient(cc *grpc.ClientConn) AgentsClient {
	return &agentsClient{cc}
}

func (c *agentsClient) AddMySQLdExporter(ctx context.Context, in *AddMySQLdExporterRequest, opts ...grpc.CallOption) (*AddMySQLdExporterResponse, error) {
	out := new(AddMySQLdExporterResponse)
	err := c.cc.Invoke(ctx, "/inventory.Agents/AddMySQLdExporter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentsServer is the server API for Agents service.
type AgentsServer interface {
	// Add mysqld_exporter agent.
	AddMySQLdExporter(context.Context, *AddMySQLdExporterRequest) (*AddMySQLdExporterResponse, error)
}

func RegisterAgentsServer(s *grpc.Server, srv AgentsServer) {
	s.RegisterService(&_Agents_serviceDesc, srv)
}

func _Agents_AddMySQLdExporter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMySQLdExporterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentsServer).AddMySQLdExporter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.Agents/AddMySQLdExporter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentsServer).AddMySQLdExporter(ctx, req.(*AddMySQLdExporterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agents_serviceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.Agents",
	HandlerType: (*AgentsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMySQLdExporter",
			Handler:    _Agents_AddMySQLdExporter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "inventory/agents.proto",
}

func init() { proto.RegisterFile("inventory/agents.proto", fileDescriptor_agents_af741dd464ceedaa) }

var fileDescriptor_agents_af741dd464ceedaa = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcb, 0x4e, 0x32, 0x31,
	0x14, 0xc7, 0xd3, 0xe1, 0x83, 0x7c, 0x14, 0xc5, 0xd8, 0x85, 0x19, 0x88, 0x89, 0x64, 0xbc, 0x84,
	0x68, 0x42, 0x15, 0x77, 0xee, 0x58, 0xb8, 0x30, 0xc1, 0x1b, 0x3e, 0xc0, 0x64, 0x4c, 0x4f, 0x48,
	0x13, 0x3c, 0x67, 0x6c, 0x0b, 0xca, 0xd6, 0x37, 0x30, 0x6e, 0x5d, 0xf8, 0x4e, 0xbe, 0x82, 0x0f,
	0x62, 0xe8, 0xe8, 0x28, 0x82, 0xb8, 0x3c, 0xb7, 0xff, 0xff, 0xd7, 0xd3, 0xc3, 0xd7, 0x34, 0x8e,
	0x00, 0x1d, 0x99, 0xb1, 0x4c, 0xfa, 0x80, 0xce, 0xb6, 0x52, 0x43, 0x8e, 0x44, 0x39, 0xcf, 0xd7,
	0xd7, 0xfb, 0x44, 0xfd, 0x01, 0xc8, 0x24, 0xd5, 0x32, 0x41, 0x24, 0x97, 0x38, 0x4d, 0xf8, 0xd1,
	0x18, 0xbd, 0x30, 0x5e, 0x3d, 0x1d, 0x5f, 0x5d, 0x76, 0xd5, 0xf1, 0x7d, 0x4a, 0xc6, 0x81, 0x11,
	0x55, 0x1e, 0x68, 0x15, 0xb2, 0x06, 0x6b, 0x2e, 0xf7, 0x02, 0xad, 0xc4, 0x36, 0x5f, 0x31, 0x43,
	0xb4, 0x31, 0x61, 0x8c, 0xa4, 0x20, 0xd6, 0x2a, 0x0c, 0x7c, 0x71, 0x69, 0x92, 0x3e, 0xc7, 0x33,
	0x52, 0x70, 0xa2, 0x44, 0x9d, 0xff, 0x1f, 0x5a, 0x30, 0x98, 0xdc, 0x40, 0x58, 0x68, 0xb0, 0x66,
	0xb9, 0x97, 0xc7, 0x93, 0x5a, 0x9a, 0x58, 0x7b, 0x47, 0x46, 0x85, 0xff, 0xb2, 0xda, 0x67, 0x2c,
	0x36, 0x78, 0x65, 0xa0, 0xad, 0x03, 0x8c, 0x27, 0xfe, 0x61, 0xd1, 0x4b, 0xf3, 0x2c, 0x75, 0x41,
	0xc6, 0x45, 0x63, 0x1e, 0x76, 0x94, 0x9a, 0x86, 0xec, 0xc1, 0xed, 0x10, 0xac, 0x9b, 0xc7, 0xc6,
	0xfe, 0x60, 0x0b, 0x16, 0xb0, 0x15, 0xa6, 0xd9, 0xa2, 0x2e, 0xaf, 0xcd, 0xb1, 0xb6, 0x29, 0xa1,
	0x05, 0x21, 0x79, 0xd1, 0xef, 0xdc, 0x3b, 0x56, 0xda, 0xb5, 0x56, 0xbe, 0xf3, 0xd6, 0x8f, 0x89,
	0xac, 0xaf, 0xfd, 0xcc, 0x78, 0xa9, 0xe3, 0x7f, 0x49, 0x3c, 0x32, 0xbe, 0x3a, 0xa3, 0x2c, 0x36,
	0xbf, 0x49, 0xfc, 0xf6, 0xe4, 0xfa, 0xd6, 0xe2, 0xa6, 0x0c, 0x2e, 0x3a, 0x78, 0x78, 0x7d, 0x7b,
	0x0a, 0xf6, 0xa2, 0x1d, 0x39, 0xda, 0x97, 0x5f, 0x47, 0x92, 0xd9, 0xcb, 0x99, 0xb9, 0x23, 0xb6,
	0x7b, 0x5d, 0xf2, 0x17, 0x71, 0xf8, 0x1e, 0x00, 0x00, 0xff, 0xff, 0x37, 0x54, 0x26, 0x40, 0x54,
	0x02, 0x00, 0x00,
}
