// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: subscribedkeyword/v1/subscribed_keyword.proto

package subscribedkeywordv1

import (
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

// SubscribedKeyword represents a keyword a user subscribed to.
type SubscribedKeyword struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id represents the id of the record in the database.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// UserId represents the user id that subscribed to the word.
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	// Keyword represents the keyword they subscribed to.
	Keyword string `protobuf:"bytes,3,opt,name=keyword,proto3" json:"keyword,omitempty"`
}

func (x *SubscribedKeyword) Reset() {
	*x = SubscribedKeyword{}
	if protoimpl.UnsafeEnabled {
		mi := &file_subscribedkeyword_v1_subscribed_keyword_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribedKeyword) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribedKeyword) ProtoMessage() {}

func (x *SubscribedKeyword) ProtoReflect() protoreflect.Message {
	mi := &file_subscribedkeyword_v1_subscribed_keyword_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribedKeyword.ProtoReflect.Descriptor instead.
func (*SubscribedKeyword) Descriptor() ([]byte, []int) {
	return file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescGZIP(), []int{0}
}

func (x *SubscribedKeyword) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SubscribedKeyword) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SubscribedKeyword) GetKeyword() string {
	if x != nil {
		return x.Keyword
	}
	return ""
}

var File_subscribedkeyword_v1_subscribed_keyword_proto protoreflect.FileDescriptor

var file_subscribedkeyword_v1_subscribed_keyword_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x6b, 0x65, 0x79, 0x77,
	0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x5f, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x6b, 0x65, 0x79, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x76, 0x31, 0x22, 0x56, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x62, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x42, 0xec, 0x01,
	0x0a, 0x18, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x76, 0x31, 0x42, 0x16, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x4b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x6e, 0x77, 0x6f, 0x73, 0x75, 0x75, 0x64, 0x6f, 0x6b, 0x61, 0x2f, 0x73, 0x73, 0x73, 0x63,
	0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x6b,
	0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x64, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x53, 0x58, 0x58, 0xaa, 0x02, 0x14, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64,
	0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x14, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x20, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x64, 0x6b, 0x65,
	0x79, 0x77, 0x6f, 0x72, 0x64, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x15, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65,
	0x64, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescOnce sync.Once
	file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescData = file_subscribedkeyword_v1_subscribed_keyword_proto_rawDesc
)

func file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescGZIP() []byte {
	file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescOnce.Do(func() {
		file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescData = protoimpl.X.CompressGZIP(file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescData)
	})
	return file_subscribedkeyword_v1_subscribed_keyword_proto_rawDescData
}

var file_subscribedkeyword_v1_subscribed_keyword_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_subscribedkeyword_v1_subscribed_keyword_proto_goTypes = []interface{}{
	(*SubscribedKeyword)(nil), // 0: subscribedkeyword.v1.SubscribedKeyword
}
var file_subscribedkeyword_v1_subscribed_keyword_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_subscribedkeyword_v1_subscribed_keyword_proto_init() }
func file_subscribedkeyword_v1_subscribed_keyword_proto_init() {
	if File_subscribedkeyword_v1_subscribed_keyword_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_subscribedkeyword_v1_subscribed_keyword_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribedKeyword); i {
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
			RawDescriptor: file_subscribedkeyword_v1_subscribed_keyword_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_subscribedkeyword_v1_subscribed_keyword_proto_goTypes,
		DependencyIndexes: file_subscribedkeyword_v1_subscribed_keyword_proto_depIdxs,
		MessageInfos:      file_subscribedkeyword_v1_subscribed_keyword_proto_msgTypes,
	}.Build()
	File_subscribedkeyword_v1_subscribed_keyword_proto = out.File
	file_subscribedkeyword_v1_subscribed_keyword_proto_rawDesc = nil
	file_subscribedkeyword_v1_subscribed_keyword_proto_goTypes = nil
	file_subscribedkeyword_v1_subscribed_keyword_proto_depIdxs = nil
}
