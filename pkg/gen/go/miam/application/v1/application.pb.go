// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: miam/application/v1/application.proto

package applicationv1

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	proto "github.com/gogo/protobuf/proto"
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

// Application is a resource provider.
type Application struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Label                string   `protobuf:"bytes,2,opt,name=label,proto3" json:"label,omitempty"`
	Urn                  string   `protobuf:"bytes,3,opt,name=urn,proto3" json:"urn,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_1703c26bc2bfdb8e, []int{0}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Application.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}

func (m *Application) XXX_Size() int {
	return m.Size()
}

func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Application) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Application) GetUrn() string {
	if m != nil {
		return m.Urn
	}
	return ""
}

func init() {
	proto.RegisterType((*Application)(nil), "miam.application.v1.Application")
}

func init() {
	proto.RegisterFile("miam/application/v1/application.proto", fileDescriptor_1703c26bc2bfdb8e)
}

var fileDescriptor_1703c26bc2bfdb8e = []byte{
	// 192 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0xcd, 0xcd, 0x4c, 0xcc,
	0xd5, 0x4f, 0x2c, 0x28, 0xc8, 0xc9, 0x4c, 0x4e, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0x33, 0x44,
	0xe6, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x83, 0x94, 0xe9, 0x21, 0x8b, 0x97, 0x19,
	0x2a, 0xb9, 0x72, 0x71, 0x3b, 0x22, 0x44, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x98, 0x32, 0x53, 0x84, 0x44, 0xb8, 0x58, 0x73, 0x12, 0x93, 0x52, 0x73,
	0x24, 0x98, 0xc0, 0x42, 0x10, 0x8e, 0x90, 0x00, 0x17, 0x73, 0x69, 0x51, 0x9e, 0x04, 0x33, 0x58,
	0x0c, 0xc4, 0x74, 0x6a, 0x66, 0x3c, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0xb9, 0x14, 0xf2, 0x8b, 0xd2, 0xf5, 0xaa, 0x52, 0xf3, 0x32, 0x4b, 0x32, 0x12, 0x8b,
	0xf4, 0xb0, 0xd8, 0xeb, 0x24, 0x80, 0x64, 0x6b, 0x00, 0xc8, 0x79, 0x01, 0x8c, 0x51, 0xbc, 0x48,
	0x6a, 0xca, 0x0c, 0x17, 0x31, 0x31, 0xfb, 0x3a, 0x46, 0xac, 0x62, 0x12, 0xf6, 0x05, 0x69, 0x47,
	0x52, 0xae, 0x17, 0x66, 0x78, 0x0a, 0x22, 0x1a, 0x83, 0x24, 0x1a, 0x13, 0x66, 0x98, 0xc4, 0x06,
	0xf6, 0xa8, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x49, 0x78, 0x45, 0x11, 0x01, 0x00, 0x00,
}

func (m *Application) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Application) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Application) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Urn) > 0 {
		i -= len(m.Urn)
		copy(dAtA[i:], m.Urn)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Urn)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Label) > 0 {
		i -= len(m.Label)
		copy(dAtA[i:], m.Label)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Label)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintApplication(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApplication(dAtA []byte, offset int, v uint64) int {
	offset -= sovApplication(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *Application) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.Label)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	l = len(m.Urn)
	if l > 0 {
		n += 1 + l + sovApplication(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovApplication(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozApplication(x uint64) (n int) {
	return sovApplication(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *Application) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApplication
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
			return fmt.Errorf("proto: Application: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Application: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Label", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Label = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Urn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApplication
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
				return ErrInvalidLengthApplication
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApplication
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Urn = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipApplication(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthApplication
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipApplication(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApplication
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
					return 0, ErrIntOverflowApplication
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowApplication
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
				return 0, ErrInvalidLengthApplication
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthApplication
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowApplication
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipApplication(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthApplication
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthApplication = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApplication   = fmt.Errorf("proto: integer overflow")
)
