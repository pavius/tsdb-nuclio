// Code generated by capnpc-go. DO NOT EDIT.

package node_common_capnp

import (
	math "math"
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type ExtAttrValue struct{ capnp.Struct }
type ExtAttrValue_Which uint16

const (
	ExtAttrValue_Which_qword          ExtAttrValue_Which = 0
	ExtAttrValue_Which_uqword         ExtAttrValue_Which = 1
	ExtAttrValue_Which_blob           ExtAttrValue_Which = 2
	ExtAttrValue_Which_notExists      ExtAttrValue_Which = 3
	ExtAttrValue_Which_str            ExtAttrValue_Which = 4
	ExtAttrValue_Which_qwordIncrement ExtAttrValue_Which = 5
	ExtAttrValue_Which_time           ExtAttrValue_Which = 6
	ExtAttrValue_Which_dfloat         ExtAttrValue_Which = 7
	ExtAttrValue_Which_floatIncrement ExtAttrValue_Which = 8
	ExtAttrValue_Which_boolean        ExtAttrValue_Which = 9
)

func (w ExtAttrValue_Which) String() string {
	const s = "qworduqwordblobnotExistsstrqwordIncrementtimedfloatfloatIncrementboolean"
	switch w {
	case ExtAttrValue_Which_qword:
		return s[0:5]
	case ExtAttrValue_Which_uqword:
		return s[5:11]
	case ExtAttrValue_Which_blob:
		return s[11:15]
	case ExtAttrValue_Which_notExists:
		return s[15:24]
	case ExtAttrValue_Which_str:
		return s[24:27]
	case ExtAttrValue_Which_qwordIncrement:
		return s[27:41]
	case ExtAttrValue_Which_time:
		return s[41:45]
	case ExtAttrValue_Which_dfloat:
		return s[45:51]
	case ExtAttrValue_Which_floatIncrement:
		return s[51:65]
	case ExtAttrValue_Which_boolean:
		return s[65:72]

	}
	return "ExtAttrValue_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// ExtAttrValue_TypeID is the unique identifier for the type ExtAttrValue.
const ExtAttrValue_TypeID = 0x9bb0f31edcf7bd65

func NewExtAttrValue(s *capnp.Segment) (ExtAttrValue, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return ExtAttrValue{st}, err
}

func NewRootExtAttrValue(s *capnp.Segment) (ExtAttrValue, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1})
	return ExtAttrValue{st}, err
}

func ReadRootExtAttrValue(msg *capnp.Message) (ExtAttrValue, error) {
	root, err := msg.RootPtr()
	return ExtAttrValue{root.Struct()}, err
}

func (s ExtAttrValue) String() string {
	str, _ := text.Marshal(0x9bb0f31edcf7bd65, s.Struct)
	return str
}

func (s ExtAttrValue) Which() ExtAttrValue_Which {
	return ExtAttrValue_Which(s.Struct.Uint16(8))
}
func (s ExtAttrValue) Qword() int64 {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != qword")
	}
	return int64(s.Struct.Uint64(0))
}

func (s ExtAttrValue) SetQword(v int64) {
	s.Struct.SetUint16(8, 0)
	s.Struct.SetUint64(0, uint64(v))
}

func (s ExtAttrValue) Uqword() uint64 {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != uqword")
	}
	return s.Struct.Uint64(0)
}

func (s ExtAttrValue) SetUqword(v uint64) {
	s.Struct.SetUint16(8, 1)
	s.Struct.SetUint64(0, v)
}

func (s ExtAttrValue) Blob() ([]byte, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != blob")
	}
	p, err := s.Struct.Ptr(0)
	return []byte(p.Data()), err
}

func (s ExtAttrValue) HasBlob() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ExtAttrValue) SetBlob(v []byte) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetData(0, v)
}

func (s ExtAttrValue) SetNotExists() {
	s.Struct.SetUint16(8, 3)

}

func (s ExtAttrValue) Str() (string, error) {
	if s.Struct.Uint16(8) != 4 {
		panic("Which() != str")
	}
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s ExtAttrValue) HasStr() bool {
	if s.Struct.Uint16(8) != 4 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ExtAttrValue) StrBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s ExtAttrValue) SetStr(v string) error {
	s.Struct.SetUint16(8, 4)
	return s.Struct.SetText(0, v)
}

func (s ExtAttrValue) QwordIncrement() int64 {
	if s.Struct.Uint16(8) != 5 {
		panic("Which() != qwordIncrement")
	}
	return int64(s.Struct.Uint64(0))
}

func (s ExtAttrValue) SetQwordIncrement(v int64) {
	s.Struct.SetUint16(8, 5)
	s.Struct.SetUint64(0, uint64(v))
}

func (s ExtAttrValue) Time() (TimeSpec, error) {
	if s.Struct.Uint16(8) != 6 {
		panic("Which() != time")
	}
	p, err := s.Struct.Ptr(0)
	return TimeSpec{Struct: p.Struct()}, err
}

func (s ExtAttrValue) HasTime() bool {
	if s.Struct.Uint16(8) != 6 {
		return false
	}
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s ExtAttrValue) SetTime(v TimeSpec) error {
	s.Struct.SetUint16(8, 6)
	return s.Struct.SetPtr(0, v.Struct.ToPtr())
}

// NewTime sets the time field to a newly
// allocated TimeSpec struct, preferring placement in s's segment.
func (s ExtAttrValue) NewTime() (TimeSpec, error) {
	s.Struct.SetUint16(8, 6)
	ss, err := NewTimeSpec(s.Struct.Segment())
	if err != nil {
		return TimeSpec{}, err
	}
	err = s.Struct.SetPtr(0, ss.Struct.ToPtr())
	return ss, err
}

