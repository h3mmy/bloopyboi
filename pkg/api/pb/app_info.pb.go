// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: proto/app_info.proto

package pb

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

type AppInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AppInfoRequest) Reset() {
	*x = AppInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfoRequest) ProtoMessage() {}

func (x *AppInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfoRequest.ProtoReflect.Descriptor instead.
func (*AppInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_app_info_proto_rawDescGZIP(), []int{0}
}

type AppInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *AppInfoResponse) Reset() {
	*x = AppInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_app_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppInfoResponse) ProtoMessage() {}

func (x *AppInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_app_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppInfoResponse.ProtoReflect.Descriptor instead.
func (*AppInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_app_info_proto_rawDescGZIP(), []int{1}
}

func (x *AppInfoResponse) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_proto_app_info_proto protoreflect.FileDescriptor

var file_proto_app_info_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x6e, 0x66, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x10, 0x0a, 0x0e,
	0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x2b,
	0x0a, 0x0f, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x47, 0x0a, 0x0b, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x41, 0x70,
	0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x2e, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x70, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x69, 0x6e,
	0x66, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_app_info_proto_rawDescOnce sync.Once
	file_proto_app_info_proto_rawDescData = file_proto_app_info_proto_rawDesc
)

func file_proto_app_info_proto_rawDescGZIP() []byte {
	file_proto_app_info_proto_rawDescOnce.Do(func() {
		file_proto_app_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_app_info_proto_rawDescData)
	})
	return file_proto_app_info_proto_rawDescData
}

var file_proto_app_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_app_info_proto_goTypes = []interface{}{
	(*AppInfoRequest)(nil),  // 0: info.AppInfoRequest
	(*AppInfoResponse)(nil), // 1: info.AppInfoResponse
}
var file_proto_app_info_proto_depIdxs = []int32{
	0, // 0: info.InfoService.AppInfo:input_type -> info.AppInfoRequest
	1, // 1: info.InfoService.AppInfo:output_type -> info.AppInfoResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_app_info_proto_init() }
func file_proto_app_info_proto_init() {
	if File_proto_app_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_app_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfoRequest); i {
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
		file_proto_app_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppInfoResponse); i {
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
			RawDescriptor: file_proto_app_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_app_info_proto_goTypes,
		DependencyIndexes: file_proto_app_info_proto_depIdxs,
		MessageInfos:      file_proto_app_info_proto_msgTypes,
	}.Build()
	File_proto_app_info_proto = out.File
	file_proto_app_info_proto_rawDesc = nil
	file_proto_app_info_proto_goTypes = nil
	file_proto_app_info_proto_depIdxs = nil
}
