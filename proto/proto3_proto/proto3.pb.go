// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3_proto/proto3.proto

package proto3_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import test_proto "github.com/golang/protobuf/proto/test_proto"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}
var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}
func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{0, 0}
}

type Message struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Hilarity             Message_Humour                     `protobuf:"varint,2,opt,name=hilarity,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                             `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm" json:"height_in_cm,omitempty"`
	Data                 []byte                             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                              `protobuf:"varint,7,opt,name=result_count,json=resultCount" json:"result_count,omitempty"`
	TrueScotsman         bool                               `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman" json:"true_scotsman,omitempty"`
	Score                float32                            `protobuf:"fixed32,9,opt,name=score" json:"score,omitempty"`
	Key                  []uint64                           `protobuf:"varint,5,rep,packed,name=key" json:"key,omitempty"`
	ShortKey             []int32                            `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey" json:"short_key,omitempty"`
	Nested               *Nested                            `protobuf:"bytes,6,opt,name=nested" json:"nested,omitempty"`
	RFunny               []Message_Humour                   `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                 `protobuf:"bytes,10,rep,name=terrain" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Proto2Field          *test_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*test_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Anything             *any.Any                           `protobuf:"bytes,14,opt,name=anything" json:"anything,omitempty"`
	ManyThings           []*any.Any                         `protobuf:"bytes,15,rep,name=many_things,json=manyThings" json:"many_things,omitempty"`
	Submessage           *Message                           `protobuf:"bytes,17,opt,name=submessage" json:"submessage,omitempty"`
	Children             []*Message                         `protobuf:"bytes,18,rep,name=children" json:"children,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{0}
}
func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (dst *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(dst, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *any.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*any.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{1}
}
func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (dst *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(dst, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{2}
}
func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (dst *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(dst, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{3}
}
func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (dst *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(dst, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_proto3_ef797ec3431c4aaa, []int{4}
}
func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (dst *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(dst, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterMapType((map[string]*test_proto.SubDefaults)(nil), "proto3_proto.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_proto.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_proto.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_proto.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_proto.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_proto.IntMaps")
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
}

func init() { proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptor_proto3_ef797ec3431c4aaa) }

var fileDescriptor_proto3_ef797ec3431c4aaa = []byte{
	// 736 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5d, 0x6f, 0xda, 0x48,
	0x14, 0x5d, 0x63, 0x3e, 0xcc, 0xb5, 0x49, 0xbc, 0xb3, 0x44, 0x3b, 0x61, 0x77, 0x25, 0x2f, 0x2b,
	0xad, 0xac, 0xd5, 0xc6, 0xd9, 0x25, 0x4a, 0x15, 0x45, 0x55, 0xab, 0x24, 0x4d, 0x54, 0x14, 0x42,
	0xd1, 0x90, 0x34, 0xea, 0x93, 0x65, 0x60, 0x00, 0xab, 0x78, 0x4c, 0xed, 0x71, 0x25, 0x3f, 0xf7,
	0x9f, 0xf4, 0x97, 0x56, 0x9e, 0x31, 0xc4, 0x89, 0x9c, 0xf6, 0x89, 0x3b, 0x67, 0xce, 0xbd, 0xe7,
	0x70, 0x7c, 0x07, 0xf6, 0xd7, 0x51, 0xc8, 0xc3, 0x23, 0x57, 0xfc, 0x1c, 0xca, 0x83, 0x23, 0x7e,
	0x90, 0x51, 0xbc, 0xea, 0xec, 0x2f, 0xc2, 0x70, 0xb1, 0xa2, 0x92, 0x32, 0x49, 0xe6, 0x87, 0x1e,
	0x4b, 0x25, 0xb1, 0xb3, 0xc7, 0x69, 0xcc, 0xf3, 0x09, 0x59, 0x29, 0xe1, 0xee, 0x17, 0x0d, 0x1a,
	0x37, 0x34, 0x8e, 0xbd, 0x05, 0x45, 0x08, 0xaa, 0xcc, 0x0b, 0x28, 0x56, 0x2c, 0xc5, 0x6e, 0x12,
	0x51, 0xa3, 0x13, 0xd0, 0x96, 0xfe, 0xca, 0x8b, 0x7c, 0x9e, 0xe2, 0x8a, 0xa5, 0xd8, 0x3b, 0xbd,
	0xdf, 0x9d, 0xa2, 0xa4, 0x93, 0x37, 0x3b, 0x6f, 0x93, 0x20, 0x4c, 0x22, 0xb2, 0x65, 0x23, 0x0b,
	0x8c, 0x25, 0xf5, 0x17, 0x4b, 0xee, 0xfa, 0xcc, 0x9d, 0x06, 0x58, 0xb5, 0x14, 0xbb, 0x45, 0x40,
	0x62, 0x7d, 0x76, 0x11, 0x64, 0x7a, 0x33, 0x8f, 0x7b, 0xb8, 0x6a, 0x29, 0xb6, 0x41, 0x44, 0x8d,
	0xfe, 0x04, 0x23, 0xa2, 0x71, 0xb2, 0xe2, 0xee, 0x34, 0x4c, 0x18, 0xc7, 0x0d, 0x4b, 0xb1, 0x55,
	0xa2, 0x4b, 0xec, 0x22, 0x83, 0xd0, 0x5f, 0xd0, 0xe2, 0x51, 0x42, 0xdd, 0x78, 0x1a, 0xf2, 0x38,
	0xf0, 0x18, 0xd6, 0x2c, 0xc5, 0xd6, 0x88, 0x91, 0x81, 0xe3, 0x1c, 0x43, 0x6d, 0xa8, 0xc5, 0xd3,
	0x30, 0xa2, 0xb8, 0x69, 0x29, 0x76, 0x85, 0xc8, 0x03, 0x32, 0x41, 0xfd, 0x48, 0x53, 0x5c, 0xb3,
	0x54, 0xbb, 0x4a, 0xb2, 0x12, 0xfd, 0x06, 0xcd, 0x78, 0x19, 0x46, 0xdc, 0xcd, 0xf0, 0x5f, 0x2c,
	0xd5, 0xae, 0x11, 0x4d, 0x00, 0xd7, 0x34, 0x45, 0xff, 0x42, 0x9d, 0xd1, 0x98, 0xd3, 0x19, 0xae,
	0x5b, 0x8a, 0xad, 0xf7, 0xda, 0x8f, 0xff, 0xfa, 0x50, 0xdc, 0x91, 0x9c, 0x83, 0x8e, 0xa1, 0x11,
	0xb9, 0xf3, 0x84, 0xb1, 0x14, 0x9b, 0x96, 0xfa, 0xc3, 0xa4, 0xea, 0xd1, 0x55, 0xc6, 0x45, 0x2f,
	0xa1, 0xc1, 0x69, 0x14, 0x79, 0x3e, 0xc3, 0x60, 0xa9, 0xb6, 0xde, 0xeb, 0x96, 0xb7, 0xdd, 0x4a,
	0xd2, 0x25, 0xe3, 0x51, 0x4a, 0x36, 0x2d, 0xe8, 0x14, 0xe4, 0x06, 0xf4, 0xdc, 0xb9, 0x4f, 0x57,
	0x33, 0xac, 0x0b, 0xa3, 0xbf, 0x3a, 0x0f, 0x5f, 0xdb, 0x19, 0x27, 0x93, 0x37, 0x74, 0xee, 0x25,
	0x2b, 0x1e, 0x13, 0x5d, 0x92, 0xaf, 0x32, 0x2e, 0xea, 0x6f, 0x7b, 0x3f, 0x7b, 0xab, 0x84, 0xe2,
	0x96, 0x90, 0xff, 0xbb, 0x5c, 0x7e, 0x24, 0x98, 0xef, 0x33, 0xa2, 0xb4, 0x90, 0x8f, 0x12, 0x08,
	0xfa, 0x0f, 0x34, 0x8f, 0xa5, 0x7c, 0xe9, 0xb3, 0x05, 0xde, 0xc9, 0xb3, 0x92, 0xbb, 0xe8, 0x6c,
	0x76, 0xd1, 0x39, 0x63, 0x29, 0xd9, 0xb2, 0xd0, 0x31, 0xe8, 0x81, 0xc7, 0x52, 0x57, 0x9c, 0x62,
	0xbc, 0x2b, 0xb4, 0xcb, 0x9b, 0x20, 0x23, 0xde, 0x0a, 0x1e, 0x3a, 0x06, 0x88, 0x93, 0x49, 0x20,
	0x4d, 0xe1, 0x9f, 0x85, 0xd4, 0x5e, 0xa9, 0x63, 0x52, 0x20, 0xa2, 0xff, 0x41, 0x9b, 0x2e, 0xfd,
	0xd5, 0x2c, 0xa2, 0x0c, 0x23, 0x21, 0xf5, 0x4c, 0xd3, 0x96, 0xd6, 0x19, 0x81, 0x51, 0x8c, 0x7c,
	0xb3, 0x3b, 0xf2, 0x71, 0x88, 0xdd, 0xf9, 0x07, 0x6a, 0x32, 0xb8, 0xca, 0x77, 0xb6, 0x43, 0x52,
	0x4e, 0x2b, 0x27, 0x4a, 0xe7, 0x1e, 0xcc, 0xa7, 0x29, 0x96, 0x4c, 0x3d, 0x78, 0x3c, 0xf5, 0xd9,
	0x4f, 0xf9, 0x30, 0xb8, 0xfb, 0x1a, 0xea, 0x72, 0xa9, 0x90, 0x0e, 0x8d, 0xbb, 0xe1, 0xf5, 0xf0,
	0xdd, 0xfd, 0xd0, 0xfc, 0x09, 0x69, 0x50, 0x1d, 0xdd, 0x0d, 0xc7, 0xa6, 0x82, 0x5a, 0xd0, 0x1c,
	0x0f, 0xce, 0x46, 0xe3, 0xdb, 0xfe, 0xc5, 0xb5, 0x59, 0x41, 0xbb, 0xa0, 0x9f, 0xf7, 0x07, 0x03,
	0xf7, 0xfc, 0xac, 0x3f, 0xb8, 0xfc, 0x60, 0xaa, 0xdd, 0x1e, 0xd4, 0xa5, 0xdd, 0xec, 0xdd, 0x4c,
	0xc4, 0x0a, 0x4b, 0x47, 0xf2, 0x90, 0xbd, 0xd4, 0x69, 0xc2, 0xa5, 0x25, 0x8d, 0x88, 0xba, 0xfb,
	0x55, 0x81, 0x9d, 0x3c, 0xb5, 0x7b, 0x9f, 0x2f, 0x6f, 0xbc, 0x35, 0x1a, 0x81, 0x31, 0x49, 0x39,
	0x75, 0x03, 0x6f, 0xbd, 0xce, 0x36, 0x41, 0x11, 0x49, 0x1f, 0x94, 0x26, 0x9d, 0xf7, 0x38, 0xe7,
	0x29, 0xa7, 0x37, 0x92, 0x9f, 0xef, 0xd5, 0xe4, 0x01, 0xe9, 0xbc, 0x02, 0xf3, 0x29, 0xa1, 0x18,
	0x99, 0x26, 0x23, 0x6b, 0x17, 0x23, 0x33, 0x8a, 0xc9, 0x7c, 0x82, 0x7a, 0x9f, 0xf1, 0xcc, 0xdb,
	0x21, 0xa8, 0x11, 0xe7, 0xb9, 0xa5, 0x3f, 0x1e, 0x5b, 0x92, 0x14, 0x87, 0x70, 0x2e, 0x2d, 0x64,
	0xcc, 0xce, 0x0b, 0xd0, 0x36, 0x40, 0x51, 0xb2, 0x56, 0x22, 0x59, 0x2b, 0x4a, 0x1e, 0x41, 0x43,
	0xce, 0x8b, 0x91, 0x0d, 0xd5, 0xc0, 0x5b, 0xc7, 0xb9, 0x68, 0xbb, 0x4c, 0x94, 0x08, 0xc6, 0xa4,
	0x2e, 0xaf, 0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x20, 0xb7, 0x04, 0xd9, 0xea, 0x05, 0x00, 0x00,
}
