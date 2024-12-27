// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v4.24.4
// source: v1alpha1/meta.proto

// Meta package defines protobuf serialization of standard COSI resources from the 'meta' namespace.

package v1alpha1

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ResourceDefinitionSpec_Sensitivity int32

const (
	ResourceDefinitionSpec_NON_SENSITIVE ResourceDefinitionSpec_Sensitivity = 0
	ResourceDefinitionSpec_SENSITIVE     ResourceDefinitionSpec_Sensitivity = 1
)

// Enum value maps for ResourceDefinitionSpec_Sensitivity.
var (
	ResourceDefinitionSpec_Sensitivity_name = map[int32]string{
		0: "NON_SENSITIVE",
		1: "SENSITIVE",
	}
	ResourceDefinitionSpec_Sensitivity_value = map[string]int32{
		"NON_SENSITIVE": 0,
		"SENSITIVE":     1,
	}
)

func (x ResourceDefinitionSpec_Sensitivity) Enum() *ResourceDefinitionSpec_Sensitivity {
	p := new(ResourceDefinitionSpec_Sensitivity)
	*p = x
	return p
}

func (x ResourceDefinitionSpec_Sensitivity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceDefinitionSpec_Sensitivity) Descriptor() protoreflect.EnumDescriptor {
	return file_v1alpha1_meta_proto_enumTypes[0].Descriptor()
}

func (ResourceDefinitionSpec_Sensitivity) Type() protoreflect.EnumType {
	return &file_v1alpha1_meta_proto_enumTypes[0]
}

func (x ResourceDefinitionSpec_Sensitivity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceDefinitionSpec_Sensitivity.Descriptor instead.
func (ResourceDefinitionSpec_Sensitivity) EnumDescriptor() ([]byte, []int) {
	return file_v1alpha1_meta_proto_rawDescGZIP(), []int{1, 0}
}

// NamespaceSpec is the protobuf serialization of the Namespace resource.
type NamespaceSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Description of the namespace.
	Description   string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NamespaceSpec) Reset() {
	*x = NamespaceSpec{}
	mi := &file_v1alpha1_meta_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceSpec) ProtoMessage() {}

func (x *NamespaceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_meta_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceSpec.ProtoReflect.Descriptor instead.
func (*NamespaceSpec) Descriptor() ([]byte, []int) {
	return file_v1alpha1_meta_proto_rawDescGZIP(), []int{0}
}

func (x *NamespaceSpec) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

// ResourceDefinitionSpec is the protobuf serialization of the ResourceDefinition resource.
type ResourceDefinitionSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Canonical type name.
	ResourceType string `protobuf:"bytes,1,opt,name=resource_type,json=resourceType,proto3" json:"resource_type,omitempty"`
	// Displayed human-readable type name.
	DisplayType string `protobuf:"bytes,2,opt,name=display_type,json=displayType,proto3" json:"display_type,omitempty"`
	// Default namespace to look for the resource if no namespace is given.
	DefaultNamespace string `protobuf:"bytes,3,opt,name=default_namespace,json=defaultNamespace,proto3" json:"default_namespace,omitempty"`
	// Human-readable aliases.
	Aliases []string `protobuf:"bytes,4,rep,name=aliases,proto3" json:"aliases,omitempty"`
	// All aliases for automatic matching.
	AllAliases []string `protobuf:"bytes,5,rep,name=all_aliases,json=allAliases,proto3" json:"all_aliases,omitempty"`
	// Additional columns to print in table output.
	PrintColumns []*ResourceDefinitionSpec_PrintColumn `protobuf:"bytes,6,rep,name=print_columns,json=printColumns,proto3" json:"print_columns,omitempty"`
	// Sensitivity indicates how secret resource of this type is.
	// The empty value represents a non-sensitive resource.
	Sensitivity   ResourceDefinitionSpec_Sensitivity `protobuf:"varint,7,opt,name=sensitivity,proto3,enum=cosi.resource.meta.ResourceDefinitionSpec_Sensitivity" json:"sensitivity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceDefinitionSpec) Reset() {
	*x = ResourceDefinitionSpec{}
	mi := &file_v1alpha1_meta_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceDefinitionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDefinitionSpec) ProtoMessage() {}

func (x *ResourceDefinitionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_meta_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDefinitionSpec.ProtoReflect.Descriptor instead.
func (*ResourceDefinitionSpec) Descriptor() ([]byte, []int) {
	return file_v1alpha1_meta_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceDefinitionSpec) GetResourceType() string {
	if x != nil {
		return x.ResourceType
	}
	return ""
}

func (x *ResourceDefinitionSpec) GetDisplayType() string {
	if x != nil {
		return x.DisplayType
	}
	return ""
}

func (x *ResourceDefinitionSpec) GetDefaultNamespace() string {
	if x != nil {
		return x.DefaultNamespace
	}
	return ""
}

func (x *ResourceDefinitionSpec) GetAliases() []string {
	if x != nil {
		return x.Aliases
	}
	return nil
}

func (x *ResourceDefinitionSpec) GetAllAliases() []string {
	if x != nil {
		return x.AllAliases
	}
	return nil
}

func (x *ResourceDefinitionSpec) GetPrintColumns() []*ResourceDefinitionSpec_PrintColumn {
	if x != nil {
		return x.PrintColumns
	}
	return nil
}

func (x *ResourceDefinitionSpec) GetSensitivity() ResourceDefinitionSpec_Sensitivity {
	if x != nil {
		return x.Sensitivity
	}
	return ResourceDefinitionSpec_NON_SENSITIVE
}

type ResourceDefinitionSpec_PrintColumn struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	JsonPath      string                 `protobuf:"bytes,2,opt,name=json_path,json=jsonPath,proto3" json:"json_path,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceDefinitionSpec_PrintColumn) Reset() {
	*x = ResourceDefinitionSpec_PrintColumn{}
	mi := &file_v1alpha1_meta_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceDefinitionSpec_PrintColumn) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceDefinitionSpec_PrintColumn) ProtoMessage() {}

