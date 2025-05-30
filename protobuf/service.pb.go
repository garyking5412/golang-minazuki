// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: protobuf/service.proto

package services

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CategoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Detail        string                 `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryResponse) Reset() {
	*x = CategoryResponse{}
	mi := &file_protobuf_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryResponse) ProtoMessage() {}

func (x *CategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryResponse.ProtoReflect.Descriptor instead.
func (*CategoryResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_service_proto_rawDescGZIP(), []int{0}
}

func (x *CategoryResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CategoryResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CategoryResponse) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type CategoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CategoryRequest) Reset() {
	*x = CategoryRequest{}
	mi := &file_protobuf_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CategoryRequest) ProtoMessage() {}

func (x *CategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CategoryRequest.ProtoReflect.Descriptor instead.
func (*CategoryRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_service_proto_rawDescGZIP(), []int{1}
}

func (x *CategoryRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ConnectRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=clientId,proto3" json:"clientId,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	mi := &file_protobuf_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_service_proto_rawDescGZIP(), []int{2}
}

func (x *ConnectRequest) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type GreetingResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Response      string                 `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GreetingResponse) Reset() {
	*x = GreetingResponse{}
	mi := &file_protobuf_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GreetingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GreetingResponse) ProtoMessage() {}

func (x *GreetingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GreetingResponse.ProtoReflect.Descriptor instead.
func (*GreetingResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_service_proto_rawDescGZIP(), []int{3}
}

func (x *GreetingResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ChatMessage struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Sender        string                 `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Timestamp     int64                  `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatMessage) Reset() {
	*x = ChatMessage{}
	mi := &file_protobuf_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatMessage) ProtoMessage() {}

func (x *ChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatMessage.ProtoReflect.Descriptor instead.
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return file_protobuf_service_proto_rawDescGZIP(), []int{4}
}

func (x *ChatMessage) GetSender() string {
	if x != nil {
		return x.Sender
	}
	return ""
}

func (x *ChatMessage) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ChatMessage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

var File_protobuf_service_proto protoreflect.FileDescriptor

const file_protobuf_service_proto_rawDesc = "" +
	"\n" +
	"\x16protobuf/service.proto\"N\n" +
	"\x10CategoryResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x16\n" +
	"\x06detail\x18\x03 \x01(\tR\x06detail\"!\n" +
	"\x0fCategoryRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\",\n" +
	"\x0eConnectRequest\x12\x1a\n" +
	"\bclientId\x18\x01 \x01(\tR\bclientId\".\n" +
	"\x10GreetingResponse\x12\x1a\n" +
	"\bresponse\x18\x01 \x01(\tR\bresponse\"]\n" +
	"\vChatMessage\x12\x16\n" +
	"\x06sender\x18\x01 \x01(\tR\x06sender\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1c\n" +
	"\ttimestamp\x18\x03 \x01(\x03R\ttimestamp2\x9f\x01\n" +
	"\x0fCategoryService\x122\n" +
	"\vGetCategory\x12\x10.CategoryRequest\x1a\x11.CategoryResponse\x120\n" +
	"\bGreeting\x12\x0f.ConnectRequest\x1a\x11.GreetingResponse0\x01\x12&\n" +
	"\x04Chat\x12\f.ChatMessage\x1a\f.ChatMessage(\x010\x01B#Z!golang-minazuki/protobuf/servicesb\x06proto3"

var (
	file_protobuf_service_proto_rawDescOnce sync.Once
	file_protobuf_service_proto_rawDescData []byte
)

func file_protobuf_service_proto_rawDescGZIP() []byte {
	file_protobuf_service_proto_rawDescOnce.Do(func() {
		file_protobuf_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_protobuf_service_proto_rawDesc), len(file_protobuf_service_proto_rawDesc)))
	})
	return file_protobuf_service_proto_rawDescData
}

var file_protobuf_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protobuf_service_proto_goTypes = []any{
	(*CategoryResponse)(nil), // 0: CategoryResponse
	(*CategoryRequest)(nil),  // 1: CategoryRequest
	(*ConnectRequest)(nil),   // 2: ConnectRequest
	(*GreetingResponse)(nil), // 3: GreetingResponse
	(*ChatMessage)(nil),      // 4: ChatMessage
}
var file_protobuf_service_proto_depIdxs = []int32{
	1, // 0: CategoryService.GetCategory:input_type -> CategoryRequest
	2, // 1: CategoryService.Greeting:input_type -> ConnectRequest
	4, // 2: CategoryService.Chat:input_type -> ChatMessage
	0, // 3: CategoryService.GetCategory:output_type -> CategoryResponse
	3, // 4: CategoryService.Greeting:output_type -> GreetingResponse
	4, // 5: CategoryService.Chat:output_type -> ChatMessage
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protobuf_service_proto_init() }
func file_protobuf_service_proto_init() {
	if File_protobuf_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_protobuf_service_proto_rawDesc), len(file_protobuf_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_service_proto_goTypes,
		DependencyIndexes: file_protobuf_service_proto_depIdxs,
		MessageInfos:      file_protobuf_service_proto_msgTypes,
	}.Build()
	File_protobuf_service_proto = out.File
	file_protobuf_service_proto_goTypes = nil
	file_protobuf_service_proto_depIdxs = nil
}
