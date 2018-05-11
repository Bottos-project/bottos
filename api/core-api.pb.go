// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/bottos-project/core/api/core-api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/bottos-project/core/api/core-api.proto

It has these top-level messages:
	PushTrxResponse
	QueryTrxRequest
	QueryTrxResponse
	QueryBlockRequest
	QueryBlockResponse
	QueryChainInfoRequest
	QueryChainInfoResponse
	QueryAccountRequest
	QueryAccountResponse
	QueryObjectByKeyReq
	QueryObjectByKeyResponse
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/bottos-project/core/common/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PushTrxResponse struct {
	Errcode uint32                  `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string                  `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *PushTrxResponse_Result `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *PushTrxResponse) Reset()                    { *m = PushTrxResponse{} }
func (m *PushTrxResponse) String() string            { return proto.CompactTextString(m) }
func (*PushTrxResponse) ProtoMessage()               {}
func (*PushTrxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PushTrxResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *PushTrxResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *PushTrxResponse) GetResult() *PushTrxResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type PushTrxResponse_Result struct {
	Trx     *types.Transaction `protobuf:"bytes,1,opt,name=trx" json:"trx"`
	TrxHash string             `protobuf:"bytes,2,opt,name=trx_hash,json=trxHash" json:"trx_hash"`
}

func (m *PushTrxResponse_Result) Reset()                    { *m = PushTrxResponse_Result{} }
func (m *PushTrxResponse_Result) String() string            { return proto.CompactTextString(m) }
func (*PushTrxResponse_Result) ProtoMessage()               {}
func (*PushTrxResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

func (m *PushTrxResponse_Result) GetTrx() *types.Transaction {
	if m != nil {
		return m.Trx
	}
	return nil
}

func (m *PushTrxResponse_Result) GetTrxHash() string {
	if m != nil {
		return m.TrxHash
	}
	return ""
}

type QueryTrxRequest struct {
	TrxHash string `protobuf:"bytes,1,opt,name=trx_hash,json=trxHash" json:"trx_hash"`
}

func (m *QueryTrxRequest) Reset()                    { *m = QueryTrxRequest{} }
func (m *QueryTrxRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryTrxRequest) ProtoMessage()               {}
func (*QueryTrxRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *QueryTrxRequest) GetTrxHash() string {
	if m != nil {
		return m.TrxHash
	}
	return ""
}

type QueryTrxResponse struct {
	Errcode uint32             `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string             `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *types.Transaction `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *QueryTrxResponse) Reset()                    { *m = QueryTrxResponse{} }
func (m *QueryTrxResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryTrxResponse) ProtoMessage()               {}
func (*QueryTrxResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *QueryTrxResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *QueryTrxResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryTrxResponse) GetResult() *types.Transaction {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryBlockRequest struct {
	BlockNum  uint32 `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num"`
	BlockHash string `protobuf:"bytes,2,opt,name=block_hash,json=blockHash" json:"block_hash"`
}

func (m *QueryBlockRequest) Reset()                    { *m = QueryBlockRequest{} }
func (m *QueryBlockRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryBlockRequest) ProtoMessage()               {}
func (*QueryBlockRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *QueryBlockRequest) GetBlockNum() uint32 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *QueryBlockRequest) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

