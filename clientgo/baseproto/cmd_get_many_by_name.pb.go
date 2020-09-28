// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_get_many_by_name.proto

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

type CmdGetManyByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageType   CmdType  `protobuf:"varint,1,opt,name=message_type,json=messageType,proto3,enum=org.anonymous.grpc.CmdType" json:"message_type,omitempty"`
	SecurityCount int32    `protobuf:"varint,2,opt,name=security_count,json=securityCount,proto3" json:"security_count,omitempty"`
	SecurityNames []string `protobuf:"bytes,3,rep,name=security_names,json=securityNames,proto3" json:"security_names,omitempty"`
}

func (x *CmdGetManyByName) Reset() {
	*x = CmdGetManyByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByName) ProtoMessage() {}

func (x *CmdGetManyByName) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByName.ProtoReflect.Descriptor instead.
func (*CmdGetManyByName) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{0}
}

func (x *CmdGetManyByName) GetMessageType() CmdType {
	if x != nil {
		return x.MessageType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdGetManyByName) GetSecurityCount() int32 {
	if x != nil {
		return x.SecurityCount
	}
	return 0
}

func (x *CmdGetManyByName) GetSecurityNames() []string {
	if x != nil {
		return x.SecurityNames
	}
	return nil
}

type CmdGetManyByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MessageSize     uint32                                      `protobuf:"varint,1,opt,name=message_size,json=messageSize,proto3" json:"message_size,omitempty"`
	SecurityCount   int32                                       `protobuf:"varint,2,opt,name=security_count,json=securityCount,proto3" json:"security_count,omitempty"`
	RequestResponse []*CmdGetManyByNameResponse_RequestResponse `protobuf:"bytes,3,rep,name=request_response,json=requestResponse,proto3" json:"request_response,omitempty"`
}

func (x *CmdGetManyByNameResponse) Reset() {
	*x = CmdGetManyByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponse) ProtoMessage() {}

func (x *CmdGetManyByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponse.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponse) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{1}
}

func (x *CmdGetManyByNameResponse) GetMessageSize() uint32 {
	if x != nil {
		return x.MessageSize
	}
	return 0
}

func (x *CmdGetManyByNameResponse) GetSecurityCount() int32 {
	if x != nil {
		return x.SecurityCount
	}
	return 0
}

func (x *CmdGetManyByNameResponse) GetRequestResponse() []*CmdGetManyByNameResponse_RequestResponse {
	if x != nil {
		return x.RequestResponse
	}
	return nil
}

type CmdGetManyByNameResponseStream struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageResponse:
	//	*CmdGetManyByNameResponseStream_MsgOnFailure_
	//	*CmdGetManyByNameResponseStream_MsgOnSuccess_
	MessageResponse isCmdGetManyByNameResponseStream_MessageResponse `protobuf_oneof:"MessageResponse"`
}

func (x *CmdGetManyByNameResponseStream) Reset() {
	*x = CmdGetManyByNameResponseStream{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponseStream) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponseStream) ProtoMessage() {}

func (x *CmdGetManyByNameResponseStream) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponseStream.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponseStream) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{2}
}

func (m *CmdGetManyByNameResponseStream) GetMessageResponse() isCmdGetManyByNameResponseStream_MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (x *CmdGetManyByNameResponseStream) GetMsgOnFailure() *CmdGetManyByNameResponseStream_MsgOnFailure {
	if x, ok := x.GetMessageResponse().(*CmdGetManyByNameResponseStream_MsgOnFailure_); ok {
		return x.MsgOnFailure
	}
	return nil
}

func (x *CmdGetManyByNameResponseStream) GetMsgOnSuccess() *CmdGetManyByNameResponseStream_MsgOnSuccess {
	if x, ok := x.GetMessageResponse().(*CmdGetManyByNameResponseStream_MsgOnSuccess_); ok {
		return x.MsgOnSuccess
	}
	return nil
}

type isCmdGetManyByNameResponseStream_MessageResponse interface {
	isCmdGetManyByNameResponseStream_MessageResponse()
}

type CmdGetManyByNameResponseStream_MsgOnFailure_ struct {
	MsgOnFailure *CmdGetManyByNameResponseStream_MsgOnFailure `protobuf:"bytes,1,opt,name=msg_on_failure,json=msgOnFailure,proto3,oneof"`
}

type CmdGetManyByNameResponseStream_MsgOnSuccess_ struct {
	MsgOnSuccess *CmdGetManyByNameResponseStream_MsgOnSuccess `protobuf:"bytes,2,opt,name=msg_on_success,json=msgOnSuccess,proto3,oneof"`
}

func (*CmdGetManyByNameResponseStream_MsgOnFailure_) isCmdGetManyByNameResponseStream_MessageResponse() {
}

func (*CmdGetManyByNameResponseStream_MsgOnSuccess_) isCmdGetManyByNameResponseStream_MessageResponse() {
}

type CmdGetManyByNameResponse_RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to MessageResponse:
	//	*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure_
	//	*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess_
	MessageResponse isCmdGetManyByNameResponse_RequestResponse_MessageResponse `protobuf_oneof:"MessageResponse"`
}

