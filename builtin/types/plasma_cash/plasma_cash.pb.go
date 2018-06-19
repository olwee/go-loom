// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

/*
Package plasma_cash is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto

It has these top-level messages:
	PlasmaBlock
	PlasmaTx
	ListTransactionsRequest
	ListTransactionsResponse
	SubmitBlockToMainnetRequest
	SubmitBlockToMainnetResponse
	PlasmaTxRequest
	PlasmaTxResponse
	DepositRequest
	DepositResponse
	GetProofRequest
	GetProofResponse
	PlasmaCashParams
	PlasmaCashInitRequest
*/
package plasma_cash

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type PlasmaBlock struct {
	Uid          *types.BigUInt `protobuf:"bytes,1,opt,name=uid" json:"uid,omitempty"`
	Transactions []*PlasmaTx    `protobuf:"bytes,2,rep,name=transactions" json:"transactions,omitempty"`
	Signature    []byte         `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	MerkleHash   []byte         `protobuf:"bytes,4,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Hash         []byte         `protobuf:"bytes,5,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *PlasmaBlock) Reset()                    { *m = PlasmaBlock{} }
func (m *PlasmaBlock) String() string            { return proto.CompactTextString(m) }
func (*PlasmaBlock) ProtoMessage()               {}
func (*PlasmaBlock) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{0} }

func (m *PlasmaBlock) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *PlasmaBlock) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *PlasmaBlock) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaBlock) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaBlock) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type PlasmaTx struct {
	Uid           uint64         `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	PreviousBlock *types.BigUInt `protobuf:"bytes,2,opt,name=previous_block,json=previousBlock" json:"previous_block,omitempty"`
	Denomination  *types.BigUInt `protobuf:"bytes,3,opt,name=denomination" json:"denomination,omitempty"`
	NewOwner      *types.Address `protobuf:"bytes,4,opt,name=new_owner,json=newOwner" json:"new_owner,omitempty"`
	Signature     []byte         `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
	Hash          []byte         `protobuf:"bytes,6,opt,name=hash,proto3" json:"hash,omitempty"`
	MerkleHash    []byte         `protobuf:"bytes,7,opt,name=merkle_hash,json=merkleHash,proto3" json:"merkle_hash,omitempty"`
	Sender        *types.Address `protobuf:"bytes,8,opt,name=sender" json:"sender,omitempty"`
}

func (m *PlasmaTx) Reset()                    { *m = PlasmaTx{} }
func (m *PlasmaTx) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTx) ProtoMessage()               {}
func (*PlasmaTx) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{1} }

func (m *PlasmaTx) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PlasmaTx) GetPreviousBlock() *types.BigUInt {
	if m != nil {
		return m.PreviousBlock
	}
	return nil
}

func (m *PlasmaTx) GetDenomination() *types.BigUInt {
	if m != nil {
		return m.Denomination
	}
	return nil
}

func (m *PlasmaTx) GetNewOwner() *types.Address {
	if m != nil {
		return m.NewOwner
	}
	return nil
}

func (m *PlasmaTx) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PlasmaTx) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *PlasmaTx) GetMerkleHash() []byte {
	if m != nil {
		return m.MerkleHash
	}
	return nil
}

func (m *PlasmaTx) GetSender() *types.Address {
	if m != nil {
		return m.Sender
	}
	return nil
}

type ListTransactionsRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
}

func (m *ListTransactionsRequest) Reset()         { *m = ListTransactionsRequest{} }
func (m *ListTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsRequest) ProtoMessage()    {}
func (*ListTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{2}
}

func (m *ListTransactionsRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

type ListTransactionsResponse struct {
	Transactions []*PlasmaTx `protobuf:"bytes,1,rep,name=transactions" json:"transactions,omitempty"`
}

func (m *ListTransactionsResponse) Reset()         { *m = ListTransactionsResponse{} }
func (m *ListTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*ListTransactionsResponse) ProtoMessage()    {}
func (*ListTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{3}
}

func (m *ListTransactionsResponse) GetTransactions() []*PlasmaTx {
	if m != nil {
		return m.Transactions
	}
	return nil
}

// This only originates from the validator
type SubmitBlockToMainnetRequest struct {
}

func (m *SubmitBlockToMainnetRequest) Reset()         { *m = SubmitBlockToMainnetRequest{} }
func (m *SubmitBlockToMainnetRequest) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetRequest) ProtoMessage()    {}
func (*SubmitBlockToMainnetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{4}
}

type SubmitBlockToMainnetResponse struct {
}

