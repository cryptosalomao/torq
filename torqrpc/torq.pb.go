// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: torq.proto

package torqrpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ChannelFlowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// From what date/time (unix timestamp)
	FromTime int64 `protobuf:"varint,1,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	// To what date/time (unix timestamp)
	ToTime int64 `protobuf:"varint,2,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	// TODO: Add "repeated" here to request multiple channels at once.
	ChanIds []uint64 `protobuf:"varint,3,rep,packed,name=chan_ids,json=chanIds,proto3" json:"chan_ids,omitempty"`
}

func (x *ChannelFlowRequest) Reset() {
	*x = ChannelFlowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelFlowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelFlowRequest) ProtoMessage() {}

func (x *ChannelFlowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelFlowRequest.ProtoReflect.Descriptor instead.
func (*ChannelFlowRequest) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{0}
}

func (x *ChannelFlowRequest) GetFromTime() int64 {
	if x != nil {
		return x.FromTime
	}
	return 0
}

func (x *ChannelFlowRequest) GetToTime() int64 {
	if x != nil {
		return x.ToTime
	}
	return 0
}

func (x *ChannelFlowRequest) GetChanIds() []uint64 {
	if x != nil {
		return x.ChanIds
	}
	return nil
}

type ChannelFlow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// What channel ID's the flow is for
	ChanIds []uint64 `protobuf:"varint,1,rep,packed,name=chan_ids,json=chanIds,proto3" json:"chan_ids,omitempty"`
	// From what date/time (unix timestamp)
	FromTime int64 `protobuf:"varint,2,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	// To what date/time (unix timestamp)
	ToTime int64 `protobuf:"varint,3,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	// Fees earned by other channels using this channels inbound liquidity.
	FeeIn uint64 `protobuf:"varint,4,opt,name=fee_in,json=feeIn,proto3" json:"fee_in,omitempty"`
	// Fees earned by this channels outbound liquidity
	FeeOut uint64 `protobuf:"varint,5,opt,name=fee_out,json=feeOut,proto3" json:"fee_out,omitempty"`
	// Amount inbound
	AmtIn uint64 `protobuf:"varint,6,opt,name=amt_in,json=amtIn,proto3" json:"amt_in,omitempty"`
	// Amount outbound
	AmtOut uint64 `protobuf:"varint,7,opt,name=amt_out,json=amtOut,proto3" json:"amt_out,omitempty"`
	// Number of forwards inbound
	CountIn int64 `protobuf:"varint,8,opt,name=count_in,json=countIn,proto3" json:"count_in,omitempty"`
	// Number of forwards outbound
	CountOut int64 `protobuf:"varint,9,opt,name=count_out,json=countOut,proto3" json:"count_out,omitempty"`
}

func (x *ChannelFlow) Reset() {
	*x = ChannelFlow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChannelFlow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChannelFlow) ProtoMessage() {}

func (x *ChannelFlow) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChannelFlow.ProtoReflect.Descriptor instead.
func (*ChannelFlow) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{1}
}

func (x *ChannelFlow) GetChanIds() []uint64 {
	if x != nil {
		return x.ChanIds
	}
	return nil
}

func (x *ChannelFlow) GetFromTime() int64 {
	if x != nil {
		return x.FromTime
	}
	return 0
}

func (x *ChannelFlow) GetToTime() int64 {
	if x != nil {
		return x.ToTime
	}
	return 0
}

func (x *ChannelFlow) GetFeeIn() uint64 {
	if x != nil {
		return x.FeeIn
	}
	return 0
}

func (x *ChannelFlow) GetFeeOut() uint64 {
	if x != nil {
		return x.FeeOut
	}
	return 0
}

func (x *ChannelFlow) GetAmtIn() uint64 {
	if x != nil {
		return x.AmtIn
	}
	return 0
}

func (x *ChannelFlow) GetAmtOut() uint64 {
	if x != nil {
		return x.AmtOut
	}
	return 0
}

func (x *ChannelFlow) GetCountIn() int64 {
	if x != nil {
		return x.CountIn
	}
	return 0
}

func (x *ChannelFlow) GetCountOut() int64 {
	if x != nil {
		return x.CountOut
	}
	return 0
}