func (x *CmdGetManyByNameResponse_RequestResponse) Reset() {
	*x = CmdGetManyByNameResponse_RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponse_RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponse_RequestResponse) ProtoMessage() {}

func (x *CmdGetManyByNameResponse_RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponse_RequestResponse.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponse_RequestResponse) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{1, 0}
}

func (m *CmdGetManyByNameResponse_RequestResponse) GetMessageResponse() isCmdGetManyByNameResponse_RequestResponse_MessageResponse {
	if m != nil {
		return m.MessageResponse
	}
	return nil
}

func (x *CmdGetManyByNameResponse_RequestResponse) GetMsgOnFailure() *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure {
	if x, ok := x.GetMessageResponse().(*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure_); ok {
		return x.MsgOnFailure
	}
	return nil
}

func (x *CmdGetManyByNameResponse_RequestResponse) GetMsgOnSuccess() *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess {
	if x, ok := x.GetMessageResponse().(*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess_); ok {
		return x.MsgOnSuccess
	}
	return nil
}

type isCmdGetManyByNameResponse_RequestResponse_MessageResponse interface {
	isCmdGetManyByNameResponse_RequestResponse_MessageResponse()
}

type CmdGetManyByNameResponse_RequestResponse_MsgOnFailure_ struct {
	MsgOnFailure *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure `protobuf:"bytes,3,opt,name=msg_on_failure,json=msgOnFailure,proto3,oneof"`
}

type CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess_ struct {
	MsgOnSuccess *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess `protobuf:"bytes,4,opt,name=msg_on_success,json=msgOnSuccess,proto3,oneof"`
}

func (*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure_) isCmdGetManyByNameResponse_RequestResponse_MessageResponse() {
}

func (*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess_) isCmdGetManyByNameResponse_RequestResponse_MessageResponse() {
}

type CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int32  `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	Mem        []byte `protobuf:"bytes,2,opt,name=mem,proto3" json:"mem,omitempty"`
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) Reset() {
	*x = CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) ProtoMessage() {}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess) GetMem() []byte {
	if x != nil {
		return x.Mem
	}
	return nil
}

type CmdGetManyByNameResponse_RequestResponse_MsgOnFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int32 `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) Reset() {
	*x = CmdGetManyByNameResponse_RequestResponse_MsgOnFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) ProtoMessage() {}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponse_RequestResponse_MsgOnFailure.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{1, 0, 1}
}

func (x *CmdGetManyByNameResponse_RequestResponse_MsgOnFailure) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

