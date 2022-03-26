// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: blog/art.proto

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

type Art struct {
	Creator      string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id           uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body         string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	Price        string `protobuf:"bytes,5,opt,name=price,proto3" json:"price,omitempty"`
	OldArt       string `protobuf:"bytes,6,opt,name=oldArt,proto3" json:"oldArt,omitempty"`
	AuthorArt    string `protobuf:"bytes,7,opt,name=authorArt,proto3" json:"authorArt,omitempty"`
	AuthorSender string `protobuf:"bytes,8,opt,name=authorSender,proto3" json:"authorSender,omitempty"`
	Published    string `protobuf:"bytes,9,opt,name=published,proto3" json:"published,omitempty"`
	Image        string `protobuf:"bytes,10,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *Art) Reset()         { *m = Art{} }
func (m *Art) String() string { return proto.CompactTextString(m) }
func (*Art) ProtoMessage()    {}
func (*Art) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfe45dc78b3bd3c9, []int{0}
}
func (m *Art) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Art) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Art.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Art) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Art.Merge(m, src)
}
func (m *Art) XXX_Size() int {
	return m.Size()
}
func (m *Art) XXX_DiscardUnknown() {
	xxx_messageInfo_Art.DiscardUnknown(m)
}

var xxx_messageInfo_Art proto.InternalMessageInfo

func (m *Art) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Art) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Art) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Art) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Art) GetPrice() string {
	if m != nil {
		return m.Price
	}
	return ""
}

func (m *Art) GetOldArt() string {
	if m != nil {
		return m.OldArt
	}
	return ""
}

func (m *Art) GetAuthorArt() string {
	if m != nil {
		return m.AuthorArt
	}
	return ""
}

func (m *Art) GetAuthorSender() string {
	if m != nil {
		return m.AuthorSender
	}
	return ""
}

func (m *Art) GetPublished() string {
	if m != nil {
		return m.Published
	}
	return ""
}

func (m *Art) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func init() {
	proto.RegisterType((*Art)(nil), "cosmonaut.blog.blog.Art")
}

func init() { proto.RegisterFile("blog/art.proto", fileDescriptor_cfe45dc78b3bd3c9) }

var fileDescriptor_cfe45dc78b3bd3c9 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x6a, 0xf3, 0x30,
	0x14, 0x85, 0x2d, 0xc7, 0x71, 0x7e, 0x8b, 0x9f, 0x0c, 0x6a, 0x29, 0x1a, 0x8a, 0x08, 0x19, 0x4a,
	0x26, 0x7b, 0xe8, 0x0b, 0x34, 0x7d, 0x84, 0x74, 0xeb, 0x66, 0x5b, 0xc2, 0x16, 0xd8, 0xb9, 0x46,
	0xbe, 0x86, 0xe6, 0x2d, 0xfa, 0x58, 0x1d, 0x33, 0x76, 0x2c, 0xf6, 0x73, 0x14, 0x8a, 0xaf, 0xda,
	0x94, 0x2e, 0xd2, 0xfd, 0xce, 0x39, 0x5c, 0x2e, 0x87, 0xaf, 0x8b, 0x06, 0xaa, 0x2c, 0x77, 0x98,
	0x76, 0x0e, 0x10, 0xc4, 0x55, 0x09, 0x7d, 0x0b, 0xc7, 0x7c, 0xc0, 0x74, 0x76, 0xe8, 0xd9, 0x7e,
	0x32, 0xbe, 0xd8, 0x3b, 0x14, 0x92, 0xaf, 0x4a, 0x67, 0x72, 0x04, 0x27, 0xd9, 0x86, 0xed, 0x92,
	0xc3, 0x0f, 0x8a, 0x35, 0x0f, 0xad, 0x96, 0xe1, 0x86, 0xed, 0xa2, 0x43, 0x68, 0xb5, 0xb8, 0xe6,
	0x4b, 0xb4, 0xd8, 0x18, 0xb9, 0xa0, 0x9c, 0x07, 0x21, 0x78, 0x54, 0x80, 0x3e, 0xc9, 0x88, 0x44,
	0x9a, 0xe7, 0x64, 0xe7, 0x6c, 0x69, 0xe4, 0xd2, 0x27, 0x09, 0xc4, 0x0d, 0x8f, 0xa1, 0xd1, 0x7b,
	0x87, 0x32, 0x26, 0xf9, 0x9b, 0xc4, 0x2d, 0x4f, 0xf2, 0x01, 0x6b, 0x70, 0xb3, 0xb5, 0x22, 0xeb,
	0x57, 0x10, 0x5b, 0xfe, 0xdf, 0xc3, 0x93, 0x39, 0x6a, 0xe3, 0xe4, 0x3f, 0x0a, 0xfc, 0xd1, 0xe6,
	0x0d, 0xdd, 0x50, 0x34, 0xb6, 0xaf, 0x8d, 0x96, 0x89, 0xdf, 0x70, 0x11, 0xe6, 0x6b, 0x6c, 0x9b,
	0x57, 0x46, 0x72, 0x7f, 0x0d, 0xc1, 0xe3, 0xc3, 0xdb, 0xa8, 0xd8, 0x79, 0x54, 0xec, 0x63, 0x54,
	0xec, 0x75, 0x52, 0xc1, 0x79, 0x52, 0xc1, 0xfb, 0xa4, 0x82, 0xe7, 0xbb, 0xca, 0x62, 0x3d, 0x14,
	0x69, 0x09, 0x6d, 0x76, 0x69, 0x2e, 0xa3, 0x4e, 0x5f, 0xfc, 0x87, 0xa7, 0xce, 0xf4, 0x45, 0x4c,
	0xed, 0xde, 0x7f, 0x05, 0x00, 0x00, 0xff, 0xff, 0xc8, 0xf5, 0xb9, 0x0e, 0x6f, 0x01, 0x00, 0x00,
}

func (m *Art) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Art) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Art) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Image) > 0 {
		i -= len(m.Image)
		copy(dAtA[i:], m.Image)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Image)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Published) > 0 {
		i -= len(m.Published)
		copy(dAtA[i:], m.Published)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Published)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.AuthorSender) > 0 {
		i -= len(m.AuthorSender)
		copy(dAtA[i:], m.AuthorSender)
		i = encodeVarintArt(dAtA, i, uint64(len(m.AuthorSender)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.AuthorArt) > 0 {
		i -= len(m.AuthorArt)
		copy(dAtA[i:], m.AuthorArt)
		i = encodeVarintArt(dAtA, i, uint64(len(m.AuthorArt)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.OldArt) > 0 {
		i -= len(m.OldArt)
		copy(dAtA[i:], m.OldArt)
		i = encodeVarintArt(dAtA, i, uint64(len(m.OldArt)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Price) > 0 {
		i -= len(m.Price)
		copy(dAtA[i:], m.Price)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Price)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintArt(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintArt(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintArt(dAtA []byte, offset int, v uint64) int {
	offset -= sovArt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Art) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovArt(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.Price)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.OldArt)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.AuthorArt)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.AuthorSender)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.Published)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	l = len(m.Image)
	if l > 0 {
		n += 1 + l + sovArt(uint64(l))
	}
	return n
}

func sovArt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozArt(x uint64) (n int) {
	return sovArt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Art) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArt
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
			return fmt.Errorf("proto: Art: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Art: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Price = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldArt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OldArt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorArt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorArt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorSender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorSender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Published", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Published = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArt
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
				return ErrInvalidLengthArt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthArt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Image = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthArt
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
func skipArt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArt
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
					return 0, ErrIntOverflowArt
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
					return 0, ErrIntOverflowArt
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
				return 0, ErrInvalidLengthArt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupArt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthArt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthArt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupArt = fmt.Errorf("proto: unexpected end of group")
)