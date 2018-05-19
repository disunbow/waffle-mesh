// Code generated by protoc-gen-go. DO NOT EDIT.
// source: route.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"

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

type RouteConfig struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	RouteDetails         []*RouteDetails `protobuf:"bytes,2,rep,name=route_details,json=routeDetails" json:"route_details,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *RouteConfig) Reset()         { *m = RouteConfig{} }
func (m *RouteConfig) String() string { return proto.CompactTextString(m) }
func (*RouteConfig) ProtoMessage()    {}
func (*RouteConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{0}
}
func (m *RouteConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteConfig.Unmarshal(m, b)
}
func (m *RouteConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteConfig.Marshal(b, m, deterministic)
}
func (dst *RouteConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteConfig.Merge(dst, src)
}
func (m *RouteConfig) XXX_Size() int {
	return xxx_messageInfo_RouteConfig.Size(m)
}
func (m *RouteConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RouteConfig proto.InternalMessageInfo

func (m *RouteConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteConfig) GetRouteDetails() []*RouteDetails {
	if m != nil {
		return m.RouteDetails
	}
	return nil
}

type RouteDetails struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Domains              []string      `protobuf:"bytes,2,rep,name=domains" json:"domains,omitempty"`
	Routes               []*RouteEntry `protobuf:"bytes,3,rep,name=routes" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *RouteDetails) Reset()         { *m = RouteDetails{} }
func (m *RouteDetails) String() string { return proto.CompactTextString(m) }
func (*RouteDetails) ProtoMessage()    {}
func (*RouteDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{1}
}
func (m *RouteDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDetails.Unmarshal(m, b)
}
func (m *RouteDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDetails.Marshal(b, m, deterministic)
}
func (dst *RouteDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDetails.Merge(dst, src)
}
func (m *RouteDetails) XXX_Size() int {
	return xxx_messageInfo_RouteDetails.Size(m)
}
func (m *RouteDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDetails.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDetails proto.InternalMessageInfo

func (m *RouteDetails) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RouteDetails) GetDomains() []string {
	if m != nil {
		return m.Domains
	}
	return nil
}

func (m *RouteDetails) GetRoutes() []*RouteEntry {
	if m != nil {
		return m.Routes
	}
	return nil
}

type RouteEntry struct {
	Match *RouteMatch `protobuf:"bytes,1,opt,name=match" json:"match,omitempty"`
	// Types that are valid to be assigned to Action:
	//	*RouteEntry_Route
	Action               isRouteEntry_Action `protobuf_oneof:"action"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RouteEntry) Reset()         { *m = RouteEntry{} }
func (m *RouteEntry) String() string { return proto.CompactTextString(m) }
func (*RouteEntry) ProtoMessage()    {}
func (*RouteEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{2}
}
func (m *RouteEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteEntry.Unmarshal(m, b)
}
func (m *RouteEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteEntry.Marshal(b, m, deterministic)
}
func (dst *RouteEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteEntry.Merge(dst, src)
}
func (m *RouteEntry) XXX_Size() int {
	return xxx_messageInfo_RouteEntry.Size(m)
}
func (m *RouteEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RouteEntry proto.InternalMessageInfo

type isRouteEntry_Action interface {
	isRouteEntry_Action()
}

type RouteEntry_Route struct {
	Route *RouteAction `protobuf:"bytes,2,opt,name=route,oneof"`
}

func (*RouteEntry_Route) isRouteEntry_Action() {}

func (m *RouteEntry) GetAction() isRouteEntry_Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *RouteEntry) GetMatch() *RouteMatch {
	if m != nil {
		return m.Match
	}
	return nil
}

func (m *RouteEntry) GetRoute() *RouteAction {
	if x, ok := m.GetAction().(*RouteEntry_Route); ok {
		return x.Route
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RouteEntry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RouteEntry_OneofMarshaler, _RouteEntry_OneofUnmarshaler, _RouteEntry_OneofSizer, []interface{}{
		(*RouteEntry_Route)(nil),
	}
}

func _RouteEntry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RouteEntry)
	// action
	switch x := m.Action.(type) {
	case *RouteEntry_Route:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Route); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RouteEntry.Action has unexpected type %T", x)
	}
	return nil
}

func _RouteEntry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RouteEntry)
	switch tag {
	case 2: // action.route
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(RouteAction)
		err := b.DecodeMessage(msg)
		m.Action = &RouteEntry_Route{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RouteEntry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RouteEntry)
	// action
	switch x := m.Action.(type) {
	case *RouteEntry_Route:
		s := proto.Size(x.Route)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RouteMatch struct {
	// Types that are valid to be assigned to PathPattern:
	//	*RouteMatch_Prefix
	//	*RouteMatch_ExactPath
	//	*RouteMatch_Regex
	PathPattern          isRouteMatch_PathPattern `protobuf_oneof:"path_pattern"`
	Headers              []*HeaderMatch           `protobuf:"bytes,4,rep,name=headers" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *RouteMatch) Reset()         { *m = RouteMatch{} }
