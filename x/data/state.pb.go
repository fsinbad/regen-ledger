// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v1/state.proto

package data

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
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
	return fileDescriptor_29cc90d94a9e9542, []int{0}
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
	return fileDescriptor_29cc90d94a9e9542, []int{1}
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
	return fileDescriptor_29cc90d94a9e9542, []int{2}
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
	return fileDescriptor_29cc90d94a9e9542, []int{3}
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
	return fileDescriptor_29cc90d94a9e9542, []int{4}
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
	proto.RegisterType((*DataID)(nil), "regen.data.v1.DataID")
	proto.RegisterType((*DataAnchor)(nil), "regen.data.v1.DataAnchor")
	proto.RegisterType((*DataAttestor)(nil), "regen.data.v1.DataAttestor")
	proto.RegisterType((*Resolver)(nil), "regen.data.v1.Resolver")
	proto.RegisterType((*DataResolver)(nil), "regen.data.v1.DataResolver")
}

func init() { proto.RegisterFile("regen/data/v1/state.proto", fileDescriptor_29cc90d94a9e9542) }

var fileDescriptor_29cc90d94a9e9542 = []byte{
	// 439 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x1c, 0xc5, 0x3b, 0x69, 0xed, 0xb6, 0xff, 0x66, 0x4b, 0x8c, 0x1e, 0x66, 0x03, 0xa6, 0x4b, 0x40,
	0xd9, 0x43, 0xcc, 0x50, 0xbd, 0x48, 0x6f, 0xca, 0x2a, 0xec, 0x35, 0x78, 0xd2, 0x83, 0x4c, 0x9b,
	0xd9, 0x74, 0x30, 0xe9, 0x94, 0xc9, 0x24, 0xfa, 0x29, 0xc4, 0xbb, 0xe0, 0xe7, 0xf1, 0xb8, 0xe0,
	0xc5, 0xa3, 0xb4, 0xdf, 0xc0, 0x4f, 0x20, 0xf9, 0x27, 0xe9, 0xae, 0xb8, 0x97, 0xbd, 0xf5, 0xdf,
	0xf7, 0xf2, 0x7e, 0x8f, 0xc7, 0xc0, 0x89, 0x16, 0xa9, 0xd8, 0xb0, 0x84, 0x1b, 0xce, 0xaa, 0x39,
	0x2b, 0x0c, 0x37, 0x22, 0xda, 0x6a, 0x65, 0x94, 0x7b, 0x8c, 0x52, 0x54, 0x4b, 0x51, 0x35, 0xf7,
	0x1e, 0xad, 0x54, 0x91, 0xab, 0x82, 0x29, 0x9d, 0xb3, 0x6a, 0xce, 0xb3, 0xed, 0x9a, 0xcf, 0xeb,
	0xa3, 0x71, 0x7b, 0xb3, 0x54, 0xa9, 0x34, 0x13, 0x0c, 0xaf, 0x65, 0x79, 0xc9, 0x8c, 0xcc, 0x45,
	0x61, 0x78, 0xbe, 0x6d, 0x0c, 0xc1, 0x6b, 0x18, 0x9e, 0x73, 0xc3, 0x2f, 0xce, 0xdd, 0x29, 0x58,
	0x32, 0xa1, 0xe4, 0x94, 0x9c, 0xd9, 0xb1, 0x25, 0x13, 0xd7, 0x81, 0xbe, 0xd4, 0x92, 0x5a, 0xa7,
	0xe4, 0x6c, 0x1c, 0xd7, 0x3f, 0x17, 0x27, 0x7f, 0xbe, 0xff, 0xfc, 0xd2, 0x7f, 0x00, 0x83, 0xda,
	0xe9, 0x8e, 0x51, 0x77, 0x08, 0x25, 0x94, 0x04, 0x97, 0x00, 0x75, 0xcc, 0xcb, 0xcd, 0x6a, 0xad,
	0xf4, 0x7f, 0x51, 0x2f, 0x60, 0x7c, 0xe0, 0x62, 0xe0, 0xe4, 0x99, 0x17, 0x35, 0xcd, 0xa2, 0xae,
	0x59, 0xf4, 0xb6, 0x73, 0xc4, 0xd7, 0xe6, 0xc5, 0x14, 0x91, 0xa3, 0x06, 0x49, 0xad, 0xe0, 0x1b,
	0x01, 0x1b, 0x41, 0xc6, 0x88, 0xc2, 0xdc, 0x82, 0xf2, 0x60, 0xc4, 0x5b, 0x0d, 0x49, 0x76, 0x7c,
	0xb8, 0xff, 0xad, 0xd1, 0xbf, 0x4b, 0x8d, 0xc7, 0x58, 0x63, 0x06, 0xc7, 0x30, 0x91, 0x49, 0x78,
	0x08, 0xb4, 0xaf, 0x61, 0x0e, 0xa1, 0xfd, 0xc0, 0xc0, 0x28, 0x16, 0x85, 0xca, 0x2a, 0x71, 0xb3,
	0xd8, 0xa0, 0x9b, 0xb3, 0xd4, 0x59, 0x37, 0x67, 0xa9, 0x33, 0x97, 0xc2, 0x51, 0xce, 0x37, 0x3c,
	0x15, 0x1a, 0xcb, 0xd8, 0x71, 0x77, 0x2e, 0x42, 0xc4, 0x3d, 0x81, 0x61, 0x9d, 0xe1, 0x10, 0xf7,
	0x08, 0xbf, 0x75, 0x88, 0x7b, 0x1f, 0x26, 0xa5, 0xce, 0xc2, 0xd6, 0xe7, 0x58, 0x94, 0xd0, 0x41,
	0xf0, 0xbe, 0x99, 0xe4, 0x16, 0x72, 0x33, 0xc9, 0x0c, 0x26, 0xba, 0xd5, 0x3e, 0xd4, 0x23, 0x62,
	0x25, 0xe8, 0xfe, 0xba, 0x48, 0x16, 0x1e, 0xe2, 0x1e, 0x82, 0x03, 0x53, 0x99, 0x84, 0x37, 0xbd,
	0xf7, 0x5e, 0xbd, 0xf9, 0xb1, 0xf3, 0xc9, 0xd5, 0xce, 0x27, 0xbf, 0x77, 0x3e, 0xf9, 0xba, 0xf7,
	0x7b, 0x57, 0x7b, 0xbf, 0xf7, 0x6b, 0xef, 0xf7, 0xde, 0x85, 0xa9, 0x34, 0xeb, 0x72, 0x19, 0xad,
	0x54, 0xce, 0xf0, 0x4d, 0x3e, 0xdd, 0x08, 0xf3, 0x49, 0xe9, 0x8f, 0xed, 0x95, 0x89, 0x24, 0x15,
	0x9a, 0x7d, 0xc6, 0x57, 0xbc, 0x1c, 0xe2, 0xc0, 0xcf, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0x43,
	0xd3, 0xf6, 0xdb, 0xda, 0x02, 0x00, 0x00,
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
