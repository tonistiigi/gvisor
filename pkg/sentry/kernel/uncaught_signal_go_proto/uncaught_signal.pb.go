// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/sentry/kernel/uncaught_signal.proto

package uncaught_signal_go_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import registers_go_proto "gvisor.googlesource.com/gvisor/pkg/sentry/arch/registers_go_proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UncaughtSignal struct {
	Tid                  int32                         `protobuf:"varint,1,opt,name=tid,proto3" json:"tid,omitempty"`
	Pid                  int32                         `protobuf:"varint,2,opt,name=pid,proto3" json:"pid,omitempty"`
	Registers            *registers_go_proto.Registers `protobuf:"bytes,3,opt,name=registers,proto3" json:"registers,omitempty"`
	SignalNumber         int32                         `protobuf:"varint,4,opt,name=signal_number,json=signalNumber,proto3" json:"signal_number,omitempty"`
	FaultAddr            uint64                        `protobuf:"varint,5,opt,name=fault_addr,json=faultAddr,proto3" json:"fault_addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *UncaughtSignal) Reset()         { *m = UncaughtSignal{} }
func (m *UncaughtSignal) String() string { return proto.CompactTextString(m) }
func (*UncaughtSignal) ProtoMessage()    {}
func (*UncaughtSignal) Descriptor() ([]byte, []int) {
	return fileDescriptor_uncaught_signal_855be4d0fd3e569d, []int{0}
}
func (m *UncaughtSignal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UncaughtSignal.Unmarshal(m, b)
}
func (m *UncaughtSignal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UncaughtSignal.Marshal(b, m, deterministic)
}
func (dst *UncaughtSignal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UncaughtSignal.Merge(dst, src)
}
func (m *UncaughtSignal) XXX_Size() int {
	return xxx_messageInfo_UncaughtSignal.Size(m)
}
func (m *UncaughtSignal) XXX_DiscardUnknown() {
	xxx_messageInfo_UncaughtSignal.DiscardUnknown(m)
}

var xxx_messageInfo_UncaughtSignal proto.InternalMessageInfo

func (m *UncaughtSignal) GetTid() int32 {
	if m != nil {
		return m.Tid
	}
	return 0
}

func (m *UncaughtSignal) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

func (m *UncaughtSignal) GetRegisters() *registers_go_proto.Registers {
	if m != nil {
		return m.Registers
	}
	return nil
}

func (m *UncaughtSignal) GetSignalNumber() int32 {
	if m != nil {
		return m.SignalNumber
	}
	return 0
}

func (m *UncaughtSignal) GetFaultAddr() uint64 {
	if m != nil {
		return m.FaultAddr
	}
	return 0
}

func init() {
	proto.RegisterType((*UncaughtSignal)(nil), "gvisor.UncaughtSignal")
}

func init() {
	proto.RegisterFile("pkg/sentry/kernel/uncaught_signal.proto", fileDescriptor_uncaught_signal_855be4d0fd3e569d)
}

var fileDescriptor_uncaught_signal_855be4d0fd3e569d = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x4d, 0x4a, 0xc6, 0x30,
	0x10, 0x86, 0x89, 0xfd, 0x81, 0xc6, 0x1f, 0x34, 0xab, 0x20, 0x88, 0x45, 0x17, 0x76, 0xd5, 0x80,
	0x9e, 0xc0, 0x0b, 0xb8, 0x88, 0xb8, 0x2e, 0x69, 0x13, 0xd3, 0xd0, 0x9a, 0x86, 0x49, 0x22, 0x78,
	0x24, 0x6f, 0x29, 0x4d, 0xd4, 0xef, 0xdb, 0x0d, 0xcf, 0xbc, 0xf3, 0xcc, 0x8b, 0x1f, 0xdc, 0xa2,
	0x99, 0x57, 0x36, 0xc0, 0x17, 0x5b, 0x14, 0x58, 0xb5, 0xb2, 0x68, 0x27, 0x11, 0xf5, 0x1c, 0x06,
	0x6f, 0xb4, 0x15, 0x6b, 0xef, 0x60, 0x0b, 0x1b, 0xa9, 0xf5, 0xa7, 0xf1, 0x1b, 0x5c, 0xdf, 0x1e,
	0x1d, 0x08, 0x98, 0x66, 0x06, 0x4a, 0x1b, 0x1f, 0x14, 0xf8, 0x1c, 0xbc, 0xfb, 0x46, 0xf8, 0xe2,
	0xed, 0x57, 0xf1, 0x9a, 0x0c, 0xe4, 0x12, 0x17, 0xc1, 0x48, 0x8a, 0x5a, 0xd4, 0x55, 0x7c, 0x1f,
	0x77, 0xe2, 0x8c, 0xa4, 0x27, 0x99, 0x38, 0x23, 0x09, 0xc3, 0xcd, 0xbf, 0x89, 0x16, 0x2d, 0xea,
	0x4e, 0x1f, 0xaf, 0xfa, 0xfc, 0xb3, 0xe7, 0x7f, 0x0b, 0x7e, 0xc8, 0x90, 0x7b, 0x7c, 0x9e, 0x0b,
	0x0e, 0x36, 0x7e, 0x8c, 0x0a, 0x68, 0x99, 0x64, 0x67, 0x19, 0xbe, 0x24, 0x46, 0x6e, 0x30, 0x7e,
	0x17, 0x71, 0x0d, 0x83, 0x90, 0x12, 0x68, 0xd5, 0xa2, 0xae, 0xe4, 0x4d, 0x22, 0xcf, 0x52, 0xc2,
	0x58, 0xa7, 0xca, 0x4f, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x62, 0x54, 0xdf, 0x06, 0x01,
	0x00, 0x00,
}
