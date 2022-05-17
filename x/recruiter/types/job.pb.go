// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: recruiter/job.proto

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

type Job struct {
	Id           uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tags         []string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	PostDeadline string   `protobuf:"bytes,5,opt,name=postDeadline,proto3" json:"postDeadline,omitempty"`
	JobDeadline  string   `protobuf:"bytes,6,opt,name=jobDeadline,proto3" json:"jobDeadline,omitempty"`
	MaxPrice     string   `protobuf:"bytes,7,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
	Location     string   `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	JobType      uint64   `protobuf:"varint,9,opt,name=jobType,proto3" json:"jobType,omitempty"`
	Status       string   `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Contractor   string   `protobuf:"bytes,11,opt,name=contractor,proto3" json:"contractor,omitempty"`
}

func (m *Job) Reset()         { *m = Job{} }
func (m *Job) String() string { return proto.CompactTextString(m) }
func (*Job) ProtoMessage()    {}
func (*Job) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3bbc81625190371, []int{0}
}
func (m *Job) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Job) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Job.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Job) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Job.Merge(m, src)
}
func (m *Job) XXX_Size() int {
	return m.Size()
}
func (m *Job) XXX_DiscardUnknown() {
	xxx_messageInfo_Job.DiscardUnknown(m)
}

var xxx_messageInfo_Job proto.InternalMessageInfo

func (m *Job) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Job) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Job) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Job) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Job) GetPostDeadline() string {
	if m != nil {
		return m.PostDeadline
	}
	return ""
}

func (m *Job) GetJobDeadline() string {
	if m != nil {
		return m.JobDeadline
	}
	return ""
}

func (m *Job) GetMaxPrice() string {
	if m != nil {
		return m.MaxPrice
	}
	return ""
}

func (m *Job) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *Job) GetJobType() uint64 {
	if m != nil {
		return m.JobType
	}
	return 0
}

func (m *Job) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Job) GetContractor() string {
	if m != nil {
		return m.Contractor
	}
	return ""
}

func init() {
	proto.RegisterType((*Job)(nil), "mohammadrezatorkaman.wework.recruiter.Job")
}

func init() { proto.RegisterFile("recruiter/job.proto", fileDescriptor_d3bbc81625190371) }

var fileDescriptor_d3bbc81625190371 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x9b, 0xf4, 0xff, 0xed, 0xa7, 0x6f, 0x30, 0x08, 0x59, 0x0c, 0x56, 0x54, 0x09, 0xa9,
	0x0b, 0xc9, 0xc0, 0xc4, 0x8a, 0x98, 0x98, 0x50, 0x61, 0x62, 0x73, 0x1c, 0xab, 0x75, 0xdb, 0xe4,
	0x46, 0xf6, 0xad, 0xda, 0xf2, 0x14, 0xbc, 0x0a, 0x6f, 0xc1, 0xd8, 0x91, 0x11, 0xb5, 0x2f, 0x82,
	0xea, 0xd0, 0x12, 0x36, 0x9f, 0xf3, 0xf3, 0xb9, 0xb2, 0xef, 0x81, 0x33, 0xab, 0x95, 0x5d, 0x1a,
	0xd2, 0x36, 0x99, 0x61, 0x1a, 0x97, 0x16, 0x09, 0xd9, 0x55, 0x8e, 0x53, 0x99, 0xe7, 0x32, 0xb3,
	0xfa, 0x55, 0x12, 0xda, 0xb9, 0xcc, 0x65, 0x11, 0xaf, 0xf4, 0x0a, 0xed, 0x3c, 0x3e, 0x05, 0x86,
	0xef, 0x21, 0x34, 0x1f, 0x30, 0x65, 0xff, 0x21, 0x34, 0x19, 0x0f, 0xa2, 0x60, 0xd4, 0x1a, 0x87,
	0x26, 0x63, 0xe7, 0xd0, 0x26, 0x43, 0x0b, 0xcd, 0xc3, 0x28, 0x18, 0xf5, 0xc7, 0x95, 0x60, 0x11,
	0x0c, 0x32, 0xed, 0x94, 0x35, 0x25, 0x19, 0x2c, 0x78, 0xd3, 0xb3, 0xba, 0xc5, 0x18, 0xb4, 0x48,
	0x4e, 0x1c, 0x6f, 0x45, 0xcd, 0x51, 0x7f, 0xec, 0xcf, 0x6c, 0x08, 0xff, 0x4a, 0x74, 0x74, 0xaf,
	0x65, 0xb6, 0x30, 0x85, 0xe6, 0x6d, 0x1f, 0xfb, 0xe3, 0x1d, 0x26, 0xcf, 0x30, 0x3d, 0x5d, 0xe9,
	0x54, 0x93, 0x6b, 0x16, 0xbb, 0x84, 0x5e, 0x2e, 0xd7, 0x8f, 0xd6, 0x28, 0xcd, 0xbb, 0x1e, 0x9f,
	0xf4, 0x81, 0x2d, 0x50, 0x49, 0xff, 0xa8, 0x5e, 0xc5, 0x8e, 0x9a, 0x71, 0xe8, 0xce, 0x30, 0x7d,
	0xde, 0x94, 0x9a, 0xf7, 0xfd, 0xf7, 0x8e, 0x92, 0x5d, 0x40, 0xc7, 0x91, 0xa4, 0xa5, 0xe3, 0xe0,
	0x33, 0x3f, 0x8a, 0x09, 0x00, 0x85, 0x05, 0x59, 0xa9, 0x08, 0x2d, 0x1f, 0x78, 0x56, 0x73, 0xee,
	0x9e, 0x3e, 0x76, 0x22, 0xd8, 0xee, 0x44, 0xf0, 0xb5, 0x13, 0xc1, 0xdb, 0x5e, 0x34, 0xb6, 0x7b,
	0xd1, 0xf8, 0xdc, 0x8b, 0xc6, 0xcb, 0xed, 0xc4, 0xd0, 0x74, 0x99, 0xc6, 0x0a, 0xf3, 0xa4, 0xbe,
	0xff, 0xeb, 0x63, 0x01, 0x49, 0x55, 0x40, 0xb2, 0x4e, 0x7e, 0x3b, 0xa3, 0x4d, 0xa9, 0x5d, 0xda,
	0xf1, 0xb5, 0xdd, 0x7c, 0x07, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xbe, 0x48, 0x76, 0xcd, 0x01, 0x00,
	0x00,
}

func (m *Job) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Job) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Job) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contractor) > 0 {
		i -= len(m.Contractor)
		copy(dAtA[i:], m.Contractor)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Contractor)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0x52
	}
	if m.JobType != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.JobType))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Location) > 0 {
		i -= len(m.Location)
		copy(dAtA[i:], m.Location)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Location)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.MaxPrice) > 0 {
		i -= len(m.MaxPrice)
		copy(dAtA[i:], m.MaxPrice)
		i = encodeVarintJob(dAtA, i, uint64(len(m.MaxPrice)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.JobDeadline) > 0 {
		i -= len(m.JobDeadline)
		copy(dAtA[i:], m.JobDeadline)
		i = encodeVarintJob(dAtA, i, uint64(len(m.JobDeadline)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PostDeadline) > 0 {
		i -= len(m.PostDeadline)
		copy(dAtA[i:], m.PostDeadline)
		i = encodeVarintJob(dAtA, i, uint64(len(m.PostDeadline)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintJob(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintJob(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
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
func (m *Job) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovJob(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovJob(uint64(l))
		}
	}
	l = len(m.PostDeadline)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.JobDeadline)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.MaxPrice)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.Location)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	if m.JobType != 0 {
		n += 1 + sovJob(uint64(m.JobType))
	}
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
	}
	l = len(m.Contractor)
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
func (m *Job) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: Job: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Job: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
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
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
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
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
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
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostDeadline", wireType)
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
			m.PostDeadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobDeadline", wireType)
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
			m.JobDeadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxPrice", wireType)
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
			m.MaxPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Location", wireType)
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
			m.Location = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobType", wireType)
			}
			m.JobType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJob
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobType |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contractor", wireType)
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
			m.Contractor = string(dAtA[iNdEx:postIndex])
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
