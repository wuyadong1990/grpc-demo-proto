// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: user/user.proto

package user

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

//  用户性别
type UserSex int32

const (
	UserSex_DEFAULT UserSex = 0
	UserSex_MEN     UserSex = 1
	UserSex_WOMEN   UserSex = 2
)

// Enum value maps for UserSex.
var (
	UserSex_name = map[int32]string{
		0: "DEFAULT",
		1: "MEN",
		2: "WOMEN",
	}
	UserSex_value = map[string]int32{
		"DEFAULT": 0,
		"MEN":     1,
		"WOMEN":   2,
	}
)

func (x UserSex) Enum() *UserSex {
	p := new(UserSex)
	*p = x
	return p
}

func (x UserSex) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserSex) Descriptor() protoreflect.EnumDescriptor {
	return file_user_user_proto_enumTypes[0].Descriptor()
}

func (UserSex) Type() protoreflect.EnumType {
	return &file_user_user_proto_enumTypes[0]
}

func (x UserSex) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserSex.Descriptor instead.
func (UserSex) EnumDescriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

//  注入标签: https://github.com/favadi/protoc-go-inject-tag
//  用户基本信息
type UserBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  @inject_tag: db:"id"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	//  @inject_tag: db:"user_name" valid:"required~用户名称必须存在"
	UserName string `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	//  @inject_tag: db:"iphone" valid:"required~手机号必须存在"
	Iphone string `protobuf:"bytes,3,opt,name=iphone,proto3" json:"iphone,omitempty"`
	//  @inject_tag: db:"password"
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	//  @inject_tag: db:"sex"
	Sex UserSex `protobuf:"varint,5,opt,name=sex,proto3,enum=user.UserSex" json:"sex,omitempty"`
}

func (x *UserBase) Reset() {
	*x = UserBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserBase) ProtoMessage() {}

func (x *UserBase) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserBase.ProtoReflect.Descriptor instead.
func (*UserBase) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserBase) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserBase) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UserBase) GetIphone() string {
	if x != nil {
		return x.Iphone
	}
	return ""
}

func (x *UserBase) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserBase) GetSex() UserSex {
	if x != nil {
		return x.Sex
	}
	return UserSex_DEFAULT
}

type UserID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  用户id
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserID) Reset() {
	*x = UserID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserID) ProtoMessage() {}

func (x *UserID) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserID.ProtoReflect.Descriptor instead.
func (*UserID) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserID) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

//  用户全部选项
type UserAllOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  用户性别
	Sex UserSex `protobuf:"varint,1,opt,name=sex,proto3,enum=user.UserSex" json:"sex,omitempty"`
	//  页数
	Page *Page `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	//  用户名称
	UserName string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
}

func (x *UserAllOption) Reset() {
	*x = UserAllOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAllOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAllOption) ProtoMessage() {}

func (x *UserAllOption) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAllOption.ProtoReflect.Descriptor instead.
func (*UserAllOption) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserAllOption) GetSex() UserSex {
	if x != nil {
		return x.Sex
	}
	return UserSex_DEFAULT
}

func (x *UserAllOption) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *UserAllOption) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

