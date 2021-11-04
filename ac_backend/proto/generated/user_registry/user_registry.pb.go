// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: user_registry.proto

package user_registry

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

type GetUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetUserRequest) Reset() {
	*x = GetUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRequest) ProtoMessage() {}

func (x *GetUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRequest.ProtoReflect.Descriptor instead.
func (*GetUserRequest) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{0}
}

func (x *GetUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	UserScope int32  `protobuf:"varint,5,opt,name=userScope,proto3" json:"userScope,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetUserResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserResponse) GetUserScope() int32 {
	if x != nil {
		return x.UserScope
	}
	return 0
}

type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName     string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName      string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email         string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password      string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	ResetPassword bool   `protobuf:"varint,5,opt,name=resetPassword,proto3" json:"resetPassword,omitempty"`
	RefreshUrl    string `protobuf:"bytes,6,opt,name=refreshUrl,proto3" json:"refreshUrl,omitempty"`
	ReturnUrl     string `protobuf:"bytes,7,opt,name=returnUrl,proto3" json:"returnUrl,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CreateUserRequest) GetResetPassword() bool {
	if x != nil {
		return x.ResetPassword
	}
	return false
}

func (x *CreateUserRequest) GetRefreshUrl() string {
	if x != nil {
		return x.RefreshUrl
	}
	return ""
}

func (x *CreateUserRequest) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

type CreateUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId                       string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RefreshUrl                   string  `protobuf:"bytes,2,opt,name=refreshUrl,proto3" json:"refreshUrl,omitempty"`
	ReturnUrl                    string  `protobuf:"bytes,3,opt,name=returnUrl,proto3" json:"returnUrl,omitempty"`
	StripeOnboardingUrl          string  `protobuf:"bytes,4,opt,name=stripeOnboardingUrl,proto3" json:"stripeOnboardingUrl,omitempty"`
	StripeOnboardingUrlExpiresAt int64   `protobuf:"varint,5,opt,name=stripeOnboardingUrlExpiresAt,proto3" json:"stripeOnboardingUrlExpiresAt,omitempty"`
	StripeAccountId              string  `protobuf:"bytes,6,opt,name=stripeAccountId,proto3" json:"stripeAccountId,omitempty"`
	StripeAccountCreatedAt       int64   `protobuf:"varint,7,opt,name=stripeAccountCreatedAt,proto3" json:"stripeAccountCreatedAt,omitempty"`
	StripeCustomerId             string  `protobuf:"bytes,8,opt,name=stripeCustomerId,proto3" json:"stripeCustomerId,omitempty"`
	JwtToken                     string  `protobuf:"bytes,9,opt,name=jwtToken,proto3" json:"jwtToken,omitempty"`
	Claims                       *Claims `protobuf:"bytes,10,opt,name=claims,proto3" json:"claims,omitempty"`
}

func (x *CreateUserResponse) Reset() {
	*x = CreateUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserResponse) ProtoMessage() {}

func (x *CreateUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserResponse.ProtoReflect.Descriptor instead.
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{3}
}

func (x *CreateUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateUserResponse) GetRefreshUrl() string {
	if x != nil {
		return x.RefreshUrl
	}
	return ""
}

func (x *CreateUserResponse) GetReturnUrl() string {
	if x != nil {
		return x.ReturnUrl
	}
	return ""
}

func (x *CreateUserResponse) GetStripeOnboardingUrl() string {
	if x != nil {
		return x.StripeOnboardingUrl
	}
	return ""
}

func (x *CreateUserResponse) GetStripeOnboardingUrlExpiresAt() int64 {
	if x != nil {
		return x.StripeOnboardingUrlExpiresAt
	}
	return 0
}

func (x *CreateUserResponse) GetStripeAccountId() string {
	if x != nil {
		return x.StripeAccountId
	}
	return ""
}

func (x *CreateUserResponse) GetStripeAccountCreatedAt() int64 {
	if x != nil {
		return x.StripeAccountCreatedAt
	}
	return 0
}

func (x *CreateUserResponse) GetStripeCustomerId() string {
	if x != nil {
		return x.StripeCustomerId
	}
	return ""
}

func (x *CreateUserResponse) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *CreateUserResponse) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

type LoginUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserRequest) Reset() {
	*x = LoginUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserRequest) ProtoMessage() {}

func (x *LoginUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserRequest.ProtoReflect.Descriptor instead.
func (*LoginUserRequest) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{4}
}

func (x *LoginUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string  `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	JwtToken string  `protobuf:"bytes,2,opt,name=jwtToken,proto3" json:"jwtToken,omitempty"`
	Claims   *Claims `protobuf:"bytes,3,opt,name=claims,proto3" json:"claims,omitempty"`
}

func (x *LoginUserResponse) Reset() {
	*x = LoginUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserResponse) ProtoMessage() {}

func (x *LoginUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserResponse.ProtoReflect.Descriptor instead.
func (*LoginUserResponse) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{5}
}

func (x *LoginUserResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *LoginUserResponse) GetJwtToken() string {
	if x != nil {
		return x.JwtToken
	}
	return ""
}

func (x *LoginUserResponse) GetClaims() *Claims {
	if x != nil {
		return x.Claims
	}
	return nil
}

type Claims struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email     string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserScope int32  `protobuf:"varint,2,opt,name=userScope,proto3" json:"userScope,omitempty"`
	Exp       int64  `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Iat       int64  `protobuf:"varint,4,opt,name=iat,proto3" json:"iat,omitempty"`
}

func (x *Claims) Reset() {
	*x = Claims{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_registry_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Claims) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Claims) ProtoMessage() {}

