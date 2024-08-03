// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: proto/gophkeep.proto

package gophkeep

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

type FileType int32

const (
	FileType_UNDEFINED FileType = 0
	FileType_BINARY    FileType = 1
	FileType_TEXT      FileType = 2
)

// Enum value maps for FileType.
var (
	FileType_name = map[int32]string{
		0: "UNDEFINED",
		1: "BINARY",
		2: "TEXT",
	}
	FileType_value = map[string]int32{
		"UNDEFINED": 0,
		"BINARY":    1,
		"TEXT":      2,
	}
)

func (x FileType) Enum() *FileType {
	p := new(FileType)
	*p = x
	return p
}

func (x FileType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_gophkeep_proto_enumTypes[0].Descriptor()
}

func (FileType) Type() protoreflect.EnumType {
	return &file_proto_gophkeep_proto_enumTypes[0]
}

func (x FileType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileType.Descriptor instead.
func (FileType) EnumDescriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{0}
}

type SaveCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta     *Meta  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SaveCredentialsRequest) Reset() {
	*x = SaveCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveCredentialsRequest) ProtoMessage() {}

func (x *SaveCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveCredentialsRequest.ProtoReflect.Descriptor instead.
func (*SaveCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{0}
}

func (x *SaveCredentialsRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SaveCredentialsRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *SaveCredentialsRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SaveBankCredentialsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta       *Meta  `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	CardNumber string `protobuf:"bytes,2,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Owner      string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	Cvv        int64  `protobuf:"varint,4,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *SaveBankCredentialsRequest) Reset() {
	*x = SaveBankCredentialsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveBankCredentialsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveBankCredentialsRequest) ProtoMessage() {}

func (x *SaveBankCredentialsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveBankCredentialsRequest.ProtoReflect.Descriptor instead.
func (*SaveBankCredentialsRequest) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{1}
}

func (x *SaveBankCredentialsRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SaveBankCredentialsRequest) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *SaveBankCredentialsRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SaveBankCredentialsRequest) GetCvv() int64 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

type SaveFileRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta   *Meta    `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Type   FileType `protobuf:"varint,3,opt,name=type,proto3,enum=proto.FileType" json:"type,omitempty"`
	Chunks int64    `protobuf:"varint,4,opt,name=chunks,proto3" json:"chunks,omitempty"`
}

func (x *SaveFileRequest) Reset() {
	*x = SaveFileRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileRequest) ProtoMessage() {}

func (x *SaveFileRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileRequest.ProtoReflect.Descriptor instead.
func (*SaveFileRequest) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{2}
}

func (x *SaveFileRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *SaveFileRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SaveFileRequest) GetType() FileType {
	if x != nil {
		return x.Type
	}
	return FileType_UNDEFINED
}

func (x *SaveFileRequest) GetChunks() int64 {
	if x != nil {
		return x.Chunks
	}
	return 0
}

type SaveFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionHash []byte `protobuf:"bytes,1,opt,name=session_hash,json=sessionHash,proto3" json:"session_hash,omitempty"`
}

func (x *SaveFileResponse) Reset() {
	*x = SaveFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveFileResponse) ProtoMessage() {}

func (x *SaveFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveFileResponse.ProtoReflect.Descriptor instead.
func (*SaveFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{3}
}

func (x *SaveFileResponse) GetSessionHash() []byte {
	if x != nil {
		return x.SessionHash
	}
	return nil
}

type FileChunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionHash []byte `protobuf:"bytes,1,opt,name=session_hash,json=sessionHash,proto3" json:"session_hash,omitempty"`
	Chunk       []byte `protobuf:"bytes,2,opt,name=chunk,proto3" json:"chunk,omitempty"`
	ChunkNumber int64  `protobuf:"varint,3,opt,name=chunk_number,json=chunkNumber,proto3" json:"chunk_number,omitempty"`
}

func (x *FileChunk) Reset() {
	*x = FileChunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileChunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileChunk) ProtoMessage() {}

func (x *FileChunk) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileChunk.ProtoReflect.Descriptor instead.
func (*FileChunk) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{4}
}

func (x *FileChunk) GetSessionHash() []byte {
	if x != nil {
		return x.SessionHash
	}
	return nil
}

func (x *FileChunk) GetChunk() []byte {
	if x != nil {
		return x.Chunk
	}
	return nil
}

func (x *FileChunk) GetChunkNumber() int64 {
	if x != nil {
		return x.ChunkNumber
	}
	return 0
}

type GetByMetaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta *Meta `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *GetByMetaRequest) Reset() {
	*x = GetByMetaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByMetaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByMetaRequest) ProtoMessage() {}

func (x *GetByMetaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByMetaRequest.ProtoReflect.Descriptor instead.
func (*GetByMetaRequest) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{5}
}

func (x *GetByMetaRequest) GetMeta() *Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

type GetCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetCredentialsResponse) Reset() {
	*x = GetCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCredentialsResponse) ProtoMessage() {}

func (x *GetCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCredentialsResponse.ProtoReflect.Descriptor instead.
func (*GetCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{6}
}

func (x *GetCredentialsResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetCredentialsResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetBankCredentialsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CardNumber string `protobuf:"bytes,1,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	Owner      string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Cvv        int64  `protobuf:"varint,3,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *GetBankCredentialsResponse) Reset() {
	*x = GetBankCredentialsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBankCredentialsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBankCredentialsResponse) ProtoMessage() {}

func (x *GetBankCredentialsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBankCredentialsResponse.ProtoReflect.Descriptor instead.
func (*GetBankCredentialsResponse) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{7}
}

func (x *GetBankCredentialsResponse) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *GetBankCredentialsResponse) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *GetBankCredentialsResponse) GetCvv() int64 {
	if x != nil {
		return x.Cvv
	}
	return 0
}

type GetFileResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetFileResponse) Reset() {
	*x = GetFileResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFileResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFileResponse) ProtoMessage() {}

func (x *GetFileResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFileResponse.ProtoReflect.Descriptor instead.
func (*GetFileResponse) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{8}
}

func (x *GetFileResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type ListAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAllRequest) Reset() {
	*x = ListAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllRequest) ProtoMessage() {}

func (x *ListAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllRequest.ProtoReflect.Descriptor instead.
func (*ListAllRequest) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{9}
}

type ListAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metas []*Meta `protobuf:"bytes,1,rep,name=metas,proto3" json:"metas,omitempty"`
}

func (x *ListAllResponse) Reset() {
	*x = ListAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllResponse) ProtoMessage() {}

func (x *ListAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllResponse.ProtoReflect.Descriptor instead.
func (*ListAllResponse) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{10}
}

func (x *ListAllResponse) GetMetas() []*Meta {
	if x != nil {
		return x.Metas
	}
	return nil
}

type Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Meta string `protobuf:"bytes,1,opt,name=meta,proto3" json:"meta,omitempty"`
}

func (x *Meta) Reset() {
	*x = Meta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Meta) ProtoMessage() {}

func (x *Meta) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Meta.ProtoReflect.Descriptor instead.
func (*Meta) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{11}
}

func (x *Meta) GetMeta() string {
	if x != nil {
		return x.Meta
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gophkeep_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gophkeep_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_proto_gophkeep_proto_rawDescGZIP(), []int{12}
}

func (x *Status) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_proto_gophkeep_proto protoreflect.FileDescriptor

var file_proto_gophkeep_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x71, 0x0a,
	0x16, 0x53, 0x61, 0x76, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0x86, 0x01, 0x0a, 0x1a, 0x53, 0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61,
	0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0x83, 0x01, 0x0a, 0x0f, 0x53, 0x61,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x23, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x22,
	0x35, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22, 0x67, 0x0a, 0x09, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x68, 0x75, 0x6e, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22,
	0x33, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x04,
	0x6d, 0x65, 0x74, 0x61, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x65, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e,
	0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x76, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0x23, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x72, 0x6c, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x34, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x05, 0x6d, 0x65, 0x74, 0x61, 0x73, 0x22, 0x1a, 0x0a, 0x04, 0x4d, 0x65,
	0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x22, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2a, 0x2f, 0x0a, 0x08, 0x46, 0x69,
	0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x49, 0x4e, 0x41, 0x52, 0x59, 0x10,
	0x01, 0x12, 0x08, 0x0a, 0x04, 0x54, 0x45, 0x58, 0x54, 0x10, 0x02, 0x32, 0xaf, 0x04, 0x0a, 0x08,
	0x47, 0x6f, 0x70, 0x68, 0x4b, 0x65, 0x65, 0x70, 0x12, 0x41, 0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x13, 0x53,
	0x61, 0x76, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x12, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x42,
	0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x61,
	0x76, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x10, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x1a, 0x0d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x28,
	0x01, 0x30, 0x01, 0x12, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x61, 0x6e, 0x6b, 0x43, 0x72,
	0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x65, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x07, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x67, 0x6f, 0x70, 0x68, 0x6b, 0x65, 0x65, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_gophkeep_proto_rawDescOnce sync.Once
	file_proto_gophkeep_proto_rawDescData = file_proto_gophkeep_proto_rawDesc
)

func file_proto_gophkeep_proto_rawDescGZIP() []byte {
	file_proto_gophkeep_proto_rawDescOnce.Do(func() {
		file_proto_gophkeep_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gophkeep_proto_rawDescData)
	})
	return file_proto_gophkeep_proto_rawDescData
}

