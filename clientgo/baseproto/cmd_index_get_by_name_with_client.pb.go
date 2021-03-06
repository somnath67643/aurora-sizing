// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_index_get_by_name_with_client.proto

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

type CmdMsgIndexGetByNameWithClient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType CmdType `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=org.anonymous.grpc.CmdType" json:"message_type,omitempty"`
	TableName   string  `protobuf:"bytes,2,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	RecordName  string  `protobuf:"bytes,3,opt,name=record_name,json=recordName,proto3" json:"record_name,omitempty"`
}

func (x *CmdMsgIndexGetByNameWithClient) Reset() {
	*x = CmdMsgIndexGetByNameWithClient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdMsgIndexGetByNameWithClient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdMsgIndexGetByNameWithClient) ProtoMessage() {}

func (x *CmdMsgIndexGetByNameWithClient) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdMsgIndexGetByNameWithClient.ProtoReflect.Descriptor instead.
func (*CmdMsgIndexGetByNameWithClient) Descriptor() ([]byte, []int) {
	return file_cmd_index_get_by_name_with_client_proto_rawDescGZIP(), []int{0}
}

func (x *CmdMsgIndexGetByNameWithClient) GetMessageType() CmdType {
	if x != nil {
		return x.MessageType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdMsgIndexGetByNameWithClient) GetTableName() string {
	if x != nil {
		return x.TableName
	}
	return ""
}

func (x *CmdMsgIndexGetByNameWithClient) GetRecordName() string {
	if x != nil {
		return x.RecordName
	}
	return ""
}

type CmdMsgIndexGetByNameWithClientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasSucceeded int32 `protobuf:"varint,1,opt,name=has_succeeded,json=hasSucceeded,proto3" json:"has_succeeded,omitempty"`
	// Types that are assignable to RequestResponse:
	//	*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess_
	//	*CmdMsgIndexGetByNameWithClientResponse_ErrorCode
	RequestResponse isCmdMsgIndexGetByNameWithClientResponse_RequestResponse `protobuf_oneof:"RequestResponse"`
}

func (x *CmdMsgIndexGetByNameWithClientResponse) Reset() {
	*x = CmdMsgIndexGetByNameWithClientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdMsgIndexGetByNameWithClientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdMsgIndexGetByNameWithClientResponse) ProtoMessage() {}

func (x *CmdMsgIndexGetByNameWithClientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdMsgIndexGetByNameWithClientResponse.ProtoReflect.Descriptor instead.
func (*CmdMsgIndexGetByNameWithClientResponse) Descriptor() ([]byte, []int) {
	return file_cmd_index_get_by_name_with_client_proto_rawDescGZIP(), []int{1}
}

func (x *CmdMsgIndexGetByNameWithClientResponse) GetHasSucceeded() int32 {
	if x != nil {
		return x.HasSucceeded
	}
	return 0
}

func (m *CmdMsgIndexGetByNameWithClientResponse) GetRequestResponse() isCmdMsgIndexGetByNameWithClientResponse_RequestResponse {
	if m != nil {
		return m.RequestResponse
	}
	return nil
}

func (x *CmdMsgIndexGetByNameWithClientResponse) GetMsgOnSuccess() *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess {
	if x, ok := x.GetRequestResponse().(*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess_); ok {
		return x.MsgOnSuccess
	}
	return nil
}

func (x *CmdMsgIndexGetByNameWithClientResponse) GetErrorCode() int32 {
	if x, ok := x.GetRequestResponse().(*CmdMsgIndexGetByNameWithClientResponse_ErrorCode); ok {
		return x.ErrorCode
	}
	return 0
}

type isCmdMsgIndexGetByNameWithClientResponse_RequestResponse interface {
	isCmdMsgIndexGetByNameWithClientResponse_RequestResponse()
}

type CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess_ struct {
	MsgOnSuccess *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess `protobuf:"bytes,2,opt,name=msg_on_success,json=msgOnSuccess,proto3,oneof"`
}

type CmdMsgIndexGetByNameWithClientResponse_ErrorCode struct {
	ErrorCode int32 `protobuf:"varint,3,opt,name=error_code,json=errorCode,proto3,oneof"`
}

func (*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess_) isCmdMsgIndexGetByNameWithClientResponse_RequestResponse() {
}

func (*CmdMsgIndexGetByNameWithClientResponse_ErrorCode) isCmdMsgIndexGetByNameWithClientResponse_RequestResponse() {
}

type IndexRecord struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecurityName string             `protobuf:"bytes,1,opt,name=security_name,json=securityName,proto3" json:"security_name,omitempty"`
	DoubleVal    map[string]float64 `protobuf:"bytes,3,rep,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
	StringVal    map[string]string  `protobuf:"bytes,4,rep,name=string_val,json=stringVal,proto3" json:"string_val,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *IndexRecord) Reset() {
	*x = IndexRecord{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IndexRecord) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IndexRecord) ProtoMessage() {}

func (x *IndexRecord) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IndexRecord.ProtoReflect.Descriptor instead.
func (*IndexRecord) Descriptor() ([]byte, []int) {
	return file_cmd_index_get_by_name_with_client_proto_rawDescGZIP(), []int{2}
}

func (x *IndexRecord) GetSecurityName() string {
	if x != nil {
		return x.SecurityName
	}
	return ""
}

func (x *IndexRecord) GetDoubleVal() map[string]float64 {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *IndexRecord) GetStringVal() map[string]string {
	if x != nil {
		return x.StringVal
	}
	return nil
}

type CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IndexRecords []*IndexRecord `protobuf:"bytes,1,rep,name=index_records,json=indexRecords,proto3" json:"index_records,omitempty"`
}

func (x *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) Reset() {
	*x = CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) ProtoMessage() {}

func (x *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_index_get_by_name_with_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess.ProtoReflect.Descriptor instead.
func (*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) Descriptor() ([]byte, []int) {
	return file_cmd_index_get_by_name_with_client_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess) GetIndexRecords() []*IndexRecord {
	if x != nil {
		return x.IndexRecords
	}
	return nil
}

var File_cmd_index_get_by_name_with_client_proto protoreflect.FileDescriptor

var file_cmd_index_get_by_name_with_client_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6d, 0x64, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f,
	0x62, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x5f, 0x77, 0x69, 0x74, 0x68, 0x5f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x0e, 0x63,
	0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01,
	0x0a, 0x1e, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xc8, 0x02, 0x0a, 0x26, 0x43, 0x6d, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x68,
	0x61, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0c, 0x68, 0x61, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64,
	0x12, 0x6f, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d,
	0x64, 0x4d, 0x73, 0x67, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x57, 0x69, 0x74, 0x68, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x1f, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f,
	0x64, 0x65, 0x1a, 0x54, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x44, 0x0a, 0x0d, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x5f, 0x72, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x0c, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x42, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xcc, 0x02, 0x0a, 0x0b,
	0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x4d, 0x0a, 0x0a, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12,
	0x4d, 0x0a, 0x0a, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x09, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x1a, 0x3c,
	0x0a, 0x0e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61,
	0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73,
	0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62,
	0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_index_get_by_name_with_client_proto_rawDescOnce sync.Once
	file_cmd_index_get_by_name_with_client_proto_rawDescData = file_cmd_index_get_by_name_with_client_proto_rawDesc
)

func file_cmd_index_get_by_name_with_client_proto_rawDescGZIP() []byte {
	file_cmd_index_get_by_name_with_client_proto_rawDescOnce.Do(func() {
		file_cmd_index_get_by_name_with_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_index_get_by_name_with_client_proto_rawDescData)
	})
	return file_cmd_index_get_by_name_with_client_proto_rawDescData
}

var file_cmd_index_get_by_name_with_client_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cmd_index_get_by_name_with_client_proto_goTypes = []interface{}{
	(*CmdMsgIndexGetByNameWithClient)(nil),                      // 0: org.anonymous.grpc.CmdMsgIndexGetByNameWithClient
	(*CmdMsgIndexGetByNameWithClientResponse)(nil),              // 1: org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse
	(*IndexRecord)(nil),                                         // 2: org.anonymous.grpc.IndexRecord
	(*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess)(nil), // 3: org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse.MsgOnSuccess
	nil,          // 4: org.anonymous.grpc.IndexRecord.DoubleValEntry
	nil,          // 5: org.anonymous.grpc.IndexRecord.StringValEntry
	(CmdType)(0), // 6: org.anonymous.grpc.CmdType
}
var file_cmd_index_get_by_name_with_client_proto_depIdxs = []int32{
	6, // 0: org.anonymous.grpc.CmdMsgIndexGetByNameWithClient.message_type:type_name -> org.anonymous.grpc.CmdType
	3, // 1: org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse.msg_on_success:type_name -> org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse.MsgOnSuccess
	4, // 2: org.anonymous.grpc.IndexRecord.double_val:type_name -> org.anonymous.grpc.IndexRecord.DoubleValEntry
	5, // 3: org.anonymous.grpc.IndexRecord.string_val:type_name -> org.anonymous.grpc.IndexRecord.StringValEntry
	2, // 4: org.anonymous.grpc.CmdMsgIndexGetByNameWithClientResponse.MsgOnSuccess.index_records:type_name -> org.anonymous.grpc.IndexRecord
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_cmd_index_get_by_name_with_client_proto_init() }
func file_cmd_index_get_by_name_with_client_proto_init() {
	if File_cmd_index_get_by_name_with_client_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_index_get_by_name_with_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdMsgIndexGetByNameWithClient); i {
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
		file_cmd_index_get_by_name_with_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdMsgIndexGetByNameWithClientResponse); i {
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
		file_cmd_index_get_by_name_with_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IndexRecord); i {
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
		file_cmd_index_get_by_name_with_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess); i {
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
	file_cmd_index_get_by_name_with_client_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CmdMsgIndexGetByNameWithClientResponse_MsgOnSuccess_)(nil),
		(*CmdMsgIndexGetByNameWithClientResponse_ErrorCode)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_index_get_by_name_with_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_index_get_by_name_with_client_proto_goTypes,
		DependencyIndexes: file_cmd_index_get_by_name_with_client_proto_depIdxs,
		MessageInfos:      file_cmd_index_get_by_name_with_client_proto_msgTypes,
	}.Build()
	File_cmd_index_get_by_name_with_client_proto = out.File
	file_cmd_index_get_by_name_with_client_proto_rawDesc = nil
	file_cmd_index_get_by_name_with_client_proto_goTypes = nil
	file_cmd_index_get_by_name_with_client_proto_depIdxs = nil
}
