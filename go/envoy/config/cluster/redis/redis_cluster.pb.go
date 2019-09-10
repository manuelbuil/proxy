// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/cluster/redis/redis_cluster.proto

package envoy_config_cluster_redis

import (
	fmt "fmt"
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

type RedisClusterConfig struct {
	// Interval between successive topology refresh requests. If not set, this defaults to 5s.
	ClusterRefreshRate *duration.Duration `protobuf:"bytes,1,opt,name=cluster_refresh_rate,json=clusterRefreshRate,proto3" json:"cluster_refresh_rate,omitempty"`
	// Timeout for topology refresh request. If not set, this defaults to 3s.
	ClusterRefreshTimeout *duration.Duration `protobuf:"bytes,2,opt,name=cluster_refresh_timeout,json=clusterRefreshTimeout,proto3" json:"cluster_refresh_timeout,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}           `json:"-"`
	XXX_unrecognized      []byte             `json:"-"`
	XXX_sizecache         int32              `json:"-"`
}

func (m *RedisClusterConfig) Reset()         { *m = RedisClusterConfig{} }
func (m *RedisClusterConfig) String() string { return proto.CompactTextString(m) }
func (*RedisClusterConfig) ProtoMessage()    {}
func (*RedisClusterConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d6593a6ec218c02, []int{0}
}

func (m *RedisClusterConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisClusterConfig.Unmarshal(m, b)
}
func (m *RedisClusterConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisClusterConfig.Marshal(b, m, deterministic)
}
func (m *RedisClusterConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisClusterConfig.Merge(m, src)
}
func (m *RedisClusterConfig) XXX_Size() int {
	return xxx_messageInfo_RedisClusterConfig.Size(m)
}
func (m *RedisClusterConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisClusterConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RedisClusterConfig proto.InternalMessageInfo

func (m *RedisClusterConfig) GetClusterRefreshRate() *duration.Duration {
	if m != nil {
		return m.ClusterRefreshRate
	}
	return nil
}

func (m *RedisClusterConfig) GetClusterRefreshTimeout() *duration.Duration {
	if m != nil {
		return m.ClusterRefreshTimeout
	}
	return nil
}

func init() {
	proto.RegisterType((*RedisClusterConfig)(nil), "envoy.config.cluster.redis.RedisClusterConfig")
}

func init() {
	proto.RegisterFile("envoy/config/cluster/redis/redis_cluster.proto", fileDescriptor_6d6593a6ec218c02)
}

var fileDescriptor_6d6593a6ec218c02 = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4b, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d,
	0xd2, 0x2f, 0x4a, 0x4d, 0xc9, 0x2c, 0x86, 0x90, 0xf1, 0x50, 0x31, 0xbd, 0x82, 0xa2, 0xfc, 0x92,
	0x7c, 0x21, 0x29, 0xb0, 0x7a, 0x3d, 0x88, 0x7a, 0x3d, 0x98, 0x1c, 0x58, 0xa5, 0x94, 0x5c, 0x7a,
	0x7e, 0x7e, 0x7a, 0x4e, 0xaa, 0x3e, 0x58, 0x65, 0x52, 0x69, 0x9a, 0x7e, 0x4a, 0x69, 0x51, 0x62,
	0x49, 0x66, 0x7e, 0x1e, 0x44, 0xaf, 0x94, 0x78, 0x59, 0x62, 0x4e, 0x66, 0x4a, 0x62, 0x49, 0xaa,
	0x3e, 0x8c, 0x01, 0x91, 0x50, 0x3a, 0xc1, 0xc8, 0x25, 0x14, 0x04, 0x32, 0xc2, 0x19, 0x62, 0x9e,
	0x33, 0xd8, 0x74, 0xa1, 0x50, 0x2e, 0x11, 0xa8, 0x05, 0xf1, 0x45, 0xa9, 0x69, 0x45, 0xa9, 0xc5,
	0x19, 0xf1, 0x45, 0x89, 0x25, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x92, 0x7a, 0x10,
	0xeb, 0xf4, 0x60, 0xd6, 0xe9, 0xb9, 0x40, 0xad, 0x73, 0xe2, 0xf8, 0xe5, 0xc4, 0xba, 0x8a, 0x91,
	0x49, 0x8b, 0x21, 0x48, 0x08, 0x6a, 0x40, 0x10, 0x44, 0x7f, 0x50, 0x62, 0x49, 0xaa, 0x50, 0x34,
	0x97, 0x38, 0xba, 0xb1, 0x25, 0x99, 0xb9, 0xa9, 0xf9, 0xa5, 0x25, 0x12, 0x4c, 0xc4, 0x9b, 0x2c,
	0x8a, 0x6a, 0x72, 0x08, 0xc4, 0x04, 0x27, 0x7b, 0x2e, 0x8d, 0xcc, 0x7c, 0x48, 0xa0, 0x16, 0x14,
	0xe5, 0x57, 0x54, 0xea, 0xe1, 0x0e, 0x2f, 0x27, 0x41, 0x64, 0x3f, 0x07, 0x80, 0xec, 0x0a, 0x60,
	0x4c, 0x62, 0x03, 0x5b, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x82, 0x93, 0x5c, 0x04, 0x99,
	0x01, 0x00, 0x00,
}
