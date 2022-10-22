//
// * Copyright 2021 The Feast Authors
// *
// * Licensed under the Apache License, Version 2.0 (the "License");
// * you may not use this file except in compliance with the License.
// * You may obtain a copy of the License at
// *
// *     https://www.apache.org/licenses/LICENSE-2.0
// *
// * Unless required by applicable law or agreed to in writing, software
// * distributed under the License is distributed on an "AS IS" BASIS,
// * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// * See the License for the specific language governing permissions and
// * limitations under the License.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/feast/core/InfraObject.proto

package core

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

// Represents a set of infrastructure objects managed by Feast
type Infra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// List of infrastructure objects managed by Feast
	InfraObjects []*InfraObject `protobuf:"bytes,1,rep,name=infra_objects,json=infraObjects,proto3" json:"infra_objects,omitempty"`
}

func (x *Infra) Reset() {
	*x = Infra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_core_InfraObject_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Infra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Infra) ProtoMessage() {}

func (x *Infra) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_core_InfraObject_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Infra.ProtoReflect.Descriptor instead.
func (*Infra) Descriptor() ([]byte, []int) {
	return file_protos_feast_core_InfraObject_proto_rawDescGZIP(), []int{0}
}

func (x *Infra) GetInfraObjects() []*InfraObject {
	if x != nil {
		return x.InfraObjects
	}
	return nil
}

// Represents a single infrastructure object managed by Feast
type InfraObject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Represents the Python class for the infrastructure object
	InfraObjectClassType string `protobuf:"bytes,1,opt,name=infra_object_class_type,json=infraObjectClassType,proto3" json:"infra_object_class_type,omitempty"`
	// The infrastructure object
	//
	// Types that are assignable to InfraObject:
	//	*InfraObject_DynamodbTable
	//	*InfraObject_DatastoreTable
	//	*InfraObject_SqliteTable
	//	*InfraObject_CustomInfra_
	InfraObject isInfraObject_InfraObject `protobuf_oneof:"infra_object"`
}

func (x *InfraObject) Reset() {
	*x = InfraObject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_core_InfraObject_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfraObject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraObject) ProtoMessage() {}

func (x *InfraObject) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_core_InfraObject_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraObject.ProtoReflect.Descriptor instead.
func (*InfraObject) Descriptor() ([]byte, []int) {
	return file_protos_feast_core_InfraObject_proto_rawDescGZIP(), []int{1}
}

func (x *InfraObject) GetInfraObjectClassType() string {
	if x != nil {
		return x.InfraObjectClassType
	}
	return ""
}

func (m *InfraObject) GetInfraObject() isInfraObject_InfraObject {
	if m != nil {
		return m.InfraObject
	}
	return nil
}

func (x *InfraObject) GetDynamodbTable() *DynamoDBTable {
	if x, ok := x.GetInfraObject().(*InfraObject_DynamodbTable); ok {
		return x.DynamodbTable
	}
	return nil
}

func (x *InfraObject) GetDatastoreTable() *DatastoreTable {
	if x, ok := x.GetInfraObject().(*InfraObject_DatastoreTable); ok {
		return x.DatastoreTable
	}
	return nil
}

func (x *InfraObject) GetSqliteTable() *SqliteTable {
	if x, ok := x.GetInfraObject().(*InfraObject_SqliteTable); ok {
		return x.SqliteTable
	}
	return nil
}

func (x *InfraObject) GetCustomInfra() *InfraObject_CustomInfra {
	if x, ok := x.GetInfraObject().(*InfraObject_CustomInfra_); ok {
		return x.CustomInfra
	}
	return nil
}

type isInfraObject_InfraObject interface {
	isInfraObject_InfraObject()
}

type InfraObject_DynamodbTable struct {
	DynamodbTable *DynamoDBTable `protobuf:"bytes,2,opt,name=dynamodb_table,json=dynamodbTable,proto3,oneof"`
}

type InfraObject_DatastoreTable struct {
	DatastoreTable *DatastoreTable `protobuf:"bytes,3,opt,name=datastore_table,json=datastoreTable,proto3,oneof"`
}

type InfraObject_SqliteTable struct {
	SqliteTable *SqliteTable `protobuf:"bytes,4,opt,name=sqlite_table,json=sqliteTable,proto3,oneof"`
}

type InfraObject_CustomInfra_ struct {
	CustomInfra *InfraObject_CustomInfra `protobuf:"bytes,100,opt,name=custom_infra,json=customInfra,proto3,oneof"`
}

func (*InfraObject_DynamodbTable) isInfraObject_InfraObject() {}

func (*InfraObject_DatastoreTable) isInfraObject_InfraObject() {}

func (*InfraObject_SqliteTable) isInfraObject_InfraObject() {}

func (*InfraObject_CustomInfra_) isInfraObject_InfraObject() {}

