// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/metrics_names.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

// MetricsNamesRequest is emty.
type MetricsNamesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsNamesRequest) Reset()         { *m = MetricsNamesRequest{} }
func (m *MetricsNamesRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsNamesRequest) ProtoMessage()    {}
func (*MetricsNamesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed854621c154122c, []int{0}
}

func (m *MetricsNamesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsNamesRequest.Unmarshal(m, b)
}
func (m *MetricsNamesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsNamesRequest.Marshal(b, m, deterministic)
}
func (m *MetricsNamesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsNamesRequest.Merge(m, src)
}
func (m *MetricsNamesRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsNamesRequest.Size(m)
}
func (m *MetricsNamesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsNamesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsNamesRequest proto.InternalMessageInfo

// MetricsNamesReply is map of stored metrics:
// key is root of metric name in db (Ex:. [m_]query_time[_sum]);
// value - Human readable name of metrics.
type MetricsNamesReply struct {
	Data                 map[string]string `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricsNamesReply) Reset()         { *m = MetricsNamesReply{} }
func (m *MetricsNamesReply) String() string { return proto.CompactTextString(m) }
func (*MetricsNamesReply) ProtoMessage()    {}
func (*MetricsNamesReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed854621c154122c, []int{1}
}

func (m *MetricsNamesReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsNamesReply.Unmarshal(m, b)
}
func (m *MetricsNamesReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsNamesReply.Marshal(b, m, deterministic)
}
func (m *MetricsNamesReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsNamesReply.Merge(m, src)
}
func (m *MetricsNamesReply) XXX_Size() int {
	return xxx_messageInfo_MetricsNamesReply.Size(m)
}
func (m *MetricsNamesReply) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsNamesReply.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsNamesReply proto.InternalMessageInfo

func (m *MetricsNamesReply) GetData() map[string]string {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*MetricsNamesRequest)(nil), "qan.v1beta1.MetricsNamesRequest")
	proto.RegisterType((*MetricsNamesReply)(nil), "qan.v1beta1.MetricsNamesReply")
	proto.RegisterMapType((map[string]string)(nil), "qan.v1beta1.MetricsNamesReply.DataEntry")
}

func init() { proto.RegisterFile("qanpb/metrics_names.proto", fileDescriptor_ed854621c154122c) }

var fileDescriptor_ed854621c154122c = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0x02, 0x41,
	0x18, 0x86, 0x19, 0xcd, 0xc2, 0x31, 0xa8, 0xa6, 0x22, 0x5b, 0x24, 0x96, 0xad, 0x83, 0x44, 0xee,
	0xa4, 0x1d, 0x0a, 0xe9, 0x62, 0x24, 0x9e, 0x0a, 0xda, 0x63, 0x97, 0xf8, 0xd4, 0xcf, 0x65, 0x69,
	0x9d, 0xd9, 0x9d, 0x19, 0x15, 0xaf, 0x5d, 0x82, 0x8e, 0xf5, 0xd3, 0xfa, 0x0b, 0xfd, 0x90, 0x70,
	0x56, 0x42, 0x13, 0x3c, 0xed, 0xce, 0xc3, 0x33, 0xf3, 0xbe, 0x33, 0x1f, 0x3d, 0x4e, 0x41, 0x24,
	0x5d, 0x3e, 0x44, 0xa3, 0xa2, 0x9e, 0x7e, 0x11, 0x30, 0x44, 0xed, 0x27, 0x4a, 0x1a, 0xc9, 0x4a,
	0x29, 0x08, 0x7f, 0x5c, 0xef, 0xa2, 0x81, 0xba, 0x53, 0x09, 0xa5, 0x0c, 0x63, 0xe4, 0x90, 0x44,
	0x1c, 0x84, 0x90, 0x06, 0x4c, 0x24, 0xc5, 0x5c, 0x75, 0x2e, 0xec, 0xa7, 0x57, 0x0b, 0x51, 0xd4,
	0xf4, 0x04, 0xc2, 0x10, 0x15, 0x97, 0x89, 0x35, 0x56, 0x6d, 0xef, 0x90, 0xee, 0x3f, 0x64, 0x79,
	0x8f, 0xb3, 0xb8, 0x00, 0xd3, 0x11, 0x6a, 0xe3, 0x7d, 0x10, 0xba, 0xb7, 0xcc, 0x93, 0x78, 0xca,
	0x6e, 0xe9, 0x46, 0x1f, 0x0c, 0x94, 0x89, 0x9b, 0xaf, 0x96, 0x1a, 0x55, 0x7f, 0xa1, 0x94, 0xbf,
	0x62, 0xfb, 0xf7, 0x60, 0xa0, 0x2d, 0x8c, 0x9a, 0x06, 0x76, 0x97, 0x73, 0x4d, 0x8b, 0x7f, 0x88,
	0xed, 0xd2, 0xfc, 0x2b, 0x4e, 0xcb, 0xc4, 0x25, 0xd5, 0x62, 0x30, 0xfb, 0x65, 0x07, 0xb4, 0x30,
	0x86, 0x78, 0x84, 0xe5, 0x9c, 0x65, 0xd9, 0xa2, 0x99, 0xbb, 0x21, 0x8d, 0x77, 0x42, 0xb7, 0x17,
	0x8f, 0x67, 0x13, 0xba, 0xd3, 0x41, 0xb3, 0x84, 0xdc, 0x35, 0x65, 0xec, 0x95, 0x9c, 0x93, 0xf5,
	0x75, 0x3d, 0xef, 0xed, 0xfb, 0xe7, 0x2b, 0x57, 0xf1, 0x8e, 0xf8, 0xf8, 0x92, 0xa7, 0x20, 0xf8,
	0xbf, 0x88, 0x26, 0x39, 0xbf, 0x7b, 0xfa, 0x6c, 0x75, 0x82, 0x36, 0xdd, 0xea, 0xe3, 0x00, 0x46,
	0xb1, 0x61, 0x4d, 0xca, 0x5a, 0xc2, 0x45, 0xa5, 0xa4, 0x72, 0x15, 0xea, 0x44, 0x0a, 0x8d, 0x3e,
	0x3b, 0xa3, 0x9e, 0xe3, 0x9e, 0xf2, 0x3e, 0x0e, 0x22, 0x11, 0x65, 0x0f, 0x6f, 0x27, 0xdb, 0x9e,
	0x79, 0xc1, 0x5c, 0x7b, 0x2e, 0x58, 0xd6, 0xdd, 0xb4, 0x73, 0xb8, 0xfa, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xb8, 0x52, 0xfc, 0xfd, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricsNamesClient is the client API for MetricsNames service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsNamesClient interface {
	// GetMetricsNames gets map of metrics names.
	GetMetricsNames(ctx context.Context, in *MetricsNamesRequest, opts ...grpc.CallOption) (*MetricsNamesReply, error)
}

type metricsNamesClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsNamesClient(cc grpc.ClientConnInterface) MetricsNamesClient {
	return &metricsNamesClient{cc}
}

func (c *metricsNamesClient) GetMetricsNames(ctx context.Context, in *MetricsNamesRequest, opts ...grpc.CallOption) (*MetricsNamesReply, error) {
	out := new(MetricsNamesReply)
	err := c.cc.Invoke(ctx, "/qan.v1beta1.MetricsNames/GetMetricsNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsNamesServer is the server API for MetricsNames service.
type MetricsNamesServer interface {
	// GetMetricsNames gets map of metrics names.
	GetMetricsNames(context.Context, *MetricsNamesRequest) (*MetricsNamesReply, error)
}

// UnimplementedMetricsNamesServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsNamesServer struct {
}

func (*UnimplementedMetricsNamesServer) GetMetricsNames(ctx context.Context, req *MetricsNamesRequest) (*MetricsNamesReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricsNames not implemented")
}

func RegisterMetricsNamesServer(s *grpc.Server, srv MetricsNamesServer) {
	s.RegisterService(&_MetricsNames_serviceDesc, srv)
}

func _MetricsNames_GetMetricsNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsNamesServer).GetMetricsNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.v1beta1.MetricsNames/GetMetricsNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsNamesServer).GetMetricsNames(ctx, req.(*MetricsNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsNames_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.v1beta1.MetricsNames",
	HandlerType: (*MetricsNamesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMetricsNames",
			Handler:    _MetricsNames_GetMetricsNames_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/metrics_names.proto",
}
