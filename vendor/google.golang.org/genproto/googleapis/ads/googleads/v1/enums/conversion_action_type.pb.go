// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/enums/conversion_action_type.proto

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

// Possible types of a conversion action.
type ConversionActionTypeEnum_ConversionActionType int32

const (
	// Not specified.
	ConversionActionTypeEnum_UNSPECIFIED ConversionActionTypeEnum_ConversionActionType = 0
	// Used for return value only. Represents value unknown in this version.
	ConversionActionTypeEnum_UNKNOWN ConversionActionTypeEnum_ConversionActionType = 1
	// Conversions that occur when a user clicks on an ad's call extension.
	ConversionActionTypeEnum_AD_CALL ConversionActionTypeEnum_ConversionActionType = 2
	// Conversions that occur when a user on a mobile device clicks a phone
	// number.
	ConversionActionTypeEnum_CLICK_TO_CALL ConversionActionTypeEnum_ConversionActionType = 3
	// Conversions that occur when a user downloads a mobile app from the Google
	// Play Store.
	ConversionActionTypeEnum_GOOGLE_PLAY_DOWNLOAD ConversionActionTypeEnum_ConversionActionType = 4
	// Conversions that occur when a user makes a purchase in an app through
	// Android billing.
	ConversionActionTypeEnum_GOOGLE_PLAY_IN_APP_PURCHASE ConversionActionTypeEnum_ConversionActionType = 5
	// Call conversions that are tracked by the advertiser and uploaded.
	ConversionActionTypeEnum_UPLOAD_CALLS ConversionActionTypeEnum_ConversionActionType = 6
	// Conversions that are tracked by the advertiser and uploaded with
	// attributed clicks.
	ConversionActionTypeEnum_UPLOAD_CLICKS ConversionActionTypeEnum_ConversionActionType = 7
	// Conversions that occur on a webpage.
	ConversionActionTypeEnum_WEBPAGE ConversionActionTypeEnum_ConversionActionType = 8
	// Conversions that occur when a user calls a dynamically-generated phone
	// number from an advertiser's website.
	ConversionActionTypeEnum_WEBSITE_CALL ConversionActionTypeEnum_ConversionActionType = 9
)

var ConversionActionTypeEnum_ConversionActionType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "AD_CALL",
	3: "CLICK_TO_CALL",
	4: "GOOGLE_PLAY_DOWNLOAD",
	5: "GOOGLE_PLAY_IN_APP_PURCHASE",
	6: "UPLOAD_CALLS",
	7: "UPLOAD_CLICKS",
	8: "WEBPAGE",
	9: "WEBSITE_CALL",
}
var ConversionActionTypeEnum_ConversionActionType_value = map[string]int32{
	"UNSPECIFIED":                 0,
	"UNKNOWN":                     1,
	"AD_CALL":                     2,
	"CLICK_TO_CALL":               3,
	"GOOGLE_PLAY_DOWNLOAD":        4,
	"GOOGLE_PLAY_IN_APP_PURCHASE": 5,
	"UPLOAD_CALLS":                6,
	"UPLOAD_CLICKS":               7,
	"WEBPAGE":                     8,
	"WEBSITE_CALL":                9,
}

func (x ConversionActionTypeEnum_ConversionActionType) String() string {
	return proto.EnumName(ConversionActionTypeEnum_ConversionActionType_name, int32(x))
}
func (ConversionActionTypeEnum_ConversionActionType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_type_ef6e965e011b0243, []int{0, 0}
}

// Container for enum describing possible types of a conversion action.
type ConversionActionTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConversionActionTypeEnum) Reset()         { *m = ConversionActionTypeEnum{} }
func (m *ConversionActionTypeEnum) String() string { return proto.CompactTextString(m) }
func (*ConversionActionTypeEnum) ProtoMessage()    {}
func (*ConversionActionTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_conversion_action_type_ef6e965e011b0243, []int{0}
}
func (m *ConversionActionTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConversionActionTypeEnum.Unmarshal(m, b)
}
func (m *ConversionActionTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConversionActionTypeEnum.Marshal(b, m, deterministic)
}
func (dst *ConversionActionTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConversionActionTypeEnum.Merge(dst, src)
}
func (m *ConversionActionTypeEnum) XXX_Size() int {
	return xxx_messageInfo_ConversionActionTypeEnum.Size(m)
}
func (m *ConversionActionTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_ConversionActionTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_ConversionActionTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ConversionActionTypeEnum)(nil), "google.ads.googleads.v1.enums.ConversionActionTypeEnum")
	proto.RegisterEnum("google.ads.googleads.v1.enums.ConversionActionTypeEnum_ConversionActionType", ConversionActionTypeEnum_ConversionActionType_name, ConversionActionTypeEnum_ConversionActionType_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/enums/conversion_action_type.proto", fileDescriptor_conversion_action_type_ef6e965e011b0243)
}

