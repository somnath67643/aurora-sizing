// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: cmd_idx_get_by_name.proto

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

type CmdIdxGetByName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSize int32   `protobuf:"varint,1,opt,name=msg_size,json=msgSize,proto3" json:"msg_size,omitempty"`
	MsgType CmdType `protobuf:"varint,2,opt,name=msg_type,json=msgType,proto3,enum=org.anonymous.grpc.CmdType" json:"msg_type,omitempty"`
	IdsName string  `protobuf:"bytes,3,opt,name=ids_name,json=idsName,proto3" json:"ids_name,omitempty"`
}

func (x *CmdIdxGetByName) Reset() {
	*x = CmdIdxGetByName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_idx_get_by_name_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdIdxGetByName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdIdxGetByName) ProtoMessage() {}

func (x *CmdIdxGetByName) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_idx_get_by_name_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdIdxGetByName.ProtoReflect.Descriptor instead.
func (*CmdIdxGetByName) Descriptor() ([]byte, []int) {
	return file_cmd_idx_get_by_name_proto_rawDescGZIP(), []int{0}
}

func (x *CmdIdxGetByName) GetMsgSize() int32 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *CmdIdxGetByName) GetMsgType() CmdType {
	if x != nil {
		return x.MsgType
	}
	return CmdType_CMD_UNDEFINED
}

func (x *CmdIdxGetByName) GetIdsName() string {
	if x != nil {
		return x.IdsName
	}
	return ""
}

type CmdIdxGetByNameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MsgSize int32 `protobuf:"varint,1,opt,name=msg_size,json=msgSize,proto3" json:"msg_size,omitempty"`
	Status  int32 `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are assignable to RequestResponse:
	//	*CmdIdxGetByNameResponse_MsgOnSuccess_
	//	*CmdIdxGetByNameResponse_MsgOnFailure_
	RequestResponse isCmdIdxGetByNameResponse_RequestResponse `protobuf_oneof:"RequestResponse"`
}

func (x *CmdIdxGetByNameResponse) Reset() {
	*x = CmdIdxGetByNameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_idx_get_by_name_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdIdxGetByNameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdIdxGetByNameResponse) ProtoMessage() {}

func (x *CmdIdxGetByNameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_idx_get_by_name_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdIdxGetByNameResponse.ProtoReflect.Descriptor instead.
func (*CmdIdxGetByNameResponse) Descriptor() ([]byte, []int) {
	return file_cmd_idx_get_by_name_proto_rawDescGZIP(), []int{1}
}

func (x *CmdIdxGetByNameResponse) GetMsgSize() int32 {
	if x != nil {
		return x.MsgSize
	}
	return 0
}

func (x *CmdIdxGetByNameResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (m *CmdIdxGetByNameResponse) GetRequestResponse() isCmdIdxGetByNameResponse_RequestResponse {
	if m != nil {
		return m.RequestResponse
	}
	return nil
}

func (x *CmdIdxGetByNameResponse) GetMsgOnSuccess() *CmdIdxGetByNameResponse_MsgOnSuccess {
	if x, ok := x.GetRequestResponse().(*CmdIdxGetByNameResponse_MsgOnSuccess_); ok {
		return x.MsgOnSuccess
	}
	return nil
}

func (x *CmdIdxGetByNameResponse) GetMsgOnFailure() *CmdIdxGetByNameResponse_MsgOnFailure {
	if x, ok := x.GetRequestResponse().(*CmdIdxGetByNameResponse_MsgOnFailure_); ok {
		return x.MsgOnFailure
	}
	return nil
}

type isCmdIdxGetByNameResponse_RequestResponse interface {
	isCmdIdxGetByNameResponse_RequestResponse()
}

type CmdIdxGetByNameResponse_MsgOnSuccess_ struct {
	MsgOnSuccess *CmdIdxGetByNameResponse_MsgOnSuccess `protobuf:"bytes,3,opt,name=msg_on_success,json=msgOnSuccess,proto3,oneof"`
}

type CmdIdxGetByNameResponse_MsgOnFailure_ struct {
	MsgOnFailure *CmdIdxGetByNameResponse_MsgOnFailure `protobuf:"bytes,4,opt,name=msg_on_failure,json=msgOnFailure,proto3,oneof"`
}

func (*CmdIdxGetByNameResponse_MsgOnSuccess_) isCmdIdxGetByNameResponse_RequestResponse() {}

func (*CmdIdxGetByNameResponse_MsgOnFailure_) isCmdIdxGetByNameResponse_RequestResponse() {}

type CmdIdxGetByNameResponse_MsgOnSuccess struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasSucceeded bool    `protobuf:"varint,1,opt,name=has_succeeded,json=hasSucceeded,proto3" json:"has_succeeded,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Time         float64 `protobuf:"fixed64,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) Reset() {
	*x = CmdIdxGetByNameResponse_MsgOnSuccess{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_idx_get_by_name_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdIdxGetByNameResponse_MsgOnSuccess) ProtoMessage() {}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_idx_get_by_name_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdIdxGetByNameResponse_MsgOnSuccess.ProtoReflect.Descriptor instead.
func (*CmdIdxGetByNameResponse_MsgOnSuccess) Descriptor() ([]byte, []int) {
	return file_cmd_idx_get_by_name_proto_rawDescGZIP(), []int{1, 0}
}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) GetHasSucceeded() bool {
	if x != nil {
		return x.HasSucceeded
	}
	return false
}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CmdIdxGetByNameResponse_MsgOnSuccess) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

