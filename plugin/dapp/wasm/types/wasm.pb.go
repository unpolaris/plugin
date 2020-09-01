// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wasm.proto

package types

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type WasmAction struct {
	// Types that are valid to be assigned to Value:
	//	*WasmAction_Create
	//	*WasmAction_Call
	Value                isWasmAction_Value `protobuf_oneof:"value"`
	Ty                   int32              `protobuf:"varint,3,opt,name=ty,proto3" json:"ty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *WasmAction) Reset()         { *m = WasmAction{} }
func (m *WasmAction) String() string { return proto.CompactTextString(m) }
func (*WasmAction) ProtoMessage()    {}
func (*WasmAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{0}
}

func (m *WasmAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmAction.Unmarshal(m, b)
}
func (m *WasmAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmAction.Marshal(b, m, deterministic)
}
func (m *WasmAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmAction.Merge(m, src)
}
func (m *WasmAction) XXX_Size() int {
	return xxx_messageInfo_WasmAction.Size(m)
}
func (m *WasmAction) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmAction.DiscardUnknown(m)
}

var xxx_messageInfo_WasmAction proto.InternalMessageInfo

type isWasmAction_Value interface {
	isWasmAction_Value()
}

type WasmAction_Create struct {
	Create *WasmCreate `protobuf:"bytes,1,opt,name=create,proto3,oneof"`
}

type WasmAction_Call struct {
	Call *WasmCall `protobuf:"bytes,2,opt,name=call,proto3,oneof"`
}

func (*WasmAction_Create) isWasmAction_Value() {}

func (*WasmAction_Call) isWasmAction_Value() {}

func (m *WasmAction) GetValue() isWasmAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *WasmAction) GetCreate() *WasmCreate {
	if x, ok := m.GetValue().(*WasmAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *WasmAction) GetCall() *WasmCall {
	if x, ok := m.GetValue().(*WasmAction_Call); ok {
		return x.Call
	}
	return nil
}

func (m *WasmAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WasmAction) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WasmAction_Create)(nil),
		(*WasmAction_Call)(nil),
	}
}

type WasmCreate struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 []byte   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmCreate) Reset()         { *m = WasmCreate{} }
func (m *WasmCreate) String() string { return proto.CompactTextString(m) }
func (*WasmCreate) ProtoMessage()    {}
func (*WasmCreate) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{1}
}

func (m *WasmCreate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmCreate.Unmarshal(m, b)
}
func (m *WasmCreate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmCreate.Marshal(b, m, deterministic)
}
func (m *WasmCreate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmCreate.Merge(m, src)
}
func (m *WasmCreate) XXX_Size() int {
	return xxx_messageInfo_WasmCreate.Size(m)
}
func (m *WasmCreate) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmCreate.DiscardUnknown(m)
}

var xxx_messageInfo_WasmCreate proto.InternalMessageInfo

func (m *WasmCreate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WasmCreate) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

type WasmCall struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Parameters           []int64  `protobuf:"varint,3,rep,packed,name=parameters,proto3" json:"parameters,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WasmCall) Reset()         { *m = WasmCall{} }
func (m *WasmCall) String() string { return proto.CompactTextString(m) }
func (*WasmCall) ProtoMessage()    {}
func (*WasmCall) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{2}
}

func (m *WasmCall) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WasmCall.Unmarshal(m, b)
}
func (m *WasmCall) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WasmCall.Marshal(b, m, deterministic)
}
func (m *WasmCall) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WasmCall.Merge(m, src)
}
func (m *WasmCall) XXX_Size() int {
	return xxx_messageInfo_WasmCall.Size(m)
}
func (m *WasmCall) XXX_DiscardUnknown() {
	xxx_messageInfo_WasmCall.DiscardUnknown(m)
}

var xxx_messageInfo_WasmCall proto.InternalMessageInfo

func (m *WasmCall) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *WasmCall) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *WasmCall) GetParameters() []int64 {
	if m != nil {
		return m.Parameters
	}
	return nil
}

type QueryCheckConract struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCheckConract) Reset()         { *m = QueryCheckConract{} }
func (m *QueryCheckConract) String() string { return proto.CompactTextString(m) }
func (*QueryCheckConract) ProtoMessage()    {}
func (*QueryCheckConract) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{3}
}

