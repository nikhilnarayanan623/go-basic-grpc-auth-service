// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: pkg/pb/user.proto

package pb

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

type SaveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *SaveUserRequest) Reset() {
	*x = SaveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserRequest) ProtoMessage() {}

func (x *SaveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserRequest.ProtoReflect.Descriptor instead.
func (*SaveUserRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{0}
}

func (x *SaveUserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *SaveUserRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *SaveUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SaveUserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type SaveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	UserId   uint32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *SaveUserResponse) Reset() {
	*x = SaveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveUserResponse) ProtoMessage() {}

func (x *SaveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveUserResponse.ProtoReflect.Descriptor instead.
func (*SaveUserResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{1}
}

func (x *SaveUserResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *SaveUserResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type FindUserByEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *FindUserByEmailRequest) Reset() {
	*x = FindUserByEmailRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserByEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserByEmailRequest) ProtoMessage() {}

func (x *FindUserByEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserByEmailRequest.ProtoReflect.Descriptor instead.
func (*FindUserByEmailRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{2}
}

func (x *FindUserByEmailRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type FindUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response  *Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	UserId    uint32    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName string    `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string    `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string    `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Password  string    `protobuf:"bytes,6,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *FindUserResponse) Reset() {
	*x = FindUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindUserResponse) ProtoMessage() {}

func (x *FindUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindUserResponse.ProtoReflect.Descriptor instead.
func (*FindUserResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_user_proto_rawDescGZIP(), []int{3}
}

func (x *FindUserResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *FindUserResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *FindUserResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *FindUserResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *FindUserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *FindUserResponse) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_pkg_pb_user_proto protoreflect.FileDescriptor

var file_pkg_pb_user_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x15, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7f,
	0x0a, 0x0f, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x55, 0x0a, 0x10, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x16, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xc3, 0x01, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x32, 0x8d, 0x01, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0f,
	0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x61, 0x76, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0a, 0x5a, 0x08,
	0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_user_proto_rawDescOnce sync.Once
	file_pkg_pb_user_proto_rawDescData = file_pkg_pb_user_proto_rawDesc
)

func file_pkg_pb_user_proto_rawDescGZIP() []byte {
	file_pkg_pb_user_proto_rawDescOnce.Do(func() {
		file_pkg_pb_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_user_proto_rawDescData)
	})
	return file_pkg_pb_user_proto_rawDescData
}

var file_pkg_pb_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_pb_user_proto_goTypes = []interface{}{
	(*SaveUserRequest)(nil),        // 0: pb.SaveUserRequest
	(*SaveUserResponse)(nil),       // 1: pb.SaveUserResponse
	(*FindUserByEmailRequest)(nil), // 2: pb.FindUserByEmailRequest
	(*FindUserResponse)(nil),       // 3: pb.FindUserResponse
	(*Response)(nil),               // 4: pb.Response
}
var file_pkg_pb_user_proto_depIdxs = []int32{
	4, // 0: pb.SaveUserResponse.response:type_name -> pb.Response
	4, // 1: pb.FindUserResponse.response:type_name -> pb.Response
	2, // 2: pb.UserService.FindUserByEmail:input_type -> pb.FindUserByEmailRequest
	0, // 3: pb.UserService.SaveUser:input_type -> pb.SaveUserRequest
	3, // 4: pb.UserService.FindUserByEmail:output_type -> pb.FindUserResponse
	1, // 5: pb.UserService.SaveUser:output_type -> pb.SaveUserResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_pb_user_proto_init() }
func file_pkg_pb_user_proto_init() {
	if File_pkg_pb_user_proto != nil {
		return
	}
	file_pkg_pb_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SaveUserResponse); i {
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
		file_pkg_pb_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserByEmailRequest); i {
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
		file_pkg_pb_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindUserResponse); i {
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
			RawDescriptor: file_pkg_pb_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_user_proto_goTypes,
		DependencyIndexes: file_pkg_pb_user_proto_depIdxs,
		MessageInfos:      file_pkg_pb_user_proto_msgTypes,
	}.Build()
	File_pkg_pb_user_proto = out.File
	file_pkg_pb_user_proto_rawDesc = nil
	file_pkg_pb_user_proto_goTypes = nil
	file_pkg_pb_user_proto_depIdxs = nil
}
