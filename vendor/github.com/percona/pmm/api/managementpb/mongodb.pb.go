// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/mongodb.proto

package managementpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
	// Node identifier on which a service is been running. Required.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// Node and Service access address (DNS name or IP). Required.
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port. Required.
	Port uint32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,5,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,7,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,8,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// MongoDB username for exporter and QAN agent access.
	Username string `protobuf:"bytes,9,opt,name=username,proto3" json:"username,omitempty"`
	// MongoDB password for exporter and QAN agent access.
	Password string `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds mongodb_exporter for provided service.
	MongodbExporter bool `protobuf:"varint,11,opt,name=mongodb_exporter,json=mongodbExporter,proto3" json:"mongodb_exporter,omitempty"`
	// If true, adds qan-mongodb-profiler-agent for provided service.
	QanMongodbProfiler bool `protobuf:"varint,12,opt,name=qan_mongodb_profiler,json=qanMongodbProfiler,proto3" json:"qan_mongodb_profiler,omitempty"`
	// Custom user-assigned labels.
	CustomLabels         map[string]string `protobuf:"bytes,20,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
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

func (m *AddMongoDBRequest) GetMongodbExporter() bool {
	if m != nil {
		return m.MongodbExporter
	}
	return false
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

func init() { proto.RegisterFile("managementpb/mongodb.proto", fileDescriptor_593aa4f9c0b43a5e) }

var fileDescriptor_593aa4f9c0b43a5e = []byte{
	// 670 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x49, 0xbb, 0xb5, 0x9b, 0xdb, 0xb1, 0xcd, 0xda, 0x21, 0x44, 0xc0, 0xa2, 0x22, 0xb4,
	0x0e, 0x68, 0x02, 0x9b, 0x84, 0xd0, 0x2e, 0xa8, 0x83, 0x1e, 0x26, 0x6d, 0x13, 0x64, 0x3b, 0x20,
	0x2e, 0x91, 0x5b, 0xbf, 0x85, 0x68, 0x89, 0x9d, 0xd9, 0x6e, 0xcb, 0x0e, 0x5c, 0xf8, 0x08, 0xf0,
	0xd1, 0xb8, 0x70, 0x43, 0x42, 0x88, 0xcf, 0x81, 0xe2, 0x38, 0x6b, 0xa6, 0x6e, 0x9c, 0x1a, 0xbf,
	0xff, 0xcf, 0x7f, 0xf7, 0xfd, 0xfd, 0x64, 0xe4, 0xa4, 0x84, 0x91, 0x08, 0x52, 0x60, 0x2a, 0x1b,
	0xfa, 0x29, 0x67, 0x11, 0xa7, 0x43, 0x2f, 0x13, 0x5c, 0x71, 0x8c, 0x66, 0x9a, 0xf3, 0x32, 0x8a,
	0xd5, 0xa7, 0xf1, 0xd0, 0x1b, 0xf1, 0xd4, 0x4f, 0xa7, 0xb1, 0x3a, 0xe7, 0x53, 0x3f, 0xe2, 0x3d,
	0x0d, 0xf6, 0x26, 0x24, 0x89, 0x29, 0x51, 0x5c, 0x48, 0xff, 0xea, 0xb3, 0xf0, 0x70, 0xee, 0x47,
	0x9c, 0x47, 0x09, 0xf8, 0x24, 0x8b, 0x7d, 0xc2, 0x18, 0x57, 0x44, 0xc5, 0x9c, 0x49, 0xa3, 0xda,
	0x31, 0x9b, 0x00, 0x53, 0x5c, 0x5c, 0x66, 0x43, 0x9f, 0x44, 0xc0, 0x54, 0xa9, 0x38, 0x55, 0x45,
	0x82, 0x98, 0xc4, 0x23, 0x28, 0xb5, 0x67, 0xfa, 0x67, 0xd4, 0x8b, 0x80, 0xf5, 0xe4, 0x94, 0x44,
	0x11, 0x08, 0x9f, 0x67, 0xda, 0x77, 0xfe, 0x8c, 0xce, 0xcf, 0x05, 0xb4, 0xde, 0xa7, 0xf4, 0x28,
	0x6f, 0xed, 0xed, 0x7e, 0x00, 0x17, 0x63, 0x90, 0x0a, 0x6f, 0xa2, 0x26, 0xe3, 0x14, 0xc2, 0x98,
	0xda, 0x96, 0x6b, 0x75, 0x97, 0xf7, 0x1b, 0xbf, 0x7f, 0x6d, 0xd6, 0x3e, 0x58, 0x41, 0x23, 0x2f,
	0x1f, 0x50, 0xbc, 0x8d, 0xda, 0xe6, 0xd8, 0x90, 0x91, 0x14, 0xec, 0xda, 0x35, 0xaa, 0x65, 0xb4,
	0x63, 0x92, 0x02, 0x76, 0x51, 0x93, 0x50, 0x2a, 0x40, 0x4a, 0xbb, 0x7e, 0x8d, 0x2a, 0xcb, 0xd8,
	0x41, 0x0b, 0x19, 0x17, 0xca, 0x5e, 0x70, 0xad, 0xee, 0x4a, 0x21, 0xaf, 0xdd, 0x09, 0x74, 0x0d,
	0x77, 0x51, 0x3b, 0x4b, 0xd3, 0x50, 0x77, 0x9f, 0xff, 0x9d, 0xc5, 0x6b, 0x16, 0x28, 0x4b, 0xd3,
	0x7e, 0x2e, 0x1d, 0x50, 0xec, 0xa2, 0x16, 0xb0, 0x49, 0x2c, 0x38, 0xcb, 0xaf, 0xc4, 0x6e, 0xe4,
	0x60, 0x50, 0x2d, 0x61, 0x1b, 0x35, 0x47, 0xc9, 0x58, 0x2a, 0x10, 0x76, 0x53, 0xab, 0xe5, 0x12,
	0x6f, 0xa1, 0x55, 0x01, 0x59, 0x12, 0x8f, 0x74, 0x36, 0xa1, 0x04, 0x65, 0x2f, 0x69, 0xe2, 0x6e,
	0xa5, 0x7c, 0x02, 0x0a, 0x3b, 0x68, 0x69, 0x2c, 0x41, 0xe8, 0x9e, 0x97, 0x35, 0x71, 0xb5, 0xce,
	0xb5, 0x8c, 0x48, 0x39, 0xe5, 0x82, 0xda, 0xa8, 0xd0, 0xca, 0x35, 0xde, 0x46, 0x6b, 0x66, 0x7a,
	0x42, 0xf8, 0x9c, 0x37, 0x06, 0xc2, 0x6e, 0xb9, 0x56, 0x77, 0x29, 0x58, 0x35, 0xf5, 0x81, 0x29,
	0xe3, 0xe7, 0x68, 0xe3, 0x82, 0xb0, 0xb0, 0xc4, 0x33, 0xc1, 0xcf, 0xe2, 0x04, 0x84, 0xdd, 0xd6,
	0x38, 0xbe, 0x20, 0xec, 0xa8, 0x90, 0xde, 0x19, 0x05, 0x9f, 0xa2, 0x95, 0xd1, 0x58, 0x2a, 0x9e,
	0x86, 0x09, 0x19, 0x42, 0x22, 0xed, 0x0d, 0xb7, 0xde, 0x6d, 0xed, 0xf8, 0xde, 0x6c, 0x42, 0xbd,
	0xb9, 0x3b, 0xf6, 0xde, 0xe8, 0x2d, 0x87, 0x7a, 0xc7, 0x80, 0x29, 0x71, 0x19, 0xb4, 0x47, 0x95,
	0x92, 0xf3, 0x1a, 0xad, 0xcf, 0x21, 0x78, 0x0d, 0xd5, 0xcf, 0xe1, 0xb2, 0x18, 0x8a, 0x20, 0xff,
	0xc4, 0x1b, 0x68, 0x71, 0x42, 0x92, 0xb1, 0x19, 0x81, 0xa0, 0x58, 0xec, 0xd5, 0x5e, 0x59, 0x9d,
	0xbf, 0x16, 0xc2, 0xd5, 0x63, 0x65, 0xc6, 0x99, 0x04, 0xbc, 0x8b, 0x9a, 0x66, 0x3c, 0xb4, 0x4d,
	0x6b, 0xe7, 0x9e, 0x77, 0x35, 0xcd, 0x9e, 0x81, 0x4f, 0x0a, 0x20, 0x28, 0x49, 0x3c, 0xb8, 0x21,
	0xbf, 0x9a, 0xde, 0xed, 0xcc, 0xef, 0x2e, 0xa3, 0x9c, 0xcf, 0xf6, 0xf4, 0x96, 0x6c, 0xeb, 0xda,
	0xaa, 0x53, 0xb1, 0x7a, 0xdf, 0x3f, 0x36, 0x6e, 0x65, 0xcc, 0x7a, 0xce, 0x6e, 0xca, 0x7f, 0xe7,
	0x0b, 0x6a, 0x1a, 0x16, 0x0b, 0x84, 0x66, 0x2d, 0xe3, 0x07, 0xff, 0xbd, 0x01, 0xe7, 0xe1, 0x6d,
	0x72, 0x91, 0x54, 0xe7, 0xf1, 0xd7, 0x1f, 0x7f, 0xbe, 0xd7, 0x36, 0x3b, 0x8e, 0x3f, 0x79, 0xe1,
	0xcf, 0x50, 0xdf, 0x70, 0x7e, 0x9f, 0xd2, 0x3d, 0xeb, 0xc9, 0x7e, 0xf8, 0xad, 0x7f, 0x1c, 0x1c,
	0xa2, 0x26, 0x85, 0x33, 0x32, 0x4e, 0x14, 0xee, 0x23, 0xdc, 0x67, 0x2e, 0x08, 0xc1, 0x85, 0x2b,
	0x8c, 0x97, 0x87, 0x9f, 0xa2, 0x6d, 0x67, 0xeb, 0x91, 0x4f, 0xe1, 0x2c, 0x66, 0x71, 0xf1, 0x1a,
	0x54, 0x1f, 0xb7, 0x41, 0x8e, 0x97, 0x27, 0x7f, 0x6c, 0x57, 0xa5, 0x61, 0x43, 0x3f, 0x15, 0xbb,
	0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x11, 0x8b, 0xa5, 0x0e, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

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
	cc *grpc.ClientConn
}

func NewMongoDBClient(cc *grpc.ClientConn) MongoDBClient {
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
