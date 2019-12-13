// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bottos-project/bottos/api/log.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SetLogConfigItemRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value"`
}

func (m *SetLogConfigItemRequest) Reset()                    { *m = SetLogConfigItemRequest{} }
func (m *SetLogConfigItemRequest) String() string            { return proto.CompactTextString(m) }
func (*SetLogConfigItemRequest) ProtoMessage()               {}
func (*SetLogConfigItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *SetLogConfigItemRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *SetLogConfigItemRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetLogConfigItemResponse struct {
	Errcode uint32 `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
}

func (m *SetLogConfigItemResponse) Reset()                    { *m = SetLogConfigItemResponse{} }
func (m *SetLogConfigItemResponse) String() string            { return proto.CompactTextString(m) }
func (*SetLogConfigItemResponse) ProtoMessage()               {}
func (*SetLogConfigItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *SetLogConfigItemResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func init() {
	proto.RegisterType((*SetLogConfigItemRequest)(nil), "api.SetLogConfigItemRequest")
	proto.RegisterType((*SetLogConfigItemResponse)(nil), "api.SetLogConfigItemResponse")
}

func init() { proto.RegisterFile("github.com/bottos-project/bottos/api/log.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xca, 0x2f, 0x29, 0xc9, 0x2f, 0xd6, 0x2d, 0x28,
	0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x81, 0x72, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x73, 0xf2, 0xd3, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x98, 0x13, 0x0b, 0x32, 0x95, 0x1c, 0xb9, 0xc4, 0x83, 0x53,
	0x4b, 0x7c, 0xf2, 0xd3, 0x9d, 0xf3, 0xf3, 0xd2, 0x32, 0xd3, 0x3d, 0x4b, 0x52, 0x73, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0xb3, 0x53, 0x2b, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x38, 0x83, 0x40, 0x4c, 0x21, 0x11, 0x2e, 0xd6, 0xb2, 0xc4, 0x9c, 0xd2, 0x54, 0x09, 0x26,
	0xb0, 0x18, 0x84, 0xa3, 0x64, 0xc2, 0x25, 0x81, 0x69, 0x44, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0xaa,
	0x90, 0x04, 0x17, 0x7b, 0x6a, 0x51, 0x51, 0x72, 0x7e, 0x4a, 0x2a, 0xd8, 0x1c, 0xde, 0x20, 0x18,
	0xd7, 0x49, 0x2d, 0x4a, 0x85, 0x18, 0xf7, 0x26, 0xb1, 0x81, 0x1d, 0x6b, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x1e, 0xea, 0xf2, 0x47, 0xde, 0x00, 0x00, 0x00,
}