func (s ExtAttrValue) Dfloat() float64 {
	if s.Struct.Uint16(8) != 7 {
		panic("Which() != dfloat")
	}
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s ExtAttrValue) SetDfloat(v float64) {
	s.Struct.SetUint16(8, 7)
	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s ExtAttrValue) FloatIncrement() float64 {
	if s.Struct.Uint16(8) != 8 {
		panic("Which() != floatIncrement")
	}
	return math.Float64frombits(s.Struct.Uint64(0))
}

func (s ExtAttrValue) SetFloatIncrement(v float64) {
	s.Struct.SetUint16(8, 8)
	s.Struct.SetUint64(0, math.Float64bits(v))
}

func (s ExtAttrValue) Boolean() bool {
	if s.Struct.Uint16(8) != 9 {
		panic("Which() != boolean")
	}
	return s.Struct.Bit(0)
}

func (s ExtAttrValue) SetBoolean(v bool) {
	s.Struct.SetUint16(8, 9)
	s.Struct.SetBit(0, v)
}

// ExtAttrValue_List is a list of ExtAttrValue.
type ExtAttrValue_List struct{ capnp.List }

// NewExtAttrValue creates a new list of ExtAttrValue.
func NewExtAttrValue_List(s *capnp.Segment, sz int32) (ExtAttrValue_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 1}, sz)
	return ExtAttrValue_List{l}, err
}

func (s ExtAttrValue_List) At(i int) ExtAttrValue { return ExtAttrValue{s.List.Struct(i)} }

func (s ExtAttrValue_List) Set(i int, v ExtAttrValue) error { return s.List.SetStruct(i, v.Struct) }

func (s ExtAttrValue_List) String() string {
	str, _ := text.MarshalList(0x9bb0f31edcf7bd65, s.List)
	return str
}

// ExtAttrValue_Promise is a wrapper for a ExtAttrValue promised by a client call.
type ExtAttrValue_Promise struct{ *capnp.Pipeline }

func (p ExtAttrValue_Promise) Struct() (ExtAttrValue, error) {
	s, err := p.Pipeline.Struct()
	return ExtAttrValue{s}, err
}

func (p ExtAttrValue_Promise) Time() TimeSpec_Promise {
	return TimeSpec_Promise{Pipeline: p.Pipeline.GetPipeline(0)}
}

const schema_8a6e4e6e3e2db81e = "x\xdaL\xcfMk\x13]\x1c\x05\xf0s\xee\xccd2" +
	"O\x13\x9a\x87{\x85nl\xabt\x11\xc5\xda\xd4V\xc4" +
	"\"\xa6\x0aY\xb8\x11\xe3\x80\xfb\xbc\x8cP\x98\xcc\xc4\xe4" +
	"\x86\x167\xee]\xf8\x05\\\xf9\x0dt!\xb8\x11\xf1;" +
	"\xb8\xf4\x1b\xb8\x12\xc1\xfa\xfe\x97\xff\xf8\x82\xcbs\x7f\x9c" +
	"s\xef\xedX\xee\x9b\xed\xe80\x00\xfa\xed\xa8&\xd9\xcb" +
	"\xe3\xb7\xab\x1f\x9e>F\x7f\x8dFV_l^-n" +
	"\x16\x0f\xd1c\xfc\x1f\xc3\x9d'\x1c\xd2>g\x0c\xec<" +
	"\xe3%\x83\xb1\x14\xe58\xdb\x1a\x95\x93pR\x16[\xbd" +
	"#\x7f\xcd\xfb\xd9\x9dA\xbe\xc8\xce\x8f\x06\xd3b\xba\xf7" +
	"\xf7(\xce\x17\xd9-\xb2\xdf\x0e\xc2\x86HH\xc0&\xbc" +
	"\x00\xa4!\x03\xa6-\x1a6\xf9C\x1c\x15\x9a\xdc\x03\xd2" +
	"\xba\x82S0\xdf\xc5\xd1\x00\xf6\x7f\x9e\x05\xd2\x86\xc2\x8a" +
	"B\xf0M\x1c\x03\xc0\x9e\xe0m u\x0a\xeb\x0a\xe1W" +
	"q\x0c\x01{\x92\xa7\x81tEaC!\xfa\"\x8e\x11" +
	"`O\xf1>\x90\xae+\x9cS\xa8}\x16\xc7\x1a`\xcf" +
	"Twl(t\x14\xe2O\xe2\xf4\xc7v\xb3zU[" +
	"aW\xa1~,\x8eu\xc0nWS\x1d\x85+\x0a\xc9" +
	"GqL\x00{\x99\xd7\x81tWa\x9f\x86k\xf7\x0e" +
	"\xcb\xd9\x98\x11\x0c#\xb0\xbb\xf8\x15\x13\x18&\xe0\xf20" +
	"/\x87l\xc2\xb0\x09JQ\xfa\xde\xd1\xc1\xdc\x83s\xd4" +
	"\xe2\xb9\x9f\xb1\x01\xc3\x06(U\xe9F1Bw\x96M" +
	"\xb2\xc2\xff\x99[\xf6\x07\x93\x8c-y\xf3\xfa\xdd\xc5\xf7" +
	"\x8f\xa6\xaf\x00\xb2\x05v\xc7w\xf3r\xe0\xb9\x04\xc3%" +
	"P\xaa\xf4o\xfd7<\x18\x96e\x9e\x0d\x0a\x12\x86\x04" +
	"\x7f\x06\x00\x00\xff\xff\xf4\x13pj"

func init() {
	schemas.Register(schema_8a6e4e6e3e2db81e,
		0x9bb0f31edcf7bd65)
}
