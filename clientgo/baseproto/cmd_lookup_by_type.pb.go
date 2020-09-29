// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_lookup_by_type.proto

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

type CmdNameLookupByType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSize        int32   `protobuf:"varint,1,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	MessageType        CmdType `protobuf:"varint,2,opt,name=message_type,json=messageType,proto3,enum=org.anonymous.grpc.CmdType" json:"message_type,omitempty"`
	GetType            GetType `protobuf:"varint,3,opt,name=get_type,json=getType,proto3,enum=org.anonymous.grpc.GetType" json:"get_type,omitempty"`
	Count              int32   `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	SecurityType       uint32  `protobuf:"varint,5,opt,name=security_type,json=securityType,proto3" json:"security_type,omitempty"`
	SecurityNamePrefix string  `protobuf:"bytes,6,opt,name=security_name_prefix,json=securityNamePrefix,proto3" json:"security_name_prefix,omitempty"`
}

func (x *CmdNameLookupByType) Reset() {
	*x = CmdNameLookupByType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_lookup_by_type_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdNameLookupByType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdNameLookupByType) ProtoMessage() {}

func (x *CmdNameLookupByType) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_lookup_by_type_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdNameLookupByType.ProtoReflect.Descriptor instead.
func (*CmdNameLookupByType) Descriptor() ([]byte, []int) {
	return file_cmd_lookup_by_type_proto_rawDescGZIP(), []int{0}
}

func (x *CmdNameLookupByType) GetMessageSize() int32 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *CmdNameLookupByType) GetMessageType() CmdType {
	if x != nil {
		return x.MessageType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdNameLookupByType) GetGetType() GetType {
	if x != nil {
		return x.GetType
	}
	return GetType_METADATA_GET_FIRST
}

func (x *CmdNameLookupByType) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CmdNameLookupByType) GetSecurityType() uint32 {
	if x != nil {
		return x.SecurityType
	}
	return 0
}

func (x *CmdNameLookupByType) GetSecurityNamePrefix() string {
	if x != nil {
		return x.SecurityNamePrefix
	}
	return ""
}

type CmdNameLookupByTypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSize   uint32   `protobuf:"varint,1,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	SecurityCount int32    `protobuf:"varint,2,opt,name=security_count,json=securityCount,proto3" json:"security_count,omitempty"`
	SecurityNames []string `protobuf:"bytes,3,rep,name=security_names,json=securityNames,proto3" json:"security_names,omitempty"`
}

func (x *CmdNameLookupByTypeResponse) Reset() {
	*x = CmdNameLookupByTypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_lookup_by_type_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdNameLookupByTypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdNameLookupByTypeResponse) ProtoMessage() {}

func (x *CmdNameLookupByTypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_lookup_by_type_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdNameLookupByTypeResponse.ProtoReflect.Descriptor instead.
func (*CmdNameLookupByTypeResponse) Descriptor() ([]byte, []int) {
	return file_cmd_lookup_by_type_proto_rawDescGZIP(), []int{1}
}

func (x *CmdNameLookupByTypeResponse) GetMessageSize() uint32 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *CmdNameLookupByTypeResponse) GetSecurityCount() int32 {
	if x != nil {
		return x.SecurityCount
	}
	return 0
}

func (x *CmdNameLookupByTypeResponse) GetSecurityNames() []string {
	if x != nil {
		return x.SecurityNames
	}
	return nil
}

type CmdNameLookupByTypeResponseStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityName string `protobuf:"bytes,1,opt,name=security_name,json=securityName,proto3" json:"security_name,omitempty"`
}

func (x *CmdNameLookupByTypeResponseStream) Reset() {
	*x = CmdNameLookupByTypeResponseStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_lookup_by_type_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdNameLookupByTypeResponseStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdNameLookupByTypeResponseStream) ProtoMessage() {}

func (x *CmdNameLookupByTypeResponseStream) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_lookup_by_type_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdNameLookupByTypeResponseStream.ProtoReflect.Descriptor instead.
func (*CmdNameLookupByTypeResponseStream) Descriptor() ([]byte, []int) {
	return file_cmd_lookup_by_type_proto_rawDescGZIP(), []int{2}
}

func (x *CmdNameLookupByTypeResponseStream) GetSecurityName() string {
	if x != nil {
		return x.SecurityName
	}
	return ""
}

var File_cmd_lookup_by_type_proto protoreflect.FileDescriptor

var file_cmd_lookup_by_type_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6d, 0x64, 0x5f, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x79, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e,
	0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x67, 0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9d,
	0x02, 0x0a, 0x13, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x67, 0x65, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x67, 0x65, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x30, 0x0a, 0x14,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x70, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x8e,
	0x01, 0x0a, 0x1b, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72,
	0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75,
	0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x22,
	0x48, 0x0a, 0x21, 0x43, 0x6d, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x4c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x42, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74,
	0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69,
	0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_lookup_by_type_proto_rawDescOnce sync.Once
	file_cmd_lookup_by_type_proto_rawDescData = file_cmd_lookup_by_type_proto_rawDesc
)

func file_cmd_lookup_by_type_proto_rawDescGZIP() []byte {
	file_cmd_lookup_by_type_proto_rawDescOnce.Do(func() {
		file_cmd_lookup_by_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_lookup_by_type_proto_rawDescData)
	})
	return file_cmd_lookup_by_type_proto_rawDescData
}

var file_cmd_lookup_by_type_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_cmd_lookup_by_type_proto_goTypes = []interface{}{
	(*CmdNameLookupByType)(nil),               // 0: org.anonymous.grpc.CmdNameLookupByType
	(*CmdNameLookupByTypeResponse)(nil),       // 1: org.anonymous.grpc.CmdNameLookupByTypeResponse
	(*CmdNameLookupByTypeResponseStream)(nil), // 2: org.anonymous.grpc.CmdNameLookupByTypeResponseStream
	(CmdType)(0), // 3: org.anonymous.grpc.CmdType
	(GetType)(0), // 4: org.anonymous.grpc.GetType
}
var file_cmd_lookup_by_type_proto_depIdxs = []int32{
	3, // 0: org.anonymous.grpc.CmdNameLookupByType.message_type:type_name -> org.anonymous.grpc.CmdType
	4, // 1: org.anonymous.grpc.CmdNameLookupByType.get_type:type_name -> org.anonymous.grpc.GetType
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_cmd_lookup_by_type_proto_init() }
func file_cmd_lookup_by_type_proto_init() {
	if File_cmd_lookup_by_type_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	file_get_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_lookup_by_type_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdNameLookupByType); i {
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
		file_cmd_lookup_by_type_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdNameLookupByTypeResponse); i {
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
		file_cmd_lookup_by_type_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdNameLookupByTypeResponseStream); i {
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
			RawDescriptor: file_cmd_lookup_by_type_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_lookup_by_type_proto_goTypes,
		DependencyIndexes: file_cmd_lookup_by_type_proto_depIdxs,
		MessageInfos:      file_cmd_lookup_by_type_proto_msgTypes,
	}.Build()
	File_cmd_lookup_by_type_proto = out.File
	file_cmd_lookup_by_type_proto_rawDesc = nil
	file_cmd_lookup_by_type_proto_goTypes = nil
	file_cmd_lookup_by_type_proto_depIdxs = nil
}
