// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

package baseline

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
type ValueType int32

const (
	ValueType_STRING ValueType = 0
	ValueType_INT    ValueType = 1
	ValueType_DOUBLE ValueType = 2
	ValueType_BOOL   ValueType = 3
)

var ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
}

var ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
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

type AnyValue struct {
	// type of the value.
	Type                 ValueType            `protobuf:"varint,1,opt,name=type,proto3,enum=baseline.ValueType" json:"type,omitempty"`
	BoolValue            bool                 `protobuf:"varint,2,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	StringValue          string               `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue             int64                `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue          float64              `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	ListValues           []*AnyValue          `protobuf:"bytes,6,rep,name=list_values,json=listValues,proto3" json:"list_values,omitempty"`
	KvlistValues         []*AttributeKeyValue `protobuf:"bytes,7,rep,name=kvlist_values,json=kvlistValues,proto3" json:"kvlist_values,omitempty"`
	BytesValue           []byte               `protobuf:"bytes,8,opt,name=bytes_value,json=bytesValue,proto3" json:"bytes_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AnyValue) Reset()         { *m = AnyValue{} }
func (m *AnyValue) String() string { return proto.CompactTextString(m) }
func (*AnyValue) ProtoMessage()    {}
func (*AnyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{1}
}

func (m *AnyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnyValue.Unmarshal(m, b)
}
func (m *AnyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnyValue.Marshal(b, m, deterministic)
}
func (m *AnyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnyValue.Merge(m, src)
}
func (m *AnyValue) XXX_Size() int {
	return xxx_messageInfo_AnyValue.Size(m)
}
func (m *AnyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AnyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AnyValue proto.InternalMessageInfo

func (m *AnyValue) GetType() ValueType {
	if m != nil {
		return m.Type
	}
	return ValueType_STRING
}

func (m *AnyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

func (m *AnyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AnyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AnyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AnyValue) GetListValues() []*AnyValue {
	if m != nil {
		return m.ListValues
	}
	return nil
}

func (m *AnyValue) GetKvlistValues() []*AttributeKeyValue {
	if m != nil {
		return m.KvlistValues
	}
	return nil
}

func (m *AnyValue) GetBytesValue() []byte {
	if m != nil {
		return m.BytesValue
	}
	return nil
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key                  string    `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                *AnyValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_555bd8c177793206, []int{2}
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

func (m *AttributeKeyValue) GetValue() *AnyValue {
	if m != nil {
		return m.Value
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
	return fileDescriptor_555bd8c177793206, []int{3}
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
	return fileDescriptor_555bd8c177793206, []int{4}
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
	proto.RegisterEnum("baseline.ValueType", ValueType_name, ValueType_value)
	proto.RegisterType((*Resource)(nil), "baseline.Resource")
	proto.RegisterType((*AnyValue)(nil), "baseline.AnyValue")
	proto.RegisterType((*AttributeKeyValue)(nil), "baseline.AttributeKeyValue")
	proto.RegisterType((*StringKeyValue)(nil), "baseline.StringKeyValue")
	proto.RegisterType((*InstrumentationLibrary)(nil), "baseline.InstrumentationLibrary")
}

func init() { proto.RegisterFile("common.proto", fileDescriptor_555bd8c177793206) }

var fileDescriptor_555bd8c177793206 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0x6a, 0xdb, 0x30,
	0x18, 0xc5, 0xa7, 0x38, 0x6d, 0xec, 0xcf, 0x69, 0xf1, 0xb4, 0x51, 0x3c, 0xba, 0x31, 0x2f, 0x37,
	0x33, 0xbb, 0x30, 0xb4, 0x85, 0x51, 0xd8, 0xcd, 0x9a, 0xee, 0x5f, 0x58, 0xa8, 0x8b, 0x9a, 0xf5,
	0x36, 0xd8, 0x8e, 0x18, 0xa2, 0x8e, 0x64, 0x24, 0x39, 0xe0, 0xcb, 0x3d, 0xe0, 0x5e, 0x64, 0x4f,
	0x31, 0x2c, 0xd9, 0x49, 0x60, 0x63, 0xbb, 0xfb, 0x7c, 0xce, 0xf9, 0x7e, 0x12, 0xc7, 0x82, 0x71,
	0x21, 0xd6, 0x6b, 0xc1, 0x93, 0x4a, 0x0a, 0x2d, 0xb0, 0x9b, 0x67, 0x8a, 0x96, 0x8c, 0xd3, 0xc9,
	0x0f, 0x04, 0x2e, 0xa1, 0x4a, 0xd4, 0xb2, 0xa0, 0xf8, 0x1d, 0x40, 0xa6, 0xb5, 0x64, 0x79, 0xad,
	0xa9, 0x0a, 0x51, 0xe4, 0xc4, 0xfe, 0xf9, 0x69, 0xd2, 0x67, 0x93, 0xab, 0xde, 0xfb, 0x4a, 0x9b,
	0xfb, 0xac, 0xac, 0x29, 0xd9, 0x8b, 0xe3, 0x4b, 0x08, 0x57, 0x52, 0x54, 0x15, 0x5d, 0x2d, 0x77,
	0xea, 0xb2, 0x10, 0x35, 0xd7, 0xe1, 0x20, 0x42, 0xf1, 0x11, 0x39, 0xe9, 0xfc, 0x2d, 0x47, 0x5d,
	0xb7, 0xee, 0xe4, 0xe7, 0x00, 0xdc, 0x2b, 0x6e, 0x91, 0xf8, 0x35, 0x0c, 0x75, 0x53, 0xd1, 0x10,
	0x45, 0x28, 0x3e, 0x3e, 0x7f, 0xb2, 0x3b, 0xdd, 0xd8, 0x8b, 0xa6, 0xa2, 0xc4, 0x04, 0xf0, 0x0b,
	0x80, 0x5c, 0x88, 0x72, 0xb9, 0x69, 0x75, 0x73, 0x82, 0x4b, 0xbc, 0x56, 0xb1, 0x9c, 0x57, 0x30,
	0x56, 0x5a, 0x32, 0xfe, 0xbd, 0x0b, 0x38, 0x11, 0x8a, 0x3d, 0xe2, 0x5b, 0xcd, 0x46, 0x4e, 0xc1,
	0x63, 0x5c, 0x77, 0xfe, 0x30, 0x42, 0xb1, 0x43, 0x5c, 0xc6, 0xf5, 0x76, 0x7f, 0x25, 0xea, 0xbc,
	0xa4, 0x9d, 0x7f, 0x10, 0xa1, 0x18, 0x11, 0xdf, 0x6a, 0x36, 0x72, 0x01, 0x7e, 0xc9, 0x54, 0x07,
	0x50, 0xe1, 0xa1, 0xe9, 0x0b, 0xef, 0xf5, 0xc5, 0xfb, 0x9a, 0xda, 0x98, 0x19, 0x15, 0x7e, 0x0f,
	0x47, 0x0f, 0x9b, 0xfd, 0xb5, 0xd1, 0xff, 0x6b, 0x1e, 0xdb, 0x8d, 0x8e, 0xf0, 0x12, 0xfc, 0xbc,
	0x69, 0xbb, 0xb5, 0x17, 0x73, 0x23, 0x14, 0x8f, 0x09, 0x18, 0xc9, 0x24, 0x26, 0x29, 0x3c, 0xfe,
	0x83, 0x81, 0x03, 0x70, 0x1e, 0x68, 0x63, 0x6a, 0xf5, 0x48, 0x3b, 0xe2, 0x18, 0x0e, 0x76, 0xdd,
	0xfd, 0xfd, 0xe2, 0x36, 0x30, 0xb9, 0x84, 0xe3, 0x3b, 0xd3, 0xdb, 0x3f, 0x68, 0x4f, 0xf7, 0x69,
	0x5e, 0xbf, 0xf9, 0x09, 0x4e, 0x66, 0x5c, 0x69, 0x59, 0xaf, 0x29, 0xd7, 0x99, 0x66, 0x82, 0xcf,
	0x59, 0x2e, 0x33, 0xd9, 0x60, 0x0c, 0x43, 0x9e, 0xad, 0x69, 0x87, 0x30, 0x33, 0x0e, 0x61, 0xb4,
	0xa1, 0x52, 0x31, 0xc1, 0x3b, 0x4a, 0xff, 0xf9, 0xe6, 0x2d, 0x78, 0xdb, 0xff, 0x8f, 0x01, 0x0e,
	0xef, 0x16, 0x64, 0x76, 0xf3, 0x39, 0x78, 0x84, 0x47, 0xe0, 0xcc, 0x6e, 0x16, 0x01, 0x6a, 0xc5,
	0x0f, 0xe9, 0xb7, 0xe9, 0xfc, 0x63, 0x30, 0xc0, 0x2e, 0x0c, 0xa7, 0x69, 0x3a, 0x0f, 0x9c, 0xe9,
	0x17, 0x78, 0xce, 0x44, 0x22, 0x2a, 0xca, 0x0b, 0xca, 0x55, 0xad, 0xec, 0xfb, 0x4f, 0xb4, 0xcc,
	0x0a, 0x9a, 0x6c, 0xce, 0xa6, 0xb0, 0x68, 0xa7, 0xdb, 0x56, 0xbc, 0x45, 0xbf, 0x06, 0xcf, 0xd2,
	0x8a, 0xf2, 0x6b, 0x9b, 0x34, 0x62, 0x62, 0xfc, 0xe4, 0xfe, 0x2c, 0x3f, 0x34, 0x9b, 0x17, 0xbf,
	0x03, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x28, 0x5e, 0xf2, 0x49, 0x03, 0x00, 0x00,
}
