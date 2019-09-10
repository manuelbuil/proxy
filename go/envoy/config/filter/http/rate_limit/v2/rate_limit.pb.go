// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/rate_limit/v2/rate_limit.proto

package envoy_config_filter_http_rate_limit_v2

import (
	fmt "fmt"
	v2 "github.com/cilium/proxy/go/envoy/config/ratelimit/v2"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type RateLimit struct {
	// The rate limit domain to use when calling the rate limit service.
	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	// Specifies the rate limit configurations to be applied with the same
	// stage number. If not set, the default stage number is 0.
	//
	// .. note::
	//
	//  The filter supports a range of 0 - 10 inclusively for stage numbers.
	Stage uint32 `protobuf:"varint,2,opt,name=stage,proto3" json:"stage,omitempty"`
	// The type of requests the filter should apply to. The supported
	// types are *internal*, *external* or *both*. A request is considered internal if
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is set to true. If
	// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>` is not set or false, a
	// request is considered external. The filter defaults to *both*, and it will apply to all request
	// types.
	RequestType string `protobuf:"bytes,3,opt,name=request_type,json=requestType,proto3" json:"request_type,omitempty"`
	// The timeout in milliseconds for the rate limit service RPC. If not
	// set, this defaults to 20ms.
	Timeout *duration.Duration `protobuf:"bytes,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// The filter's behaviour in case the rate limiting service does
	// not respond back. When it is set to true, Envoy will not allow traffic in case of
	// communication failure between rate limiting service and the proxy.
	// Defaults to false.
	FailureModeDeny bool `protobuf:"varint,5,opt,name=failure_mode_deny,json=failureModeDeny,proto3" json:"failure_mode_deny,omitempty"`
	// Specifies whether a `RESOURCE_EXHAUSTED` gRPC code must be returned instead
	// of the default `UNAVAILABLE` gRPC code for a rate limited gRPC call. The
	// HTTP code will be 200 for a gRPC response.
	RateLimitedAsResourceExhausted bool `protobuf:"varint,6,opt,name=rate_limited_as_resource_exhausted,json=rateLimitedAsResourceExhausted,proto3" json:"rate_limited_as_resource_exhausted,omitempty"`
	// Configuration for an external rate limit service provider. If not
	// specified, any calls to the rate limit service will immediately return
	// success.
	RateLimitService     *v2.RateLimitServiceConfig `protobuf:"bytes,7,opt,name=rate_limit_service,json=rateLimitService,proto3" json:"rate_limit_service,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *RateLimit) Reset()         { *m = RateLimit{} }
func (m *RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimit) ProtoMessage()    {}
func (*RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_af348a51c982d3a6, []int{0}
}

func (m *RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimit.Unmarshal(m, b)
}
func (m *RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimit.Merge(m, src)
}
func (m *RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimit.Size(m)
}
func (m *RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimit proto.InternalMessageInfo

func (m *RateLimit) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimit) GetStage() uint32 {
	if m != nil {
		return m.Stage
	}
	return 0
}

func (m *RateLimit) GetRequestType() string {
	if m != nil {
		return m.RequestType
	}
	return ""
}

func (m *RateLimit) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

func (m *RateLimit) GetFailureModeDeny() bool {
	if m != nil {
		return m.FailureModeDeny
	}
	return false
}

func (m *RateLimit) GetRateLimitedAsResourceExhausted() bool {
	if m != nil {
		return m.RateLimitedAsResourceExhausted
	}
	return false
}

func (m *RateLimit) GetRateLimitService() *v2.RateLimitServiceConfig {
	if m != nil {
		return m.RateLimitService
	}
	return nil
}

func init() {
	proto.RegisterType((*RateLimit)(nil), "envoy.config.filter.http.rate_limit.v2.RateLimit")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/rate_limit/v2/rate_limit.proto", fileDescriptor_af348a51c982d3a6)
}

var fileDescriptor_af348a51c982d3a6 = []byte{
	// 404 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xcb, 0x6e, 0xd4, 0x30,
	0x14, 0x95, 0xa7, 0x9d, 0x99, 0xd6, 0xe5, 0x51, 0xbc, 0xc1, 0x54, 0xa2, 0x84, 0x22, 0xa1, 0x51,
	0x17, 0xb6, 0x98, 0x22, 0xb1, 0x26, 0x14, 0x09, 0x21, 0x90, 0x2a, 0xc3, 0xde, 0x72, 0xc7, 0x37,
	0xa9, 0xa5, 0x24, 0x0e, 0xb6, 0x13, 0x35, 0x3f, 0xc0, 0x82, 0x4f, 0xee, 0x0a, 0x25, 0x4e, 0x66,
	0x5a, 0x56, 0xec, 0xe2, 0x7b, 0x1e, 0x3e, 0x39, 0xd7, 0xf8, 0x03, 0x54, 0xad, 0xed, 0xf8, 0xc6,
	0x56, 0x99, 0xc9, 0x79, 0x66, 0x8a, 0x00, 0x8e, 0xdf, 0x84, 0x50, 0x73, 0xa7, 0x02, 0xc8, 0xc2,
	0x94, 0x26, 0xf0, 0x76, 0x7d, 0xef, 0xc4, 0x6a, 0x67, 0x83, 0x25, 0x6f, 0x07, 0x21, 0x8b, 0x42,
	0x16, 0x85, 0xac, 0x17, 0xb2, 0x7b, 0xd4, 0x76, 0x7d, 0xf2, 0xe6, 0xc1, 0x05, 0x3d, 0xb6, 0xf3,
	0x2c, 0x7c, 0x34, 0x3b, 0x39, 0xcd, 0xad, 0xcd, 0x0b, 0xe0, 0xc3, 0xe9, 0xba, 0xc9, 0xb8, 0x6e,
	0x9c, 0x0a, 0xc6, 0x56, 0x23, 0xfe, 0xbc, 0x55, 0x85, 0xd1, 0x2a, 0x00, 0x9f, 0x3e, 0x22, 0x70,
	0xf6, 0x7b, 0x0f, 0x1f, 0x0a, 0x15, 0xe0, 0x5b, 0xef, 0x49, 0x5e, 0xe1, 0x85, 0xb6, 0xa5, 0x32,
	0x15, 0x45, 0x09, 0x5a, 0x1d, 0xa6, 0xcb, 0xbb, 0x74, 0xdf, 0xcd, 0x12, 0x24, 0xc6, 0x31, 0x79,
	0x89, 0xe7, 0x3e, 0xa8, 0x1c, 0xe8, 0x2c, 0x41, 0xab, 0xc7, 0x03, 0x7e, 0x3e, 0xa3, 0x58, 0xc4,
	0x29, 0x79, 0x8d, 0x1f, 0x39, 0xf8, 0xd5, 0x80, 0x0f, 0x32, 0x74, 0x35, 0xd0, 0xbd, 0xde, 0x45,
	0x1c, 0x8d, 0xb3, 0x9f, 0x5d, 0x0d, 0xe4, 0x02, 0x2f, 0x83, 0x29, 0xc1, 0x36, 0x81, 0xee, 0x27,
	0x68, 0x75, 0xb4, 0x7e, 0xc1, 0x62, 0x76, 0x36, 0x65, 0x67, 0x97, 0x63, 0x76, 0x31, 0x31, 0xc9,
	0x39, 0x7e, 0x96, 0x29, 0x53, 0x34, 0x0e, 0x64, 0x69, 0x35, 0x48, 0x0d, 0x55, 0x47, 0xe7, 0x09,
	0x5a, 0x1d, 0x88, 0xa7, 0x23, 0xf0, 0xdd, 0x6a, 0xb8, 0x84, 0xaa, 0x23, 0x5f, 0xf1, 0xd9, 0xae,
	0x40, 0xd0, 0x52, 0x79, 0xe9, 0xc0, 0xdb, 0xc6, 0x6d, 0x40, 0xc2, 0xed, 0x8d, 0x6a, 0x7c, 0x00,
	0x4d, 0x17, 0x83, 0xf8, 0xd4, 0x4d, 0xbf, 0x0e, 0xfa, 0xa3, 0x17, 0x23, 0xed, 0xf3, 0xc4, 0x22,
	0x06, 0x93, 0x9d, 0x97, 0xf4, 0xe0, 0x5a, 0xb3, 0x01, 0xba, 0x1c, 0x72, 0xbf, 0x63, 0x0f, 0x16,
	0xb8, 0x5d, 0x0c, 0x6b, 0xd7, 0x6c, 0xdb, 0xe8, 0x8f, 0x28, 0xf9, 0x34, 0x70, 0xd2, 0x83, 0xbb,
	0x74, 0xfe, 0x07, 0xcd, 0x8e, 0x91, 0x38, 0x76, 0xff, 0x30, 0xd2, 0x2f, 0xf8, 0xbd, 0xb1, 0xd1,
	0xb2, 0x76, 0xf6, 0xb6, 0x63, 0xff, 0xf7, 0x3c, 0xd2, 0x27, 0xdb, 0xbb, 0xae, 0xfa, 0xfe, 0xae,
	0xd0, 0xf5, 0x62, 0x28, 0xf2, 0xe2, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x79, 0x1f, 0x21, 0x57,
	0x9a, 0x02, 0x00, 0x00,
}
