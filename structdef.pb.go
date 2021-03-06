// Code generated by protoc-gen-go. DO NOT EDIT.
// source: structdef.proto

package goserbench

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProtoBufA struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	BirthDay             *int64   `protobuf:"varint,2,req,name=birthDay" json:"birthDay,omitempty"`
	Phone                *string  `protobuf:"bytes,3,req,name=phone" json:"phone,omitempty"`
	Siblings             *int32   `protobuf:"varint,4,req,name=siblings" json:"siblings,omitempty"`
	Spouse               *bool    `protobuf:"varint,5,req,name=spouse" json:"spouse,omitempty"`
	Money                *float64 `protobuf:"fixed64,6,req,name=money" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProtoBufA) Reset()         { *m = ProtoBufA{} }
func (m *ProtoBufA) String() string { return proto.CompactTextString(m) }
func (*ProtoBufA) ProtoMessage()    {}
func (*ProtoBufA) Descriptor() ([]byte, []int) {
	return fileDescriptor_structdef_788fb67217951365, []int{0}
}
func (m *ProtoBufA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProtoBufA.Unmarshal(m, b)
}
func (m *ProtoBufA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProtoBufA.Marshal(b, m, deterministic)
}
func (dst *ProtoBufA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProtoBufA.Merge(dst, src)
}
func (m *ProtoBufA) XXX_Size() int {
	return xxx_messageInfo_ProtoBufA.Size(m)
}
func (m *ProtoBufA) XXX_DiscardUnknown() {
	xxx_messageInfo_ProtoBufA.DiscardUnknown(m)
}

var xxx_messageInfo_ProtoBufA proto.InternalMessageInfo

func (m *ProtoBufA) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *ProtoBufA) GetBirthDay() int64 {
	if m != nil && m.BirthDay != nil {
		return *m.BirthDay
	}
	return 0
}

func (m *ProtoBufA) GetPhone() string {
	if m != nil && m.Phone != nil {
		return *m.Phone
	}
	return ""
}

func (m *ProtoBufA) GetSiblings() int32 {
	if m != nil && m.Siblings != nil {
		return *m.Siblings
	}
	return 0
}

func (m *ProtoBufA) GetSpouse() bool {
	if m != nil && m.Spouse != nil {
		return *m.Spouse
	}
	return false
}

func (m *ProtoBufA) GetMoney() float64 {
	if m != nil && m.Money != nil {
		return *m.Money
	}
	return 0
}

func init() {
	proto.RegisterType((*ProtoBufA)(nil), "goserbench.ProtoBufA")
}

func init() { proto.RegisterFile("structdef.proto", fileDescriptor_structdef_788fb67217951365) }

var fileDescriptor_structdef_788fb67217951365 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcc, 0x4d, 0xca, 0x83, 0x30,
	0x10, 0xc6, 0x71, 0x8c, 0x1f, 0xe8, 0x6c, 0x5e, 0x08, 0x2f, 0x25, 0x74, 0x15, 0xba, 0xca, 0xaa,
	0x77, 0x68, 0xe9, 0x01, 0x4a, 0x6e, 0xa0, 0x76, 0xfc, 0x80, 0x9a, 0x91, 0x4c, 0x5c, 0x78, 0x96,
	0x5e, 0xb6, 0x44, 0x8b, 0xbb, 0xf9, 0x0d, 0xfc, 0x1f, 0xf8, 0xe3, 0xe0, 0x97, 0x36, 0xbc, 0xb0,
	0xbb, 0xce, 0x9e, 0x02, 0x49, 0xe8, 0x89, 0xd1, 0x37, 0xe8, 0xda, 0xe1, 0xf2, 0x49, 0xa0, 0x7a,
	0xc6, 0xef, 0x7d, 0xe9, 0x6e, 0x52, 0x42, 0xe6, 0xea, 0x09, 0x55, 0xa2, 0x85, 0xa9, 0xec, 0x76,
	0xcb, 0x33, 0x94, 0xcd, 0xe8, 0xc3, 0xf0, 0xa8, 0x57, 0x25, 0xb4, 0x30, 0xa9, 0x3d, 0x2c, 0xff,
	0x21, 0x9f, 0x07, 0x72, 0xa8, 0xd2, 0x2d, 0xd8, 0x11, 0x0b, 0x1e, 0x9b, 0xf7, 0xe8, 0x7a, 0x56,
	0x99, 0x16, 0x26, 0xb7, 0x87, 0xe5, 0x09, 0x0a, 0x9e, 0x69, 0x61, 0x54, 0xb9, 0x16, 0xa6, 0xb4,
	0x3f, 0xc5, 0xa5, 0x89, 0x1c, 0xae, 0xaa, 0xd0, 0xc2, 0x24, 0x76, 0xc7, 0x37, 0x00, 0x00, 0xff,
	0xff, 0x57, 0x66, 0x93, 0x38, 0xbb, 0x00, 0x00, 0x00,
}
