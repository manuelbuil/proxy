// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/filters/network/redis_proxy/v3/redis_proxy.proto

package envoy_extensions_filters_network_redis_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cilium/proxy/go/envoy/annotations"
	v3 "github.com/cilium/proxy/go/envoy/config/core/v3"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

// ReadPolicy controls how Envoy routes read commands to Redis nodes. This is currently
// supported for Redis Cluster. All ReadPolicy settings except MASTER may return stale data
// because replication is asynchronous and requires some delay. You need to ensure that your
// application can tolerate stale data.
type RedisProxy_ConnPoolSettings_ReadPolicy int32

const (
	// Default mode. Read from the current master node.
	RedisProxy_ConnPoolSettings_MASTER RedisProxy_ConnPoolSettings_ReadPolicy = 0
	// Read from the master, but if it is unavailable, read from replica nodes.
	RedisProxy_ConnPoolSettings_PREFER_MASTER RedisProxy_ConnPoolSettings_ReadPolicy = 1
	// Read from replica nodes. If multiple replica nodes are present within a shard, a random
	// node is selected. Healthy nodes have precedent over unhealthy nodes.
	RedisProxy_ConnPoolSettings_REPLICA RedisProxy_ConnPoolSettings_ReadPolicy = 2
	// Read from the replica nodes (similar to REPLICA), but if all replicas are unavailable (not
	// present or unhealthy), read from the master.
	RedisProxy_ConnPoolSettings_PREFER_REPLICA RedisProxy_ConnPoolSettings_ReadPolicy = 3
	// Read from any node of the cluster. A random node is selected among the master and replicas,
	// healthy nodes have precedent over unhealthy nodes.
	RedisProxy_ConnPoolSettings_ANY RedisProxy_ConnPoolSettings_ReadPolicy = 4
)

var RedisProxy_ConnPoolSettings_ReadPolicy_name = map[int32]string{
	0: "MASTER",
	1: "PREFER_MASTER",
	2: "REPLICA",
	3: "PREFER_REPLICA",
	4: "ANY",
}

var RedisProxy_ConnPoolSettings_ReadPolicy_value = map[string]int32{
	"MASTER":         0,
	"PREFER_MASTER":  1,
	"REPLICA":        2,
	"PREFER_REPLICA": 3,
	"ANY":            4,
}

func (x RedisProxy_ConnPoolSettings_ReadPolicy) String() string {
	return proto.EnumName(RedisProxy_ConnPoolSettings_ReadPolicy_name, int32(x))
}

func (RedisProxy_ConnPoolSettings_ReadPolicy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0, 0, 0}
}

