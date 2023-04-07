// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: evm/deploy_new_smart_contract_proposal.proto

package types

import (
	fmt "fmt"
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

type DeployNewSmartContractProposal struct {
	Title       string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	AbiJSON     string `protobuf:"bytes,3,opt,name=abiJSON,proto3" json:"abiJSON,omitempty"`
	BytecodeHex string `protobuf:"bytes,4,opt,name=bytecodeHex,proto3" json:"bytecodeHex,omitempty"`
}

func (m *DeployNewSmartContractProposal) Reset()         { *m = DeployNewSmartContractProposal{} }
func (m *DeployNewSmartContractProposal) String() string { return proto.CompactTextString(m) }
func (*DeployNewSmartContractProposal) ProtoMessage()    {}
func (*DeployNewSmartContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_ba75a8a2093c3d91, []int{0}
}
func (m *DeployNewSmartContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DeployNewSmartContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployNewSmartContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DeployNewSmartContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployNewSmartContractProposal.Merge(m, src)
}
func (m *DeployNewSmartContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *DeployNewSmartContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployNewSmartContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DeployNewSmartContractProposal proto.InternalMessageInfo

func (m *DeployNewSmartContractProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DeployNewSmartContractProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DeployNewSmartContractProposal) GetAbiJSON() string {
	if m != nil {
		return m.AbiJSON
	}
	return ""
}

func (m *DeployNewSmartContractProposal) GetBytecodeHex() string {
	if m != nil {
		return m.BytecodeHex
	}
	return ""
}

func init() {
	proto.RegisterType((*DeployNewSmartContractProposal)(nil), "palomachain.paloma.evm.DeployNewSmartContractProposal")
}

func init() {
	proto.RegisterFile("evm/deploy_new_smart_contract_proposal.proto", fileDescriptor_ba75a8a2093c3d91)
}

var fileDescriptor_ba75a8a2093c3d91 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x2d, 0xcb, 0xd5,
	0x4f, 0x49, 0x2d, 0xc8, 0xc9, 0xaf, 0x8c, 0xcf, 0x4b, 0x2d, 0x8f, 0x2f, 0xce, 0x4d, 0x2c, 0x2a,
	0x89, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x89, 0x2f, 0x28, 0xca, 0x2f, 0xc8, 0x2f,
	0x4e, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x48, 0xcc, 0xc9, 0xcf, 0x4d,
	0x4c, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x83, 0xb0, 0xf5, 0x52, 0xcb, 0x72, 0x95, 0xa6, 0x30, 0x72,
	0xc9, 0xb9, 0x80, 0x0d, 0xf1, 0x4b, 0x2d, 0x0f, 0x06, 0x19, 0xe1, 0x0c, 0x35, 0x21, 0x00, 0x6a,
	0x80, 0x90, 0x08, 0x17, 0x6b, 0x49, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0x84, 0x23, 0xa4, 0xc0, 0xc5, 0x9d, 0x92, 0x5a, 0x9c, 0x5c, 0x94, 0x59, 0x50, 0x92, 0x99,
	0x9f, 0x27, 0xc1, 0x04, 0x96, 0x43, 0x16, 0x12, 0x92, 0xe0, 0x62, 0x4f, 0x4c, 0xca, 0xf4, 0x0a,
	0xf6, 0xf7, 0x93, 0x60, 0x06, 0xcb, 0xc2, 0xb8, 0x20, 0xbd, 0x49, 0x95, 0x25, 0xa9, 0xc9, 0xf9,
	0x29, 0xa9, 0x1e, 0xa9, 0x15, 0x12, 0x2c, 0x10, 0xbd, 0x48, 0x42, 0x4e, 0xce, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x99, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x8f, 0xe4, 0x27, 0x28, 0x5b, 0xbf, 0x42, 0x1f, 0x14, 0x2c, 0x25, 0x95, 0x05,
	0xa9, 0xc5, 0x49, 0x6c, 0x60, 0xaf, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x56, 0x38, 0x6e,
	0x42, 0x2a, 0x01, 0x00, 0x00,
}

func (m *DeployNewSmartContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployNewSmartContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployNewSmartContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BytecodeHex) > 0 {
		i -= len(m.BytecodeHex)
		copy(dAtA[i:], m.BytecodeHex)
		i = encodeVarintDeployNewSmartContractProposal(dAtA, i, uint64(len(m.BytecodeHex)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AbiJSON) > 0 {
		i -= len(m.AbiJSON)
		copy(dAtA[i:], m.AbiJSON)
		i = encodeVarintDeployNewSmartContractProposal(dAtA, i, uint64(len(m.AbiJSON)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintDeployNewSmartContractProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintDeployNewSmartContractProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDeployNewSmartContractProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovDeployNewSmartContractProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DeployNewSmartContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovDeployNewSmartContractProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovDeployNewSmartContractProposal(uint64(l))
	}
	l = len(m.AbiJSON)
	if l > 0 {
		n += 1 + l + sovDeployNewSmartContractProposal(uint64(l))
	}
	l = len(m.BytecodeHex)
	if l > 0 {
		n += 1 + l + sovDeployNewSmartContractProposal(uint64(l))
	}
	return n
}

func sovDeployNewSmartContractProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDeployNewSmartContractProposal(x uint64) (n int) {
	return sovDeployNewSmartContractProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DeployNewSmartContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDeployNewSmartContractProposal
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
			return fmt.Errorf("proto: DeployNewSmartContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployNewSmartContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployNewSmartContractProposal
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
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployNewSmartContractProposal
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
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbiJSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployNewSmartContractProposal
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
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbiJSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytecodeHex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDeployNewSmartContractProposal
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
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDeployNewSmartContractProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytecodeHex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDeployNewSmartContractProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDeployNewSmartContractProposal
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
func skipDeployNewSmartContractProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDeployNewSmartContractProposal
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
					return 0, ErrIntOverflowDeployNewSmartContractProposal
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
					return 0, ErrIntOverflowDeployNewSmartContractProposal
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
				return 0, ErrInvalidLengthDeployNewSmartContractProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDeployNewSmartContractProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDeployNewSmartContractProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDeployNewSmartContractProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDeployNewSmartContractProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDeployNewSmartContractProposal = fmt.Errorf("proto: unexpected end of group")
)
