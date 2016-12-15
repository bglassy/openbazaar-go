// Code generated by protoc-gen-go.
// source: profile.proto
// DO NOT EDIT!

/*
Package profile is a generated protocol buffer package.

It is generated from these files:
	profile.proto

It has these top-level messages:
	Profile
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Profile struct {
	Handle             string                     `protobuf:"bytes,1,opt,name=handle" json:"handle,omitempty"`
	Name               string                     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Location           string                     `protobuf:"bytes,3,opt,name=location" json:"location,omitempty"`
	About              string                     `protobuf:"bytes,4,opt,name=about" json:"about,omitempty"`
	ShortDescription   string                     `protobuf:"bytes,5,opt,name=shortDescription" json:"shortDescription,omitempty"`
	Website            string                     `protobuf:"bytes,6,opt,name=website" json:"website,omitempty"`
	Email              string                     `protobuf:"bytes,7,opt,name=email" json:"email,omitempty"`
	PhoneNumber        string                     `protobuf:"bytes,8,opt,name=phoneNumber" json:"phoneNumber,omitempty"`
	Social             []*Profile_SocialAccount   `protobuf:"bytes,9,rep,name=social" json:"social,omitempty"`
	AvgRating          uint32                     `protobuf:"varint,10,opt,name=avgRating" json:"avgRating,omitempty"`
	NumRatings         uint32                     `protobuf:"varint,11,opt,name=numRatings" json:"numRatings,omitempty"`
	Nsfw               bool                       `protobuf:"varint,12,opt,name=nsfw" json:"nsfw,omitempty"`
	Vendor             bool                       `protobuf:"varint,13,opt,name=vendor" json:"vendor,omitempty"`
	Moderator          bool                       `protobuf:"varint,14,opt,name=moderator" json:"moderator,omitempty"`
	PrimaryColor       string                     `protobuf:"bytes,15,opt,name=primaryColor" json:"primaryColor,omitempty"`
	SecondaryColor     string                     `protobuf:"bytes,16,opt,name=secondaryColor" json:"secondaryColor,omitempty"`
	TextColor          string                     `protobuf:"bytes,17,opt,name=textColor" json:"textColor,omitempty"`
	HighlightColor     string                     `protobuf:"bytes,18,opt,name=highlightColor" json:"highlightColor,omitempty"`
	HighlightTextColor string                     `protobuf:"bytes,19,opt,name=highlightTextColor" json:"highlightTextColor,omitempty"`
	AvatarHashes       *Profile_Image             `protobuf:"bytes,20,opt,name=avatarHashes" json:"avatarHashes,omitempty"`
	HeaderHashes       *Profile_Image             `protobuf:"bytes,21,opt,name=headerHashes" json:"headerHashes,omitempty"`
	FollowerCount      uint32                     `protobuf:"varint,22,opt,name=followerCount" json:"followerCount,omitempty"`
	FollowingCount     uint32                     `protobuf:"varint,23,opt,name=followingCount" json:"followingCount,omitempty"`
	ListingCount       uint32                     `protobuf:"varint,24,opt,name=listingCount" json:"listingCount,omitempty"`
	LastModified       *google_protobuf.Timestamp `protobuf:"bytes,25,opt,name=lastModified" json:"lastModified,omitempty"`
}

func (m *Profile) Reset()                    { *m = Profile{} }
func (m *Profile) String() string            { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()               {}
func (*Profile) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *Profile) GetSocial() []*Profile_SocialAccount {
	if m != nil {
		return m.Social
	}
	return nil
}

func (m *Profile) GetAvatarHashes() *Profile_Image {
	if m != nil {
		return m.AvatarHashes
	}
	return nil
}

func (m *Profile) GetHeaderHashes() *Profile_Image {
	if m != nil {
		return m.HeaderHashes
	}
	return nil
}

func (m *Profile) GetLastModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.LastModified
	}
	return nil
}

type Profile_SocialAccount struct {
	Type     string `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Proof    string `protobuf:"bytes,3,opt,name=proof" json:"proof,omitempty"`
}

func (m *Profile_SocialAccount) Reset()                    { *m = Profile_SocialAccount{} }
func (m *Profile_SocialAccount) String() string            { return proto.CompactTextString(m) }
func (*Profile_SocialAccount) ProtoMessage()               {}
func (*Profile_SocialAccount) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 0} }

type Profile_Image struct {
	Tiny     string `protobuf:"bytes,1,opt,name=tiny" json:"tiny,omitempty"`
	Small    string `protobuf:"bytes,2,opt,name=small" json:"small,omitempty"`
	Medium   string `protobuf:"bytes,3,opt,name=medium" json:"medium,omitempty"`
	Large    string `protobuf:"bytes,4,opt,name=large" json:"large,omitempty"`
	Original string `protobuf:"bytes,5,opt,name=original" json:"original,omitempty"`
}

func (m *Profile_Image) Reset()                    { *m = Profile_Image{} }
func (m *Profile_Image) String() string            { return proto.CompactTextString(m) }
func (*Profile_Image) ProtoMessage()               {}
func (*Profile_Image) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0, 1} }

func init() {
	proto.RegisterType((*Profile)(nil), "Profile")
	proto.RegisterType((*Profile_SocialAccount)(nil), "Profile.SocialAccount")
	proto.RegisterType((*Profile_Image)(nil), "Profile.Image")
}

var fileDescriptor3 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x93, 0xd1, 0x6e, 0xd3, 0x3c,
	0x14, 0xc7, 0xb5, 0x6f, 0x5b, 0xd7, 0xb9, 0xeb, 0xbe, 0x61, 0xc6, 0x30, 0x15, 0x82, 0x6a, 0x42,
	0x68, 0xe2, 0x22, 0x93, 0xca, 0x3d, 0x12, 0x1a, 0x17, 0x70, 0x01, 0x42, 0x61, 0x3c, 0x80, 0x9b,
	0x9c, 0x24, 0x96, 0xec, 0x38, 0xb2, 0x9d, 0x96, 0x8a, 0x67, 0xe4, 0x9d, 0xb0, 0x8f, 0xd3, 0xb4,
	0x19, 0xbb, 0xa8, 0xea, 0xff, 0xff, 0xfc, 0xce, 0x71, 0x7c, 0xec, 0x43, 0xa6, 0x8d, 0xd1, 0x85,
	0x90, 0x90, 0xf8, 0x7f, 0xa7, 0x67, 0xaf, 0x4b, 0xad, 0x4b, 0x09, 0xb7, 0xa8, 0x96, 0x6d, 0x71,
	0xeb, 0x84, 0x02, 0xeb, 0xb8, 0x6a, 0x22, 0x70, 0xfd, 0x67, 0x4c, 0x4e, 0xbe, 0xc7, 0x14, 0x7a,
	0x45, 0x46, 0x15, 0xaf, 0x73, 0x09, 0xec, 0x60, 0x7e, 0x70, 0x73, 0x9a, 0x76, 0x8a, 0x52, 0x72,
	0x54, 0x73, 0x05, 0xec, 0x3f, 0x74, 0x71, 0x4d, 0x67, 0x64, 0x2c, 0x75, 0xc6, 0x9d, 0xd0, 0x35,
	0x3b, 0x44, 0xbf, 0xd7, 0xf4, 0x92, 0x1c, 0xf3, 0xa5, 0x6e, 0x1d, 0x3b, 0xc2, 0x40, 0x14, 0xf4,
	0x1d, 0xb9, 0xb0, 0x95, 0x36, 0xee, 0x13, 0xd8, 0xcc, 0x88, 0x06, 0x33, 0x8f, 0x11, 0xf8, 0xc7,
	0xa7, 0x8c, 0x9c, 0xac, 0x61, 0x69, 0x85, 0x03, 0x36, 0x42, 0x64, 0x2b, 0x43, 0x6d, 0x50, 0x5c,
	0x48, 0x76, 0x12, 0x6b, 0xa3, 0xa0, 0x73, 0x32, 0x69, 0x2a, 0x5d, 0xc3, 0xb7, 0x56, 0x2d, 0xc1,
	0xb0, 0x31, 0xc6, 0xf6, 0x2d, 0x9a, 0x90, 0x91, 0xd5, 0x99, 0xe0, 0x92, 0x9d, 0xce, 0x0f, 0x6f,
	0x26, 0x8b, 0xab, 0xa4, 0x3b, 0x75, 0xf2, 0x03, 0xed, 0x8f, 0x59, 0xa6, 0xdb, 0xda, 0xa5, 0x1d,
	0x45, 0x5f, 0x92, 0x53, 0xbe, 0x2a, 0x53, 0x7f, 0xa0, 0xba, 0x64, 0xc4, 0xd7, 0x9b, 0xa6, 0x3b,
	0x83, 0xbe, 0x22, 0xa4, 0x6e, 0x55, 0x14, 0x96, 0x4d, 0x30, 0xbc, 0xe7, 0x60, 0xc7, 0x6c, 0xb1,
	0x66, 0x67, 0x3e, 0x32, 0x4e, 0x71, 0x1d, 0xba, 0xbb, 0x82, 0x3a, 0xd7, 0x86, 0x4d, 0xd1, 0xed,
	0x54, 0xd8, 0x49, 0xe9, 0x1c, 0x0c, 0x77, 0x3e, 0x74, 0x8e, 0xa1, 0x9d, 0x41, 0xaf, 0xc9, 0x59,
	0x63, 0x84, 0xe2, 0x66, 0x73, 0xa7, 0xa5, 0x07, 0xfe, 0xc7, 0xa3, 0x0d, 0x3c, 0xfa, 0x96, 0x9c,
	0x5b, 0xc8, 0x74, 0x9d, 0xf7, 0xd4, 0x05, 0x52, 0x0f, 0xdc, 0xb0, 0x93, 0x83, 0x5f, 0x2e, 0x22,
	0x4f, 0x10, 0xd9, 0x19, 0xa1, 0x4a, 0x25, 0xca, 0x4a, 0xfa, 0x5f, 0x87, 0xd0, 0x58, 0x65, 0xe8,
	0xfa, 0x4e, 0xd2, 0xde, 0xb9, 0xef, 0xcb, 0x3d, 0x45, 0xf6, 0x91, 0x08, 0x5d, 0x90, 0x33, 0xbe,
	0xe2, 0x8e, 0x9b, 0xcf, 0xdc, 0x56, 0x60, 0xd9, 0xa5, 0x27, 0x27, 0x8b, 0xf3, 0xbe, 0xff, 0x5f,
	0x14, 0x2f, 0x21, 0x1d, 0x30, 0x21, 0xa7, 0x02, 0xee, 0x7b, 0xd0, 0xe5, 0x3c, 0x7b, 0x3c, 0x67,
	0x9f, 0xa1, 0x6f, 0xc8, 0xb4, 0xd0, 0x52, 0xea, 0x35, 0x98, 0xbb, 0x70, 0x95, 0xec, 0x0a, 0xaf,
	0x65, 0x68, 0x86, 0x53, 0x46, 0xc3, 0xdf, 0x53, 0xc4, 0x9e, 0x23, 0xf6, 0xc0, 0x0d, 0x7d, 0x97,
	0xc2, 0xba, 0x9e, 0x62, 0x48, 0x0d, 0x3c, 0xfa, 0xc1, 0x33, 0xdc, 0xba, 0xaf, 0x3a, 0x17, 0x85,
	0x80, 0x9c, 0xbd, 0xc0, 0xaf, 0x9c, 0x25, 0x71, 0xe6, 0x92, 0xed, 0xcc, 0x25, 0xf7, 0xdb, 0x99,
	0x4b, 0x07, 0xfc, 0xec, 0x27, 0x99, 0x0e, 0x1e, 0x5f, 0x78, 0x36, 0x6e, 0xd3, 0x6c, 0xc7, 0x0f,
	0xd7, 0x61, 0xd0, 0x5a, 0x0b, 0x66, 0x6f, 0x00, 0x7b, 0x1d, 0x86, 0xc1, 0x6f, 0xa2, 0x8b, 0x6e,
	0x02, 0xa3, 0x98, 0xfd, 0x26, 0xc7, 0xd8, 0x1f, 0x2c, 0x27, 0xea, 0x4d, 0x5f, 0xce, 0xaf, 0x43,
	0x8a, 0x55, 0x5c, 0xca, 0xae, 0x56, 0x14, 0xe1, 0x6d, 0x2a, 0xc8, 0x45, 0xab, 0xba, 0x4a, 0x9d,
	0x0a, 0xb4, 0xe4, 0xa6, 0x84, 0xed, 0x24, 0xa3, 0x08, 0x9f, 0xa4, 0x8d, 0x28, 0x45, 0xed, 0xa7,
	0x29, 0x4e, 0x70, 0xaf, 0x97, 0x23, 0x3c, 0xf5, 0xfb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa4,
	0x63, 0x7d, 0x5d, 0x88, 0x04, 0x00, 0x00,
}
