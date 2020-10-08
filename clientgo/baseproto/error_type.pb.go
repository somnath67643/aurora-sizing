// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: error_type.proto

package baseproto

import (
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

type ErrorType int32

const (
	ErrorType_ERR_INVALID_ARGUMENTS                ErrorType = 0
	ErrorType_SDB_ERR_CONNECTION_LIMIT             ErrorType = 1
	ErrorType_SDB_ERR_AUTH_FAIL_NO_CREDENTIAL      ErrorType = 2
	ErrorType_SDB_ERR_AUTH_FAIL_INVALID_CREDENTIAL ErrorType = 3
	ErrorType_SDB_ERR_PERMISSION_DENIED            ErrorType = 4
	ErrorType_SDB_ERR_UNSUPPORTED_VERSION          ErrorType = 5
	ErrorType_SDB_ERR_INDEX_NOT_FOUND              ErrorType = 6
	ErrorType_SDB_ERR_OBJECT_ALREADY_EXISTS        ErrorType = 7
	ErrorType_ERR_UNSUPPORTED_OPERATION            ErrorType = 8
	ErrorType_SDB_ERR_OBJECT_NOT_FOUND             ErrorType = 9
	ErrorType_SDB_ERR_OBJECT_INVALID               ErrorType = 10
	ErrorType_ERR_OBJECT_NOT_FOUND                 ErrorType = 11
	ErrorType_ERR_PERMISSION_DENIED                ErrorType = 12
	ErrorType_ERR_NOT_CONSISTENT                   ErrorType = 13
	ErrorType_ERR_OVERFLOW                         ErrorType = 14
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0:  "ERR_INVALID_ARGUMENTS",
		1:  "SDB_ERR_CONNECTION_LIMIT",
		2:  "SDB_ERR_AUTH_FAIL_NO_CREDENTIAL",
		3:  "SDB_ERR_AUTH_FAIL_INVALID_CREDENTIAL",
		4:  "SDB_ERR_PERMISSION_DENIED",
		5:  "SDB_ERR_UNSUPPORTED_VERSION",
		6:  "SDB_ERR_INDEX_NOT_FOUND",
		7:  "SDB_ERR_OBJECT_ALREADY_EXISTS",
		8:  "ERR_UNSUPPORTED_OPERATION",
		9:  "SDB_ERR_OBJECT_NOT_FOUND",
		10: "SDB_ERR_OBJECT_INVALID",
		11: "ERR_OBJECT_NOT_FOUND",
		12: "ERR_PERMISSION_DENIED",
		13: "ERR_NOT_CONSISTENT",
		14: "ERR_OVERFLOW",
	}
	ErrorType_value = map[string]int32{
		"ERR_INVALID_ARGUMENTS":                0,
		"SDB_ERR_CONNECTION_LIMIT":             1,
		"SDB_ERR_AUTH_FAIL_NO_CREDENTIAL":      2,
		"SDB_ERR_AUTH_FAIL_INVALID_CREDENTIAL": 3,
		"SDB_ERR_PERMISSION_DENIED":            4,
		"SDB_ERR_UNSUPPORTED_VERSION":          5,
		"SDB_ERR_INDEX_NOT_FOUND":              6,
		"SDB_ERR_OBJECT_ALREADY_EXISTS":        7,
		"ERR_UNSUPPORTED_OPERATION":            8,
		"SDB_ERR_OBJECT_NOT_FOUND":             9,
		"SDB_ERR_OBJECT_INVALID":               10,
		"ERR_OBJECT_NOT_FOUND":                 11,
		"ERR_PERMISSION_DENIED":                12,
		"ERR_NOT_CONSISTENT":                   13,
		"ERR_OVERFLOW":                         14,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_error_type_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_error_type_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_error_type_proto_rawDescGZIP(), []int{0}
}

var File_error_type_proto protoreflect.FileDescriptor

var file_error_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2a, 0xcb, 0x03, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x53, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x53, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10, 0x01, 0x12, 0x23, 0x0a,
	0x1f, 0x53, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x5f, 0x4e, 0x4f, 0x5f, 0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c,
	0x10, 0x02, 0x12, 0x28, 0x0a, 0x24, 0x53, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x41, 0x55,
	0x54, 0x48, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f,
	0x43, 0x52, 0x45, 0x44, 0x45, 0x4e, 0x54, 0x49, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x1d, 0x0a, 0x19,
	0x53, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49,
	0x4f, 0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x04, 0x12, 0x1f, 0x0a, 0x1b, 0x53,
	0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x1b, 0x0a, 0x17,
	0x53, 0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x06, 0x12, 0x21, 0x0a, 0x1d, 0x53, 0x44, 0x42,
	0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x41, 0x4c, 0x52, 0x45,
	0x41, 0x44, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x53, 0x10, 0x07, 0x12, 0x1d, 0x0a, 0x19,
	0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x45, 0x44, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x1c, 0x0a, 0x18, 0x53,
	0x44, 0x42, 0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f,
	0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x09, 0x12, 0x1a, 0x0a, 0x16, 0x53, 0x44, 0x42,
	0x5f, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x42, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x10, 0x0a, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x42, 0x4a,
	0x45, 0x43, 0x54, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x0b, 0x12,
	0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x5f, 0x50, 0x45, 0x52, 0x4d, 0x49, 0x53, 0x53, 0x49, 0x4f,
	0x4e, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x0c, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52,
	0x52, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x53, 0x49, 0x53, 0x54, 0x45, 0x4e, 0x54,
	0x10, 0x0d, 0x12, 0x10, 0x0a, 0x0c, 0x45, 0x52, 0x52, 0x5f, 0x4f, 0x56, 0x45, 0x52, 0x46, 0x4c,
	0x4f, 0x57, 0x10, 0x0e, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34,
	0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_type_proto_rawDescOnce sync.Once
	file_error_type_proto_rawDescData = file_error_type_proto_rawDesc
)

func file_error_type_proto_rawDescGZIP() []byte {
	file_error_type_proto_rawDescOnce.Do(func() {
		file_error_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_type_proto_rawDescData)
	})
	return file_error_type_proto_rawDescData
}

var file_error_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_type_proto_goTypes = []interface{}{
	(ErrorType)(0), // 0: org.anonymous.grpc.ErrorType
}
var file_error_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_type_proto_init() }
func file_error_type_proto_init() {
	if File_error_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_type_proto_goTypes,
		DependencyIndexes: file_error_type_proto_depIdxs,
		EnumInfos:         file_error_type_proto_enumTypes,
	}.Build()
	File_error_type_proto = out.File
	file_error_type_proto_rawDesc = nil
	file_error_type_proto_goTypes = nil
	file_error_type_proto_depIdxs = nil
}
