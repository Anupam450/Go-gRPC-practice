// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blog/blogpb/blog.proto

package blogpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Blog struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthorId             string   `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title                string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blog) Reset()         { *m = Blog{} }
func (m *Blog) String() string { return proto.CompactTextString(m) }
func (*Blog) ProtoMessage()    {}
func (*Blog) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{0}
}
func (m *Blog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blog.Unmarshal(m, b)
}
func (m *Blog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blog.Marshal(b, m, deterministic)
}
func (dst *Blog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blog.Merge(dst, src)
}
func (m *Blog) XXX_Size() int {
	return xxx_messageInfo_Blog.Size(m)
}
func (m *Blog) XXX_DiscardUnknown() {
	xxx_messageInfo_Blog.DiscardUnknown(m)
}

var xxx_messageInfo_Blog proto.InternalMessageInfo

func (m *Blog) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Blog) GetAuthorId() string {
	if m != nil {
		return m.AuthorId
	}
	return ""
}

func (m *Blog) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Blog) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type CreateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogRequest) Reset()         { *m = CreateBlogRequest{} }
func (m *CreateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*CreateBlogRequest) ProtoMessage()    {}
func (*CreateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{1}
}
func (m *CreateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogRequest.Unmarshal(m, b)
}
func (m *CreateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogRequest.Marshal(b, m, deterministic)
}
func (dst *CreateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogRequest.Merge(dst, src)
}
func (m *CreateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_CreateBlogRequest.Size(m)
}
func (m *CreateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogRequest proto.InternalMessageInfo

func (m *CreateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type CreateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateBlogResponse) Reset()         { *m = CreateBlogResponse{} }
func (m *CreateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*CreateBlogResponse) ProtoMessage()    {}
func (*CreateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{2}
}
func (m *CreateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateBlogResponse.Unmarshal(m, b)
}
func (m *CreateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateBlogResponse.Marshal(b, m, deterministic)
}
func (dst *CreateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateBlogResponse.Merge(dst, src)
}
func (m *CreateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_CreateBlogResponse.Size(m)
}
func (m *CreateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateBlogResponse proto.InternalMessageInfo

func (m *CreateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type ReadBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogRequest) Reset()         { *m = ReadBlogRequest{} }
func (m *ReadBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ReadBlogRequest) ProtoMessage()    {}
func (*ReadBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{3}
}
func (m *ReadBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogRequest.Unmarshal(m, b)
}
func (m *ReadBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogRequest.Marshal(b, m, deterministic)
}
func (dst *ReadBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogRequest.Merge(dst, src)
}
func (m *ReadBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ReadBlogRequest.Size(m)
}
func (m *ReadBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogRequest proto.InternalMessageInfo

func (m *ReadBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ReadBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadBlogResponse) Reset()         { *m = ReadBlogResponse{} }
func (m *ReadBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ReadBlogResponse) ProtoMessage()    {}
func (*ReadBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{4}
}
func (m *ReadBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadBlogResponse.Unmarshal(m, b)
}
func (m *ReadBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadBlogResponse.Marshal(b, m, deterministic)
}
func (dst *ReadBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadBlogResponse.Merge(dst, src)
}
func (m *ReadBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ReadBlogResponse.Size(m)
}
func (m *ReadBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadBlogResponse proto.InternalMessageInfo

func (m *ReadBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogRequest struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogRequest) Reset()         { *m = UpdateBlogRequest{} }
func (m *UpdateBlogRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogRequest) ProtoMessage()    {}
func (*UpdateBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{5}
}
func (m *UpdateBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogRequest.Unmarshal(m, b)
}
func (m *UpdateBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogRequest.Merge(dst, src)
}
func (m *UpdateBlogRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogRequest.Size(m)
}
func (m *UpdateBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogRequest proto.InternalMessageInfo

func (m *UpdateBlogRequest) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type UpdateBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateBlogResponse) Reset()         { *m = UpdateBlogResponse{} }
func (m *UpdateBlogResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateBlogResponse) ProtoMessage()    {}
func (*UpdateBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{6}
}
func (m *UpdateBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateBlogResponse.Unmarshal(m, b)
}
func (m *UpdateBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateBlogResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateBlogResponse.Merge(dst, src)
}
func (m *UpdateBlogResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateBlogResponse.Size(m)
}
func (m *UpdateBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateBlogResponse proto.InternalMessageInfo

func (m *UpdateBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

type DeleteBlogRequest struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogRequest) Reset()         { *m = DeleteBlogRequest{} }
func (m *DeleteBlogRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogRequest) ProtoMessage()    {}
func (*DeleteBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{7}
}
func (m *DeleteBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogRequest.Unmarshal(m, b)
}
func (m *DeleteBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogRequest.Marshal(b, m, deterministic)
}
func (dst *DeleteBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogRequest.Merge(dst, src)
}
func (m *DeleteBlogRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogRequest.Size(m)
}
func (m *DeleteBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogRequest proto.InternalMessageInfo

func (m *DeleteBlogRequest) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type DeleteBlogResponse struct {
	BlogId               string   `protobuf:"bytes,1,opt,name=blog_id,json=blogId,proto3" json:"blog_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteBlogResponse) Reset()         { *m = DeleteBlogResponse{} }
