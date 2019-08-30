// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/worker/common/common.proto

package common

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
	pfs "github.com/pachyderm/pachyderm/src/client/pfs"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type State int32

const (
	State_RUNNING State = 0
	State_SUCCESS State = 1
	State_FAILED  State = 3
)

var State_name = map[int32]string{
	0: "RUNNING",
	1: "SUCCESS",
	3: "FAILED",
}

var State_value = map[string]int32{
	"RUNNING": 0,
	"SUCCESS": 1,
	"FAILED":  3,
}

func (x State) String() string {
	return proto.EnumName(State_name, int32(x))
}

func (State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{0}
}

type Input struct {
	FileInfo             *pfs.FileInfo `protobuf:"bytes,1,opt,name=file_info,json=fileInfo,proto3" json:"file_info,omitempty"`
	ParentCommit         *pfs.Commit   `protobuf:"bytes,5,opt,name=parent_commit,json=parentCommit,proto3" json:"parent_commit,omitempty"`
	Name                 string        `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Lazy                 bool          `protobuf:"varint,3,opt,name=lazy,proto3" json:"lazy,omitempty"`
	Branch               string        `protobuf:"bytes,4,opt,name=branch,proto3" json:"branch,omitempty"`
	GitURL               string        `protobuf:"bytes,6,opt,name=git_url,json=gitUrl,proto3" json:"git_url,omitempty"`
	EmptyFiles           bool          `protobuf:"varint,7,opt,name=empty_files,json=emptyFiles,proto3" json:"empty_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}
func (*Input) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{0}
}
func (m *Input) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Input) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Input.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Input) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Input.Merge(m, src)
}
func (m *Input) XXX_Size() int {
	return m.Size()
}
func (m *Input) XXX_DiscardUnknown() {
	xxx_messageInfo_Input.DiscardUnknown(m)
}

var xxx_messageInfo_Input proto.InternalMessageInfo

func (m *Input) GetFileInfo() *pfs.FileInfo {
	if m != nil {
		return m.FileInfo
	}
	return nil
}

func (m *Input) GetParentCommit() *pfs.Commit {
	if m != nil {
		return m.ParentCommit
	}
	return nil
}

func (m *Input) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Input) GetLazy() bool {
	if m != nil {
		return m.Lazy
	}
	return false
}

func (m *Input) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Input) GetGitURL() string {
	if m != nil {
		return m.GitURL
	}
	return ""
}

func (m *Input) GetEmptyFiles() bool {
	if m != nil {
		return m.EmptyFiles
	}
	return false
}

type ChunkState struct {
	State   State  `protobuf:"varint,1,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	DatumID string `protobuf:"bytes,2,opt,name=datum_id,json=datumId,proto3" json:"datum_id,omitempty"`
	// The IP address of the worker who processed this chunk
	Address              string   `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkState) Reset()         { *m = ChunkState{} }
func (m *ChunkState) String() string { return proto.CompactTextString(m) }
func (*ChunkState) ProtoMessage()    {}
func (*ChunkState) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{1}
}
func (m *ChunkState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChunkState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChunkState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChunkState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkState.Merge(m, src)
}
func (m *ChunkState) XXX_Size() int {
	return m.Size()
}
func (m *ChunkState) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkState.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkState proto.InternalMessageInfo

func (m *ChunkState) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RUNNING
}

func (m *ChunkState) GetDatumID() string {
	if m != nil {
		return m.DatumID
	}
	return ""
}

func (m *ChunkState) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MergeState struct {
	State                State       `protobuf:"varint,1,opt,name=state,proto3,enum=common.State" json:"state,omitempty"`
	Tree                 *pfs.Object `protobuf:"bytes,2,opt,name=tree,proto3" json:"tree,omitempty"`
	SizeBytes            uint64      `protobuf:"varint,3,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	StatsTree            *pfs.Object `protobuf:"bytes,4,opt,name=stats_tree,json=statsTree,proto3" json:"stats_tree,omitempty"`
	StatsSizeBytes       uint64      `protobuf:"varint,5,opt,name=stats_size_bytes,json=statsSizeBytes,proto3" json:"stats_size_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MergeState) Reset()         { *m = MergeState{} }
func (m *MergeState) String() string { return proto.CompactTextString(m) }
func (*MergeState) ProtoMessage()    {}
func (*MergeState) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{2}
}
func (m *MergeState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MergeState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MergeState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MergeState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeState.Merge(m, src)
}
func (m *MergeState) XXX_Size() int {
	return m.Size()
}
func (m *MergeState) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeState.DiscardUnknown(m)
}

var xxx_messageInfo_MergeState proto.InternalMessageInfo

