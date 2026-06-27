package grpc

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
	MeshTelemetry_SubmitReport_FullMethodName         = "/telemetry.MeshTelemetry/SubmitReport"
	MeshTelemetry_SyncBatch_FullMethodName            = "/telemetry.MeshTelemetry/SyncBatch"
	MeshTelemetry_StreamNetworkUpdates_FullMethodName = "/telemetry.MeshTelemetry/StreamNetworkUpdates"
)

// MeshTelemetryClient is the client API for MeshTelemetry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// The core service for Mesh Telemetry
type MeshTelemetryClient interface {
	// Uplink a single report to the cloud
	SubmitReport(ctx context.Context, in *MeshReport, opts ...grpc.CallOption) (*TelemetryResponse, error)
	// Bulk upload cached reports when the user gets back online
	SyncBatch(ctx context.Context, in *ReportBatch, opts ...grpc.CallOption) (*TelemetryResponse, error)
	// Real-time stream for an active node to get network-wide stats
	StreamNetworkUpdates(ctx context.Context, in *NodeIdentity, opts ...grpc.CallOption) (grpc.ServerStreamingClient[NetworkStatus], error)
}

type meshTelemetryClient struct {
	cc grpc.ClientConnInterface
}

func NewMeshTelemetryClient(cc grpc.ClientConnInterface) MeshTelemetryClient {
	return &meshTelemetryClient{cc}
}

func (c *meshTelemetryClient) SubmitReport(ctx context.Context, in *MeshReport, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, MeshTelemetry_SubmitReport_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshTelemetryClient) SyncBatch(ctx context.Context, in *ReportBatch, opts ...grpc.CallOption) (*TelemetryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TelemetryResponse)
	err := c.cc.Invoke(ctx, MeshTelemetry_SyncBatch_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *meshTelemetryClient) StreamNetworkUpdates(ctx context.Context, in *NodeIdentity, opts ...grpc.CallOption) (grpc.ServerStreamingClient[NetworkStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &MeshTelemetry_ServiceDesc.Streams[0], MeshTelemetry_StreamNetworkUpdates_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[NodeIdentity, NetworkStatus]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MeshTelemetry_StreamNetworkUpdatesClient = grpc.ServerStreamingClient[NetworkStatus]

// MeshTelemetryServer is the server API for MeshTelemetry service.
// All implementations must embed UnimplementedMeshTelemetryServer
// for forward compatibility.
//
// The core service for Mesh Telemetry
type MeshTelemetryServer interface {
	// Uplink a single report to the cloud
	SubmitReport(context.Context, *MeshReport) (*TelemetryResponse, error)
	// Bulk upload cached reports when the user gets back online
	SyncBatch(context.Context, *ReportBatch) (*TelemetryResponse, error)
	// Real-time stream for an active node to get network-wide stats
	StreamNetworkUpdates(*NodeIdentity, grpc.ServerStreamingServer[NetworkStatus]) error
	mustEmbedUnimplementedMeshTelemetryServer()
}

// UnimplementedMeshTelemetryServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMeshTelemetryServer struct{}

func (UnimplementedMeshTelemetryServer) SubmitReport(context.Context, *MeshReport) (*TelemetryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SubmitReport not implemented")
}
func (UnimplementedMeshTelemetryServer) SyncBatch(context.Context, *ReportBatch) (*TelemetryResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method SyncBatch not implemented")
}
func (UnimplementedMeshTelemetryServer) StreamNetworkUpdates(*NodeIdentity, grpc.ServerStreamingServer[NetworkStatus]) error {
	return status.Error(codes.Unimplemented, "method StreamNetworkUpdates not implemented")
}
func (UnimplementedMeshTelemetryServer) mustEmbedUnimplementedMeshTelemetryServer() {}
func (UnimplementedMeshTelemetryServer) testEmbeddedByValue()                       {}

// UnsafeMeshTelemetryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MeshTelemetryServer will
// result in compilation errors.
type UnsafeMeshTelemetryServer interface {
	mustEmbedUnimplementedMeshTelemetryServer()
}

func RegisterMeshTelemetryServer(s grpc.ServiceRegistrar, srv MeshTelemetryServer) {
	// If the following call panics, it indicates UnimplementedMeshTelemetryServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MeshTelemetry_ServiceDesc, srv)
}

func _MeshTelemetry_SubmitReport_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeshReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshTelemetryServer).SubmitReport(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeshTelemetry_SubmitReport_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshTelemetryServer).SubmitReport(ctx, req.(*MeshReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshTelemetry_SyncBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReportBatch)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MeshTelemetryServer).SyncBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MeshTelemetry_SyncBatch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MeshTelemetryServer).SyncBatch(ctx, req.(*ReportBatch))
	}
	return interceptor(ctx, in, info, handler)
}

func _MeshTelemetry_StreamNetworkUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(NodeIdentity)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MeshTelemetryServer).StreamNetworkUpdates(m, &grpc.GenericServerStream[NodeIdentity, NetworkStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type MeshTelemetry_StreamNetworkUpdatesServer = grpc.ServerStreamingServer[NetworkStatus]

// MeshTelemetry_ServiceDesc is the grpc.ServiceDesc for MeshTelemetry service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MeshTelemetry_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "telemetry.MeshTelemetry",
	HandlerType: (*MeshTelemetryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitReport",
			Handler:    _MeshTelemetry_SubmitReport_Handler,
		},
		{
			MethodName: "SyncBatch",
			Handler:    _MeshTelemetry_SyncBatch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNetworkUpdates",
			Handler:       _MeshTelemetry_StreamNetworkUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "telemetry.proto",
}
