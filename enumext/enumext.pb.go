// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/enumext.proto

package enumext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
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

var E_ExtVal = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50003,
	Name:          "common.extVal",
	Tag:           "bytes,50003,opt,name=extVal",
	Filename:      "common/enumext.proto",
}

func init() {
	proto.RegisterExtension(E_ExtVal)
}

func init() { proto.RegisterFile("common/enumext.proto", fileDescriptor_02d6928dd8ffce94) }

var fileDescriptor_02d6928dd8ffce94 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xcd, 0x31, 0x0a, 0xc2, 0x30,
	0x18, 0xc5, 0x71, 0x45, 0x28, 0xd8, 0xb1, 0x38, 0x88, 0x20, 0x54, 0x27, 0xa7, 0x2f, 0x83, 0x9b,
	0x6e, 0x42, 0x67, 0xc5, 0xa1, 0x83, 0x5b, 0x93, 0xc6, 0x36, 0x90, 0xe4, 0x85, 0x36, 0x81, 0xe2,
	0x01, 0xbc, 0xa0, 0x17, 0x12, 0x9a, 0xea, 0xfa, 0xf8, 0x3d, 0xfe, 0xe9, 0x4a, 0xc0, 0x18, 0x58,
	0x26, 0x6d, 0x30, 0x72, 0xf0, 0xe4, 0x3a, 0x78, 0x64, 0x49, 0x5c, 0x37, 0x79, 0x03, 0x34, 0x5a,
	0xb2, 0x71, 0xe5, 0xe1, 0xc9, 0x6a, 0xd9, 0x8b, 0x4e, 0x39, 0x8f, 0x2e, 0xca, 0xd3, 0x39, 0x4d,
	0xe4, 0xe0, 0xcb, 0x4a, 0x67, 0x3b, 0x8a, 0x98, 0x7e, 0x98, 0x0a, 0x1b, 0x4c, 0x59, 0xe9, 0x20,
	0xaf, 0xce, 0x2b, 0xd8, 0x7e, 0xfd, 0x79, 0x2f, 0xf2, 0xf9, 0x61, 0x79, 0x9f, 0x2e, 0x97, 0x22,
	0xdd, 0x0a, 0x18, 0x52, 0xaf, 0xb6, 0x42, 0x1b, 0xa8, 0x86, 0x76, 0xad, 0xb2, 0xe4, 0x38, 0xc5,
	0xfe, 0x6d, 0xf6, 0xd8, 0x0b, 0xd4, 0xf2, 0x2f, 0x04, 0x0c, 0x9b, 0x14, 0x73, 0x9c, 0x45, 0xc5,
	0x93, 0xb1, 0x78, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x5b, 0xa6, 0xe6, 0x08, 0xcc, 0x00, 0x00,
	0x00,
}