// Allows for custom infra objects to be added
type InfraObject_CustomInfra struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Field []byte `protobuf:"bytes,1,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *InfraObject_CustomInfra) Reset() {
	*x = InfraObject_CustomInfra{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_feast_core_InfraObject_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InfraObject_CustomInfra) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InfraObject_CustomInfra) ProtoMessage() {}

func (x *InfraObject_CustomInfra) ProtoReflect() protoreflect.Message {
	mi := &file_protos_feast_core_InfraObject_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InfraObject_CustomInfra.ProtoReflect.Descriptor instead.
func (*InfraObject_CustomInfra) Descriptor() ([]byte, []int) {
	return file_protos_feast_core_InfraObject_proto_rawDescGZIP(), []int{1, 0}
}

func (x *InfraObject_CustomInfra) GetField() []byte {
	if x != nil {
		return x.Field
	}
	return nil
}

var File_protos_feast_core_InfraObject_proto protoreflect.FileDescriptor

var file_protos_feast_core_InfraObject_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x1a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f,
	0x63, 0x6f, 0x72, 0x65, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x44, 0x79, 0x6e,
	0x61, 0x6d, 0x6f, 0x44, 0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x23, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x53, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x05, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x3c,
	0x0a, 0x0d, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x0c,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x22, 0x8c, 0x03, 0x0a,
	0x0b, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x35, 0x0a, 0x17,
	0x69, 0x6e, 0x66, 0x72, 0x61, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x0e, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x64, 0x62, 0x5f,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x66, 0x65,
	0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x79, 0x6e, 0x61, 0x6d, 0x6f, 0x44,
	0x42, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x64, 0x79, 0x6e, 0x61, 0x6d, 0x6f,
	0x64, 0x62, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x64, 0x61, 0x74, 0x61, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x0e,
	0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x3c,
	0x0a, 0x0c, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x53, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52,
	0x0b, 0x73, 0x71, 0x6c, 0x69, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x48, 0x0a, 0x0c,
	0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x72, 0x61, 0x18, 0x64, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x23, 0x2e, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x48, 0x00, 0x52, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x1a, 0x23, 0x0a, 0x0b, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x49, 0x6e, 0x66, 0x72, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x69,
	0x6e, 0x66, 0x72, 0x61, 0x5f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x42, 0x30, 0x0a, 0x10, 0x66,
	0x65, 0x61, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x42,
	0x10, 0x49, 0x6e, 0x66, 0x72, 0x61, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x5a, 0x0a, 0x66, 0x65, 0x61, 0x73, 0x74, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_feast_core_InfraObject_proto_rawDescOnce sync.Once
	file_protos_feast_core_InfraObject_proto_rawDescData = file_protos_feast_core_InfraObject_proto_rawDesc
)

func file_protos_feast_core_InfraObject_proto_rawDescGZIP() []byte {
	file_protos_feast_core_InfraObject_proto_rawDescOnce.Do(func() {
		file_protos_feast_core_InfraObject_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_feast_core_InfraObject_proto_rawDescData)
	})
	return file_protos_feast_core_InfraObject_proto_rawDescData
}

var file_protos_feast_core_InfraObject_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protos_feast_core_InfraObject_proto_goTypes = []interface{}{
	(*Infra)(nil),                   // 0: feast.core.Infra
	(*InfraObject)(nil),             // 1: feast.core.InfraObject
	(*InfraObject_CustomInfra)(nil), // 2: feast.core.InfraObject.CustomInfra
	(*DynamoDBTable)(nil),           // 3: feast.core.DynamoDBTable
	(*DatastoreTable)(nil),          // 4: feast.core.DatastoreTable
	(*SqliteTable)(nil),             // 5: feast.core.SqliteTable
}
var file_protos_feast_core_InfraObject_proto_depIdxs = []int32{
	1, // 0: feast.core.Infra.infra_objects:type_name -> feast.core.InfraObject
	3, // 1: feast.core.InfraObject.dynamodb_table:type_name -> feast.core.DynamoDBTable
	4, // 2: feast.core.InfraObject.datastore_table:type_name -> feast.core.DatastoreTable
	5, // 3: feast.core.InfraObject.sqlite_table:type_name -> feast.core.SqliteTable
	2, // 4: feast.core.InfraObject.custom_infra:type_name -> feast.core.InfraObject.CustomInfra
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_protos_feast_core_InfraObject_proto_init() }
func file_protos_feast_core_InfraObject_proto_init() {
	if File_protos_feast_core_InfraObject_proto != nil {
		return
	}
	file_protos_feast_core_DatastoreTable_proto_init()
	file_protos_feast_core_DynamoDBTable_proto_init()
	file_protos_feast_core_SqliteTable_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protos_feast_core_InfraObject_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Infra); i {
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
		file_protos_feast_core_InfraObject_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfraObject); i {
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
		file_protos_feast_core_InfraObject_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InfraObject_CustomInfra); i {
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
	file_protos_feast_core_InfraObject_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*InfraObject_DynamodbTable)(nil),
		(*InfraObject_DatastoreTable)(nil),
		(*InfraObject_SqliteTable)(nil),
		(*InfraObject_CustomInfra_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_feast_core_InfraObject_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protos_feast_core_InfraObject_proto_goTypes,
		DependencyIndexes: file_protos_feast_core_InfraObject_proto_depIdxs,
		MessageInfos:      file_protos_feast_core_InfraObject_proto_msgTypes,
	}.Build()
	File_protos_feast_core_InfraObject_proto = out.File
	file_protos_feast_core_InfraObject_proto_rawDesc = nil
	file_protos_feast_core_InfraObject_proto_goTypes = nil
	file_protos_feast_core_InfraObject_proto_depIdxs = nil
}