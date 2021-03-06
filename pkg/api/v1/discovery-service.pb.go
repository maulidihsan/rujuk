// Code generated by protoc-gen-go. DO NOT EDIT.
// source: discovery-service.proto

package v1

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

type Rumahsakit struct {
	// id ruangan
	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// nama rumah sakit
	Nama string `protobuf:"bytes,2,opt,name=nama,proto3" json:"nama,omitempty"`
	// ip rumah sakit
	Ip                   string   `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rumahsakit) Reset()         { *m = Rumahsakit{} }
func (m *Rumahsakit) String() string { return proto.CompactTextString(m) }
func (*Rumahsakit) ProtoMessage()    {}
func (*Rumahsakit) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c950d8070295dc5, []int{0}
}

func (m *Rumahsakit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rumahsakit.Unmarshal(m, b)
}
func (m *Rumahsakit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rumahsakit.Marshal(b, m, deterministic)
}
func (m *Rumahsakit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rumahsakit.Merge(m, src)
}
func (m *Rumahsakit) XXX_Size() int {
	return xxx_messageInfo_Rumahsakit.Size(m)
}
func (m *Rumahsakit) XXX_DiscardUnknown() {
	xxx_messageInfo_Rumahsakit.DiscardUnknown(m)
}

var xxx_messageInfo_Rumahsakit proto.InternalMessageInfo

func (m *Rumahsakit) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Rumahsakit) GetNama() string {
	if m != nil {
		return m.Nama
	}
	return ""
}

func (m *Rumahsakit) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type Rumahsakits struct {
	Rumahsakits          []*Rumahsakit `protobuf:"bytes,1,rep,name=rumahsakits,proto3" json:"rumahsakits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Rumahsakits) Reset()         { *m = Rumahsakits{} }
func (m *Rumahsakits) String() string { return proto.CompactTextString(m) }
func (*Rumahsakits) ProtoMessage()    {}
func (*Rumahsakits) Descriptor() ([]byte, []int) {
	return fileDescriptor_7c950d8070295dc5, []int{1}
}

func (m *Rumahsakits) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rumahsakits.Unmarshal(m, b)
}
func (m *Rumahsakits) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rumahsakits.Marshal(b, m, deterministic)
}
func (m *Rumahsakits) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rumahsakits.Merge(m, src)
}
func (m *Rumahsakits) XXX_Size() int {
	return xxx_messageInfo_Rumahsakits.Size(m)
}
func (m *Rumahsakits) XXX_DiscardUnknown() {
	xxx_messageInfo_Rumahsakits.DiscardUnknown(m)
}

var xxx_messageInfo_Rumahsakits proto.InternalMessageInfo

func (m *Rumahsakits) GetRumahsakits() []*Rumahsakit {
	if m != nil {
		return m.Rumahsakits
	}
	return nil
}

func init() {
	proto.RegisterType((*Rumahsakit)(nil), "v1.Rumahsakit")
	proto.RegisterType((*Rumahsakits)(nil), "v1.Rumahsakits")
}

func init() { proto.RegisterFile("discovery-service.proto", fileDescriptor_7c950d8070295dc5) }

