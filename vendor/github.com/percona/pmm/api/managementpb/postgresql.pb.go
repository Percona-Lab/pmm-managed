// Code generated by protoc-gen-go. DO NOT EDIT.
// source: managementpb/postgresql.proto

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

type AddPostgreSQLRequest struct {
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
	// Node and Service access address (DNS name or IP). Required.
	Address string `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	// Service Access port. Required.
	Port uint32 `protobuf:"varint,6,opt,name=port,proto3" json:"port,omitempty"`
	// The "pmm-agent" identifier which should run agents. Required.
	PmmAgentId string `protobuf:"bytes,7,opt,name=pmm_agent_id,json=pmmAgentId,proto3" json:"pmm_agent_id,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,8,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,9,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,10,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// PostgreSQL username for scraping metrics.
	Username string `protobuf:"bytes,11,opt,name=username,proto3" json:"username,omitempty"`
	// PostgreSQL password for scraping metrics.
	Password string `protobuf:"bytes,12,opt,name=password,proto3" json:"password,omitempty"`
	// If true, adds qan-postgresql-pgstatements-agent for provided service.
	QanPostgresqlPgstatementsAgent bool `protobuf:"varint,13,opt,name=qan_postgresql_pgstatements_agent,json=qanPostgresqlPgstatementsAgent,proto3" json:"qan_postgresql_pgstatements_agent,omitempty"`
	// Custom user-assigned labels.
	CustomLabels map[string]string `protobuf:"bytes,14,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Skip connection check.
	SkipConnectionCheck bool `protobuf:"varint,15,opt,name=skip_connection_check,json=skipConnectionCheck,proto3" json:"skip_connection_check,omitempty"`
	// Use TLS for database connections.
	Tls bool `protobuf:"varint,16,opt,name=tls,proto3" json:"tls,omitempty"`
	// Skip TLS certificate and hostname validation. Uses sslmode=required instead of verify-full.
	TlsSkipVerify        bool     `protobuf:"varint,17,opt,name=tls_skip_verify,json=tlsSkipVerify,proto3" json:"tls_skip_verify,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddPostgreSQLRequest) Reset()         { *m = AddPostgreSQLRequest{} }
func (m *AddPostgreSQLRequest) String() string { return proto.CompactTextString(m) }
func (*AddPostgreSQLRequest) ProtoMessage()    {}
func (*AddPostgreSQLRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e72a2ebc60b1270, []int{0}
}

func (m *AddPostgreSQLRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPostgreSQLRequest.Unmarshal(m, b)
}
func (m *AddPostgreSQLRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPostgreSQLRequest.Marshal(b, m, deterministic)
}
func (m *AddPostgreSQLRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPostgreSQLRequest.Merge(m, src)
}
func (m *AddPostgreSQLRequest) XXX_Size() int {
	return xxx_messageInfo_AddPostgreSQLRequest.Size(m)
}
func (m *AddPostgreSQLRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPostgreSQLRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddPostgreSQLRequest proto.InternalMessageInfo

func (m *AddPostgreSQLRequest) GetNodeId() string {
	if m != nil {
		return m.NodeId
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetAddNode() *AddNodeParams {
	if m != nil {
		return m.AddNode
	}
	return nil
}

func (m *AddPostgreSQLRequest) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *AddPostgreSQLRequest) GetPmmAgentId() string {
	if m != nil {
		return m.PmmAgentId
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetReplicationSet() string {
	if m != nil {
		return m.ReplicationSet
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *AddPostgreSQLRequest) GetQanPostgresqlPgstatementsAgent() bool {
	if m != nil {
		return m.QanPostgresqlPgstatementsAgent
	}
	return false
}

func (m *AddPostgreSQLRequest) GetCustomLabels() map[string]string {
	if m != nil {
		return m.CustomLabels
	}
	return nil
}

func (m *AddPostgreSQLRequest) GetSkipConnectionCheck() bool {
	if m != nil {
		return m.SkipConnectionCheck
	}
	return false
}

func (m *AddPostgreSQLRequest) GetTls() bool {
	if m != nil {
		return m.Tls
	}
	return false
}

func (m *AddPostgreSQLRequest) GetTlsSkipVerify() bool {
	if m != nil {
		return m.TlsSkipVerify
	}
	return false
}

type AddPostgreSQLResponse struct {
	Service                        *inventorypb.PostgreSQLService              `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	PostgresExporter               *inventorypb.PostgresExporter               `protobuf:"bytes,2,opt,name=postgres_exporter,json=postgresExporter,proto3" json:"postgres_exporter,omitempty"`
	QanPostgresqlPgstatementsAgent *inventorypb.QANPostgreSQLPgStatementsAgent `protobuf:"bytes,3,opt,name=qan_postgresql_pgstatements_agent,json=qanPostgresqlPgstatementsAgent,proto3" json:"qan_postgresql_pgstatements_agent,omitempty"`
	XXX_NoUnkeyedLiteral           struct{}                                    `json:"-"`
	XXX_unrecognized               []byte                                      `json:"-"`
	XXX_sizecache                  int32                                       `json:"-"`
}

func (m *AddPostgreSQLResponse) Reset()         { *m = AddPostgreSQLResponse{} }
func (m *AddPostgreSQLResponse) String() string { return proto.CompactTextString(m) }
func (*AddPostgreSQLResponse) ProtoMessage()    {}
func (*AddPostgreSQLResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6e72a2ebc60b1270, []int{1}
}

func (m *AddPostgreSQLResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddPostgreSQLResponse.Unmarshal(m, b)
}
func (m *AddPostgreSQLResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddPostgreSQLResponse.Marshal(b, m, deterministic)
}
func (m *AddPostgreSQLResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddPostgreSQLResponse.Merge(m, src)
}
func (m *AddPostgreSQLResponse) XXX_Size() int {
	return xxx_messageInfo_AddPostgreSQLResponse.Size(m)
}
func (m *AddPostgreSQLResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddPostgreSQLResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddPostgreSQLResponse proto.InternalMessageInfo

func (m *AddPostgreSQLResponse) GetService() *inventorypb.PostgreSQLService {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *AddPostgreSQLResponse) GetPostgresExporter() *inventorypb.PostgresExporter {
	if m != nil {
		return m.PostgresExporter
	}
	return nil
}

func (m *AddPostgreSQLResponse) GetQanPostgresqlPgstatementsAgent() *inventorypb.QANPostgreSQLPgStatementsAgent {
	if m != nil {
		return m.QanPostgresqlPgstatementsAgent
	}
	return nil
}

func init() {
	proto.RegisterType((*AddPostgreSQLRequest)(nil), "management.AddPostgreSQLRequest")
	proto.RegisterMapType((map[string]string)(nil), "management.AddPostgreSQLRequest.CustomLabelsEntry")
	proto.RegisterType((*AddPostgreSQLResponse)(nil), "management.AddPostgreSQLResponse")
}

func init() { proto.RegisterFile("managementpb/postgresql.proto", fileDescriptor_6e72a2ebc60b1270) }

var fileDescriptor_6e72a2ebc60b1270 = []byte{
	// 795 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x5d, 0x6f, 0x1b, 0x45,
	0x14, 0x65, 0x9d, 0xd6, 0x76, 0xae, 0xed, 0x26, 0x19, 0x5a, 0x31, 0xb8, 0x2d, 0x6c, 0x8d, 0x44,
	0x1d, 0x20, 0x5e, 0x61, 0x50, 0x85, 0xfa, 0x82, 0xdc, 0x28, 0x12, 0x91, 0xa2, 0xc8, 0x5d, 0x4b,
	0x80, 0x78, 0x59, 0x8d, 0x77, 0x6e, 0xb6, 0x2b, 0xef, 0xce, 0x4c, 0x66, 0xc6, 0x0e, 0x79, 0xe1,
	0x81, 0x67, 0x10, 0x12, 0xfc, 0x34, 0x7e, 0x00, 0x12, 0xe2, 0x87, 0xa0, 0x9d, 0x5d, 0x7f, 0xc5,
	0x12, 0xe5, 0x69, 0xe7, 0x9e, 0x73, 0xe6, 0xcc, 0xbd, 0x77, 0xee, 0x0e, 0x3c, 0xcd, 0x99, 0x60,
	0x09, 0xe6, 0x28, 0xac, 0x9a, 0x06, 0x4a, 0x1a, 0x9b, 0x68, 0x34, 0xd7, 0xd9, 0x40, 0x69, 0x69,
	0x25, 0x81, 0x35, 0xdd, 0x7d, 0x91, 0xa4, 0xf6, 0xcd, 0x7c, 0x3a, 0x88, 0x65, 0x1e, 0xe4, 0x37,
	0xa9, 0x9d, 0xc9, 0x9b, 0x20, 0x91, 0x27, 0x4e, 0x78, 0xb2, 0x60, 0x59, 0xca, 0x99, 0x95, 0xda,
	0x04, 0xab, 0x65, 0xe9, 0xd1, 0x7d, 0x92, 0x48, 0x99, 0x64, 0x18, 0x30, 0x95, 0x06, 0x4c, 0x08,
	0x69, 0x99, 0x4d, 0xa5, 0x30, 0x15, 0x4b, 0x53, 0xb1, 0x40, 0x61, 0xa5, 0xbe, 0x55, 0xd3, 0x80,
	0x25, 0x28, 0xec, 0x92, 0xe9, 0x6e, 0x32, 0x06, 0xf5, 0x22, 0x8d, 0x71, 0xc5, 0x6d, 0xa5, 0x5d,
	0x91, 0x15, 0xf7, 0x99, 0xfb, 0xc4, 0x27, 0x09, 0x8a, 0x13, 0x73, 0xc3, 0x92, 0x04, 0x75, 0x20,
	0x95, 0x3b, 0x73, 0xf7, 0xfc, 0xde, 0xaf, 0x75, 0x78, 0x38, 0xe2, 0x7c, 0x5c, 0x56, 0x3e, 0x79,
	0x7d, 0x11, 0xe2, 0xf5, 0x1c, 0x8d, 0x25, 0xef, 0x41, 0x43, 0x48, 0x8e, 0x51, 0xca, 0xa9, 0xe7,
	0x7b, 0xfd, 0xfd, 0xb0, 0x5e, 0x84, 0xe7, 0x9c, 0x3c, 0x86, 0x7d, 0x47, 0x08, 0x96, 0x23, 0xad,
	0x39, 0xaa, 0x59, 0x00, 0x97, 0x2c, 0x47, 0xf2, 0x25, 0x34, 0x19, 0xe7, 0x51, 0x11, 0xd3, 0x3d,
	0xdf, 0xeb, 0xb7, 0x86, 0xef, 0x0f, 0xd6, 0xb9, 0x0e, 0x46, 0x9c, 0x5f, 0x4a, 0x8e, 0x63, 0xa6,
	0x59, 0x6e, 0xc2, 0x06, 0x2b, 0x43, 0x72, 0x0c, 0xed, 0xaa, 0x86, 0xd2, 0xf5, 0x5e, 0xe1, 0xfa,
	0xaa, 0xfe, 0xf7, 0x5f, 0x1f, 0xd6, 0xbe, 0xf7, 0xc2, 0x56, 0xc5, 0xb9, 0x03, 0x7c, 0x28, 0x76,
	0x69, 0x34, 0x86, 0xde, 0xdf, 0x52, 0x2d, 0x61, 0xd2, 0x85, 0x7b, 0x4a, 0x6a, 0x4b, 0xeb, 0xbe,
	0xd7, 0xef, 0x94, 0xf4, 0xe1, 0x3b, 0xa1, 0xc3, 0x48, 0x1f, 0xda, 0x2a, 0xcf, 0x23, 0xd7, 0xe7,
	0xa2, 0xb2, 0xc6, 0x96, 0x05, 0xa8, 0x3c, 0x1f, 0x15, 0xd4, 0x39, 0x27, 0x3e, 0xb4, 0x50, 0x2c,
	0x52, 0x2d, 0x45, 0x91, 0x38, 0x6d, 0xba, 0x3a, 0x37, 0x21, 0x42, 0xa1, 0x11, 0x67, 0x73, 0x63,
	0x51, 0xd3, 0x7d, 0xc7, 0x2e, 0x43, 0xf2, 0x1c, 0x0e, 0x34, 0xaa, 0x2c, 0x8d, 0x5d, 0xa7, 0x23,
	0x83, 0x96, 0x82, 0x53, 0x3c, 0xd8, 0x80, 0x27, 0x68, 0x49, 0x0f, 0x9a, 0x73, 0x83, 0xda, 0xd5,
	0xdc, 0xda, 0x4a, 0x65, 0x85, 0x93, 0x2e, 0x34, 0x15, 0x33, 0xe6, 0x46, 0x6a, 0x4e, 0xdb, 0x65,
	0xb7, 0x97, 0x31, 0x39, 0x87, 0x67, 0xd7, 0x4c, 0x44, 0xeb, 0xb1, 0x8d, 0x54, 0x62, 0x2c, 0xb3,
	0xae, 0xd9, 0xa6, 0x2c, 0x93, 0x76, 0x7c, 0xaf, 0xdf, 0x0c, 0x3f, 0xb8, 0x66, 0x62, 0xbc, 0xd2,
	0x8d, 0x37, 0x64, 0xae, 0x62, 0xf2, 0x1d, 0x74, 0xe2, 0xb9, 0xb1, 0x32, 0x8f, 0x32, 0x36, 0xc5,
	0xcc, 0xd0, 0x07, 0xfe, 0x5e, 0xbf, 0x35, 0x1c, 0xde, 0xb9, 0xbd, 0x9d, 0x39, 0x19, 0x9c, 0xba,
	0x5d, 0x17, 0x6e, 0xd3, 0x99, 0xb0, 0xfa, 0x36, 0x6c, 0xc7, 0x1b, 0x10, 0x19, 0xc2, 0x23, 0x33,
	0x4b, 0x55, 0x14, 0x4b, 0x21, 0x30, 0x76, 0x0d, 0x89, 0xdf, 0x60, 0x3c, 0xa3, 0x07, 0x2e, 0xaf,
	0x77, 0x0b, 0xf2, 0x74, 0xc5, 0x9d, 0x16, 0x14, 0x39, 0x84, 0x3d, 0x9b, 0x19, 0x7a, 0xe8, 0x14,
	0xc5, 0x92, 0x7c, 0x0c, 0x07, 0x36, 0x33, 0x91, 0x73, 0x5a, 0xa0, 0x4e, 0xaf, 0x6e, 0xe9, 0x91,
	0x63, 0x3b, 0x36, 0x33, 0x93, 0x59, 0xaa, 0xbe, 0x75, 0x60, 0xf7, 0x6b, 0x38, 0xda, 0x49, 0xa8,
	0xb0, 0x9b, 0xe1, 0x6d, 0x35, 0xc6, 0xc5, 0x92, 0x3c, 0x84, 0xfb, 0x0b, 0x96, 0xcd, 0x97, 0xf3,
	0x5b, 0x06, 0x2f, 0x6b, 0x5f, 0x79, 0xbd, 0xdf, 0x6a, 0xf0, 0xe8, 0x4e, 0x9d, 0x46, 0x49, 0x61,
	0x90, 0xbc, 0x80, 0x46, 0x35, 0x88, 0xce, 0xa9, 0x35, 0x7c, 0x32, 0x58, 0xfd, 0xa1, 0x83, 0xb5,
	0x7e, 0x52, 0x6a, 0xc2, 0xa5, 0x98, 0x7c, 0x03, 0x47, 0xcb, 0x0b, 0x8a, 0xf0, 0xc7, 0x62, 0x0c,
	0x51, 0xbb, 0x73, 0x5b, 0xc3, 0xc7, 0xbb, 0x0e, 0xe6, 0xac, 0x92, 0x84, 0x87, 0xea, 0x0e, 0x42,
	0xec, 0xff, 0xb9, 0xee, 0xf2, 0xaf, 0x3b, 0xde, 0x70, 0x7e, 0x3d, 0xba, 0x5c, 0xa7, 0x37, 0x4e,
	0x26, 0xdb, 0x37, 0xff, 0xb6, 0xc9, 0x18, 0xfe, 0xe2, 0x01, 0xac, 0xf7, 0x93, 0x9f, 0xa0, 0xb3,
	0xd5, 0x1f, 0xe2, 0xbf, 0x6d, 0x44, 0xba, 0xcf, 0xfe, 0x43, 0x51, 0x36, 0xb7, 0xd7, 0xff, 0xf9,
	0xcf, 0x7f, 0xfe, 0xa8, 0xf5, 0x7a, 0x4f, 0x83, 0xc5, 0xe7, 0xc1, 0x5a, 0x1d, 0xac, 0xa5, 0xc1,
	0x88, 0xf3, 0x97, 0xde, 0x27, 0xaf, 0xa2, 0xdf, 0x47, 0x97, 0xe1, 0x05, 0x34, 0x38, 0x5e, 0xb1,
	0x79, 0x66, 0xc9, 0x08, 0xc8, 0x48, 0xf8, 0xa8, 0xb5, 0xd4, 0xbe, 0xae, 0xec, 0x06, 0xe4, 0x53,
	0x38, 0xee, 0x3e, 0xff, 0x28, 0xe0, 0x78, 0x95, 0x8a, 0xb4, 0x7c, 0xfb, 0x36, 0x9f, 0xcc, 0xb3,
	0x42, 0xbe, 0x3c, 0xfc, 0x87, 0xf6, 0x26, 0x35, 0xad, 0xbb, 0x87, 0xf1, 0x8b, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x78, 0xfb, 0x87, 0x84, 0x1b, 0x06, 0x00, 0x00,
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
	// AddPostgreSQL adds PostgreSQL Service and starts postgres exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "postgres_exporter" with provided "pmm_agent_id" and other parameters.
	AddPostgreSQL(ctx context.Context, in *AddPostgreSQLRequest, opts ...grpc.CallOption) (*AddPostgreSQLResponse, error)
}

type postgreSQLClient struct {
	cc *grpc.ClientConn
}

func NewPostgreSQLClient(cc *grpc.ClientConn) PostgreSQLClient {
	return &postgreSQLClient{cc}
}

func (c *postgreSQLClient) AddPostgreSQL(ctx context.Context, in *AddPostgreSQLRequest, opts ...grpc.CallOption) (*AddPostgreSQLResponse, error) {
	out := new(AddPostgreSQLResponse)
	err := c.cc.Invoke(ctx, "/management.PostgreSQL/AddPostgreSQL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PostgreSQLServer is the server API for PostgreSQL service.
type PostgreSQLServer interface {
	// AddPostgreSQL adds PostgreSQL Service and starts postgres exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds "postgres_exporter" with provided "pmm_agent_id" and other parameters.
	AddPostgreSQL(context.Context, *AddPostgreSQLRequest) (*AddPostgreSQLResponse, error)
}

// UnimplementedPostgreSQLServer can be embedded to have forward compatible implementations.
type UnimplementedPostgreSQLServer struct {
}

func (*UnimplementedPostgreSQLServer) AddPostgreSQL(ctx context.Context, req *AddPostgreSQLRequest) (*AddPostgreSQLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPostgreSQL not implemented")
}

func RegisterPostgreSQLServer(s *grpc.Server, srv PostgreSQLServer) {
	s.RegisterService(&_PostgreSQL_serviceDesc, srv)
}

func _PostgreSQL_AddPostgreSQL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPostgreSQLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PostgreSQLServer).AddPostgreSQL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.PostgreSQL/AddPostgreSQL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PostgreSQLServer).AddPostgreSQL(ctx, req.(*AddPostgreSQLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PostgreSQL_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.PostgreSQL",
	HandlerType: (*PostgreSQLServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPostgreSQL",
			Handler:    _PostgreSQL_AddPostgreSQL_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/postgresql.proto",
}
