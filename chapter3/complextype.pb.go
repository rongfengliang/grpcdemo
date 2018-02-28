// Code generated by protoc-gen-go. DO NOT EDIT.
// source: complextype.proto

/*
Package complextype is a generated protocol buffer package.

It is generated from these files:
	complextype.proto

It has these top-level messages:
	User
	LoginEntityUser
	LoginEntity
	TokenResponse
*/
package complextype

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import platform "github.com/rongfengliang/grpcdemo/platform"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserType int32

const (
	UserType_ADMIN     UserType = 0
	UserType_PERSONAL  UserType = 1
	UserType_PLATADMIN UserType = 2
)

var UserType_name = map[int32]string{
	0: "ADMIN",
	1: "PERSONAL",
	2: "PLATADMIN",
}
var UserType_value = map[string]int32{
	"ADMIN":     0,
	"PERSONAL":  1,
	"PLATADMIN": 2,
}

func (x UserType) String() string {
	return proto.EnumName(UserType_name, int32(x))
}
func (UserType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type User_AppType int32

const (
	User_UERAPP      User_AppType = 0
	User_DEFAULTAPP  User_AppType = 1
	User_PLATFORMAPP User_AppType = 2
)

var User_AppType_name = map[int32]string{
	0: "UERAPP",
	1: "DEFAULTAPP",
	2: "PLATFORMAPP",
}
var User_AppType_value = map[string]int32{
	"UERAPP":      0,
	"DEFAULTAPP":  1,
	"PLATFORMAPP": 2,
}

func (x User_AppType) String() string {
	return proto.EnumName(User_AppType_name, int32(x))
}
func (User_AppType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type User struct {
	Usertype UserType         `protobuf:"varint,1,opt,name=usertype,enum=UserType" json:"usertype,omitempty"`
	Export   *platform.Export `protobuf:"bytes,3,opt,name=export" json:"export,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *User) GetUsertype() UserType {
	if m != nil {
		return m.Usertype
	}
	return UserType_ADMIN
}

func (m *User) GetExport() *platform.Export {
	if m != nil {
		return m.Export
	}
	return nil
}

type LoginEntityUser struct {
	Loginentity *LoginEntity `protobuf:"bytes,1,opt,name=loginentity" json:"loginentity,omitempty"`
	User        *User        `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
}

func (m *LoginEntityUser) Reset()                    { *m = LoginEntityUser{} }
func (m *LoginEntityUser) String() string            { return proto.CompactTextString(m) }
func (*LoginEntityUser) ProtoMessage()               {}
func (*LoginEntityUser) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *LoginEntityUser) GetLoginentity() *LoginEntity {
	if m != nil {
		return m.Loginentity
	}
	return nil
}

func (m *LoginEntityUser) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type LoginEntity struct {
	Usertype UserType `protobuf:"varint,1,opt,name=usertype,enum=UserType" json:"usertype,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginEntity) Reset()                    { *m = LoginEntity{} }
func (m *LoginEntity) String() string            { return proto.CompactTextString(m) }
func (*LoginEntity) ProtoMessage()               {}
func (*LoginEntity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *LoginEntity) GetUsertype() UserType {
	if m != nil {
		return m.Usertype
	}
	return UserType_ADMIN
}

func (m *LoginEntity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginEntity) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type TokenResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *TokenResponse) Reset()                    { *m = TokenResponse{} }
func (m *TokenResponse) String() string            { return proto.CompactTextString(m) }
func (*TokenResponse) ProtoMessage()               {}
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *TokenResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *TokenResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "User")
	proto.RegisterType((*LoginEntityUser)(nil), "LoginEntityUser")
	proto.RegisterType((*LoginEntity)(nil), "LoginEntity")
	proto.RegisterType((*TokenResponse)(nil), "TokenResponse")
	proto.RegisterEnum("UserType", UserType_name, UserType_value)
	proto.RegisterEnum("User_AppType", User_AppType_name, User_AppType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserLogin service

type UserLoginClient interface {
	Platlogin(ctx context.Context, in *LoginEntityUser, opts ...grpc.CallOption) (*TokenResponse, error)
}

type userLoginClient struct {
	cc *grpc.ClientConn
}

func NewUserLoginClient(cc *grpc.ClientConn) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) Platlogin(ctx context.Context, in *LoginEntityUser, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/UserLogin/platlogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserLogin service

type UserLoginServer interface {
	Platlogin(context.Context, *LoginEntityUser) (*TokenResponse, error)
}

func RegisterUserLoginServer(s *grpc.Server, srv UserLoginServer) {
	s.RegisterService(&_UserLogin_serviceDesc, srv)
}

func _UserLogin_Platlogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginEntityUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserLoginServer).Platlogin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserLogin/Platlogin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserLoginServer).Platlogin(ctx, req.(*LoginEntityUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserLogin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserLogin",
	HandlerType: (*UserLoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "platlogin",
			Handler:    _UserLogin_Platlogin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "complextype.proto",
}

func init() { proto.RegisterFile("complextype.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xae, 0xb3, 0x38, 0xb1, 0x9f, 0xdb, 0xd4, 0x13, 0x3b, 0x78, 0xb9, 0xac, 0x18, 0x06, 0x65,
	0x30, 0x05, 0x3c, 0x18, 0x6c, 0x37, 0x43, 0x5d, 0x18, 0xb8, 0xad, 0x51, 0x9d, 0xdb, 0x2e, 0xae,
	0xad, 0x78, 0x66, 0xb6, 0x24, 0x24, 0x85, 0x25, 0x7f, 0x60, 0xbf, 0x7b, 0x48, 0x4e, 0x42, 0xb6,
	0xd3, 0x6e, 0xef, 0xfb, 0xde, 0x7b, 0x9f, 0xbf, 0xf7, 0x59, 0xf0, 0xba, 0xe6, 0x83, 0xe8, 0xe9,
	0x4e, 0xef, 0x05, 0xc5, 0x42, 0x72, 0xcd, 0x97, 0x5f, 0xda, 0x4e, 0xff, 0xd8, 0xbe, 0xe0, 0x9a,
	0x0f, 0x2b, 0xc9, 0x59, 0xbb, 0xa1, 0xac, 0xed, 0xbb, 0x8a, 0xb5, 0xab, 0x56, 0x8a, 0xba, 0xa1,
	0x03, 0x5f, 0x89, 0xbe, 0xd2, 0x1b, 0x2e, 0x87, 0x53, 0x31, 0xae, 0xc6, 0xbf, 0x1d, 0x98, 0xae,
	0x15, 0x95, 0xe8, 0x3d, 0x78, 0x5b, 0x45, 0xa5, 0x51, 0x8d, 0x9c, 0x1b, 0xe7, 0x76, 0x91, 0xf8,
	0xd8, 0x34, 0xca, 0xbd, 0xa0, 0xe4, 0xd4, 0x42, 0xef, 0x60, 0x46, 0x77, 0x82, 0x4b, 0x1d, 0xbd,
	0xba, 0x71, 0x6e, 0x83, 0x64, 0x8e, 0x33, 0x0b, 0xc9, 0x81, 0x8e, 0x3f, 0xc3, 0x3c, 0x15, 0xc2,
	0x6c, 0x21, 0x80, 0xd9, 0x3a, 0x23, 0x69, 0x51, 0x84, 0x17, 0x68, 0x01, 0x70, 0x97, 0xdd, 0xa7,
	0xeb, 0xbc, 0x34, 0xd8, 0x41, 0xd7, 0x10, 0x14, 0x79, 0x5a, 0xde, 0x3f, 0x91, 0x07, 0x43, 0x4c,
	0xe2, 0xef, 0x70, 0x9d, 0xf3, 0xb6, 0x63, 0x19, 0xd3, 0x9d, 0xde, 0x5b, 0x4b, 0x18, 0x82, 0xde,
	0x50, 0xd4, 0x52, 0xd6, 0x55, 0x90, 0x5c, 0xe2, 0xb3, 0x31, 0x72, 0x3e, 0x80, 0xde, 0xc2, 0xd4,
	0xf8, 0x8c, 0x26, 0x76, 0xd0, 0xb5, 0xf6, 0x89, 0xa5, 0xe2, 0x06, 0x82, 0xb3, 0xb5, 0xff, 0x3d,
	0x16, 0xc1, 0x94, 0x55, 0x03, 0xb5, 0x82, 0x3e, 0xb1, 0x35, 0x5a, 0x82, 0x27, 0x2a, 0xa5, 0x7e,
	0x71, 0xd9, 0xd8, 0x08, 0x7c, 0x72, 0xc2, 0xf1, 0x33, 0x5c, 0x95, 0xfc, 0x27, 0x65, 0x84, 0x2a,
	0xc1, 0x99, 0xb2, 0x02, 0x35, 0x6f, 0xc6, 0x6f, 0xb8, 0xc4, 0xd6, 0xe8, 0x0d, 0xb8, 0xda, 0x0c,
	0x1d, 0x54, 0x47, 0x80, 0x22, 0x98, 0x0f, 0x54, 0xa9, 0xaa, 0xa5, 0x07, 0xd5, 0x23, 0xfc, 0x90,
	0x80, 0x77, 0xb4, 0x86, 0x7c, 0x70, 0xd3, 0xbb, 0x87, 0x6f, 0x8f, 0xe1, 0x05, 0xba, 0x04, 0xaf,
	0xc8, 0xc8, 0xf3, 0xd3, 0x63, 0x9a, 0x87, 0x0e, 0xba, 0x02, 0xdf, 0xc4, 0x39, 0x36, 0x27, 0xc9,
	0x57, 0xf0, 0xcd, 0x8e, 0x3d, 0x19, 0x7d, 0x04, 0xdf, 0xfc, 0x74, 0x9b, 0x14, 0x0a, 0xf1, 0x3f,
	0x29, 0x2f, 0x17, 0xf8, 0x2f, 0xcf, 0x2f, 0x33, 0xfb, 0x30, 0x3e, 0xfd, 0x09, 0x00, 0x00, 0xff,
	0xff, 0xfc, 0x26, 0xb1, 0xaf, 0x68, 0x02, 0x00, 0x00,
}
