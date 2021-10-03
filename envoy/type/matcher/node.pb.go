// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.18.0
// source: envoy/type/matcher/node.proto

package envoy_type_matcher

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Specifies the way to match a Node.
// The match follows AND semantics.
type NodeMatcher struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies match criteria on the node id.
	NodeId *StringMatcher `protobuf:"bytes,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	// Specifies match criteria on the node metadata.
	NodeMetadatas []*StructMatcher `protobuf:"bytes,2,rep,name=node_metadatas,json=nodeMetadatas,proto3" json:"node_metadatas,omitempty"`
}

func (x *NodeMatcher) Reset() {
	*x = NodeMatcher{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_type_matcher_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeMatcher) ProtoMessage() {}

func (x *NodeMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_type_matcher_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeMatcher.ProtoReflect.Descriptor instead.
func (*NodeMatcher) Descriptor() ([]byte, []int) {
	return file_envoy_type_matcher_node_proto_rawDescGZIP(), []int{0}
}

func (x *NodeMatcher) GetNodeId() *StringMatcher {
	if x != nil {
		return x.NodeId
	}
	return nil
}

func (x *NodeMatcher) GetNodeMetadatas() []*StructMatcher {
	if x != nil {
		return x.NodeMetadatas
	}
	return nil
}

var File_envoy_type_matcher_node_proto protoreflect.FileDescriptor

var file_envoy_type_matcher_node_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x12, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2f,
	0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x93, 0x01, 0x0a, 0x0b, 0x4e, 0x6f, 0x64, 0x65, 0x4d, 0x61, 0x74,
	0x63, 0x68, 0x65, 0x72, 0x12, 0x3a, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x48, 0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x52, 0x0d, 0x6e, 0x6f, 0x64,
	0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x73, 0x42, 0x37, 0x0a, 0x20, 0x69, 0x6f,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f,
	0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x42, 0x09,
	0x4e, 0x6f, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06,
	0x02, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_type_matcher_node_proto_rawDescOnce sync.Once
	file_envoy_type_matcher_node_proto_rawDescData = file_envoy_type_matcher_node_proto_rawDesc
)

func file_envoy_type_matcher_node_proto_rawDescGZIP() []byte {
	file_envoy_type_matcher_node_proto_rawDescOnce.Do(func() {
		file_envoy_type_matcher_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_type_matcher_node_proto_rawDescData)
	})
	return file_envoy_type_matcher_node_proto_rawDescData
}

var file_envoy_type_matcher_node_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_type_matcher_node_proto_goTypes = []interface{}{
	(*NodeMatcher)(nil),   // 0: envoy.type.matcher.NodeMatcher
	(*StringMatcher)(nil), // 1: envoy.type.matcher.StringMatcher
	(*StructMatcher)(nil), // 2: envoy.type.matcher.StructMatcher
}
var file_envoy_type_matcher_node_proto_depIdxs = []int32{
	1, // 0: envoy.type.matcher.NodeMatcher.node_id:type_name -> envoy.type.matcher.StringMatcher
	2, // 1: envoy.type.matcher.NodeMatcher.node_metadatas:type_name -> envoy.type.matcher.StructMatcher
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_envoy_type_matcher_node_proto_init() }
func file_envoy_type_matcher_node_proto_init() {
	if File_envoy_type_matcher_node_proto != nil {
		return
	}
	file_envoy_type_matcher_string_proto_init()
	file_envoy_type_matcher_struct_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_type_matcher_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeMatcher); i {
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
			RawDescriptor: file_envoy_type_matcher_node_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_type_matcher_node_proto_goTypes,
		DependencyIndexes: file_envoy_type_matcher_node_proto_depIdxs,
		MessageInfos:      file_envoy_type_matcher_node_proto_msgTypes,
	}.Build()
	File_envoy_type_matcher_node_proto = out.File
	file_envoy_type_matcher_node_proto_rawDesc = nil
	file_envoy_type_matcher_node_proto_goTypes = nil
	file_envoy_type_matcher_node_proto_depIdxs = nil
}
