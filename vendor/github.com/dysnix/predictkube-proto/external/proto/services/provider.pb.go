// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v4.25.1
// source: proto/services/provider.proto

package services

import (
	commonproto "github.com/dysnix/predictkube-proto/external/proto/commonproto"
	enums "github.com/dysnix/predictkube-proto/external/proto/enums"
	events "github.com/dysnix/predictkube-proto/external/proto/events"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ReqPing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping *commonproto.Ping `protobuf:"bytes,1,opt,name=ping,proto3" json:"ping,omitempty"`
}

func (x *ReqPing) Reset() {
	*x = ReqPing{}
	mi := &file_proto_services_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqPing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqPing) ProtoMessage() {}

func (x *ReqPing) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqPing.ProtoReflect.Descriptor instead.
func (*ReqPing) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{0}
}

func (x *ReqPing) GetPing() *commonproto.Ping {
	if x != nil {
		return x.Ping
	}
	return nil
}

type ResPong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong *commonproto.Pong `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *ResPong) Reset() {
	*x = ResPong{}
	mi := &file_proto_services_provider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResPong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResPong) ProtoMessage() {}

func (x *ResPong) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResPong.ProtoReflect.Descriptor instead.
func (*ResPong) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{1}
}

func (x *ResPong) GetPong() *commonproto.Pong {
	if x != nil {
		return x.Pong
	}
	return nil
}

type ReqGetMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource     enums.ResourceType   `protobuf:"varint,1,opt,name=resource,proto3,enum=enums.ResourceType" json:"resource,omitempty"`
	Name         string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Namespace    string               `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Cluster      string               `protobuf:"bytes,4,opt,name=cluster,proto3" json:"cluster,omitempty"`
	StepDuration *durationpb.Duration `protobuf:"bytes,5,opt,name=stepDuration,proto3" json:"stepDuration,omitempty"`
	TimeWindow   *durationpb.Duration `protobuf:"bytes,6,opt,name=timeWindow,proto3" json:"timeWindow,omitempty"`
	// Types that are assignable to Query:
	//
	//	*ReqGetMetrics_History
	Query isReqGetMetrics_Query `protobuf_oneof:"Query"`
}

func (x *ReqGetMetrics) Reset() {
	*x = ReqGetMetrics{}
	mi := &file_proto_services_provider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqGetMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqGetMetrics) ProtoMessage() {}

func (x *ReqGetMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqGetMetrics.ProtoReflect.Descriptor instead.
func (*ReqGetMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{2}
}

func (x *ReqGetMetrics) GetResource() enums.ResourceType {
	if x != nil {
		return x.Resource
	}
	return enums.ResourceType(0)
}

func (x *ReqGetMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqGetMetrics) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ReqGetMetrics) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ReqGetMetrics) GetStepDuration() *durationpb.Duration {
	if x != nil {
		return x.StepDuration
	}
	return nil
}

func (x *ReqGetMetrics) GetTimeWindow() *durationpb.Duration {
	if x != nil {
		return x.TimeWindow
	}
	return nil
}

func (m *ReqGetMetrics) GetQuery() isReqGetMetrics_Query {
	if m != nil {
		return m.Query
	}
	return nil
}

func (x *ReqGetMetrics) GetHistory() *events.History {
	if x, ok := x.GetQuery().(*ReqGetMetrics_History); ok {
		return x.History
	}
	return nil
}

type isReqGetMetrics_Query interface {
	isReqGetMetrics_Query()
}

type ReqGetMetrics_History struct {
	History *events.History `protobuf:"bytes,7,opt,name=history,proto3,oneof"`
}

func (*ReqGetMetrics_History) isReqGetMetrics_Query() {}

type ResGetMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricValues []*commonproto.MetricValue `protobuf:"bytes,1,rep,name=metric_values,json=metricValues,proto3" json:"metric_values,omitempty"`
}

func (x *ResGetMetrics) Reset() {
	*x = ResGetMetrics{}
	mi := &file_proto_services_provider_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResGetMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResGetMetrics) ProtoMessage() {}

