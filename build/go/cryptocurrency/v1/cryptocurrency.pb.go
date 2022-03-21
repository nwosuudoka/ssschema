// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: cryptocurrency/v1/cryptocurrency.proto

package cryptocurrencyv1

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
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

// Cryptocurrency represents the schema for a cryptocurrency
type Cryptocurrency struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the id of the record in the database.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"id"`
	// Name represents the name of the currency.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" bson:"name"`
	// Symbol represnets the symbol of the currency.
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty" bson:"symbol"`
	// Logo represents the logo of the cryptocurrency.
	Logo string `protobuf:"bytes,4,opt,name=logo,proto3" json:"logo,omitempty" bson:"logo"`
	// Platforms represents the platforms where this currencies are.
	Platforms []string `protobuf:"bytes,5,rep,name=platforms,proto3" json:"platforms,omitempty" bson:"platforms"`
	// Coinmarketcap represents the coin market cap data.
	// CoinmarketCap coin_market_cap = 6;
	CoinMarketCap *CoinmarketCap `protobuf:"bytes,6,opt,name=coin_market_cap,json=coinMarketCap,proto3" json:"coin_market_cap,omitempty" bson:"coin_market_cap"`
}

func (x *Cryptocurrency) Reset() {
	*x = Cryptocurrency{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cryptocurrency) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cryptocurrency) ProtoMessage() {}

func (x *Cryptocurrency) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cryptocurrency.ProtoReflect.Descriptor instead.
func (*Cryptocurrency) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{0}
}

func (x *Cryptocurrency) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Cryptocurrency) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cryptocurrency) GetSymbol() string {
	if x != nil {
		return x.Symbol
	}
	return ""
}

func (x *Cryptocurrency) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *Cryptocurrency) GetPlatforms() []string {
	if x != nil {
		return x.Platforms
	}
	return nil
}

func (x *Cryptocurrency) GetCoinMarketCap() *CoinmarketCap {
	if x != nil {
		return x.CoinMarketCap
	}
	return nil
}

// CoinmarketCap represents the coin market cap information.
type CoinmarketCap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represnets the coinmarketcap id.
	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" bson:"platforms"`
	// CoinmarketCapListing represents the coin listing info
	Listing *CoinmarketCapListing `protobuf:"bytes,2,opt,name=listing,proto3" json:"listing,omitempty" bson:"platforms"`
	// CoinmarketCapListing holds coinmarketcap info.
	Info *CoinmarketCapMetadata `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty" bson:"platforms"`
}

func (x *CoinmarketCap) Reset() {
	*x = CoinmarketCap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinmarketCap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinmarketCap) ProtoMessage() {}

func (x *CoinmarketCap) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinmarketCap.ProtoReflect.Descriptor instead.
func (*CoinmarketCap) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{1}
}

func (x *CoinmarketCap) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CoinmarketCap) GetListing() *CoinmarketCapListing {
	if x != nil {
		return x.Listing
	}
	return nil
}

func (x *CoinmarketCap) GetInfo() *CoinmarketCapMetadata {
	if x != nil {
		return x.Info
	}
	return nil
}

// CoinmarketCapQuote quote.
type CoinmarketCapQuote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Price represnets the price change.
	Price float64 `protobuf:"fixed64,1,opt,name=price,proto3" json:"price,omitempty" bson:"price"`
	// Volume24h.
	Volume_24H float64 `protobuf:"fixed64,2,opt,name=volume_24h,json=volume24h,proto3" json:"volume_24h,omitempty" bson:"volume_24h"`
	// PercentChange1h
	PercentChange_1H float64 `protobuf:"fixed64,3,opt,name=percent_change_1h,json=percentChange1h,proto3" json:"percent_change_1h,omitempty" bson:"percent_change_1h"`
	// PercentChange24h
	PercentChange_24H float64 `protobuf:"fixed64,4,opt,name=percent_change_24h,json=percentChange24h,proto3" json:"percent_change_24h,omitempty" bson:"percent_change_24h"`
	// PercentChange7d
	PercentChange_7D float64 `protobuf:"fixed64,5,opt,name=percent_change_7d,json=percentChange7d,proto3" json:"percent_change_7d,omitempty" bson:"percent_change_7dh"`
	// MarketCap
	MarketCap float64 `protobuf:"fixed64,6,opt,name=market_cap,json=marketCap,proto3" json:"market_cap,omitempty" bson:"market_cap"`
	// LastUpdated
	LastUpdated string `protobuf:"bytes,7,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty" bson:"last_updated"`
}

