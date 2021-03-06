// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_get_by_name.proto

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

type CmdGetByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSize      int32   `protobuf:"varint,1,opt,name=msg_size,json=msgSize,proto3" json:"msg_size,omitempty"`
	MsgType      CmdType `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3,enum=org.anonymous.grpc.CmdType" json:"msg_type,omitempty"`
	SecurityName string  `protobuf:"bytes,3,opt,name=security_name,json=securityName,proto3" json:"security_name,omitempty"`
}

func (x *CmdGetByName) Reset() {
	*x = CmdGetByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_by_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetByName) ProtoMessage() {}

func (x *CmdGetByName) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_by_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetByName.ProtoReflect.Descriptor instead.
func (*CmdGetByName) Descriptor() ([]byte, []int) {
	return file_cmd_get_by_name_proto_rawDescGZIP(), []int{0}
}

func (x *CmdGetByName) GetMsgSize() int32 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *CmdGetByName) GetMsgType() CmdType {
	if x != nil {
		return x.MsgType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdGetByName) GetSecurityName() string {
	if x != nil {
		return x.SecurityName
	}
	return ""
}

type CmdGetByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSize int32 `protobuf:"varint,1,opt,name=msg_size,json=msgSize,proto3" json:"msg_size,omitempty"`
	Status  int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to Response:
	//	*CmdGetByNameResponse_ErrorType
	//	*CmdGetByNameResponse_Security
	Response isCmdGetByNameResponse_Response `protobuf_oneof:"response"`
}

func (x *CmdGetByNameResponse) Reset() {
	*x = CmdGetByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_by_name_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetByNameResponse) ProtoMessage() {}

func (x *CmdGetByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_by_name_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetByNameResponse.ProtoReflect.Descriptor instead.
func (*CmdGetByNameResponse) Descriptor() ([]byte, []int) {
	return file_cmd_get_by_name_proto_rawDescGZIP(), []int{1}
}

func (x *CmdGetByNameResponse) GetMsgSize() int32 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *CmdGetByNameResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (m *CmdGetByNameResponse) GetResponse() isCmdGetByNameResponse_Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func (x *CmdGetByNameResponse) GetErrorType() ErrorType {
	if x, ok := x.GetResponse().(*CmdGetByNameResponse_ErrorType); ok {
		return x.ErrorType
	}
	return ErrorType_ERR_INVALID_ARGUMENTS
}

func (x *CmdGetByNameResponse) GetSecurity() []byte {
	if x, ok := x.GetResponse().(*CmdGetByNameResponse_Security); ok {
		return x.Security
	}
	return nil
}

type isCmdGetByNameResponse_Response interface {
	isCmdGetByNameResponse_Response()
}

type CmdGetByNameResponse_ErrorType struct {
	ErrorType ErrorType `protobuf:"varint,3,opt,name=error_type,json=errorType,proto3,enum=org.anonymous.grpc.ErrorType,oneof"`
}

type CmdGetByNameResponse_Security struct {
	Security []byte `protobuf:"bytes,4,opt,name=security,proto3,oneof"`
}

func (*CmdGetByNameResponse_ErrorType) isCmdGetByNameResponse_Response() {}

func (*CmdGetByNameResponse_Security) isCmdGetByNameResponse_Response() {}

var File_cmd_get_by_name_proto protoreflect.FileDescriptor

var file_cmd_get_by_name_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x63, 0x6d, 0x64,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01,
	0x0a, 0x0c, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x73, 0x67,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x14, 0x43, 0x6d, 0x64, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3e, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x54, 0x79, 0x70, 0x65, 0x48, 0x00, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x08, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x42, 0x0a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d,
	0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61,
	0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cmd_get_by_name_proto_rawDescOnce sync.Once
	file_cmd_get_by_name_proto_rawDescData = file_cmd_get_by_name_proto_rawDesc
)

func file_cmd_get_by_name_proto_rawDescGZIP() []byte {
	file_cmd_get_by_name_proto_rawDescOnce.Do(func() {
		file_cmd_get_by_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_get_by_name_proto_rawDescData)
	})
	return file_cmd_get_by_name_proto_rawDescData
}

var file_cmd_get_by_name_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_get_by_name_proto_goTypes = []interface{}{
	(*CmdGetByName)(nil),         // 0: org.anonymous.grpc.CmdGetByName
	(*CmdGetByNameResponse)(nil), // 1: org.anonymous.grpc.CmdGetByNameResponse
	(CmdType)(0),                 // 2: org.anonymous.grpc.CmdType
	(ErrorType)(0),               // 3: org.anonymous.grpc.ErrorType
}
var file_cmd_get_by_name_proto_depIdxs = []int32{
	2, // 0: org.anonymous.grpc.CmdGetByName.msg_type:type_name -> org.anonymous.grpc.CmdType
	3, // 1: org.anonymous.grpc.CmdGetByNameResponse.error_type:type_name -> org.anonymous.grpc.ErrorType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_get_by_name_proto_init() }
func file_cmd_get_by_name_proto_init() {
	if File_cmd_get_by_name_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	file_error_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_get_by_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetByName); i {
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
		file_cmd_get_by_name_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetByNameResponse); i {
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
	file_cmd_get_by_name_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CmdGetByNameResponse_ErrorType)(nil),
		(*CmdGetByNameResponse_Security)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_get_by_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_get_by_name_proto_goTypes,
		DependencyIndexes: file_cmd_get_by_name_proto_depIdxs,
		MessageInfos:      file_cmd_get_by_name_proto_msgTypes,
	}.Build()
	File_cmd_get_by_name_proto = out.File
	file_cmd_get_by_name_proto_rawDesc = nil
	file_cmd_get_by_name_proto_goTypes = nil
	file_cmd_get_by_name_proto_depIdxs = nil
}