var file_proto_gophkeep_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_gophkeep_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_proto_gophkeep_proto_goTypes = []any{
	(FileType)(0),                      // 0: proto.FileType
	(*SaveCredentialsRequest)(nil),     // 1: proto.SaveCredentialsRequest
	(*SaveBankCredentialsRequest)(nil), // 2: proto.SaveBankCredentialsRequest
	(*SaveFileRequest)(nil),            // 3: proto.SaveFileRequest
	(*SaveFileResponse)(nil),           // 4: proto.SaveFileResponse
	(*FileChunk)(nil),                  // 5: proto.FileChunk
	(*GetByMetaRequest)(nil),           // 6: proto.GetByMetaRequest
	(*GetCredentialsResponse)(nil),     // 7: proto.GetCredentialsResponse
	(*GetBankCredentialsResponse)(nil), // 8: proto.GetBankCredentialsResponse
	(*GetFileResponse)(nil),            // 9: proto.GetFileResponse
	(*ListAllRequest)(nil),             // 10: proto.ListAllRequest
	(*ListAllResponse)(nil),            // 11: proto.ListAllResponse
	(*Meta)(nil),                       // 12: proto.Meta
	(*Status)(nil),                     // 13: proto.Status
}
var file_proto_gophkeep_proto_depIdxs = []int32{
	12, // 0: proto.SaveCredentialsRequest.meta:type_name -> proto.Meta
	12, // 1: proto.SaveBankCredentialsRequest.meta:type_name -> proto.Meta
	12, // 2: proto.SaveFileRequest.meta:type_name -> proto.Meta
	0,  // 3: proto.SaveFileRequest.type:type_name -> proto.FileType
	12, // 4: proto.GetByMetaRequest.meta:type_name -> proto.Meta
	12, // 5: proto.ListAllResponse.metas:type_name -> proto.Meta
	1,  // 6: proto.GophKeep.SaveCredentials:input_type -> proto.SaveCredentialsRequest
	2,  // 7: proto.GophKeep.SaveBankCredentials:input_type -> proto.SaveBankCredentialsRequest
	3,  // 8: proto.GophKeep.SaveFile:input_type -> proto.SaveFileRequest
	5,  // 9: proto.GophKeep.StartSaveFileStream:input_type -> proto.FileChunk
	6,  // 10: proto.GophKeep.GetCredentials:input_type -> proto.GetByMetaRequest
	6,  // 11: proto.GophKeep.GetBankCredentials:input_type -> proto.GetByMetaRequest
	6,  // 12: proto.GophKeep.GetFile:input_type -> proto.GetByMetaRequest
	10, // 13: proto.GophKeep.ListAll:input_type -> proto.ListAllRequest
	13, // 14: proto.GophKeep.SaveCredentials:output_type -> proto.Status
	13, // 15: proto.GophKeep.SaveBankCredentials:output_type -> proto.Status
	4,  // 16: proto.GophKeep.SaveFile:output_type -> proto.SaveFileResponse
	13, // 17: proto.GophKeep.StartSaveFileStream:output_type -> proto.Status
	7,  // 18: proto.GophKeep.GetCredentials:output_type -> proto.GetCredentialsResponse
	8,  // 19: proto.GophKeep.GetBankCredentials:output_type -> proto.GetBankCredentialsResponse
	9,  // 20: proto.GophKeep.GetFile:output_type -> proto.GetFileResponse
	11, // 21: proto.GophKeep.ListAll:output_type -> proto.ListAllResponse
	14, // [14:22] is the sub-list for method output_type
	6,  // [6:14] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_gophkeep_proto_init() }
func file_proto_gophkeep_proto_init() {
	if File_proto_gophkeep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gophkeep_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*SaveCredentialsRequest); i {
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
		file_proto_gophkeep_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*SaveBankCredentialsRequest); i {
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
		file_proto_gophkeep_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SaveFileRequest); i {
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
		file_proto_gophkeep_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SaveFileResponse); i {
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
		file_proto_gophkeep_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*FileChunk); i {
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
		file_proto_gophkeep_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GetByMetaRequest); i {
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
		file_proto_gophkeep_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*GetCredentialsResponse); i {
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
		file_proto_gophkeep_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*GetBankCredentialsResponse); i {
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
		file_proto_gophkeep_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetFileResponse); i {
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
		file_proto_gophkeep_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllRequest); i {
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
		file_proto_gophkeep_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*ListAllResponse); i {
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
		file_proto_gophkeep_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*Meta); i {
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
		file_proto_gophkeep_proto_msgTypes[12].Exporter = func(v any, i int) any {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_proto_gophkeep_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gophkeep_proto_goTypes,
		DependencyIndexes: file_proto_gophkeep_proto_depIdxs,
		EnumInfos:         file_proto_gophkeep_proto_enumTypes,
		MessageInfos:      file_proto_gophkeep_proto_msgTypes,
	}.Build()
	File_proto_gophkeep_proto = out.File
	file_proto_gophkeep_proto_rawDesc = nil
	file_proto_gophkeep_proto_goTypes = nil
	file_proto_gophkeep_proto_depIdxs = nil
}
