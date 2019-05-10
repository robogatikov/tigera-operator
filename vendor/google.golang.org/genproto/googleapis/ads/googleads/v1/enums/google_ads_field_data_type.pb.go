// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/google_ads_field_data_type.proto

package enums // import "google.golang.org/genproto/googleapis/ads/googleads/v1/enums"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// These are the various types a GoogleAdsService artifact may take on.
type GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType int32

const (
	// Unspecified
	GoogleAdsFieldDataTypeEnum_UNSPECIFIED GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 0
	// Unknown
	GoogleAdsFieldDataTypeEnum_UNKNOWN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 1
	// Maps to google.protobuf.BoolValue
	//
	// Applicable operators:  =, !=
	GoogleAdsFieldDataTypeEnum_BOOLEAN GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 2
	// Maps to google.protobuf.StringValue. It can be compared using the set of
	// operators specific to dates however.
	//
	// Applicable operators:  =, <, >, <=, >=, BETWEEN, DURING, and IN
	GoogleAdsFieldDataTypeEnum_DATE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 3
	// Maps to google.protobuf.DoubleValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_DOUBLE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 4
	// Maps to an enum. It's specific definition can be found at type_url.
	//
	// Applicable operators:  =, !=, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_ENUM GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 5
	// Maps to google.protobuf.FloatValue
	//
	// Applicable operators:  =, !=, <, >, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_FLOAT GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 6
	// Maps to google.protobuf.Int32Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT32 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 7
	// Maps to google.protobuf.Int64Value
	//
	// Applicable operators:  =, !=, <, >, <=, >=, BETWEEN, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_INT64 GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 8
	// Maps to a protocol buffer message type. The data type's details can be
	// found in type_url.
	//
	// No operators work with MESSAGE fields.
	GoogleAdsFieldDataTypeEnum_MESSAGE GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 9
	// Maps to google.protobuf.StringValue. Represents the resource name
	// (unique id) of a resource or one of its foreign keys.
	//
	// No operators work with RESOURCE_NAME fields.
	GoogleAdsFieldDataTypeEnum_RESOURCE_NAME GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 10
	// Maps to google.protobuf.StringValue.
	//
	// Applicable operators:  =, !=, LIKE, NOT LIKE, IN, NOT IN
	GoogleAdsFieldDataTypeEnum_STRING GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType = 11
)

var GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	2:  "BOOLEAN",
	3:  "DATE",
	4:  "DOUBLE",
	5:  "ENUM",
	6:  "FLOAT",
	7:  "INT32",
	8:  "INT64",
	9:  "MESSAGE",
	10: "RESOURCE_NAME",
	11: "STRING",
}
var GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_value = map[string]int32{
	"UNSPECIFIED":   0,
	"UNKNOWN":       1,
	"BOOLEAN":       2,
	"DATE":          3,
	"DOUBLE":        4,
	"ENUM":          5,
	"FLOAT":         6,
	"INT32":         7,
	"INT64":         8,
	"MESSAGE":       9,
	"RESOURCE_NAME": 10,
	"STRING":        11,
}

func (x GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) String() string {
	return proto.EnumName(GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name, int32(x))
}
func (GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_data_type_3292bb2809ac3276, []int{0, 0}
}

// Container holding the various data types.
type GoogleAdsFieldDataTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GoogleAdsFieldDataTypeEnum) Reset()         { *m = GoogleAdsFieldDataTypeEnum{} }
func (m *GoogleAdsFieldDataTypeEnum) String() string { return proto.CompactTextString(m) }
func (*GoogleAdsFieldDataTypeEnum) ProtoMessage()    {}
func (*GoogleAdsFieldDataTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_google_ads_field_data_type_3292bb2809ac3276, []int{0}
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Unmarshal(m, b)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Marshal(b, m, deterministic)
}
func (dst *GoogleAdsFieldDataTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Merge(dst, src)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_Size() int {
	return xxx_messageInfo_GoogleAdsFieldDataTypeEnum.Size(m)
}
func (m *GoogleAdsFieldDataTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_GoogleAdsFieldDataTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_GoogleAdsFieldDataTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GoogleAdsFieldDataTypeEnum)(nil), "google.ads.googleads.v1.enums.GoogleAdsFieldDataTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType", GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_name, GoogleAdsFieldDataTypeEnum_GoogleAdsFieldDataType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/google_ads_field_data_type.proto", fileDescriptor_google_ads_field_data_type_3292bb2809ac3276)
}

