// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package experimental

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

// ValueType is the enumeration of possible types that value can have.
// TODO: consolidate this with AttributeKeyValue.ValueType.
type ValueType int32

const (
	ValueType_STRING ValueType = 0
	ValueType_INT    ValueType = 1
	ValueType_DOUBLE ValueType = 2
	ValueType_BOOL   ValueType = 3
	ValueType_LIST   ValueType = 4
	ValueType_KVLIST ValueType = 5
)

var ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
	4: "LIST",
	5: "KVLIST",
}

var ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
	"LIST":   4,
	"KVLIST": 5,
}

func (x ValueType) String() string {
	return proto.EnumName(ValueType_name, int32(x))
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

// Resource information.
type Resource struct {
	// Set of labels that describe the resource.
	Attributes []*AttributeKeyValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,2,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetAttributes() []*AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type        ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=experimental.ValueType" json:"type,omitempty"`
	BoolValue   bool      `protobuf:"varint,3,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue string    `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue    int64     `protobuf:"varint,5,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue float64   `protobuf:"fixed64,6,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	//ExoticValue extended_value = 7;
	ListOrKvlistValues   []*AttributeKeyValue `protobuf:"bytes,7,rep,name=list_or_kvlist_values,json=listOrKvlistValues,proto3" json:"list_or_kvlist_values,omitempty"`
	BytesValue           []byte               `protobuf:"bytes,8,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_STRING
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AttributeKeyValue) GetListOrKvlistValues() []*AttributeKeyValue {
	if m != nil {
		return m.ListOrKvlistValues
	}
	return nil
}

func (m *AttributeKeyValue) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

// StringKeyValue is a pair of key/value strings. This is the simpler (and faster) version
// of AttributeKeyValue that only supports string values.
type StringKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringKeyValue) Reset()         { *m = StringKeyValue{} }
func (m *StringKeyValue) String() string { return proto.CompactTextString(m) }
func (*StringKeyValue) ProtoMessage()    {}
func (*StringKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
}

func (m *StringKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringKeyValue.Unmarshal(m, b)
}
func (m *StringKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringKeyValue.Marshal(b, m, deterministic)
}
func (m *StringKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringKeyValue.Merge(m, src)
}
func (m *StringKeyValue) XXX_Size() int {
	return xxx_messageInfo_StringKeyValue.Size(m)
}
func (m *StringKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringKeyValue proto.InternalMessageInfo

func (m *StringKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

// InstrumentationLibrary is a message representing the instrumentation library information
// such as the fully qualified name and version.
type InstrumentationLibrary struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentationLibrary) Reset()         { *m = InstrumentationLibrary{} }
func (m *InstrumentationLibrary) String() string { return proto.CompactTextString(m) }
func (*InstrumentationLibrary) ProtoMessage()    {}
func (*InstrumentationLibrary) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{3}
}

func (m *InstrumentationLibrary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentationLibrary.Unmarshal(m, b)
}
func (m *InstrumentationLibrary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentationLibrary.Marshal(b, m, deterministic)
}
func (m *InstrumentationLibrary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentationLibrary.Merge(m, src)
}
func (m *InstrumentationLibrary) XXX_Size() int {
	return xxx_messageInfo_InstrumentationLibrary.Size(m)
}
func (m *InstrumentationLibrary) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentationLibrary.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentationLibrary proto.InternalMessageInfo

func (m *InstrumentationLibrary) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InstrumentationLibrary) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterEnum("experimental.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*Resource)(nil), "experimental.Resource")
	proto.RegisterType((*AttributeKeyValue)(nil), "experimental.AttributeKeyValue")
	proto.RegisterType((*StringKeyValue)(nil), "experimental.StringKeyValue")
	proto.RegisterType((*InstrumentationLibrary)(nil), "experimental.InstrumentationLibrary")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 470 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xdd, 0x6e, 0xd3, 0x3c,
	0x18, 0xc7, 0x5f, 0x37, 0xfd, 0x7c, 0xda, 0x77, 0x0a, 0x16, 0x8c, 0x20, 0x40, 0x0b, 0x3d, 0x8a,
	0x40, 0x8a, 0xb4, 0x71, 0xb2, 0x33, 0x44, 0xc7, 0x57, 0xd5, 0x6a, 0x99, 0xdc, 0xd0, 0xd3, 0x28,
	0x49, 0x2d, 0x64, 0x2d, 0xb5, 0x23, 0xdb, 0xa9, 0xc8, 0x05, 0x70, 0x71, 0xdc, 0x0a, 0x57, 0x81,
	0x6c, 0x67, 0x65, 0x08, 0x09, 0x71, 0xf6, 0xf8, 0xf7, 0xff, 0xb0, 0xfb, 0x34, 0x30, 0x2b, 0xc5,
	0x7e, 0x2f, 0x78, 0x5c, 0x4b, 0xa1, 0x05, 0x9e, 0xd1, 0xaf, 0x35, 0x95, 0x6c, 0x4f, 0xb9, 0xce,
	0xab, 0xf9, 0x37, 0x04, 0x63, 0x42, 0x95, 0x68, 0x64, 0x49, 0xf1, 0x1b, 0x80, 0x5c, 0x6b, 0xc9,
	0x8a, 0x46, 0x53, 0x15, 0xa0, 0xd0, 0x8b, 0xa6, 0x17, 0x67, 0xf1, 0x7d, 0x7f, 0xfc, 0xf6, 0x4e,
	0x5f, 0xd1, 0x76, 0x9b, 0x57, 0x0d, 0x25, 0xf7, 0x22, 0xf8, 0x12, 0x82, 0x9d, 0x14, 0x75, 0x4d,
	0x77, 0xd9, 0x2f, 0x9a, 0x95, 0xa2, 0xe1, 0x3a, 0xe8, 0x85, 0x28, 0xfa, 0x9f, 0x9c, 0x76, 0xfa,
	0xb1, 0x47, 0x5d, 0x19, 0x75, 0xfe, 0xbd, 0x07, 0x0f, 0xfe, 0xe8, 0xc6, 0x3e, 0x78, 0xb7, 0xb4,
	0x0d, 0x50, 0x88, 0xa2, 0x09, 0x31, 0x23, 0x7e, 0x05, 0x7d, 0xdd, 0xd6, 0xd4, 0xb6, 0x9d, 0x5c,
	0x3c, 0xfe, 0xfd, 0x71, 0x36, 0x94, 0xb6, 0x35, 0x25, 0xd6, 0x84, 0x9f, 0x03, 0x14, 0x42, 0x54,
	0xd9, 0xc1, 0xf0, 0xc0, 0x0b, 0x51, 0x34, 0x26, 0x13, 0x43, 0x5c, 0xfb, 0x0b, 0x98, 0x29, 0x2d,
	0x19, 0xff, 0xd2, 0x19, 0xfa, 0xf6, 0x9a, 0xa9, 0x63, 0xce, 0xf2, 0x14, 0x26, 0x8c, 0xeb, 0x4e,
	0x1f, 0x84, 0x28, 0xf2, 0xc8, 0x98, 0x71, 0x7d, 0xcc, 0xef, 0x44, 0x53, 0x54, 0xb4, 0xd3, 0x87,
	0x21, 0x8a, 0x10, 0x99, 0x3a, 0xe6, 0x2c, 0x04, 0x1e, 0x55, 0x4c, 0xe9, 0x4c, 0xc8, 0xec, 0xf6,
	0x60, 0x27, 0x6b, 0x55, 0xc1, 0xe8, 0xdf, 0x96, 0x8b, 0x4d, 0x26, 0x91, 0x2b, 0x9b, 0xb5, 0x48,
	0xe1, 0x33, 0x98, 0x16, 0xad, 0xd9, 0xab, 0xbb, 0x75, 0x1c, 0xa2, 0x68, 0x46, 0xc0, 0x22, 0xeb,
	0x98, 0x5f, 0xc2, 0xc9, 0xc6, 0xfe, 0x86, 0xbf, 0xec, 0xf1, 0x21, 0x0c, 0x5c, 0xbc, 0x67, 0x99,
	0x3b, 0xcc, 0x3f, 0xc0, 0xe9, 0x92, 0x2b, 0x2d, 0x1b, 0xfb, 0x20, 0xcd, 0x04, 0x5f, 0xb3, 0x42,
	0xe6, 0xb2, 0xc5, 0x18, 0xfa, 0x3c, 0xdf, 0xd3, 0xae, 0xc2, 0xce, 0x38, 0x80, 0xd1, 0x81, 0x4a,
	0xc5, 0x04, 0xef, 0x5a, 0xee, 0x8e, 0x2f, 0xd7, 0x30, 0x39, 0xfe, 0x17, 0x18, 0x60, 0xb8, 0x49,
	0xc9, 0xf2, 0xfa, 0xa3, 0xff, 0x1f, 0x1e, 0x81, 0xb7, 0xbc, 0x4e, 0x7d, 0x64, 0xe0, 0xbb, 0xe4,
	0xf3, 0x62, 0xfd, 0xde, 0xef, 0xe1, 0x31, 0xf4, 0x17, 0x49, 0xb2, 0xf6, 0x3d, 0x33, 0xad, 0x97,
	0x9b, 0xd4, 0xef, 0x1b, 0x7d, 0xb5, 0xb5, 0xf3, 0x60, 0xf1, 0x09, 0x9e, 0x31, 0x11, 0x8b, 0x9a,
	0xf2, 0x92, 0x72, 0xd5, 0x28, 0xf7, 0x21, 0xc7, 0x5a, 0xe6, 0x25, 0x8d, 0x0f, 0xe7, 0x0b, 0x48,
	0xcd, 0x74, 0x63, 0xe0, 0x0d, 0xfa, 0xd1, 0x7b, 0x92, 0xd4, 0x94, 0x5f, 0x39, 0xa7, 0x85, 0xb1,
	0xd5, 0xe3, 0xed, 0x79, 0x31, 0xb4, 0xc9, 0xd7, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x97, 0x22,
	0x9a, 0x63, 0x12, 0x03, 0x00, 0x00,
}
