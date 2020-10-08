// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_rename_by_name.proto

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

type CmdRenameData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSize     int32     `protobuf:"varint,1,opt,name=msg_size,json=msgSize,proto3" json:"msg_size,omitempty"`
	MsgType     CmdType   `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3,enum=org.anonymous.grpc.CmdType" json:"msg_type,omitempty"`
	OldMetadata *Metadata `protobuf:"bytes,3,opt,name=old_metadata,json=oldMetadata,proto3" json:"old_metadata,omitempty"`
	NewMetadata *Metadata `protobuf:"bytes,4,opt,name=new_metadata,json=newMetadata,proto3" json:"new_metadata,omitempty"`
}

func (x *CmdRenameData) Reset() {
	*x = CmdRenameData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_rename_by_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdRenameData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRenameData) ProtoMessage() {}

func (x *CmdRenameData) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_rename_by_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRenameData.ProtoReflect.Descriptor instead.
func (*CmdRenameData) Descriptor() ([]byte, []int) {
	return file_cmd_rename_by_name_proto_rawDescGZIP(), []int{0}
}

func (x *CmdRenameData) GetMsgSize() int32 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *CmdRenameData) GetMsgType() CmdType {
	if x != nil {
		return x.MsgType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdRenameData) GetOldMetadata() *Metadata {
	if x != nil {
		return x.OldMetadata
	}
	return nil
}

func (x *CmdRenameData) GetNewMetadata() *Metadata {
	if x != nil {
		return x.NewMetadata
	}
	return nil
}

type CmdRenameDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransMsgResponse *TransMsgResponse `protobuf:"bytes,1,opt,name=trans_msg_response,json=transMsgResponse,proto3" json:"trans_msg_response,omitempty"`
}

func (x *CmdRenameDataResponse) Reset() {
	*x = CmdRenameDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_rename_by_name_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdRenameDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdRenameDataResponse) ProtoMessage() {}

func (x *CmdRenameDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_rename_by_name_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdRenameDataResponse.ProtoReflect.Descriptor instead.
func (*CmdRenameDataResponse) Descriptor() ([]byte, []int) {
	return file_cmd_rename_by_name_proto_rawDescGZIP(), []int{1}
}

func (x *CmdRenameDataResponse) GetTransMsgResponse() *TransMsgResponse {
	if x != nil {
		return x.TransMsgResponse
	}
	return nil
}

var File_cmd_rename_by_name_proto protoreflect.FileDescriptor

var file_cmd_rename_by_name_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6d, 0x64, 0x5f, 0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x62, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e,
	0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x0d, 0x43, 0x6d, 0x64,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73,
	0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73,
	0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e,
	0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3f, 0x0a,
	0x0c, 0x6f, 0x6c, 0x64, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x0b, 0x6f, 0x6c, 0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3f,
	0x0a, 0x0c, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x6b, 0x0a, 0x15, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a, 0x12, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3c, 0x50, 0x01,
	0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d,
	0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61,
	0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f,
	0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_cmd_rename_by_name_proto_rawDescOnce sync.Once
	file_cmd_rename_by_name_proto_rawDescData = file_cmd_rename_by_name_proto_rawDesc
)

func file_cmd_rename_by_name_proto_rawDescGZIP() []byte {
	file_cmd_rename_by_name_proto_rawDescOnce.Do(func() {
		file_cmd_rename_by_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_rename_by_name_proto_rawDescData)
	})
	return file_cmd_rename_by_name_proto_rawDescData
}

var file_cmd_rename_by_name_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_cmd_rename_by_name_proto_goTypes = []interface{}{
	(*CmdRenameData)(nil),         // 0: org.anonymous.grpc.CmdRenameData
	(*CmdRenameDataResponse)(nil), // 1: org.anonymous.grpc.CmdRenameDataResponse
	(CmdType)(0),                  // 2: org.anonymous.grpc.CmdType
	(*Metadata)(nil),              // 3: org.anonymous.grpc.Metadata
	(*TransMsgResponse)(nil),      // 4: org.anonymous.grpc.TransMsgResponse
}
var file_cmd_rename_by_name_proto_depIdxs = []int32{
	2, // 0: org.anonymous.grpc.CmdRenameData.msg_type:type_name -> org.anonymous.grpc.CmdType
	3, // 1: org.anonymous.grpc.CmdRenameData.old_metadata:type_name -> org.anonymous.grpc.Metadata
	3, // 2: org.anonymous.grpc.CmdRenameData.new_metadata:type_name -> org.anonymous.grpc.Metadata
	4, // 3: org.anonymous.grpc.CmdRenameDataResponse.trans_msg_response:type_name -> org.anonymous.grpc.TransMsgResponse
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cmd_rename_by_name_proto_init() }
func file_cmd_rename_by_name_proto_init() {
	if File_cmd_rename_by_name_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	file_metadata_proto_init()
	file_trans_msg_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_rename_by_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdRenameData); i {
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
		file_cmd_rename_by_name_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdRenameDataResponse); i {
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
			RawDescriptor: file_cmd_rename_by_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_rename_by_name_proto_goTypes,
		DependencyIndexes: file_cmd_rename_by_name_proto_depIdxs,
		MessageInfos:      file_cmd_rename_by_name_proto_msgTypes,
	}.Build()
	File_cmd_rename_by_name_proto = out.File
	file_cmd_rename_by_name_proto_rawDesc = nil
	file_cmd_rename_by_name_proto_goTypes = nil
	file_cmd_rename_by_name_proto_depIdxs = nil
}
