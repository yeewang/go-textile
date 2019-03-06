// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mobile.proto

package pb

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

type MobileEventType int32

const (
	MobileEventType_NODE_START     MobileEventType = 0
	MobileEventType_NODE_ONLINE    MobileEventType = 1
	MobileEventType_NODE_STOP      MobileEventType = 2
	MobileEventType_WALLET_UPDATE  MobileEventType = 10
	MobileEventType_THREAD_UPDATE  MobileEventType = 11
	MobileEventType_NOTIFICATION   MobileEventType = 12
	MobileEventType_QUERY_RESPONSE MobileEventType = 20
)

var MobileEventType_name = map[int32]string{
	0:  "NODE_START",
	1:  "NODE_ONLINE",
	2:  "NODE_STOP",
	10: "WALLET_UPDATE",
	11: "THREAD_UPDATE",
	12: "NOTIFICATION",
	20: "QUERY_RESPONSE",
}
var MobileEventType_value = map[string]int32{
	"NODE_START":     0,
	"NODE_ONLINE":    1,
	"NODE_STOP":      2,
	"WALLET_UPDATE":  10,
	"THREAD_UPDATE":  11,
	"NOTIFICATION":   12,
	"QUERY_RESPONSE": 20,
}

func (x MobileEventType) String() string {
	return proto.EnumName(MobileEventType_name, int32(x))
}
func (MobileEventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mobile_e3d61632271031cb, []int{0}
}

type MobileQueryEvent_Type int32

const (
	MobileQueryEvent_DATA  MobileQueryEvent_Type = 0
	MobileQueryEvent_DONE  MobileQueryEvent_Type = 1
	MobileQueryEvent_ERROR MobileQueryEvent_Type = 2
)

var MobileQueryEvent_Type_name = map[int32]string{
	0: "DATA",
	1: "DONE",
	2: "ERROR",
}
var MobileQueryEvent_Type_value = map[string]int32{
	"DATA":  0,
	"DONE":  1,
	"ERROR": 2,
}

func (x MobileQueryEvent_Type) String() string {
	return proto.EnumName(MobileQueryEvent_Type_name, int32(x))
}
func (MobileQueryEvent_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_mobile_e3d61632271031cb, []int{2, 0}
}

type MobileWalletAccount struct {
	Seed                 string   `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty"`
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MobileWalletAccount) Reset()         { *m = MobileWalletAccount{} }
func (m *MobileWalletAccount) String() string { return proto.CompactTextString(m) }
func (*MobileWalletAccount) ProtoMessage()    {}
func (*MobileWalletAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_e3d61632271031cb, []int{0}
}
func (m *MobileWalletAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileWalletAccount.Unmarshal(m, b)
}
func (m *MobileWalletAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileWalletAccount.Marshal(b, m, deterministic)
}
func (dst *MobileWalletAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileWalletAccount.Merge(dst, src)
}
func (m *MobileWalletAccount) XXX_Size() int {
	return xxx_messageInfo_MobileWalletAccount.Size(m)
}
func (m *MobileWalletAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileWalletAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MobileWalletAccount proto.InternalMessageInfo

func (m *MobileWalletAccount) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *MobileWalletAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MobilePreparedFiles struct {
	Dir                  *Directory        `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	Pin                  map[string]string `protobuf:"bytes,2,rep,name=pin,proto3" json:"pin,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MobilePreparedFiles) Reset()         { *m = MobilePreparedFiles{} }
func (m *MobilePreparedFiles) String() string { return proto.CompactTextString(m) }
func (*MobilePreparedFiles) ProtoMessage()    {}
func (*MobilePreparedFiles) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_e3d61632271031cb, []int{1}
}
func (m *MobilePreparedFiles) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobilePreparedFiles.Unmarshal(m, b)
}
func (m *MobilePreparedFiles) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobilePreparedFiles.Marshal(b, m, deterministic)
}
func (dst *MobilePreparedFiles) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobilePreparedFiles.Merge(dst, src)
}
func (m *MobilePreparedFiles) XXX_Size() int {
	return xxx_messageInfo_MobilePreparedFiles.Size(m)
}
func (m *MobilePreparedFiles) XXX_DiscardUnknown() {
	xxx_messageInfo_MobilePreparedFiles.DiscardUnknown(m)
}

