// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cluster.proto

package api

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

type Cluster_LbStrategy int32

const (
	Cluster_ROUND_ROBIN     Cluster_LbStrategy = 0
	Cluster_LEAST_REQUESTS  Cluster_LbStrategy = 1
	Cluster_CONSISTENT_HASH Cluster_LbStrategy = 2
	Cluster_RANDOM          Cluster_LbStrategy = 3
)

var Cluster_LbStrategy_name = map[int32]string{
	0: "ROUND_ROBIN",
	1: "LEAST_REQUESTS",
	2: "CONSISTENT_HASH",
	3: "RANDOM",
}
var Cluster_LbStrategy_value = map[string]int32{
	"ROUND_ROBIN":     0,
	"LEAST_REQUESTS":  1,
	"CONSISTENT_HASH": 2,
	"RANDOM":          3,
}

func (x Cluster_LbStrategy) String() string {
	return proto.EnumName(Cluster_LbStrategy_name, int32(x))
}
func (Cluster_LbStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cluster_a833cb308b4c1b79, []int{0, 0}
}

type Cluster struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	ConnectTimeout       int32              `protobuf:"varint,2,opt,name=connect_timeout,json=connectTimeout" json:"connect_timeout,omitempty"`
	LbStrategy           Cluster_LbStrategy `protobuf:"varint,3,opt,name=lb_strategy,json=lbStrategy,enum=api.Cluster_LbStrategy" json:"lb_strategy,omitempty"`
	LbConfig             *Cluster_LbConfig  `protobuf:"bytes,4,opt,name=lb_config,json=lbConfig" json:"lb_config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Cluster) Reset()         { *m = Cluster{} }
func (m *Cluster) String() string { return proto.CompactTextString(m) }
func (*Cluster) ProtoMessage()    {}
func (*Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_a833cb308b4c1b79, []int{0}
}
func (m *Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster.Unmarshal(m, b)
}
func (m *Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster.Marshal(b, m, deterministic)
}
func (dst *Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster.Merge(dst, src)
}
func (m *Cluster) XXX_Size() int {
	return xxx_messageInfo_Cluster.Size(m)
}
func (m *Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster proto.InternalMessageInfo

func (m *Cluster) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Cluster) GetConnectTimeout() int32 {
	if m != nil {
		return m.ConnectTimeout
	}
	return 0
}

func (m *Cluster) GetLbStrategy() Cluster_LbStrategy {
	if m != nil {
		return m.LbStrategy
	}
	return Cluster_ROUND_ROBIN
}

func (m *Cluster) GetLbConfig() *Cluster_LbConfig {
	if m != nil {
		return m.LbConfig
	}
	return nil
}

type Cluster_LbConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Cluster_LbConfig) Reset()         { *m = Cluster_LbConfig{} }
func (m *Cluster_LbConfig) String() string { return proto.CompactTextString(m) }
func (*Cluster_LbConfig) ProtoMessage()    {}
func (*Cluster_LbConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_a833cb308b4c1b79, []int{0, 0}
}
func (m *Cluster_LbConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Cluster_LbConfig.Unmarshal(m, b)
}
func (m *Cluster_LbConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Cluster_LbConfig.Marshal(b, m, deterministic)
}
func (dst *Cluster_LbConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Cluster_LbConfig.Merge(dst, src)
}
func (m *Cluster_LbConfig) XXX_Size() int {
	return xxx_messageInfo_Cluster_LbConfig.Size(m)
}
func (m *Cluster_LbConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_Cluster_LbConfig.DiscardUnknown(m)
}

var xxx_messageInfo_Cluster_LbConfig proto.InternalMessageInfo

type ClusterDiscoveryResponse struct {
	Clusters             []*Cluster `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ClusterDiscoveryResponse) Reset()         { *m = ClusterDiscoveryResponse{} }
