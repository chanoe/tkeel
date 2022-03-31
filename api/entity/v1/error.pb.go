//
//Copyright 2021 The tKeel Authors.
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: api/entity/v1/error.proto

package v1

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

// @plugins=protoc-gen-go-errors
// 错误
type Error int32

const (
	// @msg=未知类型
	// @code=UNKNOWN
	Error_ERR_UNKNOWN Error = 0
	// @msg=请求参数无效
	// @code=INVALID_ARGUMENT
	Error_ERR_INVALID_ARGUMENT Error = 1
	// @msg=请求后端存储错误
	// @code=INTERNAL
	Error_ERR_INTERNAL_STORE Error = 2
	// @msg=内部错误
	// @code=INTERNAL
	Error_ERR_INTERNAL_ERROR Error = 3
	// @msg=资源不存在
	// @code=INTERNAL
	Error_ERR_RESOURCE_NOT_FOUND Error = 4
)

// Enum value maps for Error.
var (
	Error_name = map[int32]string{
		0: "ERR_UNKNOWN",
		1: "ERR_INVALID_ARGUMENT",
		2: "ERR_INTERNAL_STORE",
		3: "ERR_INTERNAL_ERROR",
		4: "ERR_RESOURCE_NOT_FOUND",
	}
	Error_value = map[string]int32{
		"ERR_UNKNOWN":            0,
		"ERR_INVALID_ARGUMENT":   1,
		"ERR_INTERNAL_STORE":     2,
		"ERR_INTERNAL_ERROR":     3,
		"ERR_RESOURCE_NOT_FOUND": 4,
	}
)

func (x Error) Enum() *Error {
	p := new(Error)
	*p = x
	return p
}

func (x Error) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Error) Descriptor() protoreflect.EnumDescriptor {
	return file_api_entity_v1_error_proto_enumTypes[0].Descriptor()
}

func (Error) Type() protoreflect.EnumType {
	return &file_api_entity_v1_error_proto_enumTypes[0]
}

func (x Error) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Error.Descriptor instead.
func (Error) EnumDescriptor() ([]byte, []int) {
	return file_api_entity_v1_error_proto_rawDescGZIP(), []int{0}
}

var File_api_entity_v1_error_proto protoreflect.FileDescriptor

var file_api_entity_v1_error_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x69, 0x6f, 0x2e,
	0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2a, 0x7e, 0x0a, 0x05,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x52, 0x52, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e,
	0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x41, 0x52, 0x47, 0x55, 0x4d, 0x45, 0x4e, 0x54, 0x10, 0x01,
	0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x53, 0x54, 0x4f, 0x52, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x5f,
	0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x03,
	0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x04, 0x42, 0x4f, 0x0a, 0x1f,
	0x69, 0x6f, 0x2e, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74,
	0x79, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x50,
	0x01, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6b,
	0x65, 0x65, 0x6c, 0x2d, 0x69, 0x6f, 0x2f, 0x74, 0x6b, 0x65, 0x65, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_entity_v1_error_proto_rawDescOnce sync.Once
	file_api_entity_v1_error_proto_rawDescData = file_api_entity_v1_error_proto_rawDesc
)

func file_api_entity_v1_error_proto_rawDescGZIP() []byte {
	file_api_entity_v1_error_proto_rawDescOnce.Do(func() {
		file_api_entity_v1_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_entity_v1_error_proto_rawDescData)
	})
	return file_api_entity_v1_error_proto_rawDescData
}

var file_api_entity_v1_error_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_entity_v1_error_proto_goTypes = []interface{}{
	(Error)(0), // 0: io.tkeel.security.api.entity.v1.Error
}
var file_api_entity_v1_error_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_entity_v1_error_proto_init() }
func file_api_entity_v1_error_proto_init() {
	if File_api_entity_v1_error_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_entity_v1_error_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_entity_v1_error_proto_goTypes,
		DependencyIndexes: file_api_entity_v1_error_proto_depIdxs,
		EnumInfos:         file_api_entity_v1_error_proto_enumTypes,
	}.Build()
	File_api_entity_v1_error_proto = out.File
	file_api_entity_v1_error_proto_rawDesc = nil
	file_api_entity_v1_error_proto_goTypes = nil
	file_api_entity_v1_error_proto_depIdxs = nil
}
