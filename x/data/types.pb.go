// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/data/v2/types.proto

package data

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

// DigestAlgorithm is the hash digest algorithm
//
// With v2, this enum is no longer validated on-chain.
// However, this enum SHOULD still be used and updated as a registry of known digest
// algorithms and all implementations should coordinate on these values.
type DigestAlgorithm int32

const (
	// unspecified and invalid
	DigestAlgorithm_DIGEST_ALGORITHM_UNSPECIFIED DigestAlgorithm = 0
	// BLAKE2b-256
	DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256 DigestAlgorithm = 1
)

var DigestAlgorithm_name = map[int32]string{
	0: "DIGEST_ALGORITHM_UNSPECIFIED",
	1: "DIGEST_ALGORITHM_BLAKE2B_256",
}

var DigestAlgorithm_value = map[string]int32{
	"DIGEST_ALGORITHM_UNSPECIFIED": 0,
	"DIGEST_ALGORITHM_BLAKE2B_256": 1,
}

func (x DigestAlgorithm) String() string {
	return proto.EnumName(DigestAlgorithm_name, int32(x))
}

func (DigestAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{0}
}

// GraphCanonicalizationAlgorithm is the graph canonicalization algorithm
//
// With v2, this enum is no longer validated on-chain.
// However, this enum SHOULD still be used and updated as a registry of known canonicalization
// algorithms and all implementations should coordinate on these values.
type GraphCanonicalizationAlgorithm int32

const (
	// unspecified and invalid
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED GraphCanonicalizationAlgorithm = 0
	// RDFC 1.0 graph canonicalization algorithm. Essentially the same as URDNA2015 with some
	// small clarifications around escaping of escape characters.
	GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0 GraphCanonicalizationAlgorithm = 1
)

var GraphCanonicalizationAlgorithm_name = map[int32]string{
	0: "GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED",
	1: "GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0",
}

var GraphCanonicalizationAlgorithm_value = map[string]int32{
	"GRAPH_CANONICALIZATION_ALGORITHM_UNSPECIFIED": 0,
	"GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0":    1,
}

func (x GraphCanonicalizationAlgorithm) String() string {
	return proto.EnumName(GraphCanonicalizationAlgorithm_name, int32(x))
}

func (GraphCanonicalizationAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{1}
}

// GraphMerkleTree is the graph merkle tree type used for hashing, if any.
//
// With v2, this enum is no longer validated on-chain.
// However, this enum SHOULD still be used and updated as a registry of known merkle tree
// types and all implementations should coordinate on these values.
type GraphMerkleTree int32

const (
	// unspecified and valid
	GraphMerkleTree_GRAPH_MERKLE_TREE_NONE_UNSPECIFIED GraphMerkleTree = 0
)

var GraphMerkleTree_name = map[int32]string{
	0: "GRAPH_MERKLE_TREE_NONE_UNSPECIFIED",
}

var GraphMerkleTree_value = map[string]int32{
	"GRAPH_MERKLE_TREE_NONE_UNSPECIFIED": 0,
}

func (x GraphMerkleTree) String() string {
	return proto.EnumName(GraphMerkleTree_name, int32(x))
}

func (GraphMerkleTree) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{2}
}

// ContentHash specifies a hash-based content identifier for a piece of data.
// Exactly one of its fields must be set so this message behaves like a oneof.
// A protobuf oneof was not used because this caused compatibility issues with
// amino signing.
type ContentHash struct {
	// raw specifies "raw" data which does not specify a deterministic, canonical
	// encoding. Users of these hashes MUST maintain a copy of the hashed data
	// which is preserved bit by bit. All other content encodings specify a
	// deterministic, canonical encoding allowing implementations to choose from a
	// variety of alternative formats for transport and encoding while maintaining
	// the guarantee that the canonical hash will not change. The media type for
	// "raw" data is defined by the MediaType enum.
	Raw *ContentHash_Raw `protobuf:"bytes,1,opt,name=raw,proto3" json:"raw,omitempty"`
	// graph specifies graph data that conforms to the RDF data model.
	// The canonicalization algorithm used for an RDF graph is specified by
	// GraphCanonicalizationAlgorithm.
	Graph *ContentHash_Graph `protobuf:"bytes,2,opt,name=graph,proto3" json:"graph,omitempty"`
}

