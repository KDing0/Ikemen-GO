// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.2
// source: CBRRawFrame.proto

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

type CBRRawFrames struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReplayFile []*CBRRawFrames_ReplayFile `protobuf:"bytes,1,rep,name=replayFile,proto3" json:"replayFile,omitempty"`
}

func (x *CBRRawFrames) Reset() {
	*x = CBRRawFrames{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRRawFrame_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRRawFrames) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRRawFrames) ProtoMessage() {}

func (x *CBRRawFrames) ProtoReflect() protoreflect.Message {
	mi := &file_CBRRawFrame_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRRawFrames.ProtoReflect.Descriptor instead.
func (*CBRRawFrames) Descriptor() ([]byte, []int) {
	return file_CBRRawFrame_proto_rawDescGZIP(), []int{0}
}

func (x *CBRRawFrames) GetReplayFile() []*CBRRawFrames_ReplayFile {
	if x != nil {
		return x.ReplayFile
	}
	return nil
}

type CBRRawFrames_ReplayFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frame []*CBRRawFrames_Frame `protobuf:"bytes,1,rep,name=frame,proto3" json:"frame,omitempty"`
}

func (x *CBRRawFrames_ReplayFile) Reset() {
	*x = CBRRawFrames_ReplayFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRRawFrame_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRRawFrames_ReplayFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRRawFrames_ReplayFile) ProtoMessage() {}

func (x *CBRRawFrames_ReplayFile) ProtoReflect() protoreflect.Message {
	mi := &file_CBRRawFrame_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRRawFrames_ReplayFile.ProtoReflect.Descriptor instead.
func (*CBRRawFrames_ReplayFile) Descriptor() ([]byte, []int) {
	return file_CBRRawFrame_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CBRRawFrames_ReplayFile) GetFrame() []*CBRRawFrames_Frame {
	if x != nil {
		return x.Frame
	}
	return nil
}

type CBRRawFrames_Frame struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CharData []*CBRRawFrames_CharData `protobuf:"bytes,1,rep,name=charData,proto3" json:"charData,omitempty"`
}

func (x *CBRRawFrames_Frame) Reset() {
	*x = CBRRawFrames_Frame{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRRawFrame_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRRawFrames_Frame) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRRawFrames_Frame) ProtoMessage() {}

func (x *CBRRawFrames_Frame) ProtoReflect() protoreflect.Message {
	mi := &file_CBRRawFrame_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRRawFrames_Frame.ProtoReflect.Descriptor instead.
func (*CBRRawFrames_Frame) Descriptor() ([]byte, []int) {
	return file_CBRRawFrame_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CBRRawFrames_Frame) GetCharData() []*CBRRawFrames_CharData {
	if x != nil {
		return x.CharData
	}
	return nil
}

type CBRRawFrames_CharData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  int32 `protobuf:"varint,1,opt,name=input,proto3" json:"input,omitempty"`
	Facing bool  `protobuf:"varint,2,opt,name=facing,proto3" json:"facing,omitempty"`
}

func (x *CBRRawFrames_CharData) Reset() {
	*x = CBRRawFrames_CharData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CBRRawFrame_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CBRRawFrames_CharData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CBRRawFrames_CharData) ProtoMessage() {}

func (x *CBRRawFrames_CharData) ProtoReflect() protoreflect.Message {
	mi := &file_CBRRawFrame_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CBRRawFrames_CharData.ProtoReflect.Descriptor instead.
func (*CBRRawFrames_CharData) Descriptor() ([]byte, []int) {
	return file_CBRRawFrame_proto_rawDescGZIP(), []int{0, 2}
}

func (x *CBRRawFrames_CharData) GetInput() int32 {
	if x != nil {
		return x.Input
	}
	return 0
}

func (x *CBRRawFrames_CharData) GetFacing() bool {
	if x != nil {
		return x.Facing
	}
	return false
}

var File_CBRRawFrame_proto protoreflect.FileDescriptor

var file_CBRRawFrame_proto_rawDesc = []byte{
	0x0a, 0x11, 0x43, 0x42, 0x52, 0x52, 0x61, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x87, 0x02, 0x0a, 0x0c, 0x43, 0x42,
	0x52, 0x52, 0x61, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0a, 0x72, 0x65,
	0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x42, 0x52, 0x52, 0x61, 0x77, 0x46, 0x72, 0x61, 0x6d,
	0x65, 0x73, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0a, 0x72,
	0x65, 0x70, 0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x1a, 0x3c, 0x0a, 0x0a, 0x52, 0x65, 0x70,
	0x6c, 0x61, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x42,
	0x52, 0x52, 0x61, 0x77, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x52, 0x05, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x1a, 0x40, 0x0a, 0x05, 0x46, 0x72, 0x61, 0x6d, 0x65,
	0x12, 0x37, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x43, 0x42, 0x52, 0x52, 0x61, 0x77,
	0x46, 0x72, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x08, 0x63, 0x68, 0x61, 0x72, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x38, 0x0a, 0x08, 0x43, 0x68, 0x61,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66,
	0x61, 0x63, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x66, 0x61, 0x63,
	0x69, 0x6e, 0x67, 0x42, 0x06, 0x5a, 0x04, 0x2f, 0x63, 0x62, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_CBRRawFrame_proto_rawDescOnce sync.Once
	file_CBRRawFrame_proto_rawDescData = file_CBRRawFrame_proto_rawDesc
)

func file_CBRRawFrame_proto_rawDescGZIP() []byte {
	file_CBRRawFrame_proto_rawDescOnce.Do(func() {
		file_CBRRawFrame_proto_rawDescData = protoimpl.X.CompressGZIP(file_CBRRawFrame_proto_rawDescData)
	})
	return file_CBRRawFrame_proto_rawDescData
}

var file_CBRRawFrame_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_CBRRawFrame_proto_goTypes = []interface{}{
	(*CBRRawFrames)(nil),            // 0: main.CBRRawFrames
	(*CBRRawFrames_ReplayFile)(nil), // 1: main.CBRRawFrames.ReplayFile
	(*CBRRawFrames_Frame)(nil),      // 2: main.CBRRawFrames.Frame
	(*CBRRawFrames_CharData)(nil),   // 3: main.CBRRawFrames.CharData
}
var file_CBRRawFrame_proto_depIdxs = []int32{
	1, // 0: main.CBRRawFrames.replayFile:type_name -> main.CBRRawFrames.ReplayFile
	2, // 1: main.CBRRawFrames.ReplayFile.frame:type_name -> main.CBRRawFrames.Frame
	3, // 2: main.CBRRawFrames.Frame.charData:type_name -> main.CBRRawFrames.CharData
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_CBRRawFrame_proto_init() }
func file_CBRRawFrame_proto_init() {
	if File_CBRRawFrame_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_CBRRawFrame_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRRawFrames); i {
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
		file_CBRRawFrame_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRRawFrames_ReplayFile); i {
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
		file_CBRRawFrame_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRRawFrames_Frame); i {
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
		file_CBRRawFrame_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CBRRawFrames_CharData); i {
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
			RawDescriptor: file_CBRRawFrame_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CBRRawFrame_proto_goTypes,
		DependencyIndexes: file_CBRRawFrame_proto_depIdxs,
		MessageInfos:      file_CBRRawFrame_proto_msgTypes,
	}.Build()
	File_CBRRawFrame_proto = out.File
	file_CBRRawFrame_proto_rawDesc = nil
	file_CBRRawFrame_proto_goTypes = nil
	file_CBRRawFrame_proto_depIdxs = nil
}