type CmdIdxGetByNameResponse_MsgOnFailure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HasFailed bool      `protobuf:"varint,1,opt,name=has_failed,json=hasFailed,proto3" json:"has_failed,omitempty"`
	ErrorType ErrorType `protobuf:"varint,2,opt,name=error_type,json=errorType,proto3,enum=org.anonymous.grpc.ErrorType" json:"error_type,omitempty"`
}

func (x *CmdIdxGetByNameResponse_MsgOnFailure) Reset() {
	*x = CmdIdxGetByNameResponse_MsgOnFailure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_idx_get_by_name_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CmdIdxGetByNameResponse_MsgOnFailure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CmdIdxGetByNameResponse_MsgOnFailure) ProtoMessage() {}

func (x *CmdIdxGetByNameResponse_MsgOnFailure) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_idx_get_by_name_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CmdIdxGetByNameResponse_MsgOnFailure.ProtoReflect.Descriptor instead.
func (*CmdIdxGetByNameResponse_MsgOnFailure) Descriptor() ([]byte, []int) {
	return file_cmd_idx_get_by_name_proto_rawDescGZIP(), []int{1, 1}
}

func (x *CmdIdxGetByNameResponse_MsgOnFailure) GetHasFailed() bool {
	if x != nil {
		return x.HasFailed
	}
	return false
}

func (x *CmdIdxGetByNameResponse_MsgOnFailure) GetErrorType() ErrorType {
	if x != nil {
		return x.ErrorType
	}
	return ErrorType_ERR_INVALID_ARGUMENTS
}

var File_cmd_idx_get_by_name_proto protoreflect.FileDescriptor

var file_cmd_idx_get_by_name_proto_rawDesc = []byte{
	0x0a, 0x19, 0x63, 0x6d, 0x64, 0x5f, 0x69, 0x64, 0x78, 0x5f, 0x67, 0x65, 0x74, 0x5f, 0x62, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x6f, 0x72, 0x67,
	0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x1a,
	0x0e, 0x63, 0x6d, 0x64, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x7f, 0x0a, 0x0f, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x36, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1b, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75,
	0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6d, 0x64, 0x54, 0x79, 0x70, 0x65, 0x52, 0x07,
	0x6d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x64, 0x73, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x64, 0x73, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0xed, 0x03, 0x0a, 0x17, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x78, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x6d, 0x73, 0x67, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x60, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6f, 0x72, 0x67, 0x2e,
	0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x43,
	0x6d, 0x64, 0x49, 0x64, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x12, 0x60, 0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x6f, 0x6e, 0x5f, 0x66, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x6d, 0x64, 0x49, 0x64, 0x78, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0c, 0x6d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61,
	0x69, 0x6c, 0x75, 0x72, 0x65, 0x1a, 0x5b, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x5f, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x68, 0x61,
	0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x65, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x1a, 0x6b, 0x0a, 0x0c, 0x4d, 0x73, 0x67, 0x4f, 0x6e, 0x46, 0x61, 0x69, 0x6c, 0x75,
	0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x68, 0x61, 0x73, 0x5f, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61, 0x73, 0x46, 0x61, 0x69, 0x6c, 0x65,
	0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x6e, 0x6f, 0x6e,
	0x79, 0x6d, 0x6f, 0x75, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x42,
	0x11, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x3c, 0x50, 0x01, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6d, 0x6e, 0x61, 0x74, 0x68, 0x36, 0x37, 0x36, 0x34, 0x33, 0x2f,
	0x61, 0x75, 0x72, 0x6f, 0x72, 0x61, 0x2d, 0x73, 0x69, 0x7a, 0x69, 0x6e, 0x67, 0x2f, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x67, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmd_idx_get_by_name_proto_rawDescOnce sync.Once
	file_cmd_idx_get_by_name_proto_rawDescData = file_cmd_idx_get_by_name_proto_rawDesc
)