func (x *Claims) ProtoReflect() protoreflect.Message {
	mi := &file_user_registry_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Claims.ProtoReflect.Descriptor instead.
func (*Claims) Descriptor() ([]byte, []int) {
	return file_user_registry_proto_rawDescGZIP(), []int{6}
}

func (x *Claims) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Claims) GetUserScope() int32 {
	if x != nil {
		return x.UserScope
	}
	return 0
}

func (x *Claims) GetExp() int64 {
	if x != nil {
		return x.Exp
	}
	return 0
}

func (x *Claims) GetIat() int64 {
	if x != nil {
		return x.Iat
	}
	return 0
}

var File_user_registry_proto protoreflect.FileDescriptor

var file_user_registry_proto_rawDesc = []byte{
	0x0a, 0x13, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x22, 0x28, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x22, 0xe3, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x55, 0x72, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0xb2, 0x03, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e,
	0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x74, 0x75, 0x72,
	0x6e, 0x55, 0x72, 0x6c, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x4f, 0x6e,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x13, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x12, 0x42, 0x0a, 0x1c, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72, 0x6c, 0x45, 0x78, 0x70,
	0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x1c, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x72,
	0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x16, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x6a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x22, 0x44, 0x0a, 0x10,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x6f, 0x0a, 0x11, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6a, 0x77, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x06, 0x63,
	0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x52, 0x06, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x73, 0x22, 0x60, 0x0a, 0x06, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x70,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x65, 0x78, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x69, 0x61, 0x74, 0x32, 0xd1, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x4f, 0x0a, 0x1f, 0x69, 0x6f, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x11, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x17, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_user_registry_proto_rawDescOnce sync.Once
	file_user_registry_proto_rawDescData = file_user_registry_proto_rawDesc
)

func file_user_registry_proto_rawDescGZIP() []byte {
	file_user_registry_proto_rawDescOnce.Do(func() {
		file_user_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_registry_proto_rawDescData)
	})
	return file_user_registry_proto_rawDescData
}

var file_user_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_user_registry_proto_goTypes = []interface{}{
	(*GetUserRequest)(nil),     // 0: api.v1.GetUserRequest
	(*GetUserResponse)(nil),    // 1: api.v1.GetUserResponse
	(*CreateUserRequest)(nil),  // 2: api.v1.CreateUserRequest
	(*CreateUserResponse)(nil), // 3: api.v1.CreateUserResponse
	(*LoginUserRequest)(nil),   // 4: api.v1.LoginUserRequest
	(*LoginUserResponse)(nil),  // 5: api.v1.LoginUserResponse
	(*Claims)(nil),             // 6: api.v1.Claims
}
var file_user_registry_proto_depIdxs = []int32{
	6, // 0: api.v1.CreateUserResponse.claims:type_name -> api.v1.Claims
	6, // 1: api.v1.LoginUserResponse.claims:type_name -> api.v1.Claims
	0, // 2: api.v1.UserRegistry.GetUser:input_type -> api.v1.GetUserRequest
	2, // 3: api.v1.UserRegistry.CreateUser:input_type -> api.v1.CreateUserRequest
	4, // 4: api.v1.UserRegistry.LoginUser:input_type -> api.v1.LoginUserRequest
	1, // 5: api.v1.UserRegistry.GetUser:output_type -> api.v1.GetUserResponse
	3, // 6: api.v1.UserRegistry.CreateUser:output_type -> api.v1.CreateUserResponse
	5, // 7: api.v1.UserRegistry.LoginUser:output_type -> api.v1.LoginUserResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_user_registry_proto_init() }
func file_user_registry_proto_init() {
	if File_user_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserRequest); i {
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
		file_user_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
		file_user_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_user_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserResponse); i {
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
		file_user_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserRequest); i {
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
		file_user_registry_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserResponse); i {
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
		file_user_registry_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Claims); i {
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
			RawDescriptor: file_user_registry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_registry_proto_goTypes,
		DependencyIndexes: file_user_registry_proto_depIdxs,
		MessageInfos:      file_user_registry_proto_msgTypes,
	}.Build()
	File_user_registry_proto = out.File
	file_user_registry_proto_rawDesc = nil
	file_user_registry_proto_goTypes = nil
	file_user_registry_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserRegistryClient is the client API for UserRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserRegistryClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error)
}

type userRegistryClient struct {
	cc grpc.ClientConnInterface
}

func NewUserRegistryClient(cc grpc.ClientConnInterface) UserRegistryClient {
	return &userRegistryClient{cc}
}

func (c *userRegistryClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.UserRegistry/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.UserRegistry/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userRegistryClient) LoginUser(ctx context.Context, in *LoginUserRequest, opts ...grpc.CallOption) (*LoginUserResponse, error) {
	out := new(LoginUserResponse)
	err := c.cc.Invoke(ctx, "/api.v1.UserRegistry/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserRegistryServer is the server API for UserRegistry service.
type UserRegistryServer interface {
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error)
}

// UnimplementedUserRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedUserRegistryServer struct {
}

func (*UnimplementedUserRegistryServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (*UnimplementedUserRegistryServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (*UnimplementedUserRegistryServer) LoginUser(context.Context, *LoginUserRequest) (*LoginUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}

func RegisterUserRegistryServer(s *grpc.Server, srv UserRegistryServer) {
	s.RegisterService(&_UserRegistry_serviceDesc, srv)
}

func _UserRegistry_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.UserRegistry/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.UserRegistry/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserRegistry_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserRegistryServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.UserRegistry/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserRegistryServer).LoginUser(ctx, req.(*LoginUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.UserRegistry",
	HandlerType: (*UserRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserRegistry_GetUser_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _UserRegistry_CreateUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserRegistry_LoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_registry.proto",
}