func (m *SubmitBlockToMainnetResponse) Reset()         { *m = SubmitBlockToMainnetResponse{} }
func (m *SubmitBlockToMainnetResponse) String() string { return proto.CompactTextString(m) }
func (*SubmitBlockToMainnetResponse) ProtoMessage()    {}
func (*SubmitBlockToMainnetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptorPlasmaCash, []int{5}
}

type PlasmaTxRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *PlasmaTxRequest) Reset()                    { *m = PlasmaTxRequest{} }
func (m *PlasmaTxRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxRequest) ProtoMessage()               {}
func (*PlasmaTxRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{6} }

func (m *PlasmaTxRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type PlasmaTxResponse struct {
}

func (m *PlasmaTxResponse) Reset()                    { *m = PlasmaTxResponse{} }
func (m *PlasmaTxResponse) String() string            { return proto.CompactTextString(m) }
func (*PlasmaTxResponse) ProtoMessage()               {}
func (*PlasmaTxResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{7} }

// This only originates from the validator
type DepositRequest struct {
	Plasmatx *PlasmaTx `protobuf:"bytes,1,opt,name=plasmatx" json:"plasmatx,omitempty"`
}

func (m *DepositRequest) Reset()                    { *m = DepositRequest{} }
func (m *DepositRequest) String() string            { return proto.CompactTextString(m) }
func (*DepositRequest) ProtoMessage()               {}
func (*DepositRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{8} }

func (m *DepositRequest) GetPlasmatx() *PlasmaTx {
	if m != nil {
		return m.Plasmatx
	}
	return nil
}

type DepositResponse struct {
}

func (m *DepositResponse) Reset()                    { *m = DepositResponse{} }
func (m *DepositResponse) String() string            { return proto.CompactTextString(m) }
func (*DepositResponse) ProtoMessage()               {}
func (*DepositResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{9} }

type GetProofRequest struct {
	BlockHeight *types.BigUInt `protobuf:"bytes,1,opt,name=block_height,json=blockHeight" json:"block_height,omitempty"`
	Uid         *types.BigUInt `protobuf:"bytes,2,opt,name=uid" json:"uid,omitempty"`
}

func (m *GetProofRequest) Reset()                    { *m = GetProofRequest{} }
func (m *GetProofRequest) String() string            { return proto.CompactTextString(m) }
func (*GetProofRequest) ProtoMessage()               {}
func (*GetProofRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{10} }

func (m *GetProofRequest) GetBlockHeight() *types.BigUInt {
	if m != nil {
		return m.BlockHeight
	}
	return nil
}

func (m *GetProofRequest) GetUid() *types.BigUInt {
	if m != nil {
		return m.Uid
	}
	return nil
}

type GetProofResponse struct {
	Proof []byte `protobuf:"bytes,1,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *GetProofResponse) Reset()                    { *m = GetProofResponse{} }
func (m *GetProofResponse) String() string            { return proto.CompactTextString(m) }
func (*GetProofResponse) ProtoMessage()               {}
func (*GetProofResponse) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{11} }

func (m *GetProofResponse) GetProof() []byte {
	if m != nil {
		return m.Proof
	}
	return nil
}

// Initialization of state from Genesis.json
type PlasmaCashParams struct {
	BlockInterval uint64 `protobuf:"varint,1,opt,name=block_interval,json=blockInterval,proto3" json:"block_interval,omitempty"`
}

func (m *PlasmaCashParams) Reset()                    { *m = PlasmaCashParams{} }
func (m *PlasmaCashParams) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashParams) ProtoMessage()               {}
func (*PlasmaCashParams) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{12} }

func (m *PlasmaCashParams) GetBlockInterval() uint64 {
	if m != nil {
		return m.BlockInterval
	}
	return 0
}

type PlasmaCashInitRequest struct {
	Params *PlasmaCashParams `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
}

func (m *PlasmaCashInitRequest) Reset()                    { *m = PlasmaCashInitRequest{} }
func (m *PlasmaCashInitRequest) String() string            { return proto.CompactTextString(m) }
func (*PlasmaCashInitRequest) ProtoMessage()               {}
func (*PlasmaCashInitRequest) Descriptor() ([]byte, []int) { return fileDescriptorPlasmaCash, []int{13} }

func (m *PlasmaCashInitRequest) GetParams() *PlasmaCashParams {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*PlasmaBlock)(nil), "PlasmaBlock")
	proto.RegisterType((*PlasmaTx)(nil), "PlasmaTx")
	proto.RegisterType((*ListTransactionsRequest)(nil), "ListTransactionsRequest")
	proto.RegisterType((*ListTransactionsResponse)(nil), "ListTransactionsResponse")
	proto.RegisterType((*SubmitBlockToMainnetRequest)(nil), "SubmitBlockToMainnetRequest")
	proto.RegisterType((*SubmitBlockToMainnetResponse)(nil), "SubmitBlockToMainnetResponse")
	proto.RegisterType((*PlasmaTxRequest)(nil), "PlasmaTxRequest")
	proto.RegisterType((*PlasmaTxResponse)(nil), "PlasmaTxResponse")
	proto.RegisterType((*DepositRequest)(nil), "DepositRequest")
	proto.RegisterType((*DepositResponse)(nil), "DepositResponse")
	proto.RegisterType((*GetProofRequest)(nil), "GetProofRequest")
	proto.RegisterType((*GetProofResponse)(nil), "GetProofResponse")
	proto.RegisterType((*PlasmaCashParams)(nil), "PlasmaCashParams")
	proto.RegisterType((*PlasmaCashInitRequest)(nil), "PlasmaCashInitRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/plasma_cash/plasma_cash.proto", fileDescriptorPlasmaCash)
}

