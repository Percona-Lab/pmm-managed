// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.13.0
// source: telemetry/reporter/reporter_api.proto

package reporterv1

import (
	context "context"
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type ReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// One or more events to report.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ReportRequest) Reset() {
	*x = ReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportRequest) ProtoMessage() {}

func (x *ReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportRequest.ProtoReflect.Descriptor instead.
func (*ReportRequest) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{0}
}

func (x *ReportRequest) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type ReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReportResponse) Reset() {
	*x = ReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportResponse) ProtoMessage() {}

func (x *ReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_telemetry_reporter_reporter_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportResponse.ProtoReflect.Descriptor instead.
func (*ReportResponse) Descriptor() ([]byte, []int) {
	return file_telemetry_reporter_reporter_api_proto_rawDescGZIP(), []int{1}
}

var File_telemetry_reporter_reporter_api_proto protoreflect.FileDescriptor

var file_telemetry_reporter_reporter_api_proto_rawDesc = []byte{
	0x0a, 0x25, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x26, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61,
	0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65,
	0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74,
	0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5e, 0x0a, 0x0d, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4d, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f,
	0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x42, 0x06, 0xe2, 0xdf, 0x1f, 0x02, 0x60, 0x01, 0x52,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x86, 0x01, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x41, 0x50, 0x49, 0x12, 0x77, 0x0a, 0x06, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x35, 0x2e, 0x70, 0x65, 0x72, 0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c,
	0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x36, 0x2e, 0x70, 0x65, 0x72,
	0x63, 0x6f, 0x6e, 0x61, 0x2e, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x1f, 0x5a, 0x1d, 0x74, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x3b, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_telemetry_reporter_reporter_api_proto_rawDescOnce sync.Once
	file_telemetry_reporter_reporter_api_proto_rawDescData = file_telemetry_reporter_reporter_api_proto_rawDesc
)

func file_telemetry_reporter_reporter_api_proto_rawDescGZIP() []byte {
	file_telemetry_reporter_reporter_api_proto_rawDescOnce.Do(func() {
		file_telemetry_reporter_reporter_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_telemetry_reporter_reporter_api_proto_rawDescData)
	})
	return file_telemetry_reporter_reporter_api_proto_rawDescData
}

var (
	file_telemetry_reporter_reporter_api_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_telemetry_reporter_reporter_api_proto_goTypes  = []interface{}{
		(*ReportRequest)(nil),  // 0: percona.platform.telemetry.reporter.v1.ReportRequest
		(*ReportResponse)(nil), // 1: percona.platform.telemetry.reporter.v1.ReportResponse
		(*Event)(nil),          // 2: percona.platform.telemetry.reporter.v1.Event
	}
)

var file_telemetry_reporter_reporter_api_proto_depIdxs = []int32{
	2, // 0: percona.platform.telemetry.reporter.v1.ReportRequest.events:type_name -> percona.platform.telemetry.reporter.v1.Event
	0, // 1: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:input_type -> percona.platform.telemetry.reporter.v1.ReportRequest
	1, // 2: percona.platform.telemetry.reporter.v1.ReporterAPI.Report:output_type -> percona.platform.telemetry.reporter.v1.ReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_telemetry_reporter_reporter_api_proto_init() }
func file_telemetry_reporter_reporter_api_proto_init() {
	if File_telemetry_reporter_reporter_api_proto != nil {
		return
	}
	file_telemetry_reporter_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_telemetry_reporter_reporter_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_telemetry_reporter_reporter_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReportResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_telemetry_reporter_reporter_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_telemetry_reporter_reporter_api_proto_goTypes,
		DependencyIndexes: file_telemetry_reporter_reporter_api_proto_depIdxs,
		MessageInfos:      file_telemetry_reporter_reporter_api_proto_msgTypes,
	}.Build()
	File_telemetry_reporter_reporter_api_proto = out.File
	file_telemetry_reporter_reporter_api_proto_rawDesc = nil
	file_telemetry_reporter_reporter_api_proto_goTypes = nil
	file_telemetry_reporter_reporter_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConnInterface
)

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

func (*UnimplementedReporterAPIServer) Report(context.Context, *ReportRequest) (*ReportResponse, error) {
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
