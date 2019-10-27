// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type SignUpReq struct {
	UserName             string   `protobuf:"bytes,1,opt,name=UserName,proto3" json:"UserName,omitempty"`
	PassWord             string   `protobuf:"bytes,2,opt,name=PassWord,proto3" json:"PassWord,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpReq) Reset()         { *m = SignUpReq{} }
func (m *SignUpReq) String() string { return proto.CompactTextString(m) }
func (*SignUpReq) ProtoMessage()    {}
func (*SignUpReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *SignUpReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpReq.Unmarshal(m, b)
}
func (m *SignUpReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpReq.Marshal(b, m, deterministic)
}
func (m *SignUpReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpReq.Merge(m, src)
}
func (m *SignUpReq) XXX_Size() int {
	return xxx_messageInfo_SignUpReq.Size(m)
}
func (m *SignUpReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpReq.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpReq proto.InternalMessageInfo

func (m *SignUpReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *SignUpReq) GetPassWord() string {
	if m != nil {
		return m.PassWord
	}
	return ""
}

type SignUpResp struct {
	Code                 int32    `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignUpResp) Reset()         { *m = SignUpResp{} }
func (m *SignUpResp) String() string { return proto.CompactTextString(m) }
func (*SignUpResp) ProtoMessage()    {}
func (*SignUpResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *SignUpResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignUpResp.Unmarshal(m, b)
}
func (m *SignUpResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignUpResp.Marshal(b, m, deterministic)
}
func (m *SignUpResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignUpResp.Merge(m, src)
}
func (m *SignUpResp) XXX_Size() int {
	return xxx_messageInfo_SignUpResp.Size(m)
}
func (m *SignUpResp) XXX_DiscardUnknown() {
	xxx_messageInfo_SignUpResp.DiscardUnknown(m)
}

var xxx_messageInfo_SignUpResp proto.InternalMessageInfo

func (m *SignUpResp) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *SignUpResp) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SignUpReq)(nil), "proto.SignUpReq")
	proto.RegisterType((*SignUpResp)(nil), "proto.SignUpResp")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x2d, 0x4e, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0xce, 0x5c, 0x9c, 0xc1, 0x99,
	0xe9, 0x79, 0xa1, 0x05, 0x41, 0xa9, 0x85, 0x42, 0x52, 0x5c, 0x1c, 0xa1, 0xc5, 0xa9, 0x45, 0x7e,
	0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x70, 0x3e, 0x48, 0x2e, 0x20, 0xb1,
	0xb8, 0x38, 0x3c, 0xbf, 0x28, 0x45, 0x82, 0x09, 0x22, 0x07, 0xe3, 0x2b, 0x59, 0x71, 0x71, 0xc1,
	0x0c, 0x29, 0x2e, 0x10, 0x12, 0xe2, 0x62, 0x71, 0xce, 0x4f, 0x81, 0x98, 0xc0, 0x1a, 0x04, 0x66,
	0x0b, 0x49, 0x70, 0xb1, 0xe7, 0xa6, 0x16, 0x17, 0x27, 0xa6, 0xa7, 0x42, 0x35, 0xc3, 0xb8, 0x46,
	0x76, 0x5c, 0xdc, 0x20, 0x3b, 0x82, 0x53, 0x8b, 0xca, 0x32, 0x93, 0x53, 0x85, 0xf4, 0xb9, 0xd8,
	0x20, 0x46, 0x09, 0x09, 0x40, 0x1c, 0xaa, 0x07, 0x77, 0x9e, 0x94, 0x20, 0x9a, 0x48, 0x71, 0x81,
	0x12, 0x43, 0x12, 0x1b, 0x58, 0xcc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x97, 0x6b, 0x30, 0x65,
	0xdc, 0x00, 0x00, 0x00,
}