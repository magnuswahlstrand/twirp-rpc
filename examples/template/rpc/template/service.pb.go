// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package template

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

type UpdateReq struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateReq) Reset()         { *m = UpdateReq{} }
func (m *UpdateReq) String() string { return proto.CompactTextString(m) }
func (*UpdateReq) ProtoMessage()    {}
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *UpdateReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateReq.Unmarshal(m, b)
}
func (m *UpdateReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateReq.Marshal(b, m, deterministic)
}
func (m *UpdateReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateReq.Merge(m, src)
}
func (m *UpdateReq) XXX_Size() int {
	return xxx_messageInfo_UpdateReq.Size(m)
}
func (m *UpdateReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateReq proto.InternalMessageInfo

func (m *UpdateReq) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type UpdateResp struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResp) Reset()         { *m = UpdateResp{} }
func (m *UpdateResp) String() string { return proto.CompactTextString(m) }
func (*UpdateResp) ProtoMessage()    {}
func (*UpdateResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *UpdateResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResp.Unmarshal(m, b)
}
func (m *UpdateResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResp.Marshal(b, m, deterministic)
}
func (m *UpdateResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResp.Merge(m, src)
}
func (m *UpdateResp) XXX_Size() int {
	return xxx_messageInfo_UpdateResp.Size(m)
}
func (m *UpdateResp) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResp.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResp proto.InternalMessageInfo

func init() {
	proto.RegisterType((*UpdateReq)(nil), "kyett.template.UpdateReq")
	proto.RegisterType((*UpdateResp)(nil), "kyett.template.UpdateResp")
}

func init() {
	proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626)
}

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcb, 0xae, 0x4c, 0x2d, 0x29,
	0xd1, 0x2b, 0x49, 0xcd, 0x2d, 0xc8, 0x49, 0x2c, 0x49, 0x55, 0x92, 0xe7, 0xe2, 0x0c, 0x2d, 0x48,
	0x49, 0x2c, 0x49, 0x0d, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x78, 0xb8, 0xb8, 0x60, 0x0a, 0x8a, 0x0b, 0x8c,
	0x82, 0xb9, 0x04, 0x42, 0xa0, 0x5a, 0x83, 0x21, 0xe6, 0x16, 0x09, 0xd9, 0x73, 0xb1, 0x41, 0x54,
	0x08, 0x49, 0xea, 0xa1, 0x9a, 0xae, 0x07, 0x37, 0x5a, 0x4a, 0x0a, 0x97, 0x54, 0x71, 0x81, 0x13,
	0x57, 0x14, 0x07, 0x4c, 0x38, 0x89, 0x0d, 0xec, 0x4c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x1a, 0x5e, 0x50, 0x71, 0xb7, 0x00, 0x00, 0x00,
}