// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/proof/claim.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/session/types"
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

// Claim is the serialized object stored on-chain for claims pending to be proven
type Claim struct {
	SupplierAddress string `protobuf:"bytes,1,opt,name=supplier_address,json=supplierAddress,proto3" json:"supplier_address,omitempty"`
	// The session header of the session that this claim is for.
	SessionHeader *types.SessionHeader `protobuf:"bytes,2,opt,name=session_header,json=sessionHeader,proto3" json:"session_header,omitempty"`
	// Root hash returned from smt.SMST#Root().
	RootHash []byte `protobuf:"bytes,3,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_51f775f805f1488c, []int{0}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetSupplierAddress() string {
	if m != nil {
		return m.SupplierAddress
	}
	return ""
}

func (m *Claim) GetSessionHeader() *types.SessionHeader {
	if m != nil {
		return m.SessionHeader
	}
	return nil
}

func (m *Claim) GetRootHash() []byte {
	if m != nil {
		return m.RootHash
	}
	return nil
}

func init() {
	proto.RegisterType((*Claim)(nil), "poktroll.proof.Claim")
}

func init() { proto.RegisterFile("poktroll/proof/claim.proto", fileDescriptor_51f775f805f1488c) }

var fileDescriptor_51f775f805f1488c = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x59, 0x8d, 0xc6, 0xa2, 0x56, 0x43, 0x3c, 0x20, 0x26, 0x2b, 0xf1, 0xc4, 0xa5, 0x90,
	0xe8, 0x13, 0xd8, 0x26, 0x86, 0x33, 0xbd, 0x79, 0x21, 0x14, 0xd6, 0xb2, 0x29, 0x30, 0x9b, 0x9d,
	0x6d, 0xd4, 0xb7, 0xf0, 0x55, 0x4c, 0x7c, 0x08, 0x8f, 0x8d, 0x27, 0x8f, 0x06, 0x5e, 0xc4, 0xc0,
	0x02, 0xe9, 0x69, 0x32, 0xf3, 0xcd, 0x7e, 0xb3, 0xf9, 0x4d, 0x47, 0xc0, 0x46, 0x49, 0x28, 0x8a,
	0x40, 0x48, 0x80, 0x97, 0x20, 0x2d, 0x12, 0x5e, 0xfa, 0x42, 0x82, 0x02, 0x6b, 0x3a, 0x30, 0xbf,
	0x63, 0xce, 0x75, 0x0a, 0x58, 0x02, 0xc6, 0x1d, 0x0d, 0x74, 0xa3, 0x57, 0x1d, 0x3a, 0x6a, 0x90,
	0x21, 0x72, 0xa8, 0x86, 0xaa, 0xf9, 0xdd, 0x27, 0x31, 0x8f, 0x16, 0xad, 0xda, 0x5a, 0x98, 0x97,
	0xb8, 0x15, 0xa2, 0xe0, 0x4c, 0xc6, 0x49, 0x96, 0x49, 0x86, 0x68, 0x13, 0x97, 0x78, 0x93, 0xb9,
	0xfd, 0xf3, 0x35, 0xbb, 0xea, 0xad, 0x8f, 0x9a, 0x2c, 0x95, 0xe4, 0xd5, 0x3a, 0xba, 0x18, 0x5e,
	0xf4, 0x63, 0xeb, 0xc9, 0x9c, 0xf6, 0xfe, 0x38, 0x67, 0x49, 0xc6, 0xa4, 0x7d, 0xe0, 0x12, 0xef,
	0xf4, 0xfe, 0xd6, 0x1f, 0xbf, 0x3c, 0xdc, 0x5f, 0xea, 0x1a, 0x76, 0x6b, 0xd1, 0x39, 0xee, 0xb7,
	0xd6, 0x8d, 0x39, 0x91, 0x00, 0x2a, 0xce, 0x13, 0xcc, 0xed, 0x43, 0x97, 0x78, 0x67, 0xd1, 0x49,
	0x3b, 0x08, 0x13, 0xcc, 0xe7, 0xe1, 0x77, 0x4d, 0xc9, 0xae, 0xa6, 0xe4, 0xaf, 0xa6, 0xe4, 0xa3,
	0xa1, 0xc6, 0xae, 0xa1, 0xc6, 0x6f, 0x43, 0x8d, 0x67, 0x7f, 0xcd, 0x55, 0xbe, 0x5d, 0xf9, 0x29,
	0x94, 0x41, 0x7b, 0x70, 0x56, 0x31, 0xf5, 0x0a, 0x72, 0x13, 0x8c, 0x29, 0xbc, 0xf5, 0x71, 0xaa,
	0x77, 0xc1, 0x70, 0x75, 0xdc, 0x85, 0xf0, 0xf0, 0x1f, 0x00, 0x00, 0xff, 0xff, 0x97, 0x71, 0x55,
	0xe7, 0x6d, 0x01, 0x00, 0x00,
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RootHash) > 0 {
		i -= len(m.RootHash)
		copy(dAtA[i:], m.RootHash)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.RootHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SessionHeader != nil {
		{
			size, err := m.SessionHeader.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaim(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.SupplierAddress) > 0 {
		i -= len(m.SupplierAddress)
		copy(dAtA[i:], m.SupplierAddress)
		i = encodeVarintClaim(dAtA, i, uint64(len(m.SupplierAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaim(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaim(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SupplierAddress)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	if m.SessionHeader != nil {
		l = m.SessionHeader.Size()
		n += 1 + l + sovClaim(uint64(l))
	}
	l = len(m.RootHash)
	if l > 0 {
		n += 1 + l + sovClaim(uint64(l))
	}
	return n
}

func sovClaim(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaim(x uint64) (n int) {
	return sovClaim(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaim
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
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SupplierAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SupplierAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SessionHeader", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
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
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SessionHeader == nil {
				m.SessionHeader = &types.SessionHeader{}
			}
			if err := m.SessionHeader.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaim
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthClaim
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthClaim
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootHash = append(m.RootHash[:0], dAtA[iNdEx:postIndex]...)
			if m.RootHash == nil {
				m.RootHash = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaim(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaim
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
func skipClaim(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
					return 0, ErrIntOverflowClaim
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
				return 0, ErrInvalidLengthClaim
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaim
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaim
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaim        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaim          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaim = fmt.Errorf("proto: unexpected end of group")
)