type AggregatedForwards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromTime int64 `protobuf:"varint,1,opt,name=from_time,json=fromTime,proto3" json:"from_time,omitempty"`
	ToTime   int64 `protobuf:"varint,2,opt,name=to_time,json=toTime,proto3" json:"to_time,omitempty"`
	// The incoming channel ID that carried the HTLC that created the circuit.
	ChanIdIn uint64 `protobuf:"varint,3,opt,name=chan_id_in,json=chanIdIn,proto3" json:"chan_id_in,omitempty"`
	// The outgoing channel ID that carried the preimage that completed the
	// circuit.
	ChanIdOut uint64 `protobuf:"varint,4,opt,name=chan_id_out,json=chanIdOut,proto3" json:"chan_id_out,omitempty"`
	// The total fee (in satoshis) that this payment circuit carried.
	Fee uint64 `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	// The total amount (in satoshis) of the incoming HTLC that created half
	// the circuit.
	AmtIn uint64 `protobuf:"varint,6,opt,name=amt_in,json=amtIn,proto3" json:"amt_in,omitempty"`
	// The total amount (in satoshis) of the outgoing HTLC that created the
	// second half of the circuit.
	AmtOut   uint64 `protobuf:"varint,7,opt,name=amt_out,json=amtOut,proto3" json:"amt_out,omitempty"`
	CountIn  int64  `protobuf:"varint,8,opt,name=count_in,json=countIn,proto3" json:"count_in,omitempty"`
	CountOut int64  `protobuf:"varint,9,opt,name=count_out,json=countOut,proto3" json:"count_out,omitempty"`
}

func (x *AggregatedForwards) Reset() {
	*x = AggregatedForwards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AggregatedForwards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AggregatedForwards) ProtoMessage() {}

func (x *AggregatedForwards) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AggregatedForwards.ProtoReflect.Descriptor instead.
func (*AggregatedForwards) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{2}
}

func (x *AggregatedForwards) GetFromTime() int64 {
	if x != nil {
		return x.FromTime
	}
	return 0
}

func (x *AggregatedForwards) GetToTime() int64 {
	if x != nil {
		return x.ToTime
	}
	return 0
}

func (x *AggregatedForwards) GetChanIdIn() uint64 {
	if x != nil {
		return x.ChanIdIn
	}
	return 0
}

func (x *AggregatedForwards) GetChanIdOut() uint64 {
	if x != nil {
		return x.ChanIdOut
	}
	return 0
}

func (x *AggregatedForwards) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *AggregatedForwards) GetAmtIn() uint64 {
	if x != nil {
		return x.AmtIn
	}
	return 0
}

func (x *AggregatedForwards) GetAmtOut() uint64 {
	if x != nil {
		return x.AmtOut
	}
	return 0
}

func (x *AggregatedForwards) GetCountIn() int64 {
	if x != nil {
		return x.CountIn
	}
	return 0
}

func (x *AggregatedForwards) GetCountOut() int64 {
	if x != nil {
		return x.CountOut
	}
	return 0
}

type Forwards struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Forwards []*Forward `protobuf:"bytes,1,rep,name=forwards,proto3" json:"forwards,omitempty"`
}

func (x *Forwards) Reset() {
	*x = Forwards{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forwards) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forwards) ProtoMessage() {}

func (x *Forwards) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forwards.ProtoReflect.Descriptor instead.
func (*Forwards) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{3}
}

func (x *Forwards) GetForwards() []*Forward {
	if x != nil {
		return x.Forwards
	}
	return nil
}

type Forward struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The microseconds' version of TimestampNs, used by TimescaleDB
	Time int64 `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	// The number of nanoseconds elapsed since January 1, 1970 UTC when this
	// circuit was completed.
	TimeNs int64 `protobuf:"varint,2,opt,name=time_ns,json=timeNs,proto3" json:"time_ns,omitempty"`
	// The incoming channel ID that carried the HTLC that created the circuit.
	ChanIdIn uint64 `protobuf:"varint,3,opt,name=chan_id_in,json=chanIdIn,proto3" json:"chan_id_in,omitempty"`
	// The outgoing channel ID that carried the preimage that completed the
	// circuit.
	ChanIdOut uint64 `protobuf:"varint,4,opt,name=chan_id_out,json=chanIdOut,proto3" json:"chan_id_out,omitempty"`
	// The total fee (in satoshis) that this payment circuit carried.
	Fee uint64 `protobuf:"varint,5,opt,name=fee,proto3" json:"fee,omitempty"`
	// The total amount (in satoshis) of the incoming HTLC that created half
	// the circuit.
	AmtIn uint64 `protobuf:"varint,6,opt,name=amt_in,json=amtIn,proto3" json:"amt_in,omitempty"`
	// The total amount (in satoshis) of the outgoing HTLC that created the
	// second half of the circuit.
	AmtOut uint64 `protobuf:"varint,7,opt,name=amt_out,json=amtOut,proto3" json:"amt_out,omitempty"`
}

func (x *Forward) Reset() {
	*x = Forward{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Forward) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Forward) ProtoMessage() {}

func (x *Forward) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Forward.ProtoReflect.Descriptor instead.
func (*Forward) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{4}
}

