// Code generated by protoc-gen-go. DO NOT EDIT.
// source: container.proto

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	container.proto
	image.proto

It has these top-level messages:
	ContainerDeploymentRequest
	ImageDeleteRequest
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ContainerDeploymentRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ContainerDeploymentRequest) Reset()                    { *m = ContainerDeploymentRequest{} }
func (m *ContainerDeploymentRequest) String() string            { return proto.CompactTextString(m) }
func (*ContainerDeploymentRequest) ProtoMessage()               {}
func (*ContainerDeploymentRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ContainerDeploymentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*ContainerDeploymentRequest)(nil), "protocol.ContainerDeploymentRequest")
}

func init() { proto.RegisterFile("container.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 93 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xce, 0xcf, 0x2b,
	0x49, 0xcc, 0xcc, 0x4b, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0xc9,
	0xf9, 0x39, 0x4a, 0x3a, 0x5c, 0x52, 0xce, 0x30, 0x49, 0x97, 0xd4, 0x82, 0x9c, 0xfc, 0xca, 0xdc,
	0xd4, 0xbc, 0x92, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0xa6, 0xcc, 0x94, 0x24, 0x36, 0xb0, 0x3e, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x60, 0xdb, 0xb4, 0x48, 0x51, 0x00, 0x00, 0x00,
}
