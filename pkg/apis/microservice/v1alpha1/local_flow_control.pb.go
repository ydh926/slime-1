// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/apis/microservice/v1alpha1/local_flow_control.proto

package v1alpha1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type UnitType int32

const (
	UnitType_UNKNOWN UnitType = 0
	UnitType_SECOND  UnitType = 1
	UnitType_MINUTE  UnitType = 2
	UnitType_HOUR    UnitType = 3
	UnitType_DAY     UnitType = 4
)

var UnitType_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var UnitType_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x UnitType) String() string {
	return proto.EnumName(UnitType_name, int32(x))
}

func (UnitType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{0}
}

//限流算法
type FlowControlConfStatus_LimitAlgorithm int32

const (
	FlowControlConfStatus_TokenBucket   FlowControlConfStatus_LimitAlgorithm = 0
	FlowControlConfStatus_SlidingWindow FlowControlConfStatus_LimitAlgorithm = 1
)

var FlowControlConfStatus_LimitAlgorithm_name = map[int32]string{
	0: "TokenBucket",
	1: "SlidingWindow",
}

var FlowControlConfStatus_LimitAlgorithm_value = map[string]int32{
	"TokenBucket":   0,
	"SlidingWindow": 1,
}

func (x FlowControlConfStatus_LimitAlgorithm) String() string {
	return proto.EnumName(FlowControlConfStatus_LimitAlgorithm_name, int32(x))
}

func (FlowControlConfStatus_LimitAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{3, 0}
}

type FlowControlConfSpec_LimitAlgorithm int32

const (
	FlowControlConfSpec_TokenBucket   FlowControlConfSpec_LimitAlgorithm = 0
	FlowControlConfSpec_SlidingWindow FlowControlConfSpec_LimitAlgorithm = 1
)

var FlowControlConfSpec_LimitAlgorithm_name = map[int32]string{
	0: "TokenBucket",
	1: "SlidingWindow",
}

var FlowControlConfSpec_LimitAlgorithm_value = map[string]int32{
	"TokenBucket":   0,
	"SlidingWindow": 1,
}

func (x FlowControlConfSpec_LimitAlgorithm) String() string {
	return proto.EnumName(FlowControlConfSpec_LimitAlgorithm_name, int32(x))
}

func (FlowControlConfSpec_LimitAlgorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{4, 0}
}

type RateConfig struct {
	Unit                 UnitType `protobuf:"varint,1,opt,name=unit,proto3,enum=envoy.config.filter.http.local_flow_control.v2.UnitType" json:"unit,omitempty"`
	RequestsPerUnit      uint32   `protobuf:"varint,2,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RateConfig) Reset()         { *m = RateConfig{} }
func (m *RateConfig) String() string { return proto.CompactTextString(m) }
func (*RateConfig) ProtoMessage()    {}
func (*RateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{0}
}

func (m *RateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateConfig.Unmarshal(m, b)
}
func (m *RateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateConfig.Marshal(b, m, deterministic)
}
func (m *RateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateConfig.Merge(m, src)
}
func (m *RateConfig) XXX_Size() int {
	return xxx_messageInfo_RateConfig.Size(m)
}
func (m *RateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_RateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_RateConfig proto.InternalMessageInfo

func (m *RateConfig) GetUnit() UnitType {
	if m != nil {
		return m.Unit
	}
	return UnitType_UNKNOWN
}

func (m *RateConfig) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

type RateLimitDescriptorConfigStatus struct {
	Key                  string                             `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string                             `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	RateLimit            *RateConfig                        `protobuf:"bytes,3,opt,name=rate_limit,json=rateLimit,proto3" json:"rate_limit,omitempty"`
	Descriptors          []*RateLimitDescriptorConfigStatus `protobuf:"bytes,4,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *RateLimitDescriptorConfigStatus) Reset()         { *m = RateLimitDescriptorConfigStatus{} }
func (m *RateLimitDescriptorConfigStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptorConfigStatus) ProtoMessage()    {}
func (*RateLimitDescriptorConfigStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{1}
}

func (m *RateLimitDescriptorConfigStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptorConfigStatus.Unmarshal(m, b)
}
func (m *RateLimitDescriptorConfigStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptorConfigStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptorConfigStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptorConfigStatus.Merge(m, src)
}
func (m *RateLimitDescriptorConfigStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptorConfigStatus.Size(m)
}
func (m *RateLimitDescriptorConfigStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptorConfigStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptorConfigStatus proto.InternalMessageInfo

func (m *RateLimitDescriptorConfigStatus) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptorConfigStatus) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RateLimitDescriptorConfigStatus) GetRateLimit() *RateConfig {
	if m != nil {
		return m.RateLimit
	}
	return nil
}

