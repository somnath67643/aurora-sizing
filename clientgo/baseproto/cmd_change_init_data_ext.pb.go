// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_change_init_data_ext.proto

package baseproto

import (
	proto "github.com/golang/protobuf/proto"
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

type CmdChangeInitDataExt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSize    int32   `protobuf:"varint,1,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	MessageType    CmdType `protobuf:"varint,2,opt,name=message_type,json=messageType,proto3,enum=org.anonymous.grpc.CmdType" json:"message_type,omitempty"`
	RequestVersion uint32  `protobuf:"varint,3,opt,name=request_version,json=requestVersion,proto3" json:"request_version,omitempty"`
	AppNameShort   string  `protobuf:"bytes,4,opt,name=app_name_short,json=appNameShort,proto3" json:"app_name_short,omitempty"`
	UserName       string  `protobuf:"bytes,5,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	// Types that are assignable to Request:
	//	*CmdChangeInitDataExt_NameDifferenceIsNull
	//	*CmdChangeInitDataExt_AppNameFull
	Request isCmdChangeInitDataExt_Request `protobuf_oneof:"Request"`
}

func (x *CmdChangeInitDataExt) Reset() {
	*x = CmdChangeInitDataExt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_change_init_data_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdChangeInitDataExt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdChangeInitDataExt) ProtoMessage() {}

func (x *CmdChangeInitDataExt) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_change_init_data_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdChangeInitDataExt.ProtoReflect.Descriptor instead.
func (*CmdChangeInitDataExt) Descriptor() ([]byte, []int) {
	return file_cmd_change_init_data_ext_proto_rawDescGZIP(), []int{0}
}

func (x *CmdChangeInitDataExt) GetMessageSize() int32 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *CmdChangeInitDataExt) GetMessageType() CmdType {
	if x != nil {
		return x.MessageType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdChangeInitDataExt) GetRequestVersion() uint32 {
	if x != nil {
		return x.RequestVersion
	}
	return 0
}

func (x *CmdChangeInitDataExt) GetAppNameShort() string {
	if x != nil {
		return x.AppNameShort
	}
	return ""
}

func (x *CmdChangeInitDataExt) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (m *CmdChangeInitDataExt) GetRequest() isCmdChangeInitDataExt_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *CmdChangeInitDataExt) GetNameDifferenceIsNull() bool {
	if x, ok := x.GetRequest().(*CmdChangeInitDataExt_NameDifferenceIsNull); ok {
		return x.NameDifferenceIsNull
	}
	return false
}

func (x *CmdChangeInitDataExt) GetAppNameFull() string {
	if x, ok := x.GetRequest().(*CmdChangeInitDataExt_AppNameFull); ok {
		return x.AppNameFull
	}
	return ""
}

type isCmdChangeInitDataExt_Request interface {
	isCmdChangeInitDataExt_Request()
}

type CmdChangeInitDataExt_NameDifferenceIsNull struct {
	NameDifferenceIsNull bool `protobuf:"varint,6,opt,name=name_difference_is_null,json=nameDifferenceIsNull,proto3,oneof"`
}

type CmdChangeInitDataExt_AppNameFull struct {
	AppNameFull string `protobuf:"bytes,7,opt,name=app_name_full,json=appNameFull,proto3,oneof"`
}

func (*CmdChangeInitDataExt_NameDifferenceIsNull) isCmdChangeInitDataExt_Request() {}

func (*CmdChangeInitDataExt_AppNameFull) isCmdChangeInitDataExt_Request() {}

type CmdChangeInitDataExtResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResponseVersion uint32         `protobuf:"varint,1,opt,name=response_version,json=responseVersion,proto3" json:"response_version,omitempty"`
	Status          ResponseStatus `protobuf:"varint,2,opt,name=status,proto3,enum=org.anonymous.grpc.ResponseStatus" json:"status,omitempty"`
	// Types that are assignable to RequestResponse:
	//	*CmdChangeInitDataExtResponse_HasSucceeded
	//	*CmdChangeInitDataExtResponse_ErrorType
	RequestResponse isCmdChangeInitDataExtResponse_RequestResponse `protobuf_oneof:"RequestResponse"`
}

func (x *CmdChangeInitDataExtResponse) Reset() {
	*x = CmdChangeInitDataExtResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_change_init_data_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdChangeInitDataExtResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdChangeInitDataExtResponse) ProtoMessage() {}

func (x *CmdChangeInitDataExtResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_change_init_data_ext_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdChangeInitDataExtResponse.ProtoReflect.Descriptor instead.
func (*CmdChangeInitDataExtResponse) Descriptor() ([]byte, []int) {
	return file_cmd_change_init_data_ext_proto_rawDescGZIP(), []int{1}
}

func (x *CmdChangeInitDataExtResponse) GetResponseVersion() uint32 {
	if x != nil {
		return x.ResponseVersion
	}
	return 0
}

func (x *CmdChangeInitDataExtResponse) GetStatus() ResponseStatus {
	if x != nil {
		return x.Status
	}
	return ResponseStatus_SUCCESS
}

func (m *CmdChangeInitDataExtResponse) GetRequestResponse() isCmdChangeInitDataExtResponse_RequestResponse {
	if m != nil {
		return m.RequestResponse
	}
	return nil
}

func (x *CmdChangeInitDataExtResponse) GetHasSucceeded() bool {
	if x, ok := x.GetRequestResponse().(*CmdChangeInitDataExtResponse_HasSucceeded); ok {
		return x.HasSucceeded
	}
	return false
}

func (x *CmdChangeInitDataExtResponse) GetErrorType() ErrorType {
	if x, ok := x.GetRequestResponse().(*CmdChangeInitDataExtResponse_ErrorType); ok {
		return x.ErrorType
	}
	return ErrorType_ERR_INVALID_ARGUMENTS
}

type isCmdChangeInitDataExtResponse_RequestResponse interface {
	isCmdChangeInitDataExtResponse_RequestResponse()
}

type CmdChangeInitDataExtResponse_HasSucceeded struct {
	HasSucceeded bool `protobuf:"varint,3,opt,name=has_succeeded,json=hasSucceeded,proto3,oneof"`
}

type CmdChangeInitDataExtResponse_ErrorType struct {
	ErrorType ErrorType `protobuf:"varint,4,opt,name=error_type,json=errorType,proto3,enum=org.anonymous.grpc.ErrorType,oneof"`
}

func (*CmdChangeInitDataExtResponse_HasSucceeded) isCmdChangeInitDataExtResponse_RequestResponse() {}

func (*CmdChangeInitDataExtResponse_ErrorType) isCmdChangeInitDataExtResponse_RequestResponse() {}

var File_cmd_change_init_data_ext_proto protoreflect.FileDescriptor

var file_cmd_change_init_data_ext_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x63, 0x6d, 0x64, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x69,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02,
	0x0a, 0x14, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x69, 0x74, 0x44,
	0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x70, 0x70, 0x4e,
	0x61, 0x6d, 0x65, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x17, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x64, 0x69,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x5f, 0x69, 0x73, 0x5f, 0x6e, 0x75, 0x6c, 0x6c,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x14, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x69,
	0x66, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x73, 0x4e, 0x75, 0x6c, 0x6c, 0x12, 0x24,
	0x0a, 0x0d, 0x61, 0x70, 0x70, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x66, 0x75, 0x6c, 0x6c, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x46, 0x75, 0x6c, 0x6c, 0x42, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0xff, 0x01, 0x0a, 0x1c, 0x43, 0x6d, 0x64, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x49, 0x6e, 0x69,
	0x74, 0x44, 0x61, 0x74, 0x61, 0x45, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00,
	0x52, 0x0c, 0x68, 0x61, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x3e,
	0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70,
	0x65, 0x48, 0x00, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42, 0x11,
	0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61,
	0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_change_init_data_ext_proto_rawDescOnce sync.Once
	file_cmd_change_init_data_ext_proto_rawDescData = file_cmd_change_init_data_ext_proto_rawDesc
)

func file_cmd_change_init_data_ext_proto_rawDescGZIP() []byte {
	file_cmd_change_init_data_ext_proto_rawDescOnce.Do(func() {
		file_cmd_change_init_data_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_change_init_data_ext_proto_rawDescData)
	})
	return file_cmd_change_init_data_ext_proto_rawDescData
}

var file_cmd_change_init_data_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_change_init_data_ext_proto_goTypes = []interface{}{
	(*CmdChangeInitDataExt)(nil),         // 0: org.anonymous.grpc.CmdChangeInitDataExt
	(*CmdChangeInitDataExtResponse)(nil), // 1: org.anonymous.grpc.CmdChangeInitDataExtResponse
	(CmdType)(0),                         // 2: org.anonymous.grpc.CmdType
	(ResponseStatus)(0),                  // 3: org.anonymous.grpc.ResponseStatus
	(ErrorType)(0),                       // 4: org.anonymous.grpc.ErrorType
}
var file_cmd_change_init_data_ext_proto_depIdxs = []int32{
	2, // 0: org.anonymous.grpc.CmdChangeInitDataExt.message_type:type_name -> org.anonymous.grpc.CmdType
	3, // 1: org.anonymous.grpc.CmdChangeInitDataExtResponse.status:type_name -> org.anonymous.grpc.ResponseStatus
	4, // 2: org.anonymous.grpc.CmdChangeInitDataExtResponse.error_type:type_name -> org.anonymous.grpc.ErrorType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_cmd_change_init_data_ext_proto_init() }
func file_cmd_change_init_data_ext_proto_init() {
	if File_cmd_change_init_data_ext_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	file_response_status_proto_init()
	file_error_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_change_init_data_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdChangeInitDataExt); i {
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
		file_cmd_change_init_data_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdChangeInitDataExtResponse); i {
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
	file_cmd_change_init_data_ext_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*CmdChangeInitDataExt_NameDifferenceIsNull)(nil),
		(*CmdChangeInitDataExt_AppNameFull)(nil),
	}
	file_cmd_change_init_data_ext_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CmdChangeInitDataExtResponse_HasSucceeded)(nil),
		(*CmdChangeInitDataExtResponse_ErrorType)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_change_init_data_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_change_init_data_ext_proto_goTypes,
		DependencyIndexes: file_cmd_change_init_data_ext_proto_depIdxs,
		MessageInfos:      file_cmd_change_init_data_ext_proto_msgTypes,
	}.Build()
	File_cmd_change_init_data_ext_proto = out.File
	file_cmd_change_init_data_ext_proto_rawDesc = nil
	file_cmd_change_init_data_ext_proto_goTypes = nil
	file_cmd_change_init_data_ext_proto_depIdxs = nil
}