func (x *CoinmarketCapQuote) Reset() {
	*x = CoinmarketCapQuote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinmarketCapQuote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinmarketCapQuote) ProtoMessage() {}

func (x *CoinmarketCapQuote) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinmarketCapQuote.ProtoReflect.Descriptor instead.
func (*CoinmarketCapQuote) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{2}
}

func (x *CoinmarketCapQuote) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CoinmarketCapQuote) GetVolume_24H() float64 {
	if x != nil {
		return x.Volume_24H
	}
	return 0
}

func (x *CoinmarketCapQuote) GetPercentChange_1H() float64 {
	if x != nil {
		return x.PercentChange_1H
	}
	return 0
}

func (x *CoinmarketCapQuote) GetPercentChange_24H() float64 {
	if x != nil {
		return x.PercentChange_24H
	}
	return 0
}

func (x *CoinmarketCapQuote) GetPercentChange_7D() float64 {
	if x != nil {
		return x.PercentChange_7D
	}
	return 0
}

func (x *CoinmarketCapQuote) GetMarketCap() float64 {
	if x != nil {
		return x.MarketCap
	}
	return 0
}

func (x *CoinmarketCapQuote) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

// CoinmarketCapListing holds the info of the coin.
type CoinmarketCapListing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Quote represnets the quote of the cryptocurrency in different fiat currencies.
	Quote map[string]*CoinmarketCapQuote `protobuf:"bytes,1,rep,name=quote,proto3" json:"quote,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" bson:"quote"`
	// MaxSupply represents the max supply this currency can have.
	MaxSupply float64 `protobuf:"fixed64,2,opt,name=max_supply,json=maxSupply,proto3" json:"max_supply,omitempty" bson:"max_supply"`
	// TotalSupply represents the total supply this currency can have.
	TotalSupply float64 `protobuf:"fixed64,3,opt,name=total_supply,json=totalSupply,proto3" json:"total_supply,omitempty" bson:"total_supply"`
	// CMCRank represents the cmc rank of the coin.
	CmcRank int32 `protobuf:"varint,4,opt,name=cmc_rank,json=cmcRank,proto3" json:"cmc_rank,omitempty" bson:"cmc_rank"`
	// CirculatingSupply is how many is currenly in circulation.
	CirculatingSupply float64 `protobuf:"fixed64,5,opt,name=circulating_supply,json=circulatingSupply,proto3" json:"circulating_supply,omitempty" bson:"circulating_supply "`
	// LastUpdated was the last time this was updated.
	LastUpdated string `protobuf:"bytes,6,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty" bson:"last_updated"`
}

func (x *CoinmarketCapListing) Reset() {
	*x = CoinmarketCapListing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinmarketCapListing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinmarketCapListing) ProtoMessage() {}

func (x *CoinmarketCapListing) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinmarketCapListing.ProtoReflect.Descriptor instead.
func (*CoinmarketCapListing) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{3}
}

func (x *CoinmarketCapListing) GetQuote() map[string]*CoinmarketCapQuote {
	if x != nil {
		return x.Quote
	}
	return nil
}

func (x *CoinmarketCapListing) GetMaxSupply() float64 {
	if x != nil {
		return x.MaxSupply
	}
	return 0
}

func (x *CoinmarketCapListing) GetTotalSupply() float64 {
	if x != nil {
		return x.TotalSupply
	}
	return 0
}

func (x *CoinmarketCapListing) GetCmcRank() int32 {
	if x != nil {
		return x.CmcRank
	}
	return 0
}

func (x *CoinmarketCapListing) GetCirculatingSupply() float64 {
	if x != nil {
		return x.CirculatingSupply
	}
	return 0
}

func (x *CoinmarketCapListing) GetLastUpdated() string {
	if x != nil {
		return x.LastUpdated
	}
	return ""
}

