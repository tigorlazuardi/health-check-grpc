// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.16.0
// source: healthcheck.proto

package hcproto

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

type SubPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SubPayload) Reset() {
	*x = SubPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubPayload) ProtoMessage() {}

func (x *SubPayload) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubPayload.ProtoReflect.Descriptor instead.
func (*SubPayload) Descriptor() ([]byte, []int) {
	return file_healthcheck_proto_rawDescGZIP(), []int{0}
}

func (x *SubPayload) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_healthcheck_proto_rawDescGZIP(), []int{1}
}

func (x *Ack) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type HealthCheckMap struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entity map[string]*EntityStatus `protobuf:"bytes,1,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *HealthCheckMap) Reset() {
	*x = HealthCheckMap{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckMap) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckMap) ProtoMessage() {}

func (x *HealthCheckMap) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckMap.ProtoReflect.Descriptor instead.
func (*HealthCheckMap) Descriptor() ([]byte, []int) {
	return file_healthcheck_proto_rawDescGZIP(), []int{2}
}

func (x *HealthCheckMap) GetEntity() map[string]*EntityStatus {
	if x != nil {
		return x.Entity
	}
	return nil
}

type EntityStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Message string `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EntityStatus) Reset() {
	*x = EntityStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_healthcheck_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityStatus) ProtoMessage() {}

func (x *EntityStatus) ProtoReflect() protoreflect.Message {
	mi := &file_healthcheck_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityStatus.ProtoReflect.Descriptor instead.
func (*EntityStatus) Descriptor() ([]byte, []int) {
	return file_healthcheck_proto_rawDescGZIP(), []int{3}
}

func (x *EntityStatus) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EntityStatus) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EntityStatus) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_healthcheck_proto protoreflect.FileDescriptor

var file_healthcheck_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x22, 0x1c, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f,
	0x0a, 0x03, 0x41, 0x63, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0xa7, 0x01, 0x0a, 0x0e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d,
	0x61, 0x70, 0x12, 0x3f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b,
	0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x61, 0x70, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x1a, 0x54, 0x0a, 0x0b, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x2f, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x54, 0x0a, 0x0c, 0x45, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0x90, 0x01, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12,
	0x45, 0x0a, 0x09, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1b, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x4d,
	0x61, 0x70, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x17, 0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68,
	0x65, 0x63, 0x6b, 0x2e, 0x53, 0x75, 0x62, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x10,
	0x2e, 0x68, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x41, 0x63, 0x6b,
	0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x74, 0x69, 0x67, 0x6f, 0x72, 0x6c, 0x61, 0x7a, 0x75, 0x61, 0x72, 0x64, 0x69, 0x2f, 0x68,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x2d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2d, 0x67, 0x72, 0x70, 0x63,
	0x3b, 0x68, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_healthcheck_proto_rawDescOnce sync.Once
	file_healthcheck_proto_rawDescData = file_healthcheck_proto_rawDesc
)

func file_healthcheck_proto_rawDescGZIP() []byte {
	file_healthcheck_proto_rawDescOnce.Do(func() {
		file_healthcheck_proto_rawDescData = protoimpl.X.CompressGZIP(file_healthcheck_proto_rawDescData)
	})
	return file_healthcheck_proto_rawDescData
}

var file_healthcheck_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_healthcheck_proto_goTypes = []interface{}{
	(*SubPayload)(nil),     // 0: healthcheck.SubPayload
	(*Ack)(nil),            // 1: healthcheck.Ack
	(*HealthCheckMap)(nil), // 2: healthcheck.HealthCheckMap
	(*EntityStatus)(nil),   // 3: healthcheck.EntityStatus
	nil,                    // 4: healthcheck.HealthCheckMap.EntityEntry
}
var file_healthcheck_proto_depIdxs = []int32{
	4, // 0: healthcheck.HealthCheckMap.entity:type_name -> healthcheck.HealthCheckMap.EntityEntry
	3, // 1: healthcheck.HealthCheckMap.EntityEntry.value:type_name -> healthcheck.EntityStatus
	0, // 2: healthcheck.HealthCheck.Subscribe:input_type -> healthcheck.SubPayload
	0, // 3: healthcheck.HealthCheck.Unsubscribe:input_type -> healthcheck.SubPayload
	2, // 4: healthcheck.HealthCheck.Subscribe:output_type -> healthcheck.HealthCheckMap
	1, // 5: healthcheck.HealthCheck.Unsubscribe:output_type -> healthcheck.Ack
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_healthcheck_proto_init() }
func file_healthcheck_proto_init() {
	if File_healthcheck_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_healthcheck_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubPayload); i {
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
		file_healthcheck_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_healthcheck_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckMap); i {
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
		file_healthcheck_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityStatus); i {
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
			RawDescriptor: file_healthcheck_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_healthcheck_proto_goTypes,
		DependencyIndexes: file_healthcheck_proto_depIdxs,
		MessageInfos:      file_healthcheck_proto_msgTypes,
	}.Build()
	File_healthcheck_proto = out.File
	file_healthcheck_proto_rawDesc = nil
	file_healthcheck_proto_goTypes = nil
	file_healthcheck_proto_depIdxs = nil
}
