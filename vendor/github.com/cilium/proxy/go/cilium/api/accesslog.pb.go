// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v4.24.4
// source: cilium/api/accesslog.proto

package cilium

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type HttpProtocol int32

const (
	HttpProtocol_HTTP10 HttpProtocol = 0
	HttpProtocol_HTTP11 HttpProtocol = 1
	HttpProtocol_HTTP2  HttpProtocol = 2
)

// Enum value maps for HttpProtocol.
var (
	HttpProtocol_name = map[int32]string{
		0: "HTTP10",
		1: "HTTP11",
		2: "HTTP2",
	}
	HttpProtocol_value = map[string]int32{
		"HTTP10": 0,
		"HTTP11": 1,
		"HTTP2":  2,
	}
)

func (x HttpProtocol) Enum() *HttpProtocol {
	p := new(HttpProtocol)
	*p = x
	return p
}

func (x HttpProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HttpProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_cilium_api_accesslog_proto_enumTypes[0].Descriptor()
}

func (HttpProtocol) Type() protoreflect.EnumType {
	return &file_cilium_api_accesslog_proto_enumTypes[0]
}

func (x HttpProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HttpProtocol.Descriptor instead.
func (HttpProtocol) EnumDescriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{0}
}

type EntryType int32

const (
	EntryType_Request  EntryType = 0
	EntryType_Response EntryType = 1
	EntryType_Denied   EntryType = 2
)

// Enum value maps for EntryType.
var (
	EntryType_name = map[int32]string{
		0: "Request",
		1: "Response",
		2: "Denied",
	}
	EntryType_value = map[string]int32{
		"Request":  0,
		"Response": 1,
		"Denied":   2,
	}
)

func (x EntryType) Enum() *EntryType {
	p := new(EntryType)
	*p = x
	return p
}

func (x EntryType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EntryType) Descriptor() protoreflect.EnumDescriptor {
	return file_cilium_api_accesslog_proto_enumTypes[1].Descriptor()
}

func (EntryType) Type() protoreflect.EnumType {
	return &file_cilium_api_accesslog_proto_enumTypes[1]
}

func (x EntryType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EntryType.Descriptor instead.
func (EntryType) EnumDescriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{1}
}

type KeyValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *KeyValue) Reset() {
	*x = KeyValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_accesslog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValue) ProtoMessage() {}

func (x *KeyValue) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_accesslog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValue.ProtoReflect.Descriptor instead.
func (*KeyValue) Descriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{0}
}

func (x *KeyValue) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KeyValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type HttpLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HttpProtocol HttpProtocol `protobuf:"varint,1,opt,name=http_protocol,json=httpProtocol,proto3,enum=cilium.HttpProtocol" json:"http_protocol,omitempty"`
	// Request info that is also retained for the response
	Scheme string `protobuf:"bytes,2,opt,name=scheme,proto3" json:"scheme,omitempty"` // Envoy "x-forwarded-proto", e.g., "http", "https"
	Host   string `protobuf:"bytes,3,opt,name=host,proto3" json:"host,omitempty"`     // Envoy ":authority" header
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`     // Envoy ":path" header
	Method string `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"` // Envoy ":method" header
	// Request or response headers not included above
	Headers []*KeyValue `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	// Response info
	Status uint32 `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"` // Envoy ":status" header, zero for request
	// missing_headers includes both headers that were added to the
	// request, and headers that were merely logged as missing
	MissingHeaders []*KeyValue `protobuf:"bytes,8,rep,name=missing_headers,json=missingHeaders,proto3" json:"missing_headers,omitempty"`
	// rejected_headers includes headers that were flagged as unallowed,
	// which may have been removed, or merely logged and the request still
	// allowed, or the request may have been dropped due to them.
	RejectedHeaders []*KeyValue `protobuf:"bytes,9,rep,name=rejected_headers,json=rejectedHeaders,proto3" json:"rejected_headers,omitempty"`
}

func (x *HttpLogEntry) Reset() {
	*x = HttpLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_accesslog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HttpLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HttpLogEntry) ProtoMessage() {}

func (x *HttpLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_accesslog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HttpLogEntry.ProtoReflect.Descriptor instead.
func (*HttpLogEntry) Descriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{1}
}

func (x *HttpLogEntry) GetHttpProtocol() HttpProtocol {
	if x != nil {
		return x.HttpProtocol
	}
	return HttpProtocol_HTTP10
}

func (x *HttpLogEntry) GetScheme() string {
	if x != nil {
		return x.Scheme
	}
	return ""
}