type CmdGetManyByNameResponseStream_MsgOnSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int32  `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
	Mem        []byte `protobuf:"bytes,2,opt,name=mem,proto3" json:"mem,omitempty"`
}

func (x *CmdGetManyByNameResponseStream_MsgOnSuccess) Reset() {
	*x = CmdGetManyByNameResponseStream_MsgOnSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponseStream_MsgOnSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponseStream_MsgOnSuccess) ProtoMessage() {}

func (x *CmdGetManyByNameResponseStream_MsgOnSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponseStream_MsgOnSuccess.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponseStream_MsgOnSuccess) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{2, 0}
}

func (x *CmdGetManyByNameResponseStream_MsgOnSuccess) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

func (x *CmdGetManyByNameResponseStream_MsgOnSuccess) GetMem() []byte {
	if x != nil {
		return x.Mem
	}
	return nil
}

type CmdGetManyByNameResponseStream_MsgOnFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultCode int32 `protobuf:"varint,1,opt,name=result_code,json=resultCode,proto3" json:"result_code,omitempty"`
}

func (x *CmdGetManyByNameResponseStream_MsgOnFailure) Reset() {
	*x = CmdGetManyByNameResponseStream_MsgOnFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_get_many_by_name_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdGetManyByNameResponseStream_MsgOnFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdGetManyByNameResponseStream_MsgOnFailure) ProtoMessage() {}

func (x *CmdGetManyByNameResponseStream_MsgOnFailure) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_get_many_by_name_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdGetManyByNameResponseStream_MsgOnFailure.ProtoReflect.Descriptor instead.
func (*CmdGetManyByNameResponseStream_MsgOnFailure) Descriptor() ([]byte, []int) {
	return file_cmd_get_many_by_name_proto_rawDescGZIP(), []int{2, 1}
}

func (x *CmdGetManyByNameResponseStream_MsgOnFailure) GetResultCode() int32 {
	if x != nil {
		return x.ResultCode
	}
	return 0
}

var File_cmd_get_many_by_name_proto protoreflect.FileDescriptor

var file_cmd_get_many_by_name_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x6d, 0x64, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x6d, 0x61, 0x6e, 0x79, 0x5f, 0x62,
	0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x1a, 0x0e, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa0, 0x01, 0x0a, 0x10, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x3e, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73,
	0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x22, 0xce, 0x04, 0x0a, 0x18, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x73, 0x65, 0x63,
	0x75, 0x72, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x67, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79,
	0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74,
	0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0xfe, 0x02, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x71, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f,
	0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x49, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x73,
	0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73,
	0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x71, 0x0a, 0x0e, 0x6d, 0x73,
	0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x49, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61,
	0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52,
	0x0c, 0x6d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x41, 0x0a,
	0x0c, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x65, 0x6d,
	0x1a, 0x2f, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64,
	0x65, 0x42, 0x11, 0x0a, 0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xf9, 0x02, 0x0a, 0x1e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x67, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f,
	0x6e, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x3f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65,
	0x12, 0x67, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x3f, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61,
	0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d,
	0x64, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x4d, 0x73, 0x67,
	0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73, 0x67,
	0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x1a, 0x41, 0x0a, 0x0c, 0x4d, 0x73, 0x67,
	0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x65,
	0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6d, 0x65, 0x6d, 0x1a, 0x2f, 0x0a, 0x0c,
	0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x11, 0x0a,
	0x0f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f, 0x61, 0x75,
	0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_get_many_by_name_proto_rawDescOnce sync.Once
	file_cmd_get_many_by_name_proto_rawDescData = file_cmd_get_many_by_name_proto_rawDesc
)

func file_cmd_get_many_by_name_proto_rawDescGZIP() []byte {
	file_cmd_get_many_by_name_proto_rawDescOnce.Do(func() {
		file_cmd_get_many_by_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_get_many_by_name_proto_rawDescData)
	})
	return file_cmd_get_many_by_name_proto_rawDescData
}

var file_cmd_get_many_by_name_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_cmd_get_many_by_name_proto_goTypes = []interface{}{
	(*CmdGetManyByName)(nil),                                      // 0: org.anonymous.grpc.CmdGetManyByName
	(*CmdGetManyByNameResponse)(nil),                              // 1: org.anonymous.grpc.CmdGetManyByNameResponse
	(*CmdGetManyByNameResponseStream)(nil),                        // 2: org.anonymous.grpc.CmdGetManyByNameResponseStream
	(*CmdGetManyByNameResponse_RequestResponse)(nil),              // 3: org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse
	(*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess)(nil), // 4: org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.MsgOnSuccess
	(*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure)(nil), // 5: org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.MsgOnFailure
	(*CmdGetManyByNameResponseStream_MsgOnSuccess)(nil),           // 6: org.anonymous.grpc.CmdGetManyByNameResponseStream.MsgOnSuccess
	(*CmdGetManyByNameResponseStream_MsgOnFailure)(nil),           // 7: org.anonymous.grpc.CmdGetManyByNameResponseStream.MsgOnFailure
	(CmdType)(0), // 8: org.anonymous.grpc.CmdType
}
var file_cmd_get_many_by_name_proto_depIdxs = []int32{
	8, // 0: org.anonymous.grpc.CmdGetManyByName.message_type:type_name -> org.anonymous.grpc.CmdType
	3, // 1: org.anonymous.grpc.CmdGetManyByNameResponse.request_response:type_name -> org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse
	7, // 2: org.anonymous.grpc.CmdGetManyByNameResponseStream.msg_on_failure:type_name -> org.anonymous.grpc.CmdGetManyByNameResponseStream.MsgOnFailure
	6, // 3: org.anonymous.grpc.CmdGetManyByNameResponseStream.msg_on_success:type_name -> org.anonymous.grpc.CmdGetManyByNameResponseStream.MsgOnSuccess
	5, // 4: org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.msg_on_failure:type_name -> org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.MsgOnFailure
	4, // 5: org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.msg_on_success:type_name -> org.anonymous.grpc.CmdGetManyByNameResponse.RequestResponse.MsgOnSuccess
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cmd_get_many_by_name_proto_init() }
func file_cmd_get_many_by_name_proto_init() {
	if File_cmd_get_many_by_name_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_get_many_by_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByName); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponse); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponseStream); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponse_RequestResponse); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponseStream_MsgOnSuccess); i {
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
		file_cmd_get_many_by_name_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdGetManyByNameResponseStream_MsgOnFailure); i {
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
	file_cmd_get_many_by_name_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*CmdGetManyByNameResponseStream_MsgOnFailure_)(nil),
		(*CmdGetManyByNameResponseStream_MsgOnSuccess_)(nil),
	}
	file_cmd_get_many_by_name_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*CmdGetManyByNameResponse_RequestResponse_MsgOnFailure_)(nil),
		(*CmdGetManyByNameResponse_RequestResponse_MsgOnSuccess_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_get_many_by_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_get_many_by_name_proto_goTypes,
		DependencyIndexes: file_cmd_get_many_by_name_proto_depIdxs,
		MessageInfos:      file_cmd_get_many_by_name_proto_msgTypes,
	}.Build()
	File_cmd_get_many_by_name_proto = out.File
	file_cmd_get_many_by_name_proto_rawDesc = nil
	file_cmd_get_many_by_name_proto_goTypes = nil
	file_cmd_get_many_by_name_proto_depIdxs = nil
}
