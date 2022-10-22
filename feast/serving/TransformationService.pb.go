//
// Copyright 2021 The Feast Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/feast/serving/TransformationService.proto

package serving

import (
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

type TransformationServiceType int32

const (
	TransformationServiceType_TRANSFORMATION_SERVICE_TYPE_INVALID TransformationServiceType = 0
	TransformationServiceType_TRANSFORMATION_SERVICE_TYPE_PYTHON  TransformationServiceType = 1
	TransformationServiceType_TRANSFORMATION_SERVICE_TYPE_CUSTOM  TransformationServiceType = 100
)

// Enum value maps for TransformationServiceType.
var (
	TransformationServiceType_name = map[int32]string{
		0:   "TRANSFORMATION_SERVICE_TYPE_INVALID",
		1:   "TRANSFORMATION_SERVICE_TYPE_PYTHON",
		100: "TRANSFORMATION_SERVICE_TYPE_CUSTOM",
	}
	TransformationServiceType_value = map[string]int32{
		"TRANSFORMATION_SERVICE_TYPE_INVALID": 0,
		"TRANSFORMATION_SERVICE_TYPE_PYTHON":  1,
		"TRANSFORMATION_SERVICE_TYPE_CUSTOM":  100,
	}
)

func (x TransformationServiceType) Enum() *TransformationServiceType {
	p := new(TransformationServiceType)
	*p = x
	return p
}

func (x TransformationServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TransformationServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_feast_serving_TransformationService_proto_enumTypes[0].Descriptor()
}

func (TransformationServiceType) Type() protoreflect.EnumType {
	return &file_protos_feast_serving_TransformationService_proto_enumTypes[0]
}

func (x TransformationServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TransformationServiceType.Descriptor instead.
func (TransformationServiceType) EnumDescriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{0}
}

type ValueType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Value:
	//	*ValueType_ArrowValue
	Value isValueType_Value `protobuf_oneof:"value"`
}

func (x *ValueType) Reset() {
	*x = ValueType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValueType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValueType) ProtoMessage() {}

func (x *ValueType) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValueType.ProtoReflect.Descriptor instead.
func (*ValueType) Descriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{0}
}

func (m *ValueType) GetValue() isValueType_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *ValueType) GetArrowValue() []byte {
	if x, ok := x.GetValue().(*ValueType_ArrowValue); ok {
		return x.ArrowValue
	}
	return nil
}

type isValueType_Value interface {
	isValueType_Value()
}

type ValueType_ArrowValue struct {
	// Having a oneOf provides forward compatibility if we need to support compound types
	// that are not supported by arrow natively.
	ArrowValue []byte `protobuf:"bytes,1,opt,name=arrow_value,json=arrowValue,proto3,oneof"`
}

func (*ValueType_ArrowValue) isValueType_Value() {}

type GetTransformationServiceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetTransformationServiceInfoRequest) Reset() {
	*x = GetTransformationServiceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformationServiceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformationServiceInfoRequest) ProtoMessage() {}

func (x *GetTransformationServiceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformationServiceInfoRequest.ProtoReflect.Descriptor instead.
func (*GetTransformationServiceInfoRequest) Descriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{1}
}

type GetTransformationServiceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Feast version of this transformation service deployment.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Type of transformation service deployment. This is either Python, or custom
	Type                             TransformationServiceType `protobuf:"varint,2,opt,name=type,proto3,enum=feast.serving.TransformationServiceType" json:"type,omitempty"`
	TransformationServiceTypeDetails string                    `protobuf:"bytes,3,opt,name=transformation_service_type_details,json=transformationServiceTypeDetails,proto3" json:"transformation_service_type_details,omitempty"`
}

func (x *GetTransformationServiceInfoResponse) Reset() {
	*x = GetTransformationServiceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransformationServiceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransformationServiceInfoResponse) ProtoMessage() {}