var fileDescriptor_7c950d8070295dc5 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0x86, 0xb7, 0x59, 0x11, 0x9d, 0xea, 0x5a, 0xc6, 0x83, 0xa5, 0xa7, 0xd2, 0x53, 0x11, 0x2c,
	0xee, 0x7a, 0xf0, 0xa8, 0x82, 0xe8, 0x3d, 0xfe, 0x82, 0xd8, 0x0e, 0x6e, 0xdc, 0x8f, 0xc4, 0x4c,
	0x12, 0xf1, 0xdf, 0x8b, 0x11, 0xe9, 0xae, 0xb7, 0xbc, 0x99, 0x97, 0x67, 0x1e, 0x06, 0x2e, 0x06,
	0xcd, 0xbd, 0x89, 0xe4, 0xbe, 0xae, 0x98, 0x5c, 0xd4, 0x3d, 0x75, 0xd6, 0x19, 0x6f, 0x50, 0xc4,
	0x79, 0x75, 0xee, 0xc2, 0x7b, 0x58, 0xed, 0x0f, 0x9a, 0x7b, 0x00, 0x19, 0x36, 0x6a, 0xc9, 0x6a,
	0xa5, 0x3d, 0xce, 0x40, 0xe8, 0xa1, 0xcc, 0xea, 0xac, 0x3d, 0x95, 0x42, 0x0f, 0x88, 0x70, 0xb0,
	0x55, 0x1b, 0x55, 0x8a, 0x3a, 0x6b, 0x8f, 0x65, 0x7a, 0xa7, 0x8e, 0x2d, 0xa7, 0xe9, 0x47, 0x68,
	0xdb, 0xdc, 0x41, 0x3e, 0x12, 0x18, 0xaf, 0x21, 0x77, 0x63, 0x2c, 0xb3, 0x7a, 0xda, 0xe6, 0x8b,
	0x59, 0x17, 0xe7, 0xdd, 0xd8, 0x92, 0xbb, 0x95, 0xc5, 0x27, 0x14, 0x8f, 0x7f, 0xda, 0x2f, 0xbf,
	0x72, 0x78, 0x0b, 0xc5, 0x33, 0xf9, 0x87, 0xf5, 0x7a, 0x47, 0xae, 0xf8, 0x81, 0x3c, 0x91, 0xef,
	0x97, 0x92, 0x3e, 0x02, 0xb1, 0xaf, 0xce, 0xf6, 0xb1, 0xdc, 0x4c, 0xf0, 0x12, 0x8e, 0x24, 0xbd,
	0x69, 0xf6, 0xe4, 0xf0, 0xdf, 0xd6, 0xea, 0x24, 0x65, 0x62, 0x6b, 0xb6, 0x4c, 0xcd, 0xe4, 0xf5,
	0x30, 0x9d, 0xe0, 0xe6, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x87, 0x4c, 0xb3, 0x11, 0x36, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DiscoveryServiceClient is the client API for DiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DiscoveryServiceClient interface {
	GetAllRumahsakit(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Rumahsakits, error)
	Register(ctx context.Context, in *Rumahsakit, opts ...grpc.CallOption) (*Response, error)
}

type discoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewDiscoveryServiceClient(cc *grpc.ClientConn) DiscoveryServiceClient {
	return &discoveryServiceClient{cc}
}

func (c *discoveryServiceClient) GetAllRumahsakit(ctx context.Context, in *FetchRequest, opts ...grpc.CallOption) (*Rumahsakits, error) {
	out := new(Rumahsakits)
	err := c.cc.Invoke(ctx, "/v1.DiscoveryService/GetAllRumahsakit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *discoveryServiceClient) Register(ctx context.Context, in *Rumahsakit, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v1.DiscoveryService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiscoveryServiceServer is the server API for DiscoveryService service.
type DiscoveryServiceServer interface {
	GetAllRumahsakit(context.Context, *FetchRequest) (*Rumahsakits, error)
	Register(context.Context, *Rumahsakit) (*Response, error)
}

// UnimplementedDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDiscoveryServiceServer struct {
}

func (*UnimplementedDiscoveryServiceServer) GetAllRumahsakit(ctx context.Context, req *FetchRequest) (*Rumahsakits, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllRumahsakit not implemented")
}
func (*UnimplementedDiscoveryServiceServer) Register(ctx context.Context, req *Rumahsakit) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterDiscoveryServiceServer(s *grpc.Server, srv DiscoveryServiceServer) {
	s.RegisterService(&_DiscoveryService_serviceDesc, srv)
}

func _DiscoveryService_GetAllRumahsakit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).GetAllRumahsakit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DiscoveryService/GetAllRumahsakit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).GetAllRumahsakit(ctx, req.(*FetchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiscoveryService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Rumahsakit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiscoveryServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.DiscoveryService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiscoveryServiceServer).Register(ctx, req.(*Rumahsakit))
	}
	return interceptor(ctx, in, info, handler)
}

var _DiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.DiscoveryService",
	HandlerType: (*DiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllRumahsakit",
			Handler:    _DiscoveryService_GetAllRumahsakit_Handler,
		},
		{
			MethodName: "Register",
			Handler:    _DiscoveryService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "discovery-service.proto",
}