func file_cmd_idx_get_by_name_proto_rawDescGZIP() []byte {
	file_cmd_idx_get_by_name_proto_rawDescOnce.Do(func() {
		file_cmd_idx_get_by_name_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_idx_get_by_name_proto_rawDescData)
	})
	return file_cmd_idx_get_by_name_proto_rawDescData
}

var file_cmd_idx_get_by_name_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_cmd_idx_get_by_name_proto_goTypes = []interface{}{
	(*CmdIdxGetByName)(nil),                      // 0: org.anonymous.grpc.CmdIdxGetByName
	(*CmdIdxGetByNameResponse)(nil),              // 1: org.anonymous.grpc.CmdIdxGetByNameResponse
	(*CmdIdxGetByNameResponse_MsgOnSuccess)(nil), // 2: org.anonymous.grpc.CmdIdxGetByNameResponse.MsgOnSuccess
	(*CmdIdxGetByNameResponse_MsgOnFailure)(nil), // 3: org.anonymous.grpc.CmdIdxGetByNameResponse.MsgOnFailure
	(CmdType)(0),   // 4: org.anonymous.grpc.CmdType
	(ErrorType)(0), // 5: org.anonymous.grpc.ErrorType
}
var file_cmd_idx_get_by_name_proto_depIdxs = []int32{
	4, // 0: org.anonymous.grpc.CmdIdxGetByName.msg_type:type_name -> org.anonymous.grpc.CmdType
	2, // 1: org.anonymous.grpc.CmdIdxGetByNameResponse.msg_on_success:type_name -> org.anonymous.grpc.CmdIdxGetByNameResponse.MsgOnSuccess
	3, // 2: org.anonymous.grpc.CmdIdxGetByNameResponse.msg_on_failure:type_name -> org.anonymous.grpc.CmdIdxGetByNameResponse.MsgOnFailure
	5, // 3: org.anonymous.grpc.CmdIdxGetByNameResponse.MsgOnFailure.error_type:type_name -> org.anonymous.grpc.ErrorType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cmd_idx_get_by_name_proto_init() }
func file_cmd_idx_get_by_name_proto_init() {
	if File_cmd_idx_get_by_name_proto != nil {
		return
	}
	file_cmd_type_proto_init()
	file_error_type_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_cmd_idx_get_by_name_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdIdxGetByName); i {
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
		file_cmd_idx_get_by_name_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdIdxGetByNameResponse); i {
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
		file_cmd_idx_get_by_name_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdIdxGetByNameResponse_MsgOnSuccess); i {
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
		file_cmd_idx_get_by_name_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CmdIdxGetByNameResponse_MsgOnFailure); i {
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
	file_cmd_idx_get_by_name_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*CmdIdxGetByNameResponse_MsgOnSuccess_)(nil),
		(*CmdIdxGetByNameResponse_MsgOnFailure_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cmd_idx_get_by_name_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_idx_get_by_name_proto_goTypes,
		DependencyIndexes: file_cmd_idx_get_by_name_proto_depIdxs,
		MessageInfos:      file_cmd_idx_get_by_name_proto_msgTypes,
	}.Build()
	File_cmd_idx_get_by_name_proto = out.File
	file_cmd_idx_get_by_name_proto_rawDesc = nil
	file_cmd_idx_get_by_name_proto_goTypes = nil
	file_cmd_idx_get_by_name_proto_depIdxs = nil
}
