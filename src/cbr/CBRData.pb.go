// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: CBRData.proto

package cbr

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

type CBRData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayFile []*CBRData_ReplayFile `protobuf:"bytes,1,rep,name=replayFile,proto3" json:"replayFile,omitempty"`
}

func (x *CBRData) Reset() {
	*x = CBRData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRData_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRData) ProtoMessage() {}

func (x *CBRData) ProtoReflect() protoreflect.Message {
	mi := &file_CBRData_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRData.ProtoReflect.Descriptor instead.
func (*CBRData) Descriptor() ([]byte, []int) {
	return file_CBRData_proto_rawDescGZIP(), []int{0}
}

func (x *CBRData) GetReplayFile() []*CBRData_ReplayFile {
	if x != nil {
		return x.ReplayFile
	}
	return nil
}

type CBRData_ReplayFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame []*CBRData_Frame `protobuf:"bytes,1,rep,name=frame,proto3" json:"frame,omitempty"`
	Case  []*CBRData_Case  `protobuf:"bytes,2,rep,name=case,proto3" json:"case,omitempty"`
}

func (x *CBRData_ReplayFile) Reset() {
	*x = CBRData_ReplayFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRData_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRData_ReplayFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRData_ReplayFile) ProtoMessage() {}

func (x *CBRData_ReplayFile) ProtoReflect() protoreflect.Message {
	mi := &file_CBRData_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRData_ReplayFile.ProtoReflect.Descriptor instead.
func (*CBRData_ReplayFile) Descriptor() ([]byte, []int) {
	return file_CBRData_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CBRData_ReplayFile) GetFrame() []*CBRData_Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

func (x *CBRData_ReplayFile) GetCase() []*CBRData_Case {
	if x != nil {
		return x.Case
	}
	return nil
}

type CBRData_Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  int32 `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	Facing bool  `protobuf:"varint,2,opt,name=facing,proto3" json:"facing,omitempty"`
}

func (x *CBRData_Frame) Reset() {
	*x = CBRData_Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRData_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRData_Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRData_Frame) ProtoMessage() {}

func (x *CBRData_Frame) ProtoReflect() protoreflect.Message {
	mi := &file_CBRData_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRData_Frame.ProtoReflect.Descriptor instead.
func (*CBRData_Frame) Descriptor() ([]byte, []int) {
	return file_CBRData_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CBRData_Frame) GetInput() int32 {
	if x != nil {
		return x.Input
	}
	return 0
}

func (x *CBRData_Frame) GetFacing() bool {
	if x != nil {
		return x.Facing
	}
	return false
}

type CBRData_Case struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrameStartId int64 `protobuf:"varint,1,opt,name=frameStartId,proto3" json:"frameStartId,omitempty"`
	FrameEndId   int64 `protobuf:"varint,2,opt,name=frameEndId,proto3" json:"frameEndId,omitempty"`
}

func (x *CBRData_Case) Reset() {
	*x = CBRData_Case{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRData_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRData_Case) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRData_Case) ProtoMessage() {}

func (x *CBRData_Case) ProtoReflect() protoreflect.Message {
	mi := &file_CBRData_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRData_Case.ProtoReflect.Descriptor instead.
func (*CBRData_Case) Descriptor() ([]byte, []int) {
	return file_CBRData_proto_rawDescGZIP(), []int{0, 2}
}

func (x *CBRData_Case) GetFrameStartId() int64 {
	if x != nil {
		return x.FrameStartId
	}
	return 0
}

func (x *CBRData_Case) GetFrameEndId() int64 {
	if x != nil {
		return x.FrameEndId
	}
	return 0
}

var File_CBRData_proto protoreflect.FileDescriptor

var file_CBRData_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x43, 0x42, 0x52, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0xa7, 0x02, 0x0a, 0x07, 0x43, 0x42, 0x52, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x42, 0x52,
	0x44, 0x61, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x0a, 0x72, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x5f, 0x0a, 0x0a, 0x52,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x66, 0x72, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x43, 0x42, 0x52, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x05, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x63, 0x61, 0x73, 0x65, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x42, 0x52, 0x44, 0x61, 0x74,
	0x61, 0x2e, 0x43, 0x61, 0x73, 0x65, 0x52, 0x04, 0x63, 0x61, 0x73, 0x65, 0x1a, 0x35, 0x0a, 0x05,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x1a, 0x4a, 0x0a, 0x04, 0x43, 0x61, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x66,
	0x72, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x49, 0x64, 0x42,
	0x06, 0x5a, 0x04, 0x2f, 0x63, 0x62, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CBRData_proto_rawDescOnce sync.Once
	file_CBRData_proto_rawDescData = file_CBRData_proto_rawDesc
)

func file_CBRData_proto_rawDescGZIP() []byte {
	file_CBRData_proto_rawDescOnce.Do(func() {
		file_CBRData_proto_rawDescData = protoimpl.X.CompressGZIP(file_CBRData_proto_rawDescData)
	})
	return file_CBRData_proto_rawDescData
}

var file_CBRData_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_CBRData_proto_goTypes = []interface{}{
	(*CBRData)(nil),            // 0: main.CBRData
	(*CBRData_ReplayFile)(nil), // 1: main.CBRData.ReplayFile
	(*CBRData_Frame)(nil),      // 2: main.CBRData.Frame
	(*CBRData_Case)(nil),       // 3: main.CBRData.Case
}
var file_CBRData_proto_depIdxs = []int32{
	1, // 0: main.CBRData.replayFile:type_name -> main.CBRData.ReplayFile
	2, // 1: main.CBRData.ReplayFile.frame:type_name -> main.CBRData.Frame
	3, // 2: main.CBRData.ReplayFile.case:type_name -> main.CBRData.Case
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CBRData_proto_init() }
func file_CBRData_proto_init() {
	if File_CBRData_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CBRData_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRData); i {
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
		file_CBRData_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRData_ReplayFile); i {
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
		file_CBRData_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRData_Frame); i {
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
		file_CBRData_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRData_Case); i {
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
			RawDescriptor: file_CBRData_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CBRData_proto_goTypes,
		DependencyIndexes: file_CBRData_proto_depIdxs,
		MessageInfos:      file_CBRData_proto_msgTypes,
	}.Build()
	File_CBRData_proto = out.File
	file_CBRData_proto_rawDesc = nil
	file_CBRData_proto_goTypes = nil
	file_CBRData_proto_depIdxs = nil
}
