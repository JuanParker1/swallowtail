// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.17.3
// source: s.coingecko/proto/coingecko.proto

package coingeckoproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type CoinMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CoinMetadata) Reset() {
	*x = CoinMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinMetadata) ProtoMessage() {}

func (x *CoinMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinMetadata.ProtoReflect.Descriptor instead.
func (*CoinMetadata) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{0}
}

type GetAssetLatestPriceByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId string `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetPair       string `protobuf:"bytes,2,opt,name=asset_pair,json=assetPair,proto3" json:"asset_pair,omitempty"`
}

func (x *GetAssetLatestPriceByIDRequest) Reset() {
	*x = GetAssetLatestPriceByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetLatestPriceByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetLatestPriceByIDRequest) ProtoMessage() {}

func (x *GetAssetLatestPriceByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetLatestPriceByIDRequest.ProtoReflect.Descriptor instead.
func (*GetAssetLatestPriceByIDRequest) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{1}
}

func (x *GetAssetLatestPriceByIDRequest) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetAssetLatestPriceByIDRequest) GetAssetPair() string {
	if x != nil {
		return x.AssetPair
	}
	return ""
}

type GetAssetLatestPriceByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId           string  `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetSymbol               string  `protobuf:"bytes,2,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	LatestPrice               float32 `protobuf:"fixed32,3,opt,name=latest_price,json=latestPrice,proto3" json:"latest_price,omitempty"`
	PercentagePriceChange_24H float32 `protobuf:"fixed32,4,opt,name=percentage_price_change_24h,json=percentagePriceChange24h,proto3" json:"percentage_price_change_24h,omitempty"`
}

func (x *GetAssetLatestPriceByIDResponse) Reset() {
	*x = GetAssetLatestPriceByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetLatestPriceByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetLatestPriceByIDResponse) ProtoMessage() {}

func (x *GetAssetLatestPriceByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetLatestPriceByIDResponse.ProtoReflect.Descriptor instead.
func (*GetAssetLatestPriceByIDResponse) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{2}
}

func (x *GetAssetLatestPriceByIDResponse) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetAssetLatestPriceByIDResponse) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetAssetLatestPriceByIDResponse) GetLatestPrice() float32 {
	if x != nil {
		return x.LatestPrice
	}
	return 0
}

func (x *GetAssetLatestPriceByIDResponse) GetPercentagePriceChange_24H() float32 {
	if x != nil {
		return x.PercentagePriceChange_24H
	}
	return 0
}

type GetAssetLatestPriceBySymbolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetSymbol string `protobuf:"bytes,1,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	AssetPair   string `protobuf:"bytes,2,opt,name=asset_pair,json=assetPair,proto3" json:"asset_pair,omitempty"`
}

func (x *GetAssetLatestPriceBySymbolRequest) Reset() {
	*x = GetAssetLatestPriceBySymbolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetLatestPriceBySymbolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetLatestPriceBySymbolRequest) ProtoMessage() {}

func (x *GetAssetLatestPriceBySymbolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetLatestPriceBySymbolRequest.ProtoReflect.Descriptor instead.
func (*GetAssetLatestPriceBySymbolRequest) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetLatestPriceBySymbolRequest) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetAssetLatestPriceBySymbolRequest) GetAssetPair() string {
	if x != nil {
		return x.AssetPair
	}
	return ""
}

type GetAssetLatestPriceBySymbolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId           string  `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetSymbol               string  `protobuf:"bytes,2,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	LatestPrice               float32 `protobuf:"fixed32,3,opt,name=latest_price,json=latestPrice,proto3" json:"latest_price,omitempty"`
	PercentagePriceChange_24H float32 `protobuf:"fixed32,4,opt,name=percentage_price_change_24h,json=percentagePriceChange24h,proto3" json:"percentage_price_change_24h,omitempty"`
}

func (x *GetAssetLatestPriceBySymbolResponse) Reset() {
	*x = GetAssetLatestPriceBySymbolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetLatestPriceBySymbolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetLatestPriceBySymbolResponse) ProtoMessage() {}

func (x *GetAssetLatestPriceBySymbolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetLatestPriceBySymbolResponse.ProtoReflect.Descriptor instead.
func (*GetAssetLatestPriceBySymbolResponse) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssetLatestPriceBySymbolResponse) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetAssetLatestPriceBySymbolResponse) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetAssetLatestPriceBySymbolResponse) GetLatestPrice() float32 {
	if x != nil {
		return x.LatestPrice
	}
	return 0
}

func (x *GetAssetLatestPriceBySymbolResponse) GetPercentagePriceChange_24H() float32 {
	if x != nil {
		return x.PercentagePriceChange_24H
	}
	return 0
}

type GetATHBySymbolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetSymbol string `protobuf:"bytes,1,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	AssetPair   string `protobuf:"bytes,2,opt,name=asset_pair,json=assetPair,proto3" json:"asset_pair,omitempty"`
}

func (x *GetATHBySymbolRequest) Reset() {
	*x = GetATHBySymbolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATHBySymbolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATHBySymbolRequest) ProtoMessage() {}

func (x *GetATHBySymbolRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATHBySymbolRequest.ProtoReflect.Descriptor instead.
func (*GetATHBySymbolRequest) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{5}
}

func (x *GetATHBySymbolRequest) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetATHBySymbolRequest) GetAssetPair() string {
	if x != nil {
		return x.AssetPair
	}
	return ""
}

type GetATHBySymbolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId  string  `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetSymbol      string  `protobuf:"bytes,2,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	AllTimeHighPrice float32 `protobuf:"fixed32,3,opt,name=all_time_high_price,json=allTimeHighPrice,proto3" json:"all_time_high_price,omitempty"`
	CurrentPrice     float32 `protobuf:"fixed32,4,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
}

func (x *GetATHBySymbolResponse) Reset() {
	*x = GetATHBySymbolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATHBySymbolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATHBySymbolResponse) ProtoMessage() {}

func (x *GetATHBySymbolResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATHBySymbolResponse.ProtoReflect.Descriptor instead.
func (*GetATHBySymbolResponse) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{6}
}

func (x *GetATHBySymbolResponse) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetATHBySymbolResponse) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetATHBySymbolResponse) GetAllTimeHighPrice() float32 {
	if x != nil {
		return x.AllTimeHighPrice
	}
	return 0
}

func (x *GetATHBySymbolResponse) GetCurrentPrice() float32 {
	if x != nil {
		return x.CurrentPrice
	}
	return 0
}

type GetATHByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId string `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetPair       string `protobuf:"bytes,2,opt,name=asset_pair,json=assetPair,proto3" json:"asset_pair,omitempty"`
}

func (x *GetATHByIDRequest) Reset() {
	*x = GetATHByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATHByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATHByIDRequest) ProtoMessage() {}

func (x *GetATHByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATHByIDRequest.ProtoReflect.Descriptor instead.
func (*GetATHByIDRequest) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{7}
}

func (x *GetATHByIDRequest) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetATHByIDRequest) GetAssetPair() string {
	if x != nil {
		return x.AssetPair
	}
	return ""
}

type GetATHByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CoingeckoCoinId  string  `protobuf:"bytes,1,opt,name=coingecko_coin_id,json=coingeckoCoinId,proto3" json:"coingecko_coin_id,omitempty"`
	AssetSymbol      string  `protobuf:"bytes,2,opt,name=asset_symbol,json=assetSymbol,proto3" json:"asset_symbol,omitempty"`
	AllTimeHighPrice float32 `protobuf:"fixed32,3,opt,name=all_time_high_price,json=allTimeHighPrice,proto3" json:"all_time_high_price,omitempty"`
	CurrentPrice     float32 `protobuf:"fixed32,4,opt,name=current_price,json=currentPrice,proto3" json:"current_price,omitempty"`
}

func (x *GetATHByIDResponse) Reset() {
	*x = GetATHByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetATHByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetATHByIDResponse) ProtoMessage() {}

func (x *GetATHByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_s_coingecko_proto_coingecko_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetATHByIDResponse.ProtoReflect.Descriptor instead.
func (*GetATHByIDResponse) Descriptor() ([]byte, []int) {
	return file_s_coingecko_proto_coingecko_proto_rawDescGZIP(), []int{8}
}

func (x *GetATHByIDResponse) GetCoingeckoCoinId() string {
	if x != nil {
		return x.CoingeckoCoinId
	}
	return ""
}

func (x *GetATHByIDResponse) GetAssetSymbol() string {
	if x != nil {
		return x.AssetSymbol
	}
	return ""
}

func (x *GetATHByIDResponse) GetAllTimeHighPrice() float32 {
	if x != nil {
		return x.AllTimeHighPrice
	}
	return 0
}

func (x *GetATHByIDResponse) GetCurrentPrice() float32 {
	if x != nil {
		return x.CurrentPrice
	}
	return 0
}

var File_s_coingecko_proto_coingecko_proto protoreflect.FileDescriptor

var file_s_coingecko_proto_coingecko_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6f, 0x69, 0x6e, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x6b, 0x0a, 0x1e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72,
	0x22, 0xd2, 0x01, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b,
	0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x18, 0x70, 0x65, 0x72,
	0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x32, 0x34, 0x68, 0x22, 0x66, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x22, 0xd6, 0x01,
	0x0a, 0x23, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e, 0x49,
	0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x6c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x1b, 0x70, 0x65, 0x72, 0x63, 0x65,
	0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x18, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x32, 0x34, 0x68, 0x22, 0x59, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48,
	0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x61, 0x69,
	0x72, 0x22, 0xbb, 0x01, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11,
	0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63,
	0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x61,
	0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d,
	0x65, 0x48, 0x69, 0x67, 0x68, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22,
	0x5e, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b,
	0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x73, 0x73, 0x65, 0x74, 0x50, 0x61, 0x69, 0x72, 0x22,
	0xb7, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65,
	0x63, 0x6b, 0x6f, 0x5f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x43, 0x6f, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x73, 0x73, 0x65, 0x74, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x2d, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x5f, 0x68, 0x69, 0x67, 0x68, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x10, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x69, 0x67, 0x68, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x32, 0xd5, 0x02, 0x0a, 0x09, 0x63, 0x6f,
	0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b, 0x6f, 0x12, 0x5e, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x12, 0x1f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x23, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65,
	0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x42, 0x79, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x4c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x49,
	0x44, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x16,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42, 0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x54, 0x48, 0x42,
	0x79, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x3b, 0x63, 0x6f, 0x69, 0x6e, 0x67, 0x65, 0x63, 0x6b,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_s_coingecko_proto_coingecko_proto_rawDescOnce sync.Once
	file_s_coingecko_proto_coingecko_proto_rawDescData = file_s_coingecko_proto_coingecko_proto_rawDesc
)

func file_s_coingecko_proto_coingecko_proto_rawDescGZIP() []byte {
	file_s_coingecko_proto_coingecko_proto_rawDescOnce.Do(func() {
		file_s_coingecko_proto_coingecko_proto_rawDescData = protoimpl.X.CompressGZIP(file_s_coingecko_proto_coingecko_proto_rawDescData)
	})
	return file_s_coingecko_proto_coingecko_proto_rawDescData
}

var file_s_coingecko_proto_coingecko_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_s_coingecko_proto_coingecko_proto_goTypes = []interface{}{
	(*CoinMetadata)(nil),                        // 0: CoinMetadata
	(*GetAssetLatestPriceByIDRequest)(nil),      // 1: GetAssetLatestPriceByIDRequest
	(*GetAssetLatestPriceByIDResponse)(nil),     // 2: GetAssetLatestPriceByIDResponse
	(*GetAssetLatestPriceBySymbolRequest)(nil),  // 3: GetAssetLatestPriceBySymbolRequest
	(*GetAssetLatestPriceBySymbolResponse)(nil), // 4: GetAssetLatestPriceBySymbolResponse
	(*GetATHBySymbolRequest)(nil),               // 5: GetATHBySymbolRequest
	(*GetATHBySymbolResponse)(nil),              // 6: GetATHBySymbolResponse
	(*GetATHByIDRequest)(nil),                   // 7: GetATHByIDRequest
	(*GetATHByIDResponse)(nil),                  // 8: GetATHByIDResponse
}
var file_s_coingecko_proto_coingecko_proto_depIdxs = []int32{
	1, // 0: coingecko.GetAssetLatestPriceByID:input_type -> GetAssetLatestPriceByIDRequest
	3, // 1: coingecko.GetAssetLatestPriceBySymbol:input_type -> GetAssetLatestPriceBySymbolRequest
	7, // 2: coingecko.GetATHByID:input_type -> GetATHByIDRequest
	5, // 3: coingecko.GetATHBySymbol:input_type -> GetATHBySymbolRequest
	2, // 4: coingecko.GetAssetLatestPriceByID:output_type -> GetAssetLatestPriceByIDResponse
	4, // 5: coingecko.GetAssetLatestPriceBySymbol:output_type -> GetAssetLatestPriceBySymbolResponse
	8, // 6: coingecko.GetATHByID:output_type -> GetATHByIDResponse
	6, // 7: coingecko.GetATHBySymbol:output_type -> GetATHBySymbolResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_s_coingecko_proto_coingecko_proto_init() }
func file_s_coingecko_proto_coingecko_proto_init() {
	if File_s_coingecko_proto_coingecko_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_s_coingecko_proto_coingecko_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinMetadata); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetLatestPriceByIDRequest); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetLatestPriceByIDResponse); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetLatestPriceBySymbolRequest); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetLatestPriceBySymbolResponse); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATHBySymbolRequest); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATHBySymbolResponse); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATHByIDRequest); i {
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
		file_s_coingecko_proto_coingecko_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetATHByIDResponse); i {
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
			RawDescriptor: file_s_coingecko_proto_coingecko_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_s_coingecko_proto_coingecko_proto_goTypes,
		DependencyIndexes: file_s_coingecko_proto_coingecko_proto_depIdxs,
		MessageInfos:      file_s_coingecko_proto_coingecko_proto_msgTypes,
	}.Build()
	File_s_coingecko_proto_coingecko_proto = out.File
	file_s_coingecko_proto_coingecko_proto_rawDesc = nil
	file_s_coingecko_proto_coingecko_proto_goTypes = nil
	file_s_coingecko_proto_coingecko_proto_depIdxs = nil
}
