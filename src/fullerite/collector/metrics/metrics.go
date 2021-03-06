// Code generated by protoc-gen-go. DO NOT EDIT.
// source: metrics/metrics.proto

package metrics

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SampleType int32

const (
	SampleType_COUNTER   SampleType = 0
	SampleType_GAUGE     SampleType = 1
	SampleType_SUMMARY   SampleType = 2
	SampleType_HISTOGRAM SampleType = 3
	SampleType_UNTYPED   SampleType = 4
)

var SampleType_name = map[int32]string{
	0: "COUNTER",
	1: "GAUGE",
	2: "SUMMARY",
	3: "HISTOGRAM",
	4: "UNTYPED",
}

var SampleType_value = map[string]int32{
	"COUNTER":   0,
	"GAUGE":     1,
	"SUMMARY":   2,
	"HISTOGRAM": 3,
	"UNTYPED":   4,
}

func (x SampleType) String() string {
	return proto.EnumName(SampleType_name, int32(x))
}

func (SampleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{0}
}

type MetricFamilySamples struct {
	Name                 string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type                 SampleType `protobuf:"varint,2,opt,name=type,proto3,enum=com.yelp.generated.queryengine.metrics.SampleType" json:"type,omitempty"`
	Help                 string     `protobuf:"bytes,3,opt,name=help,proto3" json:"help,omitempty"`
	Samples              []*Sample  `protobuf:"bytes,4,rep,name=samples,proto3" json:"samples,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MetricFamilySamples) Reset()         { *m = MetricFamilySamples{} }
func (m *MetricFamilySamples) String() string { return proto.CompactTextString(m) }
func (*MetricFamilySamples) ProtoMessage()    {}
func (*MetricFamilySamples) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{0}
}

func (m *MetricFamilySamples) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricFamilySamples.Unmarshal(m, b)
}
func (m *MetricFamilySamples) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricFamilySamples.Marshal(b, m, deterministic)
}
func (m *MetricFamilySamples) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricFamilySamples.Merge(m, src)
}
func (m *MetricFamilySamples) XXX_Size() int {
	return xxx_messageInfo_MetricFamilySamples.Size(m)
}
func (m *MetricFamilySamples) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricFamilySamples.DiscardUnknown(m)
}

var xxx_messageInfo_MetricFamilySamples proto.InternalMessageInfo

func (m *MetricFamilySamples) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricFamilySamples) GetType() SampleType {
	if m != nil {
		return m.Type
	}
	return SampleType_COUNTER
}

func (m *MetricFamilySamples) GetHelp() string {
	if m != nil {
		return m.Help
	}
	return ""
}

func (m *MetricFamilySamples) GetSamples() []*Sample {
	if m != nil {
		return m.Samples
	}
	return nil
}

type Sample struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	LabelValues          []string `protobuf:"bytes,2,rep,name=labelValues,proto3" json:"labelValues,omitempty"`
	LabelNames           []string `protobuf:"bytes,3,rep,name=labelNames,proto3" json:"labelNames,omitempty"`
	Value                float64  `protobuf:"fixed64,4,opt,name=value,proto3" json:"value,omitempty"`
	TimestampMs          int64    `protobuf:"varint,5,opt,name=timestampMs,proto3" json:"timestampMs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Sample) Reset()         { *m = Sample{} }
func (m *Sample) String() string { return proto.CompactTextString(m) }
func (*Sample) ProtoMessage()    {}
func (*Sample) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{1}
}

func (m *Sample) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Sample.Unmarshal(m, b)
}
func (m *Sample) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Sample.Marshal(b, m, deterministic)
}
func (m *Sample) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sample.Merge(m, src)
}
func (m *Sample) XXX_Size() int {
	return xxx_messageInfo_Sample.Size(m)
}
func (m *Sample) XXX_DiscardUnknown() {
	xxx_messageInfo_Sample.DiscardUnknown(m)
}

var xxx_messageInfo_Sample proto.InternalMessageInfo

func (m *Sample) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Sample) GetLabelValues() []string {
	if m != nil {
		return m.LabelValues
	}
	return nil
}

func (m *Sample) GetLabelNames() []string {
	if m != nil {
		return m.LabelNames
	}
	return nil
}

func (m *Sample) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Sample) GetTimestampMs() int64 {
	if m != nil {
		return m.TimestampMs
	}
	return 0
}

type MetricsRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MetricsRequest) Reset()         { *m = MetricsRequest{} }
func (m *MetricsRequest) String() string { return proto.CompactTextString(m) }
func (*MetricsRequest) ProtoMessage()    {}
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{2}
}