var fileDescriptorPlasmaCash = []byte{
	// 591 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xdd, 0x6e, 0xd3, 0x30,
	0x14, 0x56, 0xbb, 0xad, 0x64, 0xa7, 0xd9, 0x9f, 0x05, 0x22, 0x1a, 0x83, 0x55, 0x91, 0x26, 0x15,
	0xc1, 0x1a, 0xb4, 0x5d, 0x00, 0x97, 0x04, 0x04, 0x2b, 0xe2, 0xa7, 0x0a, 0xe3, 0x86, 0x9b, 0xca,
	0x69, 0xbd, 0xc4, 0x6a, 0x62, 0x07, 0xdb, 0x59, 0xc7, 0x33, 0xf1, 0x02, 0x3c, 0xcd, 0x2e, 0xf6,
	0x24, 0x28, 0x8e, 0xdb, 0xa4, 0xed, 0x04, 0x82, 0x9b, 0xca, 0xe7, 0x7c, 0x3e, 0xdf, 0x39, 0xdf,
	0x17, 0x9f, 0xc2, 0xfb, 0x88, 0xaa, 0x38, 0x0f, 0x7b, 0x23, 0x9e, 0x7a, 0x09, 0xe7, 0x29, 0x23,
	0x6a, 0xca, 0xc5, 0xc4, 0x8b, 0xf8, 0x71, 0x11, 0x7a, 0x61, 0x4e, 0x13, 0x45, 0x99, 0xa7, 0x7e,
	0x64, 0x44, 0x7a, 0x59, 0x82, 0x65, 0x8a, 0x87, 0x23, 0x2c, 0xe3, 0xfa, 0xb9, 0x97, 0x09, 0xae,
	0xf8, 0xfe, 0x71, 0x8d, 0x2b, 0xe2, 0x11, 0xf7, 0x74, 0x3a, 0xcc, 0x2f, 0x74, 0xa4, 0x03, 0x7d,
	0x32, 0xd7, 0x9f, 0xfd, 0xa5, 0x75, 0xd9, 0x52, 0xff, 0x96, 0x15, 0xee, 0xcf, 0x06, 0xb4, 0x07,
	0xba, 0xad, 0x9f, 0xf0, 0xd1, 0x04, 0xed, 0xc3, 0x5a, 0x4e, 0xc7, 0x4e, 0xa3, 0xd3, 0xe8, 0xb6,
	0x4f, 0xac, 0x9e, 0x4f, 0xa3, 0xaf, 0x7d, 0xa6, 0x82, 0x22, 0x89, 0x8e, 0xc1, 0x56, 0x02, 0x33,
	0x89, 0x47, 0x8a, 0x72, 0x26, 0x9d, 0x66, 0x67, 0xad, 0xdb, 0x3e, 0xd9, 0xec, 0x95, 0xf5, 0xe7,
	0x57, 0xc1, 0x02, 0x8c, 0x0e, 0x60, 0x53, 0xd2, 0x88, 0x61, 0x95, 0x0b, 0xe2, 0xac, 0x75, 0x1a,
	0x5d, 0x3b, 0xa8, 0x12, 0xe8, 0x10, 0xda, 0x29, 0x11, 0x93, 0x84, 0x0c, 0x63, 0x2c, 0x63, 0x67,
	0x5d, 0xe3, 0x50, 0xa6, 0xce, 0xb0, 0x8c, 0x11, 0x82, 0x75, 0x8d, 0x6c, 0x68, 0x44, 0x9f, 0xdd,
	0x5f, 0x4d, 0xb0, 0x66, 0xdd, 0xd0, 0x6e, 0x35, 0xea, 0x7a, 0x39, 0xa0, 0x0f, 0xdb, 0x99, 0x20,
	0x97, 0x94, 0xe7, 0x72, 0x18, 0x16, 0x72, 0x9c, 0xe6, 0xa2, 0x0e, 0x7f, 0xef, 0xe6, 0xfa, 0x70,
	0x6b, 0x60, 0xee, 0x68, 0xc5, 0xc1, 0x56, 0x56, 0x0f, 0xd1, 0x53, 0xb0, 0xc7, 0x84, 0xf1, 0x94,
	0x32, 0x5c, 0xc8, 0xd0, 0x83, 0xd7, 0x9d, 0x58, 0x40, 0xd1, 0x29, 0x6c, 0x32, 0x32, 0x1d, 0xf2,
	0x29, 0x23, 0x42, 0x6b, 0x28, 0xae, 0xbe, 0x1a, 0x8f, 0x05, 0x91, 0xd2, 0xb7, 0x6f, 0xae, 0x0f,
	0xad, 0x4f, 0x64, 0xfa, 0xb9, 0x40, 0x03, 0x8b, 0x99, 0xd3, 0xa2, 0x31, 0x1b, 0xcb, 0xc6, 0xcc,
	0x74, 0xb7, 0x2a, 0xdd, 0xcb, 0x66, 0xdd, 0x59, 0x31, 0xab, 0x03, 0x2d, 0x49, 0xd8, 0x98, 0x08,
	0xc7, 0x5a, 0x1c, 0x22, 0x30, 0x79, 0xf7, 0x2d, 0xdc, 0xff, 0x40, 0xa5, 0x3a, 0xaf, 0x7d, 0xa1,
	0x80, 0x7c, 0xcf, 0x89, 0x54, 0xe8, 0x09, 0xd8, 0xda, 0xad, 0x61, 0x4c, 0x68, 0x14, 0xab, 0x95,
	0x8f, 0xdf, 0xd6, 0xe8, 0x99, 0x06, 0xdd, 0x3e, 0x38, 0xab, 0x3c, 0x32, 0xe3, 0x4c, 0x92, 0x95,
	0x07, 0xd2, 0xf8, 0xe3, 0x03, 0x71, 0x1f, 0xc2, 0x83, 0x2f, 0x79, 0x98, 0x52, 0xa5, 0x9d, 0x3f,
	0xe7, 0x1f, 0x31, 0x65, 0x8c, 0x28, 0x33, 0x96, 0xfb, 0x08, 0x0e, 0x6e, 0x87, 0xcb, 0x6e, 0xee,
	0x0b, 0xd8, 0x99, 0x13, 0x1b, 0x25, 0x47, 0x60, 0x95, 0x3b, 0xa4, 0xae, 0x8c, 0x8a, 0x5a, 0xf3,
	0x39, 0xe4, 0x22, 0xd8, 0xad, 0x2a, 0x0d, 0xdb, 0x73, 0xd8, 0x7e, 0x43, 0x32, 0x2e, 0xa9, 0xfa,
	0x47, 0xb2, 0x3d, 0xd8, 0x99, 0x17, 0x1a, 0xae, 0x6f, 0xb0, 0xf3, 0x8e, 0xa8, 0x81, 0xe0, 0xfc,
	0xe2, 0x7f, 0x3c, 0x9e, 0x2d, 0x61, 0xf3, 0x96, 0x25, 0x74, 0xbb, 0xb0, 0x5b, 0x71, 0x1b, 0xdf,
	0xef, 0xc2, 0x46, 0x56, 0x24, 0x34, 0xab, 0x1d, 0x94, 0x81, 0xfb, 0x72, 0xa6, 0xf2, 0x35, 0x96,
	0xf1, 0x00, 0x0b, 0x9c, 0x4a, 0x74, 0x04, 0xdb, 0xe5, 0x18, 0x94, 0x29, 0x22, 0x2e, 0x71, 0x62,
	0xd6, 0x67, 0x4b, 0x67, 0xfb, 0x26, 0xe9, 0xfa, 0x70, 0xaf, 0x2a, 0xed, 0xb3, 0xca, 0x93, 0xc7,
	0xd0, 0xca, 0x34, 0x93, 0x11, 0xb0, 0xd7, 0x5b, 0x6e, 0x11, 0x98, 0x0b, 0x61, 0x4b, 0xff, 0xc1,
	0x9c, 0xfe, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x77, 0x1a, 0x27, 0x1d, 0x0f, 0x05, 0x00, 0x00,
}