func (m *RouteMatch) String() string { return proto.CompactTextString(m) }
func (*RouteMatch) ProtoMessage()    {}
func (*RouteMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{3}
}
func (m *RouteMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteMatch.Unmarshal(m, b)
}
func (m *RouteMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteMatch.Marshal(b, m, deterministic)
}
func (dst *RouteMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteMatch.Merge(dst, src)
}
func (m *RouteMatch) XXX_Size() int {
	return xxx_messageInfo_RouteMatch.Size(m)
}
func (m *RouteMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteMatch.DiscardUnknown(m)
}

var xxx_messageInfo_RouteMatch proto.InternalMessageInfo

type isRouteMatch_PathPattern interface {
	isRouteMatch_PathPattern()
}

type RouteMatch_Prefix struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix,oneof"`
}
type RouteMatch_ExactPath struct {
	ExactPath string `protobuf:"bytes,2,opt,name=exact_path,json=exactPath,oneof"`
}
type RouteMatch_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,oneof"`
}

func (*RouteMatch_Prefix) isRouteMatch_PathPattern()    {}
func (*RouteMatch_ExactPath) isRouteMatch_PathPattern() {}
func (*RouteMatch_Regex) isRouteMatch_PathPattern()     {}

func (m *RouteMatch) GetPathPattern() isRouteMatch_PathPattern {
	if m != nil {
		return m.PathPattern
	}
	return nil
}

func (m *RouteMatch) GetPrefix() string {
	if x, ok := m.GetPathPattern().(*RouteMatch_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *RouteMatch) GetExactPath() string {
	if x, ok := m.GetPathPattern().(*RouteMatch_ExactPath); ok {
		return x.ExactPath
	}
	return ""
}

func (m *RouteMatch) GetRegex() string {
	if x, ok := m.GetPathPattern().(*RouteMatch_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *RouteMatch) GetHeaders() []*HeaderMatch {
	if m != nil {
		return m.Headers
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RouteMatch) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RouteMatch_OneofMarshaler, _RouteMatch_OneofUnmarshaler, _RouteMatch_OneofSizer, []interface{}{
		(*RouteMatch_Prefix)(nil),
		(*RouteMatch_ExactPath)(nil),
		(*RouteMatch_Regex)(nil),
	}
}

func _RouteMatch_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RouteMatch)
	// path_pattern
	switch x := m.PathPattern.(type) {
	case *RouteMatch_Prefix:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Prefix)
	case *RouteMatch_ExactPath:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ExactPath)
	case *RouteMatch_Regex:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Regex)
	case nil:
	default:
		return fmt.Errorf("RouteMatch.PathPattern has unexpected type %T", x)
	}
	return nil
}

func _RouteMatch_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RouteMatch)
	switch tag {
	case 1: // path_pattern.prefix
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PathPattern = &RouteMatch_Prefix{x}
		return true, err
	case 2: // path_pattern.exact_path
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PathPattern = &RouteMatch_ExactPath{x}
		return true, err
	case 3: // path_pattern.regex
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.PathPattern = &RouteMatch_Regex{x}
		return true, err
	default:
		return false, nil
	}
}

