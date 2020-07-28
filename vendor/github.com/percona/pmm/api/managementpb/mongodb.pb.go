// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/mongodb.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type AddMongoDBRequest struct {
	// Node identifier on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,2,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,3,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,4,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP).
	// Address (and port) or socket is required.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port.
	// Port is required when the address present.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// Service Access socket.
	// Address (and port) or socket is required.
	Socket string `protobuf:"bytes,19,opt,name=socket,proto3" json:"socket,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,7,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,8,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,9,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,10,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// MongoDB username for exporter and QAN agent access.
	Username string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	// MongoDB password for exporter and QAN agent access.
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds qan-mongodb-profiler-agent for provided service.
	QanMongodbProfiler bool `protobuf:"varint,13,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,14,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,15,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,17,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation.
	TlsSkipVerify        bool     `protobuf:"varint,18,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddMongoDBRequest) Reset()         { *m = AddMongoDBRequest{} }
func (m *AddMongoDBRequest) String() string { return proto.CompactTextString(m) }
func (*AddMongoDBRequest) ProtoMessage()    {}
func (*AddMongoDBRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_593aa4f9c0b43a5e, []int{0}
}

func (m *AddMongoDBRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMongoDBRequest.Unmarshal(m, b)
}
func (m *AddMongoDBRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMongoDBRequest.Marshal(b, m, deterministic)
}
func (m *AddMongoDBRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMongoDBRequest.Merge(m, src)
}
func (m *AddMongoDBRequest) XXX_Size() int {
	return xxx_messageInfo_AddMongoDBRequest.Size(m)
}
func (m *AddMongoDBRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMongoDBRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddMongoDBRequest proto.InternalMessageInfo

func (m *AddMongoDBRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddMongoDBRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AddMongoDBRequest) GetAddNode() *AddNodeParams {
	if m != nil {
		return m.AddNode
	}
	return nil
}

func (m *AddMongoDBRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddMongoDBRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddMongoDBRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddMongoDBRequest) GetSocket() string {
	if m != nil {
		return m.Socket
	}
	return ""
}

func (m *AddMongoDBRequest) GetPmmAgentId() string {
	if m != nil {
		return m.PmmAgentId
	}
	return ""
}

func (m *AddMongoDBRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AddMongoDBRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *AddMongoDBRequest) GetReplicationSet() string {
	if m != nil {
		return m.ReplicationSet
	}
	return ""
}

func (m *AddMongoDBRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddMongoDBRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddMongoDBRequest) GetQanMongodbProfiler() bool {
	if m != nil {
		return m.QanMongodbProfiler
	}
	return false
}

func (m *AddMongoDBRequest) GetCustomLabels() map[string]string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

func (m *AddMongoDBRequest) GetSkipConnectionCheck() bool {
	if m != nil {
		return m.SkipConnectionCheck
	}
	return false
}

func (m *AddMongoDBRequest) GetTls() bool {
	if m != nil {
		return m.Tls
	}
	return false
}

func (m *AddMongoDBRequest) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

type AddMongoDBResponse struct {
	Service              *inventorypb.MongoDBService          `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	MongodbExporter      *inventorypb.MongoDBExporter         `protobuf:"bytes,2,opt,name=mongodb_exporter,json=mongodbExporter,proto3" json:"mongodb_exporter,omitempty"`
	QanMongodbProfiler   *inventorypb.QANMongoDBProfilerAgent `protobuf:"bytes,3,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *AddMongoDBResponse) Reset()         { *m = AddMongoDBResponse{} }
func (m *AddMongoDBResponse) String() string { return proto.CompactTextString(m) }
func (*AddMongoDBResponse) ProtoMessage()    {}
func (*AddMongoDBResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_593aa4f9c0b43a5e, []int{1}
}

func (m *AddMongoDBResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddMongoDBResponse.Unmarshal(m, b)
}
func (m *AddMongoDBResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddMongoDBResponse.Marshal(b, m, deterministic)
}
func (m *AddMongoDBResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddMongoDBResponse.Merge(m, src)
}
func (m *AddMongoDBResponse) XXX_Size() int {
	return xxx_messageInfo_AddMongoDBResponse.Size(m)
}
func (m *AddMongoDBResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddMongoDBResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddMongoDBResponse proto.InternalMessageInfo

func (m *AddMongoDBResponse) GetService() *inventorypb.MongoDBService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddMongoDBResponse) GetMongodbExporter() *inventorypb.MongoDBExporter {
	if m != nil {
		return m.MongodbExporter
	}
	return nil
}

func (m *AddMongoDBResponse) GetQanMongodbProfiler() *inventorypb.QANMongoDBProfilerAgent {
	if m != nil {
		return m.QanMongodbProfiler
	}
	return nil
}

func init() {
	proto.RegisterType((*AddMongoDBRequest)(nil), "management.AddMongoDBRequest")
	proto.RegisterMapType((map[string]string)(nil), "management.AddMongoDBRequest.CustomLabelsEntry")
	proto.RegisterType((*AddMongoDBResponse)(nil), "management.AddMongoDBResponse")
}

func init() {
	proto.RegisterFile("managementpb/mongodb.proto", fileDescriptor_593aa4f9c0b43a5e)
}

var fileDescriptor_593aa4f9c0b43a5e = []byte{
	// 720 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xc7, 0xe5, 0xb4, 0x4d, 0xd2, 0x49, 0xd2, 0xb4, 0xd3, 0xde, 0x7b, 0xe7, 0xfa, 0x7e, 0x34,
	0x8a, 0x74, 0x2f, 0x01, 0xa9, 0x31, 0xa4, 0x08, 0xa1, 0x6e, 0x50, 0x5a, 0xba, 0xa8, 0x44, 0xab,
	0xe2, 0x56, 0x08, 0xb1, 0xb1, 0x26, 0x9e, 0xd3, 0xd4, 0x8a, 0x3d, 0xe3, 0xce, 0x4c, 0x52, 0xb2,
	0x60, 0xc3, 0x2b, 0xf0, 0x68, 0x3c, 0x00, 0x12, 0x42, 0x2c, 0x79, 0x06, 0x34, 0x63, 0x3b, 0x49,
	0x29, 0x65, 0x37, 0xe7, 0xfc, 0xfe, 0x73, 0x3e, 0xe6, 0x1c, 0x1b, 0xb9, 0x09, 0xe5, 0x74, 0x08,
	0x09, 0x70, 0x9d, 0x0e, 0xbc, 0x44, 0xf0, 0xa1, 0x60, 0x83, 0x6e, 0x2a, 0x85, 0x16, 0x18, 0xcd,
	0x99, 0xfb, 0x64, 0x18, 0xe9, 0xcb, 0xf1, 0xa0, 0x1b, 0x8a, 0xc4, 0x4b, 0xae, 0x23, 0x3d, 0x12,
	0xd7, 0xde, 0x50, 0xec, 0x58, 0xe1, 0xce, 0x84, 0xc6, 0x11, 0xa3, 0x5a, 0x48, 0xe5, 0xcd, 0x8e,
	0x59, 0x0c, 0xf7, 0xef, 0xa1, 0x10, 0xc3, 0x18, 0x3c, 0x9a, 0x46, 0x1e, 0xe5, 0x5c, 0x68, 0xaa,
	0x23, 0xc1, 0x55, 0x4e, 0x49, 0xc4, 0x27, 0xc0, 0xb5, 0x90, 0xd3, 0x74, 0xe0, 0xd1, 0x21, 0x70,
	0x5d, 0x10, 0x77, 0x91, 0x28, 0x90, 0x93, 0x28, 0x84, 0x19, 0xbb, 0x51, 0x73, 0x0e, 0x33, 0xd6,
	0xfe, 0xb6, 0x82, 0x36, 0xfa, 0x8c, 0x1d, 0x9b, 0x46, 0x9e, 0xef, 0xfb, 0x70, 0x35, 0x06, 0xa5,
	0xf1, 0x1f, 0xa8, 0xc2, 0x05, 0x83, 0x20, 0x62, 0xc4, 0x69, 0x39, 0x9d, 0x55, 0xbf, 0x6c, 0xcc,
	0x23, 0x86, 0xff, 0x42, 0xab, 0x16, 0x70, 0x9a, 0x00, 0x29, 0x59, 0x54, 0x35, 0x8e, 0x13, 0x9a,
	0x00, 0x7e, 0x8c, 0xaa, 0x94, 0xb1, 0xc0, 0xd8, 0x64, 0xa9, 0xe5, 0x74, 0x6a, 0xbd, 0x3f, 0xbb,
	0xf3, 0xd4, 0xdd, 0x3e, 0x63, 0x27, 0x82, 0xc1, 0x29, 0x95, 0x34, 0x51, 0x7e, 0x85, 0x66, 0x26,
	0xbe, 0x8f, 0xea, 0x79, 0x49, 0x59, 0xd4, 0x65, 0x13, 0x75, 0xbf, 0xfc, 0xf9, 0xd3, 0x76, 0xe9,
	0xb5, 0xe3, 0xd7, 0x72, 0x66, 0x13, 0x10, 0x64, 0x6e, 0x49, 0x50, 0x8a, 0xac, 0xd8, 0xdc, 0x85,
	0x89, 0x31, 0x5a, 0x4e, 0x85, 0xd4, 0xa4, 0xdc, 0x72, 0x3a, 0x0d, 0xdf, 0x9e, 0xf1, 0xef, 0xa8,
	0xac, 0x44, 0x38, 0x02, 0x4d, 0x36, 0xb3, 0x1e, 0x32, 0x0b, 0x77, 0x50, 0x3d, 0x4d, 0x92, 0xc0,
	0x3e, 0x9f, 0xe9, 0xb0, 0x72, 0x23, 0x21, 0x4a, 0x93, 0xa4, 0x6f, 0xd0, 0x11, 0xc3, 0x2d, 0x54,
	0x03, 0x3e, 0x89, 0xa4, 0xe0, 0xa6, 0x01, 0x52, 0xb5, 0x61, 0x16, 0x5d, 0xa6, 0xa2, 0x30, 0x1e,
	0x2b, 0x0d, 0x92, 0xac, 0x66, 0x15, 0xe5, 0x26, 0xbe, 0x87, 0x9a, 0x12, 0xd2, 0x38, 0x0a, 0xed,
	0x00, 0x03, 0x05, 0x9a, 0x20, 0xab, 0x58, 0x5b, 0x70, 0x9f, 0x81, 0xc6, 0x2e, 0xaa, 0x8e, 0x15,
	0x48, 0xdb, 0x7b, 0x2d, 0x7b, 0xd1, 0xc2, 0x36, 0x2c, 0xa5, 0x4a, 0x5d, 0x0b, 0xc9, 0x48, 0x3d,
	0x63, 0x85, 0x8d, 0x1f, 0xa2, 0xad, 0x2b, 0xca, 0x83, 0x7c, 0x05, 0x83, 0x54, 0x8a, 0x8b, 0x28,
	0x06, 0x49, 0x1a, 0x2d, 0xa7, 0x53, 0xf5, 0xf1, 0x15, 0xe5, 0xc7, 0x19, 0x3a, 0xcd, 0x09, 0x3e,
	0x47, 0x8d, 0x70, 0xac, 0xb4, 0x48, 0x82, 0x98, 0x0e, 0x20, 0x56, 0x64, 0xad, 0xb5, 0xd4, 0xa9,
	0xf5, 0xbc, 0x1f, 0x86, 0x74, 0x73, 0x17, 0xba, 0x07, 0xf6, 0xca, 0x0b, 0x7b, 0xe3, 0x90, 0x6b,
	0x39, 0xf5, 0xeb, 0xe1, 0x82, 0x0b, 0xf7, 0xd0, 0x6f, 0x6a, 0x14, 0xa5, 0x41, 0x28, 0x38, 0x87,
	0xd0, 0x36, 0x1b, 0x5e, 0x42, 0x38, 0x22, 0x4d, 0x5b, 0xc8, 0xa6, 0x81, 0x07, 0x33, 0x76, 0x60,
	0x10, 0x5e, 0x47, 0x4b, 0x3a, 0x56, 0x64, 0xc3, 0x2a, 0xcc, 0x11, 0xff, 0x8f, 0x9a, 0x3a, 0x56,
	0x81, 0x8d, 0x34, 0x01, 0x19, 0x5d, 0x4c, 0x09, 0xb6, 0xb4, 0xa1, 0x63, 0x75, 0x36, 0x8a, 0xd2,
	0x57, 0xd6, 0xe9, 0x3e, 0x43, 0x1b, 0xb7, 0x0a, 0x32, 0xe1, 0x46, 0x30, 0xcd, 0x57, 0xd5, 0x1c,
	0xf1, 0x16, 0x5a, 0x99, 0xd0, 0x78, 0x5c, 0xec, 0x68, 0x66, 0xec, 0x95, 0x9e, 0x3a, 0xed, 0xaf,
	0x0e, 0xc2, 0x8b, 0x4d, 0xaa, 0x54, 0x70, 0x05, 0x78, 0x17, 0x55, 0xf2, 0x4d, 0xb3, 0x61, 0xcc,
	0xea, 0xce, 0xbe, 0xa8, 0x6e, 0x2e, 0x3e, 0xcb, 0x04, 0x7e, 0xa1, 0xc4, 0x87, 0x68, 0xbd, 0x78,
	0x7e, 0x78, 0x6b, 0x76, 0x0e, 0xa4, 0x4d, 0x58, 0xeb, 0xb9, 0xb7, 0x6f, 0x1f, 0xe6, 0x0a, 0xbf,
	0x99, 0xdf, 0x29, 0x1c, 0xf8, 0xfc, 0x8e, 0x49, 0x66, 0xdf, 0x50, 0x7b, 0x21, 0xd4, 0xcb, 0xfe,
	0x49, 0x1e, 0xad, 0x18, 0xaa, 0x5d, 0xd5, 0x9f, 0x4d, 0xbb, 0xf7, 0x0e, 0x55, 0x72, 0x2d, 0x96,
	0x08, 0xcd, 0x5b, 0xc6, 0xff, 0xfc, 0x72, 0xde, 0xee, 0xbf, 0x77, 0xe1, 0xec, 0xa5, 0xda, 0xff,
	0xbd, 0xff, 0xf8, 0xe5, 0x43, 0x69, 0xbb, 0xed, 0x7a, 0x93, 0x47, 0xde, 0x5c, 0xea, 0xe5, 0x3a,
	0xaf, 0xcf, 0xd8, 0x9e, 0xf3, 0x60, 0x7f, 0xed, 0x4d, 0x7d, 0xf1, 0xb7, 0x33, 0x28, 0xdb, 0xff,
	0xcd, 0xee, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xea, 0x5c, 0xf8, 0xa6, 0x41, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MongoDBClient is the client API for MongoDB service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MongoDBClient interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error)
}

type mongoDBClient struct {
	cc grpc.ClientConnInterface
}

func NewMongoDBClient(cc grpc.ClientConnInterface) MongoDBClient {
	return &mongoDBClient{cc}
}

func (c *mongoDBClient) AddMongoDB(ctx context.Context, in *AddMongoDBRequest, opts ...grpc.CallOption) (*AddMongoDBResponse, error) {
	out := new(AddMongoDBResponse)
	err := c.cc.Invoke(ctx, "/management.MongoDB/AddMongoDB", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MongoDBServer is the server API for MongoDB service.
type MongoDBServer interface {
	// AddMongoDB adds MongoDB Service and starts several Agents.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "mongodb_exporter", and "qan_mongodb_profiler" agents
	// with provided "pmm_agent_id" and other parameters.
	AddMongoDB(context.Context, *AddMongoDBRequest) (*AddMongoDBResponse, error)
}

// UnimplementedMongoDBServer can be embedded to have forward compatible implementations.
type UnimplementedMongoDBServer struct {
}

func (*UnimplementedMongoDBServer) AddMongoDB(ctx context.Context, req *AddMongoDBRequest) (*AddMongoDBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMongoDB not implemented")
}

func RegisterMongoDBServer(s *grpc.Server, srv MongoDBServer) {
	s.RegisterService(&_MongoDB_serviceDesc, srv)
}

func _MongoDB_AddMongoDB_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMongoDBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MongoDBServer).AddMongoDB(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.MongoDB/AddMongoDB",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MongoDBServer).AddMongoDB(ctx, req.(*AddMongoDBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MongoDB_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.MongoDB",
	HandlerType: (*MongoDBServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMongoDB",
			Handler:    _MongoDB_AddMongoDB_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/mongodb.proto",
}