func (m *DeleteBlogResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteBlogResponse) ProtoMessage()    {}
func (*DeleteBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{8}
}
func (m *DeleteBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteBlogResponse.Unmarshal(m, b)
}
func (m *DeleteBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteBlogResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteBlogResponse.Merge(dst, src)
}
func (m *DeleteBlogResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteBlogResponse.Size(m)
}
func (m *DeleteBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteBlogResponse proto.InternalMessageInfo

func (m *DeleteBlogResponse) GetBlogId() string {
	if m != nil {
		return m.BlogId
	}
	return ""
}

type ListBlogRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogRequest) Reset()         { *m = ListBlogRequest{} }
func (m *ListBlogRequest) String() string { return proto.CompactTextString(m) }
func (*ListBlogRequest) ProtoMessage()    {}
func (*ListBlogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{9}
}
func (m *ListBlogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogRequest.Unmarshal(m, b)
}
func (m *ListBlogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogRequest.Marshal(b, m, deterministic)
}
func (dst *ListBlogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogRequest.Merge(dst, src)
}
func (m *ListBlogRequest) XXX_Size() int {
	return xxx_messageInfo_ListBlogRequest.Size(m)
}
func (m *ListBlogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogRequest proto.InternalMessageInfo

type ListBlogResponse struct {
	Blog                 *Blog    `protobuf:"bytes,1,opt,name=blog,proto3" json:"blog,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListBlogResponse) Reset()         { *m = ListBlogResponse{} }
func (m *ListBlogResponse) String() string { return proto.CompactTextString(m) }
func (*ListBlogResponse) ProtoMessage()    {}
func (*ListBlogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_blog_f36092d2f7e14931, []int{10}
}
func (m *ListBlogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListBlogResponse.Unmarshal(m, b)
}
func (m *ListBlogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListBlogResponse.Marshal(b, m, deterministic)
}
func (dst *ListBlogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListBlogResponse.Merge(dst, src)
}
func (m *ListBlogResponse) XXX_Size() int {
	return xxx_messageInfo_ListBlogResponse.Size(m)
}
func (m *ListBlogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListBlogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListBlogResponse proto.InternalMessageInfo

func (m *ListBlogResponse) GetBlog() *Blog {
	if m != nil {
		return m.Blog
	}
	return nil
}

func init() {
	proto.RegisterType((*Blog)(nil), "blog.Blog")
	proto.RegisterType((*CreateBlogRequest)(nil), "blog.CreateBlogRequest")
	proto.RegisterType((*CreateBlogResponse)(nil), "blog.CreateBlogResponse")
	proto.RegisterType((*ReadBlogRequest)(nil), "blog.ReadBlogRequest")
	proto.RegisterType((*ReadBlogResponse)(nil), "blog.ReadBlogResponse")
	proto.RegisterType((*UpdateBlogRequest)(nil), "blog.UpdateBlogRequest")
	proto.RegisterType((*UpdateBlogResponse)(nil), "blog.UpdateBlogResponse")
	proto.RegisterType((*DeleteBlogRequest)(nil), "blog.DeleteBlogRequest")
	proto.RegisterType((*DeleteBlogResponse)(nil), "blog.DeleteBlogResponse")
	proto.RegisterType((*ListBlogRequest)(nil), "blog.ListBlogRequest")
	proto.RegisterType((*ListBlogResponse)(nil), "blog.ListBlogResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BlogServiceClient is the client API for BlogService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogServiceClient interface {
	CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error)
	ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error)
	UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error)
	DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error)
	ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error)
}

type blogServiceClient struct {
	cc *grpc.ClientConn
}

func NewBlogServiceClient(cc *grpc.ClientConn) BlogServiceClient {
	return &blogServiceClient{cc}
}

func (c *blogServiceClient) CreateBlog(ctx context.Context, in *CreateBlogRequest, opts ...grpc.CallOption) (*CreateBlogResponse, error) {
	out := new(CreateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/CreateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ReadBlog(ctx context.Context, in *ReadBlogRequest, opts ...grpc.CallOption) (*ReadBlogResponse, error) {
	out := new(ReadBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/ReadBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) UpdateBlog(ctx context.Context, in *UpdateBlogRequest, opts ...grpc.CallOption) (*UpdateBlogResponse, error) {
	out := new(UpdateBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/UpdateBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) DeleteBlog(ctx context.Context, in *DeleteBlogRequest, opts ...grpc.CallOption) (*DeleteBlogResponse, error) {
	out := new(DeleteBlogResponse)
	err := c.cc.Invoke(ctx, "/blog.BlogService/DeleteBlog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blogServiceClient) ListBlog(ctx context.Context, in *ListBlogRequest, opts ...grpc.CallOption) (BlogService_ListBlogClient, error) {
	stream, err := c.cc.NewStream(ctx, &_BlogService_serviceDesc.Streams[0], "/blog.BlogService/ListBlog", opts...)
	if err != nil {
		return nil, err
	}
	x := &blogServiceListBlogClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BlogService_ListBlogClient interface {
	Recv() (*ListBlogResponse, error)
	grpc.ClientStream
}

type blogServiceListBlogClient struct {
	grpc.ClientStream
}

func (x *blogServiceListBlogClient) Recv() (*ListBlogResponse, error) {
	m := new(ListBlogResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlogServiceServer is the server API for BlogService service.
type BlogServiceServer interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*CreateBlogResponse, error)
	ReadBlog(context.Context, *ReadBlogRequest) (*ReadBlogResponse, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*UpdateBlogResponse, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) (*DeleteBlogResponse, error)
	ListBlog(*ListBlogRequest, BlogService_ListBlogServer) error
}

func RegisterBlogServiceServer(s *grpc.Server, srv BlogServiceServer) {
	s.RegisterService(&_BlogService_serviceDesc, srv)
}

func _BlogService_CreateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).CreateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/CreateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).CreateBlog(ctx, req.(*CreateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ReadBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).ReadBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/ReadBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).ReadBlog(ctx, req.(*ReadBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_UpdateBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).UpdateBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/UpdateBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).UpdateBlog(ctx, req.(*UpdateBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_DeleteBlog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBlogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogServiceServer).DeleteBlog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blog.BlogService/DeleteBlog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogServiceServer).DeleteBlog(ctx, req.(*DeleteBlogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlogService_ListBlog_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListBlogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BlogServiceServer).ListBlog(m, &blogServiceListBlogServer{stream})
}

type BlogService_ListBlogServer interface {
	Send(*ListBlogResponse) error
	grpc.ServerStream
}

type blogServiceListBlogServer struct {
	grpc.ServerStream
}

func (x *blogServiceListBlogServer) Send(m *ListBlogResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _BlogService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blog.BlogService",
	HandlerType: (*BlogServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlog",
			Handler:    _BlogService_CreateBlog_Handler,
		},
		{
			MethodName: "ReadBlog",
			Handler:    _BlogService_ReadBlog_Handler,
		},
		{
			MethodName: "UpdateBlog",
			Handler:    _BlogService_UpdateBlog_Handler,
		},
		{
			MethodName: "DeleteBlog",
			Handler:    _BlogService_DeleteBlog_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListBlog",
			Handler:       _BlogService_ListBlog_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "blog/blogpb/blog.proto",
}

func init() { proto.RegisterFile("blog/blogpb/blog.proto", fileDescriptor_blog_f36092d2f7e14931) }

var fileDescriptor_blog_f36092d2f7e14931 = []byte{
	// 352 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x41, 0x4f, 0xf2, 0x40,
	0x14, 0x0c, 0xfd, 0xf8, 0xa0, 0x0c, 0x89, 0xd8, 0x8d, 0x42, 0x83, 0x89, 0x31, 0x9c, 0x8c, 0x51,
	0x34, 0xe0, 0xc5, 0x78, 0x20, 0x41, 0x2f, 0x24, 0x9e, 0x30, 0x5e, 0xbc, 0x10, 0x60, 0x5f, 0xb0,
	0x49, 0xc3, 0xd6, 0x76, 0xf1, 0x37, 0xfa, 0xb3, 0xcc, 0xee, 0xb2, 0xb6, 0xb6, 0x10, 0xeb, 0x05,
	0xfa, 0xde, 0xce, 0x7b, 0x33, 0x3b, 0x93, 0x45, 0x7b, 0x11, 0x8a, 0xd5, 0xb5, 0xfa, 0x89, 0x16,
	0xfa, 0xaf, 0x1f, 0xc5, 0x42, 0x0a, 0x56, 0x55, 0xdf, 0xbd, 0x25, 0xaa, 0xe3, 0x50, 0xac, 0xd8,
	0x01, 0x9c, 0x80, 0xfb, 0x95, 0xb3, 0xca, 0x79, 0x63, 0xea, 0x04, 0x9c, 0x9d, 0xa0, 0x31, 0xdf,
	0xc8, 0x37, 0x11, 0xcf, 0x02, 0xee, 0x3b, 0xba, 0xed, 0x9a, 0xc6, 0x84, 0xb3, 0x23, 0xfc, 0x97,
	0x81, 0x0c, 0xc9, 0xff, 0xa7, 0x0f, 0x4c, 0xc1, 0x7c, 0xd4, 0x97, 0x62, 0x2d, 0x69, 0x2d, 0xfd,
	0xaa, 0xee, 0xdb, 0xb2, 0x37, 0x84, 0xf7, 0x10, 0xd3, 0x5c, 0x92, 0xa2, 0x9a, 0xd2, 0xfb, 0x86,
	0x12, 0xc9, 0x4e, 0xa1, 0x15, 0x68, 0xce, 0xe6, 0x00, 0x7d, 0x2d, 0x4d, 0x03, 0x8c, 0xb2, 0x5b,
	0xb0, 0xec, 0x50, 0x12, 0x89, 0x75, 0x42, 0xbf, 0x4e, 0x5d, 0xa0, 0x35, 0xa5, 0x39, 0xcf, 0x12,
	0x75, 0x50, 0x57, 0x47, 0xb3, 0xef, 0xfb, 0xd5, 0x54, 0x39, 0xe1, 0xbd, 0x01, 0x0e, 0x53, 0x6c,
	0xc9, 0xfd, 0x43, 0x78, 0x2f, 0x11, 0xff, 0xfb, 0x55, 0xb2, 0x43, 0x25, 0xa9, 0x2e, 0xe1, 0x3d,
	0x52, 0x48, 0x3f, 0xa9, 0xf6, 0x5e, 0xe6, 0x0a, 0x2c, 0x8b, 0xde, 0x72, 0xec, 0x85, 0x7b, 0x68,
	0x3d, 0x05, 0x89, 0xcc, 0xac, 0x56, 0x76, 0xa4, 0xad, 0x72, 0x1a, 0x07, 0x9f, 0x0e, 0x9a, 0xaa,
	0x7c, 0xa6, 0xf8, 0x23, 0x58, 0x12, 0x1b, 0x01, 0x69, 0x68, 0xac, 0x63, 0xf0, 0x85, 0xec, 0xbb,
	0x7e, 0xf1, 0x60, 0x4b, 0x78, 0x07, 0xd7, 0x66, 0xc2, 0x8e, 0x0d, 0x2a, 0x97, 0x67, 0xb7, 0x9d,
	0x6f, 0x6f, 0x47, 0x47, 0x40, 0xea, 0xb2, 0xe5, 0x2e, 0x84, 0x65, 0xb9, 0x77, 0x04, 0x32, 0x02,
	0x52, 0x0b, 0xed, 0x82, 0x42, 0x04, 0x76, 0xc1, 0x0e, 0xb7, 0xef, 0xe1, 0x5a, 0x07, 0xad, 0xf8,
	0x9c, 0xc9, 0x56, 0x7c, 0xde, 0xe8, 0x9b, 0xca, 0xd8, 0x7d, 0xad, 0x99, 0x47, 0xba, 0xa8, 0xe9,
	0x07, 0x3a, 0xfc, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xe4, 0xa6, 0x7b, 0x72, 0xba, 0x03, 0x00, 0x00,
}
