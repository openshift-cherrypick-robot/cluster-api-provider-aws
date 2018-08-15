// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/monitoring.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Monitoring configuration of the service.
//
// The example below shows how to configure monitored resources and metrics
// for monitoring. In the example, a monitored resource and two metrics are
// defined. The `library.googleapis.com/book/returned_count` metric is sent
// to both producer and consumer projects, whereas the
// `library.googleapis.com/book/overdue_count` metric is only sent to the
// consumer project.
//
//     monitored_resources:
//     - type: library.googleapis.com/branch
//       labels:
//       - key: /city
//         description: The city where the library branch is located in.
//       - key: /name
//         description: The name of the branch.
//     metrics:
//     - name: library.googleapis.com/book/returned_count
//       metric_kind: DELTA
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     - name: library.googleapis.com/book/overdue_count
//       metric_kind: GAUGE
//       value_type: INT64
//       labels:
//       - key: /customer_id
//     monitoring:
//       producer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//       consumer_destinations:
//       - monitored_resource: library.googleapis.com/branch
//         metrics:
//         - library.googleapis.com/book/returned_count
//         - library.googleapis.com/book/overdue_count
type Monitoring struct {
	// Monitoring configurations for sending metrics to the producer project.
	// There can be multiple producer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one producer destination.
	ProducerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,1,rep,name=producer_destinations,json=producerDestinations,proto3" json:"producer_destinations,omitempty"`
	// Monitoring configurations for sending metrics to the consumer project.
	// There can be multiple consumer destinations, each one must have a
	// different monitored resource type. A metric can be used in at most
	// one consumer destination.
	ConsumerDestinations []*Monitoring_MonitoringDestination `protobuf:"bytes,2,rep,name=consumer_destinations,json=consumerDestinations,proto3" json:"consumer_destinations,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Monitoring) Reset()         { *m = Monitoring{} }
func (m *Monitoring) String() string { return proto.CompactTextString(m) }
func (*Monitoring) ProtoMessage()    {}
func (*Monitoring) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitoring_4e73ce5561a2dc9c, []int{0}
}
func (m *Monitoring) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Monitoring.Unmarshal(m, b)
}
func (m *Monitoring) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Monitoring.Marshal(b, m, deterministic)
}
func (dst *Monitoring) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Monitoring.Merge(dst, src)
}
func (m *Monitoring) XXX_Size() int {
	return xxx_messageInfo_Monitoring.Size(m)
}
func (m *Monitoring) XXX_DiscardUnknown() {
	xxx_messageInfo_Monitoring.DiscardUnknown(m)
}

var xxx_messageInfo_Monitoring proto.InternalMessageInfo

func (m *Monitoring) GetProducerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ProducerDestinations
	}
	return nil
}

func (m *Monitoring) GetConsumerDestinations() []*Monitoring_MonitoringDestination {
	if m != nil {
		return m.ConsumerDestinations
	}
	return nil
}

// Configuration of a specific monitoring destination (the producer project
// or the consumer project).
type Monitoring_MonitoringDestination struct {
	// The monitored resource type. The type must be defined in
	// [Service.monitored_resources][google.api.Service.monitored_resources] section.
	MonitoredResource string `protobuf:"bytes,1,opt,name=monitored_resource,json=monitoredResource,proto3" json:"monitored_resource,omitempty"`
	// Names of the metrics to report to this monitoring destination.
	// Each name must be defined in [Service.metrics][google.api.Service.metrics] section.
	Metrics              []string `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Monitoring_MonitoringDestination) Reset()         { *m = Monitoring_MonitoringDestination{} }
func (m *Monitoring_MonitoringDestination) String() string { return proto.CompactTextString(m) }
func (*Monitoring_MonitoringDestination) ProtoMessage()    {}
func (*Monitoring_MonitoringDestination) Descriptor() ([]byte, []int) {
	return fileDescriptor_monitoring_4e73ce5561a2dc9c, []int{0, 0}
}
func (m *Monitoring_MonitoringDestination) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Monitoring_MonitoringDestination.Unmarshal(m, b)
}
func (m *Monitoring_MonitoringDestination) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Monitoring_MonitoringDestination.Marshal(b, m, deterministic)
}
func (dst *Monitoring_MonitoringDestination) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Monitoring_MonitoringDestination.Merge(dst, src)
}
func (m *Monitoring_MonitoringDestination) XXX_Size() int {
	return xxx_messageInfo_Monitoring_MonitoringDestination.Size(m)
}
func (m *Monitoring_MonitoringDestination) XXX_DiscardUnknown() {
	xxx_messageInfo_Monitoring_MonitoringDestination.DiscardUnknown(m)
}

