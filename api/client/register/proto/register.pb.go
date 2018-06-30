// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/client/register/proto/register.proto

/*
Package api_client_register is a generated protocol buffer package.

It is generated from these files:
	api/client/register/proto/register.proto

It has these top-level messages:
	RegisterReq
	RegisterResp
	UnregisterReq
	UnregisterResp
*/
package api_client_register

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

type METHOD int32

const (
	METHOD_Register   METHOD = 0
	METHOD_UnRegister METHOD = 1
)

var METHOD_name = map[int32]string{
	0: "Register",
	1: "UnRegister",
}
var METHOD_value = map[string]int32{
	"Register":   0,
	"UnRegister": 1,
}

func (x METHOD) String() string {
	return proto.EnumName(METHOD_name, int32(x))
}
func (METHOD) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RegisterReq struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *RegisterReq) Reset()                    { *m = RegisterReq{} }
func (m *RegisterReq) String() string            { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()               {}
func (*RegisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RegisterReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type RegisterResp struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Pwd      string `protobuf:"bytes,3,opt,name=pwd" json:"pwd,omitempty"`
}

func (m *RegisterResp) Reset()                    { *m = RegisterResp{} }
func (m *RegisterResp) String() string            { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()               {}
func (*RegisterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RegisterResp) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func (m *RegisterResp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegisterResp) GetPwd() string {
	if m != nil {
		return m.Pwd
	}
	return ""
}

type UnregisterReq struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *UnregisterReq) Reset()                    { *m = UnregisterReq{} }
func (m *UnregisterReq) String() string            { return proto.CompactTextString(m) }
func (*UnregisterReq) ProtoMessage()               {}
func (*UnregisterReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *UnregisterReq) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

type UnregisterResp struct {
	ClientId string `protobuf:"bytes,1,opt,name=clientId" json:"clientId,omitempty"`
}

func (m *UnregisterResp) Reset()                    { *m = UnregisterResp{} }
func (m *UnregisterResp) String() string            { return proto.CompactTextString(m) }
func (*UnregisterResp) ProtoMessage()               {}
func (*UnregisterResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *UnregisterResp) GetClientId() string {
	if m != nil {
		return m.ClientId
	}
	return ""
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "api_client_register.RegisterReq")
	proto.RegisterType((*RegisterResp)(nil), "api_client_register.RegisterResp")
	proto.RegisterType((*UnregisterReq)(nil), "api_client_register.UnregisterReq")
	proto.RegisterType((*UnregisterResp)(nil), "api_client_register.UnregisterResp")
	proto.RegisterEnum("api_client_register.METHOD", METHOD_name, METHOD_value)
}

func init() { proto.RegisterFile("api/client/register/proto/register.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0x2c, 0xc8, 0xd4,
	0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0xd1, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0xd2,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0x87, 0x73, 0xf5, 0xc0, 0x5c, 0x21, 0xe1, 0xc4, 0x82, 0xcc, 0x78,
	0x88, 0xca, 0x78, 0x98, 0x94, 0x92, 0x26, 0x17, 0x77, 0x10, 0x94, 0x1d, 0x94, 0x5a, 0x28, 0x24,
	0xc5, 0xc5, 0x01, 0x51, 0xe1, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe7, 0x2b,
	0x05, 0x70, 0xf1, 0x20, 0x94, 0x16, 0x17, 0xe0, 0x53, 0x2b, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98,
	0x9b, 0x2a, 0xc1, 0x04, 0x16, 0x07, 0xb3, 0x85, 0x04, 0xb8, 0x98, 0x0b, 0xca, 0x53, 0x24, 0x98,
	0xc1, 0x42, 0x20, 0xa6, 0x92, 0x36, 0x17, 0x6f, 0x68, 0x5e, 0x11, 0x91, 0xd6, 0xeb, 0x70, 0xf1,
	0x21, 0x2b, 0xc6, 0xef, 0x00, 0x2d, 0x35, 0x2e, 0x36, 0x5f, 0xd7, 0x10, 0x0f, 0x7f, 0x17, 0x21,
	0x1e, 0x2e, 0x0e, 0x98, 0xb3, 0x05, 0x18, 0x84, 0xf8, 0xb8, 0xb8, 0x42, 0xf3, 0xe0, 0x7c, 0xc6,
	0x24, 0x36, 0x70, 0xd8, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xde, 0xff, 0x11, 0xfe, 0x47,
	0x01, 0x00, 0x00,
}