// CoinmarketCapMetadata represents the coin metadata.
type CoinmarketCapMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Description represents the coin description.
	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty" bson:"description"`
	// DateAdded represents the date the coin was added.
	DateAdded string `protobuf:"bytes,2,opt,name=date_added,json=dateAdded,proto3" json:"date_added,omitempty" bson:"date_added"`
	// DateLaunched represents the date the coin was launched.
	DateLaunched string `protobuf:"bytes,3,opt,name=date_launched,json=dateLaunched,proto3" json:"date_launched,omitempty" bson:"date_launched"`
	// Category represents the category of the coin.
	Category string `protobuf:"bytes,4,opt,name=category,proto3" json:"category,omitempty" bson:"category"`
	// Logo represents the coin logo.
	Logo string `protobuf:"bytes,5,opt,name=logo,proto3" json:"logo,omitempty" bson:"logo"`
	// Urls represents other urls of the coin.
	Urls *CoinmarketCapUrls `protobuf:"bytes,6,opt,name=urls,proto3" json:"urls,omitempty" bson:"urls"`
}

func (x *CoinmarketCapMetadata) Reset() {
	*x = CoinmarketCapMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinmarketCapMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinmarketCapMetadata) ProtoMessage() {}

func (x *CoinmarketCapMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinmarketCapMetadata.ProtoReflect.Descriptor instead.
func (*CoinmarketCapMetadata) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{4}
}

func (x *CoinmarketCapMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CoinmarketCapMetadata) GetDateAdded() string {
	if x != nil {
		return x.DateAdded
	}
	return ""
}

func (x *CoinmarketCapMetadata) GetDateLaunched() string {
	if x != nil {
		return x.DateLaunched
	}
	return ""
}

func (x *CoinmarketCapMetadata) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CoinmarketCapMetadata) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *CoinmarketCapMetadata) GetUrls() *CoinmarketCapUrls {
	if x != nil {
		return x.Urls
	}
	return nil
}

// CoinmarketCapMetadata represents the md of the coin.
type CoinmarketCapUrls struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Website represents the website of the cmc data.
	Website []string `protobuf:"bytes,1,rep,name=website,proto3" json:"website,omitempty" bson:"website"`
	// Twitter represents the twitter info of the cmc data.
	Twitter []string `protobuf:"bytes,2,rep,name=twitter,proto3" json:"twitter,omitempty" bson:"twitter"`
	// Reddit represents the reddit info of the cmc data.
	Reddit []string `protobuf:"bytes,3,rep,name=reddit,proto3" json:"reddit,omitempty" bson:"reddit"`
	// SourceCode represents the source_code info of the cmc data.
	SourceCode []string `protobuf:"bytes,4,rep,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty" bson:"source_code"`
	// Chat represents the chat infomration of the currency.
	Chat []string `protobuf:"bytes,5,rep,name=chat,proto3" json:"chat,omitempty" bson:"chat"`
	// TechnicalDoc represents the technical doc of the data.
	TechnicalDoc []string `protobuf:"bytes,6,rep,name=technical_doc,json=technicalDoc,proto3" json:"technical_doc,omitempty" bson:"technical_doc"`
	// MessageBoards represents the technical doc of the data.
	MessageBoard []string `protobuf:"bytes,7,rep,name=message_board,json=messageBoard,proto3" json:"message_board,omitempty" bson:"message_board"`
	// Explorer represents the technical doc of the data.
	Explorer []string `protobuf:"bytes,8,rep,name=explorer,proto3" json:"explorer,omitempty" bson:"explorer"`
	// Announcement represents the technical doc of the data.
	Announcement []string `protobuf:"bytes,9,rep,name=announcement,proto3" json:"announcement,omitempty" bson:"annoucement"`
}

func (x *CoinmarketCapUrls) Reset() {
	*x = CoinmarketCapUrls{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoinmarketCapUrls) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoinmarketCapUrls) ProtoMessage() {}

