// Code generated by protoc-gen-go. DO NOT EDIT.
// source: datahub/keycodes/services.proto

package keycodes

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

// Represents a request for adding a keycode
type AddKeycodeRequest struct {
	Keycode              string   `protobuf:"bytes,1,opt,name=keycode,proto3" json:"keycode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddKeycodeRequest) Reset()         { *m = AddKeycodeRequest{} }
func (m *AddKeycodeRequest) String() string { return proto.CompactTextString(m) }
func (*AddKeycodeRequest) ProtoMessage()    {}
func (*AddKeycodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{0}
}

func (m *AddKeycodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKeycodeRequest.Unmarshal(m, b)
}
func (m *AddKeycodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKeycodeRequest.Marshal(b, m, deterministic)
}
func (m *AddKeycodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKeycodeRequest.Merge(m, src)
}
func (m *AddKeycodeRequest) XXX_Size() int {
	return xxx_messageInfo_AddKeycodeRequest.Size(m)
}
func (m *AddKeycodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKeycodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddKeycodeRequest proto.InternalMessageInfo

func (m *AddKeycodeRequest) GetKeycode() string {
	if m != nil {
		return m.Keycode
	}
	return ""
}

// Represents a response for adding a keycode
type AddKeycodeResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Keycode              *Keycode       `protobuf:"bytes,2,opt,name=keycode,proto3" json:"keycode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *AddKeycodeResponse) Reset()         { *m = AddKeycodeResponse{} }
func (m *AddKeycodeResponse) String() string { return proto.CompactTextString(m) }
func (*AddKeycodeResponse) ProtoMessage()    {}
func (*AddKeycodeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{1}
}

func (m *AddKeycodeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKeycodeResponse.Unmarshal(m, b)
}
func (m *AddKeycodeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKeycodeResponse.Marshal(b, m, deterministic)
}
func (m *AddKeycodeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKeycodeResponse.Merge(m, src)
}
func (m *AddKeycodeResponse) XXX_Size() int {
	return xxx_messageInfo_AddKeycodeResponse.Size(m)
}
func (m *AddKeycodeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKeycodeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddKeycodeResponse proto.InternalMessageInfo

func (m *AddKeycodeResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AddKeycodeResponse) GetKeycode() *Keycode {
	if m != nil {
		return m.Keycode
	}
	return nil
}

// Represents a request for retrieving keycodes detailed information
type ListKeycodesRequest struct {
	Keycodes             []string `protobuf:"bytes,1,rep,name=keycodes,proto3" json:"keycodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKeycodesRequest) Reset()         { *m = ListKeycodesRequest{} }
func (m *ListKeycodesRequest) String() string { return proto.CompactTextString(m) }
func (*ListKeycodesRequest) ProtoMessage()    {}
func (*ListKeycodesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{2}
}

func (m *ListKeycodesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKeycodesRequest.Unmarshal(m, b)
}
func (m *ListKeycodesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKeycodesRequest.Marshal(b, m, deterministic)
}
func (m *ListKeycodesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKeycodesRequest.Merge(m, src)
}
func (m *ListKeycodesRequest) XXX_Size() int {
	return xxx_messageInfo_ListKeycodesRequest.Size(m)
}
func (m *ListKeycodesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKeycodesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListKeycodesRequest proto.InternalMessageInfo

func (m *ListKeycodesRequest) GetKeycodes() []string {
	if m != nil {
		return m.Keycodes
	}
	return nil
}

// Represents a response for a retrieving keycodes detailed information request
type ListKeycodesResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Keycodes             []*Keycode     `protobuf:"bytes,2,rep,name=keycodes,proto3" json:"keycodes,omitempty"`
	Summary              *Keycode       `protobuf:"bytes,3,opt,name=summary,proto3" json:"summary,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ListKeycodesResponse) Reset()         { *m = ListKeycodesResponse{} }
func (m *ListKeycodesResponse) String() string { return proto.CompactTextString(m) }
func (*ListKeycodesResponse) ProtoMessage()    {}
func (*ListKeycodesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{3}
}

func (m *ListKeycodesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKeycodesResponse.Unmarshal(m, b)
}
func (m *ListKeycodesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKeycodesResponse.Marshal(b, m, deterministic)
}
func (m *ListKeycodesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKeycodesResponse.Merge(m, src)
}
func (m *ListKeycodesResponse) XXX_Size() int {
	return xxx_messageInfo_ListKeycodesResponse.Size(m)
}
func (m *ListKeycodesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKeycodesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListKeycodesResponse proto.InternalMessageInfo

func (m *ListKeycodesResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListKeycodesResponse) GetKeycodes() []*Keycode {
	if m != nil {
		return m.Keycodes
	}
	return nil
}

func (m *ListKeycodesResponse) GetSummary() *Keycode {
	if m != nil {
		return m.Summary
	}
	return nil
}

// Represents a request for deleting a keycode
type DeleteKeycodeRequest struct {
	Keycode              string   `protobuf:"bytes,1,opt,name=keycode,proto3" json:"keycode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKeycodeRequest) Reset()         { *m = DeleteKeycodeRequest{} }
func (m *DeleteKeycodeRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteKeycodeRequest) ProtoMessage()    {}
func (*DeleteKeycodeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{4}
}

func (m *DeleteKeycodeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKeycodeRequest.Unmarshal(m, b)
}
func (m *DeleteKeycodeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKeycodeRequest.Marshal(b, m, deterministic)
}
func (m *DeleteKeycodeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKeycodeRequest.Merge(m, src)
}
func (m *DeleteKeycodeRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteKeycodeRequest.Size(m)
}
func (m *DeleteKeycodeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKeycodeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKeycodeRequest proto.InternalMessageInfo

func (m *DeleteKeycodeRequest) GetKeycode() string {
	if m != nil {
		return m.Keycode
	}
	return ""
}

// Represents a request for generating license registration data
type GenerateRegistrationDataResponse struct {
	Status               *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Data                 string         `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GenerateRegistrationDataResponse) Reset()         { *m = GenerateRegistrationDataResponse{} }
func (m *GenerateRegistrationDataResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateRegistrationDataResponse) ProtoMessage()    {}
func (*GenerateRegistrationDataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{5}
}

func (m *GenerateRegistrationDataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRegistrationDataResponse.Unmarshal(m, b)
}
func (m *GenerateRegistrationDataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRegistrationDataResponse.Marshal(b, m, deterministic)
}
func (m *GenerateRegistrationDataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRegistrationDataResponse.Merge(m, src)
}
func (m *GenerateRegistrationDataResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateRegistrationDataResponse.Size(m)
}
func (m *GenerateRegistrationDataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRegistrationDataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRegistrationDataResponse proto.InternalMessageInfo

func (m *GenerateRegistrationDataResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GenerateRegistrationDataResponse) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

// Represents a request for activating license signature data
type ActivateRegistrationDataRequest struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateRegistrationDataRequest) Reset()         { *m = ActivateRegistrationDataRequest{} }
func (m *ActivateRegistrationDataRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateRegistrationDataRequest) ProtoMessage()    {}
func (*ActivateRegistrationDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a2f67bf7c182d88, []int{6}
}

func (m *ActivateRegistrationDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActivateRegistrationDataRequest.Unmarshal(m, b)
}
func (m *ActivateRegistrationDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActivateRegistrationDataRequest.Marshal(b, m, deterministic)
}
func (m *ActivateRegistrationDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateRegistrationDataRequest.Merge(m, src)
}
func (m *ActivateRegistrationDataRequest) XXX_Size() int {
	return xxx_messageInfo_ActivateRegistrationDataRequest.Size(m)
}
func (m *ActivateRegistrationDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateRegistrationDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateRegistrationDataRequest proto.InternalMessageInfo

func (m *ActivateRegistrationDataRequest) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*AddKeycodeRequest)(nil), "containersai.datahub.keycodes.AddKeycodeRequest")
	proto.RegisterType((*AddKeycodeResponse)(nil), "containersai.datahub.keycodes.AddKeycodeResponse")
	proto.RegisterType((*ListKeycodesRequest)(nil), "containersai.datahub.keycodes.ListKeycodesRequest")
	proto.RegisterType((*ListKeycodesResponse)(nil), "containersai.datahub.keycodes.ListKeycodesResponse")
	proto.RegisterType((*DeleteKeycodeRequest)(nil), "containersai.datahub.keycodes.DeleteKeycodeRequest")
	proto.RegisterType((*GenerateRegistrationDataResponse)(nil), "containersai.datahub.keycodes.GenerateRegistrationDataResponse")
	proto.RegisterType((*ActivateRegistrationDataRequest)(nil), "containersai.datahub.keycodes.ActivateRegistrationDataRequest")
}

func init() { proto.RegisterFile("datahub/keycodes/services.proto", fileDescriptor_0a2f67bf7c182d88) }

var fileDescriptor_0a2f67bf7c182d88 = []byte{
	// 459 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x6e, 0xd6, 0x69, 0xd0, 0x07, 0x08, 0xf1, 0x98, 0x20, 0x0a, 0x42, 0xab, 0x7c, 0x40, 0x13,
	0x52, 0xed, 0xad, 0x13, 0x57, 0x60, 0xd3, 0x10, 0x07, 0x38, 0x65, 0xb7, 0xdd, 0xdc, 0xf4, 0x51,
	0x2c, 0xd6, 0x38, 0xd8, 0xce, 0xa4, 0x5e, 0xf9, 0x71, 0x88, 0x9f, 0x85, 0x9a, 0xd8, 0x6d, 0xc7,
	0xb2, 0x85, 0xf4, 0x66, 0xc7, 0xef, 0xfb, 0xbe, 0xf7, 0xfc, 0x7d, 0x31, 0x1c, 0x4c, 0xa5, 0x93,
	0xdf, 0xcb, 0x89, 0xf8, 0x41, 0x8b, 0x4c, 0x4f, 0xc9, 0x0a, 0x4b, 0xe6, 0x5a, 0x65, 0x64, 0x79,
	0x61, 0xb4, 0xd3, 0xf8, 0x3a, 0xd3, 0xb9, 0x93, 0x2a, 0x27, 0x63, 0xa5, 0xe2, 0xbe, 0x9a, 0x87,
	0xea, 0xe4, 0x36, 0x3e, 0x2c, 0x6a, 0x7c, 0xf2, 0x6a, 0xa6, 0xf5, 0xec, 0x8a, 0x44, 0xb5, 0x9b,
	0x94, 0xdf, 0x04, 0xcd, 0x0b, 0xb7, 0xf0, 0x87, 0x2f, 0xfd, 0xa1, 0x29, 0x32, 0x61, 0x9d, 0x74,
	0xa5, 0x47, 0xb1, 0x11, 0x3c, 0x3b, 0x9d, 0x4e, 0xbf, 0xd4, 0x54, 0x29, 0xfd, 0x2c, 0xc9, 0x3a,
	0x8c, 0xe1, 0x81, 0x27, 0x8f, 0xa3, 0x61, 0x74, 0x38, 0x48, 0xc3, 0x96, 0xfd, 0x8a, 0x00, 0x37,
	0xeb, 0x6d, 0xa1, 0x73, 0x4b, 0xf8, 0x16, 0xf6, 0x6a, 0xd6, 0xaa, 0xfe, 0xd1, 0x18, 0x79, 0xad,
	0xc7, 0x4d, 0x91, 0xf1, 0x8b, 0xea, 0x24, 0xf5, 0x15, 0xf8, 0x71, 0x4d, 0xbe, 0x53, 0x15, 0xbf,
	0xe1, 0xf7, 0x4e, 0xce, 0x83, 0xd8, 0xaa, 0x89, 0x63, 0x78, 0xfe, 0x55, 0x59, 0xe7, 0xbf, 0xdb,
	0xd0, 0x75, 0x02, 0x0f, 0x03, 0x26, 0x8e, 0x86, 0xfd, 0xc3, 0x41, 0xba, 0xda, 0xb3, 0x3f, 0x11,
	0xec, 0xdf, 0xc4, 0x6c, 0xd1, 0xf9, 0xd9, 0x86, 0xc0, 0xce, 0xb0, 0xdf, 0xa1, 0xf5, 0x15, 0x6e,
	0x39, 0xbd, 0x2d, 0xe7, 0x73, 0x69, 0x16, 0x71, 0xbf, 0xdb, 0xf4, 0x1e, 0xc6, 0x8e, 0x60, 0xff,
	0x9c, 0xae, 0xc8, 0xd1, 0x7f, 0x9b, 0x36, 0x81, 0xe1, 0x67, 0xca, 0xc9, 0x48, 0x47, 0x29, 0xcd,
	0x94, 0x75, 0x46, 0x3a, 0xa5, 0xf3, 0x73, 0xe9, 0xe4, 0x56, 0xf7, 0x80, 0xb0, 0xbb, 0x6c, 0xb3,
	0xb2, 0x6f, 0x90, 0x56, 0x6b, 0xf6, 0x0e, 0x0e, 0x4e, 0x33, 0xa7, 0xae, 0x1b, 0x35, 0xea, 0x06,
	0x03, 0x2c, 0x5a, 0xc3, 0xc6, 0xbf, 0x77, 0xe1, 0x69, 0xf0, 0xe4, 0xa2, 0xfe, 0x1f, 0xd0, 0x02,
	0xac, 0x23, 0x86, 0x47, 0x2d, 0xf7, 0x73, 0x2b, 0xbd, 0xc9, 0x71, 0x07, 0x44, 0x3d, 0x3d, 0xeb,
	0xe1, 0x02, 0x1e, 0x6f, 0xe6, 0x03, 0xc7, 0x2d, 0x24, 0x0d, 0x01, 0x4c, 0x4e, 0x3a, 0x61, 0x56,
	0xd2, 0x97, 0xf0, 0xe4, 0x86, 0xa1, 0xd8, 0xc6, 0xd3, 0x64, 0x7f, 0xd2, 0x60, 0x18, 0xeb, 0x61,
	0x09, 0xf1, 0x5d, 0xd6, 0xe3, 0x8b, 0x80, 0x08, 0x2f, 0x06, 0xff, 0xb4, 0x7c, 0x31, 0x92, 0x0f,
	0x2d, 0xf2, 0x6d, 0x59, 0x62, 0x3d, 0x2c, 0x20, 0xbe, 0x2b, 0x0d, 0xf8, 0xbe, 0xcd, 0x9e, 0xfb,
	0x63, 0xd4, 0x3c, 0xe8, 0x99, 0xb8, 0x1c, 0xcd, 0x94, 0x5b, 0x12, 0x65, 0x7a, 0x2e, 0xd6, 0x0a,
	0x23, 0xa9, 0x84, 0x2c, 0x94, 0xf8, 0xf7, 0xf5, 0x9c, 0xec, 0x55, 0x53, 0x9f, 0xfc, 0x0d, 0x00,
	0x00, 0xff, 0xff, 0xa5, 0x77, 0xfb, 0x18, 0x98, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KeycodesServiceClient is the client API for KeycodesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KeycodesServiceClient interface {
	// Used to add a keycode
	AddKeycode(ctx context.Context, in *AddKeycodeRequest, opts ...grpc.CallOption) (*AddKeycodeResponse, error)
	// Used to retrieve keycodes detailed information
	ListKeycodes(ctx context.Context, in *ListKeycodesRequest, opts ...grpc.CallOption) (*ListKeycodesResponse, error)
	// Used to delete a keycode
	DeleteKeycode(ctx context.Context, in *DeleteKeycodeRequest, opts ...grpc.CallOption) (*status.Status, error)
	// Used to generate license registration data
	GenerateRegistrationData(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GenerateRegistrationDataResponse, error)
	// Used to activate license signature data
	ActivateRegistrationData(ctx context.Context, in *ActivateRegistrationDataRequest, opts ...grpc.CallOption) (*status.Status, error)
}

type keycodesServiceClient struct {
	cc *grpc.ClientConn
}

func NewKeycodesServiceClient(cc *grpc.ClientConn) KeycodesServiceClient {
	return &keycodesServiceClient{cc}
}

func (c *keycodesServiceClient) AddKeycode(ctx context.Context, in *AddKeycodeRequest, opts ...grpc.CallOption) (*AddKeycodeResponse, error) {
	out := new(AddKeycodeResponse)
	err := c.cc.Invoke(ctx, "/containersai.datahub.keycodes.KeycodesService/AddKeycode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycodesServiceClient) ListKeycodes(ctx context.Context, in *ListKeycodesRequest, opts ...grpc.CallOption) (*ListKeycodesResponse, error) {
	out := new(ListKeycodesResponse)
	err := c.cc.Invoke(ctx, "/containersai.datahub.keycodes.KeycodesService/ListKeycodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycodesServiceClient) DeleteKeycode(ctx context.Context, in *DeleteKeycodeRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/containersai.datahub.keycodes.KeycodesService/DeleteKeycode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycodesServiceClient) GenerateRegistrationData(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GenerateRegistrationDataResponse, error) {
	out := new(GenerateRegistrationDataResponse)
	err := c.cc.Invoke(ctx, "/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keycodesServiceClient) ActivateRegistrationData(ctx context.Context, in *ActivateRegistrationDataRequest, opts ...grpc.CallOption) (*status.Status, error) {
	out := new(status.Status)
	err := c.cc.Invoke(ctx, "/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeycodesServiceServer is the server API for KeycodesService service.
type KeycodesServiceServer interface {
	// Used to add a keycode
	AddKeycode(context.Context, *AddKeycodeRequest) (*AddKeycodeResponse, error)
	// Used to retrieve keycodes detailed information
	ListKeycodes(context.Context, *ListKeycodesRequest) (*ListKeycodesResponse, error)
	// Used to delete a keycode
	DeleteKeycode(context.Context, *DeleteKeycodeRequest) (*status.Status, error)
	// Used to generate license registration data
	GenerateRegistrationData(context.Context, *empty.Empty) (*GenerateRegistrationDataResponse, error)
	// Used to activate license signature data
	ActivateRegistrationData(context.Context, *ActivateRegistrationDataRequest) (*status.Status, error)
}

// UnimplementedKeycodesServiceServer can be embedded to have forward compatible implementations.
type UnimplementedKeycodesServiceServer struct {
}

func (*UnimplementedKeycodesServiceServer) AddKeycode(ctx context.Context, req *AddKeycodeRequest) (*AddKeycodeResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method AddKeycode not implemented")
}
func (*UnimplementedKeycodesServiceServer) ListKeycodes(ctx context.Context, req *ListKeycodesRequest) (*ListKeycodesResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ListKeycodes not implemented")
}
func (*UnimplementedKeycodesServiceServer) DeleteKeycode(ctx context.Context, req *DeleteKeycodeRequest) (*status.Status, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method DeleteKeycode not implemented")
}
func (*UnimplementedKeycodesServiceServer) GenerateRegistrationData(ctx context.Context, req *empty.Empty) (*GenerateRegistrationDataResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method GenerateRegistrationData not implemented")
}
func (*UnimplementedKeycodesServiceServer) ActivateRegistrationData(ctx context.Context, req *ActivateRegistrationDataRequest) (*status.Status, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method ActivateRegistrationData not implemented")
}

func RegisterKeycodesServiceServer(s *grpc.Server, srv KeycodesServiceServer) {
	s.RegisterService(&_KeycodesService_serviceDesc, srv)
}

func _KeycodesService_AddKeycode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeycodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycodesServiceServer).AddKeycode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.keycodes.KeycodesService/AddKeycode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycodesServiceServer).AddKeycode(ctx, req.(*AddKeycodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycodesService_ListKeycodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKeycodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycodesServiceServer).ListKeycodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.keycodes.KeycodesService/ListKeycodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycodesServiceServer).ListKeycodes(ctx, req.(*ListKeycodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycodesService_DeleteKeycode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKeycodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycodesServiceServer).DeleteKeycode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.keycodes.KeycodesService/DeleteKeycode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycodesServiceServer).DeleteKeycode(ctx, req.(*DeleteKeycodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycodesService_GenerateRegistrationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycodesServiceServer).GenerateRegistrationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycodesServiceServer).GenerateRegistrationData(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeycodesService_ActivateRegistrationData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRegistrationDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeycodesServiceServer).ActivateRegistrationData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeycodesServiceServer).ActivateRegistrationData(ctx, req.(*ActivateRegistrationDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeycodesService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "containersai.datahub.keycodes.KeycodesService",
	HandlerType: (*KeycodesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddKeycode",
			Handler:    _KeycodesService_AddKeycode_Handler,
		},
		{
			MethodName: "ListKeycodes",
			Handler:    _KeycodesService_ListKeycodes_Handler,
		},
		{
			MethodName: "DeleteKeycode",
			Handler:    _KeycodesService_DeleteKeycode_Handler,
		},
		{
			MethodName: "GenerateRegistrationData",
			Handler:    _KeycodesService_GenerateRegistrationData_Handler,
		},
		{
			MethodName: "ActivateRegistrationData",
			Handler:    _KeycodesService_ActivateRegistrationData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "datahub/keycodes/services.proto",
}