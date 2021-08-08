// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package chatapipb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ChatapiClient is the client API for Chatapi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatapiClient interface {
	// Getchat implements getchat.
	Getchat(ctx context.Context, in *GetchatRequest, opts ...grpc.CallOption) (*GoaChatCollection, error)
	// Postchat implements postchat.
	Postchat(ctx context.Context, in *PostchatRequest, opts ...grpc.CallOption) (*PostchatResponse, error)
}

type chatapiClient struct {
	cc grpc.ClientConnInterface
}

func NewChatapiClient(cc grpc.ClientConnInterface) ChatapiClient {
	return &chatapiClient{cc}
}

func (c *chatapiClient) Getchat(ctx context.Context, in *GetchatRequest, opts ...grpc.CallOption) (*GoaChatCollection, error) {
	out := new(GoaChatCollection)
	err := c.cc.Invoke(ctx, "/chatapi.Chatapi/Getchat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatapiClient) Postchat(ctx context.Context, in *PostchatRequest, opts ...grpc.CallOption) (*PostchatResponse, error) {
	out := new(PostchatResponse)
	err := c.cc.Invoke(ctx, "/chatapi.Chatapi/Postchat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChatapiServer is the server API for Chatapi service.
// All implementations must embed UnimplementedChatapiServer
// for forward compatibility
type ChatapiServer interface {
	// Getchat implements getchat.
	Getchat(context.Context, *GetchatRequest) (*GoaChatCollection, error)
	// Postchat implements postchat.
	Postchat(context.Context, *PostchatRequest) (*PostchatResponse, error)
	mustEmbedUnimplementedChatapiServer()
}

// UnimplementedChatapiServer must be embedded to have forward compatible implementations.
type UnimplementedChatapiServer struct {
}

func (UnimplementedChatapiServer) Getchat(context.Context, *GetchatRequest) (*GoaChatCollection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Getchat not implemented")
}
func (UnimplementedChatapiServer) Postchat(context.Context, *PostchatRequest) (*PostchatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Postchat not implemented")
}
func (UnimplementedChatapiServer) mustEmbedUnimplementedChatapiServer() {}

// UnsafeChatapiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatapiServer will
// result in compilation errors.
type UnsafeChatapiServer interface {
	mustEmbedUnimplementedChatapiServer()
}

func RegisterChatapiServer(s grpc.ServiceRegistrar, srv ChatapiServer) {
	s.RegisterService(&Chatapi_ServiceDesc, srv)
}

func _Chatapi_Getchat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetchatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatapiServer).Getchat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi.Chatapi/Getchat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatapiServer).Getchat(ctx, req.(*GetchatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Chatapi_Postchat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostchatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatapiServer).Postchat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chatapi.Chatapi/Postchat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatapiServer).Postchat(ctx, req.(*PostchatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Chatapi_ServiceDesc is the grpc.ServiceDesc for Chatapi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chatapi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chatapi.Chatapi",
	HandlerType: (*ChatapiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Getchat",
			Handler:    _Chatapi_Getchat_Handler,
		},
		{
			MethodName: "Postchat",
			Handler:    _Chatapi_Postchat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "chatapi.proto",
}
