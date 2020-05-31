// Code generated by protoc-gen-go. DO NOT EDIT.
// source: error_notice.proto

package pb

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

type StcErrorNotice struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Info                 string   `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StcErrorNotice) Reset()         { *m = StcErrorNotice{} }
func (m *StcErrorNotice) String() string { return proto.CompactTextString(m) }
func (*StcErrorNotice) ProtoMessage()    {}
func (*StcErrorNotice) Descriptor() ([]byte, []int) {
	return fileDescriptor_516c84a8d13e231f, []int{0}
}

func (m *StcErrorNotice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StcErrorNotice.Unmarshal(m, b)
}
func (m *StcErrorNotice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StcErrorNotice.Marshal(b, m, deterministic)
}
func (m *StcErrorNotice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StcErrorNotice.Merge(m, src)
}
func (m *StcErrorNotice) XXX_Size() int {
	return xxx_messageInfo_StcErrorNotice.Size(m)
}
func (m *StcErrorNotice) XXX_DiscardUnknown() {
	xxx_messageInfo_StcErrorNotice.DiscardUnknown(m)
}

var xxx_messageInfo_StcErrorNotice proto.InternalMessageInfo

func (m *StcErrorNotice) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *StcErrorNotice) GetInfo() string {
	if m != nil {
		return m.Info
	}
	return ""
}

func init() {
	proto.RegisterType((*StcErrorNotice)(nil), "pb.StcErrorNotice")
}

func init() { proto.RegisterFile("error_notice.proto", fileDescriptor_516c84a8d13e231f) }

var fileDescriptor_516c84a8d13e231f = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0x2d, 0x2a, 0xca,
	0x2f, 0x8a, 0xcf, 0xcb, 0x2f, 0xc9, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x52, 0x32, 0xe1, 0xe2, 0x0b, 0x2e, 0x49, 0x76, 0x05, 0x49, 0xfa, 0x81, 0xe5, 0x84,
	0xf8, 0xb8, 0x98, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98, 0x3c, 0x53, 0x84,
	0x84, 0xb8, 0x58, 0x32, 0xf3, 0xd2, 0xf2, 0x25, 0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0xc0, 0xec,
	0x24, 0x36, 0xb0, 0x01, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0xf0, 0x3d, 0x58, 0x56,
	0x00, 0x00, 0x00,
}