func (m *ClusterDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*ClusterDiscoveryResponse) ProtoMessage()    {}
func (*ClusterDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cluster_a833cb308b4c1b79, []int{1}
}
func (m *ClusterDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClusterDiscoveryResponse.Unmarshal(m, b)
}
func (m *ClusterDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClusterDiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *ClusterDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClusterDiscoveryResponse.Merge(dst, src)
}
func (m *ClusterDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_ClusterDiscoveryResponse.Size(m)
}
func (m *ClusterDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClusterDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClusterDiscoveryResponse proto.InternalMessageInfo

func (m *ClusterDiscoveryResponse) GetClusters() []*Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func init() {
	proto.RegisterType((*Cluster)(nil), "api.Cluster")
	proto.RegisterType((*Cluster_LbConfig)(nil), "api.Cluster.LbConfig")
	proto.RegisterType((*ClusterDiscoveryResponse)(nil), "api.ClusterDiscoveryResponse")
	proto.RegisterEnum("api.Cluster_LbStrategy", Cluster_LbStrategy_name, Cluster_LbStrategy_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ClusterDiscoveryServiceClient is the client API for ClusterDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClusterDiscoveryServiceClient interface {
	RetrieveClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*ClusterDiscoveryResponse, error)
}

type clusterDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewClusterDiscoveryServiceClient(cc *grpc.ClientConn) ClusterDiscoveryServiceClient {
	return &clusterDiscoveryServiceClient{cc}
}

func (c *clusterDiscoveryServiceClient) RetrieveClusters(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*ClusterDiscoveryResponse, error) {
	out := new(ClusterDiscoveryResponse)
	err := c.cc.Invoke(ctx, "/api.ClusterDiscoveryService/RetrieveClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClusterDiscoveryService service

type ClusterDiscoveryServiceServer interface {
	RetrieveClusters(context.Context, *DiscoveryRequest) (*ClusterDiscoveryResponse, error)
}

func RegisterClusterDiscoveryServiceServer(s *grpc.Server, srv ClusterDiscoveryServiceServer) {
	s.RegisterService(&_ClusterDiscoveryService_serviceDesc, srv)
}

func _ClusterDiscoveryService_RetrieveClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClusterDiscoveryServiceServer).RetrieveClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ClusterDiscoveryService/RetrieveClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClusterDiscoveryServiceServer).RetrieveClusters(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClusterDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ClusterDiscoveryService",
	HandlerType: (*ClusterDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveClusters",
			Handler:    _ClusterDiscoveryService_RetrieveClusters_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cluster.proto",
}

func init() { proto.RegisterFile("cluster.proto", fileDescriptor_cluster_a833cb308b4c1b79) }

var fileDescriptor_cluster_a833cb308b4c1b79 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xc1, 0x6f, 0xaa, 0x40,
	0x10, 0xc6, 0x45, 0x7c, 0x3e, 0x1d, 0xde, 0x13, 0xb2, 0x2f, 0x2f, 0x12, 0x93, 0x26, 0x84, 0x4b,
	0x39, 0x71, 0xa0, 0x97, 0x5e, 0xad, 0x90, 0x68, 0x63, 0x21, 0xdd, 0xc5, 0x33, 0x01, 0x3a, 0x35,
	0x24, 0x08, 0x94, 0x5d, 0x4d, 0xfc, 0x3f, 0xfa, 0x07, 0x37, 0x91, 0xd5, 0x1a, 0x7b, 0xfb, 0xf2,
	0x7d, 0xb3, 0xbf, 0x99, 0x9d, 0x81, 0xbf, 0x79, 0xb9, 0xe7, 0x02, 0x5b, 0xb7, 0x69, 0x6b, 0x51,
	0x13, 0x35, 0x6d, 0x8a, 0x99, 0xfe, 0x56, 0xf0, 0xbc, 0x3e, 0x60, 0x7b, 0xec, 0x5c, 0xfb, 0xb3,
	0x0f, 0xbf, 0x17, 0x5d, 0x1d, 0x21, 0x30, 0xa8, 0xd2, 0x1d, 0x9a, 0x8a, 0xa5, 0x38, 0x63, 0x7a,
	0xd2, 0xe4, 0x1e, 0xf4, 0xbc, 0xae, 0x2a, 0xcc, 0x45, 0x22, 0x8a, 0x1d, 0xd6, 0x7b, 0x61, 0xf6,
	0x2d, 0xc5, 0xf9, 0x45, 0x27, 0xd2, 0x8e, 0x3b, 0x97, 0x3c, 0x82, 0x56, 0x66, 0x09, 0x17, 0x6d,
	0x2a, 0x70, 0x7b, 0x34, 0x55, 0x4b, 0x71, 0x26, 0xde, 0xd4, 0x4d, 0x9b, 0xc2, 0x95, 0x7c, 0x77,
	0x9d, 0x31, 0x19, 0x53, 0x28, 0x2f, 0x9a, 0x78, 0x30, 0x2e, 0xb3, 0x24, 0xaf, 0xab, 0xf7, 0x62,
	0x6b, 0x0e, 0x2c, 0xc5, 0xd1, 0xbc, 0xff, 0x37, 0xef, 0x16, 0xa7, 0x90, 0x8e, 0x4a, 0xa9, 0x66,
	0x00, 0xa3, 0xb3, 0x6b, 0x53, 0x80, 0x6f, 0x32, 0xd1, 0x41, 0xa3, 0xd1, 0x26, 0xf4, 0x13, 0x1a,
	0x3d, 0xad, 0x42, 0xa3, 0x47, 0x08, 0x4c, 0xd6, 0xc1, 0x9c, 0xc5, 0x09, 0x0d, 0x5e, 0x37, 0x01,
	0x8b, 0x99, 0xa1, 0x90, 0x7f, 0xa0, 0x2f, 0xa2, 0x90, 0xad, 0x58, 0x1c, 0x84, 0x71, 0xb2, 0x9c,
	0xb3, 0xa5, 0xd1, 0x27, 0x00, 0x43, 0x3a, 0x0f, 0xfd, 0xe8, 0xc5, 0x50, 0x6d, 0x1f, 0x4c, 0xd9,
	0xdd, 0x3f, 0x2f, 0x8c, 0x22, 0x6f, 0xea, 0x8a, 0x23, 0x71, 0x60, 0x24, 0x37, 0xcb, 0x4d, 0xc5,
	0x52, 0x1d, 0xcd, 0xfb, 0x73, 0x3d, 0x2e, 0xbd, 0xa4, 0x1e, 0xc2, 0xf4, 0x96, 0xc2, 0xb0, 0x3d,
	0x14, 0x39, 0x92, 0x67, 0x30, 0x28, 0x8a, 0xb6, 0xc0, 0x03, 0xca, 0x12, 0x4e, 0xba, 0x5f, 0x5f,
	0x35, 0xfc, 0xd8, 0x23, 0x17, 0xb3, 0xbb, 0x6b, 0xfa, 0x8f, 0x71, 0xec, 0x5e, 0x36, 0x3c, 0x9d,
	0xf2, 0xe1, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x27, 0xb8, 0xa5, 0xf1, 0x01, 0x00, 0x00,
}