func _RouteMatch_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RouteMatch)
	// path_pattern
	switch x := m.PathPattern.(type) {
	case *RouteMatch_Prefix:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Prefix)))
		n += len(x.Prefix)
	case *RouteMatch_ExactPath:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ExactPath)))
		n += len(x.ExactPath)
	case *RouteMatch_Regex:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Regex)))
		n += len(x.Regex)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type HeaderMatch struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Types that are valid to be assigned to HeaderMatchPattern:
	//	*HeaderMatch_ExactMatch
	//	*HeaderMatch_RegexMatch
	HeaderMatchPattern   isHeaderMatch_HeaderMatchPattern `protobuf_oneof:"header_match_pattern"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *HeaderMatch) Reset()         { *m = HeaderMatch{} }
func (m *HeaderMatch) String() string { return proto.CompactTextString(m) }
func (*HeaderMatch) ProtoMessage()    {}
func (*HeaderMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{4}
}
func (m *HeaderMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderMatch.Unmarshal(m, b)
}
func (m *HeaderMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderMatch.Marshal(b, m, deterministic)
}
func (dst *HeaderMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderMatch.Merge(dst, src)
}
func (m *HeaderMatch) XXX_Size() int {
	return xxx_messageInfo_HeaderMatch.Size(m)
}
func (m *HeaderMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderMatch.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderMatch proto.InternalMessageInfo

type isHeaderMatch_HeaderMatchPattern interface {
	isHeaderMatch_HeaderMatchPattern()
}

type HeaderMatch_ExactMatch struct {
	ExactMatch string `protobuf:"bytes,2,opt,name=exact_match,json=exactMatch,oneof"`
}
type HeaderMatch_RegexMatch struct {
	RegexMatch string `protobuf:"bytes,3,opt,name=regex_match,json=regexMatch,oneof"`
}

func (*HeaderMatch_ExactMatch) isHeaderMatch_HeaderMatchPattern() {}
func (*HeaderMatch_RegexMatch) isHeaderMatch_HeaderMatchPattern() {}

func (m *HeaderMatch) GetHeaderMatchPattern() isHeaderMatch_HeaderMatchPattern {
	if m != nil {
		return m.HeaderMatchPattern
	}
	return nil
}

func (m *HeaderMatch) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeaderMatch) GetExactMatch() string {
	if x, ok := m.GetHeaderMatchPattern().(*HeaderMatch_ExactMatch); ok {
		return x.ExactMatch
	}
	return ""
}

func (m *HeaderMatch) GetRegexMatch() string {
	if x, ok := m.GetHeaderMatchPattern().(*HeaderMatch_RegexMatch); ok {
		return x.RegexMatch
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*HeaderMatch) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _HeaderMatch_OneofMarshaler, _HeaderMatch_OneofUnmarshaler, _HeaderMatch_OneofSizer, []interface{}{
		(*HeaderMatch_ExactMatch)(nil),
		(*HeaderMatch_RegexMatch)(nil),
	}
}

func _HeaderMatch_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*HeaderMatch)
	// header_match_pattern
	switch x := m.HeaderMatchPattern.(type) {
	case *HeaderMatch_ExactMatch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.ExactMatch)
	case *HeaderMatch_RegexMatch:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.RegexMatch)
	case nil:
	default:
		return fmt.Errorf("HeaderMatch.HeaderMatchPattern has unexpected type %T", x)
	}
	return nil
}

func _HeaderMatch_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*HeaderMatch)
	switch tag {
	case 2: // header_match_pattern.exact_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.HeaderMatchPattern = &HeaderMatch_ExactMatch{x}
		return true, err
	case 3: // header_match_pattern.regex_match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.HeaderMatchPattern = &HeaderMatch_RegexMatch{x}
		return true, err
	default:
		return false, nil
	}
}

func _HeaderMatch_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*HeaderMatch)
	// header_match_pattern
	switch x := m.HeaderMatchPattern.(type) {
	case *HeaderMatch_ExactMatch:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.ExactMatch)))
		n += len(x.ExactMatch)
	case *HeaderMatch_RegexMatch:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.RegexMatch)))
		n += len(x.RegexMatch)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RouteAction struct {
	// Types that are valid to be assigned to ClusterPattern:
	//	*RouteAction_Cluster
	//	*RouteAction_WeightedCluster
	ClusterPattern       isRouteAction_ClusterPattern `protobuf_oneof:"cluster_pattern"`
	TimeoutMs            int32                        `protobuf:"varint,3,opt,name=timeout_ms,json=timeoutMs" json:"timeout_ms,omitempty"`
	RetryStrategy        *RouteAction_RetryStrategy   `protobuf:"bytes,4,opt,name=retry_strategy,json=retryStrategy" json:"retry_strategy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RouteAction) Reset()         { *m = RouteAction{} }
