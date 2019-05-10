// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v1/resources/video.proto

package resources // import "google.golang.org/genproto/googleapis/ads/googleads/v1/resources"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

// A video.
type Video struct {
	// The resource name of the video.
	// Video resource names have the form:
	//
	// `customers/{customer_id}/videos/{video_id}`
	ResourceName string `protobuf:"bytes,1,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The ID of the video.
	Id *wrappers.StringValue `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// The owner channel id of the video.
	ChannelId *wrappers.StringValue `protobuf:"bytes,3,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	// The duration of the video in milliseconds.
	DurationMillis *wrappers.Int64Value `protobuf:"bytes,4,opt,name=duration_millis,json=durationMillis,proto3" json:"duration_millis,omitempty"`
	// The title of the video.
	Title                *wrappers.StringValue `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Video) Reset()         { *m = Video{} }
func (m *Video) String() string { return proto.CompactTextString(m) }
func (*Video) ProtoMessage()    {}
func (*Video) Descriptor() ([]byte, []int) {
	return fileDescriptor_video_71e1c02985f5b3f4, []int{0}
}
func (m *Video) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Video.Unmarshal(m, b)
}
func (m *Video) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Video.Marshal(b, m, deterministic)
}
func (dst *Video) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Video.Merge(dst, src)
}
func (m *Video) XXX_Size() int {
	return xxx_messageInfo_Video.Size(m)
}
func (m *Video) XXX_DiscardUnknown() {
	xxx_messageInfo_Video.DiscardUnknown(m)
}

var xxx_messageInfo_Video proto.InternalMessageInfo

func (m *Video) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Video) GetId() *wrappers.StringValue {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Video) GetChannelId() *wrappers.StringValue {
	if m != nil {
		return m.ChannelId
	}
	return nil
}

func (m *Video) GetDurationMillis() *wrappers.Int64Value {
	if m != nil {
		return m.DurationMillis
	}
	return nil
}

func (m *Video) GetTitle() *wrappers.StringValue {
	if m != nil {
		return m.Title
	}
	return nil
}

func init() {
	proto.RegisterType((*Video)(nil), "google.ads.googleads.v1.resources.Video")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v1/resources/video.proto", fileDescriptor_video_71e1c02985f5b3f4)
}

var fileDescriptor_video_71e1c02985f5b3f4 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x49, 0xfa, 0xf5, 0x83, 0x8e, 0xff, 0x20, 0xab, 0x50, 0x8b, 0xb4, 0x4a, 0xa1, 0x0b,
	0x9d, 0x90, 0x2a, 0x2e, 0xa6, 0xab, 0x14, 0xa1, 0x54, 0x50, 0x4a, 0x85, 0x2c, 0x24, 0x50, 0xa6,
	0x9d, 0x31, 0x0e, 0x24, 0x33, 0x61, 0x26, 0xa9, 0xcf, 0xe1, 0x2b, 0xb8, 0xf4, 0x51, 0x7c, 0x14,
	0x5f, 0xc0, 0xad, 0x24, 0x93, 0xc9, 0x46, 0xd0, 0xee, 0x0e, 0x33, 0xbf, 0x73, 0xce, 0xe5, 0x5e,
	0x70, 0x11, 0x0b, 0x11, 0x27, 0xd4, 0xc3, 0x44, 0x79, 0x5a, 0x96, 0x6a, 0xeb, 0x7b, 0x92, 0x2a,
	0x51, 0xc8, 0x0d, 0x55, 0xde, 0x96, 0x11, 0x2a, 0x60, 0x26, 0x45, 0x2e, 0x9c, 0x81, 0x66, 0x20,
	0x26, 0x0a, 0x36, 0x38, 0xdc, 0xfa, 0xb0, 0xc1, 0xbb, 0x27, 0x75, 0x62, 0x65, 0x58, 0x17, 0x4f,
	0xde, 0x8b, 0xc4, 0x59, 0x46, 0xa5, 0xd2, 0x11, 0xdd, 0x9e, 0x69, 0xcc, 0x98, 0x87, 0x39, 0x17,
	0x39, 0xce, 0x99, 0xe0, 0xf5, 0xef, 0xe9, 0xab, 0x0d, 0xda, 0x61, 0x59, 0xe8, 0x9c, 0x81, 0x03,
	0x13, 0xba, 0xe2, 0x38, 0xa5, 0xae, 0xd5, 0xb7, 0x46, 0x9d, 0xe5, 0xbe, 0x79, 0xbc, 0xc7, 0x29,
	0x75, 0xce, 0x81, 0xcd, 0x88, 0x6b, 0xf7, 0xad, 0xd1, 0xde, 0xb8, 0x57, 0x4f, 0x04, 0x4d, 0x33,
	0x7c, 0xc8, 0x25, 0xe3, 0x71, 0x88, 0x93, 0x82, 0x2e, 0x6d, 0x46, 0x9c, 0x09, 0x00, 0x9b, 0x67,
	0xcc, 0x39, 0x4d, 0x56, 0x8c, 0xb8, 0xad, 0x1d, 0x5c, 0x9d, 0x9a, 0x9f, 0x13, 0xe7, 0x06, 0x1c,
	0x91, 0x42, 0x56, 0xc3, 0xae, 0x52, 0x96, 0x24, 0x4c, 0xb9, 0xff, 0xaa, 0x84, 0xe3, 0x1f, 0x09,
	0x73, 0x9e, 0x5f, 0x5f, 0xe9, 0x80, 0x43, 0xe3, 0xb9, 0xab, 0x2c, 0xce, 0x18, 0xb4, 0x73, 0x96,
	0x27, 0xd4, 0x6d, 0xef, 0xd0, 0xae, 0xd1, 0xe9, 0x97, 0x05, 0x86, 0x1b, 0x91, 0xc2, 0x3f, 0x77,
	0x3f, 0x05, 0xd5, 0xea, 0x16, 0x65, 0xd6, 0xc2, 0x7a, 0xbc, 0xad, 0x0d, 0xb1, 0x48, 0x30, 0x8f,
	0xa1, 0x90, 0xb1, 0x17, 0x53, 0x5e, 0x35, 0x99, 0x5b, 0x67, 0x4c, 0xfd, 0x72, 0xfa, 0x49, 0xa3,
	0xde, 0xec, 0xd6, 0x2c, 0x08, 0xde, 0xed, 0xc1, 0x4c, 0x47, 0x06, 0x44, 0x41, 0x2d, 0x4b, 0x15,
	0xfa, 0x70, 0x69, 0xc8, 0x0f, 0xc3, 0x44, 0x01, 0x51, 0x51, 0xc3, 0x44, 0xa1, 0x1f, 0x35, 0xcc,
	0xa7, 0x3d, 0xd4, 0x1f, 0x08, 0x05, 0x44, 0x21, 0xd4, 0x50, 0x08, 0x85, 0x3e, 0x42, 0x0d, 0xb7,
	0xfe, 0x5f, 0x0d, 0x7b, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x09, 0xfb, 0xcf, 0xd9, 0xa6, 0x02,
	0x00, 0x00,
}
