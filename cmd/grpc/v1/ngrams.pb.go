// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/ngrams.proto

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

type LearnRequest struct {
	// body is the body of data being learned.
	Body                 string   `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LearnRequest) Reset()         { *m = LearnRequest{} }
func (m *LearnRequest) String() string { return proto.CompactTextString(m) }
func (*LearnRequest) ProtoMessage()    {}
func (*LearnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_566f5b74984976ae, []int{0}
}

func (m *LearnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearnRequest.Unmarshal(m, b)
}
func (m *LearnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearnRequest.Marshal(b, m, deterministic)
}
func (m *LearnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearnRequest.Merge(m, src)
}
func (m *LearnRequest) XXX_Size() int {
	return xxx_messageInfo_LearnRequest.Size(m)
}
func (m *LearnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LearnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LearnRequest proto.InternalMessageInfo

func (m *LearnRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type LearnResponse struct {
	// parsed_tokens is the number of tokens parsed in the training.
	ParsedTokens         int64    `protobuf:"varint,1,opt,name=parsed_tokens,json=parsedTokens,proto3" json:"parsed_tokens,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LearnResponse) Reset()         { *m = LearnResponse{} }
func (m *LearnResponse) String() string { return proto.CompactTextString(m) }
func (*LearnResponse) ProtoMessage()    {}
func (*LearnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_566f5b74984976ae, []int{1}
}

func (m *LearnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LearnResponse.Unmarshal(m, b)
}
func (m *LearnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LearnResponse.Marshal(b, m, deterministic)
}
func (m *LearnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LearnResponse.Merge(m, src)
}
func (m *LearnResponse) XXX_Size() int {
	return xxx_messageInfo_LearnResponse.Size(m)
}
func (m *LearnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LearnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LearnResponse proto.InternalMessageInfo

func (m *LearnResponse) GetParsedTokens() int64 {
	if m != nil {
		return m.ParsedTokens
	}
	return 0
}

type GenerateRequest struct {
	// limit is the target length of the output in tokens.
	Limit                int64    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateRequest) Reset()         { *m = GenerateRequest{} }
func (m *GenerateRequest) String() string { return proto.CompactTextString(m) }
func (*GenerateRequest) ProtoMessage()    {}
func (*GenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_566f5b74984976ae, []int{2}
}

func (m *GenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateRequest.Unmarshal(m, b)
}
func (m *GenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateRequest.Marshal(b, m, deterministic)
}
func (m *GenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateRequest.Merge(m, src)
}
func (m *GenerateRequest) XXX_Size() int {
	return xxx_messageInfo_GenerateRequest.Size(m)
}
func (m *GenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateRequest proto.InternalMessageInfo

func (m *GenerateRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type GenerateResponse struct {
	// body is the output that was generated.
	Body string `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
	// limit is the target length of tokens that was used (==GenerateRequest.tokens)
	Limit                int64    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GenerateResponse) Reset()         { *m = GenerateResponse{} }
func (m *GenerateResponse) String() string { return proto.CompactTextString(m) }
func (*GenerateResponse) ProtoMessage()    {}
func (*GenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_566f5b74984976ae, []int{3}
}

func (m *GenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GenerateResponse.Unmarshal(m, b)
}
func (m *GenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GenerateResponse.Marshal(b, m, deterministic)
}
func (m *GenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenerateResponse.Merge(m, src)
}
func (m *GenerateResponse) XXX_Size() int {
	return xxx_messageInfo_GenerateResponse.Size(m)
}
func (m *GenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GenerateResponse proto.InternalMessageInfo

func (m *GenerateResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *GenerateResponse) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func init() {
	proto.RegisterType((*LearnRequest)(nil), "v1.LearnRequest")
	proto.RegisterType((*LearnResponse)(nil), "v1.LearnResponse")
	proto.RegisterType((*GenerateRequest)(nil), "v1.GenerateRequest")
	proto.RegisterType((*GenerateResponse)(nil), "v1.GenerateResponse")
}

func init() { proto.RegisterFile("v1/ngrams.proto", fileDescriptor_566f5b74984976ae) }

var fileDescriptor_566f5b74984976ae = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x33, 0xd4, 0xcf,
	0x4b, 0x2f, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x54,
	0x52, 0xe2, 0xe2, 0xf1, 0x49, 0x4d, 0x2c, 0xca, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11,
	0x12, 0xe2, 0x62, 0x49, 0xca, 0x4f, 0xa9, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3,
	0x95, 0x4c, 0xb8, 0x78, 0xa1, 0x6a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x94, 0xb9, 0x78,
	0x0b, 0x12, 0x8b, 0x8a, 0x53, 0x53, 0xe2, 0x4b, 0xf2, 0xb3, 0x53, 0xf3, 0x8a, 0xc1, 0xaa, 0x99,
	0x83, 0x78, 0x20, 0x82, 0x21, 0x60, 0x31, 0x25, 0x75, 0x2e, 0x7e, 0xf7, 0xd4, 0xbc, 0xd4, 0xa2,
	0xc4, 0x92, 0x54, 0x98, 0xe1, 0x22, 0x5c, 0xac, 0x39, 0x99, 0xb9, 0x99, 0x25, 0x50, 0xf5, 0x10,
	0x8e, 0x92, 0x0d, 0x97, 0x00, 0x42, 0x21, 0xd4, 0x06, 0x2c, 0xce, 0x40, 0xe8, 0x66, 0x42, 0xd2,
	0x6d, 0x54, 0xcc, 0xc5, 0xe3, 0x07, 0xf2, 0x54, 0x70, 0x6a, 0x51, 0x59, 0x66, 0x72, 0xaa, 0x90,
	0x0e, 0x17, 0x2b, 0xd8, 0xb1, 0x42, 0x02, 0x7a, 0x65, 0x86, 0x7a, 0xc8, 0x7e, 0x93, 0x12, 0x44,
	0x12, 0x81, 0xda, 0x63, 0xca, 0xc5, 0x01, 0xb3, 0x5b, 0x48, 0x18, 0x24, 0x8d, 0xe6, 0x64, 0x29,
	0x11, 0x54, 0x41, 0x88, 0xb6, 0x24, 0x36, 0x70, 0x00, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x3d, 0x4d, 0x0d, 0x9b, 0x53, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NgramServiceClient is the client API for NgramService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NgramServiceClient interface {
	// Learn trains the ngram index on a corpus of text.
	Learn(ctx context.Context, in *LearnRequest, opts ...grpc.CallOption) (*LearnResponse, error)
	// Generate outputs a random string in the trained style.
	Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error)
}

type ngramServiceClient struct {
	cc *grpc.ClientConn
}

func NewNgramServiceClient(cc *grpc.ClientConn) NgramServiceClient {
	return &ngramServiceClient{cc}
}

func (c *ngramServiceClient) Learn(ctx context.Context, in *LearnRequest, opts ...grpc.CallOption) (*LearnResponse, error) {
	out := new(LearnResponse)
	err := c.cc.Invoke(ctx, "/v1.NgramService/Learn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ngramServiceClient) Generate(ctx context.Context, in *GenerateRequest, opts ...grpc.CallOption) (*GenerateResponse, error) {
	out := new(GenerateResponse)
	err := c.cc.Invoke(ctx, "/v1.NgramService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NgramServiceServer is the server API for NgramService service.
type NgramServiceServer interface {
	// Learn trains the ngram index on a corpus of text.
	Learn(context.Context, *LearnRequest) (*LearnResponse, error)
	// Generate outputs a random string in the trained style.
	Generate(context.Context, *GenerateRequest) (*GenerateResponse, error)
}

// UnimplementedNgramServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNgramServiceServer struct {
}

func (*UnimplementedNgramServiceServer) Learn(ctx context.Context, req *LearnRequest) (*LearnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Learn not implemented")
}
func (*UnimplementedNgramServiceServer) Generate(ctx context.Context, req *GenerateRequest) (*GenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}

func RegisterNgramServiceServer(s *grpc.Server, srv NgramServiceServer) {
	s.RegisterService(&_NgramService_serviceDesc, srv)
}

func _NgramService_Learn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LearnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NgramServiceServer).Learn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.NgramService/Learn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NgramServiceServer).Learn(ctx, req.(*LearnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NgramService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NgramServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1.NgramService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NgramServiceServer).Generate(ctx, req.(*GenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NgramService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v1.NgramService",
	HandlerType: (*NgramServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Learn",
			Handler:    _NgramService_Learn_Handler,
		},
		{
			MethodName: "Generate",
			Handler:    _NgramService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/ngrams.proto",
}
