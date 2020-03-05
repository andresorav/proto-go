// Code generated by protoc-gen-go. DO NOT EDIT.
// source: users.proto

package proto

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

type GetUser struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetUser) Reset()         { *m = GetUser{} }
func (m *GetUser) String() string { return proto.CompactTextString(m) }
func (*GetUser) ProtoMessage()    {}
func (*GetUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{0}
}

func (m *GetUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetUser.Unmarshal(m, b)
}
func (m *GetUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetUser.Marshal(b, m, deterministic)
}
func (m *GetUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetUser.Merge(m, src)
}
func (m *GetUser) XXX_Size() int {
	return xxx_messageInfo_GetUser.Size(m)
}
func (m *GetUser) XXX_DiscardUnknown() {
	xxx_messageInfo_GetUser.DiscardUnknown(m)
}

var xxx_messageInfo_GetUser proto.InternalMessageInfo

func (m *GetUser) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UserFound struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserFound) Reset()         { *m = UserFound{} }
func (m *UserFound) String() string { return proto.CompactTextString(m) }
func (*UserFound) ProtoMessage()    {}
func (*UserFound) Descriptor() ([]byte, []int) {
	return fileDescriptor_030765f334c86cea, []int{1}
}

func (m *UserFound) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserFound.Unmarshal(m, b)
}
func (m *UserFound) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserFound.Marshal(b, m, deterministic)
}
func (m *UserFound) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserFound.Merge(m, src)
}
func (m *UserFound) XXX_Size() int {
	return xxx_messageInfo_UserFound.Size(m)
}
func (m *UserFound) XXX_DiscardUnknown() {
	xxx_messageInfo_UserFound.DiscardUnknown(m)
}

var xxx_messageInfo_UserFound proto.InternalMessageInfo

func (m *UserFound) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*GetUser)(nil), "proto.GetUser")
	proto.RegisterType((*UserFound)(nil), "proto.UserFound")
}

func init() {
	proto.RegisterFile("users.proto", fileDescriptor_030765f334c86cea)
}

var fileDescriptor_030765f334c86cea = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x2d, 0x4e, 0x2d,
	0x2a, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x92, 0x5c, 0xec, 0xee,
	0xa9, 0x25, 0xa1, 0xc5, 0xa9, 0x45, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x4c, 0x99, 0x29, 0x4a, 0xea, 0x5c, 0x9c, 0x20, 0x71, 0xb7, 0xfc, 0xd2, 0xbc,
	0x14, 0x21, 0x29, 0x2e, 0x0e, 0x90, 0xee, 0xbc, 0xc4, 0xdc, 0x54, 0xa8, 0x12, 0x38, 0xdf, 0xc8,
	0x88, 0x8b, 0x15, 0xa4, 0xb0, 0x58, 0x48, 0x93, 0x8b, 0xd9, 0x3f, 0x2f, 0x55, 0x88, 0x0f, 0x62,
	0x85, 0x1e, 0xd4, 0x60, 0x29, 0x01, 0x28, 0x1f, 0x6e, 0x9a, 0x12, 0x43, 0x12, 0x1b, 0x58, 0xc8,
	0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x06, 0x11, 0x32, 0x0f, 0x94, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UsersClient is the client API for Users service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UsersClient interface {
	One(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*UserFound, error)
}

type usersClient struct {
	cc grpc.ClientConnInterface
}

func NewUsersClient(cc grpc.ClientConnInterface) UsersClient {
	return &usersClient{cc}
}

func (c *usersClient) One(ctx context.Context, in *GetUser, opts ...grpc.CallOption) (*UserFound, error) {
	out := new(UserFound)
	err := c.cc.Invoke(ctx, "/proto.Users/One", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UsersServer is the server API for Users service.
type UsersServer interface {
	One(context.Context, *GetUser) (*UserFound, error)
}

// UnimplementedUsersServer can be embedded to have forward compatible implementations.
type UnimplementedUsersServer struct {
}

func (*UnimplementedUsersServer) One(ctx context.Context, req *GetUser) (*UserFound, error) {
	return nil, status.Errorf(codes.Unimplemented, "method One not implemented")
}

func RegisterUsersServer(s *grpc.Server, srv UsersServer) {
	s.RegisterService(&_Users_serviceDesc, srv)
}

func _Users_One_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UsersServer).One(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Users/One",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UsersServer).One(ctx, req.(*GetUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _Users_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Users",
	HandlerType: (*UsersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "One",
			Handler:    _Users_One_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "users.proto",
}