// [#next-free-field: 7]
type RedisProxy struct {
	// The prefix to use when emitting :ref:`statistics <config_network_filters_redis_proxy_stats>`.
	StatPrefix string `protobuf:"bytes,1,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
	// Network settings for the connection pool to the upstream clusters.
	Settings *RedisProxy_ConnPoolSettings `protobuf:"bytes,3,opt,name=settings,proto3" json:"settings,omitempty"`
	// Indicates that latency stat should be computed in microseconds. By default it is computed in
	// milliseconds.
	LatencyInMicros bool `protobuf:"varint,4,opt,name=latency_in_micros,json=latencyInMicros,proto3" json:"latency_in_micros,omitempty"`
	// List of **unique** prefixes used to separate keys from different workloads to different
	// clusters. Envoy will always favor the longest match first in case of overlap. A catch-all
	// cluster can be used to forward commands when there is no match. Time complexity of the
	// lookups are in O(min(longest key prefix, key length)).
	//
	// Example:
	//
	// .. code-block:: yaml
	//
	//    prefix_routes:
	//      routes:
	//        - prefix: "ab"
	//          cluster: "cluster_a"
	//        - prefix: "abc"
	//          cluster: "cluster_b"
	//
	// When using the above routes, the following prefixes would be sent to:
	//
	// * 'get abc:users' would retrieve the key 'abc:users' from cluster_b.
	// * 'get ab:users' would retrieve the key 'ab:users' from cluster_a.
	// * 'get z:users' would return a NoUpstreamHost error. A :ref:`catch-all
	//   route<envoy_api_field_extensions.filters.network.redis_proxy.v3.RedisProxy.PrefixRoutes.catch_all_route>`
	//   would have retrieved the key from that cluster instead.
	//
	// See the :ref:`configuration section
	// <arch_overview_redis_configuration>` of the architecture overview for recommendations on
	// configuring the backing clusters.
	PrefixRoutes *RedisProxy_PrefixRoutes `protobuf:"bytes,5,opt,name=prefix_routes,json=prefixRoutes,proto3" json:"prefix_routes,omitempty"`
	// Authenticate Redis client connections locally by forcing downstream clients to issue a `Redis
	// AUTH command <https://redis.io/commands/auth>`_ with this password before enabling any other
	// command. If an AUTH command's password matches this password, an "OK" response will be returned
	// to the client. If the AUTH command password does not match this password, then an "ERR invalid
	// password" error will be returned. If any other command is received before AUTH when this
	// password is set, then a "NOAUTH Authentication required." error response will be sent to the
	// client. If an AUTH command is received when the password is not set, then an "ERR Client sent
	// AUTH, but no password is set" error will be returned.
	DownstreamAuthPassword *v3.DataSource `protobuf:"bytes,6,opt,name=downstream_auth_password,json=downstreamAuthPassword,proto3" json:"downstream_auth_password,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}       `json:"-"`
	XXX_unrecognized       []byte         `json:"-"`
	XXX_sizecache          int32          `json:"-"`
}

func (m *RedisProxy) Reset()         { *m = RedisProxy{} }
func (m *RedisProxy) String() string { return proto.CompactTextString(m) }
func (*RedisProxy) ProtoMessage()    {}
func (*RedisProxy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0}
}

func (m *RedisProxy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy.Unmarshal(m, b)
}
func (m *RedisProxy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy.Marshal(b, m, deterministic)
}
func (m *RedisProxy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy.Merge(m, src)
}
func (m *RedisProxy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy.Size(m)
}
func (m *RedisProxy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy proto.InternalMessageInfo

func (m *RedisProxy) GetStatPrefix() string {
	if m != nil {
		return m.StatPrefix
	}
	return ""
}

func (m *RedisProxy) GetSettings() *RedisProxy_ConnPoolSettings {
	if m != nil {
		return m.Settings
	}
	return nil
}

func (m *RedisProxy) GetLatencyInMicros() bool {
	if m != nil {
		return m.LatencyInMicros
	}
	return false
}

func (m *RedisProxy) GetPrefixRoutes() *RedisProxy_PrefixRoutes {
	if m != nil {
		return m.PrefixRoutes
	}
	return nil
}

func (m *RedisProxy) GetDownstreamAuthPassword() *v3.DataSource {
	if m != nil {
		return m.DownstreamAuthPassword
	}
	return nil
}

// Redis connection pool settings.
// [#next-free-field: 9]
type RedisProxy_ConnPoolSettings struct {
	// Per-operation timeout in milliseconds. The timer starts when the first
	// command of a pipeline is written to the backend connection. Each response received from Redis
	// resets the timer since it signifies that the next command is being processed by the backend.
	// The only exception to this behavior is when a connection to a backend is not yet established.
	// In that case, the connect timeout on the cluster will govern the timeout until the connection
	// is ready.
	OpTimeout *duration.Duration `protobuf:"bytes,1,opt,name=op_timeout,json=opTimeout,proto3" json:"op_timeout,omitempty"`
	// Use hash tagging on every redis key to guarantee that keys with the same hash tag will be
	// forwarded to the same upstream. The hash key used for determining the upstream in a
	// consistent hash ring configuration will be computed from the hash tagged key instead of the
	// whole key. The algorithm used to compute the hash tag is identical to the `redis-cluster
	// implementation <https://redis.io/topics/cluster-spec#keys-hash-tags>`_.
	//
	// Examples:
	//
	// * '{user1000}.following' and '{user1000}.followers' **will** be sent to the same upstream
	// * '{user1000}.following' and '{user1001}.following' **might** be sent to the same upstream
	EnableHashtagging bool `protobuf:"varint,2,opt,name=enable_hashtagging,json=enableHashtagging,proto3" json:"enable_hashtagging,omitempty"`
	// Accept `moved and ask redirection
	// <https://redis.io/topics/cluster-spec#redirection-and-resharding>`_ errors from upstream
	// redis servers, and retry commands to the specified target server. The target server does not
	// need to be known to the cluster manager. If the command cannot be redirected, then the
	// original error is passed downstream unchanged. By default, this support is not enabled.
	EnableRedirection bool `protobuf:"varint,3,opt,name=enable_redirection,json=enableRedirection,proto3" json:"enable_redirection,omitempty"`
	// Maximum size of encoded request buffer before flush is triggered and encoded requests
	// are sent upstream. If this is unset, the buffer flushes whenever it receives data
	// and performs no batching.
	// This feature makes it possible for multiple clients to send requests to Envoy and have
	// them batched- for example if one is running several worker processes, each with its own
	// Redis connection. There is no benefit to using this with a single downstream process.
	// Recommended size (if enabled) is 1024 bytes.
	MaxBufferSizeBeforeFlush uint32 `protobuf:"varint,4,opt,name=max_buffer_size_before_flush,json=maxBufferSizeBeforeFlush,proto3" json:"max_buffer_size_before_flush,omitempty"`
	// The encoded request buffer is flushed N milliseconds after the first request has been
	// encoded, unless the buffer size has already exceeded `max_buffer_size_before_flush`.
	// If `max_buffer_size_before_flush` is not set, this flush timer is not used. Otherwise,
	// the timer should be set according to the number of clients, overall request rate and
	// desired maximum latency for a single command. For example, if there are many requests
	// being batched together at a high rate, the buffer will likely be filled before the timer
	// fires. Alternatively, if the request rate is lower the buffer will not be filled as often
	// before the timer fires.
	// If `max_buffer_size_before_flush` is set, but `buffer_flush_timeout` is not, the latter
	// defaults to 3ms.
	BufferFlushTimeout *duration.Duration `protobuf:"bytes,5,opt,name=buffer_flush_timeout,json=bufferFlushTimeout,proto3" json:"buffer_flush_timeout,omitempty"`
	// `max_upstream_unknown_connections` controls how many upstream connections to unknown hosts
	// can be created at any given time by any given worker thread (see `enable_redirection` for
	// more details). If the host is unknown and a connection cannot be created due to enforcing
	// this limit, then redirection will fail and the original redirection error will be passed
	// downstream unchanged. This limit defaults to 100.
	MaxUpstreamUnknownConnections *wrappers.UInt32Value `protobuf:"bytes,6,opt,name=max_upstream_unknown_connections,json=maxUpstreamUnknownConnections,proto3" json:"max_upstream_unknown_connections,omitempty"`
	// Enable per-command statistics per upstream cluster, in addition to the filter level aggregate
	// count.
	EnableCommandStats bool `protobuf:"varint,8,opt,name=enable_command_stats,json=enableCommandStats,proto3" json:"enable_command_stats,omitempty"`
	// Read policy. The default is to read from the master.
	ReadPolicy           RedisProxy_ConnPoolSettings_ReadPolicy `protobuf:"varint,7,opt,name=read_policy,json=readPolicy,proto3,enum=envoy.extensions.filters.network.redis_proxy.v3.RedisProxy_ConnPoolSettings_ReadPolicy" json:"read_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *RedisProxy_ConnPoolSettings) Reset()         { *m = RedisProxy_ConnPoolSettings{} }
func (m *RedisProxy_ConnPoolSettings) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_ConnPoolSettings) ProtoMessage()    {}
func (*RedisProxy_ConnPoolSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0, 0}
}

func (m *RedisProxy_ConnPoolSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Unmarshal(m, b)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Marshal(b, m, deterministic)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.Merge(m, src)
}
func (m *RedisProxy_ConnPoolSettings) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_ConnPoolSettings.Size(m)
}
func (m *RedisProxy_ConnPoolSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_ConnPoolSettings.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_ConnPoolSettings proto.InternalMessageInfo

func (m *RedisProxy_ConnPoolSettings) GetOpTimeout() *duration.Duration {
	if m != nil {
		return m.OpTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetEnableHashtagging() bool {
	if m != nil {
		return m.EnableHashtagging
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetEnableRedirection() bool {
	if m != nil {
		return m.EnableRedirection
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetMaxBufferSizeBeforeFlush() uint32 {
	if m != nil {
		return m.MaxBufferSizeBeforeFlush
	}
	return 0
}

func (m *RedisProxy_ConnPoolSettings) GetBufferFlushTimeout() *duration.Duration {
	if m != nil {
		return m.BufferFlushTimeout
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetMaxUpstreamUnknownConnections() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxUpstreamUnknownConnections
	}
	return nil
}

func (m *RedisProxy_ConnPoolSettings) GetEnableCommandStats() bool {
	if m != nil {
		return m.EnableCommandStats
	}
	return false
}

func (m *RedisProxy_ConnPoolSettings) GetReadPolicy() RedisProxy_ConnPoolSettings_ReadPolicy {
	if m != nil {
		return m.ReadPolicy
	}
	return RedisProxy_ConnPoolSettings_MASTER
}

type RedisProxy_PrefixRoutes struct {
	// List of prefix routes.
	Routes []*RedisProxy_PrefixRoutes_Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	// Indicates that prefix matching should be case insensitive.
	CaseInsensitive bool `protobuf:"varint,2,opt,name=case_insensitive,json=caseInsensitive,proto3" json:"case_insensitive,omitempty"`
	// Optional catch-all route to forward commands that doesn't match any of the routes. The
	// catch-all route becomes required when no routes are specified.
	CatchAllRoute        *RedisProxy_PrefixRoutes_Route `protobuf:"bytes,4,opt,name=catch_all_route,json=catchAllRoute,proto3" json:"catch_all_route,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *RedisProxy_PrefixRoutes) Reset()         { *m = RedisProxy_PrefixRoutes{} }
func (m *RedisProxy_PrefixRoutes) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0, 1}
}

