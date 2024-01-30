// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v2/state.proto

package data

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

// DataID stores a compact data ID and its full IRI.
type DataID struct {
	// id is the compact automatically-generated data ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// iri is the IRI of the data which contains its full ContentHash.
	Iri string `protobuf:"bytes,2,opt,name=iri,proto3" json:"iri,omitempty"`
}

func (m *DataID) Reset()         { *m = DataID{} }
func (m *DataID) String() string { return proto.CompactTextString(m) }
func (*DataID) ProtoMessage()    {}
func (*DataID) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd916028e0fe6d62, []int{0}
}
func (m *DataID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataID.Merge(m, src)
}
func (m *DataID) XXX_Size() int {
	return m.Size()
}
func (m *DataID) XXX_DiscardUnknown() {
	xxx_messageInfo_DataID.DiscardUnknown(m)
}

var xxx_messageInfo_DataID proto.InternalMessageInfo

func (m *DataID) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DataID) GetIri() string {
	if m != nil {
		return m.Iri
	}
	return ""
}

// DataAnchor stores the anchor timestamp for a data object.
type DataAnchor struct {
	// id is the compact data ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// timestamp is the anchor timestamp for this object - the time at which
	// it was first known to the blockchain.
	Timestamp *types.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *DataAnchor) Reset()         { *m = DataAnchor{} }
func (m *DataAnchor) String() string { return proto.CompactTextString(m) }
func (*DataAnchor) ProtoMessage()    {}
func (*DataAnchor) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd916028e0fe6d62, []int{1}
}
func (m *DataAnchor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataAnchor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataAnchor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataAnchor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataAnchor.Merge(m, src)
}
func (m *DataAnchor) XXX_Size() int {
	return m.Size()
}
func (m *DataAnchor) XXX_DiscardUnknown() {
	xxx_messageInfo_DataAnchor.DiscardUnknown(m)
}

var xxx_messageInfo_DataAnchor proto.InternalMessageInfo

func (m *DataAnchor) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DataAnchor) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// DataAttestor is a join table for associating data IDs and attestors.
type DataAttestor struct {
	// id is the compact data ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// attestor is the account address of the attestor.
	Attestor []byte `protobuf:"bytes,2,opt,name=attestor,proto3" json:"attestor,omitempty"`
	// timestamp is the time at which the attestor signed this data object.
	Timestamp *types.Timestamp `protobuf:"bytes,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (m *DataAttestor) Reset()         { *m = DataAttestor{} }
func (m *DataAttestor) String() string { return proto.CompactTextString(m) }
func (*DataAttestor) ProtoMessage()    {}
func (*DataAttestor) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd916028e0fe6d62, []int{2}
}
func (m *DataAttestor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataAttestor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataAttestor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataAttestor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataAttestor.Merge(m, src)
}
func (m *DataAttestor) XXX_Size() int {
	return m.Size()
}
func (m *DataAttestor) XXX_DiscardUnknown() {
	xxx_messageInfo_DataAttestor.DiscardUnknown(m)
}

var xxx_messageInfo_DataAttestor proto.InternalMessageInfo

func (m *DataAttestor) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DataAttestor) GetAttestor() []byte {
	if m != nil {
		return m.Attestor
	}
	return nil
}

func (m *DataAttestor) GetTimestamp() *types.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

// Resolver describes a data resolver.
type Resolver struct {
	// id is the ID of the resolver.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// url is the URL of the resolver.
	Url string `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	// manager is the bytes address of the resolver manager who is allowed
	// to make calls to Msg/RegisterResolver for this resolver.
	Manager []byte `protobuf:"bytes,3,opt,name=manager,proto3" json:"manager,omitempty"`
}

func (m *Resolver) Reset()         { *m = Resolver{} }
func (m *Resolver) String() string { return proto.CompactTextString(m) }
func (*Resolver) ProtoMessage()    {}
func (*Resolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd916028e0fe6d62, []int{3}
}
func (m *Resolver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Resolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Resolver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Resolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resolver.Merge(m, src)
}
func (m *Resolver) XXX_Size() int {
	return m.Size()
}
func (m *Resolver) XXX_DiscardUnknown() {
	xxx_messageInfo_Resolver.DiscardUnknown(m)
}

var xxx_messageInfo_Resolver proto.InternalMessageInfo

func (m *Resolver) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Resolver) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Resolver) GetManager() []byte {
	if m != nil {
		return m.Manager
	}
	return nil
}

// DataResolver is a join table between data objects and resolvers and indicates
// that a resolver claims to be able to resolve this data object.
type DataResolver struct {
	// id is the compact data ID.
	Id []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// resolver_id is the ID of the resolver.
	ResolverId uint64 `protobuf:"varint,2,opt,name=resolver_id,json=resolverId,proto3" json:"resolver_id,omitempty"`
}

