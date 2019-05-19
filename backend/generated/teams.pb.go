// Code generated by protoc-gen-go. DO NOT EDIT.
// source: teams.proto

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

type GetTeamsRequest struct {
	DivisionId           string   `protobuf:"bytes,1,opt,name=division_id,json=divisionId,proto3" json:"division_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsRequest) Reset()         { *m = GetTeamsRequest{} }
func (m *GetTeamsRequest) String() string { return proto.CompactTextString(m) }
func (*GetTeamsRequest) ProtoMessage()    {}
func (*GetTeamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{0}
}

func (m *GetTeamsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsRequest.Unmarshal(m, b)
}
func (m *GetTeamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsRequest.Marshal(b, m, deterministic)
}
func (m *GetTeamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsRequest.Merge(m, src)
}
func (m *GetTeamsRequest) XXX_Size() int {
	return xxx_messageInfo_GetTeamsRequest.Size(m)
}
func (m *GetTeamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsRequest proto.InternalMessageInfo

func (m *GetTeamsRequest) GetDivisionId() string {
	if m != nil {
		return m.DivisionId
	}
	return ""
}

type GetTeamsResponse struct {
	Teams                []*Team  `protobuf:"bytes,1,rep,name=teams,proto3" json:"teams,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTeamsResponse) Reset()         { *m = GetTeamsResponse{} }
func (m *GetTeamsResponse) String() string { return proto.CompactTextString(m) }
func (*GetTeamsResponse) ProtoMessage()    {}
func (*GetTeamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{1}
}

func (m *GetTeamsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTeamsResponse.Unmarshal(m, b)
}
func (m *GetTeamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTeamsResponse.Marshal(b, m, deterministic)
}
func (m *GetTeamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTeamsResponse.Merge(m, src)
}
func (m *GetTeamsResponse) XXX_Size() int {
	return xxx_messageInfo_GetTeamsResponse.Size(m)
}
func (m *GetTeamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTeamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTeamsResponse proto.InternalMessageInfo

func (m *GetTeamsResponse) GetTeams() []*Team {
	if m != nil {
		return m.Teams
	}
	return nil
}

type Team struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DivisionId           string   `protobuf:"bytes,3,opt,name=division_id,json=divisionId,proto3" json:"division_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Team) Reset()         { *m = Team{} }
func (m *Team) String() string { return proto.CompactTextString(m) }
func (*Team) ProtoMessage()    {}
func (*Team) Descriptor() ([]byte, []int) {
	return fileDescriptor_f63f4a1b2b4dddb4, []int{2}
}

func (m *Team) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Team.Unmarshal(m, b)
}
func (m *Team) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Team.Marshal(b, m, deterministic)
}
func (m *Team) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Team.Merge(m, src)
}
func (m *Team) XXX_Size() int {
	return xxx_messageInfo_Team.Size(m)
}
func (m *Team) XXX_DiscardUnknown() {
	xxx_messageInfo_Team.DiscardUnknown(m)
}

var xxx_messageInfo_Team proto.InternalMessageInfo

func (m *Team) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Team) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Team) GetDivisionId() string {
	if m != nil {
		return m.DivisionId
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTeamsRequest)(nil), "guff.proto.GetTeamsRequest")
	proto.RegisterType((*GetTeamsResponse)(nil), "guff.proto.GetTeamsResponse")
	proto.RegisterType((*Team)(nil), "guff.proto.Team")
}

func init() { proto.RegisterFile("teams.proto", fileDescriptor_f63f4a1b2b4dddb4) }

var fileDescriptor_f63f4a1b2b4dddb4 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x49, 0x4d, 0xcc,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2f, 0x4d, 0x4b, 0x83, 0xb0, 0x95,
	0x8c, 0xb8, 0xf8, 0xdd, 0x53, 0x4b, 0x42, 0x40, 0xb2, 0x41, 0xa9, 0x85, 0xa5, 0xa9, 0xc5, 0x25,
	0x42, 0xf2, 0x5c, 0xdc, 0x29, 0x99, 0x65, 0x99, 0xc5, 0x99, 0xf9, 0x79, 0xf1, 0x99, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x5c, 0x30, 0x21, 0xcf, 0x14, 0x25, 0x2b, 0x2e, 0x01, 0x84,
	0x9e, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x35, 0x2e, 0x56, 0xb0, 0x15, 0x12, 0x8c, 0x0a,
	0xcc, 0x1a, 0xdc, 0x46, 0x02, 0x7a, 0x08, 0x3b, 0xf4, 0x40, 0x2a, 0x83, 0x20, 0xd2, 0x4a, 0xde,
	0x5c, 0x2c, 0x20, 0xae, 0x10, 0x1f, 0x17, 0x13, 0xdc, 0x6c, 0xa6, 0xcc, 0x14, 0x21, 0x21, 0x2e,
	0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x26, 0xb0, 0x08, 0x98, 0x8d, 0xee, 0x10, 0x66, 0x74, 0x87,
	0x18, 0x85, 0x73, 0xf1, 0x80, 0x5d, 0x11, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xce,
	0xc5, 0x01, 0x73, 0x98, 0x90, 0x34, 0xb2, 0x0b, 0xd0, 0xbc, 0x28, 0x25, 0x83, 0x5d, 0x12, 0xe2,
	0x17, 0x25, 0x86, 0x24, 0x36, 0xb0, 0x8c, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa4, 0xab, 0xfa,
	0xad, 0x37, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TeamsServiceClient is the client API for TeamsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TeamsServiceClient interface {
	GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error)
}

type teamsServiceClient struct {
	cc *grpc.ClientConn
}

func NewTeamsServiceClient(cc *grpc.ClientConn) TeamsServiceClient {
	return &teamsServiceClient{cc}
}

func (c *teamsServiceClient) GetTeams(ctx context.Context, in *GetTeamsRequest, opts ...grpc.CallOption) (*GetTeamsResponse, error) {
	out := new(GetTeamsResponse)
	err := c.cc.Invoke(ctx, "/guff.proto.TeamsService/GetTeams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamsServiceServer is the server API for TeamsService service.
type TeamsServiceServer interface {
	GetTeams(context.Context, *GetTeamsRequest) (*GetTeamsResponse, error)
}

// UnimplementedTeamsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTeamsServiceServer struct {
}

func (*UnimplementedTeamsServiceServer) GetTeams(ctx context.Context, req *GetTeamsRequest) (*GetTeamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTeams not implemented")
}

func RegisterTeamsServiceServer(s *grpc.Server, srv TeamsServiceServer) {
	s.RegisterService(&_TeamsService_serviceDesc, srv)
}

func _TeamsService_GetTeams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTeamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamsServiceServer).GetTeams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/guff.proto.TeamsService/GetTeams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamsServiceServer).GetTeams(ctx, req.(*GetTeamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TeamsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "guff.proto.TeamsService",
	HandlerType: (*TeamsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTeams",
			Handler:    _TeamsService_GetTeams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teams.proto",
}