var xxx_messageInfo_MobilePreparedFiles proto.InternalMessageInfo

func (m *MobilePreparedFiles) GetDir() *Directory {
	if m != nil {
		return m.Dir
	}
	return nil
}

func (m *MobilePreparedFiles) GetPin() map[string]string {
	if m != nil {
		return m.Pin
	}
	return nil
}

type MobileQueryEvent struct {
	Id                   string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type                 MobileQueryEvent_Type `protobuf:"varint,2,opt,name=type,proto3,enum=MobileQueryEvent_Type" json:"type,omitempty"`
	Data                 *QueryResult          `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Error                *Error                `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MobileQueryEvent) Reset()         { *m = MobileQueryEvent{} }
func (m *MobileQueryEvent) String() string { return proto.CompactTextString(m) }
func (*MobileQueryEvent) ProtoMessage()    {}
func (*MobileQueryEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_mobile_e3d61632271031cb, []int{2}
}
func (m *MobileQueryEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MobileQueryEvent.Unmarshal(m, b)
}
func (m *MobileQueryEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MobileQueryEvent.Marshal(b, m, deterministic)
}
func (dst *MobileQueryEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MobileQueryEvent.Merge(dst, src)
}
func (m *MobileQueryEvent) XXX_Size() int {
	return xxx_messageInfo_MobileQueryEvent.Size(m)
}
func (m *MobileQueryEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MobileQueryEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MobileQueryEvent proto.InternalMessageInfo

func (m *MobileQueryEvent) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MobileQueryEvent) GetType() MobileQueryEvent_Type {
	if m != nil {
		return m.Type
	}
	return MobileQueryEvent_DATA
}

func (m *MobileQueryEvent) GetData() *QueryResult {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *MobileQueryEvent) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
	proto.RegisterType((*MobileWalletAccount)(nil), "MobileWalletAccount")
	proto.RegisterType((*MobilePreparedFiles)(nil), "MobilePreparedFiles")
	proto.RegisterMapType((map[string]string)(nil), "MobilePreparedFiles.PinEntry")
	proto.RegisterType((*MobileQueryEvent)(nil), "MobileQueryEvent")
	proto.RegisterEnum("MobileEventType", MobileEventType_name, MobileEventType_value)
	proto.RegisterEnum("MobileQueryEvent_Type", MobileQueryEvent_Type_name, MobileQueryEvent_Type_value)
}

func init() { proto.RegisterFile("mobile.proto", fileDescriptor_mobile_e3d61632271031cb) }

