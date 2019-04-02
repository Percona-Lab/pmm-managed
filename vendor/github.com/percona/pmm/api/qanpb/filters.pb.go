// Code generated by protoc-gen-go. DO NOT EDIT.
// source: qanpb/filters.proto

package qanpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// FiltersRequest contains period for which we need filters.
type FiltersRequest struct {
	PeriodStartFrom      *timestamp.Timestamp `protobuf:"bytes,1,opt,name=period_start_from,json=periodStartFrom,proto3" json:"period_start_from,omitempty"`
	PeriodStartTo        *timestamp.Timestamp `protobuf:"bytes,2,opt,name=period_start_to,json=periodStartTo,proto3" json:"period_start_to,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *FiltersRequest) Reset()         { *m = FiltersRequest{} }
func (m *FiltersRequest) String() string { return proto.CompactTextString(m) }
func (*FiltersRequest) ProtoMessage()    {}
func (*FiltersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{0}
}

func (m *FiltersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiltersRequest.Unmarshal(m, b)
}
func (m *FiltersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiltersRequest.Marshal(b, m, deterministic)
}
func (m *FiltersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiltersRequest.Merge(m, src)
}
func (m *FiltersRequest) XXX_Size() int {
	return xxx_messageInfo_FiltersRequest.Size(m)
}
func (m *FiltersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FiltersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FiltersRequest proto.InternalMessageInfo

func (m *FiltersRequest) GetPeriodStartFrom() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartFrom
	}
	return nil
}

func (m *FiltersRequest) GetPeriodStartTo() *timestamp.Timestamp {
	if m != nil {
		return m.PeriodStartTo
	}
	return nil
}

// FiltersReply is map of labels for given period by key.
// Key is label's name and value is label's value and how many times it occur.
type FiltersReply struct {
	Labels               map[string]*ListLabels `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *FiltersReply) Reset()         { *m = FiltersReply{} }
func (m *FiltersReply) String() string { return proto.CompactTextString(m) }
func (*FiltersReply) ProtoMessage()    {}
func (*FiltersReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{1}
}

func (m *FiltersReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FiltersReply.Unmarshal(m, b)
}
func (m *FiltersReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FiltersReply.Marshal(b, m, deterministic)
}
func (m *FiltersReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FiltersReply.Merge(m, src)
}
func (m *FiltersReply) XXX_Size() int {
	return xxx_messageInfo_FiltersReply.Size(m)
}
func (m *FiltersReply) XXX_DiscardUnknown() {
	xxx_messageInfo_FiltersReply.DiscardUnknown(m)
}

var xxx_messageInfo_FiltersReply proto.InternalMessageInfo

func (m *FiltersReply) GetLabels() map[string]*ListLabels {
	if m != nil {
		return m.Labels
	}
	return nil
}