//  用户全部
type UserAll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  用户信息
	All []*UserBase `protobuf:"bytes,1,rep,name=all,proto3" json:"all,omitempty"`
	//  页数
	Page *Page `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *UserAll) Reset() {
	*x = UserAll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserAll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserAll) ProtoMessage() {}

func (x *UserAll) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserAll.ProtoReflect.Descriptor instead.
func (*UserAll) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserAll) GetAll() []*UserBase {
	if x != nil {
		return x.All
	}
	return nil
}

func (x *UserAll) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

//  空消息
type Null struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Null) Reset() {
	*x = Null{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Null) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Null) ProtoMessage() {}

func (x *Null) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Null.ProtoReflect.Descriptor instead.
func (*Null) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{4}
}

//  分页
type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//  页
	PageIndex int64 `protobuf:"varint,1,opt,name=page_index,json=pageIndex,proto3" json:"page_index,omitempty"`
	//  每页大小
	PageSize int64 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	//  总页数
	PageTotal int64 `protobuf:"varint,3,opt,name=page_total,json=pageTotal,proto3" json:"page_total,omitempty"`
	//  条数
	Count int64 `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
	//  总条数
	Total int64 `protobuf:"varint,5,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *Page) GetPageIndex() int64 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *Page) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *Page) GetPageTotal() int64 {
	if x != nil {
		return x.PageTotal
	}
	return 0
}

func (x *Page) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Page) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_user_user_proto protoreflect.FileDescriptor

var file_user_user_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8c, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x69, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x69, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x1f, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x52,
	0x03, 0x73, 0x65, 0x78, 0x22, 0x18, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x6d,
	0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x52, 0x03, 0x73, 0x65, 0x78,
	0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x4b, 0x0a,
	0x07, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x12, 0x20, 0x0a, 0x03, 0x61, 0x6c, 0x6c, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x61, 0x73, 0x65, 0x52, 0x03, 0x61, 0x6c, 0x6c, 0x12, 0x1e, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x06, 0x0a, 0x04, 0x4e, 0x75,
	0x6c, 0x6c, 0x22, 0x8d, 0x01, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70,
	0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x2a, 0x2a, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x78, 0x12, 0x0b, 0x0a,
	0x07, 0x44, 0x45, 0x46, 0x41, 0x55, 0x4c, 0x54, 0x10, 0x00, 0x12, 0x07, 0x0a, 0x03, 0x4d, 0x45,
	0x4e, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x57, 0x4f, 0x4d, 0x45, 0x4e, 0x10, 0x02, 0x32, 0x94,
	0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x0a, 0x22, 0x05, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x01, 0x2a, 0x12, 0x3c, 0x0a, 0x0a,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x2a, 0x0a,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x43, 0x0a, 0x0c, 0x55, 0x73,
	0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4f, 0x6e, 0x65, 0x12, 0x0c, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x1a, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x61, 0x73, 0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f,
	0x12, 0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62, 0x01, 0x2a, 0x12,
	0x44, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x72, 0x79, 0x41, 0x6c, 0x6c, 0x12,
	0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x41, 0x6c, 0x6c, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x6c, 0x6c, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x05, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x62, 0x01, 0x2a, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x75, 0x79, 0x61, 0x64, 0x6f, 0x6e, 0x67, 0x31, 0x39, 0x39, 0x30,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x65, 0x6d, 0x6f, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_user_proto_rawDescOnce sync.Once
	file_user_user_proto_rawDescData = file_user_user_proto_rawDesc
)

func file_user_user_proto_rawDescGZIP() []byte {
	file_user_user_proto_rawDescOnce.Do(func() {
		file_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_user_proto_rawDescData)
	})
	return file_user_user_proto_rawDescData
}

var file_user_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_user_user_proto_goTypes = []interface{}{
	(UserSex)(0),          // 0: user.UserSex
	(*UserBase)(nil),      // 1: user.UserBase
	(*UserID)(nil),        // 2: user.UserID
	(*UserAllOption)(nil), // 3: user.UserAllOption
	(*UserAll)(nil),       // 4: user.UserAll
	(*Null)(nil),          // 5: user.Null
	(*Page)(nil),          // 6: user.Page
}
var file_user_user_proto_depIdxs = []int32{
	0, // 0: user.UserBase.sex:type_name -> user.UserSex
	0, // 1: user.UserAllOption.sex:type_name -> user.UserSex
	6, // 2: user.UserAllOption.page:type_name -> user.Page
	1, // 3: user.UserAll.all:type_name -> user.UserBase
	6, // 4: user.UserAll.page:type_name -> user.Page
	1, // 5: user.UserService.UserInfo:input_type -> user.UserBase
	2, // 6: user.UserService.UserDelete:input_type -> user.UserID
	2, // 7: user.UserService.UserQueryOne:input_type -> user.UserID
	3, // 8: user.UserService.UserQueryAll:input_type -> user.UserAllOption
	1, // 9: user.UserService.UserInfo:output_type -> user.UserBase
	2, // 10: user.UserService.UserDelete:output_type -> user.UserID
	1, // 11: user.UserService.UserQueryOne:output_type -> user.UserBase
	4, // 12: user.UserService.UserQueryAll:output_type -> user.UserAll
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_user_user_proto_init() }
func file_user_user_proto_init() {
	if File_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserBase); i {
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
		file_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserID); i {
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
		file_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAllOption); i {
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
		file_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserAll); i {
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
		file_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Null); i {
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
		file_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_user_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_user_proto_goTypes,
		DependencyIndexes: file_user_user_proto_depIdxs,
		EnumInfos:         file_user_user_proto_enumTypes,
		MessageInfos:      file_user_user_proto_msgTypes,
	}.Build()
	File_user_user_proto = out.File
	file_user_user_proto_rawDesc = nil
	file_user_user_proto_goTypes = nil
	file_user_user_proto_depIdxs = nil
}