func (m *RouteAction) String() string { return proto.CompactTextString(m) }
func (*RouteAction) ProtoMessage()    {}
func (*RouteAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{5}
}
func (m *RouteAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction.Unmarshal(m, b)
}
func (m *RouteAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction.Marshal(b, m, deterministic)
}
func (dst *RouteAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction.Merge(dst, src)
}
func (m *RouteAction) XXX_Size() int {
	return xxx_messageInfo_RouteAction.Size(m)
}
func (m *RouteAction) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction proto.InternalMessageInfo

type isRouteAction_ClusterPattern interface {
	isRouteAction_ClusterPattern()
}

type RouteAction_Cluster struct {
	Cluster string `protobuf:"bytes,1,opt,name=cluster,oneof"`
}
type RouteAction_WeightedCluster struct {
	WeightedCluster *WeightedCluster `protobuf:"bytes,2,opt,name=weighted_cluster,json=weightedCluster,oneof"`
}

func (*RouteAction_Cluster) isRouteAction_ClusterPattern()         {}
func (*RouteAction_WeightedCluster) isRouteAction_ClusterPattern() {}

func (m *RouteAction) GetClusterPattern() isRouteAction_ClusterPattern {
	if m != nil {
		return m.ClusterPattern
	}
	return nil
}

func (m *RouteAction) GetCluster() string {
	if x, ok := m.GetClusterPattern().(*RouteAction_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *RouteAction) GetWeightedCluster() *WeightedCluster {
	if x, ok := m.GetClusterPattern().(*RouteAction_WeightedCluster); ok {
		return x.WeightedCluster
	}
	return nil
}

func (m *RouteAction) GetTimeoutMs() int32 {
	if m != nil {
		return m.TimeoutMs
	}
	return 0
}

func (m *RouteAction) GetRetryStrategy() *RouteAction_RetryStrategy {
	if m != nil {
		return m.RetryStrategy
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*RouteAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _RouteAction_OneofMarshaler, _RouteAction_OneofUnmarshaler, _RouteAction_OneofSizer, []interface{}{
		(*RouteAction_Cluster)(nil),
		(*RouteAction_WeightedCluster)(nil),
	}
}

func _RouteAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*RouteAction)
	// cluster_pattern
	switch x := m.ClusterPattern.(type) {
	case *RouteAction_Cluster:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Cluster)
	case *RouteAction_WeightedCluster:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WeightedCluster); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("RouteAction.ClusterPattern has unexpected type %T", x)
	}
	return nil
}

func _RouteAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*RouteAction)
	switch tag {
	case 1: // cluster_pattern.cluster
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.ClusterPattern = &RouteAction_Cluster{x}
		return true, err
	case 2: // cluster_pattern.weighted_cluster
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(WeightedCluster)
		err := b.DecodeMessage(msg)
		m.ClusterPattern = &RouteAction_WeightedCluster{msg}
		return true, err
	default:
		return false, nil
	}
}

