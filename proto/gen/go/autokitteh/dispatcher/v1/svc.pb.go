// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/dispatcher/v1/svc.proto

package dispatcherv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v1 "go.autokitteh.dev/autokitteh/proto/gen/go/autokitteh/events/v1"
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

type DispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event        *v1.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	DeploymentId string    `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
	Env          string    `protobuf:"bytes,3,opt,name=env,proto3" json:"env,omitempty"` // either env_id or full env path.
}

func (x *DispatchRequest) Reset() {
	*x = DispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchRequest) ProtoMessage() {}

func (x *DispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchRequest.ProtoReflect.Descriptor instead.
func (*DispatchRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_dispatcher_v1_svc_proto_rawDescGZIP(), []int{0}
}

func (x *DispatchRequest) GetEvent() *v1.Event {
	if x != nil {
		return x.Event
	}
	return nil
}

func (x *DispatchRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

func (x *DispatchRequest) GetEnv() string {
	if x != nil {
		return x.Env
	}
	return ""
}

type DispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *DispatchResponse) Reset() {
	*x = DispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DispatchResponse) ProtoMessage() {}

func (x *DispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DispatchResponse.ProtoReflect.Descriptor instead.
func (*DispatchResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_dispatcher_v1_svc_proto_rawDescGZIP(), []int{1}
}

func (x *DispatchResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type RedispatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId      string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EnvId        string `protobuf:"bytes,3,opt,name=env_id,json=envId,proto3" json:"env_id,omitempty"`
	DeploymentId string `protobuf:"bytes,2,opt,name=deployment_id,json=deploymentId,proto3" json:"deployment_id,omitempty"`
}

func (x *RedispatchRequest) Reset() {
	*x = RedispatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedispatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedispatchRequest) ProtoMessage() {}

func (x *RedispatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedispatchRequest.ProtoReflect.Descriptor instead.
func (*RedispatchRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_dispatcher_v1_svc_proto_rawDescGZIP(), []int{2}
}

func (x *RedispatchRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *RedispatchRequest) GetEnvId() string {
	if x != nil {
		return x.EnvId
	}
	return ""
}

func (x *RedispatchRequest) GetDeploymentId() string {
	if x != nil {
		return x.DeploymentId
	}
	return ""
}

type RedispatchResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *RedispatchResponse) Reset() {
	*x = RedispatchResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedispatchResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedispatchResponse) ProtoMessage() {}

func (x *RedispatchResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_dispatcher_v1_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedispatchResponse.ProtoReflect.Descriptor instead.
func (*RedispatchResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_dispatcher_v1_svc_proto_rawDescGZIP(), []int{3}
}

func (x *RedispatchResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

var File_autokitteh_dispatcher_v1_svc_proto protoreflect.FileDescriptor

var file_autokitteh_dispatcher_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x22, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x20,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x01,
	0x0a, 0x0f, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70,
	0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x76,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x3a, 0x77, 0xfa, 0xf7, 0x18,
	0x73, 0x1a, 0x71, 0x0a, 0x21, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x62, 0x65,
	0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64,
	0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62, 0x65, 0x20, 0x73, 0x70, 0x65,
	0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x2c, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x3d,
	0x3d, 0x20, 0x27, 0x27, 0x22, 0x37, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x74, 0x0a,
	0x11, 0x52, 0x65, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x15, 0x0a, 0x06, 0x65, 0x6e, 0x76, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6e, 0x76, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x12, 0x52, 0x65, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x32, 0xdf,
	0x01, 0x0a, 0x11, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x08, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x29, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64, 0x69,
	0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x67, 0x0a, 0x0a, 0x52, 0x65, 0x64, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0xfb, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x2e, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x42, 0x08, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4f, 0x67,
	0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64, 0x65, 0x76,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x2f, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x3b, 0x64, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x41, 0x44, 0x58, 0xaa, 0x02, 0x18, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x18, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x44, 0x69, 0x73,
	0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x24, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x1a, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x3a, 0x3a,
	0x44, 0x69, 0x73, 0x70, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_dispatcher_v1_svc_proto_rawDescOnce sync.Once
	file_autokitteh_dispatcher_v1_svc_proto_rawDescData = file_autokitteh_dispatcher_v1_svc_proto_rawDesc
)

func file_autokitteh_dispatcher_v1_svc_proto_rawDescGZIP() []byte {
	file_autokitteh_dispatcher_v1_svc_proto_rawDescOnce.Do(func() {
		file_autokitteh_dispatcher_v1_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_dispatcher_v1_svc_proto_rawDescData)
	})
	return file_autokitteh_dispatcher_v1_svc_proto_rawDescData
}

var file_autokitteh_dispatcher_v1_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_autokitteh_dispatcher_v1_svc_proto_goTypes = []interface{}{
	(*DispatchRequest)(nil),    // 0: autokitteh.dispatcher.v1.DispatchRequest
	(*DispatchResponse)(nil),   // 1: autokitteh.dispatcher.v1.DispatchResponse
	(*RedispatchRequest)(nil),  // 2: autokitteh.dispatcher.v1.RedispatchRequest
	(*RedispatchResponse)(nil), // 3: autokitteh.dispatcher.v1.RedispatchResponse
	(*v1.Event)(nil),           // 4: autokitteh.events.v1.Event
}
var file_autokitteh_dispatcher_v1_svc_proto_depIdxs = []int32{
	4, // 0: autokitteh.dispatcher.v1.DispatchRequest.event:type_name -> autokitteh.events.v1.Event
	0, // 1: autokitteh.dispatcher.v1.DispatcherService.Dispatch:input_type -> autokitteh.dispatcher.v1.DispatchRequest
	2, // 2: autokitteh.dispatcher.v1.DispatcherService.Redispatch:input_type -> autokitteh.dispatcher.v1.RedispatchRequest
	1, // 3: autokitteh.dispatcher.v1.DispatcherService.Dispatch:output_type -> autokitteh.dispatcher.v1.DispatchResponse
	3, // 4: autokitteh.dispatcher.v1.DispatcherService.Redispatch:output_type -> autokitteh.dispatcher.v1.RedispatchResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_autokitteh_dispatcher_v1_svc_proto_init() }
func file_autokitteh_dispatcher_v1_svc_proto_init() {
	if File_autokitteh_dispatcher_v1_svc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_dispatcher_v1_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchRequest); i {
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
		file_autokitteh_dispatcher_v1_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DispatchResponse); i {
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
		file_autokitteh_dispatcher_v1_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedispatchRequest); i {
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
		file_autokitteh_dispatcher_v1_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedispatchResponse); i {
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
			RawDescriptor: file_autokitteh_dispatcher_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autokitteh_dispatcher_v1_svc_proto_goTypes,
		DependencyIndexes: file_autokitteh_dispatcher_v1_svc_proto_depIdxs,
		MessageInfos:      file_autokitteh_dispatcher_v1_svc_proto_msgTypes,
	}.Build()
	File_autokitteh_dispatcher_v1_svc_proto = out.File
	file_autokitteh_dispatcher_v1_svc_proto_rawDesc = nil
	file_autokitteh_dispatcher_v1_svc_proto_goTypes = nil
	file_autokitteh_dispatcher_v1_svc_proto_depIdxs = nil
}
