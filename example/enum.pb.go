// Code generated by protoc-gen-go. DO NOT EDIT.
// source: enum.proto

package common

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

// 通用状态，复数 states
type State int32

const (
	State_STATE_UNSPECIFIED State = 0
	State_STATE_ACTIVE      State = 1
	State_STATE_INACTIVE    State = 2
	State_STATE_DELETED     State = 3
	State_STATE_DENY        State = 4
)

var State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "STATE_ACTIVE",
	2: "STATE_INACTIVE",
	3: "STATE_DELETED",
	4: "STATE_DENY",
}

var State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"STATE_ACTIVE":      1,
	"STATE_INACTIVE":    2,
	"STATE_DELETED":     3,
	"STATE_DENY":        4,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_13a9f1b5947140c8, []int{0}
}

func init() {
	proto.RegisterEnum("common.State", State_name, State_value)
}

func init() { proto.RegisterFile("enum.proto", fileDescriptor_13a9f1b5947140c8) }

var fileDescriptor_13a9f1b5947140c8 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xcd, 0x2b, 0xcd,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0x93, 0xe2,
	0x05, 0x89, 0xa5, 0x56, 0x94, 0x40, 0x84, 0xb5, 0xca, 0xb9, 0x58, 0x83, 0x4b, 0x12, 0x4b, 0x52,
	0x85, 0x44, 0xb9, 0x04, 0x83, 0x43, 0x1c, 0x43, 0x5c, 0xe3, 0x43, 0xfd, 0x82, 0x03, 0x5c, 0x9d,
	0x3d, 0xdd, 0x3c, 0x5d, 0x5d, 0x04, 0x18, 0x84, 0x14, 0xb8, 0x78, 0x20, 0xc2, 0x8e, 0xce, 0x21,
	0x9e, 0x61, 0xae, 0x02, 0x8c, 0x52, 0x7c, 0xb3, 0xb6, 0x4a, 0x70, 0xa5, 0xa5, 0xe6, 0x64, 0x56,
	0xc4, 0x97, 0xa4, 0x56, 0x94, 0x08, 0x09, 0x71, 0xf1, 0x41, 0x54, 0x78, 0xfa, 0x41, 0xd5, 0x30,
	0x09, 0x09, 0x72, 0xf1, 0x42, 0xc4, 0x5c, 0x5c, 0x7d, 0x5c, 0x43, 0x5c, 0x5d, 0x04, 0x98, 0x85,
	0xf8, 0xb8, 0xb8, 0x60, 0x42, 0x7e, 0x91, 0x02, 0x2c, 0x49, 0x6c, 0x60, 0xfb, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xd1, 0xe5, 0x0f, 0x97, 0xa4, 0x00, 0x00, 0x00,
}
