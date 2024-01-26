// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/charge/charge.proto

package go_micro_srv_charge

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type StateType int32

const (
	StateType_STATE_NORMAL      StateType = 0
	StateType_STATE_PAY_SUCCESS StateType = 1
	StateType_STATE_PAYING      StateType = 2
	StateType_STATE_REFUND      StateType = 3
)

var StateType_name = map[int32]string{
	0: "STATE_NORMAL",
	1: "STATE_PAY_SUCCESS",
	2: "STATE_PAYING",
	3: "STATE_REFUND",
}

var StateType_value = map[string]int32{
	"STATE_NORMAL":      0,
	"STATE_PAY_SUCCESS": 1,
	"STATE_PAYING":      2,
	"STATE_REFUND":      3,
}

func (x StateType) String() string {
	return proto.EnumName(StateType_name, int32(x))
}

func (StateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_7b2d19e5cbf1b47b, []int{0}
}

type ChargeReq struct {
	Uid                  int64     `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Amount               int64     `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Channel              string    `protobuf:"bytes,3,opt,name=channel,proto3" json:"channel,omitempty"`
	Subject              string    `protobuf:"bytes,4,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectId            int64     `protobuf:"varint,8,opt,name=subjectId,proto3" json:"subjectId,omitempty"`
	State                StateType `protobuf:"varint,5,opt,name=state,proto3,enum=go.micro.srv.charge.StateType" json:"state,omitempty"`
	ThirdOrderNo         string    `protobuf:"bytes,6,opt,name=third_order_no,json=thirdOrderNo,proto3" json:"third_order_no,omitempty"`
	OrderId              string    `protobuf:"bytes,7,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ChargeReq) Reset()         { *m = ChargeReq{} }
func (m *ChargeReq) String() string { return proto.CompactTextString(m) }
func (*ChargeReq) ProtoMessage()    {}
func (*ChargeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2d19e5cbf1b47b, []int{0}
}
func (m *ChargeReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChargeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChargeReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChargeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeReq.Merge(m, src)
}
func (m *ChargeReq) XXX_Size() int {
	return m.Size()
}
func (m *ChargeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeReq proto.InternalMessageInfo

func (m *ChargeReq) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *ChargeReq) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *ChargeReq) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *ChargeReq) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ChargeReq) GetSubjectId() int64 {
	if m != nil {
		return m.SubjectId
	}
	return 0
}

func (m *ChargeReq) GetState() StateType {
	if m != nil {
		return m.State
	}
	return StateType_STATE_NORMAL
}

func (m *ChargeReq) GetThirdOrderNo() string {
	if m != nil {
		return m.ThirdOrderNo
	}
	return ""
}

func (m *ChargeReq) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type ChargeResponse struct {
	State                int32    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderIdInt           int64    `protobuf:"varint,4,opt,name=order_id_int,json=orderIdInt,proto3" json:"order_id_int,omitempty"`
	Status               int32    `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChargeResponse) Reset()         { *m = ChargeResponse{} }
func (m *ChargeResponse) String() string { return proto.CompactTextString(m) }
func (*ChargeResponse) ProtoMessage()    {}
func (*ChargeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2d19e5cbf1b47b, []int{1}
}
func (m *ChargeResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChargeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ChargeResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ChargeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChargeResponse.Merge(m, src)
}
func (m *ChargeResponse) XXX_Size() int {
	return m.Size()
}
func (m *ChargeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChargeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChargeResponse proto.InternalMessageInfo

func (m *ChargeResponse) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *ChargeResponse) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *ChargeResponse) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *ChargeResponse) GetOrderIdInt() int64 {
	if m != nil {
		return m.OrderIdInt
	}
	return 0
}

func (m *ChargeResponse) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type QueryReq struct {
	OrderId              string   `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryReq) Reset()         { *m = QueryReq{} }
func (m *QueryReq) String() string { return proto.CompactTextString(m) }
func (*QueryReq) ProtoMessage()    {}
func (*QueryReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2d19e5cbf1b47b, []int{2}
}
func (m *QueryReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryReq.Merge(m, src)
}
func (m *QueryReq) XXX_Size() int {
	return m.Size()
}
func (m *QueryReq) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryReq.DiscardUnknown(m)
}

