// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package guff_proto

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

type GetCurrentUserRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserRequest) Reset()         { *m = GetCurrentUserRequest{} }
func (m *GetCurrentUserRequest) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserRequest) ProtoMessage()    {}
func (*GetCurrentUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *GetCurrentUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserRequest.Unmarshal(m, b)
}
func (m *GetCurrentUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserRequest.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserRequest.Merge(m, src)
}
func (m *GetCurrentUserRequest) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserRequest.Size(m)
}
func (m *GetCurrentUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserRequest proto.InternalMessageInfo

type GetCurrentUserResponse struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCurrentUserResponse) Reset()         { *m = GetCurrentUserResponse{} }
func (m *GetCurrentUserResponse) String() string { return proto.CompactTextString(m) }
func (*GetCurrentUserResponse) ProtoMessage()    {}
func (*GetCurrentUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *GetCurrentUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCurrentUserResponse.Unmarshal(m, b)
}
func (m *GetCurrentUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCurrentUserResponse.Marshal(b, m, deterministic)
}
func (m *GetCurrentUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCurrentUserResponse.Merge(m, src)
}
func (m *GetCurrentUserResponse) XXX_Size() int {
	return xxx_messageInfo_GetCurrentUserResponse.Size(m)
}
func (m *GetCurrentUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCurrentUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCurrentUserResponse proto.InternalMessageInfo

func (m *GetCurrentUserResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*GetCurrentUserRequest)(nil), "guff.proto.GetCurrentUserRequest")
	proto.RegisterType((*GetCurrentUserResponse)(nil), "guff.proto.GetCurrentUserResponse")
}

func init() { proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea) }

var fileDescriptor_030765f334c86cea = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2f, 0x4d, 0x4b, 0x83, 0xb0, 0x95,
	0xc4, 0xb9, 0x44, 0xdd, 0x53, 0x4b, 0x9c, 0x4b, 0x8b, 0x8a, 0x52, 0xf3, 0x4a, 0x42, 0x8b, 0x53,
	0x8b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xf4, 0xb8, 0xc4, 0xd0, 0x25, 0x8a, 0x0b,
	0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x53, 0x73, 0x13, 0x33, 0x73, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0xa3, 0x4c, 0x2e, 0x1e, 0x90, 0xaa, 0xe2, 0xe0, 0xd4, 0xa2,
	0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x48, 0x2e, 0x3e, 0x54, 0xfd, 0x42, 0x8a, 0x7a, 0x08, 0x7b, 0xf5,
	0xb0, 0x5a, 0x2a, 0xa5, 0x84, 0x4f, 0x09, 0xc4, 0x7a, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xbc, 0x31,
	0x20, 0x00, 0x00, 0xff, 0xff, 0xe3, 0x76, 0x55, 0xd9, 0xd5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UsersServiceClient is the client API for UsersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersServiceClient interface {
	GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error)
}

type usersServiceClient struct {
	cc *grpc.ClientConn
}

func NewUsersServiceClient(cc *grpc.ClientConn) UsersServiceClient {
	return &usersServiceClient{cc}
}

func (c *usersServiceClient) GetCurrentUser(ctx context.Context, in *GetCurrentUserRequest, opts ...grpc.CallOption) (*GetCurrentUserResponse, error) {
	out := new(GetCurrentUserResponse)
	err := c.cc.Invoke(ctx, "/guff.proto.UsersService/GetCurrentUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServiceServer is the server API for UsersService service.
type UsersServiceServer interface {
	GetCurrentUser(context.Context, *GetCurrentUserRequest) (*GetCurrentUserResponse, error)
}

// UnimplementedUsersServiceServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServiceServer struct {
}

func (*UnimplementedUsersServiceServer) GetCurrentUser(ctx context.Context, req *GetCurrentUserRequest) (*GetCurrentUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentUser not implemented")
}

func RegisterUsersServiceServer(s *grpc.Server, srv UsersServiceServer) {
	s.RegisterService(&_UsersService_serviceDesc, srv)
}

func _UsersService_GetCurrentUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServiceServer).GetCurrentUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guff.proto.UsersService/GetCurrentUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServiceServer).GetCurrentUser(ctx, req.(*GetCurrentUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UsersService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guff.proto.UsersService",
	HandlerType: (*UsersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentUser",
			Handler:    _UsersService_GetCurrentUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
