//
// Copyright 2020 The Feast Authors
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
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/feast/core/Feature.proto

package core

import (
	types "github.com/Golang-Tools/feasthelper/feast/types"
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

type FeatureSpecV2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the feature. Not updatable.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Value type of the feature. Not updatable.
	ValueType types.ValueType_Enum `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=feast.types.ValueType_Enum" json:"value_type,omitempty"`
	// Tags for user defined metadata on a feature
	Tags map[string]string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FeatureSpecV2) Reset() {
	*x = FeatureSpecV2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_core_Feature_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureSpecV2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureSpecV2) ProtoMessage() {}

func (x *FeatureSpecV2) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_core_Feature_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureSpecV2.ProtoReflect.Descriptor instead.
func (*FeatureSpecV2) Descriptor() ([]byte, []int) {
	return file_protos_feast_core_Feature_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureSpecV2) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FeatureSpecV2) GetValueType() types.ValueType_Enum {
	if x != nil {
		return x.ValueType
	}
	return types.ValueType_Enum(0)
}

func (x *FeatureSpecV2) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_protos_feast_core_Feature_proto protoreflect.FileDescriptor

var file_protos_feast_core_Feature_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x1e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x2f, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd1, 0x01,
	0x0a, 0x0d, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x56, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x37, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x46, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x53, 0x70, 0x65, 0x63, 0x56, 0x32, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x42, 0x2c, 0x0a, 0x10, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42, 0x0c, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x5a, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_feast_core_Feature_proto_rawDescOnce sync.Once
	file_protos_feast_core_Feature_proto_rawDescData = file_protos_feast_core_Feature_proto_rawDesc
)

func file_protos_feast_core_Feature_proto_rawDescGZIP() []byte {
	file_protos_feast_core_Feature_proto_rawDescOnce.Do(func() {
		file_protos_feast_core_Feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_feast_core_Feature_proto_rawDescData)
	})
	return file_protos_feast_core_Feature_proto_rawDescData
}

var file_protos_feast_core_Feature_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_feast_core_Feature_proto_goTypes = []interface{}{
	(*FeatureSpecV2)(nil),     // 0: feast.core.FeatureSpecV2
	nil,                       // 1: feast.core.FeatureSpecV2.TagsEntry
	(types.ValueType_Enum)(0), // 2: feast.types.ValueType.Enum
}
var file_protos_feast_core_Feature_proto_depIdxs = []int32{
	2, // 0: feast.core.FeatureSpecV2.value_type:type_name -> feast.types.ValueType.Enum
	1, // 1: feast.core.FeatureSpecV2.tags:type_name -> feast.core.FeatureSpecV2.TagsEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_feast_core_Feature_proto_init() }
func file_protos_feast_core_Feature_proto_init() {
	if File_protos_feast_core_Feature_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_feast_core_Feature_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureSpecV2); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_feast_core_Feature_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_feast_core_Feature_proto_goTypes,
		DependencyIndexes: file_protos_feast_core_Feature_proto_depIdxs,
		MessageInfos:      file_protos_feast_core_Feature_proto_msgTypes,
	}.Build()
	File_protos_feast_core_Feature_proto = out.File
	file_protos_feast_core_Feature_proto_rawDesc = nil
	file_protos_feast_core_Feature_proto_goTypes = nil
	file_protos_feast_core_Feature_proto_depIdxs = nil
}
