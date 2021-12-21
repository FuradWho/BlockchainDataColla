// 指定的当前proto语法的版本，有2和3

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: csr.proto

// 指定等会文件生成出来的package

package crs

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

// 定义request model
type CsrRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cn  string `protobuf:"bytes,1,opt,name=cn,proto3" json:"cn,omitempty"`
	Csr []byte `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
}

func (x *CsrRequest) Reset() {
	*x = CsrRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csr_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsrRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsrRequest) ProtoMessage() {}

func (x *CsrRequest) ProtoReflect() protoreflect.Message {
	mi := &file_csr_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsrRequest.ProtoReflect.Descriptor instead.
func (*CsrRequest) Descriptor() ([]byte, []int) {
	return file_csr_proto_rawDescGZIP(), []int{0}
}

func (x *CsrRequest) GetCn() string {
	if x != nil {
		return x.Cn
	}
	return ""
}

func (x *CsrRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

type CsrResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize int32  `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	Crt      []byte `protobuf:"bytes,5,opt,name=crt,proto3" json:"crt,omitempty"`
}

func (x *CsrResponse) Reset() {
	*x = CsrResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csr_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CsrResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CsrResponse) ProtoMessage() {}

func (x *CsrResponse) ProtoReflect() protoreflect.Message {
	mi := &file_csr_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CsrResponse.ProtoReflect.Descriptor instead.
func (*CsrResponse) Descriptor() ([]byte, []int) {
	return file_csr_proto_rawDescGZIP(), []int{1}
}

func (x *CsrResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CsrResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CsrResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CsrResponse) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *CsrResponse) GetCrt() []byte {
	if x != nil {
		return x.Crt
	}
	return nil
}

type CaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CaRequest) Reset() {
	*x = CaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csr_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaRequest) ProtoMessage() {}

func (x *CaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_csr_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaRequest.ProtoReflect.Descriptor instead.
func (*CaRequest) Descriptor() ([]byte, []int) {
	return file_csr_proto_rawDescGZIP(), []int{2}
}

type CaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code     string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	FileName string `protobuf:"bytes,3,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FileSize int32  `protobuf:"varint,4,opt,name=file_size,json=fileSize,proto3" json:"file_size,omitempty"`
	CaCrt    []byte `protobuf:"bytes,5,opt,name=ca_crt,json=caCrt,proto3" json:"ca_crt,omitempty"`
}

func (x *CaResponse) Reset() {
	*x = CaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_csr_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaResponse) ProtoMessage() {}

func (x *CaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_csr_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaResponse.ProtoReflect.Descriptor instead.
func (*CaResponse) Descriptor() ([]byte, []int) {
	return file_csr_proto_rawDescGZIP(), []int{3}
}

func (x *CaResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CaResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *CaResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *CaResponse) GetFileSize() int32 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *CaResponse) GetCaCrt() []byte {
	if x != nil {
		return x.CaCrt
	}
	return nil
}

var File_csr_proto protoreflect.FileDescriptor

var file_csr_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x73, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x72, 0x73,
	0x22, 0x2e, 0x0a, 0x0a, 0x43, 0x73, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x63, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x63, 0x6e, 0x12, 0x10,
	0x0a, 0x03, 0x63, 0x73, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x73, 0x72,
	0x22, 0x7f, 0x0a, 0x0b, 0x43, 0x73, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x72,
	0x74, 0x22, 0x0b, 0x0a, 0x09, 0x43, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x83,
	0x01, 0x0a, 0x0a, 0x43, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x15, 0x0a,
	0x06, 0x63, 0x61, 0x5f, 0x63, 0x72, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x63,
	0x61, 0x43, 0x72, 0x74, 0x32, 0x6b, 0x0a, 0x0a, 0x43, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x43, 0x73, 0x72, 0x12, 0x0f, 0x2e,
	0x63, 0x72, 0x73, 0x2e, 0x43, 0x73, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10,
	0x2e, 0x63, 0x72, 0x73, 0x2e, 0x43, 0x73, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x43, 0x61, 0x43, 0x72, 0x74, 0x12, 0x0e,
	0x2e, 0x63, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x63, 0x72, 0x73, 0x2e, 0x43, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_csr_proto_rawDescOnce sync.Once
	file_csr_proto_rawDescData = file_csr_proto_rawDesc
)

func file_csr_proto_rawDescGZIP() []byte {
	file_csr_proto_rawDescOnce.Do(func() {
		file_csr_proto_rawDescData = protoimpl.X.CompressGZIP(file_csr_proto_rawDescData)
	})
	return file_csr_proto_rawDescData
}

var file_csr_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_csr_proto_goTypes = []interface{}{
	(*CsrRequest)(nil),  // 0: crs.CsrRequest
	(*CsrResponse)(nil), // 1: crs.CsrResponse
	(*CaRequest)(nil),   // 2: crs.CaRequest
	(*CaResponse)(nil),  // 3: crs.CaResponse
}
var file_csr_proto_depIdxs = []int32{
	0, // 0: crs.CrsService.SendCsr:input_type -> crs.CsrRequest
	2, // 1: crs.CrsService.GetCaCrt:input_type -> crs.CaRequest
	1, // 2: crs.CrsService.SendCsr:output_type -> crs.CsrResponse
	3, // 3: crs.CrsService.GetCaCrt:output_type -> crs.CaResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_csr_proto_init() }
func file_csr_proto_init() {
	if File_csr_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_csr_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsrRequest); i {
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
		file_csr_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CsrResponse); i {
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
		file_csr_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaRequest); i {
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
		file_csr_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CaResponse); i {
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
			RawDescriptor: file_csr_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_csr_proto_goTypes,
		DependencyIndexes: file_csr_proto_depIdxs,
		MessageInfos:      file_csr_proto_msgTypes,
	}.Build()
	File_csr_proto = out.File
	file_csr_proto_rawDesc = nil
	file_csr_proto_goTypes = nil
	file_csr_proto_depIdxs = nil
}
