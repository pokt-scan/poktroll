// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: poktroll/tokenomics/event.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/pokt-network/poktroll/x/proof/types"
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

// EventClaimExpired is an event emitted during settlement whenever a claim requiring
// an on-chain proof doesn't have one. The claim cannot be settled, leading to that work
// never being rewarded.
type EventClaimExpired struct {
	Claim        *types.Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	ComputeUnits uint64       `protobuf:"varint,2,opt,name=compute_units,json=computeUnits,proto3" json:"compute_units,omitempty"`
}

func (m *EventClaimExpired) Reset()         { *m = EventClaimExpired{} }
func (m *EventClaimExpired) String() string { return proto.CompactTextString(m) }
func (*EventClaimExpired) ProtoMessage()    {}
func (*EventClaimExpired) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{0}
}
func (m *EventClaimExpired) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimExpired) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimExpired.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimExpired) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimExpired.Merge(m, src)
}
func (m *EventClaimExpired) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimExpired) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimExpired.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimExpired proto.InternalMessageInfo

func (m *EventClaimExpired) GetClaim() *types.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimExpired) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

// EventClaimSettled is an event emitted whenever a claim is settled.
// The proof_required determines whether the claim requires a proof that has been submitted or not
type EventClaimSettled struct {
	Claim         *types.Claim `protobuf:"bytes,1,opt,name=claim,proto3" json:"claim,omitempty"`
	ComputeUnits  uint64       `protobuf:"varint,2,opt,name=compute_units,json=computeUnits,proto3" json:"compute_units,omitempty"`
	ProofRequired bool         `protobuf:"varint,3,opt,name=proof_required,json=proofRequired,proto3" json:"proof_required,omitempty"`
}

func (m *EventClaimSettled) Reset()         { *m = EventClaimSettled{} }
func (m *EventClaimSettled) String() string { return proto.CompactTextString(m) }
func (*EventClaimSettled) ProtoMessage()    {}
func (*EventClaimSettled) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{1}
}
func (m *EventClaimSettled) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventClaimSettled) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventClaimSettled.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventClaimSettled) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventClaimSettled.Merge(m, src)
}
func (m *EventClaimSettled) XXX_Size() int {
	return m.Size()
}
func (m *EventClaimSettled) XXX_DiscardUnknown() {
	xxx_messageInfo_EventClaimSettled.DiscardUnknown(m)
}

var xxx_messageInfo_EventClaimSettled proto.InternalMessageInfo

func (m *EventClaimSettled) GetClaim() *types.Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

func (m *EventClaimSettled) GetComputeUnits() uint64 {
	if m != nil {
		return m.ComputeUnits
	}
	return 0
}

func (m *EventClaimSettled) GetProofRequired() bool {
	if m != nil {
		return m.ProofRequired
	}
	return false
}

// EventRelayMiningDifficultyUpdated is an event emitted whenever the relay mining difficulty is updated
// for a given service.
type EventRelayMiningDifficultyUpdated struct {
	ServiceId                string `protobuf:"bytes,1,opt,name=service_id,json=serviceId,proto3" json:"service_id,omitempty"`
	PrevTargetHashHexEncoded string `protobuf:"bytes,2,opt,name=prev_target_hash_hex_encoded,json=prevTargetHashHexEncoded,proto3" json:"prev_target_hash_hex_encoded,omitempty"`
	NewTargetHashHexEncoded  string `protobuf:"bytes,3,opt,name=new_target_hash_hex_encoded,json=newTargetHashHexEncoded,proto3" json:"new_target_hash_hex_encoded,omitempty"`
	PrevNumRelaysEma         uint64 `protobuf:"varint,4,opt,name=prev_num_relays_ema,json=prevNumRelaysEma,proto3" json:"prev_num_relays_ema,omitempty"`
	NewNumRelaysEma          uint64 `protobuf:"varint,5,opt,name=new_num_relays_ema,json=newNumRelaysEma,proto3" json:"new_num_relays_ema,omitempty"`
}

func (m *EventRelayMiningDifficultyUpdated) Reset()         { *m = EventRelayMiningDifficultyUpdated{} }
func (m *EventRelayMiningDifficultyUpdated) String() string { return proto.CompactTextString(m) }
func (*EventRelayMiningDifficultyUpdated) ProtoMessage()    {}
func (*EventRelayMiningDifficultyUpdated) Descriptor() ([]byte, []int) {
	return fileDescriptor_a78874bbf91a58c7, []int{2}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventRelayMiningDifficultyUpdated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.Merge(m, src)
}
func (m *EventRelayMiningDifficultyUpdated) XXX_Size() int {
	return m.Size()
}
func (m *EventRelayMiningDifficultyUpdated) XXX_DiscardUnknown() {
	xxx_messageInfo_EventRelayMiningDifficultyUpdated.DiscardUnknown(m)
}