func (x *ResGetMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResGetMetrics.ProtoReflect.Descriptor instead.
func (*ResGetMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{3}
}

func (x *ResGetMetrics) GetMetricValues() []*commonproto.MetricValue {
	if x != nil {
		return x.MetricValues
	}
	return nil
}

type ReqRawQueryMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace    string               `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Cluster      string               `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	StepDuration *durationpb.Duration `protobuf:"bytes,4,opt,name=stepDuration,proto3" json:"stepDuration,omitempty"`
	Raw          *events.RawQuery     `protobuf:"bytes,5,opt,name=raw,proto3" json:"raw,omitempty"`
}

func (x *ReqRawQueryMetrics) Reset() {
	*x = ReqRawQueryMetrics{}
	mi := &file_proto_services_provider_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReqRawQueryMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRawQueryMetrics) ProtoMessage() {}

func (x *ReqRawQueryMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRawQueryMetrics.ProtoReflect.Descriptor instead.
func (*ReqRawQueryMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{4}
}

func (x *ReqRawQueryMetrics) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReqRawQueryMetrics) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ReqRawQueryMetrics) GetCluster() string {
	if x != nil {
		return x.Cluster
	}
	return ""
}

func (x *ReqRawQueryMetrics) GetStepDuration() *durationpb.Duration {
	if x != nil {
		return x.StepDuration
	}
	return nil
}

func (x *ReqRawQueryMetrics) GetRaw() *events.RawQuery {
	if x != nil {
		return x.Raw
	}
	return nil
}

type ResRawQueryMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricValues []*commonproto.RawMetricValue `protobuf:"bytes,1,rep,name=metric_values,json=metricValues,proto3" json:"metric_values,omitempty"`
}

func (x *ResRawQueryMetrics) Reset() {
	*x = ResRawQueryMetrics{}
	mi := &file_proto_services_provider_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResRawQueryMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResRawQueryMetrics) ProtoMessage() {}

func (x *ResRawQueryMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_services_provider_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResRawQueryMetrics.ProtoReflect.Descriptor instead.
func (*ResRawQueryMetrics) Descriptor() ([]byte, []int) {
	return file_proto_services_provider_proto_rawDescGZIP(), []int{5}
}

func (x *ResRawQueryMetrics) GetMetricValues() []*commonproto.RawMetricValue {
	if x != nil {
		return x.MetricValues
	}
	return nil
}

var File_proto_services_provider_proto protoreflect.FileDescriptor

var file_proto_services_provider_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e,
	0x67, 0x5f, 0x70, 0x6f, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x30, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x69, 0x6e, 0x67, 0x22, 0x30, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22,
	0xbc, 0x02, 0x0a, 0x0d, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x12, 0x2f, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3d,
	0x0a, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a,
	0x0a, 0x74, 0x69, 0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x57, 0x69, 0x6e, 0x64, 0x6f, 0x77, 0x12, 0x2b, 0x0a, 0x07, 0x68, 0x69, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x48, 0x00, 0x52, 0x07, 0x68, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x42, 0x07, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x22, 0x4e,
	0x0a, 0x0d, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12,
	0x3d, 0x0a, 0x0d, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xc3,
	0x01, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x3d, 0x0a, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0c, 0x73, 0x74, 0x65, 0x70, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x22, 0x0a, 0x03, 0x72, 0x61, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52,
	0x03, 0x72, 0x61, 0x77, 0x22, 0x56, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x52, 0x61, 0x77, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x40, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x61, 0x77, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x32, 0xd7, 0x01, 0x0a,
	0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2e, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x50, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x50, 0x6f, 0x6e, 0x67, 0x22, 0x00,
	0x12, 0x40, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x17,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x17, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72,
	0x79, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x52, 0x65, 0x71, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x1a, 0x1c, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x73, 0x2e, 0x52, 0x65, 0x73, 0x52, 0x61, 0x77, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x79, 0x73, 0x6e, 0x69, 0x78, 0x2f, 0x70, 0x72, 0x65, 0x64,
	0x69, 0x63, 0x74, 0x6b, 0x75, 0x62, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_services_provider_proto_rawDescOnce sync.Once
	file_proto_services_provider_proto_rawDescData = file_proto_services_provider_proto_rawDesc
)