func (m *ContentHash) Reset()         { *m = ContentHash{} }
func (m *ContentHash) String() string { return proto.CompactTextString(m) }
func (*ContentHash) ProtoMessage()    {}
func (*ContentHash) Descriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{0}
}
func (m *ContentHash) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash.Merge(m, src)
}
func (m *ContentHash) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash proto.InternalMessageInfo

func (m *ContentHash) GetRaw() *ContentHash_Raw {
	if m != nil {
		return m.Raw
	}
	return nil
}

func (m *ContentHash) GetGraph() *ContentHash_Graph {
	if m != nil {
		return m.Graph
	}
	return nil
}

// RawVis the content hash type used for raw data.
type ContentHash_Raw struct {
	// hash represents the hash of the data based on the specified
	// digest_algorithm. It must be at least 20 bytes long and at most 64 bytes long.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm and should be a non-zero value from the DigestAlgorithm enum.
	DigestAlgorithm uint32 `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3" json:"digest_algorithm,omitempty"`
	// file_extension represents the file extension for raw data. It can be
	// must be between 2-6 characters long, must be all lower-case and should represent
	// the canonical extension for the media type.
	//
	// A list of canonical extensions which should be used is provided here
	// and SHOULD be used by implementations: txt, json, csv, xml, pdf, tiff,
	// jpg, png, svg, webp, avif, gif, apng, mpeg, mp4, webm, ogg, heic, raw.
	//
	// The above list should be updated as new media types come into common usage
	// especially when there are two or more possible extensions (i.e. jpg vs jpeg or tif vs tiff).
	FileExtension string `protobuf:"bytes,3,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
}

func (m *ContentHash_Raw) Reset()         { *m = ContentHash_Raw{} }
func (m *ContentHash_Raw) String() string { return proto.CompactTextString(m) }
func (*ContentHash_Raw) ProtoMessage()    {}
func (*ContentHash_Raw) Descriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{0, 0}
}
func (m *ContentHash_Raw) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash_Raw) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash_Raw.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash_Raw) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash_Raw.Merge(m, src)
}
func (m *ContentHash_Raw) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash_Raw) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash_Raw.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash_Raw proto.InternalMessageInfo

func (m *ContentHash_Raw) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContentHash_Raw) GetDigestAlgorithm() uint32 {
	if m != nil {
		return m.DigestAlgorithm
	}
	return 0
}

func (m *ContentHash_Raw) GetFileExtension() string {
	if m != nil {
		return m.FileExtension
	}
	return ""
}

// Graph is the content hash type used for RDF graph data.
type ContentHash_Graph struct {
	// hash represents the hash of the data based on the specified
	// digest_algorithm. It must be at least 20 bytes long and at most 64 bytes long.
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	// digest_algorithm represents the hash digest algorithm and should be a non-zero value from the DigestAlgorithm enum.
	DigestAlgorithm uint32 `protobuf:"varint,2,opt,name=digest_algorithm,json=digestAlgorithm,proto3" json:"digest_algorithm,omitempty"`
	// graph_canonicalization_algorithm represents the RDF graph
	// canonicalization algorithm and should be a non-zero value from the GraphCanonicalizationAlgorithm enum.
	CanonicalizationAlgorithm uint32 `protobuf:"varint,3,opt,name=canonicalization_algorithm,json=canonicalizationAlgorithm,proto3" json:"canonicalization_algorithm,omitempty"`
	// merkle_tree is the merkle tree type used for the graph hash, if any and should be a value from the GraphMerkleTree enum
	// or left unspecified.
	MerkleTree uint32 `protobuf:"varint,4,opt,name=merkle_tree,json=merkleTree,proto3" json:"merkle_tree,omitempty"`
}

func (m *ContentHash_Graph) Reset()         { *m = ContentHash_Graph{} }
func (m *ContentHash_Graph) String() string { return proto.CompactTextString(m) }
func (*ContentHash_Graph) ProtoMessage()    {}
func (*ContentHash_Graph) Descriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{0, 1}
}
func (m *ContentHash_Graph) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHash_Graph) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHash_Graph.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHash_Graph) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHash_Graph.Merge(m, src)
}
func (m *ContentHash_Graph) XXX_Size() int {
	return m.Size()
}
func (m *ContentHash_Graph) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHash_Graph.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHash_Graph proto.InternalMessageInfo

func (m *ContentHash_Graph) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *ContentHash_Graph) GetDigestAlgorithm() uint32 {
	if m != nil {
		return m.DigestAlgorithm
	}
	return 0
}

func (m *ContentHash_Graph) GetCanonicalizationAlgorithm() uint32 {
	if m != nil {
		return m.CanonicalizationAlgorithm
	}
	return 0
}

func (m *ContentHash_Graph) GetMerkleTree() uint32 {
	if m != nil {
		return m.MerkleTree
	}
	return 0
}

// ContentHashes contains list of content ContentHash.
type ContentHashes struct {
	// data is a list of content hashes which the resolver claims to serve.
	ContentHashes []*ContentHash `protobuf:"bytes,1,rep,name=content_hashes,json=contentHashes,proto3" json:"content_hashes,omitempty"`
}

func (m *ContentHashes) Reset()         { *m = ContentHashes{} }
func (m *ContentHashes) String() string { return proto.CompactTextString(m) }
func (*ContentHashes) ProtoMessage()    {}
func (*ContentHashes) Descriptor() ([]byte, []int) {
	return fileDescriptor_508db95308016166, []int{1}
}
func (m *ContentHashes) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentHashes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentHashes.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentHashes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentHashes.Merge(m, src)
}
func (m *ContentHashes) XXX_Size() int {
	return m.Size()
}
func (m *ContentHashes) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentHashes.DiscardUnknown(m)
}

var xxx_messageInfo_ContentHashes proto.InternalMessageInfo

func (m *ContentHashes) GetContentHashes() []*ContentHash {
	if m != nil {
		return m.ContentHashes
	}
	return nil
}

func init() {
	proto.RegisterEnum("regen.data.v2.DigestAlgorithm", DigestAlgorithm_name, DigestAlgorithm_value)
	proto.RegisterEnum("regen.data.v2.GraphCanonicalizationAlgorithm", GraphCanonicalizationAlgorithm_name, GraphCanonicalizationAlgorithm_value)
	proto.RegisterEnum("regen.data.v2.GraphMerkleTree", GraphMerkleTree_name, GraphMerkleTree_value)
	proto.RegisterType((*ContentHash)(nil), "regen.data.v2.ContentHash")
	proto.RegisterType((*ContentHash_Raw)(nil), "regen.data.v2.ContentHash.Raw")
	proto.RegisterType((*ContentHash_Graph)(nil), "regen.data.v2.ContentHash.Graph")
	proto.RegisterType((*ContentHashes)(nil), "regen.data.v2.ContentHashes")
}

func init() { proto.RegisterFile("regen/data/v2/types.proto", fileDescriptor_508db95308016166) }

var fileDescriptor_508db95308016166 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0x12, 0x41,
	0x1c, 0xc6, 0x19, 0xb7, 0x35, 0xf1, 0x8f, 0x14, 0xb2, 0x27, 0x4a, 0xcc, 0x4a, 0x48, 0x34, 0x94,
	0xb4, 0x0b, 0xae, 0xb1, 0x89, 0x07, 0x0f, 0xcb, 0x32, 0xc0, 0xa6, 0xb0, 0x34, 0x53, 0x7a, 0xe9,
	0x65, 0x32, 0x85, 0x71, 0x77, 0x03, 0xec, 0x92, 0xd9, 0xb1, 0x54, 0x6f, 0x3e, 0x80, 0x89, 0x4f,
	0xe1, 0xb3, 0x78, 0xec, 0xd1, 0xa3, 0x81, 0x17, 0x31, 0xcc, 0x5a, 0xa5, 0x8d, 0xe8, 0xc1, 0xdb,
	0xee, 0x37, 0xbf, 0xef, 0xff, 0xcd, 0x7c, 0xc9, 0x1f, 0xf6, 0x05, 0xf7, 0x79, 0x54, 0x1f, 0x33,
	0xc9, 0xea, 0x57, 0x56, 0x5d, 0xbe, 0x9f, 0xf3, 0xc4, 0x9c, 0x8b, 0x58, 0xc6, 0x7a, 0x4e, 0x1d,
	0x99, 0xeb, 0x23, 0xf3, 0xca, 0xaa, 0x7c, 0xd2, 0x20, 0xeb, 0xc4, 0x91, 0xe4, 0x91, 0xec, 0xb2,
	0x24, 0xd0, 0x1b, 0xa0, 0x09, 0xb6, 0x28, 0xa2, 0x32, 0xaa, 0x66, 0x2d, 0xc3, 0xbc, 0x03, 0x9b,
	0x1b, 0xa0, 0x49, 0xd8, 0x82, 0xac, 0x51, 0xfd, 0x18, 0x76, 0x7d, 0xc1, 0xe6, 0x41, 0xf1, 0x81,
	0xf2, 0x94, 0xff, 0xe2, 0xe9, 0xac, 0x39, 0x92, 0xe2, 0xa5, 0x09, 0x68, 0x84, 0x2d, 0x74, 0x1d,
	0x76, 0x02, 0x96, 0x04, 0x2a, 0xf1, 0x31, 0x51, 0xdf, 0xfa, 0x01, 0x14, 0xc6, 0xa1, 0xcf, 0x13,
	0x49, 0xd9, 0xd4, 0x8f, 0x45, 0x28, 0x83, 0x99, 0x9a, 0x9e, 0x23, 0xf9, 0x54, 0xb7, 0x6f, 0x65,
	0xfd, 0x19, 0xec, 0xbd, 0x0d, 0xa7, 0x9c, 0xf2, 0x6b, 0xc9, 0xa3, 0x24, 0x8c, 0xa3, 0xa2, 0x56,
	0x46, 0xd5, 0x47, 0x24, 0xb7, 0x56, 0xf1, 0xad, 0x58, 0xfa, 0x82, 0x60, 0x57, 0xa5, 0xff, 0x6f,
	0xde, 0x1b, 0x28, 0x8d, 0x58, 0x14, 0x47, 0xe1, 0x88, 0x4d, 0xc3, 0x0f, 0x4c, 0x86, 0x71, 0xb4,
	0x61, 0xd2, 0x94, 0x69, 0xff, 0x3e, 0xf1, 0xdb, 0xfe, 0x14, 0xb2, 0x33, 0x2e, 0x26, 0x53, 0x4e,
	0xa5, 0xe0, 0xbc, 0xb8, 0xa3, 0x78, 0x48, 0xa5, 0xa1, 0xe0, 0xbc, 0x42, 0x20, 0xb7, 0xd1, 0x18,
	0x4f, 0x74, 0x1b, 0xf6, 0x46, 0xa9, 0x40, 0x03, 0xa5, 0x14, 0x51, 0x59, 0xab, 0x66, 0xad, 0xd2,
	0xf6, 0x9e, 0x49, 0x6e, 0xb4, 0x39, 0xa2, 0x76, 0x0e, 0xf9, 0xd6, 0xbd, 0x67, 0x94, 0xe1, 0x49,
	0xcb, 0xed, 0xe0, 0xb3, 0x21, 0xb5, 0x7b, 0x9d, 0x01, 0x71, 0x87, 0xdd, 0x3e, 0x3d, 0xf7, 0xce,
	0x4e, 0xb1, 0xe3, 0xb6, 0x5d, 0xdc, 0x2a, 0x64, 0xfe, 0x48, 0x34, 0x7b, 0xf6, 0x09, 0xb6, 0x9a,
	0xd4, 0x7a, 0x75, 0x5c, 0x40, 0xb5, 0x8f, 0x08, 0x0c, 0xd5, 0xa9, 0xb3, 0xf5, 0xb9, 0x0d, 0x38,
	0xec, 0x10, 0xfb, 0xb4, 0x4b, 0x1d, 0xdb, 0x1b, 0x78, 0xae, 0x63, 0xf7, 0xdc, 0x0b, 0x7b, 0xe8,
	0x0e, 0xbc, 0xad, 0xb1, 0x47, 0x70, 0xf0, 0x4f, 0x07, 0x69, 0xb5, 0x1d, 0xfa, 0x82, 0x36, 0x0a,
	0xa8, 0xf6, 0x1a, 0xf2, 0xea, 0x0a, 0xfd, 0x5f, 0x0d, 0xea, 0xcf, 0xa1, 0x92, 0x4e, 0xe8, 0x63,
	0x72, 0xd2, 0xc3, 0x74, 0x48, 0x30, 0xa6, 0xde, 0xc0, 0xc3, 0x77, 0x93, 0x9a, 0xed, 0xaf, 0x4b,
	0x03, 0xdd, 0x2c, 0x0d, 0xf4, 0x7d, 0x69, 0xa0, 0xcf, 0x2b, 0x23, 0x73, 0xb3, 0x32, 0x32, 0xdf,
	0x56, 0x46, 0xe6, 0xe2, 0xd0, 0x0f, 0x65, 0xf0, 0xee, 0xd2, 0x1c, 0xc5, 0xb3, 0xba, 0x2a, 0xf9,
	0x28, 0xe2, 0x72, 0x11, 0x8b, 0xc9, 0xcf, 0xbf, 0x29, 0x1f, 0xfb, 0x5c, 0xd4, 0xaf, 0xd5, 0x7e,
	0x5d, 0x3e, 0x54, 0x7b, 0xf5, 0xf2, 0x47, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7a, 0x45, 0x83, 0x99,
	0x74, 0x03, 0x00, 0x00,
}

func (m *ContentHash) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Graph != nil {
		{
			size, err := m.Graph.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Raw != nil {
		{
			size, err := m.Raw.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContentHash_Raw) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash_Raw) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Raw) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FileExtension) > 0 {
		i -= len(m.FileExtension)
		copy(dAtA[i:], m.FileExtension)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.FileExtension)))
		i--
		dAtA[i] = 0x1a
	}
	if m.DigestAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DigestAlgorithm))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContentHash_Graph) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHash_Graph) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHash_Graph) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MerkleTree != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MerkleTree))
		i--
		dAtA[i] = 0x20
	}
	if m.CanonicalizationAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.CanonicalizationAlgorithm))
		i--
		dAtA[i] = 0x18
	}
	if m.DigestAlgorithm != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.DigestAlgorithm))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContentHashes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentHashes) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentHashes) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ContentHashes) > 0 {
		for iNdEx := len(m.ContentHashes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ContentHashes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContentHash) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Raw != nil {
		l = m.Raw.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Graph != nil {
		l = m.Graph.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ContentHash_Raw) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.DigestAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.DigestAlgorithm))
	}
	l = len(m.FileExtension)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ContentHash_Graph) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.DigestAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.DigestAlgorithm))
	}
	if m.CanonicalizationAlgorithm != 0 {
		n += 1 + sovTypes(uint64(m.CanonicalizationAlgorithm))
	}
	if m.MerkleTree != 0 {
		n += 1 + sovTypes(uint64(m.MerkleTree))
	}
	return n
}

func (m *ContentHashes) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ContentHashes) > 0 {
		for _, e := range m.ContentHashes {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContentHash) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ContentHash: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentHash: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Raw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Raw == nil {
				m.Raw = &ContentHash_Raw{}
			}
			if err := m.Raw.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Graph", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Graph == nil {
				m.Graph = &ContentHash_Graph{}
			}
			if err := m.Graph.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ContentHash_Raw) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Raw: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Raw: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileExtension", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FileExtension = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ContentHash_Graph) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Graph: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Graph: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = append(m.Hash[:0], dAtA[iNdEx:postIndex]...)
			if m.Hash == nil {
				m.Hash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DigestAlgorithm", wireType)
			}
			m.DigestAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DigestAlgorithm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CanonicalizationAlgorithm", wireType)
			}
			m.CanonicalizationAlgorithm = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CanonicalizationAlgorithm |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleTree", wireType)
			}
			m.MerkleTree = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MerkleTree |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ContentHashes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ContentHashes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentHashes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContentHashes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContentHashes = append(m.ContentHashes, &ContentHash{})
			if err := m.ContentHashes[len(m.ContentHashes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