type QueryBlockResponse struct {
	Errcode uint32                     `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string                     `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *QueryBlockResponse_Result `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *QueryBlockResponse) Reset()                    { *m = QueryBlockResponse{} }
func (m *QueryBlockResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryBlockResponse) ProtoMessage()               {}
func (*QueryBlockResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryBlockResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *QueryBlockResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryBlockResponse) GetResult() *QueryBlockResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryBlockResponse_Result struct {
	PrevBlockHash    string               `protobuf:"bytes,1,opt,name=prev_block_hash,json=prevBlockHash" json:"prev_block_hash"`
	BlockNum         uint32               `protobuf:"varint,2,opt,name=block_num,json=blockNum" json:"block_num"`
	BlockHash        string               `protobuf:"bytes,3,opt,name=block_hash,json=blockHash" json:"block_hash"`
	CursorBlockLabel uint32               `protobuf:"varint,4,opt,name=cursor_block_label,json=cursorBlockLabel" json:"cursor_block_label"`
	BlockTime        uint64               `protobuf:"varint,5,opt,name=block_time,json=blockTime" json:"block_time"`
	TrxMerkleRoot    string               `protobuf:"bytes,6,opt,name=trx_merkle_root,json=trxMerkleRoot" json:"trx_merkle_root"`
	Delegate         string               `protobuf:"bytes,7,opt,name=delegate" json:"delegate"`
	DelegateSign     string               `protobuf:"bytes,8,opt,name=delegate_sign,json=delegateSign" json:"delegate_sign"`
	Trxs             []*types.Transaction `protobuf:"bytes,9,rep,name=trxs" json:"trxs"`
}

func (m *QueryBlockResponse_Result) Reset()                    { *m = QueryBlockResponse_Result{} }
func (m *QueryBlockResponse_Result) String() string            { return proto.CompactTextString(m) }
func (*QueryBlockResponse_Result) ProtoMessage()               {}
func (*QueryBlockResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4, 0} }

func (m *QueryBlockResponse_Result) GetPrevBlockHash() string {
	if m != nil {
		return m.PrevBlockHash
	}
	return ""
}

func (m *QueryBlockResponse_Result) GetBlockNum() uint32 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *QueryBlockResponse_Result) GetBlockHash() string {
	if m != nil {
		return m.BlockHash
	}
	return ""
}

func (m *QueryBlockResponse_Result) GetCursorBlockLabel() uint32 {
	if m != nil {
		return m.CursorBlockLabel
	}
	return 0
}

func (m *QueryBlockResponse_Result) GetBlockTime() uint64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *QueryBlockResponse_Result) GetTrxMerkleRoot() string {
	if m != nil {
		return m.TrxMerkleRoot
	}
	return ""
}

func (m *QueryBlockResponse_Result) GetDelegate() string {
	if m != nil {
		return m.Delegate
	}
	return ""
}

func (m *QueryBlockResponse_Result) GetDelegateSign() string {
	if m != nil {
		return m.DelegateSign
	}
	return ""
}

func (m *QueryBlockResponse_Result) GetTrxs() []*types.Transaction {
	if m != nil {
		return m.Trxs
	}
	return nil
}

type QueryChainInfoRequest struct {
}

func (m *QueryChainInfoRequest) Reset()                    { *m = QueryChainInfoRequest{} }
func (m *QueryChainInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryChainInfoRequest) ProtoMessage()               {}
func (*QueryChainInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type QueryChainInfoResponse struct {
	Errcode uint32                         `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string                         `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *QueryChainInfoResponse_Result `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *QueryChainInfoResponse) Reset()                    { *m = QueryChainInfoResponse{} }
func (m *QueryChainInfoResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryChainInfoResponse) ProtoMessage()               {}
func (*QueryChainInfoResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *QueryChainInfoResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *QueryChainInfoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryChainInfoResponse) GetResult() *QueryChainInfoResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryChainInfoResponse_Result struct {
	HeadBlockNum          uint32 `protobuf:"varint,1,opt,name=head_block_num,json=headBlockNum" json:"head_block_num"`
	HeadBlockHash         string `protobuf:"bytes,2,opt,name=head_block_hash,json=headBlockHash" json:"head_block_hash"`
	HeadBlockTime         uint64 `protobuf:"varint,3,opt,name=head_block_time,json=headBlockTime" json:"head_block_time"`
	HeadBlockDelegate     string `protobuf:"bytes,4,opt,name=head_block_delegate,json=headBlockDelegate" json:"head_block_delegate"`
	CursorLabel           uint32 `protobuf:"varint,5,opt,name=cursor_label,json=cursorLabel" json:"cursor_label"`
	LastConsensusBlockNum uint32 `protobuf:"varint,6,opt,name=last_consensus_block_num,json=lastConsensusBlockNum" json:"last_consensus_block_num"`
}