var xxx_messageInfo_QueryReq proto.InternalMessageInfo

func (m *QueryReq) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

type QueryRsp struct {
	State                int32    `protobuf:"varint,1,opt,name=state,proto3" json:"state,omitempty"`
	OrderId              string   `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Status               int32    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRsp) Reset()         { *m = QueryRsp{} }
func (m *QueryRsp) String() string { return proto.CompactTextString(m) }
func (*QueryRsp) ProtoMessage()    {}
func (*QueryRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_7b2d19e5cbf1b47b, []int{3}
}
func (m *QueryRsp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRsp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRsp.Merge(m, src)
}
func (m *QueryRsp) XXX_Size() int {
	return m.Size()
}
func (m *QueryRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRsp.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRsp proto.InternalMessageInfo

func (m *QueryRsp) GetState() int32 {
	if m != nil {
		return m.State
	}
	return 0
}

func (m *QueryRsp) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

func (m *QueryRsp) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func init() {
	proto.RegisterEnum("go.micro.srv.charge.StateType", StateType_name, StateType_value)
	proto.RegisterType((*ChargeReq)(nil), "go.micro.srv.charge.ChargeReq")
	proto.RegisterType((*ChargeResponse)(nil), "go.micro.srv.charge.ChargeResponse")
	proto.RegisterType((*QueryReq)(nil), "go.micro.srv.charge.QueryReq")
	proto.RegisterType((*QueryRsp)(nil), "go.micro.srv.charge.QueryRsp")
}

func init() { proto.RegisterFile("proto/charge/charge.proto", fileDescriptor_7b2d19e5cbf1b47b) }

var fileDescriptor_7b2d19e5cbf1b47b = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xdf, 0x8a, 0xd3, 0x4e,
	0x14, 0xee, 0x34, 0xbf, 0xa6, 0xed, 0xa1, 0xbf, 0x12, 0xc7, 0x7f, 0xb3, 0x8b, 0x86, 0x12, 0x15,
	0x8a, 0x17, 0x11, 0x56, 0x5f, 0xa0, 0x5b, 0xab, 0x04, 0xd6, 0x76, 0x4d, 0xba, 0xe2, 0x5e, 0x85,
	0x6c, 0x32, 0x6c, 0x23, 0x6e, 0x26, 0xce, 0x24, 0x42, 0xdf, 0xc3, 0x0b, 0x1f, 0xc9, 0x4b, 0x1f,
	0x41, 0xea, 0x7b, 0x88, 0xcc, 0x4c, 0xd2, 0x74, 0xa1, 0x2c, 0x08, 0x7b, 0xd5, 0x39, 0xdf, 0x77,
	0xe6, 0x3b, 0x5f, 0xbf, 0x33, 0x04, 0x0e, 0x72, 0xce, 0x0a, 0xf6, 0x22, 0x5e, 0x45, 0xfc, 0x92,
	0x56, 0x3f, 0xae, 0xc2, 0xf0, 0xdd, 0x4b, 0xe6, 0x5e, 0xa5, 0x31, 0x67, 0xae, 0xe0, 0x5f, 0x5d,
	0x4d, 0x39, 0x7f, 0x10, 0xf4, 0xa7, 0xea, 0xe8, 0xd3, 0x2f, 0xd8, 0x02, 0xa3, 0x4c, 0x13, 0x82,
	0x46, 0x68, 0x6c, 0xf8, 0xf2, 0x88, 0x1f, 0x80, 0x19, 0x5d, 0xb1, 0x32, 0x2b, 0x48, 0x5b, 0x81,
	0x55, 0x85, 0x09, 0x74, 0xe3, 0x55, 0x94, 0x65, 0xf4, 0x33, 0x31, 0x46, 0x68, 0xdc, 0xf7, 0xeb,
	0x52, 0x32, 0xa2, 0xbc, 0xf8, 0x44, 0xe3, 0x82, 0xfc, 0xa7, 0x99, 0xaa, 0xc4, 0x8f, 0xa0, 0x5f,
	0x1d, 0xbd, 0x84, 0xf4, 0x94, 0x5c, 0x03, 0xe0, 0x57, 0xd0, 0x11, 0x45, 0x54, 0x50, 0xd2, 0x19,
	0xa1, 0xf1, 0xf0, 0xc8, 0x76, 0xf7, 0xd8, 0x75, 0x03, 0xd9, 0xb1, 0x5c, 0xe7, 0xd4, 0xd7, 0xcd,
	0xf8, 0x29, 0x0c, 0x8b, 0x55, 0xca, 0x93, 0x90, 0xf1, 0x84, 0xf2, 0x30, 0x63, 0xc4, 0x54, 0x43,
	0x07, 0x0a, 0x5d, 0x48, 0x70, 0xce, 0xf0, 0x01, 0xf4, 0x34, 0x9f, 0x26, 0xa4, 0xab, 0x4d, 0xa9,
	0xda, 0x4b, 0x9c, 0x6f, 0x08, 0x86, 0x75, 0x00, 0x22, 0x67, 0x99, 0xa0, 0xf8, 0x5e, 0xed, 0x44,
	0xe6, 0xd0, 0xa9, 0x27, 0xed, 0x6a, 0xb4, 0xaf, 0x69, 0xe0, 0x87, 0xd0, 0x2d, 0x85, 0x66, 0x0c,
	0x9d, 0x92, 0x2c, 0xbd, 0x04, 0x8f, 0x60, 0x50, 0xdf, 0x09, 0xd3, 0x4c, 0x07, 0x62, 0xf8, 0x50,
	0xdd, 0xf3, 0xb2, 0x42, 0xe6, 0x2b, 0xe5, 0x4b, 0xa1, 0xfe, 0x76, 0xc7, 0xaf, 0x2a, 0xe7, 0x19,
	0xf4, 0xde, 0x97, 0x94, 0xaf, 0xe5, 0x56, 0x76, 0x27, 0xa3, 0xeb, 0xee, 0x83, 0xba, 0x4d, 0xe4,
	0xff, 0x6e, 0xbb, 0x99, 0x6d, 0xec, 0xce, 0x7e, 0xfe, 0x11, 0xfa, 0xdb, 0x9c, 0xb1, 0x05, 0x83,
	0x60, 0x39, 0x59, 0xce, 0xc2, 0xf9, 0xc2, 0x7f, 0x37, 0x39, 0xb1, 0x5a, 0xf8, 0x3e, 0xdc, 0xd1,
	0xc8, 0xe9, 0xe4, 0x3c, 0x0c, 0xce, 0xa6, 0xd3, 0x59, 0x10, 0x58, 0xa8, 0x69, 0x3c, 0x9d, 0x9c,
	0x7b, 0xf3, 0xb7, 0x56, 0xbb, 0x41, 0xfc, 0xd9, 0x9b, 0xb3, 0xf9, 0x6b, 0xcb, 0x38, 0xda, 0xb4,
	0xc1, 0xd4, 0x61, 0xe3, 0x05, 0x98, 0x53, 0x4e, 0xa5, 0xc3, 0xfd, 0x9b, 0xde, 0x3e, 0xca, 0xc3,
	0x27, 0x37, 0xf2, 0x7a, 0x67, 0x4e, 0x0b, 0x7f, 0x80, 0xff, 0x35, 0x16, 0x94, 0x71, 0x4c, 0x85,
	0xb8, 0x2d, 0xdd, 0x13, 0x00, 0x15, 0xb1, 0x7a, 0x4b, 0xf8, 0xf1, 0xde, 0x4b, 0xf5, 0xaa, 0x0e,
	0x6f, 0xa2, 0x45, 0xee, 0xb4, 0x70, 0x04, 0xa4, 0x51, 0x3b, 0x5e, 0x2f, 0xb7, 0xaf, 0xd4, 0x4b,
	0x6e, 0xc9, 0xf0, 0xb1, 0xf5, 0x63, 0x63, 0xa3, 0x9f, 0x1b, 0x1b, 0xfd, 0xda, 0xd8, 0xe8, 0xfb,
	0x6f, 0xbb, 0x75, 0x61, 0xaa, 0x0f, 0xc0, 0xcb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x57, 0x17,
	0x9d, 0x80, 0x1d, 0x04, 0x00, 0x00,
}

func (m *ChargeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChargeReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChargeReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.SubjectId != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.SubjectId))
		i--
		dAtA[i] = 0x40
	}
	if len(m.OrderId) > 0 {
		i -= len(m.OrderId)
		copy(dAtA[i:], m.OrderId)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.OrderId)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ThirdOrderNo) > 0 {
		i -= len(m.ThirdOrderNo)
		copy(dAtA[i:], m.ThirdOrderNo)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.ThirdOrderNo)))
		i--
		dAtA[i] = 0x32
	}
	if m.State != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Channel) > 0 {
		i -= len(m.Channel)
		copy(dAtA[i:], m.Channel)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.Channel)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Amount != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x10
	}
	if m.Uid != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.Uid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ChargeResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChargeResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ChargeResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.OrderIdInt != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.OrderIdInt))
		i--
		dAtA[i] = 0x20
	}
	if m.UserId != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.UserId))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OrderId) > 0 {
		i -= len(m.OrderId)
		copy(dAtA[i:], m.OrderId)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.OrderId)))
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.OrderId) > 0 {
		i -= len(m.OrderId)
		copy(dAtA[i:], m.OrderId)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.OrderId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRsp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRsp) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRsp) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Status != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.OrderId) > 0 {
		i -= len(m.OrderId)
		copy(dAtA[i:], m.OrderId)
		i = encodeVarintCharge(dAtA, i, uint64(len(m.OrderId)))
		i--
		dAtA[i] = 0x12
	}
	if m.State != 0 {
		i = encodeVarintCharge(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintCharge(dAtA []byte, offset int, v uint64) int {
	offset -= sovCharge(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ChargeReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Uid != 0 {
		n += 1 + sovCharge(uint64(m.Uid))
	}
	if m.Amount != 0 {
		n += 1 + sovCharge(uint64(m.Amount))
	}
	l = len(m.Channel)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	if m.State != 0 {
		n += 1 + sovCharge(uint64(m.State))
	}
	l = len(m.ThirdOrderNo)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	if m.SubjectId != 0 {
		n += 1 + sovCharge(uint64(m.SubjectId))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ChargeResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovCharge(uint64(m.State))
	}
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	if m.UserId != 0 {
		n += 1 + sovCharge(uint64(m.UserId))
	}
	if m.OrderIdInt != 0 {
		n += 1 + sovCharge(uint64(m.OrderIdInt))
	}
	if m.Status != 0 {
		n += 1 + sovCharge(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueryReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *QueryRsp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.State != 0 {
		n += 1 + sovCharge(uint64(m.State))
	}
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovCharge(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovCharge(uint64(m.Status))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCharge(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCharge(x uint64) (n int) {
	return sovCharge(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChargeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCharge
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
			return fmt.Errorf("proto: ChargeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChargeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			m.Uid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Uid |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Channel", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Channel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= StateType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ThirdOrderNo", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ThirdOrderNo = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectId", wireType)
			}
			m.SubjectId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubjectId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCharge
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
func (m *ChargeResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCharge
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
			return fmt.Errorf("proto: ChargeResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChargeResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserId", wireType)
			}
			m.UserId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.UserId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderIdInt", wireType)
			}
			m.OrderIdInt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OrderIdInt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCharge
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
func (m *QueryReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCharge
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
			return fmt.Errorf("proto: QueryReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCharge
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
func (m *QueryRsp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCharge
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
			return fmt.Errorf("proto: QueryRsp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRsp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
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
				return ErrInvalidLengthCharge
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCharge
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCharge
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCharge(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCharge
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
func skipCharge(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCharge
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
					return 0, ErrIntOverflowCharge
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
					return 0, ErrIntOverflowCharge
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
				return 0, ErrInvalidLengthCharge
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCharge
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCharge
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCharge        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCharge          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCharge = fmt.Errorf("proto: unexpected end of group")
)