func (x *ResourceDefinitionSpec_PrintColumn) ProtoReflect() protoreflect.Message {
	mi := &file_v1alpha1_meta_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceDefinitionSpec_PrintColumn.ProtoReflect.Descriptor instead.
func (*ResourceDefinitionSpec_PrintColumn) Descriptor() ([]byte, []int) {
	return file_v1alpha1_meta_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResourceDefinitionSpec_PrintColumn) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ResourceDefinitionSpec_PrintColumn) GetJsonPath() string {
	if x != nil {
		return x.JsonPath
	}
	return ""
}

var File_v1alpha1_meta_proto protoreflect.FileDescriptor

var file_v1alpha1_meta_proto_rawDesc = []byte{
	0x0a, 0x13, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6d, 0x65, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x22, 0x31, 0x0a, 0x0d, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xf0, 0x03, 0x0a,
	0x16, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x2b, 0x0a, 0x11, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x61, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x61,
	0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x6c, 0x6c, 0x5f, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x6c, 0x6c,
	0x41, 0x6c, 0x69, 0x61, 0x73, 0x65, 0x73, 0x12, 0x5b, 0x0a, 0x0d, 0x70, 0x72, 0x69, 0x6e, 0x74,
	0x5f, 0x63, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36,
	0x2e, 0x63, 0x6f, 0x73, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d,
	0x65, 0x74, 0x61, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x50, 0x72, 0x69, 0x6e, 0x74,
	0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x52, 0x0c, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6c,
	0x75, 0x6d, 0x6e, 0x73, 0x12, 0x58, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x36, 0x2e, 0x63, 0x6f, 0x73, 0x69,
	0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x0b, 0x73, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x3e,
	0x0a, 0x0b, 0x50, 0x72, 0x69, 0x6e, 0x74, 0x43, 0x6f, 0x6c, 0x75, 0x6d, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x22, 0x2f,
	0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x11, 0x0a,
	0x0d, 0x4e, 0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x53, 0x45, 0x4e, 0x53, 0x49, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x42,
	0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x73, 0x69, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x72, 0x75, 0x6e, 0x74, 0x69,
	0x6d, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1alpha1_meta_proto_rawDescOnce sync.Once
	file_v1alpha1_meta_proto_rawDescData = file_v1alpha1_meta_proto_rawDesc
)

func file_v1alpha1_meta_proto_rawDescGZIP() []byte {
	file_v1alpha1_meta_proto_rawDescOnce.Do(func() {
		file_v1alpha1_meta_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1alpha1_meta_proto_rawDescData)
	})
	return file_v1alpha1_meta_proto_rawDescData
}

var file_v1alpha1_meta_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1alpha1_meta_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_v1alpha1_meta_proto_goTypes = []any{
	(ResourceDefinitionSpec_Sensitivity)(0),    // 0: cosi.resource.meta.ResourceDefinitionSpec.Sensitivity
	(*NamespaceSpec)(nil),                      // 1: cosi.resource.meta.NamespaceSpec
	(*ResourceDefinitionSpec)(nil),             // 2: cosi.resource.meta.ResourceDefinitionSpec
	(*ResourceDefinitionSpec_PrintColumn)(nil), // 3: cosi.resource.meta.ResourceDefinitionSpec.PrintColumn
}
var file_v1alpha1_meta_proto_depIdxs = []int32{
	3, // 0: cosi.resource.meta.ResourceDefinitionSpec.print_columns:type_name -> cosi.resource.meta.ResourceDefinitionSpec.PrintColumn
	0, // 1: cosi.resource.meta.ResourceDefinitionSpec.sensitivity:type_name -> cosi.resource.meta.ResourceDefinitionSpec.Sensitivity
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_v1alpha1_meta_proto_init() }
func file_v1alpha1_meta_proto_init() {
	if File_v1alpha1_meta_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1alpha1_meta_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_v1alpha1_meta_proto_goTypes,
		DependencyIndexes: file_v1alpha1_meta_proto_depIdxs,
		EnumInfos:         file_v1alpha1_meta_proto_enumTypes,
		MessageInfos:      file_v1alpha1_meta_proto_msgTypes,
	}.Build()
	File_v1alpha1_meta_proto = out.File
	file_v1alpha1_meta_proto_rawDesc = nil
	file_v1alpha1_meta_proto_goTypes = nil
	file_v1alpha1_meta_proto_depIdxs = nil
}