func _RouteAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*RouteAction)
	// cluster_pattern
	switch x := m.ClusterPattern.(type) {
	case *RouteAction_Cluster:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Cluster)))
		n += len(x.Cluster)
	case *RouteAction_WeightedCluster:
		s := proto.Size(x.WeightedCluster)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type RouteAction_RetryStrategy struct {
	RetryType            string             `protobuf:"bytes,1,opt,name=retry_type,json=retryType" json:"retry_type,omitempty"`
	RetryTimes           uint32             `protobuf:"varint,2,opt,name=retry_times,json=retryTimes" json:"retry_times,omitempty"`
	RetryTimeout         *duration.Duration `protobuf:"bytes,3,opt,name=retry_timeout,json=retryTimeout" json:"retry_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *RouteAction_RetryStrategy) Reset()         { *m = RouteAction_RetryStrategy{} }
func (m *RouteAction_RetryStrategy) String() string { return proto.CompactTextString(m) }
func (*RouteAction_RetryStrategy) ProtoMessage()    {}
func (*RouteAction_RetryStrategy) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{5, 0}
}
func (m *RouteAction_RetryStrategy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteAction_RetryStrategy.Unmarshal(m, b)
}
func (m *RouteAction_RetryStrategy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteAction_RetryStrategy.Marshal(b, m, deterministic)
}
func (dst *RouteAction_RetryStrategy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteAction_RetryStrategy.Merge(dst, src)
}
func (m *RouteAction_RetryStrategy) XXX_Size() int {
	return xxx_messageInfo_RouteAction_RetryStrategy.Size(m)
}
func (m *RouteAction_RetryStrategy) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteAction_RetryStrategy.DiscardUnknown(m)
}

var xxx_messageInfo_RouteAction_RetryStrategy proto.InternalMessageInfo

func (m *RouteAction_RetryStrategy) GetRetryType() string {
	if m != nil {
		return m.RetryType
	}
	return ""
}

func (m *RouteAction_RetryStrategy) GetRetryTimes() uint32 {
	if m != nil {
		return m.RetryTimes
	}
	return 0
}

func (m *RouteAction_RetryStrategy) GetRetryTimeout() *duration.Duration {
	if m != nil {
		return m.RetryTimeout
	}
	return nil
}

type WeightedCluster struct {
	Clusters             []*WeightedCluster_ClusterWeightPair `protobuf:"bytes,1,rep,name=clusters" json:"clusters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                             `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
	XXX_sizecache        int32                                `json:"-"`
}

func (m *WeightedCluster) Reset()         { *m = WeightedCluster{} }
func (m *WeightedCluster) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster) ProtoMessage()    {}
func (*WeightedCluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{6}
}
func (m *WeightedCluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster.Unmarshal(m, b)
}
func (m *WeightedCluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster.Marshal(b, m, deterministic)
}
func (dst *WeightedCluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster.Merge(dst, src)
}
func (m *WeightedCluster) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster.Size(m)
}
func (m *WeightedCluster) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster proto.InternalMessageInfo

func (m *WeightedCluster) GetClusters() []*WeightedCluster_ClusterWeightPair {
	if m != nil {
		return m.Clusters
	}
	return nil
}

type WeightedCluster_ClusterWeightPair struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Weight               uint32   `protobuf:"varint,2,opt,name=weight" json:"weight,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WeightedCluster_ClusterWeightPair) Reset()         { *m = WeightedCluster_ClusterWeightPair{} }
func (m *WeightedCluster_ClusterWeightPair) String() string { return proto.CompactTextString(m) }
func (*WeightedCluster_ClusterWeightPair) ProtoMessage()    {}
func (*WeightedCluster_ClusterWeightPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{6, 0}
}
func (m *WeightedCluster_ClusterWeightPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WeightedCluster_ClusterWeightPair.Unmarshal(m, b)
}
func (m *WeightedCluster_ClusterWeightPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WeightedCluster_ClusterWeightPair.Marshal(b, m, deterministic)
}
func (dst *WeightedCluster_ClusterWeightPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WeightedCluster_ClusterWeightPair.Merge(dst, src)
}
func (m *WeightedCluster_ClusterWeightPair) XXX_Size() int {
	return xxx_messageInfo_WeightedCluster_ClusterWeightPair.Size(m)
}
func (m *WeightedCluster_ClusterWeightPair) XXX_DiscardUnknown() {
	xxx_messageInfo_WeightedCluster_ClusterWeightPair.DiscardUnknown(m)
}

var xxx_messageInfo_WeightedCluster_ClusterWeightPair proto.InternalMessageInfo

func (m *WeightedCluster_ClusterWeightPair) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *WeightedCluster_ClusterWeightPair) GetWeight() uint32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

type RouteDiscoveryResponse struct {
	Result               []*RouteConfig `protobuf:"bytes,1,rep,name=result" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RouteDiscoveryResponse) Reset()         { *m = RouteDiscoveryResponse{} }
