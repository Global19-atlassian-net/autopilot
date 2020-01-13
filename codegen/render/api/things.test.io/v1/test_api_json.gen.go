// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/autopilot/codegen/render/api/test_api.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for PaintSpec
func (this *PaintSpec) MarshalJSON() ([]byte, error) {
	str, err := TestApiMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PaintSpec
func (this *PaintSpec) UnmarshalJSON(b []byte) error {
	return TestApiUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PaintColor
func (this *PaintColor) MarshalJSON() ([]byte, error) {
	str, err := TestApiMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PaintColor
func (this *PaintColor) UnmarshalJSON(b []byte) error {
	return TestApiUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for AcrylicType
func (this *AcrylicType) MarshalJSON() ([]byte, error) {
	str, err := TestApiMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for AcrylicType
func (this *AcrylicType) UnmarshalJSON(b []byte) error {
	return TestApiUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for OilType
func (this *OilType) MarshalJSON() ([]byte, error) {
	str, err := TestApiMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for OilType
func (this *OilType) UnmarshalJSON(b []byte) error {
	return TestApiUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PaintStatus
func (this *PaintStatus) MarshalJSON() ([]byte, error) {
	str, err := TestApiMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PaintStatus
func (this *PaintStatus) UnmarshalJSON(b []byte) error {
	return TestApiUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	TestApiMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	TestApiUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
