// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/events/v1/svc.proto

package eventsv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type SaveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *SaveRequest) Reset() {
	*x = SaveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveRequest) ProtoMessage() {}

func (x *SaveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveRequest.ProtoReflect.Descriptor instead.
func (*SaveRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{0}
}

func (x *SaveRequest) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type SaveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *SaveResponse) Reset() {
	*x = SaveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveResponse) ProtoMessage() {}

func (x *SaveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveResponse.ProtoReflect.Descriptor instead.
func (*SaveResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{1}
}

func (x *SaveResponse) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{2}
}

func (x *GetRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{3}
}

func (x *GetResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntegrationId string `protobuf:"bytes,1,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	ConnectionId  string `protobuf:"bytes,2,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	EventType     string `protobuf:"bytes,3,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	MaxResults    uint32 `protobuf:"varint,4,opt,name=max_results,json=maxResults,proto3" json:"max_results,omitempty"`
	Order         string `protobuf:"bytes,5,opt,name=order,proto3" json:"order,omitempty"` // ASC or DESC values are supported
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{4}
}

func (x *ListRequest) GetIntegrationId() string {
	if x != nil {
		return x.IntegrationId
	}
	return ""
}

func (x *ListRequest) GetConnectionId() string {
	if x != nil {
		return x.ConnectionId
	}
	return ""
}

func (x *ListRequest) GetEventType() string {
	if x != nil {
		return x.EventType
	}
	return ""
}

func (x *ListRequest) GetMaxResults() uint32 {
	if x != nil {
		return x.MaxResults
	}
	return 0
}

func (x *ListRequest) GetOrder() string {
	if x != nil {
		return x.Order
	}
	return ""
}

type ListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Events without their data.
	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListResponse) Reset() {
	*x = ListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponse) ProtoMessage() {}