func (m *RateLimitDescriptorConfigStatus) GetDescriptors() []*RateLimitDescriptorConfigStatus {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

type RateLimitConfStatus struct {
	Domain               string                             `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*RateLimitDescriptorConfigStatus `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *RateLimitConfStatus) Reset()         { *m = RateLimitConfStatus{} }
func (m *RateLimitConfStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfStatus) ProtoMessage()    {}
func (*RateLimitConfStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{2}
}

func (m *RateLimitConfStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfStatus.Unmarshal(m, b)
}
func (m *RateLimitConfStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitConfStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfStatus.Merge(m, src)
}
func (m *RateLimitConfStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfStatus.Size(m)
}
func (m *RateLimitConfStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfStatus proto.InternalMessageInfo

func (m *RateLimitConfStatus) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfStatus) GetDescriptors() []*RateLimitDescriptorConfigStatus {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

type FlowControlConfStatus struct {
	RateLimitConf *RateLimitConfStatus                 `protobuf:"bytes,1,opt,name=rate_limit_conf,json=rateLimitConf,proto3" json:"rate_limit_conf,omitempty"`
	Algorithm     FlowControlConfStatus_LimitAlgorithm `protobuf:"varint,2,opt,name=algorithm,proto3,enum=envoy.config.filter.http.local_flow_control.v2.FlowControlConfStatus_LimitAlgorithm" json:"algorithm,omitempty"`
	//计数器可使用最大内存限制,单位：Mbytes
	MaxMemorySize        uint64   `protobuf:"varint,3,opt,name=max_memory_size,json=maxMemorySize,proto3" json:"max_memory_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowControlConfStatus) Reset()         { *m = FlowControlConfStatus{} }
func (m *FlowControlConfStatus) String() string { return proto.CompactTextString(m) }
func (*FlowControlConfStatus) ProtoMessage()    {}
func (*FlowControlConfStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{3}
}

func (m *FlowControlConfStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowControlConfStatus.Unmarshal(m, b)
}
func (m *FlowControlConfStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowControlConfStatus.Marshal(b, m, deterministic)
}
func (m *FlowControlConfStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowControlConfStatus.Merge(m, src)
}
func (m *FlowControlConfStatus) XXX_Size() int {
	return xxx_messageInfo_FlowControlConfStatus.Size(m)
}
func (m *FlowControlConfStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowControlConfStatus.DiscardUnknown(m)
}

var xxx_messageInfo_FlowControlConfStatus proto.InternalMessageInfo

func (m *FlowControlConfStatus) GetRateLimitConf() *RateLimitConfStatus {
	if m != nil {
		return m.RateLimitConf
	}
	return nil
}

func (m *FlowControlConfStatus) GetAlgorithm() FlowControlConfStatus_LimitAlgorithm {
	if m != nil {
		return m.Algorithm
	}
	return FlowControlConfStatus_TokenBucket
}

func (m *FlowControlConfStatus) GetMaxMemorySize() uint64 {
	if m != nil {
		return m.MaxMemorySize
	}
	return 0
}

type FlowControlConfSpec struct {
	RateLimitConf *RateLimitConfSpec                 `protobuf:"bytes,1,opt,name=rate_limit_conf,json=rateLimitConf,proto3" json:"rate_limit_conf,omitempty"`
	Algorithm     FlowControlConfSpec_LimitAlgorithm `protobuf:"varint,2,opt,name=algorithm,proto3,enum=envoy.config.filter.http.local_flow_control.v2.FlowControlConfSpec_LimitAlgorithm" json:"algorithm,omitempty"`
	//计数器可使用最大内存限制,单位：Mbytes
	MaxMemorySize        uint64   `protobuf:"varint,3,opt,name=max_memory_size,json=maxMemorySize,proto3" json:"max_memory_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FlowControlConfSpec) Reset()         { *m = FlowControlConfSpec{} }
func (m *FlowControlConfSpec) String() string { return proto.CompactTextString(m) }
func (*FlowControlConfSpec) ProtoMessage()    {}
func (*FlowControlConfSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{4}
}

func (m *FlowControlConfSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FlowControlConfSpec.Unmarshal(m, b)
}
func (m *FlowControlConfSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FlowControlConfSpec.Marshal(b, m, deterministic)
}
func (m *FlowControlConfSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FlowControlConfSpec.Merge(m, src)
}
func (m *FlowControlConfSpec) XXX_Size() int {
	return xxx_messageInfo_FlowControlConfSpec.Size(m)
}
func (m *FlowControlConfSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_FlowControlConfSpec.DiscardUnknown(m)
}

var xxx_messageInfo_FlowControlConfSpec proto.InternalMessageInfo

func (m *FlowControlConfSpec) GetRateLimitConf() *RateLimitConfSpec {
	if m != nil {
		return m.RateLimitConf
	}
	return nil
}

func (m *FlowControlConfSpec) GetAlgorithm() FlowControlConfSpec_LimitAlgorithm {
	if m != nil {
		return m.Algorithm
	}
	return FlowControlConfSpec_TokenBucket
}

func (m *FlowControlConfSpec) GetMaxMemorySize() uint64 {
	if m != nil {
		return m.MaxMemorySize
	}
	return 0
}

type RateLimitDescriptorConfigSpec struct {
	Key                  string                           `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string                           `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Descriptors          []*RateLimitDescriptorConfigSpec `protobuf:"bytes,4,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	When                 string                           `protobuf:"bytes,6,opt,name=when,proto3" json:"when,omitempty"`
	Then                 string                           `protobuf:"bytes,7,opt,name=then,proto3" json:"then,omitempty"`
	Unit                 UnitType                         `protobuf:"varint,8,opt,name=unit,proto3,enum=envoy.config.filter.http.local_flow_control.v2.UnitType" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitDescriptorConfigSpec) Reset()         { *m = RateLimitDescriptorConfigSpec{} }
func (m *RateLimitDescriptorConfigSpec) String() string { return proto.CompactTextString(m) }
func (*RateLimitDescriptorConfigSpec) ProtoMessage()    {}
func (*RateLimitDescriptorConfigSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{5}
}

func (m *RateLimitDescriptorConfigSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitDescriptorConfigSpec.Unmarshal(m, b)
}
func (m *RateLimitDescriptorConfigSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitDescriptorConfigSpec.Marshal(b, m, deterministic)
}
func (m *RateLimitDescriptorConfigSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitDescriptorConfigSpec.Merge(m, src)
}
func (m *RateLimitDescriptorConfigSpec) XXX_Size() int {
	return xxx_messageInfo_RateLimitDescriptorConfigSpec.Size(m)
}
func (m *RateLimitDescriptorConfigSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitDescriptorConfigSpec.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitDescriptorConfigSpec proto.InternalMessageInfo

func (m *RateLimitDescriptorConfigSpec) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *RateLimitDescriptorConfigSpec) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *RateLimitDescriptorConfigSpec) GetDescriptors() []*RateLimitDescriptorConfigSpec {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitDescriptorConfigSpec) GetWhen() string {
	if m != nil {
		return m.When
	}
	return ""
}

func (m *RateLimitDescriptorConfigSpec) GetThen() string {
	if m != nil {
		return m.Then
	}
	return ""
}

func (m *RateLimitDescriptorConfigSpec) GetUnit() UnitType {
	if m != nil {
		return m.Unit
	}
	return UnitType_UNKNOWN
}

type RateLimitConfSpec struct {
	Domain               string                           `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*RateLimitDescriptorConfigSpec `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitConfSpec) Reset()         { *m = RateLimitConfSpec{} }
func (m *RateLimitConfSpec) String() string { return proto.CompactTextString(m) }
func (*RateLimitConfSpec) ProtoMessage()    {}
func (*RateLimitConfSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_674feea6ee15854a, []int{6}
}

func (m *RateLimitConfSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitConfSpec.Unmarshal(m, b)
}
func (m *RateLimitConfSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitConfSpec.Marshal(b, m, deterministic)
}
func (m *RateLimitConfSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitConfSpec.Merge(m, src)
}
func (m *RateLimitConfSpec) XXX_Size() int {
	return xxx_messageInfo_RateLimitConfSpec.Size(m)
}
func (m *RateLimitConfSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitConfSpec.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitConfSpec proto.InternalMessageInfo

func (m *RateLimitConfSpec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitConfSpec) GetDescriptors() []*RateLimitDescriptorConfigSpec {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func init() {
	proto.RegisterEnum("envoy.config.filter.http.local_flow_control.v2.UnitType", UnitType_name, UnitType_value)
	proto.RegisterEnum("envoy.config.filter.http.local_flow_control.v2.FlowControlConfStatus_LimitAlgorithm", FlowControlConfStatus_LimitAlgorithm_name, FlowControlConfStatus_LimitAlgorithm_value)
	proto.RegisterEnum("envoy.config.filter.http.local_flow_control.v2.FlowControlConfSpec_LimitAlgorithm", FlowControlConfSpec_LimitAlgorithm_name, FlowControlConfSpec_LimitAlgorithm_value)
	proto.RegisterType((*RateConfig)(nil), "envoy.config.filter.http.local_flow_control.v2.RateConfig")
	proto.RegisterType((*RateLimitDescriptorConfigStatus)(nil), "envoy.config.filter.http.local_flow_control.v2.RateLimitDescriptorConfigStatus")
	proto.RegisterType((*RateLimitConfStatus)(nil), "envoy.config.filter.http.local_flow_control.v2.RateLimitConfStatus")
	proto.RegisterType((*FlowControlConfStatus)(nil), "envoy.config.filter.http.local_flow_control.v2.FlowControlConfStatus")
	proto.RegisterType((*FlowControlConfSpec)(nil), "envoy.config.filter.http.local_flow_control.v2.FlowControlConfSpec")
	proto.RegisterType((*RateLimitDescriptorConfigSpec)(nil), "envoy.config.filter.http.local_flow_control.v2.RateLimitDescriptorConfigSpec")
	proto.RegisterType((*RateLimitConfSpec)(nil), "envoy.config.filter.http.local_flow_control.v2.RateLimitConfSpec")
}

func init() {
	proto.RegisterFile("pkg/apis/microservice/v1alpha1/local_flow_control.proto", fileDescriptor_674feea6ee15854a)
}

var fileDescriptor_674feea6ee15854a = []byte{
	// 656 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xc4, 0x55, 0x41, 0x6f, 0x12, 0x41,
	0x14, 0x76, 0x17, 0xa4, 0xe5, 0x11, 0xca, 0x76, 0x6a, 0x0d, 0x17, 0x63, 0xc3, 0xc1, 0x98, 0x1e,
	0x96, 0x14, 0x9b, 0x68, 0x8c, 0x97, 0x42, 0x6b, 0x34, 0xb6, 0xd0, 0x0c, 0x90, 0xa6, 0x5e, 0x36,
	0xeb, 0x32, 0xc0, 0x84, 0xdd, 0x9d, 0xed, 0xec, 0x40, 0x8b, 0x3f, 0xc0, 0x9b, 0x67, 0x2f, 0x1e,
	0xfc, 0x05, 0xfe, 0x16, 0x4f, 0xfe, 0x1e, 0x67, 0x86, 0x52, 0xb0, 0x85, 0x46, 0xda, 0xaa, 0xb7,
	0xb7, 0xdf, 0xee, 0xfb, 0xde, 0x37, 0xdf, 0xbc, 0x7d, 0x0f, 0x9e, 0x47, 0xbd, 0x4e, 0xd1, 0x8d,
	0x68, 0x5c, 0x0c, 0xa8, 0xc7, 0x59, 0x4c, 0xf8, 0x80, 0x7a, 0xa4, 0x38, 0xd8, 0x72, 0xfd, 0xa8,
	0xeb, 0x6e, 0x15, 0x7d, 0xe6, 0xb9, 0xbe, 0xd3, 0xf6, 0xd9, 0xa9, 0xe3, 0xb1, 0x50, 0x70, 0xe6,
	0xdb, 0x11, 0x67, 0x82, 0x21, 0x9b, 0x84, 0x03, 0x36, 0xb4, 0x25, 0xd8, 0xa6, 0x1d, 0xbb, 0x4d,
	0x7d, 0x41, 0xb8, 0xdd, 0x15, 0x22, 0xb2, 0x67, 0xa4, 0x0c, 0x4a, 0x85, 0x4f, 0x06, 0x00, 0x76,
	0x05, 0xa9, 0xe8, 0x04, 0xb4, 0x0f, 0xc9, 0x7e, 0x48, 0x45, 0xde, 0xd8, 0x30, 0x9e, 0xae, 0x94,
	0x5e, 0x2c, 0xc8, 0x66, 0x37, 0x65, 0x6e, 0x63, 0x18, 0x11, 0xac, 0x59, 0xd0, 0x26, 0xac, 0x72,
	0x72, 0xd2, 0x27, 0xb1, 0x88, 0x9d, 0x88, 0x70, 0x47, 0x53, 0x9b, 0x92, 0x3a, 0x8b, 0x73, 0xe3,
	0x17, 0x87, 0x84, 0xab, 0xac, 0xc2, 0x17, 0x13, 0x1e, 0x2b, 0x21, 0xfb, 0x34, 0xa0, 0x62, 0x97,
	0xc4, 0x1e, 0xa7, 0x91, 0x60, 0x7c, 0xa4, 0xab, 0x2e, 0x5c, 0xd1, 0x8f, 0x91, 0x05, 0x89, 0x1e,
	0x19, 0x6a, 0x71, 0x69, 0xac, 0x42, 0xf4, 0x00, 0xee, 0x0f, 0x5c, 0xbf, 0x4f, 0x34, 0x6b, 0x1a,
	0x8f, 0x1e, 0xd0, 0x31, 0x00, 0x97, 0x54, 0x8e, 0xaf, 0xb8, 0xf2, 0x09, 0xf9, 0x2a, 0x53, 0x7a,
	0xb9, 0xe8, 0x59, 0x26, 0xae, 0xe0, 0x34, 0x1f, 0x0b, 0x43, 0x27, 0x90, 0x69, 0x5d, 0x88, 0x8b,
	0xf3, 0xc9, 0x8d, 0x84, 0xe4, 0xae, 0xdd, 0x84, 0xfb, 0x9a, 0x83, 0xe2, 0xe9, 0x1a, 0x85, 0x6f,
	0x06, 0xac, 0x5d, 0x24, 0xa8, 0xcf, 0xce, 0xdd, 0x78, 0x08, 0xa9, 0x16, 0x0b, 0x5c, 0x1a, 0x9e,
	0x1b, 0x72, 0xfe, 0x74, 0x59, 0xa2, 0xf9, 0x0f, 0x24, 0xfe, 0x34, 0x61, 0xfd, 0xb5, 0x24, 0xa8,
	0x8c, 0xf2, 0xa7, 0x44, 0xf6, 0x20, 0x37, 0xb9, 0x0a, 0xc5, 0xdf, 0xd6, 0x6a, 0x33, 0xa5, 0xca,
	0x8d, 0x05, 0x4d, 0xd8, 0x71, 0x96, 0x4f, 0x83, 0x88, 0x43, 0xda, 0xf5, 0x3b, 0x8c, 0x53, 0xd1,
	0x0d, 0x74, 0x47, 0xac, 0x94, 0x1a, 0x8b, 0x96, 0x99, 0x79, 0x0c, 0x5b, 0xd7, 0xd8, 0x19, 0x73,
	0xe3, 0x49, 0x19, 0xf4, 0x04, 0x72, 0x81, 0x7b, 0xe6, 0x04, 0x24, 0x60, 0x7c, 0xe8, 0xc4, 0xf4,
	0x23, 0xd1, 0x0d, 0x97, 0xc4, 0x59, 0x09, 0x1f, 0x68, 0xb4, 0x2e, 0xc1, 0xc2, 0x36, 0xac, 0xfc,
	0x4e, 0x82, 0x72, 0x90, 0x69, 0xb0, 0x1e, 0x09, 0xcb, 0x7d, 0xaf, 0x47, 0x84, 0x75, 0x0f, 0xad,
	0x42, 0xb6, 0xee, 0xd3, 0x16, 0x0d, 0x3b, 0x47, 0x34, 0x6c, 0xb1, 0x53, 0xcb, 0x28, 0xfc, 0x30,
	0x61, 0xed, 0xb2, 0xa2, 0x88, 0x78, 0x88, 0xce, 0xb3, 0x75, 0xe7, 0x76, 0xb6, 0x4a, 0xee, 0xcb,
	0xa6, 0x46, 0x57, 0x4d, 0xc5, 0xb7, 0x35, 0x55, 0x96, 0xf9, 0x6f, 0x96, 0x7e, 0x37, 0xe1, 0xd1,
	0xfc, 0xe6, 0x56, 0xe6, 0xfe, 0xe9, 0x98, 0x61, 0xb3, 0x66, 0xc1, 0xc1, 0xdd, 0xfd, 0x68, 0xea,
	0x32, 0xa6, 0x2b, 0x20, 0x04, 0xc9, 0xd3, 0x2e, 0x09, 0xf3, 0x29, 0xad, 0x42, 0xc7, 0x0a, 0x13,
	0x0a, 0x5b, 0x1a, 0x61, 0x2a, 0xbe, 0x98, 0xe2, 0xcb, 0x77, 0x31, 0xc5, 0x0b, 0x5f, 0x0d, 0x58,
	0xbd, 0xd2, 0x25, 0x73, 0xa7, 0x0f, 0x9b, 0x35, 0x7d, 0xfe, 0xa2, 0x29, 0x9b, 0x65, 0x58, 0x1e,
	0x0b, 0x46, 0x19, 0x58, 0x6a, 0x56, 0xdf, 0x55, 0x6b, 0x47, 0x55, 0x79, 0xf7, 0x00, 0xa9, 0xfa,
	0x5e, 0xa5, 0x56, 0xdd, 0xb5, 0x0c, 0x15, 0x1f, 0xbc, 0xad, 0x36, 0x1b, 0x7b, 0x96, 0x89, 0x96,
	0x21, 0xf9, 0xa6, 0xd6, 0xc4, 0x56, 0x02, 0x2d, 0x41, 0x62, 0x77, 0xe7, 0xd8, 0x4a, 0x96, 0x3f,
	0x1b, 0xf0, 0x8a, 0xb2, 0x91, 0x48, 0xb9, 0x47, 0xcf, 0x86, 0x0b, 0xea, 0x2d, 0xaf, 0xef, 0x2b,
	0x78, 0xaa, 0xcd, 0x0f, 0xd5, 0x36, 0x3e, 0x34, 0xde, 0x6f, 0x0f, 0xfb, 0xa1, 0x1d, 0x12, 0x41,
	0xdc, 0x98, 0x48, 0xba, 0xa0, 0x18, 0xcb, 0x5f, 0x96, 0x14, 0xaf, 0x5f, 0xef, 0x1f, 0x52, 0x7a,
	0x99, 0x3f, 0xfb, 0x15, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x97, 0xf6, 0x24, 0x07, 0x08, 0x00, 0x00,
}
