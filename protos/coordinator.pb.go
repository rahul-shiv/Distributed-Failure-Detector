// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protos/coordinator.proto

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

// The request of the grep command on a remote machine.
type GrepRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The search string of the grep command.
	SearchString string `protobuf:"bytes,1,opt,name=searchString,proto3" json:"searchString,omitempty"`
	// The search options of the grep command.
	SearchOptions string `protobuf:"bytes,2,opt,name=searchOptions,proto3" json:"searchOptions,omitempty"`
	// The file to be searched.
	SearchFile string `protobuf:"bytes,3,opt,name=searchFile,proto3" json:"searchFile,omitempty"`
}

func (x *GrepRequest) Reset() {
	*x = GrepRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrepRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrepRequest) ProtoMessage() {}

func (x *GrepRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrepRequest.ProtoReflect.Descriptor instead.
func (*GrepRequest) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{0}
}

func (x *GrepRequest) GetSearchString() string {
	if x != nil {
		return x.SearchString
	}
	return ""
}

func (x *GrepRequest) GetSearchOptions() string {
	if x != nil {
		return x.SearchOptions
	}
	return ""
}

func (x *GrepRequest) GetSearchFile() string {
	if x != nil {
		return x.SearchFile
	}
	return ""
}

// The response of the grep command on a remote machine.
type GrepResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Single line string output of the grep command.
	GrepOutput string `protobuf:"bytes,1,opt,name=grepOutput,proto3" json:"grepOutput,omitempty"`
}

func (x *GrepResponse) Reset() {
	*x = GrepResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GrepResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GrepResponse) ProtoMessage() {}

func (x *GrepResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GrepResponse.ProtoReflect.Descriptor instead.
func (*GrepResponse) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{1}
}

func (x *GrepResponse) GetGrepOutput() string {
	if x != nil {
		return x.GrepOutput
	}
	return ""
}

type WorkerCreateLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// VM number in which the log is getting created
	VmIndex int32 `protobuf:"varint,1,opt,name=vmIndex,proto3" json:"vmIndex,omitempty"`
}

func (x *WorkerCreateLogsRequest) Reset() {
	*x = WorkerCreateLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerCreateLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerCreateLogsRequest) ProtoMessage() {}

func (x *WorkerCreateLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerCreateLogsRequest.ProtoReflect.Descriptor instead.
func (*WorkerCreateLogsRequest) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{2}
}

func (x *WorkerCreateLogsRequest) GetVmIndex() int32 {
	if x != nil {
		return x.VmIndex
	}
	return 0
}

type WorkerCreateLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkerCreateLogsResponse) Reset() {
	*x = WorkerCreateLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerCreateLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerCreateLogsResponse) ProtoMessage() {}

func (x *WorkerCreateLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerCreateLogsResponse.ProtoReflect.Descriptor instead.
func (*WorkerCreateLogsResponse) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{3}
}

type WorkerDeleteLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkerDeleteLogsRequest) Reset() {
	*x = WorkerDeleteLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDeleteLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeleteLogsRequest) ProtoMessage() {}

func (x *WorkerDeleteLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeleteLogsRequest.ProtoReflect.Descriptor instead.
func (*WorkerDeleteLogsRequest) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{4}
}

type WorkerDeleteLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WorkerDeleteLogsResponse) Reset() {
	*x = WorkerDeleteLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_coordinator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkerDeleteLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeleteLogsResponse) ProtoMessage() {}