func (x *Forward) GetTime() int64 {
	if x != nil {
		return x.Time
	}
	return 0
}

func (x *Forward) GetTimeNs() int64 {
	if x != nil {
		return x.TimeNs
	}
	return 0
}

func (x *Forward) GetChanIdIn() uint64 {
	if x != nil {
		return x.ChanIdIn
	}
	return 0
}

func (x *Forward) GetChanIdOut() uint64 {
	if x != nil {
		return x.ChanIdOut
	}
	return 0
}

func (x *Forward) GetFee() uint64 {
	if x != nil {
		return x.Fee
	}
	return 0
}

func (x *Forward) GetAmtIn() uint64 {
	if x != nil {
		return x.AmtIn
	}
	return 0
}

func (x *Forward) GetAmtOut() uint64 {
	if x != nil {
		return x.AmtOut
	}
	return 0
}

type ForwardsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ForwardsRequest) Reset() {
	*x = ForwardsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_torq_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ForwardsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ForwardsRequest) ProtoMessage() {}

func (x *ForwardsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_torq_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ForwardsRequest.ProtoReflect.Descriptor instead.
func (*ForwardsRequest) Descriptor() ([]byte, []int) {
	return file_torq_proto_rawDescGZIP(), []int{5}
}

var File_torq_proto protoreflect.FileDescriptor

var file_torq_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x6f, 0x72, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x6f,
	0x72, 0x71, 0x72, 0x70, 0x63, 0x22, 0x69, 0x0a, 0x12, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x1d, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x49, 0x64, 0x73,
	0x22, 0xfa, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x6c, 0x6f, 0x77,
	0x12, 0x1d, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x49, 0x64, 0x73, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74,
	0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x66, 0x65, 0x65, 0x5f, 0x69, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x66, 0x65, 0x65, 0x49, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x66, 0x65, 0x65, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x66,
	0x65, 0x65, 0x4f, 0x75, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d, 0x74, 0x5f, 0x69, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x6d, 0x74, 0x49, 0x6e, 0x12, 0x17, 0x0a, 0x07,
	0x61, 0x6d, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61,
	0x6d, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x8a, 0x02,
	0x0a, 0x12, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x66, 0x72, 0x6f, 0x6d, 0x54, 0x69, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x74, 0x6f, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02,
	0x30, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x49, 0x64, 0x49, 0x6e, 0x12, 0x22, 0x0a, 0x0b,
	0x63, 0x68, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x49, 0x64, 0x4f, 0x75, 0x74,
	0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66,
	0x65, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x05, 0x61, 0x6d, 0x74, 0x49, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x6d, 0x74,
	0x5f, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x61, 0x6d, 0x74, 0x4f,
	0x75, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x38, 0x0a, 0x08, 0x46, 0x6f,
	0x72, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x74, 0x6f, 0x72, 0x71, 0x72,
	0x70, 0x63, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x52, 0x08, 0x66, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x22, 0xc2, 0x01, 0x0a, 0x07, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04,
	0x74, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x01, 0x52, 0x06, 0x74, 0x69, 0x6d, 0x65, 0x4e,
	0x73, 0x12, 0x20, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x69, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x08, 0x63, 0x68, 0x61, 0x6e, 0x49,
	0x64, 0x49, 0x6e, 0x12, 0x22, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x6e, 0x5f, 0x69, 0x64, 0x5f, 0x6f,
	0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x42, 0x02, 0x30, 0x01, 0x52, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x49, 0x64, 0x4f, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x65, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x66, 0x65, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x6d, 0x74,
	0x5f, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x61, 0x6d, 0x74, 0x49, 0x6e,
	0x12, 0x17, 0x0a, 0x07, 0x61, 0x6d, 0x74, 0x5f, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x61, 0x6d, 0x74, 0x4f, 0x75, 0x74, 0x22, 0x11, 0x0a, 0x0f, 0x46, 0x6f, 0x72,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x8a, 0x01, 0x0a,
	0x07, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70, 0x63, 0x12, 0x3a, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46,
	0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70,
	0x63, 0x2e, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x11, 0x2e, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x73, 0x12, 0x43, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x2e, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70, 0x63,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x46, 0x6c, 0x6f, 0x77, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6e, 0x63, 0x61, 0x70, 0x69, 0x74, 0x61,
	0x6c, 0x2f, 0x74, 0x6f, 0x72, 0x71, 0x2f, 0x74, 0x6f, 0x72, 0x71, 0x72, 0x70, 0x63, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_torq_proto_rawDescOnce sync.Once
	file_torq_proto_rawDescData = file_torq_proto_rawDesc
)

