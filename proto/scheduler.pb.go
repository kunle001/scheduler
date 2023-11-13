// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.0
// source: proto/scheduler.proto

package proto

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

type CreateScheduleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	ExpiryDate string `protobuf:"bytes,2,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
}

func (x *CreateScheduleRequest) Reset() {
	*x = CreateScheduleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleRequest) ProtoMessage() {}

func (x *CreateScheduleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleRequest.ProtoReflect.Descriptor instead.
func (*CreateScheduleRequest) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_proto_rawDescGZIP(), []int{0}
}

func (x *CreateScheduleRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateScheduleRequest) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

type CreateScheduleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpiryDate string `protobuf:"bytes,2,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	Type       string `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
	Status     string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreatedAt  string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt  string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *CreateScheduleResponse) Reset() {
	*x = CreateScheduleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_scheduler_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateScheduleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateScheduleResponse) ProtoMessage() {}

func (x *CreateScheduleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_scheduler_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateScheduleResponse.ProtoReflect.Descriptor instead.
func (*CreateScheduleResponse) Descriptor() ([]byte, []int) {
	return file_proto_scheduler_proto_rawDescGZIP(), []int{1}
}

func (x *CreateScheduleResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateScheduleResponse) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *CreateScheduleResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *CreateScheduleResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateScheduleResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *CreateScheduleResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_proto_scheduler_proto protoreflect.FileDescriptor

var file_proto_scheduler_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x4c, 0x0a, 0x15, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79,
	0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70,
	0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb3, 0x01, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x78, 0x0a,
	0x0f, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x65, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x12, 0x27, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_scheduler_proto_rawDescOnce sync.Once
	file_proto_scheduler_proto_rawDescData = file_proto_scheduler_proto_rawDesc
)

func file_proto_scheduler_proto_rawDescGZIP() []byte {
	file_proto_scheduler_proto_rawDescOnce.Do(func() {
		file_proto_scheduler_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_scheduler_proto_rawDescData)
	})
	return file_proto_scheduler_proto_rawDescData
}

var file_proto_scheduler_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_scheduler_proto_goTypes = []interface{}{
	(*CreateScheduleRequest)(nil),  // 0: schedule_service.CreateScheduleRequest
	(*CreateScheduleResponse)(nil), // 1: schedule_service.CreateScheduleResponse
}
var file_proto_scheduler_proto_depIdxs = []int32{
	0, // 0: schedule_service.ScheduleService.CreateSchedule:input_type -> schedule_service.CreateScheduleRequest
	1, // 1: schedule_service.ScheduleService.CreateSchedule:output_type -> schedule_service.CreateScheduleResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_scheduler_proto_init() }
func file_proto_scheduler_proto_init() {
	if File_proto_scheduler_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_scheduler_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleRequest); i {
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
		file_proto_scheduler_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateScheduleResponse); i {
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
			RawDescriptor: file_proto_scheduler_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_scheduler_proto_goTypes,
		DependencyIndexes: file_proto_scheduler_proto_depIdxs,
		MessageInfos:      file_proto_scheduler_proto_msgTypes,
	}.Build()
	File_proto_scheduler_proto = out.File
	file_proto_scheduler_proto_rawDesc = nil
	file_proto_scheduler_proto_goTypes = nil
	file_proto_scheduler_proto_depIdxs = nil
}