func file_proto_services_provider_proto_rawDescGZIP() []byte {
	file_proto_services_provider_proto_rawDescOnce.Do(func() {
		file_proto_services_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_services_provider_proto_rawDescData)
	})
	return file_proto_services_provider_proto_rawDescData
}

var file_proto_services_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_services_provider_proto_goTypes = []any{
	(*ReqPing)(nil),                    // 0: services.ReqPing
	(*ResPong)(nil),                    // 1: services.ResPong
	(*ReqGetMetrics)(nil),              // 2: services.ReqGetMetrics
	(*ResGetMetrics)(nil),              // 3: services.ResGetMetrics
	(*ReqRawQueryMetrics)(nil),         // 4: services.ReqRawQueryMetrics
	(*ResRawQueryMetrics)(nil),         // 5: services.ResRawQueryMetrics
	(*commonproto.Ping)(nil),           // 6: commonproto.Ping
	(*commonproto.Pong)(nil),           // 7: commonproto.Pong
	(enums.ResourceType)(0),            // 8: enums.ResourceType
	(*durationpb.Duration)(nil),        // 9: google.protobuf.Duration
	(*events.History)(nil),             // 10: events.History
	(*commonproto.MetricValue)(nil),    // 11: commonproto.MetricValue
	(*events.RawQuery)(nil),            // 12: events.RawQuery
	(*commonproto.RawMetricValue)(nil), // 13: commonproto.RawMetricValue
}
var file_proto_services_provider_proto_depIdxs = []int32{
	6,  // 0: services.ReqPing.ping:type_name -> commonproto.Ping
	7,  // 1: services.ResPong.pong:type_name -> commonproto.Pong
	8,  // 2: services.ReqGetMetrics.resource:type_name -> enums.ResourceType
	9,  // 3: services.ReqGetMetrics.stepDuration:type_name -> google.protobuf.Duration
	9,  // 4: services.ReqGetMetrics.timeWindow:type_name -> google.protobuf.Duration
	10, // 5: services.ReqGetMetrics.history:type_name -> events.History
	11, // 6: services.ResGetMetrics.metric_values:type_name -> commonproto.MetricValue
	9,  // 7: services.ReqRawQueryMetrics.stepDuration:type_name -> google.protobuf.Duration
	12, // 8: services.ReqRawQueryMetrics.raw:type_name -> events.RawQuery
	13, // 9: services.ResRawQueryMetrics.metric_values:type_name -> commonproto.RawMetricValue
	0,  // 10: services.ProviderService.Ping:input_type -> services.ReqPing
	2,  // 11: services.ProviderService.GetMetrics:input_type -> services.ReqGetMetrics
	4,  // 12: services.ProviderService.GetRawQueryMetrics:input_type -> services.ReqRawQueryMetrics
	1,  // 13: services.ProviderService.Ping:output_type -> services.ResPong
	3,  // 14: services.ProviderService.GetMetrics:output_type -> services.ResGetMetrics
	5,  // 15: services.ProviderService.GetRawQueryMetrics:output_type -> services.ResRawQueryMetrics
	13, // [13:16] is the sub-list for method output_type
	10, // [10:13] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_proto_services_provider_proto_init() }
func file_proto_services_provider_proto_init() {
	if File_proto_services_provider_proto != nil {
		return
	}
	file_proto_services_provider_proto_msgTypes[2].OneofWrappers = []any{
		(*ReqGetMetrics_History)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_services_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_services_provider_proto_goTypes,
		DependencyIndexes: file_proto_services_provider_proto_depIdxs,
		MessageInfos:      file_proto_services_provider_proto_msgTypes,
	}.Build()
	File_proto_services_provider_proto = out.File
	file_proto_services_provider_proto_rawDesc = nil
	file_proto_services_provider_proto_goTypes = nil
	file_proto_services_provider_proto_depIdxs = nil
}
