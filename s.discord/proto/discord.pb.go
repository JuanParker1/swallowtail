// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discord.proto

package discordproto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message struct {
	SenderId             string   `protobuf:"bytes,1,opt,name=sender_id,json=senderId,proto3" json:"sender_id,omitempty"`
	Content              string   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Emoji                string   `protobuf:"bytes,3,opt,name=emoji,proto3" json:"emoji,omitempty"`
	Timestamp            string   `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetSenderId() string {
	if m != nil {
		return m.SenderId
	}
	return ""
}

func (m *Message) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Message) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func (m *Message) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

type SendMsgToChannelRequest struct {
	Msg                  *Message `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Channel              string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgToChannelRequest) Reset()         { *m = SendMsgToChannelRequest{} }
func (m *SendMsgToChannelRequest) String() string { return proto.CompactTextString(m) }
func (*SendMsgToChannelRequest) ProtoMessage()    {}
func (*SendMsgToChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{1}
}

func (m *SendMsgToChannelRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgToChannelRequest.Unmarshal(m, b)
}
func (m *SendMsgToChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgToChannelRequest.Marshal(b, m, deterministic)
}
func (m *SendMsgToChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgToChannelRequest.Merge(m, src)
}
func (m *SendMsgToChannelRequest) XXX_Size() int {
	return xxx_messageInfo_SendMsgToChannelRequest.Size(m)
}
func (m *SendMsgToChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgToChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgToChannelRequest proto.InternalMessageInfo

func (m *SendMsgToChannelRequest) GetMsg() *Message {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *SendMsgToChannelRequest) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type SendMsgToChannelResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendMsgToChannelResponse) Reset()         { *m = SendMsgToChannelResponse{} }
func (m *SendMsgToChannelResponse) String() string { return proto.CompactTextString(m) }
func (*SendMsgToChannelResponse) ProtoMessage()    {}
func (*SendMsgToChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d37f1fb6677774d8, []int{2}
}

func (m *SendMsgToChannelResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendMsgToChannelResponse.Unmarshal(m, b)
}
func (m *SendMsgToChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendMsgToChannelResponse.Marshal(b, m, deterministic)
}
func (m *SendMsgToChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendMsgToChannelResponse.Merge(m, src)
}
func (m *SendMsgToChannelResponse) XXX_Size() int {
	return xxx_messageInfo_SendMsgToChannelResponse.Size(m)
}
func (m *SendMsgToChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendMsgToChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendMsgToChannelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message)(nil), "accountproto.Message")
	proto.RegisterType((*SendMsgToChannelRequest)(nil), "accountproto.SendMsgToChannelRequest")
	proto.RegisterType((*SendMsgToChannelResponse)(nil), "accountproto.SendMsgToChannelResponse")
}

func init() { proto.RegisterFile("discord.proto", fileDescriptor_d37f1fb6677774d8) }

var fileDescriptor_d37f1fb6677774d8 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x90, 0x3b, 0x4b, 0xc4, 0x40,
	0x10, 0xc7, 0x3d, 0x4f, 0x8d, 0x19, 0x15, 0x64, 0x50, 0x5c, 0x4e, 0x0b, 0x09, 0xf8, 0xa8, 0x52,
	0x9c, 0x1f, 0xc1, 0xca, 0xe2, 0x9a, 0x68, 0x29, 0x48, 0xdc, 0x1d, 0x62, 0xc4, 0xcc, 0xc4, 0xcc,
	0xdc, 0xf7, 0x17, 0x77, 0xe3, 0x1b, 0xb9, 0xf2, 0xff, 0xd8, 0xfd, 0xff, 0x18, 0xd8, 0x0b, 0xad,
	0x7a, 0x19, 0x42, 0xd9, 0x0f, 0x62, 0x82, 0xbb, 0xb5, 0xf7, 0xb2, 0x64, 0x8b, 0xaa, 0x30, 0xc8,
	0x16, 0xa4, 0x5a, 0x37, 0x84, 0xc7, 0x90, 0x2b, 0x71, 0xa0, 0xe1, 0xa1, 0x0d, 0x6e, 0x72, 0x3a,
	0xb9, 0xcc, 0xab, 0xed, 0x64, 0xdc, 0x04, 0x74, 0x90, 0x79, 0x61, 0x23, 0x36, 0xb7, 0x1e, 0xa3,
	0x0f, 0x89, 0x07, 0xb0, 0x49, 0x9d, 0x3c, 0xb7, 0x6e, 0x1a, 0xfd, 0x24, 0xf0, 0x04, 0x72, 0x6b,
	0x3b, 0x52, 0xab, 0xbb, 0xde, 0x6d, 0xc4, 0xe4, 0xcb, 0x28, 0xee, 0xe1, 0xe8, 0x96, 0x38, 0x2c,
	0xb4, 0xb9, 0x93, 0xeb, 0xa7, 0x9a, 0x99, 0x5e, 0x2a, 0x7a, 0x5d, 0x92, 0x1a, 0x5e, 0xc0, 0xb4,
	0xd3, 0x26, 0xee, 0xef, 0xcc, 0x0f, 0xcb, 0xef, 0xb0, 0xe5, 0x48, 0x5a, 0xbd, 0x37, 0x22, 0x51,
	0x7a, 0xfa, 0x49, 0x94, 0x64, 0x31, 0x03, 0xf7, 0xf7, 0x77, 0xed, 0x85, 0x95, 0xe6, 0x0c, 0xd9,
	0x78, 0x0e, 0xf4, 0xb0, 0xff, 0xbb, 0x86, 0x67, 0x3f, 0x07, 0xff, 0x81, 0x9c, 0x9d, 0xaf, 0xaa,
	0xa5, 0xb5, 0x62, 0xed, 0x71, 0x2b, 0x36, 0xae, 0xde, 0x02, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xab,
	0x35, 0xa3, 0x85, 0x01, 0x00, 0x00,
}
