// Code generated by protoc-gen-go.
// source: cluster.proto
// DO NOT EDIT!

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	cluster.proto

It has these top-level messages:
	ClusterRequest
	ClusterReply
*/
package protobuf

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

type ClusterRequest struct {
}

func (m *ClusterRequest) Reset()                    { *m = ClusterRequest{} }
func (m *ClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*ClusterRequest) ProtoMessage()               {}
func (*ClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClusterReply struct {
}

func (m *ClusterReply) Reset()                    { *m = ClusterReply{} }
func (m *ClusterReply) String() string            { return proto.CompactTextString(m) }
func (*ClusterReply) ProtoMessage()               {}
func (*ClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*ClusterRequest)(nil), "protobuf.ClusterRequest")
	proto.RegisterType((*ClusterReply)(nil), "protobuf.ClusterReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cluster service

type ClusterClient interface {
	Join(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error)
	Leave(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error)
}

type clusterClient struct {
	cc *grpc.ClientConn
}

func NewClusterClient(cc *grpc.ClientConn) ClusterClient {
	return &clusterClient{cc}
}

func (c *clusterClient) Join(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := grpc.Invoke(ctx, "/protobuf.Cluster/Join", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clusterClient) Leave(ctx context.Context, in *ClusterRequest, opts ...grpc.CallOption) (*ClusterReply, error) {
	out := new(ClusterReply)
	err := grpc.Invoke(ctx, "/protobuf.Cluster/Leave", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cluster service

type ClusterServer interface {
	Join(context.Context, *ClusterRequest) (*ClusterReply, error)
	Leave(context.Context, *ClusterRequest) (*ClusterReply, error)
}

func RegisterClusterServer(s *grpc.Server, srv ClusterServer) {
	s.RegisterService(&_Cluster_serviceDesc, srv)
}

func _Cluster_Join_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Join(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Cluster/Join",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Join(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cluster_Leave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterServer).Leave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Cluster/Leave",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterServer).Leave(ctx, req.(*ClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cluster_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Cluster",
	HandlerType: (*ClusterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Join",
			Handler:    _Cluster_Join_Handler,
		},
		{
			MethodName: "Leave",
			Handler:    _Cluster_Leave_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x29, 0x2d,
	0x2e, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69,
	0x4a, 0x02, 0x5c, 0x7c, 0xce, 0x10, 0xa9, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x25, 0x3e,
	0x2e, 0x1e, 0xb8, 0x48, 0x41, 0x4e, 0xa5, 0x51, 0x13, 0x23, 0x17, 0x3b, 0x54, 0x40, 0xc8, 0x8a,
	0x8b, 0xc5, 0x2b, 0x3f, 0x33, 0x4f, 0x48, 0x42, 0x0f, 0x66, 0x80, 0x1e, 0xaa, 0x6e, 0x29, 0x31,
	0x2c, 0x32, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x42, 0xd6, 0x5c, 0xac, 0x3e, 0xa9, 0x89, 0x65, 0xa9,
	0xe4, 0x68, 0x4e, 0x62, 0x03, 0x4b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x25, 0xa4, 0xa7,
	0x19, 0xc8, 0x00, 0x00, 0x00,
}
