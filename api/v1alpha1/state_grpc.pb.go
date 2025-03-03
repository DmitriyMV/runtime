// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.24.4
// source: v1alpha1/state.proto

package v1alpha1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	State_Get_FullMethodName     = "/cosi.resource.State/Get"
	State_List_FullMethodName    = "/cosi.resource.State/List"
	State_Create_FullMethodName  = "/cosi.resource.State/Create"
	State_Update_FullMethodName  = "/cosi.resource.State/Update"
	State_Destroy_FullMethodName = "/cosi.resource.State/Destroy"
	State_Watch_FullMethodName   = "/cosi.resource.State/Watch"
)

// StateClient is the client API for State service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateClient interface {
	// Get a resource by type and ID.
	//
	// If a resource is not found, error is returned.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// List resources by type.
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListResponse], error)
	// Create a resource.
	//
	// If a resource already exists, Create returns an error.
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	// Update a resource.
	//
	// If a resource doesn't exist, error is returned.
	// On update current version of resource `new` in the state should match
	// curVersion, otherwise conflict error is returned.
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	// Destroy a resource.
	//
	// If a resource doesn't exist, error is returned.
	// If a resource has pending finalizers, error is returned.
	Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error)
	// Watch state of a resource by (namespace, type) or a specific resource by (namespace, type, id).
	//
	// It's fine to watch for a resource which doesn't exist yet.
	// Watch is canceled when context gets canceled.
	// Watch sends initial resource state as the very first event on the channel,
	// and then sends any updates to the resource as events.
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchResponse], error)
}

type stateClient struct {
	cc grpc.ClientConnInterface
}

func NewStateClient(cc grpc.ClientConnInterface) StateClient {
	return &stateClient{cc}
}

func (c *stateClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, State_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ListResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &State_ServiceDesc.Streams[0], State_List_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ListRequest, ListResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type State_ListClient = grpc.ServerStreamingClient[ListResponse]

func (c *stateClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, State_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, State_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) Destroy(ctx context.Context, in *DestroyRequest, opts ...grpc.CallOption) (*DestroyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DestroyResponse)
	err := c.cc.Invoke(ctx, State_Destroy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[WatchResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &State_ServiceDesc.Streams[1], State_Watch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[WatchRequest, WatchResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type State_WatchClient = grpc.ServerStreamingClient[WatchResponse]

// StateServer is the server API for State service.
// All implementations must embed UnimplementedStateServer
// for forward compatibility.
type StateServer interface {
	// Get a resource by type and ID.
	//
	// If a resource is not found, error is returned.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// List resources by type.
	List(*ListRequest, grpc.ServerStreamingServer[ListResponse]) error
	// Create a resource.
	//
	// If a resource already exists, Create returns an error.
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	// Update a resource.
	//
	// If a resource doesn't exist, error is returned.
	// On update current version of resource `new` in the state should match
	// curVersion, otherwise conflict error is returned.
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	// Destroy a resource.
	//
	// If a resource doesn't exist, error is returned.
	// If a resource has pending finalizers, error is returned.
	Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error)
	// Watch state of a resource by (namespace, type) or a specific resource by (namespace, type, id).
	//
	// It's fine to watch for a resource which doesn't exist yet.
	// Watch is canceled when context gets canceled.
	// Watch sends initial resource state as the very first event on the channel,
	// and then sends any updates to the resource as events.
	Watch(*WatchRequest, grpc.ServerStreamingServer[WatchResponse]) error
	mustEmbedUnimplementedStateServer()
}

// UnimplementedStateServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedStateServer struct{}

func (UnimplementedStateServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStateServer) List(*ListRequest, grpc.ServerStreamingServer[ListResponse]) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedStateServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedStateServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedStateServer) Destroy(context.Context, *DestroyRequest) (*DestroyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Destroy not implemented")
}
func (UnimplementedStateServer) Watch(*WatchRequest, grpc.ServerStreamingServer[WatchResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedStateServer) mustEmbedUnimplementedStateServer() {}
func (UnimplementedStateServer) testEmbeddedByValue()               {}

// UnsafeStateServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateServer will
// result in compilation errors.
type UnsafeStateServer interface {
	mustEmbedUnimplementedStateServer()
}

func RegisterStateServer(s grpc.ServiceRegistrar, srv StateServer) {
	// If the following call pancis, it indicates UnimplementedStateServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&State_ServiceDesc, srv)
}

func _State_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: State_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StateServer).List(m, &grpc.GenericServerStream[ListRequest, ListResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type State_ListServer = grpc.ServerStreamingServer[ListResponse]

func _State_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: State_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: State_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_Destroy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateServer).Destroy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: State_Destroy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateServer).Destroy(ctx, req.(*DestroyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _State_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StateServer).Watch(m, &grpc.GenericServerStream[WatchRequest, WatchResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type State_WatchServer = grpc.ServerStreamingServer[WatchResponse]

// State_ServiceDesc is the grpc.ServiceDesc for State service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var State_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cosi.resource.State",
	HandlerType: (*StateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _State_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _State_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _State_Update_Handler,
		},
		{
			MethodName: "Destroy",
			Handler:    _State_Destroy_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _State_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Watch",
			Handler:       _State_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "v1alpha1/state.proto",
}
