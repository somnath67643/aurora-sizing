// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: obj_svc.proto

package baseproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_obj_svc_proto protoreflect.FileDescriptor

var file_obj_svc_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6f, 0x62, 0x6a, 0x5f, 0x73, 0x76, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x1a, 0x11, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63,
	0x6d, 0x64, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x63, 0x6d, 0x64, 0x5f, 0x6c, 0x6f, 0x6f,
	0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x15, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e,
	0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1a, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6d, 0x64,
	0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x63, 0x6d, 0x64,
	0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x63, 0x6d, 0x64, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x63, 0x6d, 0x64, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x62, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x27, 0x63, 0x6d, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xa9, 0x10, 0x0a, 0x0a, 0x4f, 0x62,
	0x6a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x1a, 0x26, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x21, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74, 0x1a, 0x29, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b,
	0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a,
	0x2b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x15,
	0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x31, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x30, 0x01, 0x12,
	0x6a, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x2f, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x79, 0x0a, 0x15, 0x6c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d,
	0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x1a, 0x35, 0x2e,
	0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x30, 0x01, 0x12, 0x58, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x62, 0x0a, 0x0e, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x65,
	0x78, 0x74, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x1a, 0x2b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d, 0x0a, 0x17, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65,
	0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x7c, 0x0a, 0x1e, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65,
	0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x32, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x30,
	0x01, 0x12, 0x77, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74,
	0x12, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x1a, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45,
	0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x22, 0x67,
	0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e,
	0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78, 0x74, 0x1a, 0x35, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x45, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x30, 0x01, 0x12, 0x68, 0x0a, 0x10, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e,
	0x69, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x12, 0x25, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x2d,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x69,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x72, 0x0a,
	0x14, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x69, 0x74, 0x5f, 0x64, 0x61, 0x74,
	0x61, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x1a,
	0x30, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e,
	0x69, 0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x73, 0x0a, 0x15, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d,
	0x73, 0x67, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x30, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x96, 0x01, 0x0a, 0x25, 0x67, 0x65, 0x74, 0x5f, 0x69,
	0x6e, 0x64, 0x65, 0x78, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x12, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x78,
	0x74, 0x1a, 0x3a, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x30, 0x01, 0x12,
	0x89, 0x01, 0x0a, 0x1b, 0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x65, 0x73, 0x12,
	0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74,
	0x1a, 0x37, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x79, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x87, 0x01, 0x0a, 0x15,
	0x67, 0x65, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x12, 0x32, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4d, 0x73,
	0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x57,
	0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x3a, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36,
	0x34, 0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67,
	0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_obj_svc_proto_goTypes = []interface{}{
	(*CmdConnect)(nil),                             // 0: org.anonymous.grpc.CmdConnect
	(*CmdConnectExt)(nil),                          // 1: org.anonymous.grpc.CmdConnectExt
	(*CmdLookupByName)(nil),                        // 2: org.anonymous.grpc.CmdLookupByName
	(*CmdNameLookupByType)(nil),                    // 3: org.anonymous.grpc.CmdNameLookupByType
	(*CmdGetByName)(nil),                           // 4: org.anonymous.grpc.CmdGetByName
	(*CmdGetByNameExt)(nil),                        // 5: org.anonymous.grpc.CmdGetByNameExt
	(*CmdGetManyByName)(nil),                       // 6: org.anonymous.grpc.CmdGetManyByName
	(*CmdGetManyByNameExt)(nil),                    // 7: org.anonymous.grpc.CmdGetManyByNameExt
	(*CmdChangeInitData)(nil),                      // 8: org.anonymous.grpc.CmdChangeInitData
	(*CmdChangeInitDataExt)(nil),                   // 9: org.anonymous.grpc.CmdChangeInitDataExt
	(*CmdMsgIndexGetByName)(nil),                   // 10: org.anonymous.grpc.CmdMsgIndexGetByName
	(*CmdMsgIndexGetManyByNameExt)(nil),            // 11: org.anonymous.grpc.CmdMsgIndexGetManyByNameExt
	(*CmdMsgIndexGetByNameByLimit)(nil),            // 12: org.anonymous.grpc.CmdMsgIndexGetByNameByLimit
	(*CmdMsgIndexGetByNameWithClient)(nil),         // 13: org.anonymous.grpc.CmdMsgIndexGetByNameWithClient
	(*CmdConnectResponse)(nil),                     // 14: org.anonymous.grpc.CmdConnectResponse
	(*CmdConnectExtResponse)(nil),                  // 15: org.anonymous.grpc.CmdConnectExtResponse
	(*CmdLookupByNameResponse)(nil),                // 16: org.anonymous.grpc.CmdLookupByNameResponse
	(*CmdLookupByNameResponseStream)(nil),          // 17: org.anonymous.grpc.CmdLookupByNameResponseStream
	(*CmdNameLookupByTypeResponse)(nil),            // 18: org.anonymous.grpc.CmdNameLookupByTypeResponse
	(*CmdNameLookupByTypeResponseStream)(nil),      // 19: org.anonymous.grpc.CmdNameLookupByTypeResponseStream
	(*CmdGetByNameResponse)(nil),                   // 20: org.anonymous.grpc.CmdGetByNameResponse
	(*CmdGetByNameExtResponse)(nil),                // 21: org.anonymous.grpc.CmdGetByNameExtResponse
	(*CmdGetManyByNameResponse)(nil),               // 22: org.anonymous.grpc.CmdGetManyByNameResponse
	(*CmdGetManyByNameResponseStream)(nil),         // 23: org.anonymous.grpc.CmdGetManyByNameResponseStream
	(*CmdGetManyByNameExtResponse)(nil),            // 24: org.anonymous.grpc.CmdGetManyByNameExtResponse
	(*CmdGetManyByNameExtResponseStream)(nil),      // 25: org.anonymous.grpc.CmdGetManyByNameExtResponseStream
	(*CmdChangeInitDataResponse)(nil),              // 26: org.anonymous.grpc.CmdChangeInitDataResponse
	(*CmdChangeInitDataExtResponse)(nil),           // 27: org.anonymous.grpc.CmdChangeInitDataExtResponse
	(*CmdMsgIndexGetByNameResponse)(nil),           // 28: org.anonymous.grpc.CmdMsgIndexGetByNameResponse
	(*CmdMsgIndexGetManyByNameResponseStream)(nil), // 29: org.anonymous.grpc.CmdMsgIndexGetManyByNameResponseStream
	(*CmdMsgIndexGetByNameByLimitResponse)(nil),    // 30: org.anonymous.grpc.CmdMsgIndexGetByNameByLimitResponse
	(*CmdMsgIndexGetByNameWithClientResponse)(nil), // 31: org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse
}
var file_obj_svc_proto_depIdxs = []int32{
	0,  // 0: org.anonymous.grpc.ObjService.connect:input_type -> org.anonymous.grpc.CmdConnect
	1,  // 1: org.anonymous.grpc.ObjService.connect_ext:input_type -> org.anonymous.grpc.CmdConnectExt
	2,  // 2: org.anonymous.grpc.ObjService.lookup_by_name:input_type -> org.anonymous.grpc.CmdLookupByName
	2,  // 3: org.anonymous.grpc.ObjService.lookup_by_name_stream:input_type -> org.anonymous.grpc.CmdLookupByName
	3,  // 4: org.anonymous.grpc.ObjService.lookup_by_type:input_type -> org.anonymous.grpc.CmdNameLookupByType
	3,  // 5: org.anonymous.grpc.ObjService.lookup_by_type_stream:input_type -> org.anonymous.grpc.CmdNameLookupByType
	4,  // 6: org.anonymous.grpc.ObjService.get_object:input_type -> org.anonymous.grpc.CmdGetByName
	5,  // 7: org.anonymous.grpc.ObjService.get_object_ext:input_type -> org.anonymous.grpc.CmdGetByNameExt
	6,  // 8: org.anonymous.grpc.ObjService.get_object_many_by_name:input_type -> org.anonymous.grpc.CmdGetManyByName
	6,  // 9: org.anonymous.grpc.ObjService.get_object_many_by_name_stream:input_type -> org.anonymous.grpc.CmdGetManyByName
	7,  // 10: org.anonymous.grpc.ObjService.get_object_many_by_name_ext:input_type -> org.anonymous.grpc.CmdGetManyByNameExt
	7,  // 11: org.anonymous.grpc.ObjService.get_object_many_by_name_ext_stream:input_type -> org.anonymous.grpc.CmdGetManyByNameExt
	8,  // 12: org.anonymous.grpc.ObjService.change_init_data:input_type -> org.anonymous.grpc.CmdChangeInitData
	9,  // 13: org.anonymous.grpc.ObjService.change_init_data_ext:input_type -> org.anonymous.grpc.CmdChangeInitDataExt
	10, // 14: org.anonymous.grpc.ObjService.get_index_msg_by_name:input_type -> org.anonymous.grpc.CmdMsgIndexGetByName
	11, // 15: org.anonymous.grpc.ObjService.get_index_msg_many_by_name_ext_stream:input_type -> org.anonymous.grpc.CmdMsgIndexGetManyByNameExt
	12, // 16: org.anonymous.grpc.ObjService.get_index_record_in_batches:input_type -> org.anonymous.grpc.CmdMsgIndexGetByNameByLimit
	13, // 17: org.anonymous.grpc.ObjService.get_index_record_many:input_type -> org.anonymous.grpc.CmdMsgIndexGetByNameWithClient
	14, // 18: org.anonymous.grpc.ObjService.connect:output_type -> org.anonymous.grpc.CmdConnectResponse
	15, // 19: org.anonymous.grpc.ObjService.connect_ext:output_type -> org.anonymous.grpc.CmdConnectExtResponse
	16, // 20: org.anonymous.grpc.ObjService.lookup_by_name:output_type -> org.anonymous.grpc.CmdLookupByNameResponse
	17, // 21: org.anonymous.grpc.ObjService.lookup_by_name_stream:output_type -> org.anonymous.grpc.CmdLookupByNameResponseStream
	18, // 22: org.anonymous.grpc.ObjService.lookup_by_type:output_type -> org.anonymous.grpc.CmdNameLookupByTypeResponse
	19, // 23: org.anonymous.grpc.ObjService.lookup_by_type_stream:output_type -> org.anonymous.grpc.CmdNameLookupByTypeResponseStream
	20, // 24: org.anonymous.grpc.ObjService.get_object:output_type -> org.anonymous.grpc.CmdGetByNameResponse
	21, // 25: org.anonymous.grpc.ObjService.get_object_ext:output_type -> org.anonymous.grpc.CmdGetByNameExtResponse
	22, // 26: org.anonymous.grpc.ObjService.get_object_many_by_name:output_type -> org.anonymous.grpc.CmdGetManyByNameResponse
	23, // 27: org.anonymous.grpc.ObjService.get_object_many_by_name_stream:output_type -> org.anonymous.grpc.CmdGetManyByNameResponseStream
	24, // 28: org.anonymous.grpc.ObjService.get_object_many_by_name_ext:output_type -> org.anonymous.grpc.CmdGetManyByNameExtResponse
	25, // 29: org.anonymous.grpc.ObjService.get_object_many_by_name_ext_stream:output_type -> org.anonymous.grpc.CmdGetManyByNameExtResponseStream
	26, // 30: org.anonymous.grpc.ObjService.change_init_data:output_type -> org.anonymous.grpc.CmdChangeInitDataResponse
	27, // 31: org.anonymous.grpc.ObjService.change_init_data_ext:output_type -> org.anonymous.grpc.CmdChangeInitDataExtResponse
	28, // 32: org.anonymous.grpc.ObjService.get_index_msg_by_name:output_type -> org.anonymous.grpc.CmdMsgIndexGetByNameResponse
	29, // 33: org.anonymous.grpc.ObjService.get_index_msg_many_by_name_ext_stream:output_type -> org.anonymous.grpc.CmdMsgIndexGetManyByNameResponseStream
	30, // 34: org.anonymous.grpc.ObjService.get_index_record_in_batches:output_type -> org.anonymous.grpc.CmdMsgIndexGetByNameByLimitResponse
	31, // 35: org.anonymous.grpc.ObjService.get_index_record_many:output_type -> org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse
	18, // [18:36] is the sub-list for method output_type
	0,  // [0:18] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_obj_svc_proto_init() }
func file_obj_svc_proto_init() {
	if File_obj_svc_proto != nil {
		return
	}
	file_cmd_connect_proto_init()
	file_cmd_connect_ext_proto_init()
	file_cmd_lookup_by_name_proto_init()
	file_cmd_lookup_by_type_proto_init()
	file_cmd_get_by_name_proto_init()
	file_cmd_get_by_name_ext_proto_init()
	file_cmd_get_many_by_name_proto_init()
	file_cmd_get_many_by_name_ext_proto_init()
	file_cmd_change_init_data_proto_init()
	file_cmd_change_init_data_ext_proto_init()
	file_cmd_index_get_by_name_proto_init()
	file_cmd_index_get_many_by_name_proto_init()
	file_cmd_index_get_by_name_by_limit_proto_init()
	file_cmd_index_get_by_name_with_client_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_obj_svc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_obj_svc_proto_goTypes,
		DependencyIndexes: file_obj_svc_proto_depIdxs,
	}.Build()
	File_obj_svc_proto = out.File
	file_obj_svc_proto_rawDesc = nil
	file_obj_svc_proto_goTypes = nil
	file_obj_svc_proto_depIdxs = nil
}
