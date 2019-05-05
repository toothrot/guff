// Code generated by protoc-gen-go. DO NOT EDIT.
// source: divisions.proto

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

type GetDivisionsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDivisionsRequest) Reset()         { *m = GetDivisionsRequest{} }
func (m *GetDivisionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetDivisionsRequest) ProtoMessage()    {}
func (*GetDivisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6d0c2d87967cf, []int{0}
}

func (m *GetDivisionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDivisionsRequest.Unmarshal(m, b)
}
func (m *GetDivisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDivisionsRequest.Marshal(b, m, deterministic)
}
func (m *GetDivisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDivisionsRequest.Merge(m, src)
}
func (m *GetDivisionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetDivisionsRequest.Size(m)
}
func (m *GetDivisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDivisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDivisionsRequest proto.InternalMessageInfo

type GetDivisionsResponse struct {
	Divisions            []*Division `protobuf:"bytes,1,rep,name=divisions,proto3" json:"divisions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetDivisionsResponse) Reset()         { *m = GetDivisionsResponse{} }
func (m *GetDivisionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetDivisionsResponse) ProtoMessage()    {}
func (*GetDivisionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6d0c2d87967cf, []int{1}
}

func (m *GetDivisionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDivisionsResponse.Unmarshal(m, b)
}
func (m *GetDivisionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDivisionsResponse.Marshal(b, m, deterministic)
}
func (m *GetDivisionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDivisionsResponse.Merge(m, src)
}
func (m *GetDivisionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetDivisionsResponse.Size(m)
}
func (m *GetDivisionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDivisionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDivisionsResponse proto.InternalMessageInfo

func (m *GetDivisionsResponse) GetDivisions() []*Division {
	if m != nil {
		return m.Divisions
	}
	return nil
}

type Division struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Division) Reset()         { *m = Division{} }
func (m *Division) String() string { return proto.CompactTextString(m) }
func (*Division) ProtoMessage()    {}
func (*Division) Descriptor() ([]byte, []int) {
	return fileDescriptor_b3b6d0c2d87967cf, []int{2}
}

func (m *Division) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Division.Unmarshal(m, b)
}
func (m *Division) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Division.Marshal(b, m, deterministic)
}
func (m *Division) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Division.Merge(m, src)
}
func (m *Division) XXX_Size() int {
	return xxx_messageInfo_Division.Size(m)
}
func (m *Division) XXX_DiscardUnknown() {
	xxx_messageInfo_Division.DiscardUnknown(m)
}

var xxx_messageInfo_Division proto.InternalMessageInfo

func (m *Division) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*GetDivisionsRequest)(nil), "guff.proto.GetDivisionsRequest")
	proto.RegisterType((*GetDivisionsResponse)(nil), "guff.proto.GetDivisionsResponse")
	proto.RegisterType((*Division)(nil), "guff.proto.Division")
}

func init() { proto.RegisterFile("divisions.proto", fileDescriptor_b3b6d0c2d87967cf) }

var fileDescriptor_b3b6d0c2d87967cf = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0xc9, 0x2c, 0xcb,
	0x2c, 0xce, 0xcc, 0xcf, 0x2b, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2f, 0x4d,
	0x4b, 0x83, 0xb0, 0x95, 0x44, 0xb9, 0x84, 0xdd, 0x53, 0x4b, 0x5c, 0x60, 0x2a, 0x82, 0x52, 0x0b,
	0x4b, 0x53, 0x8b, 0x4b, 0x94, 0xbc, 0xb8, 0x44, 0x50, 0x85, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x85, 0x8c, 0xb8, 0x38, 0xe1, 0xa6, 0x49, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0xe8, 0x21,
	0x8c, 0xd3, 0x83, 0xe9, 0x08, 0x42, 0x28, 0x53, 0x92, 0xe2, 0xe2, 0x80, 0x09, 0x0b, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x18, 0xa5, 0x73,
	0x09, 0xc0, 0x2d, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x0a, 0xe6, 0xe2, 0x41, 0xb6,
	0x5b, 0x48, 0x1e, 0xd9, 0x02, 0x2c, 0x8e, 0x95, 0x52, 0xc0, 0xad, 0x00, 0xe2, 0x6c, 0x25, 0x86,
	0x24, 0x36, 0xb0, 0xac, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x2d, 0xbf, 0x88, 0x0d, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DivisionsServiceClient is the client API for DivisionsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DivisionsServiceClient interface {
	GetDivisions(ctx context.Context, in *GetDivisionsRequest, opts ...grpc.CallOption) (*GetDivisionsResponse, error)
}

type divisionsServiceClient struct {
	cc *grpc.ClientConn
}

func NewDivisionsServiceClient(cc *grpc.ClientConn) DivisionsServiceClient {
	return &divisionsServiceClient{cc}
}

func (c *divisionsServiceClient) GetDivisions(ctx context.Context, in *GetDivisionsRequest, opts ...grpc.CallOption) (*GetDivisionsResponse, error) {
	out := new(GetDivisionsResponse)
	err := c.cc.Invoke(ctx, "/guff.proto.DivisionsService/GetDivisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivisionsServiceServer is the server API for DivisionsService service.
type DivisionsServiceServer interface {
	GetDivisions(context.Context, *GetDivisionsRequest) (*GetDivisionsResponse, error)
}

// UnimplementedDivisionsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDivisionsServiceServer struct {
}

func (*UnimplementedDivisionsServiceServer) GetDivisions(ctx context.Context, req *GetDivisionsRequest) (*GetDivisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDivisions not implemented")
}

func RegisterDivisionsServiceServer(s *grpc.Server, srv DivisionsServiceServer) {
	s.RegisterService(&_DivisionsService_serviceDesc, srv)
}

func _DivisionsService_GetDivisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDivisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivisionsServiceServer).GetDivisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guff.proto.DivisionsService/GetDivisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivisionsServiceServer).GetDivisions(ctx, req.(*GetDivisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DivisionsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guff.proto.DivisionsService",
	HandlerType: (*DivisionsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDivisions",
			Handler:    _DivisionsService_GetDivisions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "divisions.proto",
}
