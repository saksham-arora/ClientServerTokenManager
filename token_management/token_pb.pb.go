// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1-devel
// 	protoc        v3.19.4
// source: token_management/token_pb.proto

package token_client_server_rpc

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

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type WriteTokenMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Low  uint64 `protobuf:"varint,3,opt,name=low,proto3" json:"low,omitempty"`
	Mid  uint64 `protobuf:"varint,4,opt,name=mid,proto3" json:"mid,omitempty"`
	High uint64 `protobuf:"varint,5,opt,name=high,proto3" json:"high,omitempty"`
}

func (x *WriteTokenMsg) Reset() {
	*x = WriteTokenMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteTokenMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteTokenMsg) ProtoMessage() {}

func (x *WriteTokenMsg) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteTokenMsg.ProtoReflect.Descriptor instead.
func (*WriteTokenMsg) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{1}
}

func (x *WriteTokenMsg) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *WriteTokenMsg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteTokenMsg) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WriteTokenMsg) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WriteTokenMsg) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateResponse string `protobuf:"bytes,1,opt,name=create_response,json=createResponse,proto3" json:"create_response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetCreateResponse() string {
	if x != nil {
		return x.CreateResponse
	}
	return ""
}

type WriteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateWriteResponse uint64 `protobuf:"varint,1,opt,name=create_write_response,json=createWriteResponse,proto3" json:"create_write_response,omitempty"`
}

func (x *WriteResponse) Reset() {
	*x = WriteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteResponse) ProtoMessage() {}

func (x *WriteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteResponse.ProtoReflect.Descriptor instead.
func (*WriteResponse) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{3}
}

func (x *WriteResponse) GetCreateWriteResponse() uint64 {
	if x != nil {
		return x.CreateWriteResponse
	}
	return 0
}

type WriteBroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashVal     uint64 `protobuf:"varint,1,opt,name=hash_val,json=hashVal,proto3" json:"hash_val,omitempty"`
	ReadingFlag bool   `protobuf:"varint,2,opt,name=reading_flag,json=readingFlag,proto3" json:"reading_flag,omitempty"`
	Ack         int32  `protobuf:"varint,3,opt,name=ack,proto3" json:"ack,omitempty"`
	Wts         string `protobuf:"bytes,4,opt,name=wts,proto3" json:"wts,omitempty"`
	Name        string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Low         uint64 `protobuf:"varint,6,opt,name=low,proto3" json:"low,omitempty"`
	Mid         uint64 `protobuf:"varint,7,opt,name=mid,proto3" json:"mid,omitempty"`
	High        uint64 `protobuf:"varint,8,opt,name=high,proto3" json:"high,omitempty"`
	TokenId     int32  `protobuf:"varint,9,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Server      string `protobuf:"bytes,10,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *WriteBroadcastRequest) Reset() {
	*x = WriteBroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBroadcastRequest) ProtoMessage() {}

func (x *WriteBroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBroadcastRequest.ProtoReflect.Descriptor instead.
func (*WriteBroadcastRequest) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{4}
}

func (x *WriteBroadcastRequest) GetHashVal() uint64 {
	if x != nil {
		return x.HashVal
	}
	return 0
}

func (x *WriteBroadcastRequest) GetReadingFlag() bool {
	if x != nil {
		return x.ReadingFlag
	}
	return false
}

func (x *WriteBroadcastRequest) GetAck() int32 {
	if x != nil {
		return x.Ack
	}
	return 0
}

func (x *WriteBroadcastRequest) GetWts() string {
	if x != nil {
		return x.Wts
	}
	return ""
}

func (x *WriteBroadcastRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WriteBroadcastRequest) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *WriteBroadcastRequest) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *WriteBroadcastRequest) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *WriteBroadcastRequest) GetTokenId() int32 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

func (x *WriteBroadcastRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

type WriteBroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ack int32 `protobuf:"varint,1,opt,name=ack,proto3" json:"ack,omitempty"`
}

func (x *WriteBroadcastResponse) Reset() {
	*x = WriteBroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteBroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteBroadcastResponse) ProtoMessage() {}

func (x *WriteBroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteBroadcastResponse.ProtoReflect.Descriptor instead.
func (*WriteBroadcastResponse) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{5}
}

func (x *WriteBroadcastResponse) GetAck() int32 {
	if x != nil {
		return x.Ack
	}
	return 0
}

type ReadBroadcastRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenId int32 `protobuf:"varint,1,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"` // string server=6;
}

func (x *ReadBroadcastRequest) Reset() {
	*x = ReadBroadcastRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBroadcastRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBroadcastRequest) ProtoMessage() {}

func (x *ReadBroadcastRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBroadcastRequest.ProtoReflect.Descriptor instead.
func (*ReadBroadcastRequest) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{6}
}

func (x *ReadBroadcastRequest) GetTokenId() int32 {
	if x != nil {
		return x.TokenId
	}
	return 0
}

type ReadBroadcastResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Wts      string `protobuf:"bytes,1,opt,name=wts,proto3" json:"wts,omitempty"`
	FinalVal uint64 `protobuf:"varint,2,opt,name=finalVal,proto3" json:"finalVal,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Low      uint64 `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	Mid      uint64 `protobuf:"varint,5,opt,name=mid,proto3" json:"mid,omitempty"`
	High     uint64 `protobuf:"varint,6,opt,name=high,proto3" json:"high,omitempty"`
}

func (x *ReadBroadcastResponse) Reset() {
	*x = ReadBroadcastResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_management_token_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadBroadcastResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadBroadcastResponse) ProtoMessage() {}

