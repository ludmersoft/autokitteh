// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: autokitteh/projects/v1/project.proto

package projectsv1

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

type Project struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId     string   `protobuf:"bytes,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Name          string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResourcePaths []string `protobuf:"bytes,4,rep,name=resource_paths,json=resourcePaths,proto3" json:"resource_paths,omitempty"` // can be globs.
}

func (x *Project) Reset() {
	*x = Project{}
	if protoimpl.UnsafeEnabled {
		mi := &file_autokitteh_projects_v1_project_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Project) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Project) ProtoMessage() {}

func (x *Project) ProtoReflect() protoreflect.Message {
	mi := &file_autokitteh_projects_v1_project_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Project.ProtoReflect.Descriptor instead.
func (*Project) Descriptor() ([]byte, []int) {
	return file_autokitteh_projects_v1_project_proto_rawDescGZIP(), []int{0}
}

func (x *Project) GetProjectId() string {
	if x != nil {
		return x.ProjectId
	}
	return ""
}

func (x *Project) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Project) GetResourcePaths() []string {
	if x != nil {
		return x.ResourcePaths
	}
	return nil
}

var File_autokitteh_projects_v1_project_proto protoreflect.FileDescriptor

var file_autokitteh_projects_v1_project_proto_rawDesc = []byte{
	0x0a, 0x24, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x63,
	0x0a, 0x07, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x50, 0x61,
	0x74, 0x68, 0x73, 0x42, 0xf1, 0x01, 0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x74, 0x6f,
	0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x4b, 0x67, 0x6f, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65,
	0x68, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x75,
	0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x41, 0x50, 0x58, 0xaa, 0x02, 0x16, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74,
	0x65, 0x68, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02,
	0x16, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x5c, 0x50, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x22, 0x41, 0x75, 0x74, 0x6f, 0x6b, 0x69,
	0x74, 0x74, 0x65, 0x68, 0x5c, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x18, 0x41,
	0x75, 0x74, 0x6f, 0x6b, 0x69, 0x74, 0x74, 0x65, 0x68, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_autokitteh_projects_v1_project_proto_rawDescOnce sync.Once
	file_autokitteh_projects_v1_project_proto_rawDescData = file_autokitteh_projects_v1_project_proto_rawDesc
)

func file_autokitteh_projects_v1_project_proto_rawDescGZIP() []byte {
	file_autokitteh_projects_v1_project_proto_rawDescOnce.Do(func() {
		file_autokitteh_projects_v1_project_proto_rawDescData = protoimpl.X.CompressGZIP(file_autokitteh_projects_v1_project_proto_rawDescData)
	})
	return file_autokitteh_projects_v1_project_proto_rawDescData
}

var file_autokitteh_projects_v1_project_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_autokitteh_projects_v1_project_proto_goTypes = []interface{}{
	(*Project)(nil), // 0: autokitteh.projects.v1.Project
}
var file_autokitteh_projects_v1_project_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_autokitteh_projects_v1_project_proto_init() }
func file_autokitteh_projects_v1_project_proto_init() {
	if File_autokitteh_projects_v1_project_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_autokitteh_projects_v1_project_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Project); i {
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
			RawDescriptor: file_autokitteh_projects_v1_project_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_autokitteh_projects_v1_project_proto_goTypes,
		DependencyIndexes: file_autokitteh_projects_v1_project_proto_depIdxs,
		MessageInfos:      file_autokitteh_projects_v1_project_proto_msgTypes,
	}.Build()
	File_autokitteh_projects_v1_project_proto = out.File
	file_autokitteh_projects_v1_project_proto_rawDesc = nil
	file_autokitteh_projects_v1_project_proto_goTypes = nil
	file_autokitteh_projects_v1_project_proto_depIdxs = nil
}
