// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: consensus/consensus_queue.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// message for storing the queued signed message in the internal queue
type QueuedSignedMessage struct {
	Id       uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg      *types.Any  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	SignData []*SignData `protobuf:"bytes,3,rep,name=signData,proto3" json:"signData,omitempty"`
}

func (m *QueuedSignedMessage) Reset()         { *m = QueuedSignedMessage{} }
func (m *QueuedSignedMessage) String() string { return proto.CompactTextString(m) }
func (*QueuedSignedMessage) ProtoMessage()    {}
func (*QueuedSignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cd502dda0bddc7c, []int{0}
}
func (m *QueuedSignedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueuedSignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueuedSignedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueuedSignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedSignedMessage.Merge(m, src)
}
func (m *QueuedSignedMessage) XXX_Size() int {
	return m.Size()
}
func (m *QueuedSignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedSignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedSignedMessage proto.InternalMessageInfo

func (m *QueuedSignedMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueuedSignedMessage) GetMsg() *types.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *QueuedSignedMessage) GetSignData() []*SignData {
	if m != nil {
		return m.SignData
	}
	return nil
}

type SignData struct {
	ValAddress string `protobuf:"bytes,1,opt,name=valAddress,proto3" json:"valAddress,omitempty"`
	PubKey     []byte `protobuf:"bytes,2,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	Signature  []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *SignData) Reset()         { *m = SignData{} }
func (m *SignData) String() string { return proto.CompactTextString(m) }
func (*SignData) ProtoMessage()    {}
func (*SignData) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cd502dda0bddc7c, []int{1}
}
func (m *SignData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignData.Merge(m, src)
}
func (m *SignData) XXX_Size() int {
	return m.Size()
}
func (m *SignData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignData.DiscardUnknown(m)
}

var xxx_messageInfo_SignData proto.InternalMessageInfo

func (m *SignData) GetValAddress() string {
	if m != nil {
		return m.ValAddress
	}
	return ""
}

func (m *SignData) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *SignData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*QueuedSignedMessage)(nil), "volumefi.paloma.consensus.QueuedSignedMessage")
	proto.RegisterType((*SignData)(nil), "volumefi.paloma.consensus.SignData")
}

func init() { proto.RegisterFile("consensus/consensus_queue.proto", fileDescriptor_4cd502dda0bddc7c) }

var fileDescriptor_4cd502dda0bddc7c = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4e, 0x02, 0x31,
	0x14, 0x86, 0x29, 0x63, 0x08, 0x14, 0xe3, 0xa2, 0x1a, 0x33, 0x18, 0x53, 0x09, 0x26, 0x86, 0x55,
	0x9b, 0xe0, 0x01, 0x0c, 0xc6, 0x8d, 0x31, 0x2e, 0x1c, 0x76, 0x6e, 0xb4, 0x43, 0x1f, 0xa5, 0xc9,
	0x4c, 0x3b, 0xd2, 0x29, 0x71, 0x2e, 0x61, 0x3c, 0x96, 0x4b, 0x96, 0x2e, 0x0d, 0x5c, 0xc4, 0x30,
	0xc0, 0xe0, 0xc6, 0xdd, 0xeb, 0xdf, 0xff, 0x7d, 0xf9, 0xf2, 0xf0, 0xc5, 0xd8, 0x1a, 0x07, 0xc6,
	0x79, 0xc7, 0xab, 0xe9, 0xe5, 0xcd, 0x83, 0x07, 0x96, 0xcd, 0x6c, 0x6e, 0x49, 0x67, 0x6e, 0x13,
	0x9f, 0xc2, 0x44, 0xb3, 0x4c, 0x24, 0x36, 0x15, 0xac, 0xaa, 0x9d, 0x75, 0x94, 0xb5, 0x2a, 0x01,
	0x5e, 0x16, 0x63, 0x3f, 0xe1, 0xc2, 0x14, 0x9b, 0xad, 0xde, 0x07, 0xc2, 0xc7, 0x4f, 0x6b, 0x8a,
	0x1c, 0x69, 0x65, 0x40, 0x3e, 0x82, 0x73, 0x42, 0x01, 0x39, 0xc2, 0x75, 0x2d, 0x43, 0xd4, 0x45,
	0xfd, 0x83, 0xa8, 0xae, 0x25, 0xb9, 0xc2, 0x41, 0xea, 0x54, 0x58, 0xef, 0xa2, 0x7e, 0x7b, 0x70,
	0xc2, 0x36, 0x40, 0xb6, 0x03, 0xb2, 0xa1, 0x29, 0xa2, 0x75, 0x81, 0xdc, 0xe0, 0xa6, 0xd3, 0xca,
	0xdc, 0x89, 0x5c, 0x84, 0x41, 0x37, 0xe8, 0xb7, 0x07, 0x97, 0xec, 0x5f, 0x31, 0x36, 0xda, 0x56,
	0xa3, 0x6a, 0xa9, 0xf7, 0x8a, 0x9b, 0xbb, 0x94, 0x50, 0x8c, 0xe7, 0x22, 0x19, 0x4a, 0x39, 0x03,
	0xe7, 0x4a, 0x99, 0x56, 0xf4, 0x27, 0x21, 0xa7, 0xb8, 0x91, 0xf9, 0xf8, 0x01, 0x8a, 0xd2, 0xeb,
	0x30, 0xda, 0xbe, 0xc8, 0x39, 0x6e, 0xad, 0x79, 0x22, 0xf7, 0x33, 0x08, 0x83, 0xf2, 0x6b, 0x1f,
	0xdc, 0xde, 0x7f, 0x2d, 0x29, 0x5a, 0x2c, 0x29, 0xfa, 0x59, 0x52, 0xf4, 0xb9, 0xa2, 0xb5, 0xc5,
	0x8a, 0xd6, 0xbe, 0x57, 0xb4, 0xf6, 0xcc, 0x95, 0xce, 0xa7, 0x3e, 0x66, 0x63, 0x9b, 0xf2, 0x8d,
	0xeb, 0x78, 0x2a, 0xb4, 0xd9, 0xce, 0xfc, 0x7d, 0x7f, 0x79, 0x9e, 0x17, 0x19, 0xb8, 0xb8, 0x51,
	0x1e, 0xe0, 0xfa, 0x37, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x1a, 0x94, 0x3f, 0x9d, 0x01, 0x00, 0x00,
}

func (m *QueuedSignedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueuedSignedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedSignedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignData) > 0 {
		for iNdEx := len(m.SignData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintConsensusQueue(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PubKey) > 0 {
		i -= len(m.PubKey)
		copy(dAtA[i:], m.PubKey)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.PubKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValAddress) > 0 {
		i -= len(m.ValAddress)
		copy(dAtA[i:], m.ValAddress)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.ValAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsensusQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsensusQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueuedSignedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovConsensusQueue(uint64(m.Id))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	if len(m.SignData) > 0 {
		for _, e := range m.SignData {
			l = e.Size()
			n += 1 + l + sovConsensusQueue(uint64(l))
		}
	}
	return n
}

func (m *SignData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValAddress)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.PubKey)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	return n
}

func sovConsensusQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsensusQueue(x uint64) (n int) {
	return sovConsensusQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueuedSignedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
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
			return fmt.Errorf("proto: QueuedSignedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueuedSignedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
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
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
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
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignData = append(m.SignData, &SignData{})
			if err := m.SignData[len(m.SignData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
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
func (m *SignData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
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
			return fmt.Errorf("proto: SignData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
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
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PubKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
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
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PubKey = append(m.PubKey[:0], dAtA[iNdEx:postIndex]...)
			if m.PubKey == nil {
				m.PubKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
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
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
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
func skipConsensusQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsensusQueue
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
					return 0, ErrIntOverflowConsensusQueue
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
					return 0, ErrIntOverflowConsensusQueue
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
				return 0, ErrInvalidLengthConsensusQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsensusQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsensusQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsensusQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsensusQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsensusQueue = fmt.Errorf("proto: unexpected end of group")
)