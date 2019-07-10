// Code generated by protoc-gen-go. DO NOT EDIT.
// source: project.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AutonomyProposalProject struct {
	PropProject *ProposalProject `protobuf:"bytes,1,opt,name=propProject" json:"propProject,omitempty"`
	// 董事会投票结果
	BoardResult *VotesResult `protobuf:"bytes,2,opt,name=boardResult" json:"boardResult,omitempty"`
	// 是否需要公示
	Publicity bool `protobuf:"varint,3,opt,name=publicity" json:"publicity,omitempty"`
	// 全体持票人反对票
	OpposeVotes int32 `protobuf:"varint,4,opt,name=opposeVotes" json:"opposeVotes,omitempty"`
	// 是否通过
	PubPass bool `protobuf:"varint,5,opt,name=pubPass" json:"pubPass,omitempty"`
	// 状态
	Status int32 `protobuf:"varint,6,opt,name=status" json:"status,omitempty"`
}

func (m *AutonomyProposalProject) Reset()                    { *m = AutonomyProposalProject{} }
func (m *AutonomyProposalProject) String() string            { return proto.CompactTextString(m) }
func (*AutonomyProposalProject) ProtoMessage()               {}
func (*AutonomyProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *AutonomyProposalProject) GetPropProject() *ProposalProject {
	if m != nil {
		return m.PropProject
	}
	return nil
}

func (m *AutonomyProposalProject) GetBoardResult() *VotesResult {
	if m != nil {
		return m.BoardResult
	}
	return nil
}

func (m *AutonomyProposalProject) GetPublicity() bool {
	if m != nil {
		return m.Publicity
	}
	return false
}

func (m *AutonomyProposalProject) GetOpposeVotes() int32 {
	if m != nil {
		return m.OpposeVotes
	}
	return 0
}

func (m *AutonomyProposalProject) GetPubPass() bool {
	if m != nil {
		return m.PubPass
	}
	return false
}

func (m *AutonomyProposalProject) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ProposalProject struct {
	// 提案时间
	Year  int32 `protobuf:"varint,1,opt,name=year" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day" json:"day,omitempty"`
	// 项目相关
	FirstStage   string `protobuf:"bytes,4,opt,name=firstStage" json:"firstStage,omitempty"`
	LastStage    string `protobuf:"bytes,5,opt,name=lastStage" json:"lastStage,omitempty"`
	Production   string `protobuf:"bytes,6,opt,name=production" json:"production,omitempty"`
	Description  string `protobuf:"bytes,7,opt,name=description" json:"description,omitempty"`
	Contractor   string `protobuf:"bytes,8,opt,name=contractor" json:"contractor,omitempty"`
	Amount       int32  `protobuf:"varint,9,opt,name=amount" json:"amount,omitempty"`
	AmountDetail string `protobuf:"bytes,10,opt,name=amountDetail" json:"amountDetail,omitempty"`
	// 支付相关
	ToAddr string `protobuf:"bytes,11,opt,name=toAddr" json:"toAddr,omitempty"`
	// 投票相关
	StartBlockHeight    int64 `protobuf:"varint,12,opt,name=startBlockHeight" json:"startBlockHeight,omitempty"`
	EndBlockHeight      int64 `protobuf:"varint,13,opt,name=endBlockHeight" json:"endBlockHeight,omitempty"`
	RealEndBlockHeight  int64 `protobuf:"varint,14,opt,name=realEndBlockHeight" json:"realEndBlockHeight,omitempty"`
	ProjectNeedBlockNum int32 `protobuf:"varint,15,opt,name=projectNeedBlockNum" json:"projectNeedBlockNum,omitempty"`
}

func (m *ProposalProject) Reset()                    { *m = ProposalProject{} }
func (m *ProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ProposalProject) ProtoMessage()               {}
func (*ProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{1} }

func (m *ProposalProject) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *ProposalProject) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *ProposalProject) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *ProposalProject) GetFirstStage() string {
	if m != nil {
		return m.FirstStage
	}
	return ""
}

func (m *ProposalProject) GetLastStage() string {
	if m != nil {
		return m.LastStage
	}
	return ""
}

func (m *ProposalProject) GetProduction() string {
	if m != nil {
		return m.Production
	}
	return ""
}

func (m *ProposalProject) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ProposalProject) GetContractor() string {
	if m != nil {
		return m.Contractor
	}
	return ""
}

