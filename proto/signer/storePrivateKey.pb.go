// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.3
// source: storePrivateKey.proto

package signer

import (
	blockchain "github.com/secure-protocol/protocol/proto/blockchain"
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

type GenType int32

const (
	GenType_Direct   GenType = 0 // 直接生成
	GenType_Mnemonic GenType = 1 // 助记词生成
	GenType_Derive   GenType = 2 // 派生
)

// Enum value maps for GenType.
var (
	GenType_name = map[int32]string{
		0: "Direct",
		1: "Mnemonic",
		2: "Derive",
	}
	GenType_value = map[string]int32{
		"Direct":   0,
		"Mnemonic": 1,
		"Derive":   2,
	}
)

func (x GenType) Enum() *GenType {
	p := new(GenType)
	*p = x
	return p
}

func (x GenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GenType) Descriptor() protoreflect.EnumDescriptor {
	return file_storePrivateKey_proto_enumTypes[0].Descriptor()
}

func (GenType) Type() protoreflect.EnumType {
	return &file_storePrivateKey_proto_enumTypes[0]
}

func (x GenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GenType.Descriptor instead.
func (GenType) EnumDescriptor() ([]byte, []int) {
	return file_storePrivateKey_proto_rawDescGZIP(), []int{0}
}

type Algorithm int32

const (
	Algorithm_ecdsa_secp256k1 Algorithm = 0
)

// Enum value maps for Algorithm.
var (
	Algorithm_name = map[int32]string{
		0: "ecdsa_secp256k1",
	}
	Algorithm_value = map[string]int32{
		"ecdsa_secp256k1": 0,
	}
)

func (x Algorithm) Enum() *Algorithm {
	p := new(Algorithm)
	*p = x
	return p
}

func (x Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_storePrivateKey_proto_enumTypes[1].Descriptor()
}

func (Algorithm) Type() protoreflect.EnumType {
	return &file_storePrivateKey_proto_enumTypes[1]
}

func (x Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Algorithm.Descriptor instead.
func (Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_storePrivateKey_proto_rawDescGZIP(), []int{1}
}

type StorePrivateKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountType blockchain.AccountType `protobuf:"varint,1,opt,name=AccountType,proto3,enum=blockchain.AccountType" json:"AccountType,omitempty"`
	GenType     GenType                `protobuf:"varint,2,opt,name=GenType,proto3,enum=signer.GenType" json:"GenType,omitempty"`
	Data        []byte                 `protobuf:"bytes,3,opt,name=Data,proto3" json:"Data,omitempty"`
	Encrypted   bool                   `protobuf:"varint,4,opt,name=Encrypted,proto3" json:"Encrypted,omitempty"`
	Delete      bool                   `protobuf:"varint,5,opt,name=Delete,proto3" json:"Delete,omitempty"`
	Algorithm   Algorithm              `protobuf:"varint,6,opt,name=algorithm,proto3,enum=signer.Algorithm" json:"algorithm,omitempty"`
}

func (x *StorePrivateKey) Reset() {
	*x = StorePrivateKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_storePrivateKey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StorePrivateKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StorePrivateKey) ProtoMessage() {}

func (x *StorePrivateKey) ProtoReflect() protoreflect.Message {
	mi := &file_storePrivateKey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StorePrivateKey.ProtoReflect.Descriptor instead.
func (*StorePrivateKey) Descriptor() ([]byte, []int) {
	return file_storePrivateKey_proto_rawDescGZIP(), []int{0}
}

func (x *StorePrivateKey) GetAccountType() blockchain.AccountType {
	if x != nil {
		return x.AccountType
	}
	return blockchain.AccountType(0)
}

func (x *StorePrivateKey) GetGenType() GenType {
	if x != nil {
		return x.GenType
	}
	return GenType_Direct
}

func (x *StorePrivateKey) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StorePrivateKey) GetEncrypted() bool {
	if x != nil {
		return x.Encrypted
	}
	return false
}

func (x *StorePrivateKey) GetDelete() bool {
	if x != nil {
		return x.Delete
	}
	return false
}

func (x *StorePrivateKey) GetAlgorithm() Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return Algorithm_ecdsa_secp256k1
}

var File_storePrivateKey_proto protoreflect.FileDescriptor

var file_storePrivateKey_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x1a,
	0x10, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xf2, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x50, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x29, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x0f, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x07, 0x47, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x09, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x45, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x12, 0x16, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74,
	0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x65,
	0x72, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67,
	0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x2a, 0x2f, 0x0a, 0x07, 0x47, 0x65, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x4d, 0x6e, 0x65, 0x6d, 0x6f, 0x6e, 0x69, 0x63, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x44,
	0x65, 0x72, 0x69, 0x76, 0x65, 0x10, 0x02, 0x2a, 0x20, 0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x12, 0x13, 0x0a, 0x0f, 0x65, 0x63, 0x64, 0x73, 0x61, 0x5f, 0x73, 0x65,
	0x63, 0x70, 0x32, 0x35, 0x36, 0x6b, 0x31, 0x10, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_storePrivateKey_proto_rawDescOnce sync.Once
	file_storePrivateKey_proto_rawDescData = file_storePrivateKey_proto_rawDesc
)

func file_storePrivateKey_proto_rawDescGZIP() []byte {
	file_storePrivateKey_proto_rawDescOnce.Do(func() {
		file_storePrivateKey_proto_rawDescData = protoimpl.X.CompressGZIP(file_storePrivateKey_proto_rawDescData)
	})
	return file_storePrivateKey_proto_rawDescData
}

var file_storePrivateKey_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_storePrivateKey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_storePrivateKey_proto_goTypes = []any{
	(GenType)(0),                // 0: signer.GenType
	(Algorithm)(0),              // 1: signer.Algorithm
	(*StorePrivateKey)(nil),     // 2: signer.StorePrivateKey
	(blockchain.AccountType)(0), // 3: blockchain.AccountType
}
var file_storePrivateKey_proto_depIdxs = []int32{
	3, // 0: signer.StorePrivateKey.AccountType:type_name -> blockchain.AccountType
	0, // 1: signer.StorePrivateKey.GenType:type_name -> signer.GenType
	1, // 2: signer.StorePrivateKey.algorithm:type_name -> signer.Algorithm
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_storePrivateKey_proto_init() }
func file_storePrivateKey_proto_init() {
	if File_storePrivateKey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_storePrivateKey_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*StorePrivateKey); i {
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
			RawDescriptor: file_storePrivateKey_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_storePrivateKey_proto_goTypes,
		DependencyIndexes: file_storePrivateKey_proto_depIdxs,
		EnumInfos:         file_storePrivateKey_proto_enumTypes,
		MessageInfos:      file_storePrivateKey_proto_msgTypes,
	}.Build()
	File_storePrivateKey_proto = out.File
	file_storePrivateKey_proto_rawDesc = nil
	file_storePrivateKey_proto_goTypes = nil
	file_storePrivateKey_proto_depIdxs = nil
}
