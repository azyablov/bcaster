// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: protos/bcaster.proto

package bcaster

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

type MsgState int32

const (
	MsgState_ENQUEUED MsgState = 0
	MsgState_DEQUEUED MsgState = 1
)

// Enum value maps for MsgState.
var (
	MsgState_name = map[int32]string{
		0: "ENQUEUED",
		1: "DEQUEUED",
	}
	MsgState_value = map[string]int32{
		"ENQUEUED": 0,
		"DEQUEUED": 1,
	}
)

func (x MsgState) Enum() *MsgState {
	p := new(MsgState)
	*p = x
	return p
}

func (x MsgState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgState) Descriptor() protoreflect.EnumDescriptor {
	return file_protos_bcaster_proto_enumTypes[0].Descriptor()
}

func (MsgState) Type() protoreflect.EnumType {
	return &file_protos_bcaster_proto_enumTypes[0]
}

func (x MsgState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgState.Descriptor instead.
func (MsgState) EnumDescriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{0}
}

type ShortMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg   string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	State MsgState `protobuf:"varint,2,opt,name=state,proto3,enum=bcaster.MsgState" json:"state,omitempty"`
}

func (x *ShortMsg) Reset() {
	*x = ShortMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bcaster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortMsg) ProtoMessage() {}

func (x *ShortMsg) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bcaster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortMsg.ProtoReflect.Descriptor instead.
func (*ShortMsg) Descriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{0}
}

func (x *ShortMsg) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ShortMsg) GetState() MsgState {
	if x != nil {
		return x.State
	}
	return MsgState_ENQUEUED
}

type RegistrationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegistrationRequest) Reset() {
	*x = RegistrationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bcaster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationRequest) ProtoMessage() {}

func (x *RegistrationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bcaster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationRequest.ProtoReflect.Descriptor instead.
func (*RegistrationRequest) Descriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegistrationAccept struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Timeout int32  `protobuf:"varint,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *RegistrationAccept) Reset() {
	*x = RegistrationAccept{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bcaster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrationAccept) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrationAccept) ProtoMessage() {}

func (x *RegistrationAccept) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bcaster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrationAccept.ProtoReflect.Descriptor instead.
func (*RegistrationAccept) Descriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{2}
}

func (x *RegistrationAccept) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegistrationAccept) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

type DeliveryStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Enq bool `protobuf:"varint,1,opt,name=enq,proto3" json:"enq,omitempty"`
}

func (x *DeliveryStatus) Reset() {
	*x = DeliveryStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bcaster_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryStatus) ProtoMessage() {}

func (x *DeliveryStatus) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bcaster_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryStatus.ProtoReflect.Descriptor instead.
func (*DeliveryStatus) Descriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{3}
}

func (x *DeliveryStatus) GetEnq() bool {
	if x != nil {
		return x.Enq
	}
	return false
}

type DeliveryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeliveryRequest) Reset() {
	*x = DeliveryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_bcaster_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeliveryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeliveryRequest) ProtoMessage() {}

func (x *DeliveryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_bcaster_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeliveryRequest.ProtoReflect.Descriptor instead.
func (*DeliveryRequest) Descriptor() ([]byte, []int) {
	return file_protos_bcaster_proto_rawDescGZIP(), []int{4}
}

var File_protos_bcaster_proto protoreflect.FileDescriptor

var file_protos_bcaster_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x22,
	0x4b, 0x0a, 0x08, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x27, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x62,
	0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x06, 0x22, 0x29, 0x0a, 0x13,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x22, 0x22, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6e, 0x71, 0x22,
	0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2a, 0x26, 0x0a, 0x08, 0x4d, 0x73, 0x67, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0c,
	0x0a, 0x08, 0x45, 0x4e, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x44, 0x45, 0x51, 0x55, 0x45, 0x55, 0x45, 0x44, 0x10, 0x01, 0x32, 0xd5, 0x01, 0x0a, 0x07, 0x42,
	0x43, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x11, 0x2e, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x73, 0x67, 0x1a, 0x17, 0x2e, 0x62, 0x63, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x65, 0x63, 0x76, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x18, 0x2e, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11,
	0x2e, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x4d, 0x73,
	0x67, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x2e, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74,
	0x22, 0x00, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x61, 0x7a, 0x79, 0x61, 0x62, 0x6c, 0x6f, 0x76, 0x2f, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x63, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_bcaster_proto_rawDescOnce sync.Once
	file_protos_bcaster_proto_rawDescData = file_protos_bcaster_proto_rawDesc
)

func file_protos_bcaster_proto_rawDescGZIP() []byte {
	file_protos_bcaster_proto_rawDescOnce.Do(func() {
		file_protos_bcaster_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_bcaster_proto_rawDescData)
	})
	return file_protos_bcaster_proto_rawDescData
}

var file_protos_bcaster_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_protos_bcaster_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_bcaster_proto_goTypes = []interface{}{
	(MsgState)(0),               // 0: bcaster.MsgState
	(*ShortMsg)(nil),            // 1: bcaster.ShortMsg
	(*RegistrationRequest)(nil), // 2: bcaster.RegistrationRequest
	(*RegistrationAccept)(nil),  // 3: bcaster.RegistrationAccept
	(*DeliveryStatus)(nil),      // 4: bcaster.DeliveryStatus
	(*DeliveryRequest)(nil),     // 5: bcaster.DeliveryRequest
}
var file_protos_bcaster_proto_depIdxs = []int32{
	0, // 0: bcaster.ShortMsg.state:type_name -> bcaster.MsgState
	1, // 1: bcaster.BCaster.SendShortMsg:input_type -> bcaster.ShortMsg
	5, // 2: bcaster.BCaster.RecvShortMsg:input_type -> bcaster.DeliveryRequest
	2, // 3: bcaster.BCaster.RegisterClient:input_type -> bcaster.RegistrationRequest
	4, // 4: bcaster.BCaster.SendShortMsg:output_type -> bcaster.DeliveryStatus
	1, // 5: bcaster.BCaster.RecvShortMsg:output_type -> bcaster.ShortMsg
	3, // 6: bcaster.BCaster.RegisterClient:output_type -> bcaster.RegistrationAccept
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_bcaster_proto_init() }
func file_protos_bcaster_proto_init() {
	if File_protos_bcaster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_bcaster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortMsg); i {
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
		file_protos_bcaster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationRequest); i {
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
		file_protos_bcaster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrationAccept); i {
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
		file_protos_bcaster_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryStatus); i {
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
		file_protos_bcaster_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeliveryRequest); i {
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
			RawDescriptor: file_protos_bcaster_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_bcaster_proto_goTypes,
		DependencyIndexes: file_protos_bcaster_proto_depIdxs,
		EnumInfos:         file_protos_bcaster_proto_enumTypes,
		MessageInfos:      file_protos_bcaster_proto_msgTypes,
	}.Build()
	File_protos_bcaster_proto = out.File
	file_protos_bcaster_proto_rawDesc = nil
	file_protos_bcaster_proto_goTypes = nil
	file_protos_bcaster_proto_depIdxs = nil
}
