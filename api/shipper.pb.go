// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: shipper.proto

package api

import (
	reflect "reflect"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	messages "github.com/elastic/elastic-agent-shipper/api/messages"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_shipper_proto protoreflect.FileDescriptor

var file_shipper_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x18, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73,
	0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x16, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x86, 0x02, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x12, 0x73, 0x0a, 0x0d, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x12, 0x31, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63,
	0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x73, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x84, 0x01, 0x0a, 0x16, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x41, 0x63, 0x6b, 0x6e, 0x6f, 0x77, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x34, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2e, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x41, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x65, 0x6c, 0x61, 0x73, 0x74,
	0x69, 0x63, 0x2e, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x2e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x41, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x30, 0x01, 0x42, 0x2e,
	0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6c, 0x61,
	0x73, 0x74, 0x69, 0x63, 0x2f, 0x65, 0x6c, 0x61, 0x73, 0x74, 0x69, 0x63, 0x2d, 0x61, 0x67, 0x65,
	0x6e, 0x74, 0x2d, 0x73, 0x68, 0x69, 0x70, 0x70, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_shipper_proto_goTypes = []interface{}{
	(*messages.PublishRequest)(nil),    // 0: elastic.agent.shipper.v1.messages.PublishRequest
	(*messages.StreamAcksRequest)(nil), // 1: elastic.agent.shipper.v1.messages.StreamAcksRequest
	(*messages.PublishReply)(nil),      // 2: elastic.agent.shipper.v1.messages.PublishReply
	(*messages.StreamAcksReply)(nil),   // 3: elastic.agent.shipper.v1.messages.StreamAcksReply
}
var file_shipper_proto_depIdxs = []int32{
	0, // 0: elastic.agent.shipper.v1.Producer.PublishEvents:input_type -> elastic.agent.shipper.v1.messages.PublishRequest
	1, // 1: elastic.agent.shipper.v1.Producer.StreamAcknowledgements:input_type -> elastic.agent.shipper.v1.messages.StreamAcksRequest
	2, // 2: elastic.agent.shipper.v1.Producer.PublishEvents:output_type -> elastic.agent.shipper.v1.messages.PublishReply
	3, // 3: elastic.agent.shipper.v1.Producer.StreamAcknowledgements:output_type -> elastic.agent.shipper.v1.messages.StreamAcksReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shipper_proto_init() }
func file_shipper_proto_init() {
	if File_shipper_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_shipper_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shipper_proto_goTypes,
		DependencyIndexes: file_shipper_proto_depIdxs,
	}.Build()
	File_shipper_proto = out.File
	file_shipper_proto_rawDesc = nil
	file_shipper_proto_goTypes = nil
	file_shipper_proto_depIdxs = nil
}
