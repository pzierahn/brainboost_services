// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: chat_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CollectionId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CollectionId) Reset() {
	*x = CollectionId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionId) ProtoMessage() {}

func (x *CollectionId) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionId.ProtoReflect.Descriptor instead.
func (*CollectionId) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{0}
}

func (x *CollectionId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CompletionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId   string        `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Prompt       string        `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	ModelOptions *ModelOptions `protobuf:"bytes,3,opt,name=model_options,json=modelOptions,proto3" json:"model_options,omitempty"`
}

func (x *CompletionRequest) Reset() {
	*x = CompletionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionRequest) ProtoMessage() {}

func (x *CompletionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionRequest.ProtoReflect.Descriptor instead.
func (*CompletionRequest) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{1}
}

func (x *CompletionRequest) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *CompletionRequest) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *CompletionRequest) GetModelOptions() *ModelOptions {
	if x != nil {
		return x.ModelOptions
	}
	return nil
}

type CompletionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Completion string `protobuf:"bytes,1,opt,name=completion,proto3" json:"completion,omitempty"`
}

func (x *CompletionResponse) Reset() {
	*x = CompletionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompletionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompletionResponse) ProtoMessage() {}

func (x *CompletionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompletionResponse.ProtoReflect.Descriptor instead.
func (*CompletionResponse) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{2}
}

func (x *CompletionResponse) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

type Prompt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Thread ID to post the message to
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// Collection ID to post the message to
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// Prompt to generate completion
	Prompt string `protobuf:"bytes,3,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Model options
	ModelOptions *ModelOptions `protobuf:"bytes,4,opt,name=model_options,json=modelOptions,proto3" json:"model_options,omitempty"`
	// Search options
	RetrievalOptions *RetrievalOptions `protobuf:"bytes,5,opt,name=retrieval_options,json=retrievalOptions,proto3" json:"retrieval_options,omitempty"`
	// Attachments to the prompt
	Attachments []string `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *Prompt) Reset() {
	*x = Prompt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Prompt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Prompt) ProtoMessage() {}

func (x *Prompt) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Prompt.ProtoReflect.Descriptor instead.
func (*Prompt) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{3}
}

func (x *Prompt) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *Prompt) GetCollectionId() string {
	if x != nil {
		return x.CollectionId
	}
	return ""
}

func (x *Prompt) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *Prompt) GetModelOptions() *ModelOptions {
	if x != nil {
		return x.ModelOptions
	}
	return nil
}

func (x *Prompt) GetRetrievalOptions() *RetrievalOptions {
	if x != nil {
		return x.RetrievalOptions
	}
	return nil
}

func (x *Prompt) GetAttachments() []string {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type ModelOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ModelId     string  `protobuf:"bytes,1,opt,name=model_id,json=modelId,proto3" json:"model_id,omitempty"`
	Temperature float32 `protobuf:"fixed32,2,opt,name=temperature,proto3" json:"temperature,omitempty"`
	MaxTokens   uint32  `protobuf:"varint,3,opt,name=max_tokens,json=maxTokens,proto3" json:"max_tokens,omitempty"`
	TopP        float32 `protobuf:"fixed32,4,opt,name=top_p,json=topP,proto3" json:"top_p,omitempty"`
}

func (x *ModelOptions) Reset() {
	*x = ModelOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModelOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModelOptions) ProtoMessage() {}

func (x *ModelOptions) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModelOptions.ProtoReflect.Descriptor instead.
func (*ModelOptions) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{4}
}

func (x *ModelOptions) GetModelId() string {
	if x != nil {
		return x.ModelId
	}
	return ""
}

func (x *ModelOptions) GetTemperature() float32 {
	if x != nil {
		return x.Temperature
	}
	return 0
}

func (x *ModelOptions) GetMaxTokens() uint32 {
	if x != nil {
		return x.MaxTokens
	}
	return 0
}

func (x *ModelOptions) GetTopP() float32 {
	if x != nil {
		return x.TopP
	}
	return 0
}

type RetrievalOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enabled   bool    `protobuf:"varint,1,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Threshold float32 `protobuf:"fixed32,2,opt,name=threshold,proto3" json:"threshold,omitempty"`
	Documents uint32  `protobuf:"varint,3,opt,name=documents,proto3" json:"documents,omitempty"`
}

func (x *RetrievalOptions) Reset() {
	*x = RetrievalOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetrievalOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetrievalOptions) ProtoMessage() {}

func (x *RetrievalOptions) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetrievalOptions.ProtoReflect.Descriptor instead.
func (*RetrievalOptions) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{5}
}

func (x *RetrievalOptions) GetEnabled() bool {
	if x != nil {
		return x.Enabled
	}
	return false
}

func (x *RetrievalOptions) GetThreshold() float32 {
	if x != nil {
		return x.Threshold
	}
	return 0
}

func (x *RetrievalOptions) GetDocuments() uint32 {
	if x != nil {
		return x.Documents
	}
	return 0
}

type Source struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DocumentId string             `protobuf:"bytes,1,opt,name=document_id,json=documentId,proto3" json:"document_id,omitempty"`
	Name       string             `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Fragments  []*Source_Fragment `protobuf:"bytes,3,rep,name=fragments,proto3" json:"fragments,omitempty"`
}

func (x *Source) Reset() {
	*x = Source{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source) ProtoMessage() {}

func (x *Source) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source.ProtoReflect.Descriptor instead.
func (*Source) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{6}
}

func (x *Source) GetDocumentId() string {
	if x != nil {
		return x.DocumentId
	}
	return ""
}

func (x *Source) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Source) GetFragments() []*Source_Fragment {
	if x != nil {
		return x.Fragments
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Unique ID of the message
	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	// Prompt used to generate the message
	Prompt string `protobuf:"bytes,2,opt,name=prompt,proto3" json:"prompt,omitempty"`
	// Generated completion
	Completion string `protobuf:"bytes,3,opt,name=completion,proto3" json:"completion,omitempty"`
	// Sources used to generate the completion
	Sources []*Source `protobuf:"bytes,4,rep,name=sources,proto3" json:"sources,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{7}
}

func (x *Message) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *Message) GetPrompt() string {
	if x != nil {
		return x.Prompt
	}
	return ""
}

func (x *Message) GetCompletion() string {
	if x != nil {
		return x.Completion
	}
	return ""
}

func (x *Message) GetSources() []*Source {
	if x != nil {
		return x.Sources
	}
	return nil
}

type Thread struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Messages []*Message `protobuf:"bytes,2,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (x *Thread) Reset() {
	*x = Thread{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Thread) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Thread) ProtoMessage() {}

func (x *Thread) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Thread.ProtoReflect.Descriptor instead.
func (*Thread) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{8}
}

func (x *Thread) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Thread) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type ThreadID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ThreadID) Reset() {
	*x = ThreadID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadID) ProtoMessage() {}

func (x *ThreadID) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadID.ProtoReflect.Descriptor instead.
func (*ThreadID) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{9}
}