func (x *HttpLogEntry) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *HttpLogEntry) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *HttpLogEntry) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *HttpLogEntry) GetHeaders() []*KeyValue {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *HttpLogEntry) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *HttpLogEntry) GetMissingHeaders() []*KeyValue {
	if x != nil {
		return x.MissingHeaders
	}
	return nil
}

func (x *HttpLogEntry) GetRejectedHeaders() []*KeyValue {
	if x != nil {
		return x.RejectedHeaders
	}
	return nil
}

type KafkaLogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// correlation_id is a user-supplied integer value that will be passed
	// back with the response
	CorrelationId int32 `protobuf:"varint,1,opt,name=correlation_id,json=correlationId,proto3" json:"correlation_id,omitempty"`
	// error_code is the Kafka error code being returned
	// Ref. https://kafka.apache.org/protocol#protocol_error_codes
	ErrorCode int32 `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	// api_version of the Kafka api used
	// Ref. https://kafka.apache.org/protocol#protocol_compatibility
	ApiVersion int32 `protobuf:"varint,3,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	// api_key for Kafka message
	// Reference: https://kafka.apache.org/protocol#protocol_api_keys
	ApiKey int32 `protobuf:"varint,4,opt,name=api_key,json=apiKey,proto3" json:"api_key,omitempty"`
	// Topics of the request
	// Optional, as not all messages have topics (ex. LeaveGroup, Heartbeat)
	Topics []string `protobuf:"bytes,5,rep,name=topics,proto3" json:"topics,omitempty"`
}

func (x *KafkaLogEntry) Reset() {
	*x = KafkaLogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_accesslog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaLogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaLogEntry) ProtoMessage() {}

func (x *KafkaLogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_accesslog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaLogEntry.ProtoReflect.Descriptor instead.
func (*KafkaLogEntry) Descriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{2}
}

func (x *KafkaLogEntry) GetCorrelationId() int32 {
	if x != nil {
		return x.CorrelationId
	}
	return 0
}

func (x *KafkaLogEntry) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *KafkaLogEntry) GetApiVersion() int32 {
	if x != nil {
		return x.ApiVersion
	}
	return 0
}

func (x *KafkaLogEntry) GetApiKey() int32 {
	if x != nil {
		return x.ApiKey
	}
	return 0
}

func (x *KafkaLogEntry) GetTopics() []string {
	if x != nil {
		return x.Topics
	}
	return nil
}

type L7LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Proto  string            `protobuf:"bytes,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Fields map[string]string `protobuf:"bytes,2,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *L7LogEntry) Reset() {
	*x = L7LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_accesslog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *L7LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*L7LogEntry) ProtoMessage() {}

func (x *L7LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_accesslog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use L7LogEntry.ProtoReflect.Descriptor instead.
func (*L7LogEntry) Descriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{3}
}

func (x *L7LogEntry) GetProto() string {
	if x != nil {
		return x.Proto
	}
	return ""
}

func (x *L7LogEntry) GetFields() map[string]string {
	if x != nil {
		return x.Fields
	}
	return nil
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The time that Cilium filter captured this log entry,
	// in, nanoseconds since 1/1/1970.
	Timestamp uint64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// 'true' if the request was received by an ingress listener,
	// 'false' if received by an egress listener
	IsIngress bool      `protobuf:"varint,15,opt,name=is_ingress,json=isIngress,proto3" json:"is_ingress,omitempty"`
	EntryType EntryType `protobuf:"varint,3,opt,name=entry_type,json=entryType,proto3,enum=cilium.EntryType" json:"entry_type,omitempty"`
	// Cilium network policy resource name
	PolicyName string `protobuf:"bytes,4,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// proxy_id identifies the listener this message relates to,
	// as configured via the bpf_metadata listener filter
	ProxyId uint32 `protobuf:"varint,17,opt,name=proxy_id,json=proxyId,proto3" json:"proxy_id,omitempty"`
	// Cilium rule reference
	CiliumRuleRef string `protobuf:"bytes,5,opt,name=cilium_rule_ref,json=ciliumRuleRef,proto3" json:"cilium_rule_ref,omitempty"`
	// Cilium security ID of the source and destination
	SourceSecurityId      uint32 `protobuf:"varint,6,opt,name=source_security_id,json=sourceSecurityId,proto3" json:"source_security_id,omitempty"`
	DestinationSecurityId uint32 `protobuf:"varint,16,opt,name=destination_security_id,json=destinationSecurityId,proto3" json:"destination_security_id,omitempty"`
	// These fields record the original source and destination addresses,
	// stored in ipv4:port or [ipv6]:port format.
	SourceAddress      string `protobuf:"bytes,7,opt,name=source_address,json=sourceAddress,proto3" json:"source_address,omitempty"`
	DestinationAddress string `protobuf:"bytes,8,opt,name=destination_address,json=destinationAddress,proto3" json:"destination_address,omitempty"`
	// Types that are assignable to L7:
	//
	//	*LogEntry_Http
	//	*LogEntry_Kafka
	//	*LogEntry_GenericL7
	L7 isLogEntry_L7 `protobuf_oneof:"l7"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cilium_api_accesslog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_cilium_api_accesslog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_cilium_api_accesslog_proto_rawDescGZIP(), []int{4}
}

