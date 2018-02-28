// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userlogin.proto

/*
Package userlogin is a generated protocol buffer package.

It is generated from these files:
	userlogin.proto

It has these top-level messages:
	StringMessage
	AppRequest
	AppListResponse
	AppInfo
	OperatorResponse
*/
package userlogin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type StringMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type AppRequest struct {
	Appid string `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
}

func (m *AppRequest) Reset()                    { *m = AppRequest{} }
func (m *AppRequest) String() string            { return proto.CompactTextString(m) }
func (*AppRequest) ProtoMessage()               {}
func (*AppRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AppRequest) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

type AppListResponse struct {
	Applist []*AppInfo `protobuf:"bytes,1,rep,name=applist" json:"applist,omitempty"`
	Code    int32      `protobuf:"varint,2,opt,name=code" json:"code,omitempty"`
	Message string     `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
}

func (m *AppListResponse) Reset()                    { *m = AppListResponse{} }
func (m *AppListResponse) String() string            { return proto.CompactTextString(m) }
func (*AppListResponse) ProtoMessage()               {}
func (*AppListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *AppListResponse) GetApplist() []*AppInfo {
	if m != nil {
		return m.Applist
	}
	return nil
}

func (m *AppListResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *AppListResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type AppInfo struct {
	Appid       string `protobuf:"bytes,1,opt,name=appid" json:"appid,omitempty"`
	Appname     string `protobuf:"bytes,2,opt,name=appname" json:"appname,omitempty"`
	Appdescribe string `protobuf:"bytes,3,opt,name=appdescribe" json:"appdescribe,omitempty"`
}

func (m *AppInfo) Reset()                    { *m = AppInfo{} }
func (m *AppInfo) String() string            { return proto.CompactTextString(m) }
func (*AppInfo) ProtoMessage()               {}
func (*AppInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AppInfo) GetAppid() string {
	if m != nil {
		return m.Appid
	}
	return ""
}

func (m *AppInfo) GetAppname() string {
	if m != nil {
		return m.Appname
	}
	return ""
}

func (m *AppInfo) GetAppdescribe() string {
	if m != nil {
		return m.Appdescribe
	}
	return ""
}

type OperatorResponse struct {
	Code    int32  `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	Result  string `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
}

func (m *OperatorResponse) Reset()                    { *m = OperatorResponse{} }
func (m *OperatorResponse) String() string            { return proto.CompactTextString(m) }
func (*OperatorResponse) ProtoMessage()               {}
func (*OperatorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *OperatorResponse) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *OperatorResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *OperatorResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "platform.login.StringMessage")
	proto.RegisterType((*AppRequest)(nil), "platform.login.AppRequest")
	proto.RegisterType((*AppListResponse)(nil), "platform.login.AppListResponse")
	proto.RegisterType((*AppInfo)(nil), "platform.login.AppInfo")
	proto.RegisterType((*OperatorResponse)(nil), "platform.login.OperatorResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AppList service

type AppListClient interface {
	List(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppListResponse, error)
	Delte(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*OperatorResponse, error)
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type appListClient struct {
	cc *grpc.ClientConn
}

func NewAppListClient(cc *grpc.ClientConn) AppListClient {
	return &appListClient{cc}
}

func (c *appListClient) List(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*AppListResponse, error) {
	out := new(AppListResponse)
	err := grpc.Invoke(ctx, "/platform.login.AppList/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appListClient) Delte(ctx context.Context, in *AppRequest, opts ...grpc.CallOption) (*OperatorResponse, error) {
	out := new(OperatorResponse)
	err := grpc.Invoke(ctx, "/platform.login.AppList/Delte", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appListClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := grpc.Invoke(ctx, "/platform.login.AppList/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AppList service

type AppListServer interface {
	List(context.Context, *AppRequest) (*AppListResponse, error)
	Delte(context.Context, *AppRequest) (*OperatorResponse, error)
	Echo(context.Context, *StringMessage) (*StringMessage, error)
}

func RegisterAppListServer(s *grpc.Server, srv AppListServer) {
	s.RegisterService(&_AppList_serviceDesc, srv)
}

func _AppList_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppListServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.login.AppList/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppListServer).List(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppList_Delte_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppListServer).Delte(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.login.AppList/Delte",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppListServer).Delte(ctx, req.(*AppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppList_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppListServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/platform.login.AppList/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppListServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _AppList_serviceDesc = grpc.ServiceDesc{
	ServiceName: "platform.login.AppList",
	HandlerType: (*AppListServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _AppList_List_Handler,
		},
		{
			MethodName: "Delte",
			Handler:    _AppList_Delte_Handler,
		},
		{
			MethodName: "Echo",
			Handler:    _AppList_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userlogin.proto",
}

func init() { proto.RegisterFile("userlogin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 400 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xdf, 0x6a, 0xd4, 0x40,
	0x14, 0xc6, 0x49, 0xba, 0xdb, 0x65, 0xcf, 0x6a, 0xab, 0xa3, 0xd8, 0x18, 0x14, 0xc3, 0x80, 0xb0,
	0xf4, 0x22, 0x4b, 0xeb, 0x5d, 0xef, 0x56, 0xf4, 0x42, 0x50, 0x84, 0x78, 0x23, 0xfe, 0x01, 0xa7,
	0xc9, 0x69, 0x3a, 0x30, 0x3b, 0x73, 0x9c, 0x99, 0xf4, 0x01, 0x7c, 0x05, 0x5f, 0xc2, 0xf7, 0xf1,
	0x15, 0x7c, 0x10, 0xc9, 0x24, 0x29, 0xdd, 0x6d, 0xe8, 0x55, 0xf2, 0xcd, 0x39, 0xf9, 0xcd, 0x77,
	0xbe, 0x1c, 0x38, 0x6c, 0x1c, 0x5a, 0x65, 0x6a, 0xa9, 0x73, 0xb2, 0xc6, 0x1b, 0x76, 0x40, 0x4a,
	0xf8, 0x0b, 0x63, 0x37, 0x79, 0x38, 0x4d, 0x9f, 0xd5, 0xc6, 0xd4, 0x0a, 0x57, 0x82, 0xe4, 0x4a,
	0x68, 0x6d, 0xbc, 0xf0, 0xd2, 0x68, 0xd7, 0x75, 0xf3, 0x97, 0x70, 0xff, 0x93, 0xb7, 0x52, 0xd7,
	0x1f, 0xd0, 0x39, 0x51, 0x23, 0x7b, 0x0c, 0xd3, 0x2b, 0xa1, 0x1a, 0x4c, 0xa2, 0x2c, 0x5a, 0xce,
	0x8b, 0x4e, 0x70, 0x0e, 0xb0, 0x26, 0x2a, 0xf0, 0x67, 0x83, 0xce, 0xb7, 0x3d, 0x82, 0x48, 0x56,
	0x43, 0x4f, 0x10, 0xdc, 0xc2, 0xe1, 0x9a, 0xe8, 0xbd, 0x74, 0xbe, 0x40, 0x47, 0x46, 0x3b, 0x64,
	0x27, 0x30, 0x13, 0x44, 0x4a, 0x3a, 0x9f, 0x44, 0xd9, 0xde, 0x72, 0x71, 0x7a, 0x94, 0x6f, 0xbb,
	0xcb, 0xd7, 0x44, 0xef, 0xf4, 0x85, 0x29, 0x86, 0x3e, 0xc6, 0x60, 0x52, 0x9a, 0x0a, 0x93, 0x38,
	0x8b, 0x96, 0xd3, 0x22, 0xbc, 0xb3, 0x04, 0x66, 0x9b, 0xce, 0x5e, 0xb2, 0x17, 0x6e, 0x1c, 0x24,
	0xff, 0x0a, 0xb3, 0x9e, 0x30, 0x6e, 0xaa, 0xfd, 0x54, 0x10, 0x69, 0xb1, 0xe9, 0x88, 0xf3, 0x62,
	0x90, 0x2c, 0x83, 0x85, 0x20, 0xaa, 0xd0, 0x95, 0x56, 0x9e, 0x0f, 0xe0, 0x9b, 0x47, 0xfc, 0x33,
	0x3c, 0xf8, 0x48, 0x68, 0x85, 0x37, 0xf6, 0x7a, 0xa2, 0xc1, 0x5e, 0x34, 0x6e, 0x2f, 0xde, 0xb2,
	0xc7, 0x9e, 0xc0, 0xbe, 0x45, 0xd7, 0x28, 0xdf, 0xe3, 0x7b, 0x75, 0xfa, 0x27, 0x0e, 0xbe, 0xdb,
	0xac, 0xd8, 0x37, 0x98, 0x84, 0x67, 0x3a, 0x12, 0x4d, 0x1f, 0x78, 0xfa, 0x62, 0xa4, 0x76, 0x33,
	0x68, 0x7e, 0xf4, 0xeb, 0xef, 0xbf, 0xdf, 0xf1, 0x43, 0x7e, 0x6f, 0x75, 0x75, 0x12, 0xfe, 0x74,
	0x9b, 0xe5, 0x59, 0x74, 0xcc, 0x7e, 0xc0, 0xf4, 0x0d, 0x2a, 0x8f, 0x77, 0xe2, 0xb3, 0xdd, 0xda,
	0xee, 0xd8, 0xfc, 0x69, 0xe0, 0x3f, 0xe2, 0x07, 0x03, 0xbf, 0x42, 0x85, 0x1e, 0xdb, 0x1b, 0xbe,
	0xc3, 0xe4, 0x6d, 0x79, 0x69, 0xd8, 0xf3, 0x5d, 0xc8, 0xd6, 0x5e, 0xa5, 0x77, 0x97, 0x6f, 0x0f,
	0x80, 0xe5, 0xa5, 0x39, 0x8b, 0x8e, 0x5f, 0x2f, 0xbe, 0xcc, 0xaf, 0x37, 0xfc, 0x7c, 0x3f, 0x2c,
	0xed, 0xab, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xd2, 0xec, 0x8e, 0xf5, 0x02, 0x00, 0x00,
}