func file_torq_proto_rawDescGZIP() []byte {
	file_torq_proto_rawDescOnce.Do(func() {
		file_torq_proto_rawDescData = protoimpl.X.CompressGZIP(file_torq_proto_rawDescData)
	})
	return file_torq_proto_rawDescData
}

var file_torq_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_torq_proto_goTypes = []interface{}{
	(*ChannelFlowRequest)(nil), // 0: torqrpc.ChannelFlowRequest
	(*ChannelFlow)(nil),        // 1: torqrpc.ChannelFlow
	(*AggregatedForwards)(nil), // 2: torqrpc.AggregatedForwards
	(*Forwards)(nil),           // 3: torqrpc.Forwards
	(*Forward)(nil),            // 4: torqrpc.Forward
	(*ForwardsRequest)(nil),    // 5: torqrpc.ForwardsRequest
}
var file_torq_proto_depIdxs = []int32{
	4, // 0: torqrpc.Forwards.forwards:type_name -> torqrpc.Forward
	5, // 1: torqrpc.torqrpc.GetForwards:input_type -> torqrpc.ForwardsRequest
	0, // 2: torqrpc.torqrpc.GetChannelFlow:input_type -> torqrpc.ChannelFlowRequest
	3, // 3: torqrpc.torqrpc.GetForwards:output_type -> torqrpc.Forwards
	1, // 4: torqrpc.torqrpc.GetChannelFlow:output_type -> torqrpc.ChannelFlow
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_torq_proto_init() }
func file_torq_proto_init() {
	if File_torq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_torq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelFlowRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChannelFlow); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AggregatedForwards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torq_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forwards); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torq_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Forward); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_torq_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ForwardsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_torq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_torq_proto_goTypes,
		DependencyIndexes: file_torq_proto_depIdxs,
		MessageInfos:      file_torq_proto_msgTypes,
	}.Build()
	File_torq_proto = out.File
	file_torq_proto_rawDesc = nil
	file_torq_proto_goTypes = nil
	file_torq_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TorqrpcClient is the client API for Torqrpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TorqrpcClient interface {
	GetForwards(ctx context.Context, in *ForwardsRequest, opts ...grpc.CallOption) (*Forwards, error)
	GetChannelFlow(ctx context.Context, in *ChannelFlowRequest, opts ...grpc.CallOption) (*ChannelFlow, error)
}

type torqrpcClient struct {
	cc grpc.ClientConnInterface
}

func NewTorqrpcClient(cc grpc.ClientConnInterface) TorqrpcClient {
	return &torqrpcClient{cc}
}

func (c *torqrpcClient) GetForwards(ctx context.Context, in *ForwardsRequest, opts ...grpc.CallOption) (*Forwards, error) {
	out := new(Forwards)
	err := c.cc.Invoke(ctx, "/torqrpc.torqrpc/GetForwards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *torqrpcClient) GetChannelFlow(ctx context.Context, in *ChannelFlowRequest, opts ...grpc.CallOption) (*ChannelFlow, error) {
	out := new(ChannelFlow)
	err := c.cc.Invoke(ctx, "/torqrpc.torqrpc/GetChannelFlow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TorqrpcServer is the server API for Torqrpc service.
type TorqrpcServer interface {
	GetForwards(context.Context, *ForwardsRequest) (*Forwards, error)
	GetChannelFlow(context.Context, *ChannelFlowRequest) (*ChannelFlow, error)
}

// UnimplementedTorqrpcServer can be embedded to have forward compatible implementations.
type UnimplementedTorqrpcServer struct {
}

func (*UnimplementedTorqrpcServer) GetForwards(context.Context, *ForwardsRequest) (*Forwards, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetForwards not implemented")
}
func (*UnimplementedTorqrpcServer) GetChannelFlow(context.Context, *ChannelFlowRequest) (*ChannelFlow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChannelFlow not implemented")
}

func RegisterTorqrpcServer(s *grpc.Server, srv TorqrpcServer) {
	s.RegisterService(&_Torqrpc_serviceDesc, srv)
}

func _Torqrpc_GetForwards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ForwardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqrpcServer).GetForwards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/torqrpc.torqrpc/GetForwards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqrpcServer).GetForwards(ctx, req.(*ForwardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Torqrpc_GetChannelFlow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChannelFlowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TorqrpcServer).GetChannelFlow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/torqrpc.torqrpc/GetChannelFlow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TorqrpcServer).GetChannelFlow(ctx, req.(*ChannelFlowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Torqrpc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "torqrpc.torqrpc",
	HandlerType: (*TorqrpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetForwards",
			Handler:    _Torqrpc_GetForwards_Handler,
		},
		{
			MethodName: "GetChannelFlow",
			Handler:    _Torqrpc_GetChannelFlow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "torq.proto",
}