func (x *LogEntry) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogEntry) GetIsIngress() bool {
	if x != nil {
		return x.IsIngress
	}
	return false
}

func (x *LogEntry) GetEntryType() EntryType {
	if x != nil {
		return x.EntryType
	}
	return EntryType_Request
}

func (x *LogEntry) GetPolicyName() string {
	if x != nil {
		return x.PolicyName
	}
	return ""
}

func (x *LogEntry) GetProxyId() uint32 {
	if x != nil {
		return x.ProxyId
	}
	return 0
}

func (x *LogEntry) GetCiliumRuleRef() string {
	if x != nil {
		return x.CiliumRuleRef
	}
	return ""
}

func (x *LogEntry) GetSourceSecurityId() uint32 {
	if x != nil {
		return x.SourceSecurityId
	}
	return 0
}

func (x *LogEntry) GetDestinationSecurityId() uint32 {
	if x != nil {
		return x.DestinationSecurityId
	}
	return 0
}

func (x *LogEntry) GetSourceAddress() string {
	if x != nil {
		return x.SourceAddress
	}
	return ""
}

func (x *LogEntry) GetDestinationAddress() string {
	if x != nil {
		return x.DestinationAddress
	}
	return ""
}

func (m *LogEntry) GetL7() isLogEntry_L7 {
	if m != nil {
		return m.L7
	}
	return nil
}

func (x *LogEntry) GetHttp() *HttpLogEntry {
	if x, ok := x.GetL7().(*LogEntry_Http); ok {
		return x.Http
	}
	return nil
}