func (x *ThreadID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MessageIndex struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ThreadId string `protobuf:"bytes,1,opt,name=thread_id,json=threadId,proto3" json:"thread_id,omitempty"`
	Index    uint32 `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
}

func (x *MessageIndex) Reset() {
	*x = MessageIndex{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageIndex) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageIndex) ProtoMessage() {}

func (x *MessageIndex) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageIndex.ProtoReflect.Descriptor instead.
func (*MessageIndex) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{10}
}

func (x *MessageIndex) GetThreadId() string {
	if x != nil {
		return x.ThreadId
	}
	return ""
}

func (x *MessageIndex) GetIndex() uint32 {
	if x != nil {
		return x.Index
	}
	return 0
}

type ThreadIDs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []string `protobuf:"bytes,1,rep,name=ids,proto3" json:"ids,omitempty"`
}

func (x *ThreadIDs) Reset() {
	*x = ThreadIDs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThreadIDs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThreadIDs) ProtoMessage() {}

func (x *ThreadIDs) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThreadIDs.ProtoReflect.Descriptor instead.
func (*ThreadIDs) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{11}
}

func (x *ThreadIDs) GetIds() []string {
	if x != nil {
		return x.Ids
	}
	return nil
}

type Source_Fragment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Content  string  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Position uint32  `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
	Score    float32 `protobuf:"fixed32,4,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *Source_Fragment) Reset() {
	*x = Source_Fragment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_service_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Source_Fragment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Source_Fragment) ProtoMessage() {}

func (x *Source_Fragment) ProtoReflect() protoreflect.Message {
	mi := &file_chat_service_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Source_Fragment.ProtoReflect.Descriptor instead.
func (*Source_Fragment) Descriptor() ([]byte, []int) {
	return file_chat_service_proto_rawDescGZIP(), []int{6, 0}
}

func (x *Source_Fragment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Source_Fragment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Source_Fragment) GetPosition() uint32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *Source_Fragment) GetScore() float32 {
	if x != nil {
		return x.Score
	}
	return 0
}

var File_chat_service_proto protoreflect.FileDescriptor

var file_chat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0c, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x90, 0x01, 0x0a, 0x11, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x6f, 0x63, 0x75,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64,
	0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f,
	0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70,
	0x74, 0x12, 0x42, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x34, 0x0a, 0x12, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x98, 0x02, 0x0a, 0x06,
	0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d,
	0x70, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74,
	0x12, 0x42, 0x0a, 0x0d, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f,
	0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4e, 0x0a, 0x11, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61,
	0x6c, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x10, 0x72, 0x65, 0x74, 0x72, 0x69, 0x65, 0x76, 0x61, 0x6c, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x7f, 0x0a, 0x0c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x49,
	0x64, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0b, 0x74, 0x65, 0x6d, 0x70, 0x65, 0x72, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x61, 0x78, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x6d, 0x61, 0x78, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x13, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x5f, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x74, 0x6f, 0x70, 0x50, 0x22, 0x68, 0x0a, 0x10, 0x52, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x76, 0x61, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x65, 0x6e,
	0x61, 0x62, 0x6c, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f,
	0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x09, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68,
	0x6f, 0x6c, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x22, 0xe5, 0x01, 0x0a, 0x06, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x6f, 0x63, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3e, 0x0a, 0x09, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x46, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x1a, 0x66, 0x0a, 0x08, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x07, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x07, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x22, 0x4e, 0x0a,
	0x06, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x22, 0x1a, 0x0a,
	0x08, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x41, 0x0a, 0x0c, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x68,
	0x72, 0x65, 0x61, 0x64, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1d, 0x0a, 0x09,
	0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x64, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x69, 0x64, 0x73, 0x32, 0xc1, 0x03, 0x0a, 0x04,
	0x43, 0x68, 0x61, 0x74, 0x12, 0x40, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x70, 0x74, 0x1a, 0x18, 0x2e, 0x63,
	0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3f, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x1a, 0x17,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x12, 0x4a, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x54,
	0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x73, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f,
	0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64,
	0x49, 0x44, 0x73, 0x12, 0x41, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x68, 0x72,
	0x65, 0x61, 0x64, 0x12, 0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x72, 0x65, 0x61, 0x64, 0x49, 0x44, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x54, 0x68, 0x72, 0x65, 0x61,
	0x64, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x55, 0x0a, 0x0a, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x62, 0x6f, 0x74,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x62, 0x6f, 0x74, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_chat_service_proto_rawDescOnce sync.Once
	file_chat_service_proto_rawDescData = file_chat_service_proto_rawDesc
)

func file_chat_service_proto_rawDescGZIP() []byte {
	file_chat_service_proto_rawDescOnce.Do(func() {
		file_chat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_service_proto_rawDescData)
	})
	return file_chat_service_proto_rawDescData
}

var file_chat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_chat_service_proto_goTypes = []any{
	(*CollectionId)(nil),       // 0: chatbot.chat.v1.CollectionId
	(*CompletionRequest)(nil),  // 1: chatbot.chat.v1.CompletionRequest
	(*CompletionResponse)(nil), // 2: chatbot.chat.v1.CompletionResponse
	(*Prompt)(nil),             // 3: chatbot.chat.v1.Prompt
	(*ModelOptions)(nil),       // 4: chatbot.chat.v1.ModelOptions
	(*RetrievalOptions)(nil),   // 5: chatbot.chat.v1.RetrievalOptions
	(*Source)(nil),             // 6: chatbot.chat.v1.Source
	(*Message)(nil),            // 7: chatbot.chat.v1.Message
	(*Thread)(nil),             // 8: chatbot.chat.v1.Thread
	(*ThreadID)(nil),           // 9: chatbot.chat.v1.ThreadID
	(*MessageIndex)(nil),       // 10: chatbot.chat.v1.MessageIndex
	(*ThreadIDs)(nil),          // 11: chatbot.chat.v1.ThreadIDs
	(*Source_Fragment)(nil),    // 12: chatbot.chat.v1.Source.Fragment
	(*emptypb.Empty)(nil),      // 13: google.protobuf.Empty
}
var file_chat_service_proto_depIdxs = []int32{
	4,  // 0: chatbot.chat.v1.CompletionRequest.model_options:type_name -> chatbot.chat.v1.ModelOptions
	4,  // 1: chatbot.chat.v1.Prompt.model_options:type_name -> chatbot.chat.v1.ModelOptions
	5,  // 2: chatbot.chat.v1.Prompt.retrieval_options:type_name -> chatbot.chat.v1.RetrievalOptions
	12, // 3: chatbot.chat.v1.Source.fragments:type_name -> chatbot.chat.v1.Source.Fragment
	6,  // 4: chatbot.chat.v1.Message.sources:type_name -> chatbot.chat.v1.Source
	7,  // 5: chatbot.chat.v1.Thread.messages:type_name -> chatbot.chat.v1.Message
	3,  // 6: chatbot.chat.v1.Chat.PostMessage:input_type -> chatbot.chat.v1.Prompt
	9,  // 7: chatbot.chat.v1.Chat.GetThread:input_type -> chatbot.chat.v1.ThreadID
	0,  // 8: chatbot.chat.v1.Chat.ListThreadIDs:input_type -> chatbot.chat.v1.CollectionId
	9,  // 9: chatbot.chat.v1.Chat.DeleteThread:input_type -> chatbot.chat.v1.ThreadID
	10, // 10: chatbot.chat.v1.Chat.DeleteMessageFromThread:input_type -> chatbot.chat.v1.MessageIndex
	1,  // 11: chatbot.chat.v1.Chat.Completion:input_type -> chatbot.chat.v1.CompletionRequest
	7,  // 12: chatbot.chat.v1.Chat.PostMessage:output_type -> chatbot.chat.v1.Message
	8,  // 13: chatbot.chat.v1.Chat.GetThread:output_type -> chatbot.chat.v1.Thread
	11, // 14: chatbot.chat.v1.Chat.ListThreadIDs:output_type -> chatbot.chat.v1.ThreadIDs
	13, // 15: chatbot.chat.v1.Chat.DeleteThread:output_type -> google.protobuf.Empty
	13, // 16: chatbot.chat.v1.Chat.DeleteMessageFromThread:output_type -> google.protobuf.Empty
	2,  // 17: chatbot.chat.v1.Chat.Completion:output_type -> chatbot.chat.v1.CompletionResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chat_service_proto_init() }
func file_chat_service_proto_init() {
	if File_chat_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CollectionId); i {
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
		file_chat_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CompletionRequest); i {
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
		file_chat_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CompletionResponse); i {
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
		file_chat_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*Prompt); i {
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
		file_chat_service_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ModelOptions); i {
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
		file_chat_service_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*RetrievalOptions); i {
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
		file_chat_service_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Source); i {
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
		file_chat_service_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Message); i {
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
		file_chat_service_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Thread); i {
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
		file_chat_service_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ThreadID); i {
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
		file_chat_service_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*MessageIndex); i {
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
		file_chat_service_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*ThreadIDs); i {
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
		file_chat_service_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*Source_Fragment); i {
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
			RawDescriptor: file_chat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_service_proto_goTypes,
		DependencyIndexes: file_chat_service_proto_depIdxs,
		MessageInfos:      file_chat_service_proto_msgTypes,
	}.Build()
	File_chat_service_proto = out.File
	file_chat_service_proto_rawDesc = nil
	file_chat_service_proto_goTypes = nil
	file_chat_service_proto_depIdxs = nil
}