func (m *DataResolver) Reset()         { *m = DataResolver{} }
func (m *DataResolver) String() string { return proto.CompactTextString(m) }
func (*DataResolver) ProtoMessage()    {}
func (*DataResolver) Descriptor() ([]byte, []int) {
	return fileDescriptor_fd916028e0fe6d62, []int{4}
}
func (m *DataResolver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataResolver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataResolver.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataResolver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataResolver.Merge(m, src)
}
func (m *DataResolver) XXX_Size() int {
	return m.Size()
}
func (m *DataResolver) XXX_DiscardUnknown() {
	xxx_messageInfo_DataResolver.DiscardUnknown(m)
}

var xxx_messageInfo_DataResolver proto.InternalMessageInfo

func (m *DataResolver) GetId() []byte {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DataResolver) GetResolverId() uint64 {
	if m != nil {
		return m.ResolverId
	}
	return 0
}

func init() {
	proto.RegisterType((*DataID)(nil), "regen.data.v2.DataID")
	proto.RegisterType((*DataAnchor)(nil), "regen.data.v2.DataAnchor")
	proto.RegisterType((*DataAttestor)(nil), "regen.data.v2.DataAttestor")
	proto.RegisterType((*Resolver)(nil), "regen.data.v2.Resolver")
	proto.RegisterType((*DataResolver)(nil), "regen.data.v2.DataResolver")
}

func init() { proto.RegisterFile("regen/data/v2/state.proto", fileDescriptor_fd916028e0fe6d62) }

var fileDescriptor_fd916028e0fe6d62 = []byte{
	// 435 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xb3, 0x4e, 0x48, 0x93, 0x89, 0x1b, 0x19, 0x83, 0xc4, 0xd6, 0x07, 0xa7, 0xb2, 0x04,
	0xea, 0xc1, 0x78, 0x45, 0xb8, 0xa0, 0xdc, 0x40, 0x05, 0xa9, 0x57, 0x8b, 0x13, 0x1c, 0xd0, 0x26,
	0xde, 0xba, 0x2b, 0x6c, 0x6f, 0xb5, 0x5e, 0x1b, 0x9e, 0x02, 0x71, 0x47, 0xe2, 0x79, 0x38, 0x56,
	0xe2, 0xc2, 0x11, 0x25, 0x6f, 0xc0, 0x13, 0x20, 0x8f, 0xed, 0xb4, 0x88, 0x5e, 0x7a, 0xb2, 0xc7,
	0xf3, 0xfb, 0xff, 0x7e, 0xfd, 0x1a, 0x38, 0xd2, 0x22, 0x15, 0x05, 0x4b, 0xb8, 0xe1, 0xac, 0x5e,
	0xb2, 0xd2, 0x70, 0x23, 0xa2, 0x4b, 0xad, 0x8c, 0x72, 0x0f, 0x71, 0x15, 0x35, 0xab, 0xa8, 0x5e,
	0x7a, 0x8f, 0x36, 0xaa, 0xcc, 0x55, 0xc9, 0x94, 0xce, 0x59, 0xfd, 0xac, 0x79, 0xb4, 0x3a, 0x6f,
	0x91, 0x2a, 0x95, 0x66, 0x82, 0xe1, 0xb4, 0xae, 0xce, 0x99, 0x91, 0xb9, 0x28, 0x0d, 0xcf, 0x2f,
	0x5b, 0x41, 0xf0, 0x1a, 0xc6, 0xa7, 0xdc, 0xf0, 0xb3, 0x53, 0x77, 0x0e, 0x96, 0x4c, 0x28, 0x39,
	0x26, 0x27, 0x76, 0x6c, 0xc9, 0xc4, 0x75, 0x60, 0x28, 0xb5, 0xa4, 0xd6, 0x31, 0x39, 0x99, 0xc6,
	0xcd, 0xeb, 0xea, 0xe8, 0xcf, 0xf7, 0x9f, 0x5f, 0x86, 0x0f, 0x60, 0xd4, 0x28, 0xdd, 0x29, 0xee,
	0x1d, 0x42, 0x09, 0x25, 0xc1, 0x39, 0x40, 0x63, 0xf3, 0xb2, 0xd8, 0x5c, 0x28, 0xfd, 0x9f, 0xd5,
	0x0b, 0x98, 0xee, 0xb9, 0x68, 0x38, 0x5b, 0x7a, 0x51, 0x9b, 0x2c, 0xea, 0x93, 0x45, 0x6f, 0x7b,
	0x45, 0x7c, 0x2d, 0x5e, 0xcd, 0x11, 0x39, 0x69, 0x91, 0xd4, 0x0a, 0xbe, 0x11, 0xb0, 0x11, 0x64,
	0x8c, 0x28, 0xcd, 0x2d, 0x28, 0x0f, 0x26, 0xbc, 0xdb, 0x21, 0xc9, 0x8e, 0xf7, 0xf3, 0xbf, 0x31,
	0x86, 0x77, 0x89, 0xf1, 0x18, 0x63, 0x2c, 0xe0, 0x10, 0x66, 0x32, 0x09, 0xf7, 0x86, 0xf6, 0x35,
	0xcc, 0x21, 0x74, 0x18, 0x18, 0x98, 0xc4, 0xa2, 0x54, 0x59, 0x2d, 0x6e, 0x06, 0x1b, 0xf5, 0x75,
	0x56, 0x3a, 0xeb, 0xeb, 0xac, 0x74, 0xe6, 0x52, 0x38, 0xc8, 0x79, 0xc1, 0x53, 0xa1, 0x31, 0x8c,
	0x1d, 0xf7, 0xe3, 0x2a, 0x44, 0xdc, 0x13, 0x18, 0x37, 0x1e, 0x0e, 0x71, 0x0f, 0xf0, 0x5f, 0x87,
	0xb8, 0xf7, 0x61, 0x56, 0xe9, 0x2c, 0xec, 0x74, 0x8e, 0x45, 0x09, 0x1d, 0x05, 0xef, 0xdb, 0x4a,
	0x6e, 0x21, 0xb7, 0x95, 0x2c, 0x60, 0xa6, 0xbb, 0xdd, 0x87, 0xa6, 0x44, 0x8c, 0x04, 0xfd, 0xa7,
	0xb3, 0x64, 0xe5, 0x21, 0xee, 0x21, 0x38, 0x30, 0x97, 0x49, 0x78, 0x53, 0x7b, 0xef, 0xd5, 0x9b,
	0x1f, 0x5b, 0x9f, 0x5c, 0x6d, 0x7d, 0xf2, 0x7b, 0xeb, 0x93, 0xaf, 0x3b, 0x7f, 0x70, 0xb5, 0xf3,
	0x07, 0xbf, 0x76, 0xfe, 0xe0, 0x5d, 0x98, 0x4a, 0x73, 0x51, 0xad, 0xa3, 0x8d, 0xca, 0x19, 0x5e,
	0xe3, 0xd3, 0x42, 0x98, 0x4f, 0x4a, 0x7f, 0xec, 0xa6, 0x4c, 0x24, 0xa9, 0xd0, 0xec, 0x33, 0xde,
	0xef, 0x7a, 0x8c, 0x05, 0x3f, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x02, 0xe5, 0xfa, 0xd5, 0xd4,
	0x02, 0x00, 0x00,
}

