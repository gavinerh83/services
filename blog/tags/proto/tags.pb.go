// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/micro/services/blog/tags/proto/tags.proto

package tags

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

type Tag struct {
	// Type is useful for namespacing and listing across resources,
	// ie. list tags for posts, customers etc.
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Slug                 string   `protobuf:"bytes,2,opt,name=slug,proto3" json:"slug,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Count                int64    `protobuf:"varint,5,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Tag) GetSlug() string {
	if m != nil {
		return m.Slug
	}
	return ""
}

func (m *Tag) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Tag) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Tag) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type AddRequest struct {
	ResourceID           string   `protobuf:"bytes,1,opt,name=resourceID,proto3" json:"resourceID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	ResourceCreated      int64    `protobuf:"varint,4,opt,name=resourceCreated,proto3" json:"resourceCreated,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddRequest) Reset()         { *m = AddRequest{} }
func (m *AddRequest) String() string { return proto.CompactTextString(m) }
func (*AddRequest) ProtoMessage()    {}
func (*AddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{1}
}

func (m *AddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddRequest.Unmarshal(m, b)
}
func (m *AddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddRequest.Marshal(b, m, deterministic)
}
func (m *AddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddRequest.Merge(m, src)
}
func (m *AddRequest) XXX_Size() int {
	return xxx_messageInfo_AddRequest.Size(m)
}
func (m *AddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddRequest proto.InternalMessageInfo

func (m *AddRequest) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

func (m *AddRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *AddRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AddRequest) GetResourceCreated() int64 {
	if m != nil {
		return m.ResourceCreated
	}
	return 0
}

type AddResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddResponse) Reset()         { *m = AddResponse{} }
func (m *AddResponse) String() string { return proto.CompactTextString(m) }
func (*AddResponse) ProtoMessage()    {}
func (*AddResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{2}
}

func (m *AddResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddResponse.Unmarshal(m, b)
}
func (m *AddResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddResponse.Marshal(b, m, deterministic)
}
func (m *AddResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddResponse.Merge(m, src)
}
func (m *AddResponse) XXX_Size() int {
	return xxx_messageInfo_AddResponse.Size(m)
}
func (m *AddResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddResponse proto.InternalMessageInfo

type RemoveRequest struct {
	ResourceID           string   `protobuf:"bytes,1,opt,name=resourceID,proto3" json:"resourceID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{3}
}

func (m *RemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveRequest.Unmarshal(m, b)
}
func (m *RemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveRequest.Marshal(b, m, deterministic)
}
func (m *RemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveRequest.Merge(m, src)
}
func (m *RemoveRequest) XXX_Size() int {
	return xxx_messageInfo_RemoveRequest.Size(m)
}
func (m *RemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveRequest proto.InternalMessageInfo

func (m *RemoveRequest) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

func (m *RemoveRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *RemoveRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type RemoveResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RemoveResponse) Reset()         { *m = RemoveResponse{} }
func (m *RemoveResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveResponse) ProtoMessage()    {}
func (*RemoveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{4}
}

func (m *RemoveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RemoveResponse.Unmarshal(m, b)
}
func (m *RemoveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RemoveResponse.Marshal(b, m, deterministic)
}
func (m *RemoveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveResponse.Merge(m, src)
}
func (m *RemoveResponse) XXX_Size() int {
	return xxx_messageInfo_RemoveResponse.Size(m)
}
func (m *RemoveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveResponse proto.InternalMessageInfo

type UpdateRequest struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{5}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *UpdateRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UpdateRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{6}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

// ListRequest: list either by resource id or type.
// Optionally filter by min or max count.
type ListRequest struct {
	ResourceID           string   `protobuf:"bytes,1,opt,name=resourceID,proto3" json:"resourceID,omitempty"`
	Type                 string   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	MinCount             int64    `protobuf:"varint,3,opt,name=minCount,proto3" json:"minCount,omitempty"`
	MaxCount             int64    `protobuf:"varint,4,opt,name=maxCount,proto3" json:"maxCount,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{7}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

func (m *ListRequest) GetResourceID() string {
	if m != nil {
		return m.ResourceID
	}
	return ""
}

func (m *ListRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ListRequest) GetMinCount() int64 {
	if m != nil {
		return m.MinCount
	}
	return 0
}

func (m *ListRequest) GetMaxCount() int64 {
	if m != nil {
		return m.MaxCount
	}
	return 0
}

type ListResponse struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fc76523216c4be8e, []int{8}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Tag)(nil), "tags.Tag")
	proto.RegisterType((*AddRequest)(nil), "tags.AddRequest")
	proto.RegisterType((*AddResponse)(nil), "tags.AddResponse")
	proto.RegisterType((*RemoveRequest)(nil), "tags.RemoveRequest")
	proto.RegisterType((*RemoveResponse)(nil), "tags.RemoveResponse")
	proto.RegisterType((*UpdateRequest)(nil), "tags.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "tags.UpdateResponse")
	proto.RegisterType((*ListRequest)(nil), "tags.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "tags.ListResponse")
}

func init() {
	proto.RegisterFile("github.com/micro/services/blog/tags/proto/tags.proto", fileDescriptor_fc76523216c4be8e)
}

var fileDescriptor_fc76523216c4be8e = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x41, 0x4f, 0xab, 0x40,
	0x10, 0x2e, 0x5d, 0xda, 0xbc, 0x0e, 0xaf, 0xef, 0xf5, 0xed, 0xeb, 0x81, 0x90, 0x68, 0x08, 0x27,
	0x0e, 0x5a, 0x92, 0xaa, 0x3f, 0xa0, 0xa9, 0x17, 0x13, 0x4f, 0xa4, 0x1e, 0x8c, 0x27, 0x0a, 0x1b,
	0xdc, 0xa4, 0xb0, 0xc8, 0x2e, 0x8d, 0x8d, 0x17, 0xff, 0xa3, 0x7f, 0xc8, 0xb0, 0x0b, 0xb8, 0xad,
	0x8d, 0x17, 0xbd, 0xcd, 0x7c, 0xc3, 0xcc, 0xf7, 0xed, 0x37, 0x03, 0x5c, 0xa6, 0x54, 0x3c, 0x56,
	0xeb, 0x59, 0xcc, 0xb2, 0x20, 0xa3, 0x71, 0xc9, 0x02, 0x4e, 0xca, 0x2d, 0x8d, 0x09, 0x0f, 0xd6,
	0x1b, 0x96, 0x06, 0x22, 0x4a, 0x79, 0x50, 0x94, 0x4c, 0x30, 0x19, 0xce, 0x64, 0x88, 0xcd, 0x3a,
	0xf6, 0x5e, 0x00, 0xad, 0xa2, 0x14, 0x63, 0x30, 0xc5, 0xae, 0x20, 0xb6, 0xe1, 0x1a, 0xfe, 0x28,
	0x94, 0x71, 0x8d, 0xf1, 0x4d, 0x95, 0xda, 0x7d, 0x85, 0xd5, 0x31, 0x9e, 0xc2, 0x40, 0x50, 0xb1,
	0x21, 0x36, 0x92, 0xa0, 0x4a, 0xb0, 0x0b, 0x56, 0x42, 0x78, 0x5c, 0xd2, 0x42, 0x50, 0x96, 0xdb,
	0xa6, 0xac, 0xe9, 0x50, 0xdd, 0x17, 0xb3, 0x2a, 0x17, 0xf6, 0xc0, 0x35, 0x7c, 0x14, 0xaa, 0xc4,
	0x7b, 0x35, 0x00, 0x16, 0x49, 0x12, 0x92, 0xa7, 0x8a, 0x70, 0x81, 0x4f, 0x01, 0x4a, 0xc2, 0x59,
	0x55, 0xc6, 0xe4, 0xe6, 0xba, 0x91, 0xa2, 0x21, 0x9d, 0xc8, 0xbe, 0x26, 0xf2, 0xb8, 0x20, 0x1f,
	0xfe, 0xb6, 0x7d, 0xcb, 0x92, 0x44, 0x82, 0x24, 0x52, 0x14, 0x0a, 0x0f, 0x61, 0x6f, 0x0c, 0x96,
	0x54, 0xc0, 0x0b, 0x96, 0x73, 0xe2, 0xdd, 0xc3, 0x38, 0x24, 0x19, 0xdb, 0x92, 0x1f, 0xd7, 0xe4,
	0x4d, 0xe0, 0x4f, 0x3b, 0xba, 0x21, 0x7b, 0x80, 0xf1, 0x5d, 0x91, 0x44, 0xa2, 0x23, 0x3b, 0xb6,
	0x85, 0x6e, 0x58, 0xff, 0x0b, 0xc7, 0xd1, 0x27, 0xc7, 0x6b, 0xba, 0x76, 0x78, 0x43, 0xb7, 0x03,
	0xeb, 0x96, 0x72, 0xf1, 0x9d, 0x97, 0x39, 0xf0, 0x2b, 0xa3, 0xf9, 0x52, 0x6e, 0x12, 0x49, 0x43,
	0xbb, 0x5c, 0xd6, 0xa2, 0x67, 0x55, 0x33, 0x9b, 0x5a, 0x93, 0x7b, 0xe7, 0xf0, 0x5b, 0x51, 0x2b,
	0x29, 0xf8, 0x04, 0xe4, 0xf5, 0xd9, 0x86, 0x8b, 0x7c, 0x6b, 0x3e, 0x9a, 0xc9, 0xb3, 0x5c, 0x45,
	0x69, 0x28, 0xe1, 0xf9, 0x9b, 0x01, 0xe6, 0x2a, 0x4a, 0x39, 0x3e, 0x03, 0xb4, 0x48, 0x12, 0x3c,
	0x51, 0x1f, 0x7c, 0x9c, 0x8a, 0xf3, 0x4f, 0x43, 0x9a, 0xe7, 0xf5, 0xf0, 0x15, 0x0c, 0x95, 0xc3,
	0xf8, 0xbf, 0x2a, 0xef, 0xad, 0xd2, 0x99, 0xee, 0x83, 0x5d, 0x5b, 0x00, 0x66, 0x2d, 0x0e, 0x37,
	0x33, 0x35, 0x8f, 0x1c, 0xac, 0x43, 0x3a, 0x8f, 0xb2, 0xb6, 0xe5, 0xd9, 0xdb, 0x62, 0xcb, 0x73,
	0xe0, 0x7e, 0x6f, 0x3d, 0x94, 0xff, 0xdd, 0xc5, 0x7b, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xb1,
	0x3f, 0xea, 0xaf, 0x03, 0x00, 0x00,
}
