// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.15.8
// source: proto/testing/test.proto

package testing

import (
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Person_PhoneType int32

const (
	Person_MOBILE Person_PhoneType = 0
	Person_HOME   Person_PhoneType = 1
	Person_WORK   Person_PhoneType = 2
)

// Enum value maps for Person_PhoneType.
var (
	Person_PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	Person_PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x Person_PhoneType) Enum() *Person_PhoneType {
	p := new(Person_PhoneType)
	*p = x
	return p
}

func (x Person_PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Person_PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_testing_test_proto_enumTypes[0].Descriptor()
}

func (Person_PhoneType) Type() protoreflect.EnumType {
	return &file_proto_testing_test_proto_enumTypes[0]
}

func (x Person_PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Person_PhoneType.Descriptor instead.
func (Person_PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{3, 0}
}

type TestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo string `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar string `protobuf:"bytes,2,opt,name=bar,proto3" json:"bar,omitempty" spec-name:"foo" ssz-size:"32"`
}

func (x *TestMessage) Reset() {
	*x = TestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestMessage) ProtoMessage() {}

func (x *TestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestMessage.ProtoReflect.Descriptor instead.
func (*TestMessage) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestMessage) GetFoo() string {
	if x != nil {
		return x.Foo
	}
	return ""
}

func (x *TestMessage) GetBar() string {
	if x != nil {
		return x.Bar
	}
	return ""
}

type TestNestedMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Fuzz string       `protobuf:"bytes,1,opt,name=fuzz,proto3" json:"fuzz,omitempty"`
	Msg  *TestMessage `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *TestNestedMessage) Reset() {
	*x = TestNestedMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestNestedMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestNestedMessage) ProtoMessage() {}

func (x *TestNestedMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestNestedMessage.ProtoReflect.Descriptor instead.
func (*TestNestedMessage) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{1}
}

func (x *TestNestedMessage) GetFuzz() string {
	if x != nil {
		return x.Fuzz
	}
	return ""
}

func (x *TestNestedMessage) GetMsg() *TestMessage {
	if x != nil {
		return x.Msg
	}
	return nil
}

type Puzzle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Challenge string `protobuf:"bytes,1,opt,name=challenge,proto3" json:"challenge,omitempty"`
	Answer    string `protobuf:"bytes,2,opt,name=answer,proto3" json:"answer,omitempty"`
}

func (x *Puzzle) Reset() {
	*x = Puzzle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Puzzle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Puzzle) ProtoMessage() {}

func (x *Puzzle) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Puzzle.ProtoReflect.Descriptor instead.
func (*Puzzle) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{2}
}

func (x *Puzzle) GetChallenge() string {
	if x != nil {
		return x.Challenge
	}
	return ""
}

