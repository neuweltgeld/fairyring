// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fairyring/fairyring/key_share.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type KeyShare struct {
	Validator           string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	BlockHeight         uint64 `protobuf:"varint,2,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	KeyShare            string `protobuf:"bytes,3,opt,name=keyShare,proto3" json:"keyShare,omitempty"`
	ReceivedTimestamp   uint64 `protobuf:"varint,4,opt,name=receivedTimestamp,proto3" json:"receivedTimestamp,omitempty"`
	ReceivedBlockHeight uint64 `protobuf:"varint,5,opt,name=receivedBlockHeight,proto3" json:"receivedBlockHeight,omitempty"`
}

func (m *KeyShare) Reset()         { *m = KeyShare{} }
func (m *KeyShare) String() string { return proto.CompactTextString(m) }
func (*KeyShare) ProtoMessage()    {}
func (*KeyShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e52ec66195b00f3, []int{0}
}
func (m *KeyShare) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *KeyShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_KeyShare.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *KeyShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyShare.Merge(m, src)
}
func (m *KeyShare) XXX_Size() int {
	return m.Size()
}
func (m *KeyShare) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyShare.DiscardUnknown(m)
}

var xxx_messageInfo_KeyShare proto.InternalMessageInfo

func (m *KeyShare) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *KeyShare) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *KeyShare) GetKeyShare() string {
	if m != nil {
		return m.KeyShare
	}
	return ""
}

func (m *KeyShare) GetReceivedTimestamp() uint64 {
	if m != nil {
		return m.ReceivedTimestamp
	}
	return 0
}

func (m *KeyShare) GetReceivedBlockHeight() uint64 {
	if m != nil {
		return m.ReceivedBlockHeight
	}
	return 0
}

func init() {
	proto.RegisterType((*KeyShare)(nil), "fairyring.fairyring.KeyShare")
}

func init() {
	proto.RegisterFile("fairyring/fairyring/key_share.proto", fileDescriptor_1e52ec66195b00f3)
}

var fileDescriptor_1e52ec66195b00f3 = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x4b, 0xcc, 0x2c,
	0xaa, 0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0x47, 0xb0, 0xb2, 0x53, 0x2b, 0xe3, 0x8b, 0x33, 0x12, 0x8b,
	0x52, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x84, 0xe1, 0x52, 0x7a, 0x70, 0x96, 0xd2, 0x31,
	0x46, 0x2e, 0x0e, 0xef, 0xd4, 0xca, 0x60, 0x90, 0x3a, 0x21, 0x19, 0x2e, 0xce, 0xb2, 0xc4, 0x9c,
	0xcc, 0x94, 0xc4, 0x92, 0xfc, 0x22, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x90,
	0x02, 0x17, 0x77, 0x52, 0x4e, 0x7e, 0x72, 0xb6, 0x47, 0x6a, 0x66, 0x7a, 0x46, 0x89, 0x04, 0x93,
	0x02, 0xa3, 0x06, 0x4b, 0x10, 0xb2, 0x90, 0x90, 0x14, 0x17, 0x47, 0x36, 0xd4, 0x2c, 0x09, 0x66,
	0xb0, 0x76, 0x38, 0x5f, 0x48, 0x87, 0x4b, 0xb0, 0x28, 0x35, 0x39, 0x35, 0xb3, 0x2c, 0x35, 0x25,
	0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x05, 0x6c, 0x06, 0xa6, 0x84, 0x90,
	0x01, 0x97, 0x30, 0x4c, 0xd0, 0x09, 0xc9, 0x4e, 0x56, 0xb0, 0x7a, 0x6c, 0x52, 0x4e, 0xa6, 0x27,
	0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c,
	0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x25, 0x8d, 0x08, 0x92, 0x0a, 0xa4, 0xe0,
	0x29, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x87, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x2a, 0x16, 0x13, 0x3b, 0x42, 0x01, 0x00, 0x00,
}

func (m *KeyShare) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *KeyShare) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *KeyShare) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ReceivedBlockHeight != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.ReceivedBlockHeight))
		i--
		dAtA[i] = 0x28
	}
	if m.ReceivedTimestamp != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.ReceivedTimestamp))
		i--
		dAtA[i] = 0x20
	}
	if len(m.KeyShare) > 0 {
		i -= len(m.KeyShare)
		copy(dAtA[i:], m.KeyShare)
		i = encodeVarintKeyShare(dAtA, i, uint64(len(m.KeyShare)))
		i--
		dAtA[i] = 0x1a
	}
	if m.BlockHeight != 0 {
		i = encodeVarintKeyShare(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintKeyShare(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintKeyShare(dAtA []byte, offset int, v uint64) int {
	offset -= sovKeyShare(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *KeyShare) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovKeyShare(uint64(l))
	}
	if m.BlockHeight != 0 {
		n += 1 + sovKeyShare(uint64(m.BlockHeight))
	}
	l = len(m.KeyShare)
	if l > 0 {
		n += 1 + l + sovKeyShare(uint64(l))
	}
	if m.ReceivedTimestamp != 0 {
		n += 1 + sovKeyShare(uint64(m.ReceivedTimestamp))
	}
	if m.ReceivedBlockHeight != 0 {
		n += 1 + sovKeyShare(uint64(m.ReceivedBlockHeight))
	}
	return n
}

func sovKeyShare(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozKeyShare(x uint64) (n int) {
	return sovKeyShare(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *KeyShare) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKeyShare
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
			return fmt.Errorf("proto: KeyShare: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: KeyShare: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
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
				return ErrInvalidLengthKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field KeyShare", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
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
				return ErrInvalidLengthKeyShare
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthKeyShare
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.KeyShare = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedTimestamp", wireType)
			}
			m.ReceivedTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedBlockHeight", wireType)
			}
			m.ReceivedBlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKeyShare
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedBlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipKeyShare(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthKeyShare
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
func skipKeyShare(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKeyShare
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
					return 0, ErrIntOverflowKeyShare
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
					return 0, ErrIntOverflowKeyShare
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
				return 0, ErrInvalidLengthKeyShare
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupKeyShare
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthKeyShare
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthKeyShare        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKeyShare          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupKeyShare = fmt.Errorf("proto: unexpected end of group")
)