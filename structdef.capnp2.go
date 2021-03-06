// Code generated by capnpc-go. DO NOT EDIT.

package goserbench

import (
	math "math"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Capnp2A struct{ capnp.Struct }

// Capnp2A_TypeID is the unique identifier for the type Capnp2A.
const Capnp2A_TypeID = 0xa6f4401be1b430d1

func NewCapnp2A(s *capnp.Segment) (Capnp2A, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return Capnp2A{st}, err
}

func NewRootCapnp2A(s *capnp.Segment) (Capnp2A, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2})
	return Capnp2A{st}, err
}

func ReadRootCapnp2A(msg *capnp.Message) (Capnp2A, error) {
	root, err := msg.RootPtr()
	return Capnp2A{root.Struct()}, err
}

func (s Capnp2A) String() string {
	str, _ := text.Marshal(0xa6f4401be1b430d1, s.Struct)
	return str
}

func (s Capnp2A) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Capnp2A) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Capnp2A) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Capnp2A) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Capnp2A) BirthDay() int64 {
	return int64(s.Struct.Uint64(0))
}

func (s Capnp2A) SetBirthDay(v int64) {
	s.Struct.SetUint64(0, uint64(v))
}

func (s Capnp2A) Phone() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Capnp2A) HasPhone() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Capnp2A) PhoneBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Capnp2A) SetPhone(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Capnp2A) Siblings() int32 {
	return int32(s.Struct.Uint32(8))
}

func (s Capnp2A) SetSiblings(v int32) {
	s.Struct.SetUint32(8, uint32(v))
}

func (s Capnp2A) Spouse() bool {
	return s.Struct.Bit(96)
}

func (s Capnp2A) SetSpouse(v bool) {
	s.Struct.SetBit(96, v)
}

func (s Capnp2A) Money() float64 {
	return math.Float64frombits(s.Struct.Uint64(16))
}

func (s Capnp2A) SetMoney(v float64) {
	s.Struct.SetUint64(16, math.Float64bits(v))
}

// Capnp2A_List is a list of Capnp2A.
type Capnp2A_List struct{ capnp.List }

// NewCapnp2A creates a new list of Capnp2A.
func NewCapnp2A_List(s *capnp.Segment, sz int32) (Capnp2A_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 24, PointerCount: 2}, sz)
	return Capnp2A_List{l}, err
}

func (s Capnp2A_List) At(i int) Capnp2A { return Capnp2A{s.List.Struct(i)} }

func (s Capnp2A_List) Set(i int, v Capnp2A) error { return s.List.SetStruct(i, v.Struct) }

func (s Capnp2A_List) String() string {
	str, _ := text.MarshalList(0xa6f4401be1b430d1, s.List)
	return str
}

// Capnp2A_Promise is a wrapper for a Capnp2A promised by a client call.
type Capnp2A_Promise struct{ *capnp.Pipeline }

func (p Capnp2A_Promise) Struct() (Capnp2A, error) {
	s, err := p.Pipeline.Struct()
	return Capnp2A{s}, err
}

const schema_e18e32ede031d499 = "x\xdaD\xcb1Kzq\x18\xc5\xf1s\x9e\xdf\xfd\xfd" +
	"\x1d\x14\xfe=p\x87\xa0\xc1=(\xd2\xd1%\xa3\xa6\xa6" +
	"\x9e\xa95\xb5[\x0ay\xbdxu\xf0M4\xb5\x08\x09" +
	".\x85\x81C\x81AC\xd0\xd4\x124\x08\xbd\x80|\x01" +
	"m\xed7nK\xd3\x81\xc3\xe7\xbb6\xa9K\xc5\xbf\x10" +
	"\xb0\xd0\xff\xcb\x96;\x8b\xd5F\xfd\xfb\x16\xa6t\xd9\xf8" +
	"\xa3\xf2\xf9U\xbd\\\xc1K\x01\xd0\xd77]\x16\x80\xca" +
	"\xfb1\xb1\x95\xa5\x83\xfe\xb058\x8d\xe4l\xbb\xd5H" +
	"\xe2\xa4Z\xdb\xff\x9d=\xe0\x88\xb4u\x17\x00\x01\x01\x1d" +
	"o\x02v\xe5hS!\x192\xff&\x87\x80];\xda" +
	"L\xa8\xc2\x90\x02\xe8M\x15\xb0\xa9\xa3\xcd\x85\xea$\xa4" +
	"\x03\xf4.\x973G[\x0858\x09\x19\x00\xfaP\x03" +
	"l\xeehOB\xf5\x12\xd2\x03\xfa\x98\xe7\xf7\x8e\xf6," +
	"\xfc\x1f7\xba\x11K\x10\x96\xc0\xac\xd9\xe9\x0f\xda\x07\x8d" +
	"\x11\x00z\x08=XN\xda\xbd\xf8O\xa4\x9d\xe6E'" +
	">Os\x11@\x18\x80\xbbi\xd2\x1b\xa6\x11\x09!\xc1" +
	"r\xb7\x17G#\x16!,\x82?\x01\x00\x00\xff\xffq" +
	"\xfdA\xf1"

func init() {
	schemas.Register(schema_e18e32ede031d499,
		0xa6f4401be1b430d1)
}
