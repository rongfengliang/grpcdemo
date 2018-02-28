// Code generated by protoc-gen-go. DO NOT EDIT.
// source: complextype.proto

/*
Package complextype is a generated protocol buffer package.

It is generated from these files:
	complextype.proto

It has these top-level messages:
	User
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

type LoginEntity struct {
	Usertype UserType `protobuf:"varint,1,opt,name=usertype,enum=UserType" json:"usertype,omitempty"`
	Name     string   `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Password string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
}

func (m *LoginEntity) Reset()                    { *m = LoginEntity{} }
func (m *LoginEntity) String() string            { return proto.CompactTextString(m) }
func (*LoginEntity) ProtoMessage()               {}
func (*LoginEntity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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
func (*TokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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
	Platlogin(ctx context.Context, in *LoginEntity, opts ...grpc.CallOption) (*TokenResponse, error)
}

type userLoginClient struct {
	cc *grpc.ClientConn
}

func NewUserLoginClient(cc *grpc.ClientConn) UserLoginClient {
	return &userLoginClient{cc}
}

func (c *userLoginClient) Platlogin(ctx context.Context, in *LoginEntity, opts ...grpc.CallOption) (*TokenResponse, error) {
	out := new(TokenResponse)
	err := grpc.Invoke(ctx, "/UserLogin/platlogin", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserLogin service

type UserLoginServer interface {
	Platlogin(context.Context, *LoginEntity) (*TokenResponse, error)
}

func RegisterUserLoginServer(s *grpc.Server, srv UserLoginServer) {
	s.RegisterService(&_UserLogin_serviceDesc, srv)
}

func _UserLogin_Platlogin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginEntity)
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
		return srv.(UserLoginServer).Platlogin(ctx, req.(*LoginEntity))
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
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x5d, 0xeb, 0xda, 0x30,
	0x14, 0xc6, 0xad, 0xf3, 0xa5, 0x3d, 0xbe, 0xcc, 0x85, 0x5d, 0x14, 0x6f, 0x26, 0x85, 0x81, 0xdb,
	0x45, 0x85, 0x0e, 0x84, 0x5d, 0x16, 0xac, 0x30, 0xa8, 0x5a, 0x62, 0xfd, 0x00, 0xb5, 0x8d, 0x5d,
	0x59, 0x9b, 0x84, 0x24, 0x32, 0xfd, 0x02, 0xfb, 0xdc, 0x23, 0xa9, 0xca, 0x76, 0xf7, 0xbf, 0x3b,
	0xbf, 0x73, 0xfa, 0x3c, 0x7d, 0x4e, 0x0e, 0x7c, 0xc8, 0x59, 0xc3, 0x6b, 0x72, 0x53, 0x77, 0x4e,
	0x7c, 0x2e, 0x98, 0x62, 0xf3, 0xef, 0x65, 0xa5, 0x7e, 0x5e, 0xcf, 0x7e, 0xce, 0x9a, 0x95, 0x60,
	0xb4, 0xbc, 0x10, 0x5a, 0xd6, 0x55, 0x46, 0xcb, 0x55, 0x29, 0x78, 0x5e, 0x90, 0x86, 0xad, 0x78,
	0x9d, 0xa9, 0x0b, 0x13, 0xcd, 0xab, 0x68, 0xa5, 0xde, 0x1f, 0x0b, 0x7a, 0x27, 0x49, 0x04, 0xfa,
	0x0c, 0xf6, 0x55, 0x12, 0xa1, 0x5d, 0x5d, 0x6b, 0x61, 0x2d, 0xa7, 0x81, 0xe3, 0xeb, 0x41, 0x7a,
	0xe7, 0x04, 0xbf, 0x46, 0xe8, 0x13, 0x0c, 0xc8, 0x8d, 0x33, 0xa1, 0xdc, 0x77, 0x0b, 0x6b, 0x39,
	0x0a, 0x86, 0x7e, 0x64, 0x10, 0x3f, 0xda, 0xde, 0x1a, 0x86, 0x21, 0xe7, 0x5a, 0x85, 0x00, 0x06,
	0xa7, 0x08, 0x87, 0x49, 0x32, 0xeb, 0xa0, 0x29, 0xc0, 0x26, 0xda, 0x86, 0xa7, 0x38, 0xd5, 0x6c,
	0xa1, 0xf7, 0x30, 0x4a, 0xe2, 0x30, 0xdd, 0x1e, 0xf0, 0x4e, 0x37, 0xba, 0x5e, 0x01, 0xa3, 0x98,
	0x95, 0x15, 0x8d, 0xa8, 0xaa, 0xd4, 0xfd, 0xad, 0x71, 0x10, 0xf4, 0x68, 0xd6, 0x10, 0xb7, 0xbb,
	0xb0, 0x96, 0x0e, 0x36, 0x35, 0x9a, 0x83, 0xcd, 0x33, 0x29, 0x7f, 0x33, 0x51, 0x98, 0x90, 0x0e,
	0x7e, 0xb1, 0x77, 0x84, 0x49, 0xca, 0x7e, 0x11, 0x8a, 0x89, 0xe4, 0x8c, 0x4a, 0x63, 0x90, 0xb3,
	0xa2, 0xfd, 0x47, 0x1f, 0x9b, 0x1a, 0x7d, 0x84, 0xbe, 0xd2, 0x1f, 0x3d, 0x5c, 0x5b, 0x40, 0x2e,
	0x0c, 0x1b, 0x22, 0x65, 0x56, 0x92, 0x87, 0xeb, 0x13, 0xbf, 0x06, 0x60, 0x3f, 0xa3, 0x21, 0x07,
	0xfa, 0xe1, 0x66, 0xf7, 0x63, 0x3f, 0xeb, 0xa0, 0x31, 0xd8, 0x49, 0x84, 0x8f, 0x87, 0x7d, 0x18,
	0xcf, 0x2c, 0x34, 0x01, 0x47, 0x2f, 0xdc, 0x0e, 0xbb, 0xc1, 0x1a, 0x1c, 0xad, 0x31, 0x2b, 0xa3,
	0x2f, 0xe0, 0xe8, 0xb3, 0xd4, 0x06, 0xc6, 0xfe, 0x3f, 0xef, 0x30, 0x9f, 0xfa, 0xff, 0xe5, 0x3d,
	0x0f, 0xcc, 0xd9, 0xbe, 0xfd, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xf0, 0xe2, 0x6f, 0x13, 0x06, 0x02,
	0x00, 0x00,
}
