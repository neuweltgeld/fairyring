// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/pep/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgSubmitEncryptedTx struct {
	Creator           string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Data              string `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	TargetBlockHeight uint64 `protobuf:"varint,3,opt,name=targetBlockHeight,proto3" json:"targetBlockHeight,omitempty"`
}

func (m *MsgSubmitEncryptedTx) Reset()         { *m = MsgSubmitEncryptedTx{} }
func (m *MsgSubmitEncryptedTx) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEncryptedTx) ProtoMessage()    {}
func (*MsgSubmitEncryptedTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6953e463911e1ec, []int{0}
}
func (m *MsgSubmitEncryptedTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEncryptedTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEncryptedTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEncryptedTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEncryptedTx.Merge(m, src)
}
func (m *MsgSubmitEncryptedTx) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEncryptedTx) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEncryptedTx.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEncryptedTx proto.InternalMessageInfo

func (m *MsgSubmitEncryptedTx) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitEncryptedTx) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MsgSubmitEncryptedTx) GetTargetBlockHeight() uint64 {
	if m != nil {
		return m.TargetBlockHeight
	}
	return 0
}

type MsgSubmitEncryptedTxResponse struct {
}

func (m *MsgSubmitEncryptedTxResponse) Reset()         { *m = MsgSubmitEncryptedTxResponse{} }
func (m *MsgSubmitEncryptedTxResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitEncryptedTxResponse) ProtoMessage()    {}
func (*MsgSubmitEncryptedTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6953e463911e1ec, []int{1}
}
func (m *MsgSubmitEncryptedTxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitEncryptedTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitEncryptedTxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitEncryptedTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitEncryptedTxResponse.Merge(m, src)
}
func (m *MsgSubmitEncryptedTxResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitEncryptedTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitEncryptedTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitEncryptedTxResponse proto.InternalMessageInfo

// this line is used by starport scaffolding # proto/tx/message
type MsgCreateAggregatedKeyShare struct {
	Creator   string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Height    uint64 `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Data      string `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	PublicKey string `protobuf:"bytes,4,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
}

func (m *MsgCreateAggregatedKeyShare) Reset()         { *m = MsgCreateAggregatedKeyShare{} }
func (m *MsgCreateAggregatedKeyShare) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAggregatedKeyShare) ProtoMessage()    {}
func (*MsgCreateAggregatedKeyShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6953e463911e1ec, []int{2}
}
func (m *MsgCreateAggregatedKeyShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAggregatedKeyShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAggregatedKeyShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAggregatedKeyShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAggregatedKeyShare.Merge(m, src)
}
func (m *MsgCreateAggregatedKeyShare) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAggregatedKeyShare) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAggregatedKeyShare.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAggregatedKeyShare proto.InternalMessageInfo

func (m *MsgCreateAggregatedKeyShare) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateAggregatedKeyShare) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *MsgCreateAggregatedKeyShare) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func (m *MsgCreateAggregatedKeyShare) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type MsgCreateAggregatedKeyShareResponse struct {
}

func (m *MsgCreateAggregatedKeyShareResponse) Reset()         { *m = MsgCreateAggregatedKeyShareResponse{} }
func (m *MsgCreateAggregatedKeyShareResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateAggregatedKeyShareResponse) ProtoMessage()    {}
func (*MsgCreateAggregatedKeyShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f6953e463911e1ec, []int{3}
}
func (m *MsgCreateAggregatedKeyShareResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateAggregatedKeyShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateAggregatedKeyShareResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateAggregatedKeyShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateAggregatedKeyShareResponse.Merge(m, src)
}
func (m *MsgCreateAggregatedKeyShareResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateAggregatedKeyShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateAggregatedKeyShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateAggregatedKeyShareResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgSubmitEncryptedTx)(nil), "fairyring.pep.MsgSubmitEncryptedTx")
	proto.RegisterType((*MsgSubmitEncryptedTxResponse)(nil), "fairyring.pep.MsgSubmitEncryptedTxResponse")
	proto.RegisterType((*MsgCreateAggregatedKeyShare)(nil), "fairyring.pep.MsgCreateAggregatedKeyShare")
	proto.RegisterType((*MsgCreateAggregatedKeyShareResponse)(nil), "fairyring.pep.MsgCreateAggregatedKeyShareResponse")
}

