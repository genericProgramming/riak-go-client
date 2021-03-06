// Code generated by protoc-gen-go.
// source: riak.proto
// DO NOT EDIT!

/*
Package riak is a generated protocol buffer package.

It is generated from these files:
	riak.proto

It has these top-level messages:
	RpbErrorResp
	RpbGetServerInfoResp
	RpbPair
	RpbGetBucketReq
	RpbGetBucketResp
	RpbSetBucketReq
	RpbResetBucketReq
	RpbGetBucketTypeReq
	RpbSetBucketTypeReq
	RpbModFun
	RpbCommitHook
	RpbBucketProps
	RpbAuthReq
*/
package riak

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Used by riak_repl bucket fixup
type RpbBucketProps_RpbReplMode int32

const (
	RpbBucketProps_FALSE    RpbBucketProps_RpbReplMode = 0
	RpbBucketProps_REALTIME RpbBucketProps_RpbReplMode = 1
	RpbBucketProps_FULLSYNC RpbBucketProps_RpbReplMode = 2
	RpbBucketProps_TRUE     RpbBucketProps_RpbReplMode = 3
)

var RpbBucketProps_RpbReplMode_name = map[int32]string{
	0: "FALSE",
	1: "REALTIME",
	2: "FULLSYNC",
	3: "TRUE",
}
var RpbBucketProps_RpbReplMode_value = map[string]int32{
	"FALSE":    0,
	"REALTIME": 1,
	"FULLSYNC": 2,
	"TRUE":     3,
}

func (x RpbBucketProps_RpbReplMode) Enum() *RpbBucketProps_RpbReplMode {
	p := new(RpbBucketProps_RpbReplMode)
	*p = x
	return p
}
func (x RpbBucketProps_RpbReplMode) String() string {
	return proto.EnumName(RpbBucketProps_RpbReplMode_name, int32(x))
}
func (x *RpbBucketProps_RpbReplMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(RpbBucketProps_RpbReplMode_value, data, "RpbBucketProps_RpbReplMode")
	if err != nil {
		return err
	}
	*x = RpbBucketProps_RpbReplMode(value)
	return nil
}
func (RpbBucketProps_RpbReplMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{11, 0}
}

// Error response - may be generated for any Req
type RpbErrorResp struct {
	Errmsg           []byte  `protobuf:"bytes,1,req,name=errmsg" json:"errmsg,omitempty"`
	Errcode          *uint32 `protobuf:"varint,2,req,name=errcode" json:"errcode,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbErrorResp) Reset()                    { *m = RpbErrorResp{} }
func (m *RpbErrorResp) String() string            { return proto.CompactTextString(m) }
func (*RpbErrorResp) ProtoMessage()               {}
func (*RpbErrorResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RpbErrorResp) GetErrmsg() []byte {
	if m != nil {
		return m.Errmsg
	}
	return nil
}

func (m *RpbErrorResp) GetErrcode() uint32 {
	if m != nil && m.Errcode != nil {
		return *m.Errcode
	}
	return 0
}

// Get server info request - no message defined, just send RpbGetServerInfoReq message code
type RpbGetServerInfoResp struct {
	Node             []byte `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	ServerVersion    []byte `protobuf:"bytes,2,opt,name=server_version" json:"server_version,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetServerInfoResp) Reset()                    { *m = RpbGetServerInfoResp{} }
func (m *RpbGetServerInfoResp) String() string            { return proto.CompactTextString(m) }
func (*RpbGetServerInfoResp) ProtoMessage()               {}
func (*RpbGetServerInfoResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RpbGetServerInfoResp) GetNode() []byte {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *RpbGetServerInfoResp) GetServerVersion() []byte {
	if m != nil {
		return m.ServerVersion
	}
	return nil
}

// Key/value pair - used for user metadata, indexes, search doc fields
type RpbPair struct {
	Key              []byte `protobuf:"bytes,1,req,name=key" json:"key,omitempty"`
	Value            []byte `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbPair) Reset()                    { *m = RpbPair{} }