func (x *LogEntry) GetKafka() *KafkaLogEntry {
	if x, ok := x.GetL7().(*LogEntry_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (x *LogEntry) GetGenericL7() *L7LogEntry {
	if x, ok := x.GetL7().(*LogEntry_GenericL7); ok {
		return x.GenericL7
	}
	return nil
}

type isLogEntry_L7 interface {
	isLogEntry_L7()
}

type LogEntry_Http struct {
	Http *HttpLogEntry `protobuf:"bytes,100,opt,name=http,proto3,oneof"`
}

type LogEntry_Kafka struct {
	Kafka *KafkaLogEntry `protobuf:"bytes,101,opt,name=kafka,proto3,oneof"`
}

type LogEntry_GenericL7 struct {
	GenericL7 *L7LogEntry `protobuf:"bytes,102,opt,name=generic_l7,json=genericL7,proto3,oneof"`
}

func (*LogEntry_Http) isLogEntry_L7() {}

func (*LogEntry_Kafka) isLogEntry_L7() {}

func (*LogEntry_GenericL7) isLogEntry_L7() {}

var File_cilium_api_accesslog_proto protoreflect.FileDescriptor

var file_cilium_api_accesslog_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x69,
	0x6c, 0x69, 0x75, 0x6d, 0x22, 0x32, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xdd, 0x02, 0x0a, 0x0c, 0x48, 0x74, 0x74,
	0x70, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x39, 0x0a, 0x0d, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x14, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x6f, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2a, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x39, 0x0a, 0x0f, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x5f, 0x68, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x69, 0x6c, 0x69,
	0x75, 0x6d, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6e, 0x67, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x3b, 0x0a, 0x10, 0x72,
	0x65, 0x6a, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x4b,
	0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0f, 0x72, 0x65, 0x6a, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x22, 0xa7, 0x01, 0x0a, 0x0d, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x61, 0x70, 0x69, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x61, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x73, 0x22, 0x95, 0x01, 0x0a, 0x0a, 0x4c, 0x37, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x36, 0x0a, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x4c, 0x37, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x1a,
	0x39, 0x0a, 0x0b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x04, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x69, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x30, 0x0a, 0x0a, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75,
	0x6d, 0x2e, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x74,
	0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x5f, 0x72, 0x75, 0x6c,
	0x65, 0x5f, 0x72, 0x65, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x69, 0x6c,
	0x69, 0x75, 0x6d, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x66, 0x12, 0x2c, 0x0a, 0x12, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x65,
	0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x36, 0x0a, 0x17, 0x64, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x15, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x49, 0x64,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x0a, 0x13, 0x64, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x64, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x2a, 0x0a, 0x04, 0x68, 0x74, 0x74, 0x70,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e,
	0x48, 0x74, 0x74, 0x70, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x04,
	0x68, 0x74, 0x74, 0x70, 0x12, 0x2d, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x65, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2e, 0x4b, 0x61, 0x66,
	0x6b, 0x61, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x05, 0x6b, 0x61,
	0x66, 0x6b, 0x61, 0x12, 0x33, 0x0a, 0x0a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x5f, 0x6c,
	0x37, 0x18, 0x66, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d,
	0x2e, 0x4c, 0x37, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x48, 0x00, 0x52, 0x09, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x4c, 0x37, 0x42, 0x04, 0x0a, 0x02, 0x6c, 0x37, 0x2a, 0x31,
	0x0a, 0x0c, 0x48, 0x74, 0x74, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x0a,
	0x0a, 0x06, 0x48, 0x54, 0x54, 0x50, 0x31, 0x30, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x48, 0x54,
	0x54, 0x50, 0x31, 0x31, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x54, 0x54, 0x50, 0x32, 0x10,
	0x02, 0x2a, 0x32, 0x0a, 0x09, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x65, 0x6e,
	0x69, 0x65, 0x64, 0x10, 0x02, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x78, 0x79,
	0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x69, 0x6c, 0x69, 0x75, 0x6d, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x63,
	0x69, 0x6c, 0x69, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cilium_api_accesslog_proto_rawDescOnce sync.Once
	file_cilium_api_accesslog_proto_rawDescData = file_cilium_api_accesslog_proto_rawDesc
)

func file_cilium_api_accesslog_proto_rawDescGZIP() []byte {
	file_cilium_api_accesslog_proto_rawDescOnce.Do(func() {
		file_cilium_api_accesslog_proto_rawDescData = protoimpl.X.CompressGZIP(file_cilium_api_accesslog_proto_rawDescData)
	})
	return file_cilium_api_accesslog_proto_rawDescData
}

var file_cilium_api_accesslog_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_cilium_api_accesslog_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_cilium_api_accesslog_proto_goTypes = []interface{}{
	(HttpProtocol)(0),     // 0: cilium.HttpProtocol
	(EntryType)(0),        // 1: cilium.EntryType
	(*KeyValue)(nil),      // 2: cilium.KeyValue
	(*HttpLogEntry)(nil),  // 3: cilium.HttpLogEntry
	(*KafkaLogEntry)(nil), // 4: cilium.KafkaLogEntry
	(*L7LogEntry)(nil),    // 5: cilium.L7LogEntry
	(*LogEntry)(nil),      // 6: cilium.LogEntry
	nil,                   // 7: cilium.L7LogEntry.FieldsEntry
}
var file_cilium_api_accesslog_proto_depIdxs = []int32{
	0, // 0: cilium.HttpLogEntry.http_protocol:type_name -> cilium.HttpProtocol
	2, // 1: cilium.HttpLogEntry.headers:type_name -> cilium.KeyValue
	2, // 2: cilium.HttpLogEntry.missing_headers:type_name -> cilium.KeyValue
	2, // 3: cilium.HttpLogEntry.rejected_headers:type_name -> cilium.KeyValue
	7, // 4: cilium.L7LogEntry.fields:type_name -> cilium.L7LogEntry.FieldsEntry
	1, // 5: cilium.LogEntry.entry_type:type_name -> cilium.EntryType
	3, // 6: cilium.LogEntry.http:type_name -> cilium.HttpLogEntry
	4, // 7: cilium.LogEntry.kafka:type_name -> cilium.KafkaLogEntry
	5, // 8: cilium.LogEntry.generic_l7:type_name -> cilium.L7LogEntry
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_cilium_api_accesslog_proto_init() }
func file_cilium_api_accesslog_proto_init() {
	if File_cilium_api_accesslog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cilium_api_accesslog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValue); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cilium_api_accesslog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HttpLogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cilium_api_accesslog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaLogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cilium_api_accesslog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*L7LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_cilium_api_accesslog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_cilium_api_accesslog_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*LogEntry_Http)(nil),
		(*LogEntry_Kafka)(nil),
		(*LogEntry_GenericL7)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_cilium_api_accesslog_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cilium_api_accesslog_proto_goTypes,
		DependencyIndexes: file_cilium_api_accesslog_proto_depIdxs,
		EnumInfos:         file_cilium_api_accesslog_proto_enumTypes,
		MessageInfos:      file_cilium_api_accesslog_proto_msgTypes,
	}.Build()
	File_cilium_api_accesslog_proto = out.File
	file_cilium_api_accesslog_proto_rawDesc = nil
	file_cilium_api_accesslog_proto_goTypes = nil
	file_cilium_api_accesslog_proto_depIdxs = nil
}