func (m *QueryCheckConract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCheckConract.Unmarshal(m, b)
}
func (m *QueryCheckConract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCheckConract.Marshal(b, m, deterministic)
}
func (m *QueryCheckConract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCheckConract.Merge(m, src)
}
func (m *QueryCheckConract) XXX_Size() int {
	return xxx_messageInfo_QueryCheckConract.Size(m)
}
func (m *QueryCheckConract) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCheckConract.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCheckConract proto.InternalMessageInfo

func (m *QueryCheckConract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryCreateTransaction struct {
	Ty                   int32    `protobuf:"varint,1,opt,name=ty,proto3" json:"ty,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code                 []byte   `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Method               string   `protobuf:"bytes,4,opt,name=method,proto3" json:"method,omitempty"`
	Parameters           []int64  `protobuf:"varint,5,rep,packed,name=parameters,proto3" json:"parameters,omitempty"`
	Fee                  int64    `protobuf:"varint,6,opt,name=fee,proto3" json:"fee,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryCreateTransaction) Reset()         { *m = QueryCreateTransaction{} }
func (m *QueryCreateTransaction) String() string { return proto.CompactTextString(m) }
func (*QueryCreateTransaction) ProtoMessage()    {}
func (*QueryCreateTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{4}
}

func (m *QueryCreateTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryCreateTransaction.Unmarshal(m, b)
}
func (m *QueryCreateTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryCreateTransaction.Marshal(b, m, deterministic)
}
func (m *QueryCreateTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryCreateTransaction.Merge(m, src)
}
func (m *QueryCreateTransaction) XXX_Size() int {
	return xxx_messageInfo_QueryCreateTransaction.Size(m)
}
func (m *QueryCreateTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryCreateTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_QueryCreateTransaction proto.InternalMessageInfo

func (m *QueryCreateTransaction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

func (m *QueryCreateTransaction) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryCreateTransaction) GetCode() []byte {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *QueryCreateTransaction) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *QueryCreateTransaction) GetParameters() []int64 {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *QueryCreateTransaction) GetFee() int64 {
	if m != nil {
		return m.Fee
	}
	return 0
}

type CustomLog struct {
	Info                 []string `protobuf:"bytes,1,rep,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CustomLog) Reset()         { *m = CustomLog{} }
func (m *CustomLog) String() string { return proto.CompactTextString(m) }
func (*CustomLog) ProtoMessage()    {}
func (*CustomLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{5}
}

func (m *CustomLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CustomLog.Unmarshal(m, b)
}
func (m *CustomLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CustomLog.Marshal(b, m, deterministic)
}
func (m *CustomLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CustomLog.Merge(m, src)
}
func (m *CustomLog) XXX_Size() int {
	return xxx_messageInfo_CustomLog.Size(m)
}
func (m *CustomLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CustomLog.DiscardUnknown(m)
}

var xxx_messageInfo_CustomLog proto.InternalMessageInfo

func (m *CustomLog) GetInfo() []string {
	if m != nil {
		return m.Info
	}
	return nil
}

type CreateContractLog struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateContractLog) Reset()         { *m = CreateContractLog{} }
func (m *CreateContractLog) String() string { return proto.CompactTextString(m) }
func (*CreateContractLog) ProtoMessage()    {}
func (*CreateContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{6}
}

func (m *CreateContractLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateContractLog.Unmarshal(m, b)
}
func (m *CreateContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateContractLog.Marshal(b, m, deterministic)
}
func (m *CreateContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateContractLog.Merge(m, src)
}
func (m *CreateContractLog) XXX_Size() int {
	return xxx_messageInfo_CreateContractLog.Size(m)
}
func (m *CreateContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_CreateContractLog proto.InternalMessageInfo

func (m *CreateContractLog) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateContractLog) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type CallContractLog struct {
	Contract             string   `protobuf:"bytes,1,opt,name=contract,proto3" json:"contract,omitempty"`
	Method               string   `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Result               int64    `protobuf:"varint,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CallContractLog) Reset()         { *m = CallContractLog{} }
func (m *CallContractLog) String() string { return proto.CompactTextString(m) }
func (*CallContractLog) ProtoMessage()    {}
func (*CallContractLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{7}
}

func (m *CallContractLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallContractLog.Unmarshal(m, b)
}
func (m *CallContractLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallContractLog.Marshal(b, m, deterministic)
}
func (m *CallContractLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallContractLog.Merge(m, src)
}
func (m *CallContractLog) XXX_Size() int {
	return xxx_messageInfo_CallContractLog.Size(m)
}
func (m *CallContractLog) XXX_DiscardUnknown() {
	xxx_messageInfo_CallContractLog.DiscardUnknown(m)
}

var xxx_messageInfo_CallContractLog proto.InternalMessageInfo

func (m *CallContractLog) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func (m *CallContractLog) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *CallContractLog) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type LocalDataLog struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	PreValue             []byte   `protobuf:"bytes,2,opt,name=preValue,proto3" json:"preValue,omitempty"`
	CurValue             []byte   `protobuf:"bytes,3,opt,name=curValue,proto3" json:"curValue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LocalDataLog) Reset()         { *m = LocalDataLog{} }
func (m *LocalDataLog) String() string { return proto.CompactTextString(m) }
func (*LocalDataLog) ProtoMessage()    {}
func (*LocalDataLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_7d78909ad64e3bbb, []int{8}
}

func (m *LocalDataLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LocalDataLog.Unmarshal(m, b)
}
func (m *LocalDataLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LocalDataLog.Marshal(b, m, deterministic)
}
func (m *LocalDataLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LocalDataLog.Merge(m, src)
}
func (m *LocalDataLog) XXX_Size() int {
	return xxx_messageInfo_LocalDataLog.Size(m)
}
func (m *LocalDataLog) XXX_DiscardUnknown() {
	xxx_messageInfo_LocalDataLog.DiscardUnknown(m)
}

var xxx_messageInfo_LocalDataLog proto.InternalMessageInfo

func (m *LocalDataLog) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *LocalDataLog) GetPreValue() []byte {
	if m != nil {
		return m.PreValue
	}
	return nil
}

func (m *LocalDataLog) GetCurValue() []byte {
	if m != nil {
		return m.CurValue
	}
	return nil
}

func init() {
	proto.RegisterType((*WasmAction)(nil), "types.wasmAction")
	proto.RegisterType((*WasmCreate)(nil), "types.wasmCreate")
	proto.RegisterType((*WasmCall)(nil), "types.wasmCall")
	proto.RegisterType((*QueryCheckConract)(nil), "types.queryCheckConract")
	proto.RegisterType((*QueryCreateTransaction)(nil), "types.queryCreateTransaction")
	proto.RegisterType((*CustomLog)(nil), "types.customLog")
	proto.RegisterType((*CreateContractLog)(nil), "types.createContractLog")
	proto.RegisterType((*CallContractLog)(nil), "types.callContractLog")
	proto.RegisterType((*LocalDataLog)(nil), "types.localDataLog")
}

func init() { proto.RegisterFile("wasm.proto", fileDescriptor_7d78909ad64e3bbb) }

var fileDescriptor_7d78909ad64e3bbb = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0x3d, 0x0f, 0xd3, 0x30,
	0x10, 0x6d, 0xe2, 0x26, 0x34, 0x47, 0x45, 0x5b, 0x0f, 0x55, 0xc4, 0x00, 0x91, 0x25, 0x44, 0x24,
	0xa4, 0x0e, 0xc0, 0xc6, 0x04, 0x61, 0xe8, 0xc0, 0x64, 0x21, 0xc4, 0x02, 0x92, 0x71, 0x5d, 0x5a,
	0xd5, 0x89, 0x83, 0xe3, 0x80, 0xf2, 0x57, 0xf8, 0xb5, 0xc8, 0x1f, 0x0d, 0x11, 0xaa, 0x3a, 0xb0,
	0xdd, 0xc7, 0xbb, 0x7b, 0x77, 0xf7, 0x0e, 0xe0, 0x17, 0xeb, 0xea, 0x5d, 0xab, 0x95, 0x51, 0x38,
	0x31, 0x43, 0x2b, 0x3a, 0x32, 0xf8, 0xe0, 0x5b, 0x6e, 0xce, 0xaa, 0xc1, 0x2f, 0x20, 0xe5, 0x5a,
	0x30, 0x23, 0xf2, 0xa8, 0x88, 0xca, 0x87, 0x2f, 0x37, 0x3b, 0x87, 0xda, 0x59, 0x48, 0xe5, 0x12,
	0xfb, 0x19, 0x0d, 0x10, 0xfc, 0x0c, 0xe6, 0x9c, 0x49, 0x99, 0xc7, 0x0e, 0xba, 0x9a, 0x42, 0x99,
	0x94, 0xfb, 0x19, 0x75, 0x69, 0xfc, 0x08, 0x62, 0x33, 0xe4, 0xa8, 0x88, 0xca, 0x84, 0xc6, 0x66,
	0x78, 0xf7, 0x00, 0x92, 0x9f, 0x4c, 0xf6, 0x82, 0xbc, 0xf6, 0xd4, 0xbe, 0x2f, 0xc6, 0x30, 0x6f,
	0x58, 0xed, 0x89, 0x33, 0xea, 0x6c, 0x1b, 0xe3, 0xea, 0x20, 0x1c, 0xc3, 0x92, 0x3a, 0x9b, 0x7c,
	0x85, 0xc5, 0x95, 0x02, 0x3f, 0x86, 0x05, 0x57, 0x8d, 0xd1, 0x8c, 0x9b, 0x50, 0x37, 0xfa, 0x78,
	0x0b, 0x69, 0x2d, 0xcc, 0x49, 0x1d, 0x5c, 0x75, 0x46, 0x83, 0x87, 0x9f, 0x00, 0xb4, 0x4c, 0xb3,
	0x5a, 0x18, 0xa1, 0xbb, 0x1c, 0x15, 0xa8, 0x44, 0x74, 0x12, 0x21, 0xcf, 0x61, 0xf3, 0xa3, 0x17,
	0x7a, 0xa8, 0x4e, 0x82, 0x5f, 0x2a, 0xd5, 0xb8, 0x66, 0x37, 0x86, 0x23, 0xbf, 0x23, 0xd8, 0x7a,
	0xa4, 0x5b, 0xe0, 0xa3, 0x66, 0x4d, 0xc7, 0xfc, 0x19, 0xfd, 0xca, 0xd1, 0x75, 0xe5, 0xb1, 0x3c,
	0xbe, 0xb1, 0x1b, 0xfa, 0xbb, 0xdb, 0x64, 0xe6, 0xf9, 0x9d, 0x99, 0x93, 0x7f, 0x67, 0xc6, 0x6b,
	0x40, 0x47, 0x21, 0xf2, 0xb4, 0x88, 0x4a, 0x44, 0xad, 0x49, 0x9e, 0x42, 0xc6, 0xfb, 0xce, 0xa8,
	0xfa, 0x83, 0xfa, 0x6e, 0xa9, 0xce, 0xcd, 0x51, 0xe5, 0x51, 0x81, 0x2c, 0xbd, 0xb5, 0xc9, 0x1b,
	0xd8, 0x78, 0x19, 0xab, 0x70, 0xb0, 0x00, 0xbc, 0xab, 0x41, 0x16, 0x34, 0xf8, 0x02, 0x2b, 0x2b,
	0xed, 0xb4, 0xf4, 0x7f, 0xa4, 0xd8, 0x42, 0xaa, 0x45, 0xd7, 0x4b, 0xe3, 0x8e, 0x80, 0x68, 0xf0,
	0xc8, 0x67, 0x58, 0x4a, 0xc5, 0x99, 0x7c, 0xcf, 0x0c, 0xb3, 0xbd, 0xd7, 0x80, 0x2e, 0xc2, 0xdf,
	0x73, 0x49, 0xad, 0x69, 0xd9, 0x5a, 0x2d, 0x3e, 0xd9, 0x37, 0x0a, 0xcf, 0x31, 0xfa, 0x6e, 0x92,
	0x5e, 0xfb, 0x9c, 0x3f, 0xee, 0xe8, 0x7f, 0x4b, 0xdd, 0xef, 0xbf, 0xfa, 0x13, 0x00, 0x00, 0xff,
	0xff, 0xb4, 0xb4, 0xce, 0xf0, 0x09, 0x03, 0x00, 0x00,
}