func (m *RouteDiscoveryResponse) String() string { return proto.CompactTextString(m) }
func (*RouteDiscoveryResponse) ProtoMessage()    {}
func (*RouteDiscoveryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_route_7902da5439c7a5c4, []int{7}
}
func (m *RouteDiscoveryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteDiscoveryResponse.Unmarshal(m, b)
}
func (m *RouteDiscoveryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteDiscoveryResponse.Marshal(b, m, deterministic)
}
func (dst *RouteDiscoveryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteDiscoveryResponse.Merge(dst, src)
}
func (m *RouteDiscoveryResponse) XXX_Size() int {
	return xxx_messageInfo_RouteDiscoveryResponse.Size(m)
}
func (m *RouteDiscoveryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteDiscoveryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RouteDiscoveryResponse proto.InternalMessageInfo

func (m *RouteDiscoveryResponse) GetResult() []*RouteConfig {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterType((*RouteConfig)(nil), "api.RouteConfig")
	proto.RegisterType((*RouteDetails)(nil), "api.RouteDetails")
	proto.RegisterType((*RouteEntry)(nil), "api.RouteEntry")
	proto.RegisterType((*RouteMatch)(nil), "api.RouteMatch")
	proto.RegisterType((*HeaderMatch)(nil), "api.HeaderMatch")
	proto.RegisterType((*RouteAction)(nil), "api.RouteAction")
	proto.RegisterType((*RouteAction_RetryStrategy)(nil), "api.RouteAction.RetryStrategy")
	proto.RegisterType((*WeightedCluster)(nil), "api.WeightedCluster")
	proto.RegisterType((*WeightedCluster_ClusterWeightPair)(nil), "api.WeightedCluster.ClusterWeightPair")
	proto.RegisterType((*RouteDiscoveryResponse)(nil), "api.RouteDiscoveryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RouteDiscoveryServiceClient is the client API for RouteDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RouteDiscoveryServiceClient interface {
	RetrieveRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*RouteDiscoveryResponse, error)
}

type routeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRouteDiscoveryServiceClient(cc *grpc.ClientConn) RouteDiscoveryServiceClient {
	return &routeDiscoveryServiceClient{cc}
}

func (c *routeDiscoveryServiceClient) RetrieveRoutes(ctx context.Context, in *DiscoveryRequest, opts ...grpc.CallOption) (*RouteDiscoveryResponse, error) {
	out := new(RouteDiscoveryResponse)
	err := c.cc.Invoke(ctx, "/api.RouteDiscoveryService/RetrieveRoutes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RouteDiscoveryService service

type RouteDiscoveryServiceServer interface {
	RetrieveRoutes(context.Context, *DiscoveryRequest) (*RouteDiscoveryResponse, error)
}

func RegisterRouteDiscoveryServiceServer(s *grpc.Server, srv RouteDiscoveryServiceServer) {
	s.RegisterService(&_RouteDiscoveryService_serviceDesc, srv)
}

func _RouteDiscoveryService_RetrieveRoutes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RouteDiscoveryServiceServer).RetrieveRoutes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.RouteDiscoveryService/RetrieveRoutes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RouteDiscoveryServiceServer).RetrieveRoutes(ctx, req.(*DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RouteDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.RouteDiscoveryService",
	HandlerType: (*RouteDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveRoutes",
			Handler:    _RouteDiscoveryService_RetrieveRoutes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "route.proto",
}

func init() { proto.RegisterFile("route.proto", fileDescriptor_route_7902da5439c7a5c4) }

var fileDescriptor_route_7902da5439c7a5c4 = []byte{
	// 655 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x6f, 0x1a, 0x31,
	0x10, 0x85, 0x10, 0x48, 0x18, 0x20, 0x24, 0x56, 0x82, 0xb6, 0x54, 0xcd, 0xc7, 0x4a, 0x6d, 0x51,
	0x0f, 0x1b, 0x29, 0x95, 0x7a, 0x6c, 0x15, 0x92, 0x54, 0x5c, 0x22, 0x45, 0x4e, 0xa4, 0xaa, 0x27,
	0xe4, 0xc0, 0x64, 0xb1, 0x04, 0xbb, 0x5b, 0xdb, 0x9b, 0x84, 0x53, 0x7f, 0x42, 0x6f, 0xbd, 0xf4,
	0x87, 0xf6, 0x5a, 0xed, 0xd8, 0x0b, 0x4b, 0xca, 0x09, 0xfc, 0xde, 0x9b, 0x99, 0xe7, 0xe7, 0xb5,
	0xa1, 0xa1, 0xe2, 0xd4, 0x60, 0x90, 0xa8, 0xd8, 0xc4, 0xac, 0x22, 0x12, 0xd9, 0x3d, 0x0c, 0xe3,
	0x38, 0x9c, 0xe2, 0x29, 0x41, 0xf7, 0xe9, 0xc3, 0xe9, 0x38, 0x55, 0xc2, 0xc8, 0x38, 0xb2, 0xa2,
	0x6e, 0x7b, 0x2c, 0xf5, 0x28, 0x7e, 0x44, 0x35, 0xb7, 0x80, 0xff, 0x1d, 0x1a, 0x3c, 0x6b, 0x72,
	0x11, 0x47, 0x0f, 0x32, 0x64, 0x0c, 0x36, 0x23, 0x31, 0x43, 0xaf, 0x7c, 0x5c, 0xee, 0xd5, 0x39,
	0xfd, 0x67, 0x9f, 0xa0, 0x45, 0x73, 0x86, 0x63, 0x34, 0x42, 0x4e, 0xb5, 0xb7, 0x71, 0x5c, 0xe9,
	0x35, 0xce, 0xf6, 0x02, 0x91, 0xc8, 0x80, 0x8a, 0x2f, 0x2d, 0xc1, 0x9b, 0xaa, 0xb0, 0xf2, 0x11,
	0x9a, 0x45, 0x76, 0x6d, 0x6f, 0x0f, 0xb6, 0xc6, 0xf1, 0x4c, 0xc8, 0xc8, 0x76, 0xad, 0xf3, 0x7c,
	0xc9, 0xde, 0x43, 0x8d, 0xba, 0x69, 0xaf, 0x42, 0xe3, 0xda, 0xcb, 0x71, 0x57, 0x91, 0x51, 0x73,
	0xee, 0x68, 0x3f, 0x04, 0x58, 0xa2, 0xec, 0x2d, 0x54, 0x67, 0xc2, 0x8c, 0x26, 0x34, 0x65, 0xa5,
	0xea, 0x3a, 0x83, 0xb9, 0x65, 0x59, 0x0f, 0xaa, 0x54, 0xee, 0x6d, 0x90, 0x6c, 0x77, 0x29, 0x3b,
	0x1f, 0x65, 0x71, 0x0d, 0x4a, 0xdc, 0x0a, 0xfa, 0xdb, 0x50, 0x13, 0x04, 0xf9, 0x7f, 0xca, 0x6e,
	0x12, 0x75, 0x62, 0x1e, 0xd4, 0x12, 0x85, 0x0f, 0xf2, 0xd9, 0x6e, 0x68, 0x50, 0xe2, 0x6e, 0xcd,
	0x8e, 0x00, 0xf0, 0x59, 0x8c, 0xcc, 0x30, 0x11, 0x66, 0x42, 0x13, 0x32, 0xb6, 0x4e, 0xd8, 0x8d,
	0x30, 0x13, 0xd6, 0x81, 0xaa, 0xc2, 0x10, 0x9f, 0xbd, 0x8a, 0xe3, 0xec, 0x92, 0x7d, 0x80, 0xad,
	0x09, 0x8a, 0x31, 0x2a, 0xed, 0x6d, 0xd2, 0xa6, 0xad, 0xaf, 0x01, 0x61, 0xd6, 0x7f, 0x2e, 0xe8,
	0xef, 0x40, 0x33, 0x6b, 0x9f, 0xcd, 0x30, 0xa8, 0x22, 0xff, 0x27, 0x34, 0x0a, 0xba, 0xb5, 0x61,
	0x9f, 0x40, 0xc3, 0xfa, 0xb2, 0x09, 0xe5, 0xc6, 0xac, 0x59, 0x5b, 0x76, 0x02, 0x0d, 0xb2, 0xe2,
	0x24, 0xb9, 0x3f, 0x20, 0x90, 0x24, 0xfd, 0x0e, 0xec, 0x5b, 0x0f, 0x56, 0xb3, 0x30, 0xf0, 0x77,
	0xc3, 0x7d, 0x4a, 0x36, 0x41, 0xd6, 0x85, 0xad, 0xd1, 0x34, 0xd5, 0x06, 0xd5, 0x22, 0xa0, 0x1c,
	0x60, 0xe7, 0xb0, 0xfb, 0x84, 0x32, 0x9c, 0x18, 0x1c, 0x0f, 0x73, 0x91, 0x3d, 0x89, 0x7d, 0xda,
	0xf1, 0x37, 0x47, 0x5e, 0x58, 0x6e, 0x50, 0xe2, 0xed, 0xa7, 0x55, 0x88, 0xbd, 0x01, 0x30, 0x72,
	0x86, 0x71, 0x6a, 0x86, 0x33, 0x4d, 0x46, 0xab, 0xbc, 0xee, 0x90, 0x6b, 0xcd, 0xae, 0x60, 0x47,
	0xa1, 0x51, 0xf3, 0xa1, 0x36, 0x4a, 0x18, 0x0c, 0xe7, 0xde, 0x26, 0xf5, 0x3f, 0x7c, 0x79, 0xd2,
	0x01, 0xcf, 0x64, 0xb7, 0x4e, 0xc5, 0x5b, 0xaa, 0xb8, 0xec, 0xfe, 0x2a, 0x43, 0x6b, 0x45, 0x90,
	0xcd, 0xb5, 0x8d, 0xcd, 0x3c, 0xc9, 0xe3, 0xad, 0x13, 0x72, 0x37, 0x4f, 0x90, 0x1d, 0x65, 0x01,
	0x12, 0x2d, 0x67, 0xa8, 0x69, 0x53, 0x2d, 0x6e, 0x2b, 0xee, 0x32, 0x84, 0x7d, 0x86, 0xd6, 0x52,
	0x10, 0xa7, 0x86, 0xac, 0x37, 0xce, 0x5e, 0x05, 0xf6, 0xe6, 0x06, 0xf9, 0xcd, 0x0d, 0x2e, 0xdd,
	0xcd, 0xe5, 0xcd, 0x45, 0x75, 0x9c, 0x9a, 0xfe, 0x1e, 0xb4, 0x5d, 0x62, 0x8b, 0xe4, 0x7f, 0x97,
	0xa1, 0xfd, 0x22, 0x31, 0xd6, 0x87, 0x6d, 0x27, 0xd3, 0x5e, 0x99, 0xbe, 0xa5, 0x77, 0xeb, 0x92,
	0x0d, 0xdc, 0xaf, 0x85, 0x6f, 0x84, 0x54, 0x7c, 0x51, 0xd7, 0xfd, 0x02, 0x7b, 0xff, 0xd1, 0x6b,
	0x3f, 0xac, 0x0e, 0xd4, 0xec, 0xf1, 0xb8, 0xfd, 0xba, 0x95, 0xdf, 0x87, 0x8e, 0x7d, 0x01, 0xf2,
	0x47, 0x87, 0xa3, 0x4e, 0xe2, 0x48, 0x23, 0xeb, 0x41, 0x4d, 0xa1, 0x4e, 0xa7, 0xc6, 0x99, 0x2b,
	0x5c, 0x40, 0xfb, 0x12, 0x71, 0xc7, 0x9f, 0x0d, 0xe1, 0x60, 0xb5, 0xc7, 0x2d, 0xaa, 0x47, 0x39,
	0x42, 0xf6, 0x15, 0x76, 0xb2, 0x93, 0x91, 0xf8, 0x88, 0x24, 0xd0, 0xec, 0x80, 0x9a, 0x14, 0x86,
	0xfd, 0x48, 0x51, 0x9b, 0xee, 0xeb, 0xc2, 0x43, 0xf5, 0xd2, 0x88, 0x5f, 0xba, 0xaf, 0x51, 0xe2,
	0x1f, 0xff, 0x05, 0x00, 0x00, 0xff, 0xff, 0xfa, 0x87, 0x6b, 0xe4, 0x4d, 0x05, 0x00, 0x00,
}