var fileDescriptor_mobile_e3d61632271031cb = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xcd, 0x6a, 0xdb, 0x40,
	0x18, 0x8c, 0x7e, 0x9c, 0xc6, 0x9f, 0x6c, 0x67, 0xbb, 0x0d, 0x45, 0x84, 0x14, 0x8c, 0xa1, 0x10,
	0x72, 0x50, 0xc1, 0x85, 0x52, 0x7a, 0x53, 0xa3, 0x0d, 0x35, 0xb8, 0x92, 0xb2, 0x56, 0x08, 0xed,
	0xc5, 0xc8, 0xd6, 0x47, 0x59, 0xaa, 0x48, 0xea, 0x6a, 0xed, 0xa2, 0x87, 0xe8, 0x1b, 0xf4, 0x35,
	0xfa, 0x7e, 0x45, 0x2b, 0xe9, 0x52, 0x7a, 0x9b, 0x19, 0xcd, 0x0c, 0xf3, 0x89, 0x85, 0xc9, 0x53,
	0xb9, 0x13, 0x39, 0x7a, 0x95, 0x2c, 0x55, 0x79, 0x09, 0x47, 0x81, 0x3f, 0x7b, 0xec, 0xfc, 0x38,
	0xa0, 0x6c, 0x7a, 0x32, 0x7d, 0xc2, 0xba, 0x4e, 0xbf, 0xf5, 0xbe, 0xc5, 0x2d, 0xbc, 0xf8, 0xac,
	0x73, 0x8f, 0x69, 0x9e, 0xa3, 0xf2, 0xf7, 0xfb, 0xf2, 0x50, 0x28, 0x4a, 0xc1, 0xae, 0x11, 0x33,
	0xd7, 0x98, 0x1b, 0xd7, 0x63, 0xae, 0x31, 0x75, 0xe1, 0x59, 0x9a, 0x65, 0x12, 0xeb, 0xda, 0x35,
	0xb5, 0x3c, 0xd0, 0xc5, 0x6f, 0x63, 0x68, 0x89, 0x25, 0x56, 0xa9, 0xc4, 0xec, 0x4e, 0xe4, 0x58,
	0xd3, 0x2b, 0xb0, 0x32, 0x21, 0x75, 0x89, 0xb3, 0x04, 0x2f, 0x10, 0x12, 0xf7, 0xaa, 0x94, 0x0d,
	0x6f, 0x65, 0xfa, 0x06, 0xac, 0x4a, 0x14, 0xae, 0x39, 0xb7, 0xae, 0x9d, 0xe5, 0x2b, 0xef, 0x3f,
	0x05, 0x5e, 0x2c, 0x0a, 0x56, 0xa8, 0x36, 0x50, 0x89, 0xe2, 0xf2, 0x1d, 0x9c, 0x0d, 0x02, 0x25,
	0x60, 0x7d, 0xc7, 0xa6, 0xdf, 0xd7, 0x42, 0x7a, 0x01, 0xa3, 0x63, 0x9a, 0x1f, 0xb0, 0x1f, 0xd7,
	0x91, 0x0f, 0xe6, 0x7b, 0x63, 0xf1, 0xc7, 0x00, 0xd2, 0xb5, 0xdf, 0xb7, 0x3f, 0x82, 0x1d, 0xb1,
	0x50, 0x74, 0x06, 0xa6, 0x18, 0xee, 0x33, 0x45, 0x46, 0x6f, 0xc0, 0x56, 0x4d, 0xd5, 0xa5, 0x67,
	0xcb, 0x97, 0xde, 0xbf, 0x01, 0x2f, 0x69, 0x2a, 0xe4, 0xda, 0x43, 0xe7, 0x60, 0x67, 0xa9, 0x4a,
	0x5d, 0x4b, 0x1f, 0x36, 0xf1, 0xb4, 0x8b, 0x63, 0x7d, 0xc8, 0x15, 0xd7, 0x5f, 0xe8, 0x15, 0x8c,
	0x50, 0xca, 0x52, 0xba, 0xb6, 0xb6, 0x9c, 0x7a, 0xac, 0x65, 0xbc, 0x13, 0x17, 0xaf, 0xc1, 0x6e,
	0xdb, 0xe8, 0x19, 0xd8, 0x81, 0x9f, 0xf8, 0xe4, 0x44, 0xa3, 0x28, 0x64, 0xc4, 0xa0, 0x63, 0x18,
	0x31, 0xce, 0x23, 0x4e, 0xcc, 0x9b, 0x5f, 0x06, 0x9c, 0x77, 0x33, 0xf4, 0x02, 0x1d, 0x99, 0x01,
	0x84, 0x51, 0xc0, 0xb6, 0x9b, 0xc4, 0xe7, 0x09, 0x39, 0xa1, 0xe7, 0xe0, 0x68, 0x1e, 0x85, 0xeb,
	0x95, 0xce, 0x4f, 0x61, 0xdc, 0x1b, 0xa2, 0x98, 0x98, 0xf4, 0x39, 0x4c, 0x1f, 0xfd, 0xf5, 0x9a,
	0x25, 0xdb, 0x87, 0x38, 0xf0, 0x13, 0x46, 0xa0, 0x95, 0x92, 0x4f, 0x9c, 0xf9, 0xc1, 0x20, 0x39,
	0x94, 0xc0, 0x24, 0x8c, 0x92, 0xd5, 0xdd, 0xea, 0xd6, 0x4f, 0x56, 0x51, 0x48, 0x26, 0x94, 0xc2,
	0xec, 0xfe, 0x81, 0xf1, 0x2f, 0x5b, 0xce, 0x36, 0x71, 0x14, 0x6e, 0x18, 0xb9, 0xf8, 0x68, 0x7f,
	0x35, 0xab, 0xdd, 0xee, 0x54, 0x3f, 0x9c, 0xb7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x94,
	0x48, 0xa4, 0x70, 0x02, 0x00, 0x00,
}