func (m *RpbPair) String() string            { return proto.CompactTextString(m) }
func (*RpbPair) ProtoMessage()               {}
func (*RpbPair) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RpbPair) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RpbPair) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketReq) Reset()                    { *m = RpbGetBucketReq{} }
func (m *RpbGetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketReq) ProtoMessage()               {}
func (*RpbGetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RpbGetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbGetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties response
type RpbGetBucketResp struct {
	Props            *RpbBucketProps `protobuf:"bytes,1,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbGetBucketResp) Reset()                    { *m = RpbGetBucketResp{} }
func (m *RpbGetBucketResp) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketResp) ProtoMessage()               {}
func (*RpbGetBucketResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RpbGetBucketResp) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketReq struct {
	Bucket           []byte          `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	Type             []byte          `protobuf:"bytes,3,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketReq) Reset()                    { *m = RpbSetBucketReq{} }
func (m *RpbSetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbSetBucketReq) ProtoMessage()               {}
func (*RpbSetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *RpbSetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbSetBucketReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

func (m *RpbSetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Reset bucket properties request
type RpbResetBucketReq struct {
	Bucket           []byte `protobuf:"bytes,1,req,name=bucket" json:"bucket,omitempty"`
	Type             []byte `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbResetBucketReq) Reset()                    { *m = RpbResetBucketReq{} }
func (m *RpbResetBucketReq) String() string            { return proto.CompactTextString(m) }
func (*RpbResetBucketReq) ProtoMessage()               {}
func (*RpbResetBucketReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RpbResetBucketReq) GetBucket() []byte {
	if m != nil {
		return m.Bucket
	}
	return nil
}

func (m *RpbResetBucketReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Get bucket properties request
type RpbGetBucketTypeReq struct {
	Type             []byte `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbGetBucketTypeReq) Reset()                    { *m = RpbGetBucketTypeReq{} }
func (m *RpbGetBucketTypeReq) String() string            { return proto.CompactTextString(m) }
func (*RpbGetBucketTypeReq) ProtoMessage()               {}
func (*RpbGetBucketTypeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RpbGetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

// Set bucket properties request
type RpbSetBucketTypeReq struct {
	Type             []byte          `protobuf:"bytes,1,req,name=type" json:"type,omitempty"`
	Props            *RpbBucketProps `protobuf:"bytes,2,req,name=props" json:"props,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *RpbSetBucketTypeReq) Reset()                    { *m = RpbSetBucketTypeReq{} }
func (m *RpbSetBucketTypeReq) String() string            { return proto.CompactTextString(m) }
func (*RpbSetBucketTypeReq) ProtoMessage()               {}
func (*RpbSetBucketTypeReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RpbSetBucketTypeReq) GetType() []byte {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *RpbSetBucketTypeReq) GetProps() *RpbBucketProps {
	if m != nil {
		return m.Props
	}
	return nil
}

// Module-Function pairs for commit hooks and other bucket properties
// that take functions
type RpbModFun struct {
	Module           []byte `protobuf:"bytes,1,req,name=module" json:"module,omitempty"`
	Function         []byte `protobuf:"bytes,2,req,name=function" json:"function,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbModFun) Reset()                    { *m = RpbModFun{} }
func (m *RpbModFun) String() string            { return proto.CompactTextString(m) }
func (*RpbModFun) ProtoMessage()               {}
func (*RpbModFun) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *RpbModFun) GetModule() []byte {
	if m != nil {
		return m.Module
	}
	return nil
}

func (m *RpbModFun) GetFunction() []byte {
	if m != nil {
		return m.Function
	}
	return nil
}

// A commit hook, which may either be a modfun or a JavaScript named
// function
type RpbCommitHook struct {
	Modfun           *RpbModFun `protobuf:"bytes,1,opt,name=modfun" json:"modfun,omitempty"`
	Name             []byte     `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte     `json:"-"`
}

func (m *RpbCommitHook) Reset()                    { *m = RpbCommitHook{} }
func (m *RpbCommitHook) String() string            { return proto.CompactTextString(m) }
func (*RpbCommitHook) ProtoMessage()               {}
func (*RpbCommitHook) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RpbCommitHook) GetModfun() *RpbModFun {
	if m != nil {
		return m.Modfun
	}
	return nil
}

func (m *RpbCommitHook) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

// Bucket properties
type RpbBucketProps struct {
	// Declared in riak_core_app
	NVal          *uint32          `protobuf:"varint,1,opt,name=n_val" json:"n_val,omitempty"`
	AllowMult     *bool            `protobuf:"varint,2,opt,name=allow_mult" json:"allow_mult,omitempty"`
	LastWriteWins *bool            `protobuf:"varint,3,opt,name=last_write_wins" json:"last_write_wins,omitempty"`
	Precommit     []*RpbCommitHook `protobuf:"bytes,4,rep,name=precommit" json:"precommit,omitempty"`
	HasPrecommit  *bool            `protobuf:"varint,5,opt,name=has_precommit,def=0" json:"has_precommit,omitempty"`
	Postcommit    []*RpbCommitHook `protobuf:"bytes,6,rep,name=postcommit" json:"postcommit,omitempty"`
	HasPostcommit *bool            `protobuf:"varint,7,opt,name=has_postcommit,def=0" json:"has_postcommit,omitempty"`
	ChashKeyfun   *RpbModFun       `protobuf:"bytes,8,opt,name=chash_keyfun" json:"chash_keyfun,omitempty"`
	// Declared in riak_kv_app
	Linkfun     *RpbModFun `protobuf:"bytes,9,opt,name=linkfun" json:"linkfun,omitempty"`
	OldVclock   *uint32    `protobuf:"varint,10,opt,name=old_vclock" json:"old_vclock,omitempty"`
	YoungVclock *uint32    `protobuf:"varint,11,opt,name=young_vclock" json:"young_vclock,omitempty"`
	BigVclock   *uint32    `protobuf:"varint,12,opt,name=big_vclock" json:"big_vclock,omitempty"`
	SmallVclock *uint32    `protobuf:"varint,13,opt,name=small_vclock" json:"small_vclock,omitempty"`
	Pr          *uint32    `protobuf:"varint,14,opt,name=pr" json:"pr,omitempty"`
	R           *uint32    `protobuf:"varint,15,opt,name=r" json:"r,omitempty"`
	W           *uint32    `protobuf:"varint,16,opt,name=w" json:"w,omitempty"`
	Pw          *uint32    `protobuf:"varint,17,opt,name=pw" json:"pw,omitempty"`
	Dw          *uint32    `protobuf:"varint,18,opt,name=dw" json:"dw,omitempty"`
	Rw          *uint32    `protobuf:"varint,19,opt,name=rw" json:"rw,omitempty"`
	BasicQuorum *bool      `protobuf:"varint,20,opt,name=basic_quorum" json:"basic_quorum,omitempty"`
	NotfoundOk  *bool      `protobuf:"varint,21,opt,name=notfound_ok" json:"notfound_ok,omitempty"`
	// Used by riak_kv_multi_backend
	Backend []byte `protobuf:"bytes,22,opt,name=backend" json:"backend,omitempty"`
	// Used by riak_search bucket fixup
	Search *bool                       `protobuf:"varint,23,opt,name=search" json:"search,omitempty"`
	Repl   *RpbBucketProps_RpbReplMode `protobuf:"varint,24,opt,name=repl,enum=RpbBucketProps_RpbReplMode" json:"repl,omitempty"`
	// Search index
	SearchIndex []byte `protobuf:"bytes,25,opt,name=search_index" json:"search_index,omitempty"`
	// KV Datatypes
	Datatype []byte `protobuf:"bytes,26,opt,name=datatype" json:"datatype,omitempty"`
	// KV strong consistency
	Consistent *bool `protobuf:"varint,27,opt,name=consistent" json:"consistent,omitempty"`
	// KV fast path
	WriteOnce *bool `protobuf:"varint,28,opt,name=write_once" json:"write_once,omitempty"`
	// Hyperlolog DT Precision
	HllPrecision     *uint32 `protobuf:"varint,29,opt,name=hll_precision" json:"hll_precision,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbBucketProps) Reset()                    { *m = RpbBucketProps{} }
func (m *RpbBucketProps) String() string            { return proto.CompactTextString(m) }
func (*RpbBucketProps) ProtoMessage()               {}
func (*RpbBucketProps) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

const Default_RpbBucketProps_HasPrecommit bool = false
const Default_RpbBucketProps_HasPostcommit bool = false

func (m *RpbBucketProps) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

func (m *RpbBucketProps) GetAllowMult() bool {
	if m != nil && m.AllowMult != nil {
		return *m.AllowMult
	}
	return false
}

func (m *RpbBucketProps) GetLastWriteWins() bool {
	if m != nil && m.LastWriteWins != nil {
		return *m.LastWriteWins
	}
	return false
}

func (m *RpbBucketProps) GetPrecommit() []*RpbCommitHook {
	if m != nil {
		return m.Precommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPrecommit() bool {
	if m != nil && m.HasPrecommit != nil {
		return *m.HasPrecommit
	}
	return Default_RpbBucketProps_HasPrecommit
}

func (m *RpbBucketProps) GetPostcommit() []*RpbCommitHook {
	if m != nil {
		return m.Postcommit
	}
	return nil
}

func (m *RpbBucketProps) GetHasPostcommit() bool {
	if m != nil && m.HasPostcommit != nil {
		return *m.HasPostcommit
	}
	return Default_RpbBucketProps_HasPostcommit
}

func (m *RpbBucketProps) GetChashKeyfun() *RpbModFun {
	if m != nil {
		return m.ChashKeyfun
	}
	return nil
}

func (m *RpbBucketProps) GetLinkfun() *RpbModFun {
	if m != nil {
		return m.Linkfun
	}
	return nil
}

func (m *RpbBucketProps) GetOldVclock() uint32 {
	if m != nil && m.OldVclock != nil {
		return *m.OldVclock
	}
	return 0
}

func (m *RpbBucketProps) GetYoungVclock() uint32 {
	if m != nil && m.YoungVclock != nil {
		return *m.YoungVclock
	}
	return 0
}

func (m *RpbBucketProps) GetBigVclock() uint32 {
	if m != nil && m.BigVclock != nil {
		return *m.BigVclock
	}
	return 0
}

func (m *RpbBucketProps) GetSmallVclock() uint32 {
	if m != nil && m.SmallVclock != nil {
		return *m.SmallVclock
	}
	return 0
}

func (m *RpbBucketProps) GetPr() uint32 {
	if m != nil && m.Pr != nil {
		return *m.Pr
	}
	return 0
}

func (m *RpbBucketProps) GetR() uint32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return 0
}

func (m *RpbBucketProps) GetW() uint32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

func (m *RpbBucketProps) GetPw() uint32 {
	if m != nil && m.Pw != nil {
		return *m.Pw
	}
	return 0
}

func (m *RpbBucketProps) GetDw() uint32 {
	if m != nil && m.Dw != nil {
		return *m.Dw
	}
	return 0
}

func (m *RpbBucketProps) GetRw() uint32 {
	if m != nil && m.Rw != nil {
		return *m.Rw
	}
	return 0
}

func (m *RpbBucketProps) GetBasicQuorum() bool {
	if m != nil && m.BasicQuorum != nil {
		return *m.BasicQuorum
	}
	return false
}

func (m *RpbBucketProps) GetNotfoundOk() bool {
	if m != nil && m.NotfoundOk != nil {
		return *m.NotfoundOk
	}
	return false
}

func (m *RpbBucketProps) GetBackend() []byte {
	if m != nil {
		return m.Backend
	}
	return nil
}

func (m *RpbBucketProps) GetSearch() bool {
	if m != nil && m.Search != nil {
		return *m.Search
	}
	return false
}

func (m *RpbBucketProps) GetRepl() RpbBucketProps_RpbReplMode {
	if m != nil && m.Repl != nil {
		return *m.Repl
	}
	return RpbBucketProps_FALSE
}

func (m *RpbBucketProps) GetSearchIndex() []byte {
	if m != nil {
		return m.SearchIndex
	}
	return nil
}

func (m *RpbBucketProps) GetDatatype() []byte {
	if m != nil {
		return m.Datatype
	}
	return nil
}

func (m *RpbBucketProps) GetConsistent() bool {
	if m != nil && m.Consistent != nil {
		return *m.Consistent
	}
	return false
}

func (m *RpbBucketProps) GetWriteOnce() bool {
	if m != nil && m.WriteOnce != nil {
		return *m.WriteOnce
	}
	return false
}

func (m *RpbBucketProps) GetHllPrecision() uint32 {
	if m != nil && m.HllPrecision != nil {
		return *m.HllPrecision
	}
	return 0
}

// Authentication request
type RpbAuthReq struct {
	User             []byte `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	Password         []byte `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbAuthReq) Reset()                    { *m = RpbAuthReq{} }
func (m *RpbAuthReq) String() string            { return proto.CompactTextString(m) }
func (*RpbAuthReq) ProtoMessage()               {}
func (*RpbAuthReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *RpbAuthReq) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *RpbAuthReq) GetPassword() []byte {
	if m != nil {
		return m.Password
	}
	return nil
}

func init() {
	proto.RegisterType((*RpbErrorResp)(nil), "RpbErrorResp")
	proto.RegisterType((*RpbGetServerInfoResp)(nil), "RpbGetServerInfoResp")
	proto.RegisterType((*RpbPair)(nil), "RpbPair")
	proto.RegisterType((*RpbGetBucketReq)(nil), "RpbGetBucketReq")
	proto.RegisterType((*RpbGetBucketResp)(nil), "RpbGetBucketResp")
	proto.RegisterType((*RpbSetBucketReq)(nil), "RpbSetBucketReq")
	proto.RegisterType((*RpbResetBucketReq)(nil), "RpbResetBucketReq")
	proto.RegisterType((*RpbGetBucketTypeReq)(nil), "RpbGetBucketTypeReq")
	proto.RegisterType((*RpbSetBucketTypeReq)(nil), "RpbSetBucketTypeReq")
	proto.RegisterType((*RpbModFun)(nil), "RpbModFun")
	proto.RegisterType((*RpbCommitHook)(nil), "RpbCommitHook")
	proto.RegisterType((*RpbBucketProps)(nil), "RpbBucketProps")
	proto.RegisterType((*RpbAuthReq)(nil), "RpbAuthReq")
	proto.RegisterEnum("RpbBucketProps_RpbReplMode", RpbBucketProps_RpbReplMode_name, RpbBucketProps_RpbReplMode_value)
}

var fileDescriptor0 = []byte{
	// 738 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x54, 0xdb, 0x6e, 0xdb, 0x46,
	0x10, 0xad, 0x6e, 0x96, 0x34, 0xba, 0x9a, 0xbe, 0x6d, 0x7d, 0x29, 0x64, 0x16, 0x05, 0x5a, 0xa0,
	0x95, 0x51, 0xbd, 0xb5, 0x08, 0x02, 0x58, 0x86, 0x9c, 0x18, 0x90, 0x13, 0x43, 0xb2, 0x1f, 0xf2,
	0x44, 0x50, 0xe4, 0xca, 0x22, 0x44, 0x72, 0xe9, 0xdd, 0xa5, 0x19, 0x7f, 0x61, 0x7e, 0x2b, 0xc3,
	0x91, 0x28, 0xf9, 0x06, 0x27, 0x0f, 0x04, 0xb8, 0x67, 0xce, 0x99, 0x39, 0x3b, 0x3b, 0xbb, 0x00,
	0xd2, 0xb3, 0xe7, 0xdd, 0x48, 0x0a, 0x2d, 0xcc, 0x13, 0xa8, 0x8f, 0xa2, 0xc9, 0x40, 0x4a, 0x21,
	0x47, 0x5c, 0x45, 0x46, 0x13, 0x36, 0xb8, 0x94, 0x81, 0xba, 0x65, 0xb9, 0x4e, 0xfe, 0xcf, 0xba,
	0xd1, 0x82, 0x32, 0xae, 0x1d, 0xe1, 0x72, 0x96, 0x47, 0xa0, 0x61, 0xbe, 0x83, 0x6d, 0x14, 0x7c,
	0xe0, 0x7a, 0xcc, 0xe5, 0x3d, 0x97, 0x17, 0xe1, 0x54, 0x90, 0xb0, 0x0e, 0xc5, 0x30, 0x65, 0xe5,
	0x3a, 0x39, 0x94, 0xed, 0x42, 0x53, 0x51, 0xdc, 0xc2, 0x4f, 0x79, 0x22, 0x44, 0x35, 0xe2, 0xe6,
	0x1f, 0x50, 0x46, 0xf5, 0x95, 0xed, 0x49, 0xa3, 0x06, 0x85, 0x39, 0x7f, 0x58, 0x96, 0x69, 0x40,
	0xe9, 0xde, 0xf6, 0x63, 0xbe, 0xa4, 0x9d, 0x40, 0x6b, 0x51, 0xa4, 0x1f, 0x3b, 0x73, 0xae, 0x47,
	0xfc, 0x2e, 0x35, 0x36, 0xa1, 0xc5, 0x52, 0x81, 0xf5, 0xf4, 0x43, 0x94, 0x09, 0x7a, 0xd0, 0x7e,
	0x2a, 0x40, 0x47, 0xbf, 0x41, 0x09, 0xf7, 0x18, 0x29, 0x12, 0xd4, 0x7a, 0xad, 0x2e, 0x32, 0x16,
	0xe1, 0xab, 0x14, 0x36, 0x3f, 0x53, 0x91, 0xf1, 0x5b, 0x45, 0x56, 0x29, 0xf2, 0xaf, 0xa6, 0x58,
	0x99, 0x28, 0x90, 0x89, 0x7f, 0x61, 0x13, 0xe3, 0x58, 0xfb, 0xe7, 0x7d, 0xff, 0x0e, 0x5b, 0x8f,
	0x7d, 0x5f, 0x63, 0x24, 0x15, 0x65, 0x24, 0x92, 0x98, 0x67, 0x44, 0x1a, 0xbf, 0x49, 0xfa, 0x91,
	0x55, 0xf3, 0x1f, 0xa8, 0x22, 0x72, 0x29, 0xdc, 0xf3, 0x38, 0x4c, 0x4d, 0x05, 0xc2, 0x8d, 0xfd,
	0x4c, 0xdc, 0x86, 0xca, 0x34, 0x0e, 0x1d, 0xbd, 0x38, 0xa8, 0xb4, 0xe6, 0x7f, 0xd0, 0x40, 0xfa,
	0x99, 0x08, 0x02, 0x4f, 0x7f, 0x14, 0x62, 0x6e, 0xec, 0x93, 0x04, 0x59, 0x74, 0xc2, 0xb5, 0x1e,
	0x74, 0xd7, 0xe9, 0xd2, 0xb3, 0xb7, 0x83, 0x6c, 0x4f, 0xdf, 0x4a, 0xd0, 0x7c, 0xd6, 0x27, 0x3c,
	0xde, 0xd0, 0xc2, 0x03, 0x26, 0x6d, 0xc3, 0x30, 0x00, 0x6c, 0xdf, 0x17, 0x89, 0x15, 0xc4, 0xbe,
	0x26, 0x55, 0xc5, 0xd8, 0x83, 0x96, 0x6f, 0x2b, 0x6d, 0x25, 0xd2, 0xd3, 0xdc, 0x4a, 0xbc, 0x50,
	0x51, 0x57, 0x2b, 0xc6, 0x31, 0x54, 0x23, 0xc9, 0x1d, 0x72, 0xc2, 0x8a, 0x9d, 0x02, 0xd6, 0x6e,
	0x76, 0x9f, 0x7a, 0x3b, 0x84, 0xc6, 0xcc, 0x56, 0xd6, 0x9a, 0x56, 0x4a, 0x95, 0xff, 0x97, 0xa6,
	0xb6, 0xaf, 0xb8, 0x61, 0x02, 0x44, 0x42, 0xe9, 0x65, 0x68, 0xe3, 0xd5, 0x0c, 0x47, 0xd0, 0xa4,
	0x0c, 0x6b, 0x5e, 0xf9, 0x71, 0x8a, 0x0e, 0xd4, 0x1d, 0x8c, 0xcf, 0x2c, 0x9c, 0xd8, 0xb4, 0x05,
	0x95, 0x17, 0x2d, 0x38, 0x80, 0xb2, 0xef, 0x85, 0xf3, 0x34, 0x58, 0x7d, 0x11, 0xc4, 0xfd, 0x0a,
	0xdf, 0xb5, 0xee, 0x1d, 0x5f, 0x38, 0x73, 0x06, 0xd4, 0x83, 0x6d, 0xa8, 0x3f, 0x88, 0x38, 0xbc,
	0xcd, 0xd0, 0x5a, 0xd6, 0x99, 0x89, 0xb7, 0xc2, 0xea, 0x19, 0x53, 0x05, 0xd8, 0xaf, 0x0c, 0x6d,
	0x10, 0x0a, 0x90, 0x8f, 0x24, 0x6b, 0xd2, 0x7f, 0x15, 0x72, 0x92, 0xb5, 0xb2, 0xdf, 0x84, 0xb5,
	0x57, 0x8c, 0x84, 0x6d, 0x66, 0xff, 0x6e, 0xc2, 0x8c, 0xec, 0x5f, 0x26, 0x6c, 0x2b, 0xcb, 0x3d,
	0xb1, 0x95, 0xe7, 0x58, 0x77, 0xb1, 0x90, 0x71, 0xc0, 0xb6, 0xa9, 0xe5, 0x5b, 0x50, 0x0b, 0x85,
	0x9e, 0xa2, 0x3d, 0xd7, 0x12, 0x73, 0xb6, 0x43, 0x20, 0xbe, 0x04, 0x13, 0x1b, 0x8f, 0x34, 0x74,
	0xd9, 0x2e, 0xdd, 0x71, 0x1c, 0x22, 0xc5, 0x6d, 0xe9, 0xcc, 0xd8, 0x1e, 0x11, 0xfe, 0x82, 0xa2,
	0xe4, 0x91, 0xcf, 0x18, 0xae, 0x9a, 0xbd, 0x83, 0x67, 0x03, 0xd8, 0xa5, 0xab, 0x11, 0xf9, 0xd8,
	0x12, 0x4e, 0x5b, 0x22, 0xa9, 0xe5, 0x85, 0x2e, 0xff, 0xca, 0x7e, 0xa5, 0x84, 0x38, 0x85, 0xae,
	0xad, 0x6d, 0x1a, 0xea, 0x7d, 0x42, 0xd0, 0xab, 0x23, 0x42, 0xe5, 0x29, 0xcd, 0x43, 0xcd, 0x0e,
	0xa8, 0x0c, 0x62, 0x8b, 0x19, 0x11, 0xa1, 0xc3, 0xd9, 0x21, 0x61, 0x3b, 0x38, 0x00, 0xd8, 0xa0,
	0x74, 0x00, 0x3c, 0x7a, 0x6d, 0x8e, 0xd2, 0xdd, 0x99, 0xef, 0xa1, 0xf6, 0xb8, 0x6a, 0x15, 0x4a,
	0xe7, 0xa7, 0xc3, 0xf1, 0xa0, 0xfd, 0x0b, 0x4e, 0x6c, 0x65, 0x34, 0x38, 0x1d, 0x5e, 0x5f, 0x5c,
	0x0e, 0xda, 0xb9, 0x74, 0x75, 0x7e, 0x33, 0x1c, 0x8e, 0xbf, 0x7c, 0x3a, 0x6b, 0xe7, 0x8d, 0x0a,
	0x14, 0xaf, 0x47, 0x37, 0x83, 0x76, 0xc1, 0xfc, 0x1b, 0x00, 0xf5, 0xa7, 0xb1, 0x9e, 0x2d, 0xef,
	0x5b, 0x8c, 0x8f, 0xda, 0xfa, 0xca, 0x44, 0xb6, 0x52, 0x89, 0x90, 0xee, 0xe2, 0xca, 0xf4, 0x8f,
	0x61, 0x0f, 0x67, 0xa7, 0x8b, 0xfd, 0x9c, 0x89, 0xee, 0xfa, 0x89, 0x9d, 0xc4, 0xd3, 0xfe, 0xc6,
	0x08, 0x97, 0x57, 0xfd, 0xef, 0x01, 0x00, 0x00, 0xff, 0xff, 0x16, 0x03, 0xd1, 0x19, 0x7a, 0x05,
	0x00, 0x00,
}