func init() { proto.RegisterFile("fairyring/pep/tx.proto", fileDescriptor_f6953e463911e1ec) }

var fileDescriptor_f6953e463911e1ec = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0x02, 0x31,
	0x10, 0xc6, 0x8d, 0x8a, 0xc5, 0x81, 0x1e, 0x0c, 0xad, 0x2c, 0x56, 0x82, 0xac, 0x14, 0xa4, 0x2d,
	0x2b, 0xd8, 0x27, 0xa8, 0xa5, 0x50, 0x10, 0x2f, 0xda, 0x53, 0x2f, 0x12, 0xd7, 0x69, 0x5c, 0xb4,
	0x6e, 0x48, 0x22, 0x98, 0xde, 0xfa, 0x06, 0x7d, 0xac, 0x1e, 0x3d, 0xf6, 0x58, 0xf4, 0xd2, 0xc7,
	0x28, 0xae, 0x5d, 0xad, 0xf8, 0x07, 0x6f, 0x3b, 0x33, 0xdf, 0xe6, 0xfb, 0x65, 0xf2, 0x41, 0xfe,
	0x85, 0x07, 0xca, 0xaa, 0x60, 0x24, 0xaa, 0x12, 0x65, 0xd5, 0x4c, 0x3c, 0xa9, 0x42, 0x13, 0xd2,
	0xd3, 0x55, 0xdf, 0x93, 0x28, 0x0b, 0x95, 0x4d, 0x19, 0x17, 0x42, 0xa1, 0xe0, 0x06, 0x7b, 0x9d,
	0x01, 0xda, 0x8e, 0xee, 0x73, 0x85, 0xcb, 0x1f, 0x5d, 0x05, 0x67, 0x4d, 0x2d, 0xda, 0xe3, 0xee,
	0x6b, 0x60, 0x1e, 0x46, 0xbe, 0xb2, 0xd2, 0x60, 0xef, 0x69, 0x42, 0x1d, 0x38, 0xf1, 0x15, 0x72,
	0x13, 0x2a, 0x87, 0x94, 0x48, 0x25, 0xdb, 0x8a, 0x4b, 0x4a, 0x21, 0xdd, 0xe3, 0x86, 0x3b, 0xc9,
	0xa8, 0x1d, 0x7d, 0xd3, 0x1b, 0xc8, 0x19, 0xae, 0x04, 0x9a, 0xfa, 0x30, 0xf4, 0x07, 0x8f, 0x18,
	0x88, 0xbe, 0x71, 0x52, 0x25, 0x52, 0x49, 0xb7, 0xb6, 0x07, 0x2e, 0x83, 0xe2, 0x2e, 0xcf, 0x16,
	0x6a, 0x19, 0x8e, 0x34, 0xba, 0xef, 0x04, 0x2e, 0x9a, 0x5a, 0xdc, 0x2f, 0x0c, 0xf1, 0x6e, 0xc5,
	0xde, 0x40, 0xdb, 0x5e, 0x90, 0x1f, 0x60, 0xcb, 0x43, 0xa6, 0xbf, 0x34, 0x4f, 0x46, 0xe6, 0x7f,
	0xd5, 0x8a, 0x39, 0xf5, 0x8f, 0xb9, 0x08, 0x59, 0x39, 0xee, 0x0e, 0x03, 0xbf, 0x81, 0xd6, 0x49,
	0x47, 0x83, 0x75, 0xc3, 0xbd, 0x84, 0xf2, 0x01, 0x84, 0x18, 0xb5, 0xf6, 0x43, 0x20, 0xd5, 0xd4,
	0x82, 0x22, 0xe4, 0xb6, 0x77, 0x58, 0xf6, 0x36, 0x5e, 0xc5, 0xdb, 0x75, 0xe9, 0xc2, 0xf5, 0x11,
	0xa2, 0xd8, 0x8e, 0xbe, 0x81, 0xb3, 0x77, 0x2b, 0x57, 0xdb, 0x07, 0xed, 0xd3, 0x16, 0x6a, 0xc7,
	0x6b, 0x63, 0xef, 0x7a, 0xf5, 0x73, 0xc6, 0xc8, 0x74, 0xc6, 0xc8, 0xf7, 0x8c, 0x91, 0x8f, 0x39,
	0x4b, 0x4c, 0xe7, 0x2c, 0xf1, 0x35, 0x67, 0x89, 0xe7, 0xf3, 0x75, 0xda, 0x26, 0xcb, 0x58, 0x5a,
	0x89, 0xba, 0x9b, 0x89, 0x12, 0x76, 0xfb, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x0b, 0x1d, 0x25,
	0xb4, 0x02, 0x00, 0x00,
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
	SubmitEncryptedTx(ctx context.Context, in *MsgSubmitEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateAggregatedKeyShare(ctx context.Context, in *MsgCreateAggregatedKeyShare, opts ...grpc.CallOption) (*MsgCreateAggregatedKeyShareResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitEncryptedTx(ctx context.Context, in *MsgSubmitEncryptedTx, opts ...grpc.CallOption) (*MsgSubmitEncryptedTxResponse, error) {
	out := new(MsgSubmitEncryptedTxResponse)
	err := c.cc.Invoke(ctx, "/fairyring.pep.Msg/SubmitEncryptedTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateAggregatedKeyShare(ctx context.Context, in *MsgCreateAggregatedKeyShare, opts ...grpc.CallOption) (*MsgCreateAggregatedKeyShareResponse, error) {
	out := new(MsgCreateAggregatedKeyShareResponse)
	err := c.cc.Invoke(ctx, "/fairyring.pep.Msg/CreateAggregatedKeyShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitEncryptedTx(context.Context, *MsgSubmitEncryptedTx) (*MsgSubmitEncryptedTxResponse, error)
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateAggregatedKeyShare(context.Context, *MsgCreateAggregatedKeyShare) (*MsgCreateAggregatedKeyShareResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitEncryptedTx(ctx context.Context, req *MsgSubmitEncryptedTx) (*MsgSubmitEncryptedTxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitEncryptedTx not implemented")
}
func (*UnimplementedMsgServer) CreateAggregatedKeyShare(ctx context.Context, req *MsgCreateAggregatedKeyShare) (*MsgCreateAggregatedKeyShareResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAggregatedKeyShare not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitEncryptedTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitEncryptedTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitEncryptedTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairyring.pep.Msg/SubmitEncryptedTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitEncryptedTx(ctx, req.(*MsgSubmitEncryptedTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateAggregatedKeyShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateAggregatedKeyShare)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateAggregatedKeyShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fairyring.pep.Msg/CreateAggregatedKeyShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateAggregatedKeyShare(ctx, req.(*MsgCreateAggregatedKeyShare))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fairyring.pep.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitEncryptedTx",
			Handler:    _Msg_SubmitEncryptedTx_Handler,
		},
		{
			MethodName: "CreateAggregatedKeyShare",
			Handler:    _Msg_CreateAggregatedKeyShare_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fairyring/pep/tx.proto",
}

func (m *MsgSubmitEncryptedTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEncryptedTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEncryptedTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TargetBlockHeight != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TargetBlockHeight))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitEncryptedTxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitEncryptedTxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitEncryptedTxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateAggregatedKeyShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAggregatedKeyShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAggregatedKeyShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PublicKey) > 0 {
		i -= len(m.PublicKey)
		copy(dAtA[i:], m.PublicKey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.PublicKey)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Height != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateAggregatedKeyShareResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateAggregatedKeyShareResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateAggregatedKeyShareResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgSubmitEncryptedTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.TargetBlockHeight != 0 {
		n += 1 + sovTx(uint64(m.TargetBlockHeight))
	}
	return n
}

func (m *MsgSubmitEncryptedTxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateAggregatedKeyShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTx(uint64(m.Height))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.PublicKey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateAggregatedKeyShareResponse) Size() (n int) {
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
func (m *MsgSubmitEncryptedTx) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitEncryptedTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEncryptedTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TargetBlockHeight", wireType)
			}
			m.TargetBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TargetBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgSubmitEncryptedTxResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgSubmitEncryptedTxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitEncryptedTxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgCreateAggregatedKeyShare) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateAggregatedKeyShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAggregatedKeyShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
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
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
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
			m.PublicKey = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateAggregatedKeyShareResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateAggregatedKeyShareResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateAggregatedKeyShareResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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