var xxx_messageInfo_Monitoring_MonitoringDestination proto.InternalMessageInfo

func (m *Monitoring_MonitoringDestination) GetMonitoredResource() string {
	if m != nil {
		return m.MonitoredResource
	}
	return ""
}

func (m *Monitoring_MonitoringDestination) GetMetrics() []string {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func init() {
	proto.RegisterType((*Monitoring)(nil), "google.api.Monitoring")
	proto.RegisterType((*Monitoring_MonitoringDestination)(nil), "google.api.Monitoring.MonitoringDestination")
}

func init() {
	proto.RegisterFile("google/api/monitoring.proto", fileDescriptor_monitoring_4e73ce5561a2dc9c)
}

var fileDescriptor_monitoring_4e73ce5561a2dc9c = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x51, 0xcd, 0x4a, 0xc4, 0x30,
	0x10, 0xa6, 0x55, 0x94, 0x8d, 0xa0, 0x58, 0x5c, 0x28, 0xab, 0x87, 0xc5, 0xd3, 0x1e, 0xb4, 0x05,
	0x3d, 0x7a, 0x72, 0x51, 0xc4, 0x83, 0x50, 0x7a, 0xf4, 0xb2, 0xc6, 0x74, 0x0c, 0x03, 0xdb, 0x99,
	0x9a, 0xa4, 0x3e, 0x90, 0xcf, 0xe0, 0x03, 0xca, 0x36, 0xed, 0x36, 0x8a, 0x27, 0x6f, 0x99, 0x7c,
	0x7f, 0xc3, 0x37, 0xe2, 0x54, 0x33, 0xeb, 0x35, 0xe4, 0xb2, 0xc1, 0xbc, 0x66, 0x42, 0xc7, 0x06,
	0x49, 0x67, 0x8d, 0x61, 0xc7, 0x89, 0xf0, 0x60, 0x26, 0x1b, 0x9c, 0x9d, 0x05, 0x44, 0x49, 0xc4,
	0x4e, 0x3a, 0x64, 0xb2, 0x9e, 0x79, 0xfe, 0x15, 0x0b, 0xf1, 0xb4, 0x95, 0x27, 0x52, 0x4c, 0x1b,
	0xc3, 0x55, 0xab, 0xc0, 0xac, 0x2a, 0xb0, 0x0e, 0xc9, 0xb3, 0xd3, 0x68, 0xbe, 0xb3, 0x38, 0xb8,
	0xba, 0xc8, 0x46, 0xe3, 0x6c, 0x94, 0x05, 0xcf, 0xbb, 0x51, 0x54, 0x9e, 0x0c, 0x56, 0xc1, 0xa7,
	0xdd, 0x44, 0x28, 0x26, 0xdb, 0xd6, 0xbf, 0x23, 0xe2, 0xff, 0x44, 0x0c, 0x56, 0x61, 0xc4, 0xec,
	0x45, 0x4c, 0xff, 0xa4, 0x27, 0x97, 0x22, 0xe9, 0xbb, 0x82, 0x6a, 0x65, 0xc0, 0x72, 0x6b, 0x14,
	0xa4, 0xd1, 0x3c, 0x5a, 0x4c, 0xca, 0xe3, 0x2d, 0x52, 0xf6, 0x40, 0x92, 0x8a, 0xfd, 0x1a, 0x9c,
	0x41, 0xe5, 0x97, 0x9b, 0x94, 0xc3, 0xb8, 0x7c, 0x17, 0x87, 0x8a, 0xeb, 0x60, 0xd5, 0xe5, 0xd1,
	0x98, 0x58, 0x6c, 0x9a, 0x2d, 0xa2, 0xe7, 0xfb, 0x1e, 0xd6, 0xbc, 0x96, 0xa4, 0x33, 0x36, 0x3a,
	0xd7, 0x40, 0x5d, 0xef, 0xb9, 0x87, 0x64, 0x83, 0xb6, 0x3b, 0x8c, 0x05, 0xf3, 0x81, 0x0a, 0x14,
	0xd3, 0x1b, 0xea, 0x9b, 0x1f, 0xd3, 0x67, 0xbc, 0xfb, 0x70, 0x5b, 0x3c, 0xbe, 0xee, 0x75, 0xc2,
	0xeb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x51, 0x35, 0xf3, 0xe2, 0xf9, 0x01, 0x00, 0x00,
}