func (m *MergeState) GetState() State {
	if m != nil {
		return m.State
	}
	return State_RUNNING
}

func (m *MergeState) GetTree() *pfs.Object {
	if m != nil {
		return m.Tree
	}
	return nil
}

func (m *MergeState) GetSizeBytes() uint64 {
	if m != nil {
		return m.SizeBytes
	}
	return 0
}

func (m *MergeState) GetStatsTree() *pfs.Object {
	if m != nil {
		return m.StatsTree
	}
	return nil
}

func (m *MergeState) GetStatsSizeBytes() uint64 {
	if m != nil {
		return m.StatsSizeBytes
	}
	return 0
}

type Plan struct {
	Chunks               []int64  `protobuf:"varint,1,rep,packed,name=chunks,proto3" json:"chunks,omitempty"`
	Merges               int64    `protobuf:"varint,2,opt,name=merges,proto3" json:"merges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Plan) Reset()         { *m = Plan{} }
func (m *Plan) String() string { return proto.CompactTextString(m) }
func (*Plan) ProtoMessage()    {}
func (*Plan) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb6c79ddd9db74, []int{3}
}
func (m *Plan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Plan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Plan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Plan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Plan.Merge(m, src)
}
func (m *Plan) XXX_Size() int {
	return m.Size()
}
func (m *Plan) XXX_DiscardUnknown() {
	xxx_messageInfo_Plan.DiscardUnknown(m)
}

var xxx_messageInfo_Plan proto.InternalMessageInfo

func (m *Plan) GetChunks() []int64 {
	if m != nil {
		return m.Chunks
	}
	return nil
}

func (m *Plan) GetMerges() int64 {
	if m != nil {
		return m.Merges
	}
	return 0
}

func init() {
	proto.RegisterEnum("common.State", State_name, State_value)
	proto.RegisterType((*Input)(nil), "common.Input")
	proto.RegisterType((*ChunkState)(nil), "common.ChunkState")
	proto.RegisterType((*MergeState)(nil), "common.MergeState")
	proto.RegisterType((*Plan)(nil), "common.Plan")
}

func init() { proto.RegisterFile("server/worker/common/common.proto", fileDescriptor_91fb6c79ddd9db74) }

var fileDescriptor_91fb6c79ddd9db74 = []byte{
	// 525 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xdd, 0x8a, 0xd3, 0x40,
	0x14, 0xde, 0xb1, 0x69, 0xd2, 0x9e, 0xda, 0xa5, 0x0c, 0x8b, 0x84, 0x05, 0xdb, 0xda, 0x05, 0x29,
	0x05, 0x1b, 0x59, 0x41, 0xaf, 0x6d, 0xbb, 0xbb, 0x04, 0xd6, 0x2a, 0x53, 0x7b, 0xe3, 0x4d, 0x48,
	0x93, 0x69, 0x3a, 0x6e, 0xfe, 0x98, 0x99, 0x2a, 0xdd, 0x27, 0xf1, 0x8d, 0xf4, 0xd2, 0x27, 0x58,
	0xa4, 0x5e, 0xfa, 0x12, 0x32, 0x33, 0x29, 0x2c, 0xe8, 0xcd, 0x5e, 0x84, 0x7c, 0x3f, 0x27, 0xdf,
	0x9c, 0x73, 0x92, 0xc0, 0x33, 0x41, 0xf9, 0x17, 0xca, 0xbd, 0xaf, 0x05, 0xbf, 0xa1, 0xdc, 0x8b,
	0x8a, 0x2c, 0x2b, 0xf2, 0xea, 0x36, 0x2e, 0x79, 0x21, 0x0b, 0x6c, 0x1b, 0x76, 0x7a, 0x12, 0xa5,
	0x8c, 0xe6, 0xd2, 0x2b, 0xd7, 0x42, 0x5d, 0xc6, 0x3d, 0x3d, 0x49, 0x8a, 0xa4, 0xd0, 0xd0, 0x53,
	0xc8, 0xa8, 0x83, 0x3f, 0x08, 0xea, 0x7e, 0x5e, 0x6e, 0x25, 0x1e, 0x41, 0x73, 0xcd, 0x52, 0x1a,
	0xb0, 0x7c, 0x5d, 0xb8, 0xa8, 0x8f, 0x86, 0xad, 0xf3, 0xf6, 0x58, 0x3d, 0x7e, 0xc9, 0x52, 0xea,
	0xe7, 0xeb, 0x82, 0x34, 0xd6, 0x15, 0xc2, 0x2f, 0xa1, 0x5d, 0x86, 0x9c, 0xe6, 0x32, 0x50, 0x47,
	0x32, 0xe9, 0xd6, 0x75, 0x7d, 0x4b, 0xd7, 0x4f, 0xb5, 0x44, 0x1e, 0x9b, 0x0a, 0xc3, 0x30, 0x06,
	0x2b, 0x0f, 0x33, 0xea, 0x3e, 0xea, 0xa3, 0x61, 0x93, 0x68, 0xac, 0xb4, 0x34, 0xbc, 0xdd, 0xb9,
	0xb5, 0x3e, 0x1a, 0x36, 0x88, 0xc6, 0xf8, 0x09, 0xd8, 0x2b, 0x1e, 0xe6, 0xd1, 0xc6, 0xb5, 0x74,
	0x65, 0xc5, 0xf0, 0x19, 0x38, 0x09, 0x93, 0xc1, 0x96, 0xa7, 0xae, 0xad, 0x8c, 0x09, 0xec, 0xef,
	0x7a, 0xf6, 0x15, 0x93, 0x4b, 0x72, 0x4d, 0xec, 0x84, 0xc9, 0x25, 0x4f, 0x71, 0x0f, 0x5a, 0x34,
	0x2b, 0xe5, 0x2e, 0x50, 0x8d, 0x0a, 0xd7, 0xd1, 0xb9, 0xa0, 0x25, 0x35, 0x84, 0x18, 0x08, 0x80,
	0xe9, 0x66, 0x9b, 0xdf, 0x2c, 0x64, 0x28, 0x29, 0x3e, 0x83, 0xba, 0x50, 0x40, 0x4f, 0x7b, 0x7c,
	0xde, 0x1e, 0x57, 0xdb, 0xd4, 0x2e, 0x31, 0x1e, 0x7e, 0x0e, 0x8d, 0x38, 0x94, 0xdb, 0x2c, 0x60,
	0xb1, 0x69, 0x7e, 0xd2, 0xda, 0xdf, 0xf5, 0x9c, 0x99, 0xd2, 0xfc, 0x19, 0x71, 0xb4, 0xe9, 0xc7,
	0xd8, 0x05, 0x27, 0x8c, 0x63, 0x4e, 0x85, 0xd0, 0xf3, 0x34, 0xc9, 0x81, 0x0e, 0xbe, 0x23, 0x80,
	0x77, 0x94, 0x27, 0xf4, 0x01, 0xa7, 0xf6, 0xc0, 0x92, 0x9c, 0x9a, 0x75, 0x1d, 0xf6, 0xfa, 0x7e,
	0xf5, 0x99, 0x46, 0x92, 0x68, 0x03, 0x3f, 0x05, 0x10, 0xec, 0x96, 0x06, 0xab, 0x9d, 0xa4, 0xe6,
	0x44, 0x8b, 0x34, 0x95, 0x32, 0x51, 0x02, 0x1e, 0x01, 0xa8, 0x20, 0x11, 0xe8, 0x14, 0xeb, 0xdf,
	0x94, 0xa6, 0xb6, 0x3f, 0xaa, 0xa8, 0x21, 0x74, 0x4c, 0xed, 0xbd, 0xc0, 0xba, 0x0e, 0x3c, 0xd6,
	0xfa, 0xe2, 0x90, 0x3a, 0x78, 0x0d, 0xd6, 0x87, 0x34, 0xcc, 0xd5, 0x4b, 0x8a, 0xd4, 0x1a, 0x85,
	0x8b, 0xfa, 0xb5, 0x61, 0x8d, 0x54, 0x4c, 0xe9, 0x99, 0x1a, 0x54, 0xe8, 0xbe, 0x6b, 0xa4, 0x62,
	0xa3, 0x17, 0x50, 0x37, 0xb3, 0xb7, 0xc0, 0x21, 0xcb, 0xf9, 0xdc, 0x9f, 0x5f, 0x75, 0x8e, 0x14,
	0x59, 0x2c, 0xa7, 0xd3, 0x8b, 0xc5, 0xa2, 0x83, 0x30, 0x80, 0x7d, 0xf9, 0xd6, 0xbf, 0xbe, 0x98,
	0x75, 0x6a, 0x13, 0xff, 0xc7, 0xbe, 0x8b, 0x7e, 0xee, 0xbb, 0xe8, 0xd7, 0xbe, 0x8b, 0xbe, 0xfd,
	0xee, 0x1e, 0x7d, 0x7a, 0x93, 0x30, 0xb9, 0xd9, 0xae, 0xd4, 0xaa, 0xbc, 0x32, 0x8c, 0x36, 0xbb,
	0x98, 0xf2, 0xfb, 0x48, 0xf0, 0xc8, 0xfb, 0xdf, 0xff, 0xb1, 0xb2, 0xf5, 0x57, 0xfe, 0xea, 0x6f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x86, 0x68, 0x20, 0x3e, 0x03, 0x00, 0x00,
}

func (m *Input) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Input) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Input) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.EmptyFiles {
		i--
		if m.EmptyFiles {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.GitURL) > 0 {
		i -= len(m.GitURL)
		copy(dAtA[i:], m.GitURL)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.GitURL)))
		i--
		dAtA[i] = 0x32
	}
	if m.ParentCommit != nil {
		{
			size, err := m.ParentCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Branch) > 0 {
		i -= len(m.Branch)
		copy(dAtA[i:], m.Branch)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Branch)))
		i--
		dAtA[i] = 0x22
	}
	if m.Lazy {
		i--
		if m.Lazy {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.FileInfo != nil {
		{
			size, err := m.FileInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ChunkState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChunkState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChunkState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DatumID) > 0 {
		i -= len(m.DatumID)
		copy(dAtA[i:], m.DatumID)
		i = encodeVarintCommon(dAtA, i, uint64(len(m.DatumID)))
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MergeState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MergeState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MergeState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.StatsSizeBytes != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.StatsSizeBytes))
		i--
		dAtA[i] = 0x28
	}
	if m.StatsTree != nil {
		{
			size, err := m.StatsTree.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.SizeBytes != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.SizeBytes))
		i--
		dAtA[i] = 0x18
	}
	if m.Tree != nil {
		{
			size, err := m.Tree.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintCommon(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Plan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Plan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Plan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Merges != 0 {
		i = encodeVarintCommon(dAtA, i, uint64(m.Merges))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Chunks) > 0 {
		dAtA6 := make([]byte, len(m.Chunks)*10)
		var j5 int
		for _, num1 := range m.Chunks {
			num := uint64(num1)
			for num >= 1<<7 {
				dAtA6[j5] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j5++
			}
			dAtA6[j5] = uint8(num)
			j5++
		}
		i -= j5
		copy(dAtA[i:], dAtA6[:j5])
		i = encodeVarintCommon(dAtA, i, uint64(j5))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Input) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FileInfo != nil {
		l = m.FileInfo.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.Lazy {
		n += 2
	}
	l = len(m.Branch)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.ParentCommit != nil {
		l = m.ParentCommit.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.GitURL)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.EmptyFiles {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChunkState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovCommon(uint64(m.State))
	}
	l = len(m.DatumID)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *MergeState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovCommon(uint64(m.State))
	}
	if m.Tree != nil {
		l = m.Tree.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.SizeBytes != 0 {
		n += 1 + sovCommon(uint64(m.SizeBytes))
	}
	if m.StatsTree != nil {
		l = m.StatsTree.Size()
		n += 1 + l + sovCommon(uint64(l))
	}
	if m.StatsSizeBytes != 0 {
		n += 1 + sovCommon(uint64(m.StatsSizeBytes))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Plan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Chunks) > 0 {
		l = 0
		for _, e := range m.Chunks {
			l += sovCommon(uint64(e))
		}
		n += 1 + sovCommon(uint64(l)) + l
	}
	if m.Merges != 0 {
		n += 1 + sovCommon(uint64(m.Merges))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Input) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Input: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Input: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FileInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.FileInfo == nil {
				m.FileInfo = &pfs.FileInfo{}
			}
			if err := m.FileInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lazy", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Lazy = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Branch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ParentCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ParentCommit == nil {
				m.ParentCommit = &pfs.Commit{}
			}
			if err := m.ParentCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GitURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GitURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EmptyFiles", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.EmptyFiles = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *ChunkState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: ChunkState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChunkState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DatumID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DatumID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *MergeState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: MergeState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MergeState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= State(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Tree == nil {
				m.Tree = &pfs.Object{}
			}
			if err := m.Tree.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SizeBytes", wireType)
			}
			m.SizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SizeBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsTree", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
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
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StatsTree == nil {
				m.StatsTree = &pfs.Object{}
			}
			if err := m.StatsTree.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StatsSizeBytes", wireType)
			}
			m.StatsSizeBytes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StatsSizeBytes |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func (m *Plan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
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
			return fmt.Errorf("proto: Plan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Plan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType == 0 {
				var v int64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Chunks = append(m.Chunks, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthCommon
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthCommon
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Chunks) == 0 {
					m.Chunks = make([]int64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v int64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= int64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Chunks = append(m.Chunks, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Chunks", wireType)
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Merges", wireType)
			}
			m.Merges = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Merges |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCommon
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
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
					return 0, ErrIntOverflowCommon
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
				return 0, ErrInvalidLengthCommon
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
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
				next, err := skipCommon(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthCommon
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
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)
