// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: managementpb/haproxy.proto

package managementpb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	inventorypb "github.com/percona/pmm/api/inventorypb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
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

type AddHAProxyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Node identifier on which an external exporter is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeId string `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Node name on which a service and node is been running.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	NodeName string `protobuf:"bytes,18,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	// Create a new Node with those parameters.
	// add_node always should be passed with address.
	// Exactly one of these parameters should be present: node_id, node_name, add_node.
	AddNode *AddNodeParams `protobuf:"bytes,19,opt,name=add_node,json=addNode,proto3" json:"add_node,omitempty"`
	// Node and Exporter access address (DNS name or IP).
	// address always should be passed with add_node.
	Address string `protobuf:"bytes,20,opt,name=address,proto3" json:"address,omitempty"`
	// Unique across all Services user-defined name. Required.
	ServiceName string `protobuf:"bytes,2,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// HTTP basic auth username for collecting metrics.
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	// HTTP basic auth password for collecting metrics.
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	// Scheme to generate URI to exporter metrics endpoints.
	Scheme string `protobuf:"bytes,5,opt,name=scheme,proto3" json:"scheme,omitempty"`
	// Path under which metrics are exposed, used to generate URI.
	MetricsPath string `protobuf:"bytes,6,opt,name=metrics_path,json=metricsPath,proto3" json:"metrics_path,omitempty"`
	// Listen port for scraping metrics.
	ListenPort uint32 `protobuf:"varint,7,opt,name=listen_port,json=listenPort,proto3" json:"listen_port,omitempty"`
	// Environment name.
	Environment string `protobuf:"bytes,9,opt,name=environment,proto3" json:"environment,omitempty"`
	// Cluster name.
	Cluster string `protobuf:"bytes,10,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Replication set name.
	ReplicationSet string `protobuf:"bytes,11,opt,name=replication_set,json=replicationSet,proto3" json:"replication_set,omitempty"`
	// Custom user-assigned labels for Service.
	CustomLabels map[string]string `protobuf:"bytes,15,rep,name=custom_labels,json=customLabels,proto3" json:"custom_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Defines metrics flow model for this exporter.
	// Metrics could be pushed to the server with vmagent,
	// pulled by the server, or the server could choose behavior automatically.
	// Node with registered pmm_agent_id must present at pmm-server
	// in case of push metrics_mode.
	MetricsMode MetricsMode `protobuf:"varint,16,opt,name=metrics_mode,json=metricsMode,proto3,enum=management.MetricsMode" json:"metrics_mode,omitempty"`
}

func (x *AddHAProxyRequest) Reset() {
	*x = AddHAProxyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_haproxy_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHAProxyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHAProxyRequest) ProtoMessage() {}

func (x *AddHAProxyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_haproxy_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHAProxyRequest.ProtoReflect.Descriptor instead.
func (*AddHAProxyRequest) Descriptor() ([]byte, []int) {
	return file_managementpb_haproxy_proto_rawDescGZIP(), []int{0}
}

func (x *AddHAProxyRequest) GetNodeId() string {
	if x != nil {
		return x.NodeId
	}
	return ""
}

func (x *AddHAProxyRequest) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *AddHAProxyRequest) GetAddNode() *AddNodeParams {
	if x != nil {
		return x.AddNode
	}
	return nil
}

func (x *AddHAProxyRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *AddHAProxyRequest) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AddHAProxyRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AddHAProxyRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AddHAProxyRequest) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *AddHAProxyRequest) GetMetricsPath() string {
	if x != nil {
		return x.MetricsPath
	}
	return ""
}

func (x *AddHAProxyRequest) GetListenPort() uint32 {
	if x != nil {
		return x.ListenPort
	}
	return 0
}

func (x *AddHAProxyRequest) GetEnvironment() string {
	if x != nil {
		return x.Environment
	}
	return ""
}

func (x *AddHAProxyRequest) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *AddHAProxyRequest) GetReplicationSet() string {
	if x != nil {
		return x.ReplicationSet
	}
	return ""
}

func (x *AddHAProxyRequest) GetCustomLabels() map[string]string {
	if x != nil {
		return x.CustomLabels
	}
	return nil
}

func (x *AddHAProxyRequest) GetMetricsMode() MetricsMode {
	if x != nil {
		return x.MetricsMode
	}
	return MetricsMode_AUTO
}

type AddHAProxyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service          *inventorypb.HAProxyService   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	ExternalExporter *inventorypb.ExternalExporter `protobuf:"bytes,2,opt,name=external_exporter,json=externalExporter,proto3" json:"external_exporter,omitempty"`
}

func (x *AddHAProxyResponse) Reset() {
	*x = AddHAProxyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_managementpb_haproxy_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddHAProxyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddHAProxyResponse) ProtoMessage() {}

func (x *AddHAProxyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_managementpb_haproxy_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddHAProxyResponse.ProtoReflect.Descriptor instead.
func (*AddHAProxyResponse) Descriptor() ([]byte, []int) {
	return file_managementpb_haproxy_proto_rawDescGZIP(), []int{1}
}

func (x *AddHAProxyResponse) GetService() *inventorypb.HAProxyService {
	if x != nil {
		return x.Service
	}
	return nil
}

func (x *AddHAProxyResponse) GetExternalExporter() *inventorypb.ExternalExporter {
	if x != nil {
		return x.ExternalExporter
	}
	return nil
}

var File_managementpb_haproxy_proto protoreflect.FileDescriptor

var file_managementpb_haproxy_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x68,
	0x61, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x77, 0x69, 0x74, 0x6b, 0x6f, 0x77, 0x2f, 0x67, 0x6f, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74,
	0x6f, 0x72, 0x79, 0x70, 0x62, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x70, 0x62, 0x2f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1a, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70, 0x62, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x05, 0x0a,
	0x11, 0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x34, 0x0a, 0x08, 0x61, 0x64, 0x64, 0x5f,
	0x6e, 0x6f, 0x64, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x29, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06,
	0xe2, 0xdf, 0x1f, 0x02, 0x58, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x50, 0x61, 0x74, 0x68, 0x12, 0x2b, 0x0a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x5f, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0a, 0xe2, 0xdf, 0x1f,
	0x06, 0x10, 0x00, 0x18, 0x80, 0x80, 0x04, 0x52, 0x0a, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x50,
	0x6f, 0x72, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f, 0x6e, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x6e, 0x76, 0x69, 0x72, 0x6f,
	0x6e, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12,
	0x27, 0x0a, 0x0f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x74, 0x12, 0x54, 0x0a, 0x0d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2f, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64,
	0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x3a,
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x4d, 0x6f, 0x64, 0x65, 0x1a, 0x3f, 0x0a, 0x11, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x93, 0x01, 0x0a, 0x12,
	0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x33, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e,
	0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x5f, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x45,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x52,
	0x10, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x72, 0x32, 0x7d, 0x0a, 0x07, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x72, 0x0a, 0x0a,
	0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x12, 0x1d, 0x2e, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f,
	0x78, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1f, 0x22, 0x1a, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x48, 0x41, 0x50, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x41, 0x64, 0x64, 0x3a, 0x01, 0x2a,
	0x42, 0x1f, 0x5a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_managementpb_haproxy_proto_rawDescOnce sync.Once
	file_managementpb_haproxy_proto_rawDescData = file_managementpb_haproxy_proto_rawDesc
)

func file_managementpb_haproxy_proto_rawDescGZIP() []byte {
	file_managementpb_haproxy_proto_rawDescOnce.Do(func() {
		file_managementpb_haproxy_proto_rawDescData = protoimpl.X.CompressGZIP(file_managementpb_haproxy_proto_rawDescData)
	})
	return file_managementpb_haproxy_proto_rawDescData
}

var file_managementpb_haproxy_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_managementpb_haproxy_proto_goTypes = []interface{}{
	(*AddHAProxyRequest)(nil),            // 0: management.AddHAProxyRequest
	(*AddHAProxyResponse)(nil),           // 1: management.AddHAProxyResponse
	nil,                                  // 2: management.AddHAProxyRequest.CustomLabelsEntry
	(*AddNodeParams)(nil),                // 3: management.AddNodeParams
	(MetricsMode)(0),                     // 4: management.MetricsMode
	(*inventorypb.HAProxyService)(nil),   // 5: inventory.HAProxyService
	(*inventorypb.ExternalExporter)(nil), // 6: inventory.ExternalExporter
}
var file_managementpb_haproxy_proto_depIdxs = []int32{
	3, // 0: management.AddHAProxyRequest.add_node:type_name -> management.AddNodeParams
	2, // 1: management.AddHAProxyRequest.custom_labels:type_name -> management.AddHAProxyRequest.CustomLabelsEntry
	4, // 2: management.AddHAProxyRequest.metrics_mode:type_name -> management.MetricsMode
	5, // 3: management.AddHAProxyResponse.service:type_name -> inventory.HAProxyService
	6, // 4: management.AddHAProxyResponse.external_exporter:type_name -> inventory.ExternalExporter
	0, // 5: management.HAProxy.AddHAProxy:input_type -> management.AddHAProxyRequest
	1, // 6: management.HAProxy.AddHAProxy:output_type -> management.AddHAProxyResponse
	6, // [6:7] is the sub-list for method output_type
	5, // [5:6] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_managementpb_haproxy_proto_init() }
func file_managementpb_haproxy_proto_init() {
	if File_managementpb_haproxy_proto != nil {
		return
	}
	file_managementpb_metrics_proto_init()
	file_managementpb_service_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_managementpb_haproxy_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHAProxyRequest); i {
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
		file_managementpb_haproxy_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddHAProxyResponse); i {
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
			RawDescriptor: file_managementpb_haproxy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_managementpb_haproxy_proto_goTypes,
		DependencyIndexes: file_managementpb_haproxy_proto_depIdxs,
		MessageInfos:      file_managementpb_haproxy_proto_msgTypes,
	}.Build()
	File_managementpb_haproxy_proto = out.File
	file_managementpb_haproxy_proto_rawDesc = nil
	file_managementpb_haproxy_proto_goTypes = nil
	file_managementpb_haproxy_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// HAProxyClient is the client API for HAProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HAProxyClient interface {
	// AddHAProxy adds HAProxy service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddHAProxy(ctx context.Context, in *AddHAProxyRequest, opts ...grpc.CallOption) (*AddHAProxyResponse, error)
}

type hAProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewHAProxyClient(cc grpc.ClientConnInterface) HAProxyClient {
	return &hAProxyClient{cc}
}

func (c *hAProxyClient) AddHAProxy(ctx context.Context, in *AddHAProxyRequest, opts ...grpc.CallOption) (*AddHAProxyResponse, error) {
	out := new(AddHAProxyResponse)
	err := c.cc.Invoke(ctx, "/management.HAProxy/AddHAProxy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HAProxyServer is the server API for HAProxy service.
type HAProxyServer interface {
	// AddHAProxy adds HAProxy service and adds external exporter.
	// It automatically adds a service to inventory, which is running on provided "node_id",
	// then adds an "external exporter" agent to inventory, which is running on provided "runs_on_node_id".
	AddHAProxy(context.Context, *AddHAProxyRequest) (*AddHAProxyResponse, error)
}

// UnimplementedHAProxyServer can be embedded to have forward compatible implementations.
type UnimplementedHAProxyServer struct {
}

func (*UnimplementedHAProxyServer) AddHAProxy(context.Context, *AddHAProxyRequest) (*AddHAProxyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHAProxy not implemented")
}

func RegisterHAProxyServer(s *grpc.Server, srv HAProxyServer) {
	s.RegisterService(&_HAProxy_serviceDesc, srv)
}

func _HAProxy_AddHAProxy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHAProxyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HAProxyServer).AddHAProxy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/management.HAProxy/AddHAProxy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HAProxyServer).AddHAProxy(ctx, req.(*AddHAProxyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _HAProxy_serviceDesc = grpc.ServiceDesc{
	ServiceName: "management.HAProxy",
	HandlerType: (*HAProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddHAProxy",
			Handler:    _HAProxy_AddHAProxy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "managementpb/haproxy.proto",
}
