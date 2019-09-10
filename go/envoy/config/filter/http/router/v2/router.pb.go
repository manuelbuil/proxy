// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/router/v2/router.proto

package envoy_config_filter_http_router_v2

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/config/filter/accesslog/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Router struct {
	// Whether the router generates dynamic cluster statistics. Defaults to
	// true. Can be disabled in high performance scenarios.
	DynamicStats *wrappers.BoolValue `protobuf:"bytes,1,opt,name=dynamic_stats,json=dynamicStats,proto3" json:"dynamic_stats,omitempty"`
	// Whether to start a child span for egress routed calls. This can be
	// useful in scenarios where other filters (auth, ratelimit, etc.) make
	// outbound calls and have child spans rooted at the same ingress
	// parent. Defaults to false.
	StartChildSpan bool `protobuf:"varint,2,opt,name=start_child_span,json=startChildSpan,proto3" json:"start_child_span,omitempty"`
	// Configuration for HTTP upstream logs emitted by the router. Upstream logs
	// are configured in the same way as access logs, but each log entry represents
	// an upstream request. Presuming retries are configured, multiple upstream
	// requests may be made for each downstream (inbound) request.
	UpstreamLog []*v2.AccessLog `protobuf:"bytes,3,rep,name=upstream_log,json=upstreamLog,proto3" json:"upstream_log,omitempty"`
	// Do not add any additional *x-envoy-* headers to requests or responses. This
	// only affects the :ref:`router filter generated *x-envoy-* headers
	// <config_http_filters_router_headers_set>`, other Envoy filters and the HTTP
	// connection manager may continue to set *x-envoy-* headers.
	SuppressEnvoyHeaders bool `protobuf:"varint,4,opt,name=suppress_envoy_headers,json=suppressEnvoyHeaders,proto3" json:"suppress_envoy_headers,omitempty"`
	// Specifies a list of HTTP headers to strictly validate. Envoy will reject a
	// request and respond with HTTP status 400 if the request contains an invalid
	// value for any of the headers listed in this field. Strict header checking
	// is only supported for the following headers:
	//
	// Value must be a ','-delimited list (i.e. no spaces) of supported retry
	// policy values:
	//
	// * :ref:`config_http_filters_router_x-envoy-retry-grpc-on`
	// * :ref:`config_http_filters_router_x-envoy-retry-on`
	//
	// Value must be an integer:
	//
	// * :ref:`config_http_filters_router_x-envoy-max-retries`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-timeout-ms`
	// * :ref:`config_http_filters_router_x-envoy-upstream-rq-per-try-timeout-ms`
	StrictCheckHeaders   []string `protobuf:"bytes,5,rep,name=strict_check_headers,json=strictCheckHeaders,proto3" json:"strict_check_headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Router) Reset()         { *m = Router{} }
func (m *Router) String() string { return proto.CompactTextString(m) }
func (*Router) ProtoMessage()    {}
func (*Router) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc1f525510d06eb8, []int{0}
}

func (m *Router) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Router.Unmarshal(m, b)
}
func (m *Router) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Router.Marshal(b, m, deterministic)
}
func (m *Router) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Router.Merge(m, src)
}
func (m *Router) XXX_Size() int {
	return xxx_messageInfo_Router.Size(m)
}
func (m *Router) XXX_DiscardUnknown() {
	xxx_messageInfo_Router.DiscardUnknown(m)
}

var xxx_messageInfo_Router proto.InternalMessageInfo

func (m *Router) GetDynamicStats() *wrappers.BoolValue {
	if m != nil {
		return m.DynamicStats
	}
	return nil
}

func (m *Router) GetStartChildSpan() bool {
	if m != nil {
		return m.StartChildSpan
	}
	return false
}

func (m *Router) GetUpstreamLog() []*v2.AccessLog {
	if m != nil {
		return m.UpstreamLog
	}
	return nil
}

func (m *Router) GetSuppressEnvoyHeaders() bool {
	if m != nil {
		return m.SuppressEnvoyHeaders
	}
	return false
}

func (m *Router) GetStrictCheckHeaders() []string {
	if m != nil {
		return m.StrictCheckHeaders
	}
	return nil
}

func init() {
	proto.RegisterType((*Router)(nil), "envoy.config.filter.http.router.v2.Router")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/router/v2/router.proto", fileDescriptor_cc1f525510d06eb8)
}

var fileDescriptor_cc1f525510d06eb8 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8a, 0xd4, 0x40,
	0x10, 0xa5, 0x1d, 0x5d, 0x34, 0xb3, 0xca, 0x12, 0x57, 0x1d, 0xe6, 0xb0, 0x0c, 0x73, 0x90, 0x01,
	0xe9, 0xce, 0x12, 0xbd, 0xcb, 0x66, 0x11, 0x3c, 0x2c, 0xb2, 0xf4, 0x82, 0xd7, 0xd0, 0x9b, 0xe9,
	0xc9, 0x34, 0x26, 0xa9, 0xb6, 0xba, 0x13, 0x27, 0x3f, 0x20, 0x08, 0x1e, 0x44, 0xfc, 0x23, 0x0f,
	0xfe, 0x93, 0x27, 0xe9, 0xee, 0x64, 0x55, 0x18, 0xd8, 0x5b, 0xd5, 0x7b, 0xf5, 0x5e, 0x55, 0x57,
	0x75, 0x94, 0xc8, 0xa6, 0x83, 0x3e, 0x29, 0xa0, 0xd9, 0xa8, 0x32, 0xd9, 0xa8, 0xca, 0x4a, 0x4c,
	0xb6, 0xd6, 0xea, 0x04, 0xa1, 0x75, 0x71, 0x97, 0x0e, 0x11, 0xd3, 0x08, 0x16, 0xe2, 0xa5, 0x17,
	0xb0, 0x20, 0x60, 0x41, 0xc0, 0x9c, 0x80, 0x0d, 0x65, 0x5d, 0x3a, 0x3f, 0xdd, 0x67, 0x2a, 0x8a,
	0x42, 0x1a, 0x53, 0x41, 0xe9, 0x2c, 0x6f, 0x92, 0xe0, 0x3a, 0x3f, 0x29, 0x01, 0xca, 0x4a, 0x26,
	0x3e, 0xbb, 0x6e, 0x37, 0xc9, 0x27, 0x14, 0x5a, 0x4b, 0x34, 0x03, 0xff, 0xac, 0x13, 0x95, 0x5a,
	0x0b, 0x2b, 0x93, 0x31, 0x08, 0xc4, 0xf2, 0xe7, 0x24, 0x3a, 0xe0, 0xbe, 0x71, 0xfc, 0x3a, 0x7a,
	0xb8, 0xee, 0x1b, 0x51, 0xab, 0x22, 0x37, 0x56, 0x58, 0x33, 0x23, 0x0b, 0xb2, 0x9a, 0xa6, 0x73,
	0x16, 0xbc, 0xd9, 0xe8, 0xcd, 0x32, 0x80, 0xea, 0xbd, 0xa8, 0x5a, 0xc9, 0x0f, 0x07, 0xc1, 0x95,
	0xab, 0x8f, 0x57, 0xd1, 0x91, 0xb1, 0x02, 0x6d, 0x5e, 0x6c, 0x55, 0xb5, 0xce, 0x8d, 0x16, 0xcd,
	0xec, 0xce, 0x82, 0xac, 0xee, 0xf3, 0x47, 0x1e, 0x3f, 0x77, 0xf0, 0x95, 0x16, 0x4d, 0xfc, 0x2e,
	0x3a, 0x6c, 0xb5, 0xb1, 0x28, 0x45, 0x9d, 0x57, 0x50, 0xce, 0x26, 0x8b, 0xc9, 0x6a, 0x9a, 0xbe,
	0x60, 0xfb, 0x76, 0xf3, 0xf7, 0xa9, 0x5d, 0xca, 0xce, 0x7c, 0x72, 0x01, 0x25, 0x9f, 0x8e, 0x06,
	0x17, 0x50, 0xc6, 0xaf, 0xa2, 0xa7, 0xa6, 0xd5, 0x1a, 0xa5, 0x31, 0xb9, 0xf7, 0xc8, 0xb7, 0x52,
	0xac, 0x25, 0x9a, 0xd9, 0x5d, 0xdf, 0xff, 0x78, 0x64, 0xdf, 0x38, 0xf2, 0x6d, 0xe0, 0xe2, 0x5f,
	0x24, 0x3a, 0x36, 0x16, 0x55, 0xe1, 0x26, 0x96, 0xc5, 0x87, 0x1b, 0xd1, 0xbd, 0xc5, 0x64, 0xf5,
	0x20, 0xfb, 0x41, 0x7e, 0x67, 0xdf, 0xc8, 0x77, 0xf2, 0x95, 0x2c, 0xbf, 0x10, 0xfc, 0x4c, 0xf8,
	0xc9, 0x8e, 0x7a, 0x73, 0x3a, 0xb6, 0xa6, 0xf8, 0x91, 0x5a, 0x55, 0x4b, 0x68, 0x2d, 0xad, 0x0d,
	0x7f, 0xbe, 0x8f, 0xd7, 0x12, 0xa9, 0xc5, 0xfe, 0xdf, 0xba, 0xc7, 0x63, 0x5d, 0x2d, 0x76, 0x14,
	0xa5, 0x45, 0x25, 0x0d, 0x7f, 0x32, 0x82, 0x0e, 0xe8, 0x69, 0x89, 0xba, 0xa0, 0xd0, 0xf0, 0xa3,
	0xff, 0x61, 0x68, 0x78, 0x1c, 0x46, 0x3e, 0x77, 0x13, 0x0f, 0x2f, 0xc9, 0xce, 0xa2, 0x53, 0x05,
	0x61, 0x7b, 0x1a, 0x61, 0xd7, 0xb3, 0xdb, 0x3f, 0x59, 0x36, 0x0d, 0x67, 0xbf, 0x74, 0x57, 0xbd,
	0x24, 0xd7, 0x07, 0xfe, 0xbc, 0x2f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xe7, 0xac, 0xb7,
	0xd1, 0x02, 0x00, 0x00,
}