var xxx_messageInfo_EventRelayMiningDifficultyUpdated proto.InternalMessageInfo

func (m *EventRelayMiningDifficultyUpdated) GetServiceId() string {
	if m != nil {
		return m.ServiceId
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevTargetHashHexEncoded() string {
	if m != nil {
		return m.PrevTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetNewTargetHashHexEncoded() string {
	if m != nil {
		return m.NewTargetHashHexEncoded
	}
	return ""
}

func (m *EventRelayMiningDifficultyUpdated) GetPrevNumRelaysEma() uint64 {
	if m != nil {
		return m.PrevNumRelaysEma
	}
	return 0
}

func (m *EventRelayMiningDifficultyUpdated) GetNewNumRelaysEma() uint64 {
	if m != nil {
		return m.NewNumRelaysEma
	}
	return 0
}

func init() {
	proto.RegisterType((*EventClaimExpired)(nil), "poktroll.tokenomics.EventClaimExpired")
	proto.RegisterType((*EventClaimSettled)(nil), "poktroll.tokenomics.EventClaimSettled")
	proto.RegisterType((*EventRelayMiningDifficultyUpdated)(nil), "poktroll.tokenomics.EventRelayMiningDifficultyUpdated")
}

func init() { proto.RegisterFile("poktroll/tokenomics/event.proto", fileDescriptor_a78874bbf91a58c7) }

var fileDescriptor_a78874bbf91a58c7 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xcb, 0x8e, 0xd3, 0x30,
	0x14, 0x86, 0xeb, 0xb9, 0x20, 0x6a, 0x18, 0x2e, 0x1e, 0x21, 0xa2, 0x01, 0x42, 0x29, 0x42, 0xaa,
	0x34, 0x9a, 0x44, 0x62, 0xb6, 0x88, 0x05, 0x10, 0x69, 0x58, 0x0c, 0x8b, 0xc0, 0x6c, 0xd8, 0x58,
	0x9e, 0xe4, 0xb4, 0xb1, 0x1a, 0x5f, 0xb0, 0x9d, 0x26, 0x7d, 0x07, 0x16, 0x3c, 0x00, 0x0f, 0xc4,
	0x72, 0x96, 0x2c, 0x51, 0xfb, 0x22, 0xc8, 0x6e, 0x29, 0x15, 0x82, 0x25, 0xdb, 0xff, 0x7c, 0xff,
	0xf9, 0xcf, 0xb1, 0x0f, 0x7e, 0xac, 0xd5, 0xd4, 0x19, 0x55, 0xd7, 0xa9, 0x53, 0x53, 0x90, 0x4a,
	0xf0, 0xc2, 0xa6, 0x30, 0x03, 0xe9, 0x12, 0x6d, 0x94, 0x53, 0xe4, 0xf0, 0x17, 0x90, 0xfc, 0x06,
	0x8e, 0x8e, 0x36, 0x2e, 0x6d, 0x94, 0x1a, 0xa7, 0x45, 0xcd, 0xb8, 0x58, 0x19, 0x86, 0x80, 0xef,
	0x66, 0xde, 0xff, 0xda, 0x6b, 0x59, 0xa7, 0xb9, 0x81, 0x92, 0x1c, 0xe3, 0xfd, 0xc0, 0x44, 0x68,
	0x80, 0x46, 0x37, 0x9e, 0xdf, 0x4b, 0x36, 0x5d, 0x43, 0x83, 0x24, 0xc0, 0xf9, 0x8a, 0x21, 0x4f,
	0xf1, 0x41, 0xa1, 0x84, 0x6e, 0x1c, 0xd0, 0x46, 0x72, 0x67, 0xa3, 0x9d, 0x01, 0x1a, 0xed, 0xe5,
	0x37, 0xd7, 0xe2, 0x85, 0xd7, 0x86, 0x9f, 0xd1, 0x76, 0xce, 0x7b, 0x70, 0xae, 0xfe, 0x1f, 0x39,
	0xe4, 0x19, 0xbe, 0x15, 0xac, 0xd4, 0xc0, 0xa7, 0xc6, 0xef, 0x12, 0xed, 0x0e, 0xd0, 0xe8, 0x7a,
	0x7e, 0x10, 0xd4, 0x7c, 0x2d, 0x0e, 0xbf, 0xee, 0xe0, 0x27, 0x61, 0x9c, 0x1c, 0x6a, 0x36, 0x3f,
	0xe7, 0x92, 0xcb, 0xc9, 0x1b, 0x3e, 0x1e, 0xf3, 0xa2, 0xa9, 0xdd, 0xfc, 0x42, 0x97, 0xcc, 0x41,
	0x49, 0x1e, 0x61, 0x6c, 0xc1, 0xcc, 0x78, 0x01, 0x94, 0x97, 0x61, 0xc6, 0x7e, 0xde, 0x5f, 0x2b,
	0x6f, 0x4b, 0xf2, 0x12, 0x3f, 0xd4, 0x06, 0x66, 0xd4, 0x31, 0x33, 0x01, 0x47, 0x2b, 0x66, 0x2b,
	0x5a, 0x41, 0x47, 0x41, 0x16, 0xaa, 0x84, 0x32, 0xcc, 0xd7, 0xcf, 0x23, 0xcf, 0x7c, 0x08, 0xc8,
	0x19, 0xb3, 0xd5, 0x19, 0x74, 0xd9, 0xaa, 0x4e, 0x5e, 0xe0, 0x07, 0x12, 0xda, 0x7f, 0xda, 0x77,
	0x83, 0xfd, 0xbe, 0x84, 0xf6, 0xaf, 0xee, 0x13, 0x7c, 0x18, 0xd2, 0x65, 0x23, 0xa8, 0xf1, 0x5b,
	0x58, 0x0a, 0x82, 0x45, 0x7b, 0xe1, 0x51, 0xee, 0xf8, 0xd2, 0xbb, 0x46, 0x84, 0xf5, 0x6c, 0x26,
	0x18, 0x39, 0xc6, 0xc4, 0x87, 0xfd, 0x41, 0xef, 0x07, 0xfa, 0xb6, 0x84, 0x76, 0x1b, 0x7e, 0x75,
	0xfe, 0x6d, 0x11, 0xa3, 0xab, 0x45, 0x8c, 0x7e, 0x2c, 0x62, 0xf4, 0x65, 0x19, 0xf7, 0xae, 0x96,
	0x71, 0xef, 0xfb, 0x32, 0xee, 0x7d, 0x3c, 0x9d, 0x70, 0x57, 0x35, 0x97, 0x49, 0xa1, 0x44, 0xea,
	0x3f, 0xeb, 0x44, 0x82, 0x6b, 0x95, 0x99, 0xa6, 0x9b, 0x13, 0xeb, 0xb6, 0x4f, 0xd3, 0xcd, 0x35,
	0xd8, 0xcb, 0x6b, 0xe1, 0xd4, 0x4e, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0xee, 0x9a, 0x02, 0xb4,
	0xbe, 0x02, 0x00, 0x00,
}

func (m *EventClaimExpired) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimExpired) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimExpired) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventClaimSettled) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventClaimSettled) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventClaimSettled) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProofRequired {
		i--
		if m.ProofRequired {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.ComputeUnits != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.ComputeUnits))
		i--
		dAtA[i] = 0x10
	}
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventRelayMiningDifficultyUpdated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventRelayMiningDifficultyUpdated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventRelayMiningDifficultyUpdated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NewNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.NewNumRelaysEma))
		i--
		dAtA[i] = 0x28
	}
	if m.PrevNumRelaysEma != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.PrevNumRelaysEma))
		i--
		dAtA[i] = 0x20
	}
	if len(m.NewTargetHashHexEncoded) > 0 {
		i -= len(m.NewTargetHashHexEncoded)
		copy(dAtA[i:], m.NewTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.NewTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PrevTargetHashHexEncoded) > 0 {
		i -= len(m.PrevTargetHashHexEncoded)
		copy(dAtA[i:], m.PrevTargetHashHexEncoded)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PrevTargetHashHexEncoded)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceId) > 0 {
		i -= len(m.ServiceId)
		copy(dAtA[i:], m.ServiceId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.ServiceId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventClaimExpired) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.ComputeUnits))
	}
	return n
}

func (m *EventClaimSettled) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.ComputeUnits != 0 {
		n += 1 + sovEvent(uint64(m.ComputeUnits))
	}
	if m.ProofRequired {
		n += 2
	}
	return n
}

func (m *EventRelayMiningDifficultyUpdated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PrevTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.NewTargetHashHexEncoded)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.PrevNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.PrevNumRelaysEma))
	}
	if m.NewNumRelaysEma != 0 {
		n += 1 + sovEvent(uint64(m.NewNumRelaysEma))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventClaimExpired) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimExpired: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimExpired: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventClaimSettled) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventClaimSettled: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventClaimSettled: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &types.Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ComputeUnits", wireType)
			}
			m.ComputeUnits = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ComputeUnits |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofRequired", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ProofRequired = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *EventRelayMiningDifficultyUpdated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventRelayMiningDifficultyUpdated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrevTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewTargetHashHexEncoded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NewTargetHashHexEncoded = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrevNumRelaysEma", wireType)
			}
			m.PrevNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PrevNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewNumRelaysEma", wireType)
			}
			m.NewNumRelaysEma = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NewNumRelaysEma |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
