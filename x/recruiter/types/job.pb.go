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
	Id           uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title        string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Tags         string `protobuf:"bytes,4,opt,name=tags,proto3" json:"tags,omitempty"`
	PostDeadline string `protobuf:"bytes,5,opt,name=postDeadline,proto3" json:"postDeadline,omitempty"`
	JobDeadline  string `protobuf:"bytes,6,opt,name=jobDeadline,proto3" json:"jobDeadline,omitempty"`
	MaxPrice     string `protobuf:"bytes,7,opt,name=maxPrice,proto3" json:"maxPrice,omitempty"`
	Location     string `protobuf:"bytes,8,opt,name=location,proto3" json:"location,omitempty"`
	JobType      uint64 `protobuf:"varint,9,opt,name=jobType,proto3" json:"jobType,omitempty"`
	Status       string `protobuf:"bytes,10,opt,name=status,proto3" json:"status,omitempty"`
	Contractor   string `protobuf:"bytes,11,opt,name=contractor,proto3" json:"contractor,omitempty"`
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

func (m *Job) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
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
	// 318 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x86, 0x9b, 0xf4, 0xff, 0xf4, 0xd3, 0x37, 0x18, 0x84, 0x2c, 0x06, 0xab, 0xaa, 0x84, 0xd4,
	0x85, 0x64, 0x60, 0x62, 0x45, 0x4c, 0x4c, 0xa8, 0x30, 0xb1, 0x39, 0x8e, 0xd5, 0xba, 0x6d, 0x72,
	0x22, 0xfb, 0x54, 0x6d, 0xb9, 0x0a, 0x6e, 0x85, 0xbb, 0x60, 0xec, 0xc8, 0x88, 0xda, 0x1b, 0x41,
	0x75, 0x68, 0x08, 0x5b, 0x9e, 0xf7, 0xc9, 0x39, 0xb2, 0xfd, 0xc2, 0x99, 0xd5, 0xca, 0xae, 0x0c,
	0x69, 0x1b, 0xcf, 0x31, 0x89, 0x0a, 0x8b, 0x84, 0xec, 0x2a, 0xc3, 0x99, 0xcc, 0x32, 0x99, 0x5a,
	0xfd, 0x2a, 0x09, 0xed, 0x42, 0x66, 0x32, 0x8f, 0xd6, 0x7a, 0x8d, 0x76, 0x11, 0x55, 0x03, 0xa3,
	0xf7, 0x10, 0x9a, 0x0f, 0x98, 0xb0, 0xff, 0x10, 0x9a, 0x94, 0x07, 0xc3, 0x60, 0xdc, 0x9a, 0x84,
	0x26, 0x65, 0xe7, 0xd0, 0x26, 0x43, 0x4b, 0xcd, 0xc3, 0x61, 0x30, 0xee, 0x4f, 0x4a, 0x60, 0x43,
	0x18, 0xa4, 0xda, 0x29, 0x6b, 0x0a, 0x32, 0x98, 0xf3, 0xa6, 0x77, 0xf5, 0x88, 0x31, 0x68, 0x91,
	0x9c, 0x3a, 0xde, 0xf2, 0xca, 0x7f, 0xb3, 0x11, 0xfc, 0x2b, 0xd0, 0xd1, 0xbd, 0x96, 0xe9, 0xd2,
	0xe4, 0x9a, 0xb7, 0xbd, 0xfb, 0x93, 0x1d, 0x37, 0xcf, 0x31, 0xa9, 0x7e, 0xe9, 0x94, 0x9b, 0x6b,
	0x11, 0xbb, 0x84, 0x5e, 0x26, 0x37, 0x8f, 0xd6, 0x28, 0xcd, 0xbb, 0x5e, 0x57, 0x7c, 0x74, 0x4b,
	0x54, 0xd2, 0x1f, 0xaa, 0x57, 0xba, 0x13, 0x33, 0x0e, 0xdd, 0x39, 0x26, 0xcf, 0xdb, 0x42, 0xf3,
	0xbe, 0xbf, 0xde, 0x09, 0xd9, 0x05, 0x74, 0x1c, 0x49, 0x5a, 0x39, 0x0e, 0x7e, 0xe6, 0x87, 0x98,
	0x00, 0x50, 0x98, 0x93, 0x95, 0x8a, 0xd0, 0xf2, 0x81, 0x77, 0xb5, 0xe4, 0xee, 0xe9, 0x63, 0x2f,
	0x82, 0xdd, 0x5e, 0x04, 0x5f, 0x7b, 0x11, 0xbc, 0x1d, 0x44, 0x63, 0x77, 0x10, 0x8d, 0xcf, 0x83,
	0x68, 0xbc, 0xdc, 0x4e, 0x0d, 0xcd, 0x56, 0x49, 0xa4, 0x30, 0x8b, 0xeb, 0xef, 0x7f, 0x7d, 0x2a,
	0x20, 0x2e, 0x0b, 0x88, 0x37, 0xf1, 0x6f, 0x67, 0xb4, 0x2d, 0xb4, 0x4b, 0x3a, 0xbe, 0xb6, 0x9b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xfd, 0x97, 0x43, 0xcd, 0x01, 0x00, 0x00,
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
		i -= len(m.Tags)
		copy(dAtA[i:], m.Tags)
		i = encodeVarintJob(dAtA, i, uint64(len(m.Tags)))
		i--
		dAtA[i] = 0x22
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
	l = len(m.Tags)
	if l > 0 {
		n += 1 + l + sovJob(uint64(l))
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
			m.Tags = string(dAtA[iNdEx:postIndex])
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