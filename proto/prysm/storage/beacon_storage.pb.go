// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/prysm/storage/beacon_storage.proto

package storage

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

type BeaconStateForStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to StorageState:
	//	*BeaconStateForStorage_StateV1
	StorageState isBeaconStateForStorage_StorageState `protobuf_oneof:"storageState"`
}

func (x *BeaconStateForStorage) Reset() {
	*x = BeaconStateForStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_storage_beacon_storage_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BeaconStateForStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BeaconStateForStorage) ProtoMessage() {}

func (x *BeaconStateForStorage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_storage_beacon_storage_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BeaconStateForStorage.ProtoReflect.Descriptor instead.
func (*BeaconStateForStorage) Descriptor() ([]byte, []int) {
	return file_proto_prysm_storage_beacon_storage_proto_rawDescGZIP(), []int{0}
}

func (m *BeaconStateForStorage) GetStorageState() isBeaconStateForStorage_StorageState {
	if m != nil {
		return m.StorageState
	}
	return nil
}

func (x *BeaconStateForStorage) GetStateV1() *StorageBeaconStateV1 {
	if x, ok := x.GetStorageState().(*BeaconStateForStorage_StateV1); ok {
		return x.StateV1
	}
	return nil
}

type isBeaconStateForStorage_StorageState interface {
	isBeaconStateForStorage_StorageState()
}

type BeaconStateForStorage_StateV1 struct {
	StateV1 *StorageBeaconStateV1 `protobuf:"bytes,1,opt,name=state_v1,json=stateV1,proto3,oneof"`
}

func (*BeaconStateForStorage_StateV1) isBeaconStateForStorage_StorageState() {}

type StorageBeaconStateV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version          Version `protobuf:"varint,1001,opt,name=version,proto3,enum=ethereum.eth.storage.Version" json:"version,omitempty"`
	ValidatorIndexes []byte  `protobuf:"bytes,1002,opt,name=validatorIndexes,proto3" json:"validatorIndexes,omitempty"`
	State            []byte  `protobuf:"bytes,2000,opt,name=state,proto3" json:"state,omitempty"`
}

func (x *StorageBeaconStateV1) Reset() {
	*x = StorageBeaconStateV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_prysm_storage_beacon_storage_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorageBeaconStateV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorageBeaconStateV1) ProtoMessage() {}

func (x *StorageBeaconStateV1) ProtoReflect() protoreflect.Message {
	mi := &file_proto_prysm_storage_beacon_storage_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorageBeaconStateV1.ProtoReflect.Descriptor instead.
func (*StorageBeaconStateV1) Descriptor() ([]byte, []int) {
	return file_proto_prysm_storage_beacon_storage_proto_rawDescGZIP(), []int{1}
}

func (x *StorageBeaconStateV1) GetVersion() Version {
	if x != nil {
		return x.Version
	}
	return Version_UNKNOWN
}

func (x *StorageBeaconStateV1) GetValidatorIndexes() []byte {
	if x != nil {
		return x.ValidatorIndexes
	}
	return nil
}

func (x *StorageBeaconStateV1) GetState() []byte {
	if x != nil {
		return x.State
	}
	return nil
}

var File_proto_prysm_storage_beacon_storage_proto protoreflect.FileDescriptor

var file_proto_prysm_storage_beacon_storage_proto_rawDesc = []byte{
	0x0a, 0x28, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x62, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x74, 0x68, 0x65,
	0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x1a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x79, 0x73, 0x6d, 0x2f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x15, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x46, 0x6f, 0x72, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x47, 0x0a, 0x08,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a,
	0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x42, 0x65, 0x61,
	0x63, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56, 0x31, 0x48, 0x00, 0x52, 0x07, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x56, 0x31, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x14, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x42, 0x65, 0x61, 0x63, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x56, 0x31, 0x12, 0x38,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1d, 0x2e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x2e, 0x65, 0x74, 0x68, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x10, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x73, 0x18, 0xea, 0x07, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x10, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0xd0,
	0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_prysm_storage_beacon_storage_proto_rawDescOnce sync.Once
	file_proto_prysm_storage_beacon_storage_proto_rawDescData = file_proto_prysm_storage_beacon_storage_proto_rawDesc
)

func file_proto_prysm_storage_beacon_storage_proto_rawDescGZIP() []byte {
	file_proto_prysm_storage_beacon_storage_proto_rawDescOnce.Do(func() {
		file_proto_prysm_storage_beacon_storage_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_prysm_storage_beacon_storage_proto_rawDescData)
	})
	return file_proto_prysm_storage_beacon_storage_proto_rawDescData
}

var file_proto_prysm_storage_beacon_storage_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_prysm_storage_beacon_storage_proto_goTypes = []interface{}{
	(*BeaconStateForStorage)(nil), // 0: ethereum.eth.storage.BeaconStateForStorage
	(*StorageBeaconStateV1)(nil),  // 1: ethereum.eth.storage.StorageBeaconStateV1
	(Version)(0),                  // 2: ethereum.eth.storage.Version
}
var file_proto_prysm_storage_beacon_storage_proto_depIdxs = []int32{
	1, // 0: ethereum.eth.storage.BeaconStateForStorage.state_v1:type_name -> ethereum.eth.storage.StorageBeaconStateV1
	2, // 1: ethereum.eth.storage.StorageBeaconStateV1.version:type_name -> ethereum.eth.storage.Version
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_prysm_storage_beacon_storage_proto_init() }
func file_proto_prysm_storage_beacon_storage_proto_init() {
	if File_proto_prysm_storage_beacon_storage_proto != nil {
		return
	}
	file_proto_prysm_storage_version_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_prysm_storage_beacon_storage_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BeaconStateForStorage); i {
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
		file_proto_prysm_storage_beacon_storage_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StorageBeaconStateV1); i {
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
	file_proto_prysm_storage_beacon_storage_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*BeaconStateForStorage_StateV1)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_prysm_storage_beacon_storage_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_prysm_storage_beacon_storage_proto_goTypes,
		DependencyIndexes: file_proto_prysm_storage_beacon_storage_proto_depIdxs,
		MessageInfos:      file_proto_prysm_storage_beacon_storage_proto_msgTypes,
	}.Build()
	File_proto_prysm_storage_beacon_storage_proto = out.File
	file_proto_prysm_storage_beacon_storage_proto_rawDesc = nil
	file_proto_prysm_storage_beacon_storage_proto_goTypes = nil
	file_proto_prysm_storage_beacon_storage_proto_depIdxs = nil
}
