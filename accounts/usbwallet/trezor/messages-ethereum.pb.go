// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages-dogeum.proto

package trezor

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//*
// Request: Ask device for public key corresponding to address_n path
// @start
// @next DogeumPublicKey
// @next Failure
type DogeumGetPublicKey struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	ShowDisplay          *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumGetPublicKey) Reset()         { *m = DogeumGetPublicKey{} }
func (m *DogeumGetPublicKey) String() string { return proto.CompactTextString(m) }
func (*DogeumGetPublicKey) ProtoMessage()    {}
func (*DogeumGetPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{0}
}

func (m *DogeumGetPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumGetPublicKey.Unmarshal(m, b)
}
func (m *DogeumGetPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumGetPublicKey.Marshal(b, m, deterministic)
}
func (m *DogeumGetPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumGetPublicKey.Merge(m, src)
}
func (m *DogeumGetPublicKey) XXX_Size() int {
	return xxx_messageInfo_DogeumGetPublicKey.Size(m)
}
func (m *DogeumGetPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumGetPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumGetPublicKey proto.InternalMessageInfo

func (m *DogeumGetPublicKey) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *DogeumGetPublicKey) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

//*
// Response: Contains public key derived from device private seed
// @end
type DogeumPublicKey struct {
	Node                 *HDNodeType `protobuf:"bytes,1,opt,name=node" json:"node,omitempty"`
	Xpub                 *string     `protobuf:"bytes,2,opt,name=xpub" json:"xpub,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DogeumPublicKey) Reset()         { *m = DogeumPublicKey{} }
func (m *DogeumPublicKey) String() string { return proto.CompactTextString(m) }
func (*DogeumPublicKey) ProtoMessage()    {}
func (*DogeumPublicKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{1}
}

func (m *DogeumPublicKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumPublicKey.Unmarshal(m, b)
}
func (m *DogeumPublicKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumPublicKey.Marshal(b, m, deterministic)
}
func (m *DogeumPublicKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumPublicKey.Merge(m, src)
}
func (m *DogeumPublicKey) XXX_Size() int {
	return xxx_messageInfo_DogeumPublicKey.Size(m)
}
func (m *DogeumPublicKey) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumPublicKey.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumPublicKey proto.InternalMessageInfo

func (m *DogeumPublicKey) GetNode() *HDNodeType {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *DogeumPublicKey) GetXpub() string {
	if m != nil && m.Xpub != nil {
		return *m.Xpub
	}
	return ""
}

//*
// Request: Ask device for Dogeum address corresponding to address_n path
// @start
// @next DogeumAddress
// @next Failure
type DogeumGetAddress struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	ShowDisplay          *bool    `protobuf:"varint,2,opt,name=show_display,json=showDisplay" json:"show_display,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumGetAddress) Reset()         { *m = DogeumGetAddress{} }
func (m *DogeumGetAddress) String() string { return proto.CompactTextString(m) }
func (*DogeumGetAddress) ProtoMessage()    {}
func (*DogeumGetAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{2}
}

func (m *DogeumGetAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumGetAddress.Unmarshal(m, b)
}
func (m *DogeumGetAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumGetAddress.Marshal(b, m, deterministic)
}
func (m *DogeumGetAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumGetAddress.Merge(m, src)
}
func (m *DogeumGetAddress) XXX_Size() int {
	return xxx_messageInfo_DogeumGetAddress.Size(m)
}
func (m *DogeumGetAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumGetAddress.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumGetAddress proto.InternalMessageInfo

func (m *DogeumGetAddress) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *DogeumGetAddress) GetShowDisplay() bool {
	if m != nil && m.ShowDisplay != nil {
		return *m.ShowDisplay
	}
	return false
}

//*
// Response: Contains an Dogeum address derived from device private seed
// @end
type DogeumAddress struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	AddressHex           *string  `protobuf:"bytes,2,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumAddress) Reset()         { *m = DogeumAddress{} }
func (m *DogeumAddress) String() string { return proto.CompactTextString(m) }
func (*DogeumAddress) ProtoMessage()    {}
func (*DogeumAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{3}
}

func (m *DogeumAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumAddress.Unmarshal(m, b)
}
func (m *DogeumAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumAddress.Marshal(b, m, deterministic)
}
func (m *DogeumAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumAddress.Merge(m, src)
}
func (m *DogeumAddress) XXX_Size() int {
	return xxx_messageInfo_DogeumAddress.Size(m)
}
func (m *DogeumAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumAddress.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumAddress proto.InternalMessageInfo

func (m *DogeumAddress) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *DogeumAddress) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

//*
// Request: Ask device to sign transaction
// All fields are optional from the protocol's point of view. Each field defaults to value `0` if missing.
// Note: the first at most 1024 bytes of data MUST be transmitted as part of this message.
// @start
// @next DogeumTxRequest
// @next Failure
type DogeumSignTx struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Nonce                []byte   `protobuf:"bytes,2,opt,name=nonce" json:"nonce,omitempty"`
	GasPrice             []byte   `protobuf:"bytes,3,opt,name=gas_price,json=gasPrice" json:"gas_price,omitempty"`
	GasLimit             []byte   `protobuf:"bytes,4,opt,name=gas_limit,json=gasLimit" json:"gas_limit,omitempty"`
	ToBin                []byte   `protobuf:"bytes,5,opt,name=toBin" json:"toBin,omitempty"`
	ToHex                *string  `protobuf:"bytes,11,opt,name=toHex" json:"toHex,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value" json:"value,omitempty"`
	DataInitialChunk     []byte   `protobuf:"bytes,7,opt,name=data_initial_chunk,json=dataInitialChunk" json:"data_initial_chunk,omitempty"`
	DataLength           *uint32  `protobuf:"varint,8,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	ChainId              *uint32  `protobuf:"varint,9,opt,name=chain_id,json=chainId" json:"chain_id,omitempty"`
	TxType               *uint32  `protobuf:"varint,10,opt,name=tx_type,json=txType" json:"tx_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumSignTx) Reset()         { *m = DogeumSignTx{} }
func (m *DogeumSignTx) String() string { return proto.CompactTextString(m) }
func (*DogeumSignTx) ProtoMessage()    {}
func (*DogeumSignTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{4}
}

func (m *DogeumSignTx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumSignTx.Unmarshal(m, b)
}
func (m *DogeumSignTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumSignTx.Marshal(b, m, deterministic)
}
func (m *DogeumSignTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumSignTx.Merge(m, src)
}
func (m *DogeumSignTx) XXX_Size() int {
	return xxx_messageInfo_DogeumSignTx.Size(m)
}
func (m *DogeumSignTx) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumSignTx.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumSignTx proto.InternalMessageInfo

func (m *DogeumSignTx) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *DogeumSignTx) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *DogeumSignTx) GetGasPrice() []byte {
	if m != nil {
		return m.GasPrice
	}
	return nil
}

func (m *DogeumSignTx) GetGasLimit() []byte {
	if m != nil {
		return m.GasLimit
	}
	return nil
}

func (m *DogeumSignTx) GetToBin() []byte {
	if m != nil {
		return m.ToBin
	}
	return nil
}

func (m *DogeumSignTx) GetToHex() string {
	if m != nil && m.ToHex != nil {
		return *m.ToHex
	}
	return ""
}

func (m *DogeumSignTx) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *DogeumSignTx) GetDataInitialChunk() []byte {
	if m != nil {
		return m.DataInitialChunk
	}
	return nil
}

func (m *DogeumSignTx) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *DogeumSignTx) GetChainId() uint32 {
	if m != nil && m.ChainId != nil {
		return *m.ChainId
	}
	return 0
}

func (m *DogeumSignTx) GetTxType() uint32 {
	if m != nil && m.TxType != nil {
		return *m.TxType
	}
	return 0
}

//*
// Response: Device asks for more data from transaction payload, or returns the signature.
// If data_length is set, device awaits that many more bytes of payload.
// Otherwise, the signature_* fields contain the computed transaction signature. All three fields will be present.
// @end
// @next DogeumTxAck
type DogeumTxRequest struct {
	DataLength           *uint32  `protobuf:"varint,1,opt,name=data_length,json=dataLength" json:"data_length,omitempty"`
	SignatureV           *uint32  `protobuf:"varint,2,opt,name=signature_v,json=signatureV" json:"signature_v,omitempty"`
	SignatureR           []byte   `protobuf:"bytes,3,opt,name=signature_r,json=signatureR" json:"signature_r,omitempty"`
	SignatureS           []byte   `protobuf:"bytes,4,opt,name=signature_s,json=signatureS" json:"signature_s,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumTxRequest) Reset()         { *m = DogeumTxRequest{} }
func (m *DogeumTxRequest) String() string { return proto.CompactTextString(m) }
func (*DogeumTxRequest) ProtoMessage()    {}
func (*DogeumTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{5}
}

func (m *DogeumTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumTxRequest.Unmarshal(m, b)
}
func (m *DogeumTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumTxRequest.Marshal(b, m, deterministic)
}
func (m *DogeumTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumTxRequest.Merge(m, src)
}
func (m *DogeumTxRequest) XXX_Size() int {
	return xxx_messageInfo_DogeumTxRequest.Size(m)
}
func (m *DogeumTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumTxRequest proto.InternalMessageInfo

func (m *DogeumTxRequest) GetDataLength() uint32 {
	if m != nil && m.DataLength != nil {
		return *m.DataLength
	}
	return 0
}

func (m *DogeumTxRequest) GetSignatureV() uint32 {
	if m != nil && m.SignatureV != nil {
		return *m.SignatureV
	}
	return 0
}

func (m *DogeumTxRequest) GetSignatureR() []byte {
	if m != nil {
		return m.SignatureR
	}
	return nil
}

func (m *DogeumTxRequest) GetSignatureS() []byte {
	if m != nil {
		return m.SignatureS
	}
	return nil
}

//*
// Request: Transaction payload data.
// @next DogeumTxRequest
type DogeumTxAck struct {
	DataChunk            []byte   `protobuf:"bytes,1,opt,name=data_chunk,json=dataChunk" json:"data_chunk,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumTxAck) Reset()         { *m = DogeumTxAck{} }
func (m *DogeumTxAck) String() string { return proto.CompactTextString(m) }
func (*DogeumTxAck) ProtoMessage()    {}
func (*DogeumTxAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{6}
}

func (m *DogeumTxAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumTxAck.Unmarshal(m, b)
}
func (m *DogeumTxAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumTxAck.Marshal(b, m, deterministic)
}
func (m *DogeumTxAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumTxAck.Merge(m, src)
}
func (m *DogeumTxAck) XXX_Size() int {
	return xxx_messageInfo_DogeumTxAck.Size(m)
}
func (m *DogeumTxAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumTxAck.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumTxAck proto.InternalMessageInfo

func (m *DogeumTxAck) GetDataChunk() []byte {
	if m != nil {
		return m.DataChunk
	}
	return nil
}

//*
// Request: Ask device to sign message
// @start
// @next DogeumMessageSignature
// @next Failure
type DogeumSignMessage struct {
	AddressN             []uint32 `protobuf:"varint,1,rep,name=address_n,json=addressN" json:"address_n,omitempty"`
	Message              []byte   `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumSignMessage) Reset()         { *m = DogeumSignMessage{} }
func (m *DogeumSignMessage) String() string { return proto.CompactTextString(m) }
func (*DogeumSignMessage) ProtoMessage()    {}
func (*DogeumSignMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{7}
}

func (m *DogeumSignMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumSignMessage.Unmarshal(m, b)
}
func (m *DogeumSignMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumSignMessage.Marshal(b, m, deterministic)
}
func (m *DogeumSignMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumSignMessage.Merge(m, src)
}
func (m *DogeumSignMessage) XXX_Size() int {
	return xxx_messageInfo_DogeumSignMessage.Size(m)
}
func (m *DogeumSignMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumSignMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumSignMessage proto.InternalMessageInfo

func (m *DogeumSignMessage) GetAddressN() []uint32 {
	if m != nil {
		return m.AddressN
	}
	return nil
}

func (m *DogeumSignMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

//*
// Response: Signed message
// @end
type DogeumMessageSignature struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	AddressHex           *string  `protobuf:"bytes,3,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumMessageSignature) Reset()         { *m = DogeumMessageSignature{} }
func (m *DogeumMessageSignature) String() string { return proto.CompactTextString(m) }
func (*DogeumMessageSignature) ProtoMessage()    {}
func (*DogeumMessageSignature) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{8}
}

func (m *DogeumMessageSignature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumMessageSignature.Unmarshal(m, b)
}
func (m *DogeumMessageSignature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumMessageSignature.Marshal(b, m, deterministic)
}
func (m *DogeumMessageSignature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumMessageSignature.Merge(m, src)
}
func (m *DogeumMessageSignature) XXX_Size() int {
	return xxx_messageInfo_DogeumMessageSignature.Size(m)
}
func (m *DogeumMessageSignature) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumMessageSignature.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumMessageSignature proto.InternalMessageInfo

func (m *DogeumMessageSignature) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *DogeumMessageSignature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *DogeumMessageSignature) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

//*
// Request: Ask device to verify message
// @start
// @next Success
// @next Failure
type DogeumVerifyMessage struct {
	AddressBin           []byte   `protobuf:"bytes,1,opt,name=addressBin" json:"addressBin,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	Message              []byte   `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	AddressHex           *string  `protobuf:"bytes,4,opt,name=addressHex" json:"addressHex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DogeumVerifyMessage) Reset()         { *m = DogeumVerifyMessage{} }
func (m *DogeumVerifyMessage) String() string { return proto.CompactTextString(m) }
func (*DogeumVerifyMessage) ProtoMessage()    {}
func (*DogeumVerifyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb33f46ba915f15c, []int{9}
}

func (m *DogeumVerifyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DogeumVerifyMessage.Unmarshal(m, b)
}
func (m *DogeumVerifyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DogeumVerifyMessage.Marshal(b, m, deterministic)
}
func (m *DogeumVerifyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DogeumVerifyMessage.Merge(m, src)
}
func (m *DogeumVerifyMessage) XXX_Size() int {
	return xxx_messageInfo_DogeumVerifyMessage.Size(m)
}
func (m *DogeumVerifyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DogeumVerifyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DogeumVerifyMessage proto.InternalMessageInfo

func (m *DogeumVerifyMessage) GetAddressBin() []byte {
	if m != nil {
		return m.AddressBin
	}
	return nil
}

func (m *DogeumVerifyMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *DogeumVerifyMessage) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *DogeumVerifyMessage) GetAddressHex() string {
	if m != nil && m.AddressHex != nil {
		return *m.AddressHex
	}
	return ""
}

func init() {
	proto.RegisterType((*DogeumGetPublicKey)(nil), "hw.trezor.messages.dogeum.DogeumGetPublicKey")
	proto.RegisterType((*DogeumPublicKey)(nil), "hw.trezor.messages.dogeum.DogeumPublicKey")
	proto.RegisterType((*DogeumGetAddress)(nil), "hw.trezor.messages.dogeum.DogeumGetAddress")
	proto.RegisterType((*DogeumAddress)(nil), "hw.trezor.messages.dogeum.DogeumAddress")
	proto.RegisterType((*DogeumSignTx)(nil), "hw.trezor.messages.dogeum.DogeumSignTx")
	proto.RegisterType((*DogeumTxRequest)(nil), "hw.trezor.messages.dogeum.DogeumTxRequest")
	proto.RegisterType((*DogeumTxAck)(nil), "hw.trezor.messages.dogeum.DogeumTxAck")
	proto.RegisterType((*DogeumSignMessage)(nil), "hw.trezor.messages.dogeum.DogeumSignMessage")
	proto.RegisterType((*DogeumMessageSignature)(nil), "hw.trezor.messages.dogeum.DogeumMessageSignature")
	proto.RegisterType((*DogeumVerifyMessage)(nil), "hw.trezor.messages.dogeum.DogeumVerifyMessage")
}

func init() { proto.RegisterFile("messages-dogeum.proto", fileDescriptor_cb33f46ba915f15c) }

var fileDescriptor_cb33f46ba915f15c = []byte{
	// 593 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xb4, 0x49, 0x26, 0x0d, 0x1f, 0xa6, 0x55, 0x17, 0x0a, 0x34, 0x18, 0x21, 0xe5,
	0x00, 0x3e, 0x70, 0x43, 0xe2, 0xd2, 0x52, 0x44, 0x2b, 0x4a, 0x55, 0xdc, 0xa8, 0x57, 0x6b, 0x63,
	0x6f, 0xe3, 0x55, 0x9d, 0xdd, 0xe0, 0x5d, 0xb7, 0x0e, 0x7f, 0x82, 0x23, 0xff, 0x87, 0x5f, 0x86,
	0xf6, 0x2b, 0x71, 0x52, 0x54, 0x0e, 0xbd, 0x65, 0xde, 0xbc, 0x7d, 0xf3, 0x66, 0xf4, 0x62, 0xd8,
	0x99, 0x10, 0x21, 0xf0, 0x98, 0x88, 0x77, 0x44, 0x66, 0xa4, 0x20, 0xe5, 0x24, 0x9c, 0x16, 0x5c,
	0x72, 0x7f, 0x37, 0xbb, 0x09, 0x65, 0x41, 0x7e, 0xf2, 0x22, 0x74, 0x94, 0xd0, 0x51, 0x9e, 0x6d,
	0xcf, 0x5f, 0x25, 0x7c, 0x32, 0xe1, 0xcc, 0xbc, 0x09, 0x2e, 0x60, 0xeb, 0xb3, 0xa5, 0x7c, 0x21,
	0xf2, 0xac, 0x1c, 0xe5, 0x34, 0xf9, 0x4a, 0x66, 0xfe, 0x2e, 0x74, 0x70, 0x9a, 0x16, 0x44, 0x88,
	0x98, 0x21, 0xaf, 0xdf, 0x18, 0xf4, 0xa2, 0xb6, 0x05, 0x4e, 0xfd, 0x57, 0xb0, 0x29, 0x32, 0x7e,
	0x13, 0xa7, 0x54, 0x4c, 0x73, 0x3c, 0x43, 0x6b, 0x7d, 0x6f, 0xd0, 0x8e, 0xba, 0x0a, 0x3b, 0x34,
	0x50, 0x30, 0x82, 0xc7, 0x4e, 0x77, 0x21, 0xfa, 0x01, 0x9a, 0x8c, 0xa7, 0x04, 0x79, 0x7d, 0x6f,
	0xd0, 0x7d, 0xff, 0x26, 0xfc, 0x87, 0x5f, 0x6b, 0xee, 0xe8, 0xf0, 0x94, 0xa7, 0x64, 0x38, 0x9b,
	0x92, 0x48, 0x3f, 0xf1, 0x7d, 0x68, 0x56, 0xd3, 0x72, 0xa4, 0x47, 0x75, 0x22, 0xfd, 0x3b, 0x18,
	0x82, 0x5f, 0xf3, 0xbe, 0x6f, 0xdc, 0xdd, 0xdb, 0xf9, 0x77, 0x78, 0xe8, 0x54, 0x9d, 0xe4, 0x4b,
	0x00, 0xab, 0x70, 0x40, 0x99, 0x76, 0xbf, 0x19, 0xd5, 0x90, 0x5a, 0xff, 0x88, 0x54, 0xd6, 0x62,
	0x0d, 0x09, 0xfe, 0xac, 0xc1, 0x03, 0xa7, 0x79, 0x4e, 0xc7, 0x6c, 0x58, 0xdd, 0xed, 0x72, 0x0b,
	0xd6, 0x19, 0x67, 0x09, 0xd1, 0x52, 0x9b, 0x91, 0x29, 0xd4, 0x93, 0x31, 0x16, 0xf1, 0xb4, 0xa0,
	0x09, 0x41, 0x0d, 0xdd, 0x69, 0x8f, 0xb1, 0x38, 0x53, 0xb5, 0x6b, 0xe6, 0x74, 0x42, 0x25, 0x6a,
	0xce, 0x9b, 0x27, 0xaa, 0x56, 0x7a, 0x92, 0x2b, 0xeb, 0xeb, 0x46, 0x4f, 0x17, 0x06, 0x55, 0x86,
	0xbb, 0xda, 0xb0, 0x29, 0x14, 0x7a, 0x8d, 0xf3, 0x92, 0xa0, 0x0d, 0xc3, 0xd5, 0x85, 0xff, 0x16,
	0xfc, 0x14, 0x4b, 0x1c, 0x53, 0x46, 0x25, 0xc5, 0x79, 0x9c, 0x64, 0x25, 0xbb, 0x42, 0x2d, 0x4d,
	0x79, 0xa4, 0x3a, 0xc7, 0xa6, 0xf1, 0x49, 0xe1, 0xfe, 0x1e, 0x74, 0x35, 0x3b, 0x27, 0x6c, 0x2c,
	0x33, 0xd4, 0xee, 0x7b, 0x83, 0x5e, 0x04, 0x0a, 0x3a, 0xd1, 0x88, 0xff, 0x14, 0xda, 0x49, 0x86,
	0x29, 0x8b, 0x69, 0x8a, 0x3a, 0xba, 0xdb, 0xd2, 0xf5, 0x71, 0xea, 0xef, 0x40, 0x4b, 0x56, 0xb1,
	0x9c, 0x4d, 0x09, 0x02, 0xdd, 0xd9, 0x90, 0x95, 0xca, 0x41, 0xf0, 0xdb, 0x5b, 0x44, 0x6a, 0x58,
	0x45, 0xe4, 0x47, 0x49, 0x84, 0x5c, 0x1d, 0xe5, 0xdd, 0x1a, 0xb5, 0x07, 0x5d, 0x41, 0xc7, 0x0c,
	0xcb, 0xb2, 0x20, 0xf1, 0xb5, 0xbe, 0x68, 0x2f, 0x82, 0x39, 0x74, 0xb1, 0x4c, 0x28, 0xec, 0x61,
	0x17, 0x84, 0x68, 0x99, 0x20, 0xec, 0x71, 0x17, 0x84, 0xf3, 0x20, 0x84, 0xde, 0xc2, 0xd8, 0x7e,
	0x72, 0xe5, 0xbf, 0x00, 0xed, 0xc0, 0x5e, 0xc9, 0xe4, 0xa5, 0xa3, 0x10, 0x7d, 0x9e, 0xe0, 0x04,
	0x9e, 0xd4, 0xd3, 0xf0, 0xcd, 0x64, 0xff, 0xee, 0x48, 0x20, 0x68, 0xd9, 0xff, 0x88, 0x0d, 0x85,
	0x2b, 0x83, 0x0a, 0x90, 0x53, 0xb3, 0x4a, 0xe7, 0xce, 0xda, 0x7f, 0x83, 0xfb, 0x1c, 0x3a, 0xf3,
	0x3d, 0xac, 0xee, 0x02, 0x58, 0x89, 0x75, 0xe3, 0x56, 0xac, 0x7f, 0x79, 0xb0, 0xed, 0x46, 0x5f,
	0x90, 0x82, 0x5e, 0xce, 0xdc, 0x2a, 0xf7, 0x9b, 0x5b, 0xdb, 0xb5, 0xb1, 0xb4, 0xeb, 0x8a, 0xa3,
	0xe6, 0xaa, 0xa3, 0x83, 0x8f, 0xf0, 0x3a, 0xe1, 0x93, 0x50, 0x60, 0xc9, 0x45, 0x46, 0x73, 0x3c,
	0x12, 0xee, 0x03, 0x93, 0xd3, 0x91, 0xf9, 0xe2, 0x8d, 0xca, 0xcb, 0x83, 0xed, 0xa1, 0x06, 0xad,
	0x5b, 0xb7, 0xc2, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x8a, 0xce, 0x81, 0xc8, 0x59, 0x05, 0x00,
	0x00,
}