var fileDescriptor_google_ads_field_data_type_3292bb2809ac3276 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcf, 0x8e, 0x93, 0x40,
	0x18, 0x17, 0x76, 0xb7, 0xbb, 0x3b, 0x8d, 0x71, 0xe4, 0xe0, 0x61, 0x75, 0x0f, 0xbb, 0x0f, 0x30,
	0x04, 0x6b, 0x3c, 0x8c, 0x89, 0xc9, 0x50, 0xa6, 0x84, 0xd8, 0x0e, 0x4d, 0x81, 0x9a, 0x18, 0x12,
	0x32, 0x3a, 0x48, 0x48, 0x5a, 0x86, 0x74, 0x68, 0x93, 0xbe, 0x8e, 0x47, 0x4f, 0x3e, 0x87, 0x17,
	0xdf, 0xc3, 0x93, 0x8f, 0x60, 0x06, 0x0a, 0xa7, 0xea, 0x85, 0xfc, 0xf8, 0x7e, 0x7f, 0x3e, 0xf8,
	0x7d, 0xe0, 0x7d, 0x21, 0x65, 0xb1, 0xc9, 0x6d, 0x2e, 0x94, 0xdd, 0x41, 0x8d, 0x0e, 0x8e, 0x9d,
	0x57, 0xfb, 0x6d, 0x3f, 0xca, 0xb8, 0x50, 0xd9, 0xd7, 0x32, 0xdf, 0x88, 0x4c, 0xf0, 0x86, 0x67,
	0xcd, 0xb1, 0xce, 0x51, 0xbd, 0x93, 0x8d, 0xb4, 0xee, 0x3b, 0x05, 0xe2, 0x42, 0xa1, 0xc1, 0x8f,
	0x0e, 0x0e, 0x6a, 0xfd, 0x77, 0xaf, 0xfa, 0xf8, 0xba, 0xb4, 0x79, 0x55, 0xc9, 0x86, 0x37, 0xa5,
	0xac, 0x54, 0x67, 0x7e, 0xfc, 0x65, 0x80, 0x3b, 0xbf, 0x15, 0x10, 0xa1, 0x66, 0x3a, 0xdf, 0xe3,
	0x0d, 0x8f, 0x8f, 0x75, 0x4e, 0xab, 0xfd, 0xf6, 0xf1, 0x87, 0x01, 0x5e, 0x9c, 0xa7, 0xad, 0x67,
	0x60, 0x9c, 0xb0, 0x68, 0x49, 0xa7, 0xc1, 0x2c, 0xa0, 0x1e, 0x7c, 0x62, 0x8d, 0xc1, 0x75, 0xc2,
	0x3e, 0xb0, 0xf0, 0x23, 0x83, 0x86, 0x7e, 0x71, 0xc3, 0x70, 0x4e, 0x09, 0x83, 0xa6, 0x75, 0x03,
	0x2e, 0x3d, 0x12, 0x53, 0x78, 0x61, 0x01, 0x30, 0xf2, 0xc2, 0xc4, 0x9d, 0x53, 0x78, 0xa9, 0xa7,
	0x94, 0x25, 0x0b, 0x78, 0x65, 0xdd, 0x82, 0xab, 0xd9, 0x3c, 0x24, 0x31, 0x1c, 0x69, 0x18, 0xb0,
	0x78, 0xf2, 0x1a, 0x5e, 0x9f, 0xe0, 0xdb, 0x37, 0xf0, 0x46, 0xa7, 0x2d, 0x68, 0x14, 0x11, 0x9f,
	0xc2, 0x5b, 0xeb, 0x39, 0x78, 0xba, 0xa2, 0x51, 0x98, 0xac, 0xa6, 0x34, 0x63, 0x64, 0x41, 0x21,
	0xd0, 0xb1, 0x51, 0xbc, 0x0a, 0x98, 0x0f, 0xc7, 0xee, 0x1f, 0x03, 0x3c, 0x7c, 0x91, 0x5b, 0xf4,
	0xdf, 0x56, 0xdc, 0x97, 0xe7, 0xff, 0x6a, 0xa9, 0x4b, 0x59, 0x1a, 0x9f, 0xdc, 0x93, 0xbb, 0x90,
	0x1b, 0x5e, 0x15, 0x48, 0xee, 0x0a, 0xbb, 0xc8, 0xab, 0xb6, 0xb2, 0xfe, 0x46, 0x75, 0xa9, 0xfe,
	0x71, 0xb2, 0x77, 0xed, 0xf3, 0x9b, 0x79, 0xe1, 0x13, 0xf2, 0xdd, 0xbc, 0xef, 0x36, 0x21, 0x22,
	0x14, 0x1a, 0x96, 0xa2, 0xb5, 0x83, 0x74, 0xc1, 0xea, 0x67, 0xcf, 0xa7, 0x44, 0xa8, 0x74, 0xe0,
	0xd3, 0xb5, 0x93, 0xb6, 0xfc, 0x6f, 0xf3, 0xa1, 0x1b, 0x62, 0x4c, 0x84, 0xc2, 0x78, 0x50, 0x60,
	0xbc, 0x76, 0x30, 0x6e, 0x35, 0x9f, 0x47, 0xed, 0x87, 0x4d, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff,
	0x3a, 0xb3, 0x98, 0xee, 0x4a, 0x02, 0x00, 0x00,
}