func (m *QueryChainInfoResponse_Result) Reset()         { *m = QueryChainInfoResponse_Result{} }
func (m *QueryChainInfoResponse_Result) String() string { return proto.CompactTextString(m) }
func (*QueryChainInfoResponse_Result) ProtoMessage()    {}
func (*QueryChainInfoResponse_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

func (m *QueryChainInfoResponse_Result) GetHeadBlockNum() uint32 {
	if m != nil {
		return m.HeadBlockNum
	}
	return 0
}

func (m *QueryChainInfoResponse_Result) GetHeadBlockHash() string {
	if m != nil {
		return m.HeadBlockHash
	}
	return ""
}

func (m *QueryChainInfoResponse_Result) GetHeadBlockTime() uint64 {
	if m != nil {
		return m.HeadBlockTime
	}
	return 0
}

func (m *QueryChainInfoResponse_Result) GetHeadBlockDelegate() string {
	if m != nil {
		return m.HeadBlockDelegate
	}
	return ""
}

func (m *QueryChainInfoResponse_Result) GetCursorLabel() uint32 {
	if m != nil {
		return m.CursorLabel
	}
	return 0
}

func (m *QueryChainInfoResponse_Result) GetLastConsensusBlockNum() uint32 {
	if m != nil {
		return m.LastConsensusBlockNum
	}
	return 0
}

type QueryAccountRequest struct {
	AccountName string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
}

func (m *QueryAccountRequest) Reset()                    { *m = QueryAccountRequest{} }
func (m *QueryAccountRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountRequest) ProtoMessage()               {}
func (*QueryAccountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *QueryAccountRequest) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

type QueryAccountResponse struct {
	Errcode uint32                       `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string                       `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *QueryAccountResponse_Result `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *QueryAccountResponse) Reset()                    { *m = QueryAccountResponse{} }
func (m *QueryAccountResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountResponse) ProtoMessage()               {}
func (*QueryAccountResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *QueryAccountResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *QueryAccountResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryAccountResponse) GetResult() *QueryAccountResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryAccountResponse_Result struct {
	AccountName   string `protobuf:"bytes,1,opt,name=account_name,json=accountName" json:"account_name"`
	Pubkey        string `protobuf:"bytes,2,opt,name=pubkey" json:"pubkey"`
	Balance       uint64 `protobuf:"varint,3,opt,name=balance" json:"balance"`
	StakedBalance uint64 `protobuf:"varint,4,opt,name=staked_balance,json=stakedBalance" json:"staked_balance"`
}

func (m *QueryAccountResponse_Result) Reset()                    { *m = QueryAccountResponse_Result{} }
func (m *QueryAccountResponse_Result) String() string            { return proto.CompactTextString(m) }
func (*QueryAccountResponse_Result) ProtoMessage()               {}
func (*QueryAccountResponse_Result) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8, 0} }

func (m *QueryAccountResponse_Result) GetAccountName() string {
	if m != nil {
		return m.AccountName
	}
	return ""
}

func (m *QueryAccountResponse_Result) GetPubkey() string {
	if m != nil {
		return m.Pubkey
	}
	return ""
}

func (m *QueryAccountResponse_Result) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *QueryAccountResponse_Result) GetStakedBalance() uint64 {
	if m != nil {
		return m.StakedBalance
	}
	return 0
}

type QueryObjectByKeyReq struct {
	Contract string `protobuf:"bytes,1,opt,name=contract" json:"contract"`
	Object   string `protobuf:"bytes,2,opt,name=object" json:"object"`
	Key      string `protobuf:"bytes,3,opt,name=key" json:"key"`
}

func (m *QueryObjectByKeyReq) Reset()                    { *m = QueryObjectByKeyReq{} }
func (m *QueryObjectByKeyReq) String() string            { return proto.CompactTextString(m) }
func (*QueryObjectByKeyReq) ProtoMessage()               {}
func (*QueryObjectByKeyReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *QueryObjectByKeyReq) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *QueryObjectByKeyReq) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *QueryObjectByKeyReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type QueryObjectByKeyResponse struct {
	Errcode uint32                           `protobuf:"varint,1,opt,name=errcode" json:"errcode"`
	Msg     string                           `protobuf:"bytes,2,opt,name=msg" json:"msg"`
	Result  *QueryObjectByKeyResponse_Result `protobuf:"bytes,3,opt,name=result" json:"result"`
}

func (m *QueryObjectByKeyResponse) Reset()                    { *m = QueryObjectByKeyResponse{} }
func (m *QueryObjectByKeyResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryObjectByKeyResponse) ProtoMessage()               {}
func (*QueryObjectByKeyResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *QueryObjectByKeyResponse) GetErrcode() uint32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *QueryObjectByKeyResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *QueryObjectByKeyResponse) GetResult() *QueryObjectByKeyResponse_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

type QueryObjectByKeyResponse_Result struct {
	Contract string `protobuf:"bytes,1,opt,name=contract" json:"contract"`
	Object   string `protobuf:"bytes,2,opt,name=object" json:"object"`
	Key      string `protobuf:"bytes,3,opt,name=key" json:"key"`
	Value    string `protobuf:"bytes,4,opt,name=value" json:"value"`
}

func (m *QueryObjectByKeyResponse_Result) Reset()         { *m = QueryObjectByKeyResponse_Result{} }
func (m *QueryObjectByKeyResponse_Result) String() string { return proto.CompactTextString(m) }
func (*QueryObjectByKeyResponse_Result) ProtoMessage()    {}
func (*QueryObjectByKeyResponse_Result) Descriptor() ([]byte, []int) {
	return fileDescriptor0, []int{10, 0}
}

func (m *QueryObjectByKeyResponse_Result) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *QueryObjectByKeyResponse_Result) GetObject() string {
	if m != nil {
		return m.Object
	}
	return ""
}

func (m *QueryObjectByKeyResponse_Result) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *QueryObjectByKeyResponse_Result) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*PushTrxResponse)(nil), "api.PushTrxResponse")
	proto.RegisterType((*PushTrxResponse_Result)(nil), "api.PushTrxResponse.Result")
	proto.RegisterType((*QueryTrxRequest)(nil), "api.QueryTrxRequest")
	proto.RegisterType((*QueryTrxResponse)(nil), "api.QueryTrxResponse")
	proto.RegisterType((*QueryBlockRequest)(nil), "api.QueryBlockRequest")
	proto.RegisterType((*QueryBlockResponse)(nil), "api.QueryBlockResponse")
	proto.RegisterType((*QueryBlockResponse_Result)(nil), "api.QueryBlockResponse.Result")
	proto.RegisterType((*QueryChainInfoRequest)(nil), "api.QueryChainInfoRequest")
	proto.RegisterType((*QueryChainInfoResponse)(nil), "api.QueryChainInfoResponse")
	proto.RegisterType((*QueryChainInfoResponse_Result)(nil), "api.QueryChainInfoResponse.Result")
	proto.RegisterType((*QueryAccountRequest)(nil), "api.QueryAccountRequest")
	proto.RegisterType((*QueryAccountResponse)(nil), "api.QueryAccountResponse")
	proto.RegisterType((*QueryAccountResponse_Result)(nil), "api.QueryAccountResponse.Result")
	proto.RegisterType((*QueryObjectByKeyReq)(nil), "api.QueryObjectByKeyReq")
	proto.RegisterType((*QueryObjectByKeyResponse)(nil), "api.QueryObjectByKeyResponse")
	proto.RegisterType((*QueryObjectByKeyResponse_Result)(nil), "api.QueryObjectByKeyResponse.Result")
}

func init() { proto.RegisterFile("github.com/bottos-project/core/api/core-api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 922 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5f, 0x6f, 0xe3, 0x44,
	0x10, 0x4f, 0xe2, 0x34, 0x49, 0xa7, 0xe9, 0x9f, 0xdb, 0x6b, 0x7b, 0xbe, 0x54, 0x87, 0x72, 0xa6,
	0x54, 0x15, 0xba, 0x4b, 0x44, 0x4f, 0xc0, 0x81, 0x40, 0xe8, 0x52, 0x90, 0x38, 0x15, 0xee, 0x20,
	0xed, 0x13, 0x3c, 0x58, 0x1b, 0x67, 0x49, 0x7c, 0xb5, 0xbd, 0xee, 0xee, 0xba, 0x4a, 0x9e, 0xe1,
	0x09, 0x89, 0x8f, 0xc1, 0xc7, 0xe0, 0x91, 0x17, 0x3e, 0x0a, 0x9f, 0x02, 0xed, 0xae, 0xd7, 0x76,
	0xdc, 0xb4, 0x95, 0x22, 0xde, 0xbc, 0xb3, 0xbf, 0x9d, 0xf9, 0xcd, 0xfc, 0x66, 0x77, 0x0c, 0x1f,
	0x4d, 0x7c, 0x31, 0x4d, 0x46, 0x3d, 0x8f, 0x86, 0xfd, 0x11, 0x15, 0x82, 0xf2, 0xe7, 0x31, 0xa3,
	0xef, 0x88, 0x27, 0xfa, 0x1e, 0x65, 0xa4, 0x8f, 0x63, 0x5f, 0x7d, 0x3c, 0xc7, 0xb1, 0xdf, 0x8b,
	0x19, 0x15, 0x14, 0x59, 0x38, 0xf6, 0x3b, 0x5f, 0xde, 0x73, 0xce, 0xa3, 0x61, 0x48, 0xa3, 0xbe,
	0x98, 0xc7, 0x84, 0xf7, 0x05, 0xc3, 0x11, 0xc7, 0x9e, 0xf0, 0x69, 0xa4, 0x7d, 0x38, 0x7f, 0x55,
	0x61, 0xfb, 0x87, 0x84, 0x4f, 0x2f, 0xd8, 0x6c, 0x48, 0x78, 0x4c, 0x23, 0x4e, 0x90, 0x0d, 0x4d,
	0xc2, 0x98, 0x47, 0xc7, 0xc4, 0xae, 0x76, 0xab, 0xc7, 0x9b, 0x43, 0xb3, 0x44, 0x3b, 0x60, 0x85,
	0x7c, 0x62, 0xd7, 0xba, 0xd5, 0xe3, 0xf5, 0xa1, 0xfc, 0x44, 0x2f, 0xa0, 0xc1, 0x08, 0x4f, 0x02,
	0x61, 0x5b, 0xdd, 0xea, 0xf1, 0xc6, 0xc9, 0x41, 0x4f, 0xf2, 0x2b, 0x79, 0xec, 0x0d, 0x15, 0x64,
	0x98, 0x42, 0x3b, 0xaf, 0xa1, 0xa1, 0x2d, 0xe8, 0x10, 0x2c, 0xc1, 0x66, 0x2a, 0xcc, 0xc6, 0x09,
	0xea, 0x29, 0x96, 0xbd, 0x8b, 0x9c, 0xe5, 0x50, 0x6e, 0xa3, 0xc7, 0xd0, 0x12, 0x6c, 0xe6, 0x4e,
	0x31, 0x9f, 0xa6, 0xb1, 0x9b, 0x82, 0xcd, 0xbe, 0xc5, 0x7c, 0xea, 0x3c, 0x83, 0xed, 0x1f, 0x13,
	0xc2, 0xe6, 0x2a, 0xda, 0x55, 0x42, 0xb8, 0x58, 0x40, 0x57, 0x17, 0xd1, 0xef, 0x60, 0x27, 0x47,
	0xaf, 0x90, 0xed, 0x87, 0xa5, 0x6c, 0x97, 0x31, 0x4e, 0x11, 0xce, 0x5b, 0x78, 0xa0, 0x62, 0x0d,
	0x02, 0xea, 0x5d, 0x1a, 0x6e, 0x07, 0xb0, 0x3e, 0x92, 0x6b, 0x37, 0x4a, 0xc2, 0x34, 0x5c, 0x4b,
	0x19, 0xde, 0x24, 0x21, 0x7a, 0x02, 0xa0, 0x37, 0x0b, 0x89, 0x6a, 0xb8, 0x22, 0xff, 0xb7, 0x05,
	0xa8, 0xe8, 0x71, 0x05, 0xfe, 0x9f, 0x94, 0xf8, 0xbf, 0xa7, 0xd4, 0xba, 0xe9, 0xb4, 0x2c, 0xd8,
	0x3f, 0xb5, 0x4c, 0xb1, 0x23, 0xd8, 0x8e, 0x19, 0xb9, 0x76, 0x0b, 0x4c, 0x75, 0x91, 0x37, 0xa5,
	0x79, 0x60, 0xd8, 0x2e, 0x66, 0x5a, 0xbb, 0x33, 0x53, 0xab, 0x94, 0x29, 0x7a, 0x06, 0xc8, 0x4b,
	0x18, 0xa7, 0x2c, 0x8d, 0x12, 0xe0, 0x11, 0x09, 0xec, 0xba, 0x72, 0xb2, 0xa3, 0x77, 0x54, 0xa0,
	0xef, 0xa4, 0x3d, 0x77, 0x26, 0xfc, 0x90, 0xd8, 0x6b, 0xdd, 0xea, 0x71, 0x3d, 0x75, 0x76, 0xe1,
	0x87, 0x44, 0x12, 0x96, 0xed, 0x10, 0x12, 0x76, 0x19, 0x10, 0x97, 0x51, 0x2a, 0xec, 0x86, 0x26,
	0x2c, 0xd8, 0xec, 0x7b, 0x65, 0x1d, 0x52, 0x2a, 0x50, 0x07, 0x5a, 0x63, 0x12, 0x90, 0x09, 0x16,
	0xc4, 0x6e, 0x2a, 0x40, 0xb6, 0x46, 0xef, 0xc3, 0xa6, 0xf9, 0x76, 0xb9, 0x3f, 0x89, 0xec, 0x96,
	0x02, 0xb4, 0x8d, 0xf1, 0xdc, 0x9f, 0x44, 0xe8, 0x08, 0xea, 0x82, 0xcd, 0xb8, 0xbd, 0xde, 0xb5,
	0x6e, 0x69, 0x0d, 0xb5, 0xef, 0x3c, 0x82, 0x3d, 0x55, 0xf1, 0xd3, 0x29, 0xf6, 0xa3, 0xd7, 0xd1,
	0x2f, 0x34, 0x6d, 0x0e, 0xe7, 0x0f, 0x0b, 0xf6, 0xcb, 0x3b, 0x2b, 0x88, 0xfc, 0x79, 0x49, 0x64,
	0x27, 0x17, 0xf9, 0x86, 0xe3, 0xb2, 0xd0, 0xbf, 0xd7, 0x0a, 0x57, 0x73, 0x6b, 0x4a, 0xf0, 0xd8,
	0x2d, 0xf7, 0x6b, 0x5b, 0x5a, 0x07, 0x46, 0xc9, 0x23, 0xd8, 0x2e, 0xa0, 0x0a, 0x8d, 0xbb, 0x99,
	0xc1, 0x94, 0xa4, 0x8b, 0x38, 0xa5, 0x94, 0xa5, 0x94, 0xca, 0x71, 0x4a, 0xad, 0x1e, 0x3c, 0x2c,
	0xe0, 0x32, 0x41, 0xea, 0xca, 0xe7, 0x83, 0x0c, 0xfb, 0xb5, 0x51, 0xe6, 0x29, 0xb4, 0xd3, 0x56,
	0xd1, 0x4d, 0xb2, 0xa6, 0x38, 0x6e, 0x68, 0x9b, 0xee, 0x8f, 0x4f, 0xc1, 0x0e, 0x30, 0x17, 0xae,
	0x27, 0x13, 0x8e, 0x78, 0xc2, 0x0b, 0x29, 0x35, 0x14, 0x7c, 0x4f, 0xee, 0x9f, 0x9a, 0x6d, 0x93,
	0x9b, 0xf3, 0x12, 0x1e, 0xaa, 0xaa, 0xbd, 0xf2, 0x3c, 0x9a, 0x44, 0xc2, 0xdc, 0xe1, 0xa7, 0xd0,
	0xc6, 0xda, 0xe2, 0x46, 0x38, 0x24, 0x69, 0xfb, 0x6f, 0xa4, 0xb6, 0x37, 0x38, 0x24, 0xce, 0xaf,
	0x35, 0xd8, 0x5d, 0x3c, 0xba, 0x82, 0x8e, 0x2f, 0x4b, 0x3a, 0x76, 0x73, 0x1d, 0x4b, 0x6e, 0xcb,
	0x2a, 0xfe, 0x56, 0xcd, 0x54, 0xbc, 0x9f, 0x2c, 0xda, 0x87, 0x46, 0x9c, 0x8c, 0x2e, 0xc9, 0x3c,
	0x0d, 0x9e, 0xae, 0x24, 0xd7, 0x11, 0x0e, 0x70, 0xe4, 0x19, 0xa9, 0xcc, 0x12, 0x7d, 0x00, 0x5b,
	0x5c, 0xe0, 0x4b, 0x32, 0x76, 0x0d, 0xa0, 0xae, 0xb5, 0xd4, 0xd6, 0x81, 0x36, 0x3a, 0x3f, 0xa7,
	0xf5, 0x7b, 0x3b, 0x92, 0xe3, 0x68, 0x30, 0x3f, 0x23, 0xf3, 0x21, 0xb9, 0x92, 0x17, 0xcd, 0xa3,
	0x91, 0x60, 0xd8, 0x13, 0x29, 0x9d, 0x6c, 0x2d, 0xb9, 0x50, 0x85, 0x36, 0x5c, 0xf4, 0x4a, 0x56,
	0x47, 0x12, 0xd4, 0x2f, 0x85, 0xfc, 0x74, 0xfe, 0xad, 0x82, 0x7d, 0xd3, 0xfb, 0x0a, 0x65, 0xfe,
	0xa2, 0x54, 0xe6, 0xc3, 0xbc, 0xcc, 0x4b, 0x5c, 0x97, 0x4b, 0x3d, 0xce, 0x2a, 0xfd, 0xbf, 0xa4,
	0x85, 0x76, 0x61, 0xed, 0x1a, 0x07, 0x89, 0xe9, 0x78, 0xbd, 0x38, 0xf9, 0xd3, 0x82, 0xe6, 0x29,
	0x65, 0xe4, 0x55, 0xec, 0xa3, 0x8f, 0xa1, 0x99, 0x8e, 0x57, 0xb4, 0xe4, 0x8d, 0xe9, 0xec, 0x2e,
	0x1b, 0xc0, 0x4e, 0x05, 0x7d, 0x06, 0x2d, 0x33, 0xfa, 0xd0, 0x6e, 0x9e, 0x62, 0x3e, 0x37, 0x3b,
	0x7b, 0x25, 0x6b, 0x76, 0xf4, 0x2b, 0x80, 0x7c, 0x44, 0xa0, 0xfd, 0x1b, 0x33, 0x43, 0x1f, 0x7f,
	0x74, 0xcb, 0x2c, 0x71, 0x2a, 0xe8, 0x0c, 0xb6, 0x16, 0x9f, 0x1f, 0xd4, 0x59, 0xfa, 0x26, 0x69,
	0x47, 0x07, 0x77, 0xbc, 0x57, 0x4e, 0x05, 0x7d, 0x03, 0xed, 0xe2, 0x1d, 0x40, 0xf6, 0x92, 0x6b,
	0xa1, 0x1d, 0x3d, 0xbe, 0xf5, 0xc2, 0x38, 0x15, 0x74, 0x9e, 0xbe, 0xb5, 0x46, 0xe3, 0x73, 0xc1,
	0xfc, 0x68, 0x72, 0x26, 0xfb, 0xfe, 0x96, 0x06, 0xb8, 0xea, 0x3c, 0xb9, 0xb3, 0x35, 0x9c, 0xca,
	0xe0, 0xf0, 0x27, 0xe7, 0xfe, 0xdf, 0xb8, 0x51, 0x43, 0xfd, 0x7a, 0xbd, 0xf8, 0x2f, 0x00, 0x00,
	0xff, 0xff, 0x89, 0xb3, 0xf1, 0x1f, 0xf3, 0x09, 0x00, 0x00,
}