func (m *DataID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Iri) > 0 {
		i -= len(m.Iri)
		copy(dAtA[i:], m.Iri)
		i = encodeVarintState(dAtA, i, uint64(len(m.Iri)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataAnchor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataAnchor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataAnchor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DataAttestor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataAttestor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataAttestor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Timestamp != nil {
		{
			size, err := m.Timestamp.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Attestor) > 0 {
		i -= len(m.Attestor)
		copy(dAtA[i:], m.Attestor)
		i = encodeVarintState(dAtA, i, uint64(len(m.Attestor)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Resolver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Resolver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Resolver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Manager) > 0 {
		i -= len(m.Manager)
		copy(dAtA[i:], m.Manager)
		i = encodeVarintState(dAtA, i, uint64(len(m.Manager)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Url) > 0 {
		i -= len(m.Url)
		copy(dAtA[i:], m.Url)
		i = encodeVarintState(dAtA, i, uint64(len(m.Url)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DataResolver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataResolver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataResolver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ResolverId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.ResolverId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintState(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DataID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Iri)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *DataAnchor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *DataAttestor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Attestor)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.Timestamp != nil {
		l = m.Timestamp.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *Resolver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Manager)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *DataResolver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.ResolverId != 0 {
		n += 1 + sovState(uint64(m.ResolverId))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DataID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: DataID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Iri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Iri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *DataAnchor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: DataAnchor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataAnchor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *DataAttestor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: DataAttestor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataAttestor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Attestor", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Attestor = append(m.Attestor[:0], dAtA[iNdEx:postIndex]...)
			if m.Attestor == nil {
				m.Attestor = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Timestamp == nil {
				m.Timestamp = &types.Timestamp{}
			}
			if err := m.Timestamp.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *Resolver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Resolver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Resolver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Manager", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Manager = append(m.Manager[:0], dAtA[iNdEx:postIndex]...)
			if m.Manager == nil {
				m.Manager = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *DataResolver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: DataResolver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataResolver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = append(m.Id[:0], dAtA[iNdEx:postIndex]...)
			if m.Id == nil {
				m.Id = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolverId", wireType)
			}
			m.ResolverId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolverId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
