// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: protos/distributed_grep.proto

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
type ClientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The search string of the grep command.
	QueryString string `protobuf:"bytes,1,opt,name=queryString,proto3" json:"queryString,omitempty"`
	// The search options of the grep command.
	Args string `protobuf:"bytes,2,opt,name=args,proto3" json:"args,omitempty"`
	// The file to be searched.
	QueryFile string `protobuf:"bytes,3,opt,name=queryFile,proto3" json:"queryFile,omitempty"`
}

func (x *ClientRequest) Reset() {
	*x = ClientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientRequest) ProtoMessage() {}

func (x *ClientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientRequest.ProtoReflect.Descriptor instead.
func (*ClientRequest) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{0}
}

func (x *ClientRequest) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

func (x *ClientRequest) GetArgs() string {
	if x != nil {
		return x.Args
	}
	return ""
}

func (x *ClientRequest) GetQueryFile() string {
	if x != nil {
		return x.QueryFile
	}
	return ""
}

// The response of the distributed grep command.
type CoordinatorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//array with responses of the grep command on the remote workers.
	Result []string `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
	//cumulative of lines matched.
	Total int64 `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *CoordinatorResponse) Reset() {
	*x = CoordinatorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CoordinatorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CoordinatorResponse) ProtoMessage() {}

func (x *CoordinatorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CoordinatorResponse.ProtoReflect.Descriptor instead.
func (*CoordinatorResponse) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{1}
}

func (x *CoordinatorResponse) GetResult() []string {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *CoordinatorResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type CreateLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLogsRequest) Reset() {
	*x = CreateLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogsRequest) ProtoMessage() {}

func (x *CreateLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogsRequest.ProtoReflect.Descriptor instead.
func (*CreateLogsRequest) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{2}
}

type CreateLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateLogsResponse) Reset() {
	*x = CreateLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLogsResponse) ProtoMessage() {}

func (x *CreateLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLogsResponse.ProtoReflect.Descriptor instead.
func (*CreateLogsResponse) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{3}
}

type DeleteLogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLogsRequest) Reset() {
	*x = DeleteLogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogsRequest) ProtoMessage() {}

func (x *DeleteLogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogsRequest.ProtoReflect.Descriptor instead.
func (*DeleteLogsRequest) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{4}
}

type DeleteLogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLogsResponse) Reset() {
	*x = DeleteLogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_distributed_grep_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLogsResponse) ProtoMessage() {}

func (x *DeleteLogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_distributed_grep_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLogsResponse.ProtoReflect.Descriptor instead.
func (*DeleteLogsResponse) Descriptor() ([]byte, []int) {
	return file_protos_distributed_grep_proto_rawDescGZIP(), []int{5}
}

var File_protos_distributed_grep_proto protoreflect.FileDescriptor

var file_protos_distributed_grep_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x64, 0x5f, 0x67, 0x72, 0x65, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x63, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x22, 0x43, 0x0a, 0x13, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x13, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x13, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f,
	0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xbf, 0x01, 0x0a, 0x0f, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64, 0x47,
	0x72, 0x65, 0x70, 0x12, 0x32, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x47, 0x72, 0x65, 0x70, 0x12,
	0x0e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x43, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x6f, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0e, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c,
	0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x15, 0x5a, 0x13, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x64,
	0x5f, 0x67, 0x72, 0x65, 0x70, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_distributed_grep_proto_rawDescOnce sync.Once
	file_protos_distributed_grep_proto_rawDescData = file_protos_distributed_grep_proto_rawDesc
)

func file_protos_distributed_grep_proto_rawDescGZIP() []byte {
	file_protos_distributed_grep_proto_rawDescOnce.Do(func() {
		file_protos_distributed_grep_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_distributed_grep_proto_rawDescData)
	})
	return file_protos_distributed_grep_proto_rawDescData
}

var file_protos_distributed_grep_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_distributed_grep_proto_goTypes = []interface{}{
	(*ClientRequest)(nil),       // 0: ClientRequest
	(*CoordinatorResponse)(nil), // 1: CoordinatorResponse
	(*CreateLogsRequest)(nil),   // 2: CreateLogsRequest
	(*CreateLogsResponse)(nil),  // 3: CreateLogsResponse
	(*DeleteLogsRequest)(nil),   // 4: DeleteLogsRequest
	(*DeleteLogsResponse)(nil),  // 5: DeleteLogsResponse
}
var file_protos_distributed_grep_proto_depIdxs = []int32{
	0, // 0: DistributedGrep.CallGrep:input_type -> ClientRequest
	2, // 1: DistributedGrep.CallCreateLogs:input_type -> CreateLogsRequest
	4, // 2: DistributedGrep.CallDeleteLogs:input_type -> DeleteLogsRequest
	1, // 3: DistributedGrep.CallGrep:output_type -> CoordinatorResponse
	3, // 4: DistributedGrep.CallCreateLogs:output_type -> CreateLogsResponse
	5, // 5: DistributedGrep.CallDeleteLogs:output_type -> DeleteLogsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_distributed_grep_proto_init() }
func file_protos_distributed_grep_proto_init() {
	if File_protos_distributed_grep_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_distributed_grep_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientRequest); i {
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
		file_protos_distributed_grep_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CoordinatorResponse); i {
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
		file_protos_distributed_grep_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogsRequest); i {
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
		file_protos_distributed_grep_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLogsResponse); i {
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
		file_protos_distributed_grep_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogsRequest); i {
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
		file_protos_distributed_grep_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLogsResponse); i {
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
			RawDescriptor: file_protos_distributed_grep_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_distributed_grep_proto_goTypes,
		DependencyIndexes: file_protos_distributed_grep_proto_depIdxs,
		MessageInfos:      file_protos_distributed_grep_proto_msgTypes,
	}.Build()
	File_protos_distributed_grep_proto = out.File
	file_protos_distributed_grep_proto_rawDesc = nil
	file_protos_distributed_grep_proto_goTypes = nil
	file_protos_distributed_grep_proto_depIdxs = nil
}
