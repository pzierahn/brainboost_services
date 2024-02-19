// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: document_service.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	DocumentService_List_FullMethodName          = "/endpoint.brainboost.documents.v2.DocumentService/List"
	DocumentService_Get_FullMethodName           = "/endpoint.brainboost.documents.v2.DocumentService/Get"
	DocumentService_Delete_FullMethodName        = "/endpoint.brainboost.documents.v2.DocumentService/Delete"
	DocumentService_IndexDocument_FullMethodName = "/endpoint.brainboost.documents.v2.DocumentService/IndexDocument"
	DocumentService_Search_FullMethodName        = "/endpoint.brainboost.documents.v2.DocumentService/Search"
	DocumentService_GetReferences_FullMethodName = "/endpoint.brainboost.documents.v2.DocumentService/GetReferences"
)

// DocumentServiceClient is the client API for DocumentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DocumentServiceClient interface {
	List(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error)
	Get(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*Document, error)
	Delete(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*emptypb.Empty, error)
	IndexDocument(ctx context.Context, in *IndexJob, opts ...grpc.CallOption) (DocumentService_IndexDocumentClient, error)
	Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResults, error)
	GetReferences(ctx context.Context, in *ReferenceIDs, opts ...grpc.CallOption) (*References, error)
}

type documentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDocumentServiceClient(cc grpc.ClientConnInterface) DocumentServiceClient {
	return &documentServiceClient{cc}
}

func (c *documentServiceClient) List(ctx context.Context, in *DocumentFilter, opts ...grpc.CallOption) (*DocumentList, error) {
	out := new(DocumentList)
	err := c.cc.Invoke(ctx, DocumentService_List_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Get(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*Document, error) {
	out := new(Document)
	err := c.cc.Invoke(ctx, DocumentService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) Delete(ctx context.Context, in *DocumentID, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, DocumentService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) IndexDocument(ctx context.Context, in *IndexJob, opts ...grpc.CallOption) (DocumentService_IndexDocumentClient, error) {
	stream, err := c.cc.NewStream(ctx, &DocumentService_ServiceDesc.Streams[0], DocumentService_IndexDocument_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &documentServiceIndexDocumentClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DocumentService_IndexDocumentClient interface {
	Recv() (*IndexProgress, error)
	grpc.ClientStream
}

type documentServiceIndexDocumentClient struct {
	grpc.ClientStream
}

func (x *documentServiceIndexDocumentClient) Recv() (*IndexProgress, error) {
	m := new(IndexProgress)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *documentServiceClient) Search(ctx context.Context, in *SearchQuery, opts ...grpc.CallOption) (*SearchResults, error) {
	out := new(SearchResults)
	err := c.cc.Invoke(ctx, DocumentService_Search_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *documentServiceClient) GetReferences(ctx context.Context, in *ReferenceIDs, opts ...grpc.CallOption) (*References, error) {
	out := new(References)
	err := c.cc.Invoke(ctx, DocumentService_GetReferences_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DocumentServiceServer is the server API for DocumentService service.
// All implementations must embed UnimplementedDocumentServiceServer
// for forward compatibility
type DocumentServiceServer interface {
	List(context.Context, *DocumentFilter) (*DocumentList, error)
	Get(context.Context, *DocumentID) (*Document, error)
	Delete(context.Context, *DocumentID) (*emptypb.Empty, error)
	IndexDocument(*IndexJob, DocumentService_IndexDocumentServer) error
	Search(context.Context, *SearchQuery) (*SearchResults, error)
	GetReferences(context.Context, *ReferenceIDs) (*References, error)
	mustEmbedUnimplementedDocumentServiceServer()
}

// UnimplementedDocumentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDocumentServiceServer struct {
}

func (UnimplementedDocumentServiceServer) List(context.Context, *DocumentFilter) (*DocumentList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDocumentServiceServer) Get(context.Context, *DocumentID) (*Document, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedDocumentServiceServer) Delete(context.Context, *DocumentID) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDocumentServiceServer) IndexDocument(*IndexJob, DocumentService_IndexDocumentServer) error {
	return status.Errorf(codes.Unimplemented, "method IndexDocument not implemented")
}
func (UnimplementedDocumentServiceServer) Search(context.Context, *SearchQuery) (*SearchResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedDocumentServiceServer) GetReferences(context.Context, *ReferenceIDs) (*References, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReferences not implemented")
}
func (UnimplementedDocumentServiceServer) mustEmbedUnimplementedDocumentServiceServer() {}

// UnsafeDocumentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DocumentServiceServer will
// result in compilation errors.
type UnsafeDocumentServiceServer interface {
	mustEmbedUnimplementedDocumentServiceServer()
}

func RegisterDocumentServiceServer(s grpc.ServiceRegistrar, srv DocumentServiceServer) {
	s.RegisterService(&DocumentService_ServiceDesc, srv)
}

func _DocumentService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).List(ctx, req.(*DocumentFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Get(ctx, req.(*DocumentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Delete(ctx, req.(*DocumentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_IndexDocument_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(IndexJob)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DocumentServiceServer).IndexDocument(m, &documentServiceIndexDocumentServer{stream})
}

type DocumentService_IndexDocumentServer interface {
	Send(*IndexProgress) error
	grpc.ServerStream
}

type documentServiceIndexDocumentServer struct {
	grpc.ServerStream
}

func (x *documentServiceIndexDocumentServer) Send(m *IndexProgress) error {
	return x.ServerStream.SendMsg(m)
}

func _DocumentService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).Search(ctx, req.(*SearchQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _DocumentService_GetReferences_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReferenceIDs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DocumentServiceServer).GetReferences(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DocumentService_GetReferences_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DocumentServiceServer).GetReferences(ctx, req.(*ReferenceIDs))
	}
	return interceptor(ctx, in, info, handler)
}

// DocumentService_ServiceDesc is the grpc.ServiceDesc for DocumentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DocumentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "endpoint.brainboost.documents.v2.DocumentService",
	HandlerType: (*DocumentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _DocumentService_List_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _DocumentService_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DocumentService_Delete_Handler,
		},
		{
			MethodName: "Search",
			Handler:    _DocumentService_Search_Handler,
		},
		{
			MethodName: "GetReferences",
			Handler:    _DocumentService_GetReferences_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "IndexDocument",
			Handler:       _DocumentService_IndexDocument_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "document_service.proto",
}