func (x *GetTransformationServiceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransformationServiceInfoResponse.ProtoReflect.Descriptor instead.
func (*GetTransformationServiceInfoResponse) Descriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransformationServiceInfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetTransformationServiceInfoResponse) GetType() TransformationServiceType {
	if x != nil {
		return x.Type
	}
	return TransformationServiceType_TRANSFORMATION_SERVICE_TYPE_INVALID
}

func (x *GetTransformationServiceInfoResponse) GetTransformationServiceTypeDetails() string {
	if x != nil {
		return x.TransformationServiceTypeDetails
	}
	return ""
}

type TransformFeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OnDemandFeatureViewName string     `protobuf:"bytes,1,opt,name=on_demand_feature_view_name,json=onDemandFeatureViewName,proto3" json:"on_demand_feature_view_name,omitempty"`
	Project                 string     `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	TransformationInput     *ValueType `protobuf:"bytes,3,opt,name=transformation_input,json=transformationInput,proto3" json:"transformation_input,omitempty"`
}

func (x *TransformFeaturesRequest) Reset() {
	*x = TransformFeaturesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformFeaturesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformFeaturesRequest) ProtoMessage() {}

func (x *TransformFeaturesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformFeaturesRequest.ProtoReflect.Descriptor instead.
func (*TransformFeaturesRequest) Descriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{3}
}

func (x *TransformFeaturesRequest) GetOnDemandFeatureViewName() string {
	if x != nil {
		return x.OnDemandFeatureViewName
	}
	return ""
}

func (x *TransformFeaturesRequest) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *TransformFeaturesRequest) GetTransformationInput() *ValueType {
	if x != nil {
		return x.TransformationInput
	}
	return nil
}

type TransformFeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransformationOutput *ValueType `protobuf:"bytes,3,opt,name=transformation_output,json=transformationOutput,proto3" json:"transformation_output,omitempty"`
}

func (x *TransformFeaturesResponse) Reset() {
	*x = TransformFeaturesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransformFeaturesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransformFeaturesResponse) ProtoMessage() {}

func (x *TransformFeaturesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_serving_TransformationService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransformFeaturesResponse.ProtoReflect.Descriptor instead.
func (*TransformFeaturesResponse) Descriptor() ([]byte, []int) {
	return file_protos_feast_serving_TransformationService_proto_rawDescGZIP(), []int{4}
}

func (x *TransformFeaturesResponse) GetTransformationOutput() *ValueType {
	if x != nil {
		return x.TransformationOutput
	}
	return nil
}

var File_protos_feast_serving_TransformationService_proto protoreflect.FileDescriptor

var file_protos_feast_serving_TransformationService_proto_rawDesc = []byte{
	0x0a, 0x30, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x22, 0x37, 0x0a, 0x09, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0b, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x72, 0x72, 0x6f, 0x77, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x42, 0x07, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x25, 0x0a, 0x23, 0x47, 0x65,
	0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xcd, 0x01, 0x0a, 0x24, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x28, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x4d, 0x0a, 0x23, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x20, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x73, 0x22, 0xbf, 0x01, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x46,
	0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c,
	0x0a, 0x1b, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x5f, 0x66, 0x65, 0x61, 0x74,
	0x75, 0x72, 0x65, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x17, 0x6f, 0x6e, 0x44, 0x65, 0x6d, 0x61, 0x6e, 0x64, 0x46, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x56, 0x69, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x4b, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x13,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x22, 0x6a, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d,
	0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4d, 0x0a, 0x15, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x18, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x2a,
	0x94, 0x01, 0x0a, 0x19, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x27, 0x0a,
	0x23, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46,
	0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x59, 0x54, 0x48, 0x4f, 0x4e, 0x10, 0x01, 0x12, 0x26,
	0x0a, 0x22, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x46, 0x4f, 0x52, 0x4d, 0x41, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x55,
	0x53, 0x54, 0x4f, 0x4d, 0x10, 0x64, 0x32, 0x89, 0x02, 0x0a, 0x15, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x87, 0x01, 0x0a, 0x1c, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x32, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e,
	0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x33, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x66, 0x0a, 0x11, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12,
	0x27, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e,
	0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74,
	0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f,
	0x72, 0x6d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x43, 0x0a, 0x13, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x42, 0x1d, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x41, 0x50, 0x49, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x5a, 0x0d, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_feast_serving_TransformationService_proto_rawDescOnce sync.Once
	file_protos_feast_serving_TransformationService_proto_rawDescData = file_protos_feast_serving_TransformationService_proto_rawDesc
)

func file_protos_feast_serving_TransformationService_proto_rawDescGZIP() []byte {
	file_protos_feast_serving_TransformationService_proto_rawDescOnce.Do(func() {
		file_protos_feast_serving_TransformationService_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_feast_serving_TransformationService_proto_rawDescData)
	})
	return file_protos_feast_serving_TransformationService_proto_rawDescData
}

var file_protos_feast_serving_TransformationService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_feast_serving_TransformationService_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_feast_serving_TransformationService_proto_goTypes = []interface{}{
	(TransformationServiceType)(0),               // 0: feast.serving.TransformationServiceType
	(*ValueType)(nil),                            // 1: feast.serving.ValueType
	(*GetTransformationServiceInfoRequest)(nil),  // 2: feast.serving.GetTransformationServiceInfoRequest
	(*GetTransformationServiceInfoResponse)(nil), // 3: feast.serving.GetTransformationServiceInfoResponse
	(*TransformFeaturesRequest)(nil),             // 4: feast.serving.TransformFeaturesRequest
	(*TransformFeaturesResponse)(nil),            // 5: feast.serving.TransformFeaturesResponse
}
var file_protos_feast_serving_TransformationService_proto_depIdxs = []int32{
	0, // 0: feast.serving.GetTransformationServiceInfoResponse.type:type_name -> feast.serving.TransformationServiceType
	1, // 1: feast.serving.TransformFeaturesRequest.transformation_input:type_name -> feast.serving.ValueType
	1, // 2: feast.serving.TransformFeaturesResponse.transformation_output:type_name -> feast.serving.ValueType
	2, // 3: feast.serving.TransformationService.GetTransformationServiceInfo:input_type -> feast.serving.GetTransformationServiceInfoRequest
	4, // 4: feast.serving.TransformationService.TransformFeatures:input_type -> feast.serving.TransformFeaturesRequest
	3, // 5: feast.serving.TransformationService.GetTransformationServiceInfo:output_type -> feast.serving.GetTransformationServiceInfoResponse
	5, // 6: feast.serving.TransformationService.TransformFeatures:output_type -> feast.serving.TransformFeaturesResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_protos_feast_serving_TransformationService_proto_init() }
func file_protos_feast_serving_TransformationService_proto_init() {
	if File_protos_feast_serving_TransformationService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_feast_serving_TransformationService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValueType); i {
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
		file_protos_feast_serving_TransformationService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformationServiceInfoRequest); i {
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
		file_protos_feast_serving_TransformationService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransformationServiceInfoResponse); i {
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
		file_protos_feast_serving_TransformationService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformFeaturesRequest); i {
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
		file_protos_feast_serving_TransformationService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransformFeaturesResponse); i {
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
	file_protos_feast_serving_TransformationService_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ValueType_ArrowValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_feast_serving_TransformationService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_feast_serving_TransformationService_proto_goTypes,
		DependencyIndexes: file_protos_feast_serving_TransformationService_proto_depIdxs,
		EnumInfos:         file_protos_feast_serving_TransformationService_proto_enumTypes,
		MessageInfos:      file_protos_feast_serving_TransformationService_proto_msgTypes,
	}.Build()
	File_protos_feast_serving_TransformationService_proto = out.File
	file_protos_feast_serving_TransformationService_proto_rawDesc = nil
	file_protos_feast_serving_TransformationService_proto_goTypes = nil
	file_protos_feast_serving_TransformationService_proto_depIdxs = nil
}