func (x *CoinmarketCapUrls) ProtoReflect() protoreflect.Message {
	mi := &file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoinmarketCapUrls.ProtoReflect.Descriptor instead.
func (*CoinmarketCapUrls) Descriptor() ([]byte, []int) {
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP(), []int{5}
}

func (x *CoinmarketCapUrls) GetWebsite() []string {
	if x != nil {
		return x.Website
	}
	return nil
}

func (x *CoinmarketCapUrls) GetTwitter() []string {
	if x != nil {
		return x.Twitter
	}
	return nil
}

func (x *CoinmarketCapUrls) GetReddit() []string {
	if x != nil {
		return x.Reddit
	}
	return nil
}

func (x *CoinmarketCapUrls) GetSourceCode() []string {
	if x != nil {
		return x.SourceCode
	}
	return nil
}

func (x *CoinmarketCapUrls) GetChat() []string {
	if x != nil {
		return x.Chat
	}
	return nil
}

func (x *CoinmarketCapUrls) GetTechnicalDoc() []string {
	if x != nil {
		return x.TechnicalDoc
	}
	return nil
}

func (x *CoinmarketCapUrls) GetMessageBoard() []string {
	if x != nil {
		return x.MessageBoard
	}
	return nil
}

func (x *CoinmarketCapUrls) GetExplorer() []string {
	if x != nil {
		return x.Explorer
	}
	return nil
}

func (x *CoinmarketCapUrls) GetAnnouncement() []string {
	if x != nil {
		return x.Announcement
	}
	return nil
}

var File_cryptocurrency_v1_cryptocurrency_proto protoreflect.FileDescriptor

var file_cryptocurrency_v1_cryptocurrency_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc4, 0x02, 0x0a, 0x0e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x63, 0x79, 0x12, 0x1e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x24, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x06, 0x73, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x22, 0x52, 0x06, 0x73,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x6c, 0x6f, 0x67, 0x6f, 0x22, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x33, 0x0a, 0x09, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x42, 0x15,
	0x9a, 0x84, 0x9e, 0x03, 0x10, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66,
	0x6f, 0x72, 0x6d, 0x73, 0x22, 0x52, 0x09, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73,
	0x12, 0x65, 0x0a, 0x0f, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f,
	0x63, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f,
	0x69, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x42, 0x1b, 0x9a, 0x84, 0x9e,
	0x03, 0x16, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6f, 0x69, 0x6e, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x22, 0x52, 0x0d, 0x63, 0x6f, 0x69, 0x6e, 0x4d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x22, 0xe5, 0x01, 0x0a, 0x0d, 0x43, 0x6f, 0x69, 0x6e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x22, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x58, 0x0a, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x43, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03,
	0x10, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73,
	0x22, 0x52, 0x07, 0x6c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x53, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69,
	0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70,
	0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d, 0x73, 0x22, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22,
	0xcd, 0x03, 0x0a, 0x12, 0x43, 0x6f, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x70, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x01, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x62, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x35, 0x0a, 0x0a, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x22, 0x52, 0x09, 0x76, 0x6f, 0x6c,
	0x75, 0x6d, 0x65, 0x32, 0x34, 0x68, 0x12, 0x49, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x31, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x01, 0x42, 0x1d, 0x9a, 0x84, 0x9e, 0x03, 0x18, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65,
	0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x31, 0x68, 0x22,
	0x52, 0x0f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x31,
	0x68, 0x12, 0x4c, 0x0a, 0x12, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61,
	0x6e, 0x67, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x42, 0x1e, 0x9a,
	0x84, 0x9e, 0x03, 0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e,
	0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x32, 0x34, 0x68, 0x22, 0x52, 0x10, 0x70,
	0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x32, 0x34, 0x68, 0x12,
	0x4a, 0x0a, 0x11, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x5f, 0x37, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x1e, 0x9a, 0x84, 0x9e, 0x03,
	0x19, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x5f, 0x63,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x5f, 0x37, 0x64, 0x68, 0x22, 0x52, 0x0f, 0x70, 0x65, 0x72, 0x63,
	0x65, 0x6e, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x37, 0x64, 0x12, 0x35, 0x0a, 0x0a, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x42,
	0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x5f, 0x63, 0x61, 0x70, 0x22, 0x52, 0x09, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43,
	0x61, 0x70, 0x12, 0x3b, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62,
	0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x22, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22,
	0x86, 0x04, 0x0a, 0x14, 0x43, 0x6f, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61,
	0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x5b, 0x0a, 0x05, 0x71, 0x75, 0x6f, 0x74,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x32, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2e, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x11, 0x9a, 0x84, 0x9e,
	0x03, 0x0c, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x22, 0x52, 0x05,
	0x71, 0x75, 0x6f, 0x74, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75, 0x70,
	0x70, 0x6c, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11,
	0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x61, 0x78, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x22, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0c,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x22, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x08, 0x63, 0x6d, 0x63,
	0x5f, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0x9a, 0x84, 0x9e,
	0x03, 0x0f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x6d, 0x63, 0x5f, 0x72, 0x61, 0x6e, 0x6b,
	0x22, 0x52, 0x07, 0x63, 0x6d, 0x63, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x4e, 0x0a, 0x12, 0x63, 0x69,
	0x72, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x79,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x42, 0x1f, 0x9a, 0x84, 0x9e, 0x03, 0x1a, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x73,
	0x75, 0x70, 0x70, 0x6c, 0x79, 0x20, 0x22, 0x52, 0x11, 0x63, 0x69, 0x72, 0x63, 0x75, 0x6c, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x53, 0x75, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x3b, 0x0a, 0x0c, 0x6c, 0x61,
	0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x13, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x1a, 0x5f, 0x0a, 0x0a, 0x51, 0x75, 0x6f, 0x74, 0x65,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x51, 0x75, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xed, 0x02, 0x0a, 0x15, 0x43, 0x6f, 0x69,
	0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x39, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x0a, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x65, 0x64, 0x22, 0x52, 0x09, 0x64, 0x61, 0x74, 0x65, 0x41,
	0x64, 0x64, 0x65, 0x64, 0x12, 0x3e, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x19, 0x9a, 0x84, 0x9e,
	0x03, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x61, 0x75,
	0x6e, 0x63, 0x68, 0x65, 0x64, 0x22, 0x52, 0x0c, 0x64, 0x61, 0x74, 0x65, 0x4c, 0x61, 0x75, 0x6e,
	0x63, 0x68, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03, 0x0f, 0x62, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x52, 0x08, 0x63, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x6c, 0x6f, 0x67, 0x6f, 0x22, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x6f, 0x12, 0x4a, 0x0a, 0x04,
	0x75, 0x72, 0x6c, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x63, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6f, 0x69, 0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x55, 0x72, 0x6c, 0x73,
	0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x72, 0x6c,
	0x73, 0x22, 0x52, 0x04, 0x75, 0x72, 0x6c, 0x73, 0x22, 0xec, 0x03, 0x0a, 0x11, 0x43, 0x6f, 0x69,
	0x6e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x43, 0x61, 0x70, 0x55, 0x72, 0x6c, 0x73, 0x12, 0x2d,
	0x0a, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x13, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x77, 0x65, 0x62, 0x73,
	0x69, 0x74, 0x65, 0x22, 0x52, 0x07, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x12, 0x2d, 0x0a,
	0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x42, 0x13,
	0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x77, 0x69, 0x74, 0x74,
	0x65, 0x72, 0x22, 0x52, 0x07, 0x74, 0x77, 0x69, 0x74, 0x74, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x06,
	0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84,
	0x9e, 0x03, 0x0d, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x22,
	0x52, 0x06, 0x72, 0x65, 0x64, 0x64, 0x69, 0x74, 0x12, 0x38, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x17, 0x9a,
	0x84, 0x9e, 0x03, 0x12, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x52, 0x0a, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x68, 0x61, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x68, 0x61,
	0x74, 0x22, 0x52, 0x04, 0x63, 0x68, 0x61, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x6f, 0x63, 0x22, 0x52, 0x0c, 0x74, 0x65, 0x63, 0x68,
	0x6e, 0x69, 0x63, 0x61, 0x6c, 0x44, 0x6f, 0x63, 0x12, 0x3e, 0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x42,
	0x19, 0x9a, 0x84, 0x9e, 0x03, 0x14, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x22, 0x52, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x30, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x6c,
	0x6f, 0x72, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x42, 0x14, 0x9a, 0x84, 0x9e, 0x03,
	0x0f, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x22,
	0x52, 0x08, 0x65, 0x78, 0x70, 0x6c, 0x6f, 0x72, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x0c, 0x61, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09,
	0x42, 0x17, 0x9a, 0x84, 0x9e, 0x03, 0x12, 0x62, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x52, 0x0c, 0x61, 0x6e, 0x6e, 0x6f, 0x75,
	0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0xd4, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x13, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63,
	0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6e, 0x77, 0x6f, 0x73, 0x75, 0x75, 0x64, 0x6f, 0x6b, 0x61, 0x2f,
	0x73, 0x73, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x43, 0x58,
	0x58, 0xaa, 0x02, 0x11, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x11, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1d, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cryptocurrency_v1_cryptocurrency_proto_rawDescOnce sync.Once
	file_cryptocurrency_v1_cryptocurrency_proto_rawDescData = file_cryptocurrency_v1_cryptocurrency_proto_rawDesc
)

