package jsonerror

import (
	reflect "reflect"
	sync "sync"
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	"log"
)

const (
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Error         *Error_Status `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *Error) gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (x *Error) gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Error) gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6() {
}
func (x *Error) gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6() protoreflect.Message {
	mi := &file_error_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Error) gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0}
}
func (x *Error) gologoo__GetError_545f4b099e3e9f98b438c6c0ca1f43b6() *Error_Status {
	if x != nil {
		return x.Error
	}
	return nil
}

type Error_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
	Code          int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message       string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Status        code.Code    `protobuf:"varint,4,opt,name=status,proto3,enum=google.rpc.Code" json:"status,omitempty"`
	Details       []*anypb.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Error_Status) gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6() {
	*x = Error_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_error_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}
func (x *Error_Status) gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6() string {
	return protoimpl.X.MessageStringOf(x)
}
func (*Error_Status) gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6() {
}
func (x *Error_Status) gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6() protoreflect.Message {
	mi := &file_error_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}
func (*Error_Status) gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6() ([]byte, []int) {
	return file_error_proto_rawDescGZIP(), []int{0, 0}
}
func (x *Error_Status) gologoo__GetCode_545f4b099e3e9f98b438c6c0ca1f43b6() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}
func (x *Error_Status) gologoo__GetMessage_545f4b099e3e9f98b438c6c0ca1f43b6() string {
	if x != nil {
		return x.Message
	}
	return ""
}
func (x *Error_Status) gologoo__GetStatus_545f4b099e3e9f98b438c6c0ca1f43b6() code.Code {
	if x != nil {
		return x.Status
	}
	return code.Code(0)
}
func (x *Error_Status) gologoo__GetDetails_545f4b099e3e9f98b438c6c0ca1f43b6() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_error_proto protoreflect.FileDescriptor
var file_error_proto_rawDesc = []byte{0x0a, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x29, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x1a, 0x90, 0x01, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x67, 0x61, 0x78, 0x2d, 0x67, 0x6f, 0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x6a, 0x73, 0x6f, 0x6e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33}
var (
	file_error_proto_rawDescOnce sync.Once
	file_error_proto_rawDescData = file_error_proto_rawDesc
)

func gologoo__file_error_proto_rawDescGZIP_545f4b099e3e9f98b438c6c0ca1f43b6() []byte {
	file_error_proto_rawDescOnce.Do(func() {
		file_error_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_proto_rawDescData)
	})
	return file_error_proto_rawDescData
}

var file_error_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_error_proto_goTypes = []interface {
}{(*Error)(nil), (*Error_Status)(nil), (code.Code)(0), (*anypb.Any)(nil)}
var file_error_proto_depIdxs = []int32{1, 2, 3, 3, 3, 3, 3, 0}

func gologoo__init_545f4b099e3e9f98b438c6c0ca1f43b6() {
	file_error_proto_init()
}
func gologoo__file_error_proto_init_545f4b099e3e9f98b438c6c0ca1f43b6() {
	if File_error_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_error_proto_msgTypes[0].Exporter = func(v interface {
		}, i int) interface {
		} {
			switch v := v.(*Error); i {
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
		file_error_proto_msgTypes[1].Exporter = func(v interface {
		}, i int) interface {
		} {
			switch v := v.(*Error_Status); i {
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
	type x struct {
	}
	out := protoimpl.TypeBuilder{File: protoimpl.DescBuilder{GoPackagePath: reflect.TypeOf(x{}).PkgPath(), RawDescriptor: file_error_proto_rawDesc, NumEnums: 0, NumMessages: 2, NumExtensions: 0, NumServices: 0}, GoTypes: file_error_proto_goTypes, DependencyIndexes: file_error_proto_depIdxs, MessageInfos: file_error_proto_msgTypes}.Build()
	File_error_proto = out.File
	file_error_proto_rawDesc = nil
	file_error_proto_goTypes = nil
	file_error_proto_depIdxs = nil
}
func (x *Error) Reset() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	x.gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func (x *Error) String() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (recv *Error) ProtoMessage() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	recv.gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func (x *Error) ProtoReflect() protoreflect.Message {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (recv *Error) Descriptor() ([]byte, []int) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0, r1 := recv.gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func (x *Error) GetError() *Error_Status {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GetError_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__GetError_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (x *Error_Status) Reset() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	x.gologoo__Reset_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func (x *Error_Status) String() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__String_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (recv *Error_Status) ProtoMessage() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	recv.gologoo__ProtoMessage_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func (x *Error_Status) ProtoReflect() protoreflect.Message {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__ProtoReflect_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (recv *Error_Status) Descriptor() ([]byte, []int) {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0, r1 := recv.gologoo__Descriptor_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v %v\n", r0, r1)
	return r0, r1
}
func (x *Error_Status) GetCode() int32 {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GetCode_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__GetCode_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (x *Error_Status) GetMessage() string {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GetMessage_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__GetMessage_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (x *Error_Status) GetStatus() code.Code {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GetStatus_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__GetStatus_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func (x *Error_Status) GetDetails() []*anypb.Any {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__GetDetails_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := x.gologoo__GetDetails_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func file_error_proto_rawDescGZIP() []byte {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__file_error_proto_rawDescGZIP_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	r0 := gologoo__file_error_proto_rawDescGZIP_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("Output: %v\n", r0)
	return r0
}
func init() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__init_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	gologoo__init_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
func file_error_proto_init() {
	log.SetFlags(19)
	log.Printf("📨 Call %s\n", "gologoo__file_error_proto_init_545f4b099e3e9f98b438c6c0ca1f43b6")
	log.Printf("Input : (none)\n")
	gologoo__file_error_proto_init_545f4b099e3e9f98b438c6c0ca1f43b6()
	log.Printf("🚚 Output: %v\n", "(none)")
	return
}
