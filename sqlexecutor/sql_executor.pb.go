// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: hicon-sm/sql_executor.proto

package sqlexecutor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_hicon_sm_sql_executor_proto protoreflect.FileDescriptor

var file_hicon_sm_sql_executor_proto_rawDesc = string([]byte{
	0x0a, 0x1b, 0x68, 0x69, 0x63, 0x6f, 0x6e, 0x2d, 0x73, 0x6d, 0x2f, 0x73, 0x71, 0x6c, 0x5f, 0x65,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x9d, 0x05, 0x0a, 0x0b, 0x53, 0x51, 0x4c,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x3c, 0x0a, 0x0c, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79,
	0x50, 0x4b, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00,
	0x12, 0x37, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x07, 0x46, 0x69, 0x6e,
	0x64, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x22, 0x00, 0x12, 0x34, 0x0a, 0x04, 0x45, 0x78, 0x65, 0x63, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x42, 0x75, 0x6c, 0x6b,
	0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79,
	0x50, 0x4b, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00,
	0x12, 0x39, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x12, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0e, 0x42,
	0x75, 0x6c, 0x6b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x50, 0x4b, 0x12, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x4b, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x1a,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0f, 0x42, 0x75, 0x6c, 0x6b, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x57, 0x69, 0x74, 0x68, 0x54, 0x78, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79,
	0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x73, 0x71,
	0x6c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var file_hicon_sm_sql_executor_proto_goTypes = []any{
	(*anypb.Any)(nil), // 0: google.protobuf.Any
}
var file_hicon_sm_sql_executor_proto_depIdxs = []int32{
	0,  // 0: SQLExecutor.UpsertConfig:input_type -> google.protobuf.Any
	0,  // 1: SQLExecutor.FindByPK:input_type -> google.protobuf.Any
	0,  // 2: SQLExecutor.FindOne:input_type -> google.protobuf.Any
	0,  // 3: SQLExecutor.FindAll:input_type -> google.protobuf.Any
	0,  // 4: SQLExecutor.Exec:input_type -> google.protobuf.Any
	0,  // 5: SQLExecutor.BulkInsert:input_type -> google.protobuf.Any
	0,  // 6: SQLExecutor.UpdateByPK:input_type -> google.protobuf.Any
	0,  // 7: SQLExecutor.UpdateAll:input_type -> google.protobuf.Any
	0,  // 8: SQLExecutor.BulkUpdateByPK:input_type -> google.protobuf.Any
	0,  // 9: SQLExecutor.DeleteByPK:input_type -> google.protobuf.Any
	0,  // 10: SQLExecutor.BulkWriteWithTx:input_type -> google.protobuf.Any
	0,  // 11: SQLExecutor.UpsertConfig:output_type -> google.protobuf.Any
	0,  // 12: SQLExecutor.FindByPK:output_type -> google.protobuf.Any
	0,  // 13: SQLExecutor.FindOne:output_type -> google.protobuf.Any
	0,  // 14: SQLExecutor.FindAll:output_type -> google.protobuf.Any
	0,  // 15: SQLExecutor.Exec:output_type -> google.protobuf.Any
	0,  // 16: SQLExecutor.BulkInsert:output_type -> google.protobuf.Any
	0,  // 17: SQLExecutor.UpdateByPK:output_type -> google.protobuf.Any
	0,  // 18: SQLExecutor.UpdateAll:output_type -> google.protobuf.Any
	0,  // 19: SQLExecutor.BulkUpdateByPK:output_type -> google.protobuf.Any
	0,  // 20: SQLExecutor.DeleteByPK:output_type -> google.protobuf.Any
	0,  // 21: SQLExecutor.BulkWriteWithTx:output_type -> google.protobuf.Any
	11, // [11:22] is the sub-list for method output_type
	0,  // [0:11] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_hicon_sm_sql_executor_proto_init() }
func file_hicon_sm_sql_executor_proto_init() {
	if File_hicon_sm_sql_executor_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_hicon_sm_sql_executor_proto_rawDesc), len(file_hicon_sm_sql_executor_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hicon_sm_sql_executor_proto_goTypes,
		DependencyIndexes: file_hicon_sm_sql_executor_proto_depIdxs,
	}.Build()
	File_hicon_sm_sql_executor_proto = out.File
	file_hicon_sm_sql_executor_proto_goTypes = nil
	file_hicon_sm_sql_executor_proto_depIdxs = nil
}