func (x *ListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponse.ProtoReflect.Descriptor instead.
func (*ListResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{5}
}

func (x *ListResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

type ListEventRecordsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventId string     `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	State   EventState `protobuf:"varint,3,opt,name=state,proto3,enum=autokitteh.events.v1.EventState" json:"state,omitempty"`
}

func (x *ListEventRecordsRequest) Reset() {
	*x = ListEventRecordsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventRecordsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventRecordsRequest) ProtoMessage() {}

func (x *ListEventRecordsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventRecordsRequest.ProtoReflect.Descriptor instead.
func (*ListEventRecordsRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{6}
}

func (x *ListEventRecordsRequest) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *ListEventRecordsRequest) GetState() EventState {
	if x != nil {
		return x.State
	}
	return EventState_EVENT_STATE_UNSPECIFIED
}

type ListEventRecordsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Records []*EventRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
}

func (x *ListEventRecordsResponse) Reset() {
	*x = ListEventRecordsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventRecordsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventRecordsResponse) ProtoMessage() {}

func (x *ListEventRecordsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventRecordsResponse.ProtoReflect.Descriptor instead.
func (*ListEventRecordsResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{7}
}

func (x *ListEventRecordsResponse) GetRecords() []*EventRecord {
	if x != nil {
		return x.Records
	}
	return nil
}

type AddEventRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Record *EventRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (x *AddEventRecordRequest) Reset() {
	*x = AddEventRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRecordRequest) ProtoMessage() {}

func (x *AddEventRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRecordRequest.ProtoReflect.Descriptor instead.
func (*AddEventRecordRequest) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{8}
}

func (x *AddEventRecordRequest) GetRecord() *EventRecord {
	if x != nil {
		return x.Record
	}
	return nil
}

type AddEventRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddEventRecordResponse) Reset() {
	*x = AddEventRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_events_v1_svc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEventRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEventRecordResponse) ProtoMessage() {}

func (x *AddEventRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_events_v1_svc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEventRecordResponse.ProtoReflect.Descriptor instead.
func (*AddEventRecordResponse) Descriptor() ([]byte, []int) {
	return file_autokitteh_events_v1_svc_proto_rawDescGZIP(), []int{9}
}

var File_autokitteh_events_v1_svc_proto protoreflect.FileDescriptor

var file_autokitteh_events_v1_svc_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x14, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x3a, 0xde, 0x01, 0xfa, 0xf7, 0x18, 0xd9, 0x01,
	0x1a, 0x68, 0x0a, 0x1c, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6e, 0x67, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x12, 0x15, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x1a, 0x31, 0x68, 0x61, 0x73, 0x28, 0x74, 0x68, 0x69,
	0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x26, 0x26, 0x20, 0x74, 0x68, 0x69, 0x73,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x1a, 0x6d, 0x0a, 0x1d, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x5f, 0x6d, 0x75,
	0x73, 0x74, 0x5f, 0x62, 0x65, 0x5f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x6d, 0x75, 0x73, 0x74, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x62,
	0x65, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x1a, 0x2c, 0x68, 0x61, 0x73,
	0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x29, 0x20, 0x26, 0x26, 0x20,
	0x74, 0x68, 0x69, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x20, 0x3d, 0x3d, 0x20, 0x27, 0x27, 0x22, 0x33, 0x0a, 0x0c, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x31,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x22, 0x40, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x31, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x22, 0xaf, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x6d, 0x61, 0x78, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x43, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0xf7, 0x18, 0x04, 0x72, 0x02,
	0x10, 0x01, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x41, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x61, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x42, 0x09, 0xfa, 0xf7,
	0x18, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x65,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x42, 0x0c,
	0xfa, 0xf7, 0x18, 0x08, 0x92, 0x01, 0x05, 0x22, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x39, 0x0a, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x52, 0x06, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x3a, 0x70, 0xfa, 0xf7, 0x18, 0x6c,
	0x1a, 0x6a, 0x0a, 0x26, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x68, 0x61, 0x76,
	0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x10, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6e, 0x67, 0x20, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x20, 0x69, 0x64, 0x1a, 0x2e, 0x68, 0x61,
	0x73, 0x28, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x29, 0x20, 0x26,
	0x26, 0x20, 0x74, 0x68, 0x69, 0x73, 0x2e, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x20, 0x21, 0x3d, 0x20, 0x27, 0x27, 0x22, 0x18, 0x0a, 0x16,
	0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xd9, 0x03, 0x0a, 0x0d, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65,
	0x12, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x20,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x21, 0x2e, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x71, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x2d, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74,
	0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x2b, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0xdf, 0x01, 0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42,
	0x08, 0x53, 0x76, 0x63, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x6f, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x45, 0x58, 0xaa, 0x02, 0x14, 0x41, 0x75, 0x74,
	0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x14, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x20, 0x41, 0x75, 0x74, 0x6f, 0x6b,
	0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x16, 0x41, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x3a, 0x3a, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_events_v1_svc_proto_rawDescOnce sync.Once
	file_autokitteh_events_v1_svc_proto_rawDescData = file_autokitteh_events_v1_svc_proto_rawDesc
)

func file_autokitteh_events_v1_svc_proto_rawDescGZIP() []byte {
	file_autokitteh_events_v1_svc_proto_rawDescOnce.Do(func() {
		file_autokitteh_events_v1_svc_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_events_v1_svc_proto_rawDescData)
	})
	return file_autokitteh_events_v1_svc_proto_rawDescData
}

var file_autokitteh_events_v1_svc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_autokitteh_events_v1_svc_proto_goTypes = []interface{}{
	(*SaveRequest)(nil),              // 0: autokitteh.events.v1.SaveRequest
	(*SaveResponse)(nil),             // 1: autokitteh.events.v1.SaveResponse
	(*GetRequest)(nil),               // 2: autokitteh.events.v1.GetRequest
	(*GetResponse)(nil),              // 3: autokitteh.events.v1.GetResponse
	(*ListRequest)(nil),              // 4: autokitteh.events.v1.ListRequest
	(*ListResponse)(nil),             // 5: autokitteh.events.v1.ListResponse
	(*ListEventRecordsRequest)(nil),  // 6: autokitteh.events.v1.ListEventRecordsRequest
	(*ListEventRecordsResponse)(nil), // 7: autokitteh.events.v1.ListEventRecordsResponse
	(*AddEventRecordRequest)(nil),    // 8: autokitteh.events.v1.AddEventRecordRequest
	(*AddEventRecordResponse)(nil),   // 9: autokitteh.events.v1.AddEventRecordResponse
	(*Event)(nil),                    // 10: autokitteh.events.v1.Event
	(EventState)(0),                  // 11: autokitteh.events.v1.EventState
	(*EventRecord)(nil),              // 12: autokitteh.events.v1.EventRecord
}
var file_autokitteh_events_v1_svc_proto_depIdxs = []int32{
	10, // 0: autokitteh.events.v1.SaveRequest.event:type_name -> autokitteh.events.v1.Event
	10, // 1: autokitteh.events.v1.GetResponse.event:type_name -> autokitteh.events.v1.Event
	10, // 2: autokitteh.events.v1.ListResponse.events:type_name -> autokitteh.events.v1.Event
	11, // 3: autokitteh.events.v1.ListEventRecordsRequest.state:type_name -> autokitteh.events.v1.EventState
	12, // 4: autokitteh.events.v1.ListEventRecordsResponse.records:type_name -> autokitteh.events.v1.EventRecord
	12, // 5: autokitteh.events.v1.AddEventRecordRequest.record:type_name -> autokitteh.events.v1.EventRecord
	0,  // 6: autokitteh.events.v1.EventsService.Save:input_type -> autokitteh.events.v1.SaveRequest
	2,  // 7: autokitteh.events.v1.EventsService.Get:input_type -> autokitteh.events.v1.GetRequest
	4,  // 8: autokitteh.events.v1.EventsService.List:input_type -> autokitteh.events.v1.ListRequest
	6,  // 9: autokitteh.events.v1.EventsService.ListEventRecords:input_type -> autokitteh.events.v1.ListEventRecordsRequest
	8,  // 10: autokitteh.events.v1.EventsService.AddEventRecord:input_type -> autokitteh.events.v1.AddEventRecordRequest
	1,  // 11: autokitteh.events.v1.EventsService.Save:output_type -> autokitteh.events.v1.SaveResponse
	3,  // 12: autokitteh.events.v1.EventsService.Get:output_type -> autokitteh.events.v1.GetResponse
	5,  // 13: autokitteh.events.v1.EventsService.List:output_type -> autokitteh.events.v1.ListResponse
	7,  // 14: autokitteh.events.v1.EventsService.ListEventRecords:output_type -> autokitteh.events.v1.ListEventRecordsResponse
	9,  // 15: autokitteh.events.v1.EventsService.AddEventRecord:output_type -> autokitteh.events.v1.AddEventRecordResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_autokitteh_events_v1_svc_proto_init() }
func file_autokitteh_events_v1_svc_proto_init() {
	if File_autokitteh_events_v1_svc_proto != nil {
		return
	}
	file_autokitteh_events_v1_event_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_events_v1_svc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveRequest); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveResponse); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListResponse); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventRecordsRequest); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventRecordsResponse); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRecordRequest); i {
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
		file_autokitteh_events_v1_svc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEventRecordResponse); i {
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
			RawDescriptor: file_autokitteh_events_v1_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_autokitteh_events_v1_svc_proto_goTypes,
		DependencyIndexes: file_autokitteh_events_v1_svc_proto_depIdxs,
		MessageInfos:      file_autokitteh_events_v1_svc_proto_msgTypes,
	}.Build()
	File_autokitteh_events_v1_svc_proto = out.File
	file_autokitteh_events_v1_svc_proto_rawDesc = nil
	file_autokitteh_events_v1_svc_proto_goTypes = nil
	file_autokitteh_events_v1_svc_proto_depIdxs = nil
}