func (m *RedisProxy_PrefixRoutes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes.Size(m)
}
func (m *RedisProxy_PrefixRoutes) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes) GetRoutes() []*RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes) GetCaseInsensitive() bool {
	if m != nil {
		return m.CaseInsensitive
	}
	return false
}

func (m *RedisProxy_PrefixRoutes) GetCatchAllRoute() *RedisProxy_PrefixRoutes_Route {
	if m != nil {
		return m.CatchAllRoute
	}
	return nil
}

type RedisProxy_PrefixRoutes_Route struct {
	// String prefix that must match the beginning of the keys. Envoy will always favor the
	// longest match.
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Indicates if the prefix needs to be removed from the key when forwarded.
	RemovePrefix bool `protobuf:"varint,2,opt,name=remove_prefix,json=removePrefix,proto3" json:"remove_prefix,omitempty"`
	// Upstream cluster to forward the command to.
	Cluster string `protobuf:"bytes,3,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// Indicates that the route has a request mirroring policy.
	RequestMirrorPolicy  []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy `protobuf:"bytes,4,rep,name=request_mirror_policy,json=requestMirrorPolicy,proto3" json:"request_mirror_policy,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                             `json:"-"`
	XXX_unrecognized     []byte                                               `json:"-"`
	XXX_sizecache        int32                                                `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route) Reset()         { *m = RedisProxy_PrefixRoutes_Route{} }
func (m *RedisProxy_PrefixRoutes_Route) String() string { return proto.CompactTextString(m) }
func (*RedisProxy_PrefixRoutes_Route) ProtoMessage()    {}
func (*RedisProxy_PrefixRoutes_Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0, 1, 0}
}

func (m *RedisProxy_PrefixRoutes_Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRemovePrefix() bool {
	if m != nil {
		return m.RemovePrefix
	}
	return false
}

func (m *RedisProxy_PrefixRoutes_Route) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route) GetRequestMirrorPolicy() []*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy {
	if m != nil {
		return m.RequestMirrorPolicy
	}
	return nil
}

// The router is capable of shadowing traffic from one cluster to another. The current
// implementation is "fire and forget," meaning Envoy will not wait for the shadow cluster to
// respond before returning the response from the primary cluster. All normal statistics are
// collected for the shadow cluster making this feature useful for testing.
type RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy struct {
	// Specifies the cluster that requests will be mirrored to. The cluster must
	// exist in the cluster manager configuration.
	Cluster string `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	// If not specified or the runtime key is not present, all requests to the target cluster
	// will be mirrored.
	//
	// If specified, Envoy will lookup the runtime key to get the percentage of requests to the
	// mirror.
	RuntimeFraction *v3.RuntimeFractionalPercent `protobuf:"bytes,2,opt,name=runtime_fraction,json=runtimeFraction,proto3" json:"runtime_fraction,omitempty"`
	// Set this to TRUE to only mirror write commands, this is effectively replicating the
	// writes in a "fire and forget" manner.
	ExcludeReadCommands  bool     `protobuf:"varint,3,opt,name=exclude_read_commands,json=excludeReadCommands,proto3" json:"exclude_read_commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Reset() {
	*m = RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy{}
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) String() string {
	return proto.CompactTextString(m)
}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) ProtoMessage() {}
func (*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{0, 1, 0, 0}
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Unmarshal(m, b)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Marshal(b, m, deterministic)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Merge(m, src)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_Size() int {
	return xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.Size(m)
}
func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy proto.InternalMessageInfo

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetCluster() string {
	if m != nil {
		return m.Cluster
	}
	return ""
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetRuntimeFraction() *v3.RuntimeFractionalPercent {
	if m != nil {
		return m.RuntimeFraction
	}
	return nil
}

func (m *RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy) GetExcludeReadCommands() bool {
	if m != nil {
		return m.ExcludeReadCommands
	}
	return false
}

// RedisProtocolOptions specifies Redis upstream protocol options. This object is used in
// :ref:`typed_extension_protocol_options<envoy_api_field_config.cluster.v3.Cluster.typed_extension_protocol_options>`,
// keyed by the name `envoy.redis_proxy`.
type RedisProtocolOptions struct {
	// Upstream server password as defined by the `requirepass` directive
	// <https://redis.io/topics/config>`_ in the server's configuration file.
	AuthPassword         *v3.DataSource `protobuf:"bytes,1,opt,name=auth_password,json=authPassword,proto3" json:"auth_password,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RedisProtocolOptions) Reset()         { *m = RedisProtocolOptions{} }
func (m *RedisProtocolOptions) String() string { return proto.CompactTextString(m) }
func (*RedisProtocolOptions) ProtoMessage()    {}
func (*RedisProtocolOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2438571bcc1970a0, []int{1}
}

func (m *RedisProtocolOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RedisProtocolOptions.Unmarshal(m, b)
}
func (m *RedisProtocolOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RedisProtocolOptions.Marshal(b, m, deterministic)
}
func (m *RedisProtocolOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RedisProtocolOptions.Merge(m, src)
}
func (m *RedisProtocolOptions) XXX_Size() int {
	return xxx_messageInfo_RedisProtocolOptions.Size(m)
}
func (m *RedisProtocolOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_RedisProtocolOptions.DiscardUnknown(m)
}

var xxx_messageInfo_RedisProtocolOptions proto.InternalMessageInfo

func (m *RedisProtocolOptions) GetAuthPassword() *v3.DataSource {
	if m != nil {
		return m.AuthPassword
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.extensions.filters.network.redis_proxy.v3.RedisProxy_ConnPoolSettings_ReadPolicy", RedisProxy_ConnPoolSettings_ReadPolicy_name, RedisProxy_ConnPoolSettings_ReadPolicy_value)
	proto.RegisterType((*RedisProxy)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProxy")
	proto.RegisterType((*RedisProxy_ConnPoolSettings)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProxy.ConnPoolSettings")
	proto.RegisterType((*RedisProxy_PrefixRoutes)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProxy.PrefixRoutes")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProxy.PrefixRoutes.Route")
	proto.RegisterType((*RedisProxy_PrefixRoutes_Route_RequestMirrorPolicy)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProxy.PrefixRoutes.Route.RequestMirrorPolicy")
	proto.RegisterType((*RedisProtocolOptions)(nil), "envoy.extensions.filters.network.redis_proxy.v3.RedisProtocolOptions")
}

func init() {
	proto.RegisterFile("envoy/extensions/filters/network/redis_proxy/v3/redis_proxy.proto", fileDescriptor_2438571bcc1970a0)
}

var fileDescriptor_2438571bcc1970a0 = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcb, 0x6f, 0x1b, 0x45,
	0x18, 0x67, 0x6d, 0xc7, 0x71, 0xc7, 0x71, 0xe3, 0x4c, 0x1f, 0x18, 0xab, 0x14, 0x37, 0xbd, 0x18,
	0x24, 0x76, 0x91, 0x23, 0x84, 0x64, 0xf1, 0x90, 0x9d, 0x87, 0x92, 0xd0, 0xb4, 0xd6, 0xa4, 0x01,
	0x55, 0x42, 0xac, 0xc6, 0xbb, 0x63, 0x7b, 0xd5, 0xf5, 0xcc, 0x32, 0x33, 0xeb, 0x38, 0x3d, 0x95,
	0x4a, 0x48, 0x08, 0xee, 0x1c, 0x38, 0x73, 0x42, 0x70, 0x86, 0x7f, 0x80, 0x7f, 0xaa, 0x27, 0x34,
	0x0f, 0xdb, 0x1b, 0x92, 0x02, 0x49, 0x73, 0xda, 0x9d, 0xef, 0xf7, 0xcd, 0xef, 0x7b, 0x7f, 0x03,
	0x3a, 0x84, 0x4e, 0xd8, 0x89, 0x47, 0xa6, 0x92, 0x50, 0x11, 0x31, 0x2a, 0xbc, 0x41, 0x14, 0x4b,
	0xc2, 0x85, 0x47, 0x89, 0x3c, 0x66, 0xfc, 0xa9, 0xc7, 0x49, 0x18, 0x09, 0x3f, 0xe1, 0x6c, 0x7a,
	0xe2, 0x4d, 0x36, 0xb2, 0x47, 0x37, 0xe1, 0x4c, 0x32, 0xe8, 0x69, 0x0a, 0x77, 0x41, 0xe1, 0x5a,
	0x0a, 0xd7, 0x52, 0xb8, 0xd9, 0x3b, 0x93, 0x8d, 0xfa, 0x3b, 0xc6, 0x66, 0xc0, 0xe8, 0x20, 0x1a,
	0x7a, 0x01, 0xe3, 0x44, 0x11, 0xf7, 0xb1, 0x20, 0x86, 0xb1, 0x7e, 0x77, 0xc8, 0xd8, 0x30, 0x26,
	0x9e, 0x3e, 0xf5, 0xd3, 0x81, 0x17, 0xa6, 0x1c, 0xcb, 0x88, 0xd1, 0x57, 0xe1, 0xc7, 0x1c, 0x27,
	0x89, 0xb2, 0x68, 0xf0, 0x46, 0x1a, 0x26, 0xd8, 0xc3, 0x94, 0x32, 0xa9, 0xaf, 0x09, 0x4f, 0x28,
	0xd7, 0x64, 0x34, 0x99, 0x59, 0xb8, 0x77, 0x46, 0x63, 0x42, 0xb8, 0x72, 0x3e, 0xa2, 0x43, 0xab,
	0x72, 0xdf, 0x78, 0x99, 0xd5, 0x09, 0x49, 0xc2, 0x49, 0x90, 0xf5, 0xe4, 0xcd, 0x09, 0x8e, 0xa3,
	0x10, 0x4b, 0xe2, 0xcd, 0x7e, 0x0c, 0xb0, 0xfe, 0x2d, 0x04, 0x00, 0xa9, 0xb0, 0x7b, 0x2a, 0x6a,
	0xd8, 0x04, 0x65, 0x21, 0xb1, 0xf4, 0x13, 0x4e, 0x06, 0xd1, 0xb4, 0xe6, 0x34, 0x9c, 0xe6, 0xb5,
	0xee, 0xf2, 0xcb, 0x6e, 0x81, 0xe7, 0x1a, 0x0e, 0x02, 0x0a, 0xeb, 0x69, 0x08, 0x72, 0x50, 0x12,
	0x44, 0xca, 0x88, 0x0e, 0x45, 0x2d, 0xdf, 0x70, 0x9a, 0xe5, 0xd6, 0x03, 0xf7, 0x82, 0x09, 0x76,
	0x17, 0x86, 0xdd, 0x4d, 0x46, 0x69, 0x8f, 0xb1, 0xf8, 0xd0, 0x72, 0x76, 0x4b, 0x2f, 0xbb, 0x4b,
	0x3f, 0x38, 0xb9, 0xaa, 0x83, 0xe6, 0x76, 0xe0, 0x7b, 0x60, 0x2d, 0xc6, 0x92, 0xd0, 0xe0, 0xc4,
	0x8f, 0xa8, 0x3f, 0x8e, 0x02, 0xce, 0x44, 0xad, 0xd0, 0x70, 0x9a, 0x25, 0xb4, 0x6a, 0x81, 0x3d,
	0x7a, 0xa0, 0xc5, 0x70, 0x0c, 0x2a, 0x26, 0x08, 0x9f, 0xb3, 0x54, 0x12, 0x51, 0x5b, 0xd2, 0x4e,
	0xee, 0xbe, 0x8e, 0x93, 0x26, 0x74, 0xa4, 0xf9, 0xd0, 0x4a, 0x92, 0x39, 0xc1, 0x3e, 0xa8, 0x85,
	0xec, 0x98, 0x0a, 0xc9, 0x09, 0x1e, 0xfb, 0x38, 0x95, 0x23, 0x3f, 0xc1, 0x42, 0x1c, 0x33, 0x1e,
	0xd6, 0x8a, 0xda, 0x72, 0xc3, 0x5a, 0x36, 0xed, 0xe4, 0xaa, 0x76, 0x52, 0xf4, 0x5b, 0x58, 0xe2,
	0x43, 0x96, 0xf2, 0x80, 0x74, 0x8b, 0x7f, 0xfe, 0xf1, 0xe3, 0x2f, 0x39, 0x07, 0xdd, 0x5e, 0x30,
	0x75, 0x52, 0x39, 0xea, 0x59, 0x9e, 0xfa, 0x77, 0x45, 0x50, 0xfd, 0x67, 0x9e, 0x60, 0x17, 0x00,
	0x96, 0xf8, 0x32, 0x1a, 0x13, 0x96, 0x4a, 0x5d, 0xb0, 0x72, 0xeb, 0x2d, 0xd7, 0x34, 0x9e, 0x3b,
	0x6b, 0x3c, 0x77, 0xcb, 0x36, 0xa6, 0x4e, 0xeb, 0xaf, 0x4e, 0xae, 0xe4, 0xa0, 0x6b, 0x2c, 0x79,
	0x6c, 0x6e, 0xc1, 0xf7, 0x01, 0x24, 0x14, 0xf7, 0x63, 0xe2, 0x8f, 0xb0, 0x18, 0x49, 0x3c, 0x1c,
	0x46, 0x74, 0x58, 0xcb, 0xe9, 0xc4, 0xae, 0x19, 0x64, 0x77, 0x01, 0x64, 0xd4, 0x55, 0xaa, 0x38,
	0x09, 0x14, 0xb3, 0x6e, 0x82, 0xb9, 0x3a, 0x5a, 0x00, 0xf0, 0x53, 0x70, 0x67, 0x8c, 0xa7, 0x7e,
	0x3f, 0x1d, 0x0c, 0x08, 0xf7, 0x45, 0xf4, 0x8c, 0xf8, 0x7d, 0x32, 0x60, 0x9c, 0xf8, 0x83, 0x38,
	0x15, 0x23, 0x5d, 0xc0, 0x0a, 0xaa, 0x8d, 0xf1, 0xb4, 0xab, 0x55, 0x0e, 0xa3, 0x67, 0xa4, 0xab,
	0x15, 0x76, 0x14, 0x0e, 0x3f, 0x07, 0x37, 0xed, 0x5d, 0xad, 0x3f, 0x8f, 0x75, 0xe9, 0x3f, 0x62,
	0x45, 0xd0, 0x5c, 0xd3, 0x2c, 0xb3, 0x50, 0x09, 0x68, 0x28, 0x67, 0xd2, 0xc4, 0x56, 0x2a, 0xa5,
	0x4f, 0x29, 0x3b, 0xa6, 0x7e, 0xc0, 0x28, 0x35, 0xfe, 0x0a, 0x5b, 0xaf, 0x3b, 0x67, 0x88, 0x8f,
	0xf6, 0xa8, 0xdc, 0x68, 0x7d, 0x81, 0xe3, 0x94, 0xa0, 0xb7, 0xc7, 0x78, 0x7a, 0x64, 0x49, 0x8e,
	0x0c, 0xc7, 0xe6, 0x82, 0x02, 0x7e, 0x00, 0x6e, 0xda, 0x14, 0x05, 0x6c, 0x3c, 0xc6, 0x34, 0xf4,
	0xd5, 0xe8, 0x88, 0x5a, 0x49, 0x27, 0xc9, 0xa6, 0x6f, 0xd3, 0x40, 0x87, 0x0a, 0x81, 0x2f, 0x1c,
	0x50, 0xe6, 0x04, 0x87, 0x7e, 0xc2, 0xe2, 0x28, 0x38, 0xa9, 0x2d, 0x37, 0x9c, 0xe6, 0xf5, 0xd6,
	0x97, 0x57, 0x39, 0x53, 0x2e, 0x22, 0x38, 0xec, 0x69, 0x7a, 0xdd, 0x07, 0x2f, 0xf4, 0x78, 0x01,
	0x3e, 0x97, 0xae, 0x1f, 0xa9, 0x65, 0x30, 0x3b, 0x41, 0x00, 0x8a, 0x07, 0x9d, 0xc3, 0xc7, 0xdb,
	0xa8, 0xfa, 0x06, 0x5c, 0x03, 0x95, 0x1e, 0xda, 0xde, 0xd9, 0x46, 0xbe, 0x15, 0x39, 0xb0, 0x0c,
	0x96, 0xd1, 0x76, 0xef, 0xc1, 0xde, 0x66, 0xa7, 0x9a, 0x83, 0x10, 0x5c, 0xb7, 0xf8, 0x4c, 0x96,
	0x87, 0xcb, 0x20, 0xdf, 0x79, 0xf8, 0xa4, 0x5a, 0x68, 0x1f, 0xfc, 0xfc, 0xd7, 0xf7, 0x77, 0x77,
	0xc1, 0xce, 0xa9, 0x01, 0x30, 0x71, 0x9c, 0x1f, 0x46, 0xeb, 0xdf, 0xc2, 0xa8, 0x3f, 0x2f, 0x81,
	0x95, 0xec, 0x28, 0xc2, 0x01, 0x28, 0xda, 0x21, 0x77, 0x1a, 0xf9, 0x66, 0xb9, 0xf5, 0xf0, 0xaa,
	0x86, 0xdc, 0xd5, 0x1f, 0x64, 0xd9, 0xe1, 0xbb, 0xa0, 0x1a, 0x60, 0x41, 0xfc, 0x88, 0xce, 0xf7,
	0xb4, 0x9d, 0x92, 0x55, 0x25, 0xdf, 0x5b, 0x88, 0xe1, 0x04, 0xac, 0x06, 0x58, 0x06, 0x23, 0x1f,
	0xc7, 0xb1, 0xd9, 0x40, 0xba, 0xcf, 0xaf, 0xde, 0xb7, 0x8a, 0x36, 0xd3, 0x89, 0x63, 0x7d, 0xac,
	0x3f, 0x5f, 0x02, 0x4b, 0xfa, 0x0f, 0xde, 0x06, 0xc5, 0xec, 0x16, 0x47, 0xf6, 0x04, 0xef, 0x83,
	0x0a, 0x27, 0x63, 0x36, 0x21, 0xb3, 0x25, 0x6f, 0x22, 0x58, 0x31, 0x42, 0xbb, 0xdd, 0xef, 0x81,
	0xe5, 0x20, 0x4e, 0x85, 0x24, 0x5c, 0xcf, 0x75, 0xe6, 0x0d, 0x98, 0xc9, 0xe1, 0x4f, 0x0e, 0xb8,
	0xc5, 0xc9, 0x37, 0x29, 0x11, 0xd2, 0x1f, 0x47, 0x9c, 0x33, 0x3e, 0x6b, 0xdd, 0x82, 0x2e, 0x42,
	0xff, 0x6a, 0x03, 0x75, 0x91, 0xb1, 0x75, 0xa0, 0x4d, 0x99, 0x0e, 0x45, 0x37, 0xf8, 0x59, 0x61,
	0xfd, 0xb7, 0x1c, 0xb8, 0x71, 0x8e, 0x72, 0x36, 0x26, 0xe7, 0x15, 0x31, 0x3d, 0x01, 0x55, 0x9e,
	0x52, 0xb5, 0x60, 0xfc, 0x01, 0xc7, 0x66, 0xaf, 0xe5, 0x74, 0xd9, 0xdc, 0xf3, 0xb7, 0x37, 0x32,
	0xda, 0x3b, 0x56, 0x19, 0xc7, 0x3d, 0xc2, 0x03, 0x42, 0x25, 0x5a, 0xe5, 0xa7, 0x11, 0xd8, 0x02,
	0xb7, 0xc8, 0x34, 0x88, 0xd3, 0x50, 0x6d, 0x4d, 0x1c, 0xce, 0xf6, 0x82, 0xb0, 0x7b, 0xf3, 0x86,
	0x05, 0xd5, 0xf8, 0xd9, 0xbd, 0x20, 0xda, 0x81, 0x9a, 0x9b, 0xaf, 0xc1, 0x57, 0x97, 0x9b, 0x9b,
	0xff, 0x97, 0xc3, 0xf6, 0x23, 0x65, 0x64, 0x1f, 0xec, 0x5e, 0x95, 0x91, 0xf6, 0x9e, 0x22, 0xdc,
	0x02, 0xdd, 0xd7, 0x27, 0xdc, 0x2f, 0x94, 0xf2, 0xd5, 0x02, 0x5a, 0x5b, 0x4c, 0x92, 0x2d, 0x54,
	0xfb, 0x63, 0x65, 0xe3, 0x23, 0xf0, 0xe1, 0xa5, 0x6c, 0xec, 0x17, 0x4a, 0xb9, 0x6a, 0x7e, 0x5e,
	0xf5, 0xf5, 0xdf, 0x1d, 0x70, 0x73, 0x86, 0x4a, 0x16, 0xb0, 0xf8, 0x51, 0x62, 0xb6, 0xf8, 0x01,
	0xa8, 0x9c, 0x7e, 0xc9, 0x9d, 0x0b, 0xbe, 0xe4, 0x2b, 0x38, 0xf3, 0x7e, 0xb7, 0x77, 0x94, 0xd3,
	0x1d, 0xf0, 0xd9, 0x25, 0x9c, 0xce, 0xba, 0xd5, 0x45, 0xe0, 0x93, 0x88, 0x19, 0x1f, 0x8c, 0xe2,
	0x05, 0x07, 0xad, 0xbb, 0xba, 0xc8, 0x85, 0xe6, 0xee, 0x39, 0xfd, 0xa2, 0x7e, 0xe5, 0x36, 0xfe,
	0x0e, 0x00, 0x00, 0xff, 0xff, 0x17, 0x18, 0x41, 0x8e, 0x68, 0x0b, 0x00, 0x00,
}