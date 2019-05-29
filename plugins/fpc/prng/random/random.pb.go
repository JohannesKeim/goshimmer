// Code generated by protoc-gen-go. DO NOT EDIT.
// source: random/random.proto

package random

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

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{0}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

type Random struct {
	Index                int64    `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Random) Reset()         { *m = Random{} }
func (m *Random) String() string { return proto.CompactTextString(m) }
func (*Random) ProtoMessage()    {}
func (*Random) Descriptor() ([]byte, []int) {
	return fileDescriptor_e5da2919a686585f, []int{1}
}

func (m *Random) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Random.Unmarshal(m, b)
}
func (m *Random) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Random.Marshal(b, m, deterministic)
}
func (m *Random) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Random.Merge(m, src)
}
func (m *Random) XXX_Size() int {
	return xxx_messageInfo_Random.Size(m)
}
func (m *Random) XXX_DiscardUnknown() {
	xxx_messageInfo_Random.DiscardUnknown(m)
}

var xxx_messageInfo_Random proto.InternalMessageInfo

func (m *Random) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Random) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto.RegisterType((*Void)(nil), "random.Void")
	proto.RegisterType((*Random)(nil), "random.Random")
}

func init() { proto.RegisterFile("random/random.proto", fileDescriptor_e5da2919a686585f) }

var fileDescriptor_e5da2919a686585f = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x12,
	0x1b, 0x17, 0x4b, 0x58, 0x7e, 0x66, 0x8a, 0x92, 0x09, 0x17, 0x5b, 0x10, 0x58, 0x44, 0x48, 0x84,
	0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x39, 0x08, 0xc2, 0x01,
	0x89, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x30, 0x06, 0x41, 0x38, 0x46,
	0x0e, 0x5c, 0xfc, 0x10, 0x5d, 0xee, 0xa9, 0x79, 0xa9, 0x45, 0x89, 0x25, 0xf9, 0x45, 0x42, 0xba,
	0x5c, 0x9c, 0xc1, 0xa5, 0x49, 0xc5, 0xc9, 0x45, 0x99, 0x49, 0xa9, 0x42, 0x3c, 0x7a, 0x50, 0x4b,
	0x41, 0x76, 0x48, 0xf1, 0xc1, 0x78, 0x10, 0x3d, 0x4a, 0x0c, 0x06, 0x8c, 0x49, 0x6c, 0x60, 0xe7,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x12, 0x1f, 0xf1, 0x4a, 0xa5, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RandomGeneratorClient is the client API for RandomGenerator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RandomGeneratorClient interface {
	Subscribe(ctx context.Context, in *Void, opts ...grpc.CallOption) (RandomGenerator_SubscribeClient, error)
}

type randomGeneratorClient struct {
	cc *grpc.ClientConn
}

func NewRandomGeneratorClient(cc *grpc.ClientConn) RandomGeneratorClient {
	return &randomGeneratorClient{cc}
}

func (c *randomGeneratorClient) Subscribe(ctx context.Context, in *Void, opts ...grpc.CallOption) (RandomGenerator_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RandomGenerator_serviceDesc.Streams[0], "/random.RandomGenerator/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &randomGeneratorSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RandomGenerator_SubscribeClient interface {
	Recv() (*Random, error)
	grpc.ClientStream
}

type randomGeneratorSubscribeClient struct {
	grpc.ClientStream
}

func (x *randomGeneratorSubscribeClient) Recv() (*Random, error) {
	m := new(Random)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RandomGeneratorServer is the server API for RandomGenerator service.
type RandomGeneratorServer interface {
	Subscribe(*Void, RandomGenerator_SubscribeServer) error
}

// UnimplementedRandomGeneratorServer can be embedded to have forward compatible implementations.
type UnimplementedRandomGeneratorServer struct {
}

func (*UnimplementedRandomGeneratorServer) Subscribe(req *Void, srv RandomGenerator_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterRandomGeneratorServer(s *grpc.Server, srv RandomGeneratorServer) {
	s.RegisterService(&_RandomGenerator_serviceDesc, srv)
}

func _RandomGenerator_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Void)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RandomGeneratorServer).Subscribe(m, &randomGeneratorSubscribeServer{stream})
}

type RandomGenerator_SubscribeServer interface {
	Send(*Random) error
	grpc.ServerStream
}

type randomGeneratorSubscribeServer struct {
	grpc.ServerStream
}

func (x *randomGeneratorSubscribeServer) Send(m *Random) error {
	return x.ServerStream.SendMsg(m)
}

var _RandomGenerator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "random.RandomGenerator",
	HandlerType: (*RandomGeneratorServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _RandomGenerator_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "random/random.proto",
}
