// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: third_party/tink/proto/aes_ctr.proto

package aes_ctr_go_proto

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

type AesCtrParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IvSize uint32 `protobuf:"varint,1,opt,name=iv_size,json=ivSize,proto3" json:"iv_size,omitempty"`
}

func (x *AesCtrParams) Reset() {
	*x = AesCtrParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrParams) ProtoMessage() {}

func (x *AesCtrParams) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrParams.ProtoReflect.Descriptor instead.
func (*AesCtrParams) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_proto_rawDescGZIP(), []int{0}
}

func (x *AesCtrParams) GetIvSize() uint32 {
	if x != nil {
		return x.IvSize
	}
	return 0
}

type AesCtrKeyFormat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Params  *AesCtrParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	KeySize uint32        `protobuf:"varint,2,opt,name=key_size,json=keySize,proto3" json:"key_size,omitempty"`
}

func (x *AesCtrKeyFormat) Reset() {
	*x = AesCtrKeyFormat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrKeyFormat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrKeyFormat) ProtoMessage() {}

func (x *AesCtrKeyFormat) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrKeyFormat.ProtoReflect.Descriptor instead.
func (*AesCtrKeyFormat) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_proto_rawDescGZIP(), []int{1}
}

func (x *AesCtrKeyFormat) GetParams() *AesCtrParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AesCtrKeyFormat) GetKeySize() uint32 {
	if x != nil {
		return x.KeySize
	}
	return 0
}

// key_type: type.googleapis.com/google.crypto.tink.AesCtrKey
type AesCtrKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version  uint32        `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Params   *AesCtrParams `protobuf:"bytes,2,opt,name=params,proto3" json:"params,omitempty"`
	KeyValue []byte        `protobuf:"bytes,3,opt,name=key_value,json=keyValue,proto3" json:"key_value,omitempty"`
}

func (x *AesCtrKey) Reset() {
	*x = AesCtrKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AesCtrKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AesCtrKey) ProtoMessage() {}

func (x *AesCtrKey) ProtoReflect() protoreflect.Message {
	mi := &file_third_party_tink_proto_aes_ctr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AesCtrKey.ProtoReflect.Descriptor instead.
func (*AesCtrKey) Descriptor() ([]byte, []int) {
	return file_third_party_tink_proto_aes_ctr_proto_rawDescGZIP(), []int{2}
}

func (x *AesCtrKey) GetVersion() uint32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *AesCtrKey) GetParams() *AesCtrParams {
	if x != nil {
		return x.Params
	}
	return nil
}

func (x *AesCtrKey) GetKeyValue() []byte {
	if x != nil {
		return x.KeyValue
	}
	return nil
}

var File_third_party_tink_proto_aes_ctr_proto protoreflect.FileDescriptor

var file_third_party_tink_proto_aes_ctr_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x68, 0x69, 0x72, 0x64, 0x5f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x74, 0x69,
	0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73, 0x5f, 0x63, 0x74, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x22, 0x27, 0x0a, 0x0c, 0x41, 0x65,
	0x73, 0x43, 0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x76,
	0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x69, 0x76, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x66, 0x0a, 0x0f, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x4b, 0x65, 0x79,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x65, 0x73, 0x43,
	0x74, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73,
	0x12, 0x19, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6b, 0x65, 0x79, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x7c, 0x0a, 0x09, 0x41,
	0x65, 0x73, 0x43, 0x74, 0x72, 0x4b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x12, 0x38, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x2e, 0x74, 0x69, 0x6e, 0x6b, 0x2e, 0x41, 0x65, 0x73, 0x43, 0x74, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x52, 0x06, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6b, 0x65, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x08, 0x6b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x4f, 0x0a, 0x1c, 0x63, 0x6f, 0x6d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x74,
	0x69, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x74,
	0x69, 0x6e, 0x6b, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x65, 0x73, 0x5f, 0x63, 0x74,
	0x72, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_third_party_tink_proto_aes_ctr_proto_rawDescOnce sync.Once
	file_third_party_tink_proto_aes_ctr_proto_rawDescData = file_third_party_tink_proto_aes_ctr_proto_rawDesc
)

func file_third_party_tink_proto_aes_ctr_proto_rawDescGZIP() []byte {
	file_third_party_tink_proto_aes_ctr_proto_rawDescOnce.Do(func() {
		file_third_party_tink_proto_aes_ctr_proto_rawDescData = protoimpl.X.CompressGZIP(file_third_party_tink_proto_aes_ctr_proto_rawDescData)
	})
	return file_third_party_tink_proto_aes_ctr_proto_rawDescData
}

var file_third_party_tink_proto_aes_ctr_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_third_party_tink_proto_aes_ctr_proto_goTypes = []interface{}{
	(*AesCtrParams)(nil),    // 0: google.crypto.tink.AesCtrParams
	(*AesCtrKeyFormat)(nil), // 1: google.crypto.tink.AesCtrKeyFormat
	(*AesCtrKey)(nil),       // 2: google.crypto.tink.AesCtrKey
}
var file_third_party_tink_proto_aes_ctr_proto_depIdxs = []int32{
	0, // 0: google.crypto.tink.AesCtrKeyFormat.params:type_name -> google.crypto.tink.AesCtrParams
	0, // 1: google.crypto.tink.AesCtrKey.params:type_name -> google.crypto.tink.AesCtrParams
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_third_party_tink_proto_aes_ctr_proto_init() }
func file_third_party_tink_proto_aes_ctr_proto_init() {
	if File_third_party_tink_proto_aes_ctr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_third_party_tink_proto_aes_ctr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrParams); i {
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
		file_third_party_tink_proto_aes_ctr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrKeyFormat); i {
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
		file_third_party_tink_proto_aes_ctr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AesCtrKey); i {
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
			RawDescriptor: file_third_party_tink_proto_aes_ctr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_third_party_tink_proto_aes_ctr_proto_goTypes,
		DependencyIndexes: file_third_party_tink_proto_aes_ctr_proto_depIdxs,
		MessageInfos:      file_third_party_tink_proto_aes_ctr_proto_msgTypes,
	}.Build()
	File_third_party_tink_proto_aes_ctr_proto = out.File
	file_third_party_tink_proto_aes_ctr_proto_rawDesc = nil
	file_third_party_tink_proto_aes_ctr_proto_goTypes = nil
	file_third_party_tink_proto_aes_ctr_proto_depIdxs = nil
}
