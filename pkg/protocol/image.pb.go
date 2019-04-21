// Code generated by protoc-gen-go. DO NOT EDIT.
// source: image.proto

package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeleteImageRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteImageRequest) Reset()                    { *m = DeleteImageRequest{} }
func (m *DeleteImageRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageRequest) ProtoMessage()               {}
func (*DeleteImageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *DeleteImageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteImageReply struct {
}

func (m *DeleteImageReply) Reset()                    { *m = DeleteImageReply{} }
func (m *DeleteImageReply) String() string            { return proto.CompactTextString(m) }
func (*DeleteImageReply) ProtoMessage()               {}
func (*DeleteImageReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GenerateTokenRequest struct {
	Image    string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Duration int64  `protobuf:"varint,2,opt,name=duration" json:"duration,omitempty"`
}

func (m *GenerateTokenRequest) Reset()                    { *m = GenerateTokenRequest{} }
func (m *GenerateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*GenerateTokenRequest) ProtoMessage()               {}
func (*GenerateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GenerateTokenRequest) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *GenerateTokenRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type GenerateTokenReply struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GenerateTokenReply) Reset()                    { *m = GenerateTokenReply{} }
func (m *GenerateTokenReply) String() string            { return proto.CompactTextString(m) }
func (*GenerateTokenReply) ProtoMessage()               {}
func (*GenerateTokenReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GenerateTokenReply) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*DeleteImageRequest)(nil), "protocol.DeleteImageRequest")
	proto.RegisterType((*DeleteImageReply)(nil), "protocol.DeleteImageReply")
	proto.RegisterType((*GenerateTokenRequest)(nil), "protocol.GenerateTokenRequest")
	proto.RegisterType((*GenerateTokenReply)(nil), "protocol.GenerateTokenReply")
}

func init() { proto.RegisterFile("image.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 161 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xce, 0xcc, 0x4d, 0x4c,
	0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9, 0xf9, 0x39, 0x4a, 0x2a,
	0x5c, 0x42, 0x2e, 0xa9, 0x39, 0xa9, 0x25, 0xa9, 0x9e, 0x20, 0xe9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6,
	0xcc, 0x14, 0x25, 0x21, 0x2e, 0x01, 0x14, 0x55, 0x05, 0x39, 0x95, 0x4a, 0x1e, 0x5c, 0x22, 0xee,
	0xa9, 0x79, 0xa9, 0x45, 0x89, 0x25, 0xa9, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x30, 0xbd, 0x22, 0x5c,
	0xac, 0x60, 0xab, 0xa0, 0xda, 0x21, 0x1c, 0x21, 0x29, 0x2e, 0x8e, 0x94, 0xd2, 0xa2, 0xc4, 0x92,
	0xcc, 0xfc, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x38, 0x5f, 0x49, 0x8b, 0x4b, 0x08,
	0xcd, 0xa4, 0x82, 0x9c, 0x4a, 0x90, 0x39, 0x25, 0x20, 0x1e, 0xcc, 0x1c, 0x30, 0x27, 0x89, 0x0d,
	0xec, 0x72, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x97, 0x62, 0x25, 0x7c, 0xcf, 0x00, 0x00,
	0x00,
}