func (m *MetricsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsRequest.Unmarshal(m, b)
}
func (m *MetricsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsRequest.Marshal(b, m, deterministic)
}
func (m *MetricsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsRequest.Merge(m, src)
}
func (m *MetricsRequest) XXX_Size() int {
	return xxx_messageInfo_MetricsRequest.Size(m)
}
func (m *MetricsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsRequest proto.InternalMessageInfo

type MetricsResponse struct {
	MetricFamilySample   []*MetricFamilySamples `protobuf:"bytes,1,rep,name=metricFamilySample,proto3" json:"metricFamilySample,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *MetricsResponse) Reset()         { *m = MetricsResponse{} }
func (m *MetricsResponse) String() string { return proto.CompactTextString(m) }
func (*MetricsResponse) ProtoMessage()    {}
func (*MetricsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cb1af212157fe6fb, []int{3}
}

func (m *MetricsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricsResponse.Unmarshal(m, b)
}
func (m *MetricsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricsResponse.Marshal(b, m, deterministic)
}
func (m *MetricsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricsResponse.Merge(m, src)
}
func (m *MetricsResponse) XXX_Size() int {
	return xxx_messageInfo_MetricsResponse.Size(m)
}
func (m *MetricsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MetricsResponse proto.InternalMessageInfo

func (m *MetricsResponse) GetMetricFamilySample() []*MetricFamilySamples {
	if m != nil {
		return m.MetricFamilySample
	}
	return nil
}

func init() {
	proto.RegisterEnum("com.yelp.generated.queryengine.metrics.SampleType", SampleType_name, SampleType_value)
	proto.RegisterType((*MetricFamilySamples)(nil), "com.yelp.generated.queryengine.metrics.MetricFamilySamples")
	proto.RegisterType((*Sample)(nil), "com.yelp.generated.queryengine.metrics.Sample")
	proto.RegisterType((*MetricsRequest)(nil), "com.yelp.generated.queryengine.metrics.MetricsRequest")
	proto.RegisterType((*MetricsResponse)(nil), "com.yelp.generated.queryengine.metrics.MetricsResponse")
}

func init() { proto.RegisterFile("metrics/metrics.proto", fileDescriptor_cb1af212157fe6fb) }

var fileDescriptor_cb1af212157fe6fb = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0x86, 0xd7, 0x4d, 0xba, 0x55, 0xa7, 0x62, 0x89, 0x0c, 0x48, 0x16, 0x07, 0x14, 0xe5, 0x80,
	0x22, 0x0e, 0xa9, 0x14, 0x10, 0x1c, 0x38, 0x15, 0xe8, 0x76, 0x39, 0xa4, 0x8b, 0xdc, 0x16, 0x69,
	0xb9, 0xac, 0xbc, 0x65, 0x28, 0xd1, 0xc6, 0x89, 0x37, 0x4e, 0x91, 0x22, 0x21, 0xae, 0xbc, 0x01,
	0xaf, 0xc5, 0x2b, 0x21, 0xdb, 0x2d, 0x54, 0xa2, 0x87, 0xf6, 0x94, 0xf1, 0x9f, 0xf9, 0x3f, 0xdb,
	0xff, 0xc8, 0xf0, 0x48, 0x62, 0x53, 0xe7, 0x4b, 0x3d, 0xdc, 0x7c, 0x13, 0x55, 0x57, 0x4d, 0x45,
	0x9f, 0x2e, 0x2b, 0x99, 0xb4, 0x58, 0xa8, 0x64, 0x85, 0x25, 0xd6, 0xa2, 0xc1, 0xcf, 0xc9, 0xdd,
	0x1a, 0xeb, 0x16, 0xcb, 0x55, 0x5e, 0x62, 0xb2, 0xe9, 0x8e, 0x7e, 0x13, 0x78, 0x90, 0xd9, 0xfa,
	0x5c, 0xc8, 0xbc, 0x68, 0x67, 0x42, 0xaa, 0x02, 0x35, 0xa5, 0xe0, 0x97, 0x42, 0x22, 0x23, 0x21,
	0x89, 0xfb, 0xdc, 0xd6, 0xf4, 0x1c, 0xfc, 0xa6, 0x55, 0xc8, 0x3a, 0x21, 0x89, 0xcf, 0xd2, 0x34,
	0x39, 0x6c, 0x8b, 0xc4, 0x21, 0xe7, 0xad, 0x42, 0x6e, 0xfd, 0x86, 0xfd, 0x15, 0x0b, 0xc5, 0x3c,
	0xc7, 0x36, 0x35, 0xbd, 0x80, 0x9e, 0x76, 0x5b, 0x33, 0x3f, 0xf4, 0xe2, 0x41, 0x9a, 0x1c, 0x87,
	0xe7, 0x5b, 0x7b, 0xf4, 0x8b, 0xc0, 0xa9, 0xd3, 0xf6, 0x5e, 0x22, 0x84, 0x41, 0x21, 0x6e, 0xb0,
	0xf8, 0x28, 0x8a, 0x35, 0x6a, 0xd6, 0x09, 0xbd, 0xb8, 0xcf, 0x77, 0x25, 0xfa, 0x04, 0xc0, 0x2e,
	0xa7, 0x42, 0xa2, 0x66, 0x9e, 0x6d, 0xd8, 0x51, 0xe8, 0x43, 0xe8, 0x7e, 0x33, 0x9d, 0xcc, 0x0f,
	0x49, 0x4c, 0xb8, 0x5b, 0x18, 0x6e, 0x93, 0x4b, 0xd4, 0x8d, 0x90, 0x2a, 0xd3, 0xac, 0x1b, 0x92,
	0xd8, 0xe3, 0xbb, 0x52, 0x14, 0xc0, 0x99, 0x4b, 0x5a, 0x73, 0xbc, 0x5b, 0xa3, 0x6e, 0xa2, 0x1f,
	0x70, 0xff, 0xaf, 0xa2, 0x55, 0x55, 0x6a, 0xa4, 0xb7, 0x40, 0xe5, 0x7f, 0xe3, 0x60, 0xc4, 0x46,
	0xf2, 0xfa, 0xd0, 0x48, 0xf6, 0x0c, 0x94, 0xef, 0xc1, 0x3e, 0xcb, 0x00, 0xfe, 0x0d, 0x87, 0x0e,
	0xa0, 0xf7, 0xf6, 0x72, 0x31, 0x9d, 0x8f, 0x79, 0x70, 0x42, 0xfb, 0xd0, 0x9d, 0x8c, 0x16, 0x93,
	0x71, 0x40, 0x8c, 0x3e, 0x5b, 0x64, 0xd9, 0x88, 0x5f, 0x05, 0x1d, 0x7a, 0x0f, 0xfa, 0x17, 0xef,
	0x67, 0xf3, 0xcb, 0x09, 0x1f, 0x65, 0x81, 0x67, 0xfe, 0x2d, 0xa6, 0xf3, 0xab, 0x0f, 0xe3, 0x77,
	0x81, 0x9f, 0xfe, 0x24, 0xd0, 0xdb, 0xdc, 0x87, 0x7e, 0x87, 0xde, 0xe6, 0x34, 0xf4, 0xe5, 0x71,
	0xc7, 0xde, 0xa6, 0xf3, 0xf8, 0xd5, 0xd1, 0x3e, 0x97, 0x61, 0x74, 0xf2, 0xe6, 0xc5, 0xa7, 0xd4,
	0xf8, 0x96, 0x55, 0xad, 0x0c, 0x64, 0x58, 0x8b, 0xf2, 0x36, 0x2f, 0x57, 0xd7, 0xaa, 0x10, 0xcd,
	0x97, 0xaa, 0x96, 0x43, 0x8b, 0xb9, 0x76, 0x9c, 0xed, 0xcb, 0xb9, 0x39, 0xb5, 0x4f, 0xe7, 0xf9,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9d, 0x07, 0x0a, 0xa4, 0x53, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MetricsClient is the client API for Metrics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsClient interface {
	Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error)
}

type metricsClient struct {
	cc *grpc.ClientConn
}

func NewMetricsClient(cc *grpc.ClientConn) MetricsClient {
	return &metricsClient{cc}
}

func (c *metricsClient) Metrics(ctx context.Context, in *MetricsRequest, opts ...grpc.CallOption) (*MetricsResponse, error) {
	out := new(MetricsResponse)
	err := c.cc.Invoke(ctx, "/com.yelp.generated.queryengine.metrics.Metrics/metrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServer is the server API for Metrics service.
type MetricsServer interface {
	Metrics(context.Context, *MetricsRequest) (*MetricsResponse, error)
}

// UnimplementedMetricsServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServer struct {
}

func (*UnimplementedMetricsServer) Metrics(ctx context.Context, req *MetricsRequest) (*MetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Metrics not implemented")
}

func RegisterMetricsServer(s *grpc.Server, srv MetricsServer) {
	s.RegisterService(&_Metrics_serviceDesc, srv)
}

func _Metrics_Metrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServer).Metrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.yelp.generated.queryengine.metrics.Metrics/Metrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServer).Metrics(ctx, req.(*MetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Metrics_serviceDesc = grpc.ServiceDesc{
	ServiceName: "com.yelp.generated.queryengine.metrics.Metrics",
	HandlerType: (*MetricsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "metrics",
			Handler:    _Metrics_Metrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metrics/metrics.proto",
}
