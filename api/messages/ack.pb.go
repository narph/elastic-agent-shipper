// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: messages/ack.proto

package messages

import (
	reflect "reflect"
	sync "sync"

	status "google.golang.org/genproto/googleapis/rpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StreamAcksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Optional. Requests acknowledgements originating only from this source.
	Source *Source `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// Optional. Requests acknowledgements for events for this data stream only.
	DataStream *DataStream `protobuf:"bytes,2,opt,name=data_stream,json=dataStream,proto3" json:"data_stream,omitempty"`
}

func (x *StreamAcksRequest) Reset() {
	*x = StreamAcksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ack_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAcksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAcksRequest) ProtoMessage() {}

func (x *StreamAcksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ack_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAcksRequest.ProtoReflect.Descriptor instead.
func (*StreamAcksRequest) Descriptor() ([]byte, []int) {
	return file_messages_ack_proto_rawDescGZIP(), []int{0}
}

func (x *StreamAcksRequest) GetSource() *Source {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *StreamAcksRequest) GetDataStream() *DataStream {
	if x != nil {
		return x.DataStream
	}
	return nil
}

type StreamAcksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acks []*Acknowledgement `protobuf:"bytes,1,rep,name=acks,proto3" json:"acks,omitempty"`
}

func (x *StreamAcksReply) Reset() {
	*x = StreamAcksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ack_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamAcksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamAcksReply) ProtoMessage() {}

func (x *StreamAcksReply) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ack_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamAcksReply.ProtoReflect.Descriptor instead.
func (*StreamAcksReply) Descriptor() ([]byte, []int) {
	return file_messages_ack_proto_rawDescGZIP(), []int{1}
}

func (x *StreamAcksReply) GetAcks() []*Acknowledgement {
	if x != nil {
		return x.Acks
	}
	return nil
}

type Acknowledgement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Timestamp of the event being acknowledged.
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// ID derived from the queue offset that uniquely identifies the event within the shipper
	// system. Used to track the progress of events through the queue system.
	QueueId string `protobuf:"bytes,2,opt,name=queue_id,json=queueId,proto3" json:"queue_id,omitempty"`
	// Optional. Event ID provided when the event was published, if it exists.
	EventId string `protobuf:"bytes,3,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	// Optional. Error status indicating the message failed to send.
	Error *status.Status `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Acknowledgement) Reset() {
	*x = Acknowledgement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_messages_ack_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acknowledgement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acknowledgement) ProtoMessage() {}

func (x *Acknowledgement) ProtoReflect() protoreflect.Message {
	mi := &file_messages_ack_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acknowledgement.ProtoReflect.Descriptor instead.
func (*Acknowledgement) Descriptor() ([]byte, []int) {
	return file_messages_ack_proto_rawDescGZIP(), []int{2}
}

func (x *Acknowledgement) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Acknowledgement) GetQueueId() string {
	if x != nil {
		return x.QueueId
	}
	return ""
}

func (x *Acknowledgement) GetEventId() string {
	if x != nil {
		return x.EventId
	}
	return ""
}

func (x *Acknowledgement) GetError() *status.Status {
	if x != nil {
		return x.Error
	}
	return nil
}

var File_messages_ack_proto protoreflect.FileDescriptor

var file_messages_ack_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x6b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x01, 0x0a, 0x11, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x41, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x4e, 0x0a, 0x0b, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x22, 0x59, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65,
	0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x22, 0xab, 0x01,
	0x0a, 0x0f, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x37, 0x5a, 0x35, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69,
	0x63, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2d,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_messages_ack_proto_rawDescOnce sync.Once
	file_messages_ack_proto_rawDescData = file_messages_ack_proto_rawDesc
)

func file_messages_ack_proto_rawDescGZIP() []byte {
	file_messages_ack_proto_rawDescOnce.Do(func() {
		file_messages_ack_proto_rawDescData = protoimpl.X.CompressGZIP(file_messages_ack_proto_rawDescData)
	})
	return file_messages_ack_proto_rawDescData
}

var file_messages_ack_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_messages_ack_proto_goTypes = []interface{}{
	(*StreamAcksRequest)(nil),     // 0: elastic.agent.shipper.v1.messages.StreamAcksRequest
	(*StreamAcksReply)(nil),       // 1: elastic.agent.shipper.v1.messages.StreamAcksReply
	(*Acknowledgement)(nil),       // 2: elastic.agent.shipper.v1.messages.Acknowledgement
	(*Source)(nil),                // 3: elastic.agent.shipper.v1.messages.Source
	(*DataStream)(nil),            // 4: elastic.agent.shipper.v1.messages.DataStream
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*status.Status)(nil),         // 6: google.rpc.Status
}
var file_messages_ack_proto_depIdxs = []int32{
	3, // 0: elastic.agent.shipper.v1.messages.StreamAcksRequest.source:type_name -> elastic.agent.shipper.v1.messages.Source
	4, // 1: elastic.agent.shipper.v1.messages.StreamAcksRequest.data_stream:type_name -> elastic.agent.shipper.v1.messages.DataStream
	2, // 2: elastic.agent.shipper.v1.messages.StreamAcksReply.acks:type_name -> elastic.agent.shipper.v1.messages.Acknowledgement
	5, // 3: elastic.agent.shipper.v1.messages.Acknowledgement.timestamp:type_name -> google.protobuf.Timestamp
	6, // 4: elastic.agent.shipper.v1.messages.Acknowledgement.error:type_name -> google.rpc.Status
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_messages_ack_proto_init() }
func file_messages_ack_proto_init() {
	if File_messages_ack_proto != nil {
		return
	}
	file_messages_publish_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_messages_ack_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAcksRequest); i {
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
		file_messages_ack_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamAcksReply); i {
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
		file_messages_ack_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acknowledgement); i {
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
			RawDescriptor: file_messages_ack_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_messages_ack_proto_goTypes,
		DependencyIndexes: file_messages_ack_proto_depIdxs,
		MessageInfos:      file_messages_ack_proto_msgTypes,
	}.Build()
	File_messages_ack_proto = out.File
	file_messages_ack_proto_rawDesc = nil
	file_messages_ack_proto_goTypes = nil
	file_messages_ack_proto_depIdxs = nil
}
