// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/supplier/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types1 "github.com/pokt-network/poktroll/x/shared/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgUpdateParams is the Msg/UpdateParams request type.
type MsgUpdateParams struct {
	// authority is the address that controls the module (defaults to x/gov unless overwritten).
	Authority string `protobuf:"bytes,1,opt,name=authority,proto3" json:"authority,omitempty"`
	// params defines the x/supplier parameters to update.
	// NOTE: All parameters must be supplied.
	Params Params `protobuf:"bytes,2,opt,name=params,proto3" json:"params"`
}

func (m *MsgUpdateParams) Reset()         { *m = MsgUpdateParams{} }
func (m *MsgUpdateParams) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParams) ProtoMessage()    {}
func (*MsgUpdateParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{0}
}
func (m *MsgUpdateParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParams.Merge(m, src)
}
func (m *MsgUpdateParams) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParams) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParams.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParams proto.InternalMessageInfo

func (m *MsgUpdateParams) GetAuthority() string {
	if m != nil {
		return m.Authority
	}
	return ""
}

func (m *MsgUpdateParams) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// MsgUpdateParamsResponse defines the response structure for executing a
// MsgUpdateParams message.
type MsgUpdateParamsResponse struct {
}

func (m *MsgUpdateParamsResponse) Reset()         { *m = MsgUpdateParamsResponse{} }
func (m *MsgUpdateParamsResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateParamsResponse) ProtoMessage()    {}
func (*MsgUpdateParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{1}
}
func (m *MsgUpdateParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateParamsResponse.Merge(m, src)
}
func (m *MsgUpdateParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateParamsResponse proto.InternalMessageInfo

type MsgStakeSupplier struct {
	Address  string                          `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Stake    *types.Coin                     `protobuf:"bytes,2,opt,name=stake,proto3" json:"stake,omitempty"`
	Services []*types1.SupplierServiceConfig `protobuf:"bytes,3,rep,name=services,proto3" json:"services,omitempty"`
}

func (m *MsgStakeSupplier) Reset()         { *m = MsgStakeSupplier{} }
func (m *MsgStakeSupplier) String() string { return proto.CompactTextString(m) }
func (*MsgStakeSupplier) ProtoMessage()    {}
func (*MsgStakeSupplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{2}
}
func (m *MsgStakeSupplier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStakeSupplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStakeSupplier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStakeSupplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeSupplier.Merge(m, src)
}
func (m *MsgStakeSupplier) XXX_Size() int {
	return m.Size()
}
func (m *MsgStakeSupplier) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeSupplier.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeSupplier proto.InternalMessageInfo

func (m *MsgStakeSupplier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgStakeSupplier) GetStake() *types.Coin {
	if m != nil {
		return m.Stake
	}
	return nil
}

func (m *MsgStakeSupplier) GetServices() []*types1.SupplierServiceConfig {
	if m != nil {
		return m.Services
	}
	return nil
}

type MsgStakeSupplierResponse struct {
}

func (m *MsgStakeSupplierResponse) Reset()         { *m = MsgStakeSupplierResponse{} }
func (m *MsgStakeSupplierResponse) String() string { return proto.CompactTextString(m) }
func (*MsgStakeSupplierResponse) ProtoMessage()    {}
func (*MsgStakeSupplierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{3}
}
func (m *MsgStakeSupplierResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgStakeSupplierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgStakeSupplierResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgStakeSupplierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgStakeSupplierResponse.Merge(m, src)
}
func (m *MsgStakeSupplierResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgStakeSupplierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgStakeSupplierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgStakeSupplierResponse proto.InternalMessageInfo

type MsgUnstakeSupplier struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgUnstakeSupplier) Reset()         { *m = MsgUnstakeSupplier{} }
func (m *MsgUnstakeSupplier) String() string { return proto.CompactTextString(m) }
func (*MsgUnstakeSupplier) ProtoMessage()    {}
func (*MsgUnstakeSupplier) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{4}
}
func (m *MsgUnstakeSupplier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnstakeSupplier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnstakeSupplier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnstakeSupplier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnstakeSupplier.Merge(m, src)
}
func (m *MsgUnstakeSupplier) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnstakeSupplier) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnstakeSupplier.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnstakeSupplier proto.InternalMessageInfo

func (m *MsgUnstakeSupplier) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgUnstakeSupplierResponse struct {
}

func (m *MsgUnstakeSupplierResponse) Reset()         { *m = MsgUnstakeSupplierResponse{} }
func (m *MsgUnstakeSupplierResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUnstakeSupplierResponse) ProtoMessage()    {}
func (*MsgUnstakeSupplierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_63b974929807ef57, []int{5}
}
func (m *MsgUnstakeSupplierResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnstakeSupplierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnstakeSupplierResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnstakeSupplierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnstakeSupplierResponse.Merge(m, src)
}
func (m *MsgUnstakeSupplierResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnstakeSupplierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnstakeSupplierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnstakeSupplierResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgUpdateParams)(nil), "poktroll.supplier.MsgUpdateParams")
	proto.RegisterType((*MsgUpdateParamsResponse)(nil), "poktroll.supplier.MsgUpdateParamsResponse")
	proto.RegisterType((*MsgStakeSupplier)(nil), "poktroll.supplier.MsgStakeSupplier")
	proto.RegisterType((*MsgStakeSupplierResponse)(nil), "poktroll.supplier.MsgStakeSupplierResponse")
	proto.RegisterType((*MsgUnstakeSupplier)(nil), "poktroll.supplier.MsgUnstakeSupplier")
	proto.RegisterType((*MsgUnstakeSupplierResponse)(nil), "poktroll.supplier.MsgUnstakeSupplierResponse")
}

func init() { proto.RegisterFile("poktroll/supplier/tx.proto", fileDescriptor_63b974929807ef57) }

var fileDescriptor_63b974929807ef57 = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xb1, 0x6f, 0xd3, 0x4e,
	0x18, 0x8d, 0x1b, 0xb5, 0xbf, 0x5f, 0xae, 0x45, 0xa5, 0x56, 0xa5, 0x3a, 0x16, 0x98, 0xc8, 0x15,
	0x28, 0x0a, 0x8a, 0x4f, 0x09, 0x52, 0x87, 0x8a, 0x85, 0x74, 0x44, 0x91, 0x90, 0x23, 0x18, 0x18,
	0x40, 0x97, 0xe4, 0xb8, 0x58, 0x89, 0x7d, 0xd6, 0xdd, 0x25, 0xb4, 0x1b, 0x62, 0x64, 0xe2, 0xcf,
	0x60, 0xcc, 0xc0, 0xc0, 0xc6, 0x5a, 0xb6, 0x8a, 0x89, 0x09, 0xa1, 0x64, 0xc8, 0xbf, 0x81, 0xec,
	0xbb, 0x73, 0x88, 0x13, 0xd4, 0x4a, 0x2c, 0x89, 0x7d, 0xef, 0x7d, 0xdf, 0x7b, 0xef, 0xfb, 0x7c,
	0xc0, 0x8e, 0xe9, 0x50, 0x30, 0x3a, 0x1a, 0x41, 0x3e, 0x8e, 0xe3, 0x51, 0x80, 0x19, 0x14, 0xe7,
	0x5e, 0xcc, 0xa8, 0xa0, 0xe6, 0x81, 0xc6, 0x3c, 0x8d, 0xd9, 0x07, 0x28, 0x0c, 0x22, 0x0a, 0xd3,
	0x5f, 0xc9, 0xb2, 0x8f, 0x7a, 0x94, 0x87, 0x94, 0xc3, 0x90, 0x13, 0x38, 0x69, 0x24, 0x7f, 0x0a,
	0x28, 0x4b, 0xe0, 0x75, 0xfa, 0x06, 0xe5, 0x8b, 0x82, 0x0e, 0x09, 0x25, 0x54, 0x9e, 0x27, 0x4f,
	0xea, 0xd4, 0x51, 0x9d, 0xba, 0x88, 0x63, 0x38, 0x69, 0x74, 0xb1, 0x40, 0x0d, 0xd8, 0xa3, 0x41,
	0xa4, 0xf1, 0x75, 0xaf, 0x31, 0x62, 0x28, 0xd4, 0x5d, 0xef, 0x2e, 0xf1, 0x01, 0x62, 0xb8, 0x0f,
	0x39, 0x66, 0x93, 0xa0, 0x87, 0x25, 0xec, 0x7e, 0x35, 0xc0, 0x7e, 0x9b, 0x93, 0xe7, 0x71, 0x1f,
	0x09, 0xfc, 0x2c, 0x2d, 0x34, 0x4f, 0x40, 0x09, 0x8d, 0xc5, 0x80, 0xb2, 0x40, 0x5c, 0x58, 0x46,
	0xc5, 0xa8, 0x96, 0x5a, 0xd6, 0xf7, 0xcf, 0xf5, 0x43, 0xe5, 0xf6, 0x49, 0xbf, 0xcf, 0x30, 0xe7,
	0x1d, 0xc1, 0x82, 0x88, 0xf8, 0x4b, 0xaa, 0xf9, 0x18, 0xec, 0x48, 0x69, 0x6b, 0xab, 0x62, 0x54,
	0x77, 0x9b, 0x65, 0x6f, 0x6d, 0x56, 0x9e, 0x94, 0x68, 0x95, 0x2e, 0x7f, 0xde, 0x2b, 0x7c, 0x5a,
	0x4c, 0x6b, 0x86, 0xaf, 0x6a, 0x4e, 0x4f, 0xde, 0x2f, 0xa6, 0xb5, 0x65, 0xb7, 0x0f, 0x8b, 0x69,
	0xed, 0x38, 0xf3, 0x7e, 0xbe, 0x4c, 0x97, 0x73, 0xeb, 0x96, 0xc1, 0x51, 0xee, 0xc8, 0xc7, 0x3c,
	0xa6, 0x11, 0xc7, 0xee, 0x37, 0x03, 0xdc, 0x6e, 0x73, 0xd2, 0x11, 0x68, 0x88, 0x3b, 0xaa, 0xde,
	0x6c, 0x82, 0xff, 0x90, 0x4c, 0x70, 0x6d, 0x36, 0x4d, 0x34, 0x21, 0xd8, 0xe6, 0x49, 0x93, 0x2c,
	0x98, 0xa2, 0x27, 0x4b, 0xf1, 0xd4, 0x52, 0xbc, 0x33, 0x1a, 0x44, 0xbe, 0xe4, 0x99, 0x2d, 0xf0,
	0xbf, 0x9a, 0x33, 0xb7, 0x8a, 0x95, 0x62, 0x75, 0xb7, 0xf9, 0xe0, 0x8f, 0x61, 0xa4, 0x8b, 0xf0,
	0xb4, 0xa3, 0x8e, 0x24, 0x9e, 0xd1, 0xe8, 0x4d, 0x40, 0xfc, 0xac, 0xee, 0x74, 0x2f, 0x19, 0x88,
	0xb6, 0xe0, 0xda, 0xc0, 0xca, 0x47, 0xc9, 0x72, 0xbe, 0x00, 0x66, 0x32, 0x82, 0x88, 0xff, 0x6b,
	0xd0, 0x9c, 0xe6, 0x1d, 0x60, 0xaf, 0xf7, 0xd5, 0xaa, 0xcd, 0x2f, 0x5b, 0xa0, 0xd8, 0xe6, 0xc4,
	0x7c, 0x05, 0xf6, 0x56, 0x3e, 0x1f, 0x77, 0xc3, 0xda, 0x73, 0x1b, 0xb2, 0x6b, 0xd7, 0x73, 0xb4,
	0x8e, 0x89, 0xc0, 0xad, 0xd5, 0x0d, 0x1e, 0x6f, 0x2e, 0x5e, 0x21, 0xd9, 0x0f, 0x6f, 0x40, 0xca,
	0x24, 0x08, 0xd8, 0xcf, 0x4f, 0xef, 0xfe, 0x5f, 0x1c, 0xae, 0xd2, 0xec, 0xfa, 0x8d, 0x68, 0x5a,
	0xc8, 0xde, 0x7e, 0x97, 0x7c, 0xf3, 0xad, 0xa7, 0x97, 0x33, 0xc7, 0xb8, 0x9a, 0x39, 0xc6, 0xaf,
	0x99, 0x63, 0x7c, 0x9c, 0x3b, 0x85, 0xab, 0xb9, 0x53, 0xf8, 0x31, 0x77, 0x0a, 0x2f, 0x1b, 0x24,
	0x10, 0x83, 0x71, 0xd7, 0xeb, 0xd1, 0x10, 0x26, 0x9d, 0xeb, 0x11, 0x16, 0x6f, 0x29, 0x1b, 0xc2,
	0x4d, 0x57, 0x41, 0x5c, 0xc4, 0x98, 0x77, 0x77, 0xd2, 0x9b, 0xfc, 0xe8, 0x77, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x5d, 0x53, 0x46, 0x3c, 0xb6, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	StakeSupplier(ctx context.Context, in *MsgStakeSupplier, opts ...grpc.CallOption) (*MsgStakeSupplierResponse, error)
	UnstakeSupplier(ctx context.Context, in *MsgUnstakeSupplier, opts ...grpc.CallOption) (*MsgUnstakeSupplierResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Msg/UpdateParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StakeSupplier(ctx context.Context, in *MsgStakeSupplier, opts ...grpc.CallOption) (*MsgStakeSupplierResponse, error) {
	out := new(MsgStakeSupplierResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Msg/StakeSupplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UnstakeSupplier(ctx context.Context, in *MsgUnstakeSupplier, opts ...grpc.CallOption) (*MsgUnstakeSupplierResponse, error) {
	out := new(MsgUnstakeSupplierResponse)
	err := c.cc.Invoke(ctx, "/poktroll.supplier.Msg/UnstakeSupplier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	StakeSupplier(context.Context, *MsgStakeSupplier) (*MsgStakeSupplierResponse, error)
	UnstakeSupplier(context.Context, *MsgUnstakeSupplier) (*MsgUnstakeSupplierResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpdateParams(ctx context.Context, req *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (*UnimplementedMsgServer) StakeSupplier(ctx context.Context, req *MsgStakeSupplier) (*MsgStakeSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StakeSupplier not implemented")
}
func (*UnimplementedMsgServer) UnstakeSupplier(ctx context.Context, req *MsgUnstakeSupplier) (*MsgUnstakeSupplierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnstakeSupplier not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Msg/UpdateParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StakeSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStakeSupplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StakeSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Msg/StakeSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StakeSupplier(ctx, req.(*MsgStakeSupplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UnstakeSupplier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUnstakeSupplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UnstakeSupplier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/poktroll.supplier.Msg/UnstakeSupplier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UnstakeSupplier(ctx, req.(*MsgUnstakeSupplier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "poktroll.supplier.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "StakeSupplier",
			Handler:    _Msg_StakeSupplier_Handler,
		},
		{
			MethodName: "UnstakeSupplier",
			Handler:    _Msg_UnstakeSupplier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "poktroll/supplier/tx.proto",
}

func (m *MsgUpdateParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParams) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParams) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Authority) > 0 {
		i -= len(m.Authority)
		copy(dAtA[i:], m.Authority)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Authority)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgStakeSupplier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStakeSupplier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStakeSupplier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Stake != nil {
		{
			size, err := m.Stake.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgStakeSupplierResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgStakeSupplierResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgStakeSupplierResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUnstakeSupplier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnstakeSupplier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnstakeSupplier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUnstakeSupplierResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnstakeSupplierResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnstakeSupplierResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpdateParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Authority)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Params.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgStakeSupplier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Stake != nil {
		l = m.Stake.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgStakeSupplierResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUnstakeSupplier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUnstakeSupplierResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpdateParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authority", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Authority = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStakeSupplier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStakeSupplier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStakeSupplier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stake", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stake == nil {
				m.Stake = &types.Coin{}
			}
			if err := m.Stake.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &types1.SupplierServiceConfig{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgStakeSupplierResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgStakeSupplierResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgStakeSupplierResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnstakeSupplier) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnstakeSupplier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnstakeSupplier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUnstakeSupplierResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUnstakeSupplierResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnstakeSupplierResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