var fileDescriptor_conversion_action_type_ef6e965e011b0243 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xdf, 0x8e, 0x94, 0x30,
	0x14, 0xc6, 0x85, 0xd5, 0x5d, 0xed, 0x6a, 0xac, 0x64, 0x2f, 0xd6, 0x3f, 0x1b, 0xb3, 0xfb, 0x00,
	0x25, 0xc4, 0xbb, 0x7a, 0x55, 0xa0, 0x22, 0x59, 0x02, 0x8d, 0x0c, 0x43, 0x34, 0x24, 0x04, 0x07,
	0x42, 0x26, 0xd9, 0x69, 0xc9, 0x94, 0x99, 0x64, 0x5e, 0xc7, 0x4b, 0x1f, 0xc5, 0x5b, 0xdf, 0xc2,
	0x0b, 0xe3, 0x23, 0x98, 0xb6, 0x42, 0xbc, 0x18, 0xbd, 0x81, 0xd3, 0xf3, 0x9d, 0xef, 0x77, 0xda,
	0x73, 0x00, 0xee, 0x85, 0xe8, 0xef, 0x3a, 0xb7, 0x69, 0xa5, 0x6b, 0x42, 0x15, 0xed, 0x3d, 0xb7,
	0xe3, 0xbb, 0x8d, 0x74, 0x57, 0x82, 0xef, 0xbb, 0xad, 0x5c, 0x0b, 0x5e, 0x37, 0xab, 0x51, 0xfd,
	0xc6, 0xc3, 0xd0, 0xa1, 0x61, 0x2b, 0x46, 0xe1, 0x5c, 0x19, 0x03, 0x6a, 0x5a, 0x89, 0x66, 0x2f,
	0xda, 0x7b, 0x48, 0x7b, 0x5f, 0xbc, 0x9a, 0xd0, 0xc3, 0xda, 0x6d, 0x38, 0x17, 0x63, 0xa3, 0x00,
	0xd2, 0x98, 0x6f, 0x7e, 0x59, 0xe0, 0x32, 0x98, 0xe9, 0x44, 0xc3, 0x17, 0x87, 0xa1, 0xa3, 0x7c,
	0xb7, 0xb9, 0xf9, 0x6e, 0x81, 0x8b, 0x63, 0xa2, 0xf3, 0x14, 0x9c, 0x17, 0x69, 0xce, 0x68, 0x10,
	0xbf, 0x8b, 0x69, 0x08, 0xef, 0x39, 0xe7, 0xe0, 0xac, 0x48, 0x6f, 0xd3, 0xac, 0x4c, 0xa1, 0xa5,
	0x0e, 0x24, 0xac, 0x03, 0x92, 0x24, 0xd0, 0x76, 0x9e, 0x81, 0x27, 0x41, 0x12, 0x07, 0xb7, 0xf5,
	0x22, 0x33, 0xa9, 0x13, 0xe7, 0x12, 0x5c, 0x44, 0x59, 0x16, 0x25, 0xb4, 0x66, 0x09, 0xf9, 0x58,
	0x87, 0x59, 0x99, 0x26, 0x19, 0x09, 0xe1, 0x7d, 0xe7, 0x35, 0x78, 0xf9, 0xb7, 0x12, 0xa7, 0x35,
	0x61, 0xac, 0x66, 0xc5, 0x87, 0xe0, 0x3d, 0xc9, 0x29, 0x7c, 0xe0, 0x40, 0xf0, 0xb8, 0x60, 0xaa,
	0x58, 0xb3, 0x72, 0x78, 0xaa, 0xf8, 0x53, 0x46, 0xb5, 0xc9, 0xe1, 0x99, 0xea, 0x5f, 0x52, 0x9f,
	0x91, 0x88, 0xc2, 0x87, 0xca, 0x51, 0x52, 0x3f, 0x8f, 0x17, 0xd4, 0xb4, 0x7f, 0xe4, 0xff, 0xb4,
	0xc0, 0xf5, 0x4a, 0x6c, 0xd0, 0x7f, 0xc7, 0xe6, 0x3f, 0x3f, 0xf6, 0x70, 0xa6, 0x66, 0xc6, 0xac,
	0x4f, 0xfe, 0x1f, 0x6f, 0x2f, 0xee, 0x1a, 0xde, 0x23, 0xb1, 0xed, 0xdd, 0xbe, 0xe3, 0x7a, 0xa2,
	0xd3, 0xfa, 0x86, 0xb5, 0xfc, 0xc7, 0x36, 0xdf, 0xea, 0xef, 0x17, 0xfb, 0x24, 0x22, 0xe4, 0xab,
	0x7d, 0x15, 0x19, 0x14, 0x69, 0x25, 0x32, 0xa1, 0x8a, 0x96, 0x1e, 0x52, 0x1b, 0x90, 0xdf, 0x26,
	0xbd, 0x22, 0xad, 0xac, 0x66, 0xbd, 0x5a, 0x7a, 0x95, 0xd6, 0x7f, 0xd8, 0xd7, 0x26, 0x89, 0x31,
	0x69, 0x25, 0xc6, 0x73, 0x05, 0xc6, 0x4b, 0x0f, 0x63, 0x5d, 0xf3, 0xf9, 0x54, 0x5f, 0xec, 0xcd,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58, 0xf5, 0x21, 0xad, 0x65, 0x02, 0x00, 0x00,
}