func (x *WorkerDeleteLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_coordinator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeleteLogsResponse.ProtoReflect.Descriptor instead.
func (*WorkerDeleteLogsResponse) Descriptor() ([]byte, []int) {
	return file_protos_coordinator_proto_rawDescGZIP(), []int{5}
}

var File_protos_coordinator_proto protoreflect.FileDescriptor

var file_protos_coordinator_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x67, 0x72, 0x65, 0x70,
	0x22, 0x77, 0x0a, 0x0b, 0x47, 0x72, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x22, 0x0a, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x69, 0x6c, 0x65, 0x22, 0x2e, 0x0a, 0x0c, 0x47, 0x72, 0x65,
	0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x67, 0x72, 0x65,
	0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x67,
	0x72, 0x65, 0x70, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x33, 0x0a, 0x17, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x6d, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x1a,
	0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19, 0x0a, 0x17, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1a, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x32, 0xf2, 0x01, 0x0a, 0x0d, 0x47, 0x72, 0x65, 0x70, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x66,
	0x61, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x47, 0x72, 0x65, 0x70, 0x12, 0x11, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x2e, 0x47, 0x72, 0x65, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12,
	0x2e, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x47, 0x72, 0x65, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a, 0x14, 0x43, 0x61, 0x6c, 0x6c, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x72, 0x65, 0x70, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72,
	0x65, 0x70, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x57, 0x0a,
	0x14, 0x43, 0x61, 0x6c, 0x6c, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x65, 0x72, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_coordinator_proto_rawDescOnce sync.Once
	file_protos_coordinator_proto_rawDescData = file_protos_coordinator_proto_rawDesc
)

func file_protos_coordinator_proto_rawDescGZIP() []byte {
	file_protos_coordinator_proto_rawDescOnce.Do(func() {
		file_protos_coordinator_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_coordinator_proto_rawDescData)
	})
	return file_protos_coordinator_proto_rawDescData
}

var file_protos_coordinator_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_coordinator_proto_goTypes = []interface{}{
	(*GrepRequest)(nil),              // 0: grep.GrepRequest
	(*GrepResponse)(nil),             // 1: grep.GrepResponse
	(*WorkerCreateLogsRequest)(nil),  // 2: grep.WorkerCreateLogsRequest
	(*WorkerCreateLogsResponse)(nil), // 3: grep.WorkerCreateLogsResponse
	(*WorkerDeleteLogsRequest)(nil),  // 4: grep.WorkerDeleteLogsRequest
	(*WorkerDeleteLogsResponse)(nil), // 5: grep.WorkerDeleteLogsResponse
}
var file_protos_coordinator_proto_depIdxs = []int32{
	0, // 0: grep.GrepInterface.Grep:input_type -> grep.GrepRequest
	2, // 1: grep.GrepInterface.CallWorkerCreateLogs:input_type -> grep.WorkerCreateLogsRequest
	4, // 2: grep.GrepInterface.CallWorkerDeleteLogs:input_type -> grep.WorkerDeleteLogsRequest
	1, // 3: grep.GrepInterface.Grep:output_type -> grep.GrepResponse
	3, // 4: grep.GrepInterface.CallWorkerCreateLogs:output_type -> grep.WorkerCreateLogsResponse
	5, // 5: grep.GrepInterface.CallWorkerDeleteLogs:output_type -> grep.WorkerDeleteLogsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_coordinator_proto_init() }
func file_protos_coordinator_proto_init() {
	if File_protos_coordinator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_coordinator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrepRequest); i {
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
		file_protos_coordinator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GrepResponse); i {
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
		file_protos_coordinator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerCreateLogsRequest); i {
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
		file_protos_coordinator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerCreateLogsResponse); i {
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
		file_protos_coordinator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDeleteLogsRequest); i {
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
		file_protos_coordinator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkerDeleteLogsResponse); i {
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
			RawDescriptor: file_protos_coordinator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_coordinator_proto_goTypes,
		DependencyIndexes: file_protos_coordinator_proto_depIdxs,
		MessageInfos:      file_protos_coordinator_proto_msgTypes,
	}.Build()
	File_protos_coordinator_proto = out.File
	file_protos_coordinator_proto_rawDesc = nil
	file_protos_coordinator_proto_goTypes = nil
	file_protos_coordinator_proto_depIdxs = nil
}
