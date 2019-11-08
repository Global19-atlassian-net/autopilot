// Code generated by protoc-gen-go. DO NOT EDIT.
// source: autopilot-operator.proto

package v1

import (
	fmt "fmt"
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

// MeshProviders provide an interface to monitoring and managing a specific
// mesh.
// AutoPilot does not abstract the mesh API - AutoPilot developers must
// still reason able about Provider-specific CRDs. AutoPilot's job is to
// abstract operational concerns such as discovering control plane configuration
// and monitoring metrics.
type MeshProvider int32

const (
	// the Operator will utilize the Service Mesh Interface (SMI) for metrics and configuration.
	// Compatible with multiple meshes (may require installation of an SMI Adapter).
	MeshProvider_SMI MeshProvider = 0
	// the Operator will utilize Istio mesh for metrics and configuration
	MeshProvider_Istio MeshProvider = 1
	// the Operator will utilize a locally deployed Prometheus instance for metrics
	MeshProvider_Custom MeshProvider = 2
)

var MeshProvider_name = map[int32]string{
	0: "SMI",
	1: "Istio",
	2: "Custom",
}

var MeshProvider_value = map[string]int32{
	"SMI":    0,
	"Istio":  1,
	"Custom": 2,
}

func (x MeshProvider) String() string {
	return proto.EnumName(MeshProvider_name, int32(x))
}

func (MeshProvider) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

// The AutoPilotOperator file is the bootstrap
// Configuration file for the Operator.
// It is stored and mounted to the operator as a Kubernetes ConfigMap.
// The Operator will hot-reload when the configuration file changes.
// Default name is 'autopilot-operator.yaml' and should be stored in the project root.
type AutoPilotOperator struct {
	// version of the operator
	// used for logging and metrics
	// default is "0.0.1"
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// meshProvider determines how the operator will connect to a service mesh
	// Default is "SMI"
	MeshProvider MeshProvider `protobuf:"varint,2,opt,name=meshProvider,proto3,enum=autopilot.MeshProvider" json:"meshProvider,omitempty"`
	// controlPlaneNs is the namespace the control plane lives in
	// Default is "istio-system"
	ControlPlaneNs string `protobuf:"bytes,3,opt,name=controlPlaneNs,proto3" json:"controlPlaneNs,omitempty"`
	// workInterval to sets the interval at which CRD workers resync.
	// Default is 5s
	WorkInterval *duration.Duration `protobuf:"bytes,4,opt,name=workInterval,proto3" json:"workInterval,omitempty"`
	// Serve metrics on this address. Set to empty string to disable metrics
	// defaults to ":9090"
	MetricsAddr string `protobuf:"bytes,5,opt,name=metricsAddr,proto3" json:"metricsAddr,omitempty"`
	// Enable leader election. This will prevent more than one operator from running at a time
	// defaults to true
	EnableLeaderElection bool `protobuf:"varint,6,opt,name=enableLeaderElection,proto3" json:"enableLeaderElection,omitempty"`
	// if non-empty, watchNamespace will restrict the Operator to watching resources in a single namespace
	// if empty (default), the Operator must have Cluster-scope RBAC permissions (ClusterRole/Binding)
	// can also be set via the WATCH_NAMESPACE environment variable
	WatchNamespace       string   `protobuf:"bytes,7,opt,name=watchNamespace,proto3" json:"watchNamespace,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AutoPilotOperator) Reset()         { *m = AutoPilotOperator{} }
func (m *AutoPilotOperator) String() string { return proto.CompactTextString(m) }
func (*AutoPilotOperator) ProtoMessage()    {}
func (*AutoPilotOperator) Descriptor() ([]byte, []int) {
	return fileDescriptor_56f975433f2c607a, []int{0}
}

func (m *AutoPilotOperator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AutoPilotOperator.Unmarshal(m, b)
}
func (m *AutoPilotOperator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AutoPilotOperator.Marshal(b, m, deterministic)
}
func (m *AutoPilotOperator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AutoPilotOperator.Merge(m, src)
}
func (m *AutoPilotOperator) XXX_Size() int {
	return xxx_messageInfo_AutoPilotOperator.Size(m)
}
func (m *AutoPilotOperator) XXX_DiscardUnknown() {
	xxx_messageInfo_AutoPilotOperator.DiscardUnknown(m)
}

var xxx_messageInfo_AutoPilotOperator proto.InternalMessageInfo

func (m *AutoPilotOperator) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AutoPilotOperator) GetMeshProvider() MeshProvider {
	if m != nil {
		return m.MeshProvider
	}
	return MeshProvider_SMI
}

func (m *AutoPilotOperator) GetControlPlaneNs() string {
	if m != nil {
		return m.ControlPlaneNs
	}
	return ""
}

func (m *AutoPilotOperator) GetWorkInterval() *duration.Duration {
	if m != nil {
		return m.WorkInterval
	}
	return nil
}

func (m *AutoPilotOperator) GetMetricsAddr() string {
	if m != nil {
		return m.MetricsAddr
	}
	return ""
}

func (m *AutoPilotOperator) GetEnableLeaderElection() bool {
	if m != nil {
		return m.EnableLeaderElection
	}
	return false
}

func (m *AutoPilotOperator) GetWatchNamespace() string {
	if m != nil {
		return m.WatchNamespace
	}
	return ""
}

func init() {
	proto.RegisterEnum("autopilot.MeshProvider", MeshProvider_name, MeshProvider_value)
	proto.RegisterType((*AutoPilotOperator)(nil), "autopilot.AutoPilotOperator")
}

func init() { proto.RegisterFile("autopilot-operator.proto", fileDescriptor_56f975433f2c607a) }

var fileDescriptor_56f975433f2c607a = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x4f, 0x4f, 0xc2, 0x40,
	0x10, 0xc5, 0x2d, 0x08, 0xc8, 0x42, 0x08, 0x6e, 0x4c, 0x5c, 0x3d, 0x98, 0x46, 0xa3, 0x69, 0x4c,
	0xd8, 0x46, 0x3c, 0x1a, 0x0f, 0xf8, 0xe7, 0x40, 0x22, 0x48, 0xea, 0xcd, 0xdb, 0xb6, 0x1d, 0x61,
	0xe3, 0xb6, 0xd3, 0xec, 0x6e, 0xcb, 0xc7, 0xf4, 0x2b, 0x19, 0xca, 0x9f, 0x50, 0xe3, 0x71, 0xde,
	0x7b, 0x79, 0x33, 0xbf, 0x21, 0x4c, 0xe4, 0x16, 0x33, 0xa9, 0xd0, 0x0e, 0x30, 0x03, 0x2d, 0x2c,
	0x6a, 0x9e, 0x69, 0xb4, 0x48, 0xdb, 0x3b, 0xe7, 0xfc, 0x62, 0x8e, 0x38, 0x57, 0xe0, 0x97, 0x46,
	0x98, 0x7f, 0xf9, 0x71, 0xae, 0x85, 0x95, 0x98, 0xae, 0xa3, 0x97, 0x3f, 0x35, 0x72, 0x3c, 0xca,
	0x2d, 0xce, 0x56, 0xe9, 0xf7, 0x4d, 0x0d, 0x65, 0xa4, 0x55, 0x80, 0x36, 0x12, 0x53, 0xe6, 0xb8,
	0x8e, 0xd7, 0x0e, 0xb6, 0x23, 0x7d, 0x20, 0xdd, 0x04, 0xcc, 0x62, 0xa6, 0xb1, 0x90, 0x31, 0x68,
	0x56, 0x73, 0x1d, 0xaf, 0x37, 0x3c, 0xe5, 0xbb, 0x8d, 0x7c, 0xb2, 0x67, 0x07, 0x95, 0x30, 0xbd,
	0x21, 0xbd, 0x08, 0x53, 0xab, 0x51, 0xcd, 0x94, 0x48, 0x61, 0x6a, 0x58, 0xbd, 0x6c, 0xff, 0xa3,
	0xd2, 0x47, 0xd2, 0x5d, 0xa2, 0xfe, 0x1e, 0xa7, 0x16, 0x74, 0x21, 0x14, 0x3b, 0x74, 0x1d, 0xaf,
	0x33, 0x3c, 0xe3, 0x6b, 0x16, 0xbe, 0x65, 0xe1, 0x2f, 0x1b, 0x96, 0xa0, 0x12, 0xa7, 0x2e, 0xe9,
	0x24, 0x60, 0xb5, 0x8c, 0xcc, 0x28, 0x8e, 0x35, 0x6b, 0x94, 0x3b, 0xf6, 0x25, 0x3a, 0x24, 0x27,
	0x90, 0x8a, 0x50, 0xc1, 0x1b, 0x88, 0x18, 0xf4, 0xab, 0x82, 0x68, 0xd5, 0xc3, 0x9a, 0xae, 0xe3,
	0x1d, 0x05, 0xff, 0x7a, 0xab, 0xe3, 0x97, 0xc2, 0x46, 0x8b, 0xa9, 0x48, 0xc0, 0x64, 0x22, 0x02,
	0xd6, 0x5a, 0x1f, 0x5f, 0x55, 0x6f, 0x39, 0xe9, 0xee, 0xbf, 0x80, 0xb6, 0x48, 0xfd, 0x63, 0x32,
	0xee, 0x1f, 0xd0, 0x36, 0x69, 0x8c, 0x8d, 0x95, 0xd8, 0x77, 0x28, 0x21, 0xcd, 0xe7, 0xdc, 0x58,
	0x4c, 0xfa, 0xb5, 0xa7, 0xeb, 0xcf, 0xab, 0xb9, 0xb4, 0x8b, 0x3c, 0xe4, 0x11, 0x26, 0xbe, 0x41,
	0x85, 0x03, 0x89, 0xfe, 0xee, 0x9f, 0xbe, 0xc8, 0xa4, 0x5f, 0xdc, 0x85, 0xcd, 0x92, 0xfa, 0xfe,
	0x37, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x4f, 0x6c, 0xc7, 0xf6, 0x01, 0x00, 0x00,
}