func (m *ProposalProject) GetAmount() int32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ProposalProject) GetAmountDetail() string {
	if m != nil {
		return m.AmountDetail
	}
	return ""
}

func (m *ProposalProject) GetToAddr() string {
	if m != nil {
		return m.ToAddr
	}
	return ""
}

func (m *ProposalProject) GetStartBlockHeight() int64 {
	if m != nil {
		return m.StartBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetEndBlockHeight() int64 {
	if m != nil {
		return m.EndBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetRealEndBlockHeight() int64 {
	if m != nil {
		return m.RealEndBlockHeight
	}
	return 0
}

func (m *ProposalProject) GetProjectNeedBlockNum() int32 {
	if m != nil {
		return m.ProjectNeedBlockNum
	}
	return 0
}

type RevokeProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *RevokeProposalProject) Reset()                    { *m = RevokeProposalProject{} }
func (m *RevokeProposalProject) String() string            { return proto.CompactTextString(m) }
func (*RevokeProposalProject) ProtoMessage()               {}
func (*RevokeProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

func (m *RevokeProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type TerminateProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *TerminateProposalProject) Reset()                    { *m = TerminateProposalProject{} }
func (m *TerminateProposalProject) String() string            { return proto.CompactTextString(m) }
func (*TerminateProposalProject) ProtoMessage()               {}
func (*TerminateProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

func (m *TerminateProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

// receipt
type ReceiptProposalProject struct {
	Prev    *AutonomyProposalProject `protobuf:"bytes,1,opt,name=prev" json:"prev,omitempty"`
	Current *AutonomyProposalProject `protobuf:"bytes,2,opt,name=current" json:"current,omitempty"`
}

func (m *ReceiptProposalProject) Reset()                    { *m = ReceiptProposalProject{} }
func (m *ReceiptProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReceiptProposalProject) ProtoMessage()               {}
func (*ReceiptProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

func (m *ReceiptProposalProject) GetPrev() *AutonomyProposalProject {
	if m != nil {
		return m.Prev
	}
	return nil
}

func (m *ReceiptProposalProject) GetCurrent() *AutonomyProposalProject {
	if m != nil {
		return m.Current
	}
	return nil
}

type LocalProposalProject struct {
	PropPrj  *AutonomyProposalProject `protobuf:"bytes,1,opt,name=propPrj" json:"propPrj,omitempty"`
	Comments []string                 `protobuf:"bytes,2,rep,name=comments" json:"comments,omitempty"`
}

func (m *LocalProposalProject) Reset()                    { *m = LocalProposalProject{} }
func (m *LocalProposalProject) String() string            { return proto.CompactTextString(m) }
func (*LocalProposalProject) ProtoMessage()               {}
func (*LocalProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

func (m *LocalProposalProject) GetPropPrj() *AutonomyProposalProject {
	if m != nil {
		return m.PropPrj
	}
	return nil
}

func (m *LocalProposalProject) GetComments() []string {
	if m != nil {
		return m.Comments
	}
	return nil
}

// query
type ReplyQueryProposalProject struct {
	ProposalID string `protobuf:"bytes,1,opt,name=proposalID" json:"proposalID,omitempty"`
}

func (m *ReplyQueryProposalProject) Reset()                    { *m = ReplyQueryProposalProject{} }
func (m *ReplyQueryProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReplyQueryProposalProject) ProtoMessage()               {}
func (*ReplyQueryProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ReplyQueryProposalProject) GetProposalID() string {
	if m != nil {
		return m.ProposalID
	}
	return ""
}

type ReplyProposalProject struct {
	PropProjects []*LocalProposalProject `protobuf:"bytes,1,rep,name=propProjects" json:"propProjects,omitempty"`
}

func (m *ReplyProposalProject) Reset()                    { *m = ReplyProposalProject{} }
func (m *ReplyProposalProject) String() string            { return proto.CompactTextString(m) }
func (*ReplyProposalProject) ProtoMessage()               {}
func (*ReplyProposalProject) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{7} }

func (m *ReplyProposalProject) GetPropProjects() []*LocalProposalProject {
	if m != nil {
		return m.PropProjects
	}
	return nil
}

func init() {
	proto.RegisterType((*AutonomyProposalProject)(nil), "types.AutonomyProposalProject")
	proto.RegisterType((*ProposalProject)(nil), "types.ProposalProject")
	proto.RegisterType((*RevokeProposalProject)(nil), "types.RevokeProposalProject")
	proto.RegisterType((*TerminateProposalProject)(nil), "types.TerminateProposalProject")
	proto.RegisterType((*ReceiptProposalProject)(nil), "types.ReceiptProposalProject")
	proto.RegisterType((*LocalProposalProject)(nil), "types.LocalProposalProject")
	proto.RegisterType((*ReplyQueryProposalProject)(nil), "types.ReplyQueryProposalProject")
	proto.RegisterType((*ReplyProposalProject)(nil), "types.ReplyProposalProject")
}

func init() { proto.RegisterFile("project.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 579 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0x4f, 0x6f, 0xd4, 0x3c,
	0x10, 0xc6, 0x95, 0x6e, 0xd3, 0xed, 0xce, 0xf6, 0x9f, 0xfc, 0xf6, 0x2d, 0xa6, 0xa0, 0x2a, 0xca,
	0x01, 0x45, 0x1c, 0x56, 0xa8, 0x20, 0x51, 0xc1, 0x01, 0x15, 0x15, 0x09, 0x24, 0x54, 0x15, 0x83,
	0xe0, 0xec, 0x4d, 0x4c, 0x9b, 0x36, 0x89, 0x2d, 0x7b, 0x52, 0x29, 0x5f, 0x80, 0xaf, 0xcc, 0x81,
	0x0b, 0xca, 0x38, 0x4b, 0xd3, 0xb0, 0x08, 0x7a, 0xf3, 0x3c, 0xf3, 0xfc, 0xac, 0x19, 0xdb, 0x63,
	0xd8, 0x34, 0x56, 0x5f, 0xaa, 0x14, 0x67, 0xc6, 0x6a, 0xd4, 0x2c, 0xc4, 0xc6, 0x28, 0xb7, 0xbf,
	0x59, 0xa4, 0xba, 0x2c, 0x75, 0xe5, 0xd5, 0xf8, 0x47, 0x00, 0xf7, 0x8e, 0x6b, 0xd4, 0x95, 0x2e,
	0x9b, 0x33, 0xab, 0x8d, 0x76, 0xb2, 0x38, 0xf3, 0x1c, 0x3b, 0x82, 0xa9, 0xb1, 0xda, 0x74, 0x21,
	0x0f, 0xa2, 0x20, 0x99, 0x1e, 0xee, 0xcd, 0x68, 0x9f, 0xd9, 0xc0, 0x2c, 0xfa, 0x56, 0xf6, 0x0c,
	0xa6, 0x73, 0x2d, 0x6d, 0x26, 0x94, 0xab, 0x0b, 0xe4, 0x2b, 0x44, 0xb2, 0x8e, 0xfc, 0xac, 0x51,
	0x39, 0x9f, 0x11, 0x7d, 0x1b, 0x7b, 0x08, 0x13, 0x53, 0xcf, 0x8b, 0x3c, 0xcd, 0xb1, 0xe1, 0xa3,
	0x28, 0x48, 0xd6, 0xc5, 0x8d, 0xc0, 0x22, 0x98, 0x6a, 0x63, 0xb4, 0x53, 0xc4, 0xf3, 0xd5, 0x28,
	0x48, 0x42, 0xd1, 0x97, 0x18, 0x87, 0xb1, 0xa9, 0xe7, 0x67, 0xd2, 0x39, 0x1e, 0x12, 0xbd, 0x08,
	0xd9, 0x1e, 0xac, 0x39, 0x94, 0x58, 0x3b, 0xbe, 0x46, 0x58, 0x17, 0xc5, 0xdf, 0x47, 0xb0, 0x3d,
	0xec, 0x9a, 0xc1, 0x6a, 0xa3, 0xa4, 0xa5, 0x76, 0x43, 0x41, 0x6b, 0xb6, 0x0b, 0x61, 0xa9, 0x2b,
	0xbc, 0xa0, 0x4e, 0x42, 0xe1, 0x03, 0xb6, 0x03, 0xa3, 0x4c, 0xfa, 0x4a, 0x43, 0xd1, 0x2e, 0xd9,
	0x01, 0xc0, 0xd7, 0xdc, 0x3a, 0xfc, 0x88, 0xf2, 0x5c, 0x51, 0x89, 0x13, 0xd1, 0x53, 0xda, 0x0e,
	0x0b, 0xb9, 0x48, 0x87, 0x94, 0xbe, 0x11, 0x5a, 0xda, 0x58, 0x9d, 0xd5, 0x29, 0xe6, 0xba, 0xa2,
	0x4a, 0x27, 0xa2, 0xa7, 0xb4, 0x27, 0x90, 0x29, 0x97, 0xda, 0xdc, 0x90, 0x61, 0x4c, 0x86, 0xbe,
	0xd4, 0xee, 0x90, 0xea, 0x0a, 0xad, 0x4c, 0x51, 0x5b, 0xbe, 0xee, 0x77, 0xb8, 0x51, 0xda, 0x73,
	0x90, 0xa5, 0xae, 0x2b, 0xe4, 0x13, 0x7f, 0x0e, 0x3e, 0x62, 0x31, 0x6c, 0xf8, 0xd5, 0x89, 0x42,
	0x99, 0x17, 0x1c, 0x88, 0xbc, 0xa5, 0xb5, 0x2c, 0xea, 0xe3, 0x2c, 0xb3, 0x7c, 0x4a, 0xd9, 0x2e,
	0x62, 0x8f, 0x61, 0xc7, 0xa1, 0xb4, 0xf8, 0xba, 0xd0, 0xe9, 0xd5, 0x5b, 0x95, 0x9f, 0x5f, 0x20,
	0xdf, 0x88, 0x82, 0x64, 0x24, 0x7e, 0xd3, 0xd9, 0x23, 0xd8, 0x52, 0x55, 0xd6, 0x77, 0x6e, 0x92,
	0x73, 0xa0, 0xb2, 0x19, 0x30, 0xab, 0x64, 0xf1, 0xe6, 0xb6, 0x77, 0x8b, 0xbc, 0x4b, 0x32, 0xec,
	0x09, 0xfc, 0xd7, 0x3d, 0xf6, 0x53, 0xa5, 0x7c, 0xe6, 0xb4, 0x2e, 0xf9, 0x36, 0x35, 0xb9, 0x2c,
	0x15, 0x3f, 0x87, 0xff, 0x85, 0xba, 0xd6, 0x57, 0x6a, 0x78, 0xfd, 0xfe, 0x12, 0x48, 0x7a, 0x77,
	0x42, 0x8f, 0xc0, 0x5f, 0x42, 0xa7, 0xc4, 0x2f, 0x80, 0x7f, 0x52, 0xb6, 0xcc, 0x2b, 0x89, 0x77,
	0x66, 0xbf, 0x05, 0xb0, 0x27, 0x54, 0xaa, 0x72, 0x83, 0x43, 0xf4, 0x10, 0x56, 0x8d, 0x55, 0xd7,
	0xdd, 0x90, 0x1d, 0x74, 0xa3, 0xf2, 0x87, 0xc9, 0x14, 0xe4, 0x65, 0x47, 0x30, 0x4e, 0x6b, 0x6b,
	0x55, 0xb5, 0x98, 0xb0, 0xbf, 0x61, 0x0b, 0x7b, 0x5c, 0xc0, 0xee, 0x7b, 0x9d, 0x52, 0x62, 0x30,
	0xf1, 0x63, 0x3f, 0xc6, 0x97, 0xff, 0x58, 0xc8, 0xc2, 0xce, 0xf6, 0x61, 0xbd, 0xfd, 0x57, 0x54,
	0x85, 0x8e, 0xaf, 0x44, 0xa3, 0x64, 0x22, 0x7e, 0xc5, 0xf1, 0x4b, 0xb8, 0x2f, 0x94, 0x29, 0x9a,
	0x0f, 0xb5, 0xb2, 0xcd, 0x5d, 0xcf, 0xec, 0x0b, 0xec, 0x12, 0x3c, 0xe4, 0x5e, 0xc1, 0x46, 0xef,
	0xc7, 0x71, 0x3c, 0x88, 0x46, 0xc9, 0xf4, 0xf0, 0x41, 0x57, 0xef, 0xb2, 0xee, 0xc4, 0x2d, 0x60,
	0xbe, 0x46, 0x1f, 0xe0, 0xd3, 0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xe1, 0xe0, 0x86, 0xbb, 0x27,
	0x05, 0x00, 0x00,
}