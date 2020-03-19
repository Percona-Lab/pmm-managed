// Code generated by protoc-gen-go. DO NOT EDIT.
// source: telemetry/reporter/reporter_api.proto

package reporterv1

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReportRequest struct {
	// One or more events to report.
	Events               []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportRequest) Reset()         { *m = ReportRequest{} }
func (m *ReportRequest) String() string { return proto.CompactTextString(m) }
func (*ReportRequest) ProtoMessage()    {}
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_127842dcbe7a0194, []int{0}
}

func (m *ReportRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportRequest.Unmarshal(m, b)
}
func (m *ReportRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportRequest.Marshal(b, m, deterministic)
}
func (m *ReportRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportRequest.Merge(m, src)
}
func (m *ReportRequest) XXX_Size() int {
	return xxx_messageInfo_ReportRequest.Size(m)
}
func (m *ReportRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReportRequest proto.InternalMessageInfo

func (m *ReportRequest) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

type ReportResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReportResponse) Reset()         { *m = ReportResponse{} }
func (m *ReportResponse) String() string { return proto.CompactTextString(m) }
func (*ReportResponse) ProtoMessage()    {}
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_127842dcbe7a0194, []int{1}
}

func (m *ReportResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReportResponse.Unmarshal(m, b)
}
func (m *ReportResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReportResponse.Marshal(b, m, deterministic)
}
func (m *ReportResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReportResponse.Merge(m, src)
}
func (m *ReportResponse) XXX_Size() int {
	return xxx_messageInfo_ReportResponse.Size(m)
}
func (m *ReportResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReportResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReportResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReportRequest)(nil), "percona.platform.telemetry.reporter.v1.ReportRequest")
	proto.RegisterType((*ReportResponse)(nil), "percona.platform.telemetry.reporter.v1.ReportResponse")
}

func init() {
	proto.RegisterFile("telemetry/reporter/reporter_api.proto", fileDescriptor_127842dcbe7a0194)
}

var fileDescriptor_127842dcbe7a0194 = []byte{
	// 240 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2d, 0x49, 0xcd, 0x49,
	0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4, 0x2f, 0x4a, 0x2d, 0xc8, 0x2f, 0x2a, 0x49, 0x2d, 0x82, 0x33,
	0xe2, 0x13, 0x0b, 0x32, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xd4, 0x0a, 0x52, 0x8b, 0x92,
	0xf3, 0xf3, 0x12, 0xf5, 0x0a, 0x72, 0x12, 0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0xf5, 0xe0, 0xfa, 0xf4,
	0x60, 0xca, 0xf5, 0xca, 0x0c, 0xa5, 0xcc, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3,
	0x73, 0xf5, 0x73, 0xcb, 0x33, 0x4b, 0xb2, 0xf3, 0xcb, 0xf5, 0xd3, 0xf3, 0x75, 0xc1, 0x86, 0xe8,
	0x96, 0x25, 0xe6, 0x64, 0xa6, 0x24, 0x96, 0xe4, 0x17, 0x15, 0xeb, 0xc3, 0x99, 0x10, 0xf3, 0xa5,
	0xe4, 0xb0, 0x38, 0x23, 0xb5, 0x2c, 0x35, 0xaf, 0x04, 0x22, 0xaf, 0x14, 0xc7, 0xc5, 0x1b, 0x04,
	0x16, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xf2, 0xe5, 0x62, 0x03, 0xcb, 0x17, 0x4b,
	0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0xe9, 0xea, 0x11, 0xe7, 0x42, 0x3d, 0x57, 0x90, 0x2e, 0x27,
	0xb6, 0x47, 0xf7, 0xe5, 0x99, 0x12, 0x18, 0x83, 0xa0, 0x86, 0x28, 0x09, 0x70, 0xf1, 0xc1, 0xcc,
	0x2f, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x35, 0x6a, 0x63, 0xe4, 0xe2, 0x0e, 0x82, 0xea, 0x73, 0x0c,
	0xf0, 0x14, 0x2a, 0xe7, 0x62, 0x83, 0x70, 0x85, 0x4c, 0x89, 0xb5, 0x0a, 0xc5, 0xc5, 0x52, 0x66,
	0xa4, 0x6a, 0x83, 0x38, 0xc4, 0x89, 0x27, 0x8a, 0x0b, 0x26, 0x5b, 0x66, 0x98, 0xc4, 0x06, 0x0e,
	0x0f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0xa1, 0x19, 0x32, 0xb8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ReporterAPIClient is the client API for ReporterAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ReporterAPIClient interface {
	// Report submits several telemetry events to the server.
	Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error)
}

type reporterAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewReporterAPIClient(cc grpc.ClientConnInterface) ReporterAPIClient {
	return &reporterAPIClient{cc}
}

func (c *reporterAPIClient) Report(ctx context.Context, in *ReportRequest, opts ...grpc.CallOption) (*ReportResponse, error) {
	out := new(ReportResponse)
	err := c.cc.Invoke(ctx, "/percona.platform.telemetry.reporter.v1.ReporterAPI/Report", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReporterAPIServer is the server API for ReporterAPI service.
type ReporterAPIServer interface {
	// Report submits several telemetry events to the server.
	Report(context.Context, *ReportRequest) (*ReportResponse, error)
}

// UnimplementedReporterAPIServer can be embedded to have forward compatible implementations.
type UnimplementedReporterAPIServer struct {
}

func (*UnimplementedReporterAPIServer) Report(ctx context.Context, req *ReportRequest) (*ReportResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Report not implemented")
}

func RegisterReporterAPIServer(s *grpc.Server, srv ReporterAPIServer) {
	s.RegisterService(&_ReporterAPI_serviceDesc, srv)
}

func _ReporterAPI_Report_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReporterAPIServer).Report(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/percona.platform.telemetry.reporter.v1.ReporterAPI/Report",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReporterAPIServer).Report(ctx, req.(*ReportRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ReporterAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "percona.platform.telemetry.reporter.v1.ReporterAPI",
	HandlerType: (*ReporterAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Report",
			Handler:    _ReporterAPI_Report_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "telemetry/reporter/reporter_api.proto",
}
