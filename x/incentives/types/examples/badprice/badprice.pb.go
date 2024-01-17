// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: slinky/incentives/v1/badprice.proto

package badprice

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

// BadPriceIncentive is a message that contains the information about a bad
// price that was submitted by a validator.
//
// NOTE: This is an example of a bad price incentive. It is not used in
// production.
type BadPriceIncentive struct {
	// Validator is the address of the validator that submitted the bad price.
	Validator string `protobuf:"bytes,1,opt,name=validator,proto3" json:"validator,omitempty"`
	// Amount is the amount to slash.
	Amount string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *BadPriceIncentive) Reset()         { *m = BadPriceIncentive{} }
func (m *BadPriceIncentive) String() string { return proto.CompactTextString(m) }
func (*BadPriceIncentive) ProtoMessage()    {}
func (*BadPriceIncentive) Descriptor() ([]byte, []int) {
	return fileDescriptor_90b575131cd0f9c4, []int{0}
}
func (m *BadPriceIncentive) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BadPriceIncentive) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BadPriceIncentive.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BadPriceIncentive) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BadPriceIncentive.Merge(m, src)
}
func (m *BadPriceIncentive) XXX_Size() int {
	return m.Size()
}
func (m *BadPriceIncentive) XXX_DiscardUnknown() {
	xxx_messageInfo_BadPriceIncentive.DiscardUnknown(m)
}

var xxx_messageInfo_BadPriceIncentive proto.InternalMessageInfo

func (m *BadPriceIncentive) GetValidator() string {
	if m != nil {
		return m.Validator
	}
	return ""
}

func (m *BadPriceIncentive) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func init() {
	proto.RegisterType((*BadPriceIncentive)(nil), "slinky.incentives.v1.BadPriceIncentive")
}

func init() {
	proto.RegisterFile("slinky/incentives/v1/badprice.proto", fileDescriptor_90b575131cd0f9c4)
}

var fileDescriptor_90b575131cd0f9c4 = []byte{
	// 255 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0xce, 0xc9, 0xcc,
	0xcb, 0xae, 0xd4, 0xcf, 0xcc, 0x4b, 0x4e, 0xcd, 0x2b, 0xc9, 0x2c, 0x4b, 0x2d, 0xd6, 0x2f, 0x33,
	0xd4, 0x4f, 0x4a, 0x4c, 0x29, 0x28, 0xca, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x81, 0x28, 0xd2, 0x43, 0x28, 0xd2, 0x2b, 0x33, 0x94, 0x12, 0x4c, 0xcc, 0xcd, 0xcc, 0xcb,
	0xd7, 0x07, 0x93, 0x10, 0x85, 0x52, 0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x60, 0x9e,
	0x3e, 0x84, 0x03, 0x91, 0x52, 0x9a, 0xc8, 0xc8, 0x25, 0xe8, 0x94, 0x98, 0x12, 0x00, 0x32, 0xd6,
	0x13, 0x66, 0x8e, 0x90, 0x0c, 0x17, 0x67, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0x7e, 0x91,
	0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x42, 0x40, 0x48, 0x8c, 0x8b, 0x2d, 0x31, 0x37, 0xbf,
	0x34, 0xaf, 0x44, 0x82, 0x09, 0x2c, 0x05, 0xe5, 0x59, 0xb9, 0x9d, 0xda, 0xa2, 0x2b, 0x87, 0xcd,
	0x4d, 0x7a, 0x70, 0x93, 0xbb, 0x9e, 0x6f, 0xd0, 0x92, 0x87, 0xfa, 0x2d, 0xbf, 0x28, 0x31, 0x39,
	0x27, 0x55, 0x1f, 0xc3, 0x76, 0xa7, 0xc8, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c,
	0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63,
	0x88, 0xb2, 0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xce,
	0x2c, 0xd0, 0xcd, 0x4d, 0x2d, 0xd3, 0x87, 0x1a, 0x57, 0x81, 0x1c, 0x58, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0xfa, 0xa9, 0x15, 0x89, 0xb9, 0x05, 0x39, 0xa9, 0xc5, 0xf0, 0x80, 0x4b, 0x62, 0x03, 0xfb,
	0xda, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x72, 0xb7, 0x35, 0xcf, 0x60, 0x01, 0x00, 0x00,
}

func (m *BadPriceIncentive) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BadPriceIncentive) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BadPriceIncentive) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintBadprice(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Validator) > 0 {
		i -= len(m.Validator)
		copy(dAtA[i:], m.Validator)
		i = encodeVarintBadprice(dAtA, i, uint64(len(m.Validator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBadprice(dAtA []byte, offset int, v uint64) int {
	offset -= sovBadprice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BadPriceIncentive) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Validator)
	if l > 0 {
		n += 1 + l + sovBadprice(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovBadprice(uint64(l))
	}
	return n
}

func sovBadprice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBadprice(x uint64) (n int) {
	return sovBadprice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BadPriceIncentive) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBadprice
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
			return fmt.Errorf("proto: BadPriceIncentive: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BadPriceIncentive: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadprice
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
				return ErrInvalidLengthBadprice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadprice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBadprice
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
				return ErrInvalidLengthBadprice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBadprice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBadprice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBadprice
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
func skipBadprice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBadprice
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
					return 0, ErrIntOverflowBadprice
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
					return 0, ErrIntOverflowBadprice
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
				return 0, ErrInvalidLengthBadprice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBadprice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBadprice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBadprice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBadprice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBadprice = fmt.Errorf("proto: unexpected end of group")
)
