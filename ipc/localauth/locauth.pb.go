// Code generated by protoc-gen-go. DO NOT EDIT.
// source: locauth.proto

package localauth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e483730694cbd353, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type ATResponse struct {
	AccessToken          string   `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ATResponse) Reset()         { *m = ATResponse{} }
func (m *ATResponse) String() string { return proto.CompactTextString(m) }
func (*ATResponse) ProtoMessage()    {}
func (*ATResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e483730694cbd353, []int{1}
}

func (m *ATResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ATResponse.Unmarshal(m, b)
}
func (m *ATResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ATResponse.Marshal(b, m, deterministic)
}
func (m *ATResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ATResponse.Merge(m, src)
}
func (m *ATResponse) XXX_Size() int {
	return xxx_messageInfo_ATResponse.Size(m)
}
func (m *ATResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ATResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ATResponse proto.InternalMessageInfo

func (m *ATResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

type UCRequest struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UCRequest) Reset()         { *m = UCRequest{} }
func (m *UCRequest) String() string { return proto.CompactTextString(m) }
func (*UCRequest) ProtoMessage()    {}
func (*UCRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e483730694cbd353, []int{2}
}

func (m *UCRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UCRequest.Unmarshal(m, b)
}
func (m *UCRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UCRequest.Marshal(b, m, deterministic)
}
func (m *UCRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UCRequest.Merge(m, src)
}
func (m *UCRequest) XXX_Size() int {
	return xxx_messageInfo_UCRequest.Size(m)
}
func (m *UCRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UCRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UCRequest proto.InternalMessageInfo

func (m *UCRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type PingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e483730694cbd353, []int{3}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "localauth.Empty")
	proto.RegisterType((*ATResponse)(nil), "localauth.ATResponse")
	proto.RegisterType((*UCRequest)(nil), "localauth.UCRequest")
	proto.RegisterType((*PingResponse)(nil), "localauth.PingResponse")
}

func init() { proto.RegisterFile("locauth.proto", fileDescriptor_e483730694cbd353) }

var fileDescriptor_e483730694cbd353 = []byte{
	// 226 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xc9, 0x4f, 0x4e,
	0x2c, 0x2d, 0xc9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x04, 0x71, 0x73, 0x40, 0x02,
	0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9, 0x05, 0x25, 0x95, 0x4a, 0x7a, 0x5c, 0x5c, 0x8e, 0x21, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x0a, 0x5c, 0xdc, 0x89, 0xc9, 0xc9, 0xa9, 0xc5,
	0xc5, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4a,
	0xf2, 0x5c, 0x9c, 0xa1, 0xce, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25, 0x42, 0x42, 0x5c, 0x2c,
	0xc9, 0xf9, 0x29, 0xa9, 0x50, 0x75, 0x60, 0xb6, 0x92, 0x06, 0x17, 0x4f, 0x40, 0x66, 0x5e, 0x3a,
	0xdc, 0x48, 0x09, 0x2e, 0xf6, 0xdc, 0xd4, 0xe2, 0xe2, 0xc4, 0x74, 0x98, 0x32, 0x18, 0xd7, 0x68,
	0x2b, 0x23, 0x17, 0xa7, 0x0f, 0xc8, 0x45, 0x8e, 0xa5, 0x25, 0x19, 0x42, 0xc6, 0x5c, 0x2c, 0x20,
	0x7d, 0x42, 0x02, 0x7a, 0x70, 0x57, 0xea, 0x81, 0x9d, 0x28, 0x25, 0x8e, 0x24, 0x82, 0x6c, 0xb4,
	0x12, 0x83, 0x90, 0x35, 0x17, 0x9f, 0x7b, 0x6a, 0x89, 0x23, 0xc2, 0x7d, 0x58, 0xb4, 0x8b, 0x22,
	0x89, 0x20, 0xbc, 0xaa, 0xc4, 0x20, 0x64, 0xc6, 0xc5, 0x15, 0x5a, 0x90, 0x92, 0x58, 0x92, 0xea,
	0x9c, 0x9f, 0x92, 0x2a, 0x24, 0x82, 0xa4, 0x0c, 0xee, 0x43, 0x29, 0x0c, 0xe3, 0x94, 0x18, 0x92,
	0xd8, 0xc0, 0xa1, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x71, 0x2e, 0x12, 0x5e, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LocalAuthClient is the client API for LocalAuth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LocalAuthClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error)
	GetAccessToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ATResponse, error)
	UpdateCode(ctx context.Context, in *UCRequest, opts ...grpc.CallOption) (*Empty, error)
}

type localAuthClient struct {
	cc *grpc.ClientConn
}

func NewLocalAuthClient(cc *grpc.ClientConn) LocalAuthClient {
	return &localAuthClient{cc}
}

func (c *localAuthClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/localauth.LocalAuth/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localAuthClient) GetAccessToken(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*ATResponse, error) {
	out := new(ATResponse)
	err := c.cc.Invoke(ctx, "/localauth.LocalAuth/GetAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *localAuthClient) UpdateCode(ctx context.Context, in *UCRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/localauth.LocalAuth/UpdateCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocalAuthServer is the server API for LocalAuth service.
type LocalAuthServer interface {
	Ping(context.Context, *Empty) (*PingResponse, error)
	GetAccessToken(context.Context, *Empty) (*ATResponse, error)
	UpdateCode(context.Context, *UCRequest) (*Empty, error)
}

// UnimplementedLocalAuthServer can be embedded to have forward compatible implementations.
type UnimplementedLocalAuthServer struct {
}

func (*UnimplementedLocalAuthServer) Ping(ctx context.Context, req *Empty) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedLocalAuthServer) GetAccessToken(ctx context.Context, req *Empty) (*ATResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessToken not implemented")
}
func (*UnimplementedLocalAuthServer) UpdateCode(ctx context.Context, req *UCRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCode not implemented")
}

func RegisterLocalAuthServer(s *grpc.Server, srv LocalAuthServer) {
	s.RegisterService(&_LocalAuth_serviceDesc, srv)
}

func _LocalAuth_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalAuthServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/localauth.LocalAuth/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalAuthServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalAuth_GetAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalAuthServer).GetAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/localauth.LocalAuth/GetAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalAuthServer).GetAccessToken(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocalAuth_UpdateCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UCRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocalAuthServer).UpdateCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/localauth.LocalAuth/UpdateCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocalAuthServer).UpdateCode(ctx, req.(*UCRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LocalAuth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "localauth.LocalAuth",
	HandlerType: (*LocalAuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _LocalAuth_Ping_Handler,
		},
		{
			MethodName: "GetAccessToken",
			Handler:    _LocalAuth_GetAccessToken_Handler,
		},
		{
			MethodName: "UpdateCode",
			Handler:    _LocalAuth_UpdateCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "locauth.proto",
}