func (x *Puzzle) GetAnswer() string {
	if x != nil {
		return x.Answer
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id          int32                 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Email       string                `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phones      []*Person_PhoneNumber `protobuf:"bytes,4,rep,name=phones,proto3" json:"phones,omitempty"`
	LastUpdated *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{3}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Person) GetPhones() []*Person_PhoneNumber {
	if x != nil {
		return x.Phones
	}
	return nil
}

func (x *Person) GetLastUpdated() *timestamp.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type AddressBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *AddressBook) Reset() {
	*x = AddressBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressBook) ProtoMessage() {}

func (x *AddressBook) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressBook.ProtoReflect.Descriptor instead.
func (*AddressBook) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{4}
}

func (x *AddressBook) GetPeople() []*Person {
	if x != nil {
		return x.People
	}
	return nil
}

type TestSimpleMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foo []byte `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	Bar uint64 `protobuf:"varint,2,opt,name=bar,proto3" json:"bar,omitempty"`
}

func (x *TestSimpleMessage) Reset() {
	*x = TestSimpleMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestSimpleMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestSimpleMessage) ProtoMessage() {}

func (x *TestSimpleMessage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestSimpleMessage.ProtoReflect.Descriptor instead.
func (*TestSimpleMessage) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{5}
}

func (x *TestSimpleMessage) GetFoo() []byte {
	if x != nil {
		return x.Foo
	}
	return nil
}

func (x *TestSimpleMessage) GetBar() uint64 {
	if x != nil {
		return x.Bar
	}
	return 0
}

type Person_PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string           `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   Person_PhoneType `protobuf:"varint,2,opt,name=type,proto3,enum=testing.Person_PhoneType" json:"type,omitempty"`
}

func (x *Person_PhoneNumber) Reset() {
	*x = Person_PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_testing_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_PhoneNumber) ProtoMessage() {}

func (x *Person_PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_proto_testing_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_PhoneNumber.ProtoReflect.Descriptor instead.
func (*Person_PhoneNumber) Descriptor() ([]byte, []int) {
	return file_proto_testing_test_proto_rawDescGZIP(), []int{3, 0}
}

func (x *Person_PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Person_PhoneNumber) GetType() Person_PhoneType {
	if x != nil {
		return x.Type
	}
	return Person_MOBILE
}

var file_proto_testing_test_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60001,
		Name:          "testing.ssz_size",
		Tag:           "bytes,60001,opt,name=ssz_size",
		Filename:      "proto/testing/test.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         60003,
		Name:          "testing.spec_name",
		Tag:           "bytes,60003,opt,name=spec_name",
		Filename:      "proto/testing/test.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string ssz_size = 60001;
	E_SszSize = &file_proto_testing_test_proto_extTypes[0]
	// optional string spec_name = 60003;
	E_SpecName = &file_proto_testing_test_proto_extTypes[1]
)

var File_proto_testing_test_proto protoreflect.FileDescriptor

var file_proto_testing_test_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0b, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x1f, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x8a, 0xa6, 0x1d, 0x02, 0x33, 0x32, 0x9a, 0xa6, 0x1d, 0x03,
	0x66, 0x6f, 0x6f, 0x52, 0x03, 0x62, 0x61, 0x72, 0x22, 0x4f, 0x0a, 0x11, 0x54, 0x65, 0x73, 0x74,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x75, 0x7a, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x75, 0x7a,
	0x7a, 0x12, 0x26, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3e, 0x0a, 0x06, 0x50, 0x75, 0x7a,
	0x7a, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x6e, 0x67,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22, 0xb9, 0x02, 0x0a, 0x06, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x33,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x1a, 0x54, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x2b, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57,
	0x4f, 0x52, 0x4b, 0x10, 0x02, 0x22, 0x36, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x22, 0x37, 0x0a,
	0x11, 0x54, 0x65, 0x73, 0x74, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x66, 0x6f, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x62, 0x61, 0x72, 0x3a, 0x3a, 0x0a, 0x08, 0x73, 0x73, 0x7a, 0x5f, 0x73, 0x69,
	0x7a, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xe1, 0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x73, 0x7a, 0x53, 0x69,
	0x7a, 0x65, 0x3a, 0x3c, 0x0a, 0x09, 0x73, 0x70, 0x65, 0x63, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe3,
	0xd4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x70, 0x65, 0x63, 0x4e, 0x61, 0x6d, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_testing_test_proto_rawDescOnce sync.Once
	file_proto_testing_test_proto_rawDescData = file_proto_testing_test_proto_rawDesc
)

func file_proto_testing_test_proto_rawDescGZIP() []byte {
	file_proto_testing_test_proto_rawDescOnce.Do(func() {
		file_proto_testing_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_testing_test_proto_rawDescData)
	})
	return file_proto_testing_test_proto_rawDescData
}

var file_proto_testing_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_testing_test_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_testing_test_proto_goTypes = []interface{}{
	(Person_PhoneType)(0),             // 0: testing.Person.PhoneType
	(*TestMessage)(nil),               // 1: testing.TestMessage
	(*TestNestedMessage)(nil),         // 2: testing.TestNestedMessage
	(*Puzzle)(nil),                    // 3: testing.Puzzle
	(*Person)(nil),                    // 4: testing.Person
	(*AddressBook)(nil),               // 5: testing.AddressBook
	(*TestSimpleMessage)(nil),         // 6: testing.TestSimpleMessage
	(*Person_PhoneNumber)(nil),        // 7: testing.Person.PhoneNumber
	(*timestamp.Timestamp)(nil),       // 8: google.protobuf.Timestamp
	(*descriptorpb.FieldOptions)(nil), // 9: google.protobuf.FieldOptions
}
var file_proto_testing_test_proto_depIdxs = []int32{
	1, // 0: testing.TestNestedMessage.msg:type_name -> testing.TestMessage
	7, // 1: testing.Person.phones:type_name -> testing.Person.PhoneNumber
	8, // 2: testing.Person.last_updated:type_name -> google.protobuf.Timestamp
	4, // 3: testing.AddressBook.people:type_name -> testing.Person
	0, // 4: testing.Person.PhoneNumber.type:type_name -> testing.Person.PhoneType
	9, // 5: testing.ssz_size:extendee -> google.protobuf.FieldOptions
	9, // 6: testing.spec_name:extendee -> google.protobuf.FieldOptions
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	5, // [5:7] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_testing_test_proto_init() }
func file_proto_testing_test_proto_init() {
	if File_proto_testing_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_testing_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestMessage); i {
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
		file_proto_testing_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestNestedMessage); i {
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
		file_proto_testing_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Puzzle); i {
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
		file_proto_testing_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_proto_testing_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressBook); i {
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
		file_proto_testing_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestSimpleMessage); i {
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
		file_proto_testing_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_PhoneNumber); i {
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
			RawDescriptor: file_proto_testing_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_proto_testing_test_proto_goTypes,
		DependencyIndexes: file_proto_testing_test_proto_depIdxs,
		EnumInfos:         file_proto_testing_test_proto_enumTypes,
		MessageInfos:      file_proto_testing_test_proto_msgTypes,
		ExtensionInfos:    file_proto_testing_test_proto_extTypes,
	}.Build()
	File_proto_testing_test_proto = out.File
	file_proto_testing_test_proto_rawDesc = nil
	file_proto_testing_test_proto_goTypes = nil
	file_proto_testing_test_proto_depIdxs = nil
}