func file_cryptocurrency_v1_cryptocurrency_proto_rawDescGZIP() []byte {
	file_cryptocurrency_v1_cryptocurrency_proto_rawDescOnce.Do(func() {
		file_cryptocurrency_v1_cryptocurrency_proto_rawDescData = protoimpl.X.CompressGZIP(file_cryptocurrency_v1_cryptocurrency_proto_rawDescData)
	})
	return file_cryptocurrency_v1_cryptocurrency_proto_rawDescData
}

var file_cryptocurrency_v1_cryptocurrency_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_cryptocurrency_v1_cryptocurrency_proto_goTypes = []interface{}{
	(*Cryptocurrency)(nil),        // 0: cryptocurrency.v1.Cryptocurrency
	(*CoinmarketCap)(nil),         // 1: cryptocurrency.v1.CoinmarketCap
	(*CoinmarketCapQuote)(nil),    // 2: cryptocurrency.v1.CoinmarketCapQuote
	(*CoinmarketCapListing)(nil),  // 3: cryptocurrency.v1.CoinmarketCapListing
	(*CoinmarketCapMetadata)(nil), // 4: cryptocurrency.v1.CoinmarketCapMetadata
	(*CoinmarketCapUrls)(nil),     // 5: cryptocurrency.v1.CoinmarketCapUrls
	nil,                           // 6: cryptocurrency.v1.CoinmarketCapListing.QuoteEntry
}
var file_cryptocurrency_v1_cryptocurrency_proto_depIdxs = []int32{
	1, // 0: cryptocurrency.v1.Cryptocurrency.coin_market_cap:type_name -> cryptocurrency.v1.CoinmarketCap
	3, // 1: cryptocurrency.v1.CoinmarketCap.listing:type_name -> cryptocurrency.v1.CoinmarketCapListing
	4, // 2: cryptocurrency.v1.CoinmarketCap.info:type_name -> cryptocurrency.v1.CoinmarketCapMetadata
	6, // 3: cryptocurrency.v1.CoinmarketCapListing.quote:type_name -> cryptocurrency.v1.CoinmarketCapListing.QuoteEntry
	5, // 4: cryptocurrency.v1.CoinmarketCapMetadata.urls:type_name -> cryptocurrency.v1.CoinmarketCapUrls
	2, // 5: cryptocurrency.v1.CoinmarketCapListing.QuoteEntry.value:type_name -> cryptocurrency.v1.CoinmarketCapQuote
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_cryptocurrency_v1_cryptocurrency_proto_init() }
func file_cryptocurrency_v1_cryptocurrency_proto_init() {
	if File_cryptocurrency_v1_cryptocurrency_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cryptocurrency); i {
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
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinmarketCap); i {
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
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinmarketCapQuote); i {
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
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinmarketCapListing); i {
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
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinmarketCapMetadata); i {
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
		file_cryptocurrency_v1_cryptocurrency_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoinmarketCapUrls); i {
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
			RawDescriptor: file_cryptocurrency_v1_cryptocurrency_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cryptocurrency_v1_cryptocurrency_proto_goTypes,
		DependencyIndexes: file_cryptocurrency_v1_cryptocurrency_proto_depIdxs,
		MessageInfos:      file_cryptocurrency_v1_cryptocurrency_proto_msgTypes,
	}.Build()
	File_cryptocurrency_v1_cryptocurrency_proto = out.File
	file_cryptocurrency_v1_cryptocurrency_proto_rawDesc = nil
	file_cryptocurrency_v1_cryptocurrency_proto_goTypes = nil
	file_cryptocurrency_v1_cryptocurrency_proto_depIdxs = nil
}
