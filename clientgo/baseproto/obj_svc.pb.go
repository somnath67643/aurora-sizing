// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: obj_svc.proto

package baseproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

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
	0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65,
	0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xcd, 0x04, 0x0a, 0x0a, 0x4f, 0x62, 0x6a, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1e,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x1a, 0x26,
	0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5b, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x5f, 0x65, 0x78, 0x74, 0x12, 0x21, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74, 0x1a, 0x29, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d,
	0x64, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x45, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x2b, 0x2e, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x43, 0x6d, 0x64, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6a, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75,
	0x70, 0x5f, 0x62, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x27, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x54, 0x79,
	0x70, 0x65, 0x1a, 0x2f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c,
	0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x58, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x12, 0x20, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e,
	0x61, 0x6d, 0x65, 0x1a, 0x28, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a,
	0x0f, 0x67, 0x65, 0x74, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79,
	0x12, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x2c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47,
	0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34,
	0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_obj_svc_proto_goTypes = []interface{}{
	(*CmdConnect)(nil),                  // 0: org.anonymous.grpc.CmdConnect
	(*CmdConnectExt)(nil),               // 1: org.anonymous.grpc.CmdConnectExt
	(*CmdLookupByName)(nil),             // 2: org.anonymous.grpc.CmdLookupByName
	(*CmdNameLookupByType)(nil),         // 3: org.anonymous.grpc.CmdNameLookupByType
	(*CmdGetByName)(nil),                // 4: org.anonymous.grpc.CmdGetByName
	(*CmdGetManyByName)(nil),            // 5: org.anonymous.grpc.CmdGetManyByName
	(*CmdConnectResponse)(nil),          // 6: org.anonymous.grpc.CmdConnectResponse
	(*CmdConnectExtResponse)(nil),       // 7: org.anonymous.grpc.CmdConnectExtResponse
	(*CmdLookupByNameResponse)(nil),     // 8: org.anonymous.grpc.CmdLookupByNameResponse
	(*CmdNameLookupByTypeResponse)(nil), // 9: org.anonymous.grpc.CmdNameLookupByTypeResponse
	(*CmdGetByNameResponse)(nil),        // 10: org.anonymous.grpc.CmdGetByNameResponse
	(*CmdGetManyByNameResponse)(nil),    // 11: org.anonymous.grpc.CmdGetManyByNameResponse
}
var file_obj_svc_proto_depIdxs = []int32{
	0,  // 0: org.anonymous.grpc.ObjService.connect:input_type -> org.anonymous.grpc.CmdConnect
	1,  // 1: org.anonymous.grpc.ObjService.connect_ext:input_type -> org.anonymous.grpc.CmdConnectExt
	2,  // 2: org.anonymous.grpc.ObjService.lookup_by_name:input_type -> org.anonymous.grpc.CmdLookupByName
	3,  // 3: org.anonymous.grpc.ObjService.lookup_by_type:input_type -> org.anonymous.grpc.CmdNameLookupByType
	4,  // 4: org.anonymous.grpc.ObjService.get_object:input_type -> org.anonymous.grpc.CmdGetByName
	5,  // 5: org.anonymous.grpc.ObjService.get_object_many:input_type -> org.anonymous.grpc.CmdGetManyByName
	6,  // 6: org.anonymous.grpc.ObjService.connect:output_type -> org.anonymous.grpc.CmdConnectResponse
	7,  // 7: org.anonymous.grpc.ObjService.connect_ext:output_type -> org.anonymous.grpc.CmdConnectExtResponse
	8,  // 8: org.anonymous.grpc.ObjService.lookup_by_name:output_type -> org.anonymous.grpc.CmdLookupByNameResponse
	9,  // 9: org.anonymous.grpc.ObjService.lookup_by_type:output_type -> org.anonymous.grpc.CmdNameLookupByTypeResponse
	10, // 10: org.anonymous.grpc.ObjService.get_object:output_type -> org.anonymous.grpc.CmdGetByNameResponse
	11, // 11: org.anonymous.grpc.ObjService.get_object_many:output_type -> org.anonymous.grpc.CmdGetManyByNameResponse
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
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
	file_cmd_get_many_by_name_proto_init()
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
