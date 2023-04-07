// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/job.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type JobDefinition struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ABI     string `protobuf:"bytes,2,opt,name=ABI,proto3" json:"ABI,omitempty"`
}

func (m *JobDefinition) Reset()         { *m = JobDefinition{} }
func (m *JobDefinition) String() string { return proto.CompactTextString(m) }
func (*JobDefinition) ProtoMessage()    {}
func (*JobDefinition) Descriptor() ([]byte, []int) {
	return fileDescriptor_33609e408ca6cc36, []int{0}
}
func (m *JobDefinition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobDefinition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JobDefinition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JobDefinition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobDefinition.Merge(m, src)
}
func (m *JobDefinition) XXX_Size() int {
	return m.Size()
}
func (m *JobDefinition) XXX_DiscardUnknown() {
	xxx_messageInfo_JobDefinition.DiscardUnknown(m)
}

var xxx_messageInfo_JobDefinition proto.InternalMessageInfo

func (m *JobDefinition) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *JobDefinition) GetABI() string {
	if m != nil {
		return m.ABI
	}
	return ""
}

type JobPayload struct {
	HexPayload string `protobuf:"bytes,1,opt,name=hexPayload,proto3" json:"hexPayload,omitempty"`
}

func (m *JobPayload) Reset()         { *m = JobPayload{} }
func (m *JobPayload) String() string { return proto.CompactTextString(m) }
func (*JobPayload) ProtoMessage()    {}
func (*JobPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_33609e408ca6cc36, []int{1}
}
func (m *JobPayload) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JobPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JobPayload.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JobPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobPayload.Merge(m, src)
}
func (m *JobPayload) XXX_Size() int {
	return m.Size()
}
func (m *JobPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_JobPayload.DiscardUnknown(m)
}

var xxx_messageInfo_JobPayload proto.InternalMessageInfo

func (m *JobPayload) GetHexPayload() string {
	if m != nil {
		return m.HexPayload
	}
	return ""
}

func init() {
	proto.RegisterType((*JobDefinition)(nil), "palomachain.paloma.evm.JobDefinition")
	proto.RegisterType((*JobPayload)(nil), "palomachain.paloma.evm.JobPayload")
}

func init() { proto.RegisterFile("evm/job.proto", fileDescriptor_33609e408ca6cc36) }

var fileDescriptor_33609e408ca6cc36 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2d, 0xcb, 0xd5,
	0xcf, 0xca, 0x4f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x48, 0xcc, 0xc9, 0xcf,
	0x4d, 0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xb0, 0xf5, 0x52, 0xcb, 0x72, 0xa5, 0x44, 0xd2,
	0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf4, 0x41, 0x2c, 0x88, 0x6a, 0x25, 0x6b, 0x2e, 0x5e, 0xaf, 0xfc,
	0x24, 0x97, 0xd4, 0xb4, 0xcc, 0xbc, 0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x21, 0x09, 0x2e, 0xf6, 0xc4,
	0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x18, 0x57, 0x48,
	0x80, 0x8b, 0xd9, 0xd1, 0xc9, 0x53, 0x82, 0x09, 0x2c, 0x0a, 0x62, 0x2a, 0xe9, 0x70, 0x71, 0x79,
	0xe5, 0x27, 0x05, 0x24, 0x56, 0xe6, 0xe4, 0x27, 0xa6, 0x08, 0xc9, 0x71, 0x71, 0x65, 0xa4, 0x56,
	0x40, 0x79, 0x50, 0xcd, 0x48, 0x22, 0x4e, 0xce, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7,
	0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c,
	0xc7, 0x10, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0xe4,
	0x7a, 0x28, 0x5b, 0xbf, 0x42, 0x1f, 0xe4, 0xc3, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0,
	0xb3, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x80, 0x1f, 0x0a, 0x8a, 0xf5, 0x00, 0x00, 0x00,
}

func (m *JobDefinition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobDefinition) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JobDefinition) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ABI) > 0 {
		i -= len(m.ABI)
		copy(dAtA[i:], m.ABI)
		i = encodeVarintJob(dAtA, i, uint64(len(m.ABI)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JobPayload) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JobPayload) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JobPayload) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.HexPayload) > 0 {
		i -= len(m.HexPayload)
		copy(dAtA[i:], m.HexPayload)
		i = encodeVarintJob(dAtA, i, uint64(len(m.HexPayload)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJob(dAtA []byte, offset int, v uint64) int {
	offset -= sovJob(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JobDefinition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.ABI)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	return n
}

func (m *JobPayload) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.HexPayload)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	return n
}

func sovJob(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJob(x uint64) (n int) {
	return sovJob(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JobDefinition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJob
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
			return fmt.Errorf("proto: JobDefinition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobDefinition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ABI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ABI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJob
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
func (m *JobPayload) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJob
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
			return fmt.Errorf("proto: JobPayload: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JobPayload: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HexPayload", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return ErrInvalidLengthJob
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJob
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HexPayload = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJob(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJob
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
func skipJob(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJob
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
					return 0, ErrIntOverflowJob
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
					return 0, ErrIntOverflowJob
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
				return 0, ErrInvalidLengthJob
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJob
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJob
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJob        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJob          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJob = fmt.Errorf("proto: unexpected end of group")
)