// ListLabels is list of label's values: duplicates are impossible.
type ListLabels struct {
	Name                 []*ValueAndCount `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ListLabels) Reset()         { *m = ListLabels{} }
func (m *ListLabels) String() string { return proto.CompactTextString(m) }
func (*ListLabels) ProtoMessage()    {}
func (*ListLabels) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{2}
}

func (m *ListLabels) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListLabels.Unmarshal(m, b)
}
func (m *ListLabels) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListLabels.Marshal(b, m, deterministic)
}
func (m *ListLabels) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListLabels.Merge(m, src)
}
func (m *ListLabels) XXX_Size() int {
	return xxx_messageInfo_ListLabels.Size(m)
}
func (m *ListLabels) XXX_DiscardUnknown() {
	xxx_messageInfo_ListLabels.DiscardUnknown(m)
}

var xxx_messageInfo_ListLabels proto.InternalMessageInfo

func (m *ListLabels) GetName() []*ValueAndCount {
	if m != nil {
		return m.Name
	}
	return nil
}

// ValueAndCount is label values and how many times this value occur.
type ValueAndCount struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueAndCount) Reset()         { *m = ValueAndCount{} }
func (m *ValueAndCount) String() string { return proto.CompactTextString(m) }
func (*ValueAndCount) ProtoMessage()    {}
func (*ValueAndCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_2662fef59f58e02c, []int{3}
}

func (m *ValueAndCount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueAndCount.Unmarshal(m, b)
}
func (m *ValueAndCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueAndCount.Marshal(b, m, deterministic)
}
func (m *ValueAndCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueAndCount.Merge(m, src)
}
func (m *ValueAndCount) XXX_Size() int {
	return xxx_messageInfo_ValueAndCount.Size(m)
}
func (m *ValueAndCount) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueAndCount.DiscardUnknown(m)
}

var xxx_messageInfo_ValueAndCount proto.InternalMessageInfo

func (m *ValueAndCount) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ValueAndCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*FiltersRequest)(nil), "qan.FiltersRequest")
	proto.RegisterType((*FiltersReply)(nil), "qan.FiltersReply")
	proto.RegisterMapType((map[string]*ListLabels)(nil), "qan.FiltersReply.LabelsEntry")
	proto.RegisterType((*ListLabels)(nil), "qan.ListLabels")
	proto.RegisterType((*ValueAndCount)(nil), "qan.ValueAndCount")
}

func init() { proto.RegisterFile("qanpb/filters.proto", fileDescriptor_2662fef59f58e02c) }

var fileDescriptor_2662fef59f58e02c = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4a, 0xe3, 0x50,
	0x14, 0xc6, 0x49, 0x33, 0x6d, 0x99, 0xd3, 0xe9, 0x74, 0x7a, 0x3b, 0x8b, 0x10, 0xe6, 0x4f, 0x09,
	0xcc, 0x50, 0x66, 0x91, 0xcb, 0x54, 0x05, 0xa9, 0x2b, 0x2b, 0xb6, 0x20, 0x75, 0x13, 0x8b, 0x88,
	0x9b, 0x72, 0x63, 0x6f, 0x4b, 0x30, 0xb9, 0x37, 0x7f, 0x4e, 0x0a, 0xdd, 0xfa, 0x06, 0xe2, 0xd6,
	0xb7, 0xf2, 0x15, 0x7c, 0x10, 0xc9, 0xbd, 0xa9, 0xb6, 0xb8, 0x70, 0x97, 0x73, 0xbe, 0xef, 0xfc,
	0xf2, 0x9d, 0x9c, 0x40, 0x27, 0x61, 0x22, 0xf6, 0xe9, 0x22, 0x08, 0x91, 0xa7, 0x99, 0x1b, 0xa7,
	0x12, 0x25, 0x31, 0x13, 0x26, 0xec, 0x1f, 0x4b, 0x29, 0x97, 0x21, 0xa7, 0x2c, 0x0e, 0x28, 0x13,
	0x42, 0x22, 0xc3, 0x40, 0x8a, 0xd2, 0x62, 0xff, 0x2e, 0x55, 0x55, 0xf9, 0xf9, 0x82, 0x62, 0x10,
	0xf1, 0x0c, 0x59, 0x14, 0x6b, 0x83, 0xf3, 0x68, 0xc0, 0xd7, 0x91, 0xa6, 0x7a, 0x3c, 0xc9, 0x79,
	0x86, 0x64, 0x04, 0xed, 0x98, 0xa7, 0x81, 0x9c, 0xcf, 0x32, 0x64, 0x29, 0xce, 0x16, 0xa9, 0x8c,
	0x2c, 0xa3, 0x6b, 0xf4, 0x1a, 0x7d, 0xdb, 0xd5, 0x3c, 0x77, 0xc3, 0x73, 0xa7, 0x1b, 0x9e, 0xd7,
	0xd2, 0x43, 0x17, 0xc5, 0xcc, 0x28, 0x95, 0x11, 0x19, 0x42, 0x6b, 0x87, 0x83, 0xd2, 0xaa, 0x7c,
	0x48, 0x69, 0x6e, 0x51, 0xa6, 0xd2, 0xb9, 0x37, 0xe0, 0xcb, 0x6b, 0xbc, 0x38, 0x5c, 0x93, 0x03,
	0xa8, 0x85, 0xcc, 0xe7, 0x61, 0x66, 0x19, 0x5d, 0xb3, 0xd7, 0xe8, 0xff, 0x74, 0x13, 0x26, 0xdc,
	0x6d, 0x8b, 0x3b, 0x51, 0xfa, 0xa9, 0xc0, 0x74, 0xed, 0x95, 0x66, 0xfb, 0x0c, 0x1a, 0x5b, 0x6d,
	0xf2, 0x0d, 0xcc, 0x5b, 0xbe, 0x56, 0x4b, 0x7d, 0xf6, 0x8a, 0x47, 0xf2, 0x07, 0xaa, 0x2b, 0x16,
	0xe6, 0xbc, 0x8c, 0xd8, 0x52, 0xd8, 0x49, 0x90, 0xa1, 0x1e, 0xf3, 0xb4, 0x3a, 0xa8, 0x1c, 0x1a,
	0xce, 0x3e, 0xc0, 0x9b, 0x40, 0xfe, 0xc2, 0x27, 0xc1, 0x22, 0x5e, 0xc6, 0x21, 0x6a, 0xee, 0xb2,
	0xf0, 0x1e, 0x8b, 0xf9, 0x89, 0xcc, 0x05, 0x7a, 0x4a, 0x77, 0x8e, 0xa0, 0xb9, 0xd3, 0x26, 0xdf,
	0x37, 0x6f, 0xd4, 0x29, 0x74, 0x51, 0x74, 0x6f, 0x0a, 0x59, 0xe5, 0x30, 0x3d, 0x5d, 0xf4, 0xaf,
	0xa0, 0x5e, 0xae, 0x48, 0xce, 0xc1, 0x1c, 0x73, 0x24, 0x9d, 0xdd, 0xbd, 0xd5, 0xe5, 0xec, 0xf6,
	0xbb, 0x8f, 0xe1, 0xfc, 0xba, 0x7b, 0x7a, 0x7e, 0xa8, 0x58, 0x4e, 0x87, 0xae, 0xfe, 0xd3, 0x84,
	0x09, 0x5a, 0xaa, 0x74, 0xcc, 0x71, 0x60, 0xfc, 0x1b, 0xd6, 0xaf, 0xab, 0xea, 0xd7, 0xf2, 0x6b,
	0xea, 0x18, 0x7b, 0x2f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x20, 0x6d, 0x91, 0x6a, 0x02, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// FiltersClient is the client API for Filters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FiltersClient interface {
	// Get gets map of metrics names.
	Get(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (*FiltersReply, error)
}

type filtersClient struct {
	cc *grpc.ClientConn
}

func NewFiltersClient(cc *grpc.ClientConn) FiltersClient {
	return &filtersClient{cc}
}

func (c *filtersClient) Get(ctx context.Context, in *FiltersRequest, opts ...grpc.CallOption) (*FiltersReply, error) {
	out := new(FiltersReply)
	err := c.cc.Invoke(ctx, "/qan.Filters/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FiltersServer is the server API for Filters service.
type FiltersServer interface {
	// Get gets map of metrics names.
	Get(context.Context, *FiltersRequest) (*FiltersReply, error)
}

func RegisterFiltersServer(s *grpc.Server, srv FiltersServer) {
	s.RegisterService(&_Filters_serviceDesc, srv)
}

func _Filters_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FiltersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FiltersServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/qan.Filters/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FiltersServer).Get(ctx, req.(*FiltersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Filters_serviceDesc = grpc.ServiceDesc{
	ServiceName: "qan.Filters",
	HandlerType: (*FiltersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Filters_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "qanpb/filters.proto",
}