func (x *ReadBroadcastResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_management_token_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadBroadcastResponse.ProtoReflect.Descriptor instead.
func (*ReadBroadcastResponse) Descriptor() ([]byte, []int) {
	return file_token_management_token_pb_proto_rawDescGZIP(), []int{7}
}

func (x *ReadBroadcastResponse) GetWts() string {
	if x != nil {
		return x.Wts
	}
	return ""
}

func (x *ReadBroadcastResponse) GetFinalVal() uint64 {
	if x != nil {
		return x.FinalVal
	}
	return 0
}

func (x *ReadBroadcastResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ReadBroadcastResponse) GetLow() uint64 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *ReadBroadcastResponse) GetMid() uint64 {
	if x != nil {
		return x.Mid
	}
	return 0
}

func (x *ReadBroadcastResponse) GetHigh() uint64 {
	if x != nil {
		return x.High
	}
	return 0
}

var File_token_management_token_pb_proto protoreflect.FileDescriptor

var file_token_management_token_pb_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6b, 0x0a, 0x0d,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x73, 0x67, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03,
	0x6c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x22, 0x33, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x43,
	0x0a, 0x0d, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x32, 0x0a, 0x15, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x13,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0xf8, 0x01, 0x0a, 0x15, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x07, 0x68, 0x61, 0x73, 0x68, 0x56, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x61, 0x64,
	0x69, 0x6e, 0x67, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b,
	0x72, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x67, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x12, 0x10, 0x0a,
	0x03, 0x77, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x74, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2a,
	0x0a, 0x16, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x63, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x63, 0x6b, 0x22, 0x31, 0x0a, 0x14, 0x52, 0x65,
	0x61, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x91, 0x01,
	0x0a, 0x15, 0x52, 0x65, 0x61, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x77, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x77, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69, 0x6e,
	0x61, 0x6c, 0x56, 0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x68, 0x69, 0x67,
	0x68, 0x32, 0x81, 0x04, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x12, 0x47, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x65, 0x77, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1a, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x09, 0x52,
	0x65, 0x61, 0x64, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x1a, 0x1f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x4d, 0x73, 0x67, 0x1a, 0x1f, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x09, 0x44, 0x72, 0x6f, 0x70, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x1a, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0e, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x12, 0x27, 0x2e, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x42, 0x72,
	0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x62, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61,
	0x73, 0x74, 0x12, 0x26, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63,
	0x61, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x63, 0x61, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x72, 0x70, 0x63, 0x3b, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x5f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x5f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_management_token_pb_proto_rawDescOnce sync.Once
	file_token_management_token_pb_proto_rawDescData = file_token_management_token_pb_proto_rawDesc
)

func file_token_management_token_pb_proto_rawDescGZIP() []byte {
	file_token_management_token_pb_proto_rawDescOnce.Do(func() {
		file_token_management_token_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_management_token_pb_proto_rawDescData)
	})
	return file_token_management_token_pb_proto_rawDescData
}

var file_token_management_token_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_token_management_token_pb_proto_goTypes = []interface{}{
	(*Token)(nil),                  // 0: token_management.Token
	(*WriteTokenMsg)(nil),          // 1: token_management.WriteTokenMsg
	(*Response)(nil),               // 2: token_management.Response
	(*WriteResponse)(nil),          // 3: token_management.WriteResponse
	(*WriteBroadcastRequest)(nil),  // 4: token_management.WriteBroadcastRequest
	(*WriteBroadcastResponse)(nil), // 5: token_management.WriteBroadcastResponse
	(*ReadBroadcastRequest)(nil),   // 6: token_management.ReadBroadcastRequest
	(*ReadBroadcastResponse)(nil),  // 7: token_management.ReadBroadcastResponse
}
var file_token_management_token_pb_proto_depIdxs = []int32{
	0, // 0: token_management.TokenManager.CreateNewToken:input_type -> token_management.Token
	0, // 1: token_management.TokenManager.ReadToken:input_type -> token_management.Token
	1, // 2: token_management.TokenManager.WriteToken:input_type -> token_management.WriteTokenMsg
	0, // 3: token_management.TokenManager.DropToken:input_type -> token_management.Token
	4, // 4: token_management.TokenManager.WriteBroadcast:input_type -> token_management.WriteBroadcastRequest
	6, // 5: token_management.TokenManager.ReadBroadcast:input_type -> token_management.ReadBroadcastRequest
	2, // 6: token_management.TokenManager.CreateNewToken:output_type -> token_management.Response
	3, // 7: token_management.TokenManager.ReadToken:output_type -> token_management.WriteResponse
	3, // 8: token_management.TokenManager.WriteToken:output_type -> token_management.WriteResponse
	2, // 9: token_management.TokenManager.DropToken:output_type -> token_management.Response
	5, // 10: token_management.TokenManager.WriteBroadcast:output_type -> token_management.WriteBroadcastResponse
	7, // 11: token_management.TokenManager.ReadBroadcast:output_type -> token_management.ReadBroadcastResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_token_management_token_pb_proto_init() }
func file_token_management_token_pb_proto_init() {
	if File_token_management_token_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_management_token_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_token_management_token_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteTokenMsg); i {
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
		file_token_management_token_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_token_management_token_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteResponse); i {
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
		file_token_management_token_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBroadcastRequest); i {
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
		file_token_management_token_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteBroadcastResponse); i {
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
		file_token_management_token_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBroadcastRequest); i {
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
		file_token_management_token_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadBroadcastResponse); i {
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
			RawDescriptor: file_token_management_token_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_management_token_pb_proto_goTypes,
		DependencyIndexes: file_token_management_token_pb_proto_depIdxs,
		MessageInfos:      file_token_management_token_pb_proto_msgTypes,
	}.Build()
	File_token_management_token_pb_proto = out.File
	file_token_management_token_pb_proto_rawDesc = nil
	file_token_management_token_pb_proto_goTypes = nil
	file_token_management_token_pb_proto_depIdxs = nil
}