// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package dispatcher

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

type SessionConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionConfig) Reset()         { *m = SessionConfig{} }
func (m *SessionConfig) String() string { return proto.CompactTextString(m) }
func (*SessionConfig) ProtoMessage()    {}
func (*SessionConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{0}
}

func (m *SessionConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionConfig.Unmarshal(m, b)
}
func (m *SessionConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionConfig.Marshal(b, m, deterministic)
}
func (m *SessionConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionConfig.Merge(m, src)
}
func (m *SessionConfig) XXX_Size() int {
	return xxx_messageInfo_SessionConfig.Size(m)
}
func (m *SessionConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SessionConfig proto.InternalMessageInfo

type Config struct {
	Settings             *SessionConfig `protobuf:"bytes,1,opt,name=settings,proto3" json:"settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_3eaf2c85e69e9ea4, []int{1}
}

func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (m *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(m, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetSettings() *SessionConfig {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*SessionConfig)(nil), "xflash.server.trojan.app.dispatcher.SessionConfig")
	proto.RegisterType((*Config)(nil), "xflash.server.trojan.app.dispatcher.Config")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_3eaf2c85e69e9ea4) }

var fileDescriptor_3eaf2c85e69e9ea4 = []byte{
	// 194 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0xae, 0x48, 0xcb, 0x49, 0x2c, 0xce,
	0xd0, 0x2b, 0x4e, 0x2d, 0x2a, 0x4b, 0x2d, 0xd2, 0x2b, 0x29, 0xca, 0xcf, 0x4a, 0xcc, 0xd3, 0x4b,
	0x2c, 0x28, 0xd0, 0x4b, 0xc9, 0x2c, 0x2e, 0x48, 0x2c, 0x49, 0xce, 0x48, 0x2d, 0x52, 0x12, 0xe5,
	0xe2, 0x0d, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0x73, 0x06, 0xeb, 0xf5, 0x62, 0xe1, 0x60, 0x14,
	0x60, 0x52, 0x8a, 0xe0, 0x62, 0x83, 0xf0, 0x85, 0xfc, 0xb8, 0x38, 0x8a, 0x53, 0x4b, 0x4a, 0x32,
	0xf3, 0xd2, 0x8b, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x8c, 0xf4, 0x88, 0x30, 0x58, 0x0f,
	0xc5, 0xd4, 0x20, 0xb8, 0x19, 0x4e, 0xed, 0x8c, 0x5c, 0xea, 0xc9, 0xf9, 0xb9, 0xc4, 0x98, 0x11,
	0xc0, 0x18, 0x65, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xd1,
	0xa1, 0x5b, 0x90, 0x98, 0x97, 0x92, 0xa8, 0x0f, 0xd1, 0xa7, 0x0b, 0xd1, 0xa7, 0x9f, 0x58, 0x50,
	0xa0, 0x8f, 0xd0, 0xb7, 0x8a, 0x09, 0xbb, 0xd7, 0x1d, 0x0b, 0x0a, 0xf4, 0x5c, 0xe0, 0xaa, 0x92,
	0xd8, 0xc0, 0xc1, 0x64, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x25, 0xbf, 0x81, 0xc8, 0x36, 0x01,
	0x00, 0x00,
}
