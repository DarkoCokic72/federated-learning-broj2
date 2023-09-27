// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.1
// source: train.proto

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

type TrainReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SenderAddress     string `protobuf:"bytes,1,opt,name=SenderAddress,proto3" json:"SenderAddress,omitempty"`
	SenderId          string `protobuf:"bytes,2,opt,name=SenderId,proto3" json:"SenderId,omitempty"`
	MarshalledWeights []byte `protobuf:"bytes,3,opt,name=marshalledWeights,proto3" json:"marshalledWeights,omitempty"`
}

func (x *TrainReq) Reset() {
	*x = TrainReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainReq) ProtoMessage() {}

func (x *TrainReq) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainReq.ProtoReflect.Descriptor instead.
func (*TrainReq) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{0}
}

func (x *TrainReq) GetSenderAddress() string {
	if x != nil {
		return x.SenderAddress
	}
	return ""
}

func (x *TrainReq) GetSenderId() string {
	if x != nil {
		return x.SenderId
	}
	return ""
}

func (x *TrainReq) GetMarshalledWeights() []byte {
	if x != nil {
		return x.MarshalledWeights
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Weights *WeightsList `protobuf:"bytes,1,opt,name=Weights,proto3" json:"Weights,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetWeights() *WeightsList {
	if x != nil {
		return x.Weights
	}
	return nil
}

type WeightsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InnerList []*WeightsListInner `protobuf:"bytes,1,rep,name=inner_list,json=innerList,proto3" json:"inner_list,omitempty"`
}

func (x *WeightsList) Reset() {
	*x = WeightsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightsList) ProtoMessage() {}

func (x *WeightsList) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightsList.ProtoReflect.Descriptor instead.
func (*WeightsList) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{2}
}

func (x *WeightsList) GetInnerList() []*WeightsListInner {
	if x != nil {
		return x.InnerList
	}
	return nil
}

type WeightsListInner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ValueList []float32 `protobuf:"fixed32,1,rep,packed,name=value_list,json=valueList,proto3" json:"value_list,omitempty"`
}

func (x *WeightsListInner) Reset() {
	*x = WeightsListInner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeightsListInner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeightsListInner) ProtoMessage() {}

func (x *WeightsListInner) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeightsListInner.ProtoReflect.Descriptor instead.
func (*WeightsListInner) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{3}
}

func (x *WeightsListInner) GetValueList() []float32 {
	if x != nil {
		return x.ValueList
	}
	return nil
}

type TrainResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *TrainResp) Reset() {
	*x = TrainResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrainResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrainResp) ProtoMessage() {}

func (x *TrainResp) ProtoReflect() protoreflect.Message {
	mi := &file_train_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrainResp.ProtoReflect.Descriptor instead.
func (*TrainResp) Descriptor() ([]byte, []int) {
	return file_train_proto_rawDescGZIP(), []int{4}
}

func (x *TrainResp) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_train_proto protoreflect.FileDescriptor

var file_train_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7a, 0x0a, 0x08, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x12, 0x24, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x64,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x11, 0x6d,
	0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x6c, 0x65, 0x64, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73,
	0x22, 0x38, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07,
	0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x07, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x22, 0x45, 0x0a, 0x0b, 0x57, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0a, 0x69, 0x6e, 0x6e,
	0x65, 0x72, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x49, 0x6e, 0x6e, 0x65, 0x72, 0x52, 0x09, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x31, 0x0a, 0x10, 0x57, 0x65, 0x69, 0x67, 0x68, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x02, 0x52, 0x09, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x22, 0x1f, 0x0a, 0x09, 0x54, 0x72, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x20, 0x5a, 0x1e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2d, 0x6c, 0x65, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x2d, 0x62, 0x72, 0x6f, 0x6a,
	0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_train_proto_rawDescOnce sync.Once
	file_train_proto_rawDescData = file_train_proto_rawDesc
)

func file_train_proto_rawDescGZIP() []byte {
	file_train_proto_rawDescOnce.Do(func() {
		file_train_proto_rawDescData = protoimpl.X.CompressGZIP(file_train_proto_rawDescData)
	})
	return file_train_proto_rawDescData
}

var file_train_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_train_proto_goTypes = []interface{}{
	(*TrainReq)(nil),         // 0: proto.TrainReq
	(*Response)(nil),         // 1: proto.Response
	(*WeightsList)(nil),      // 2: proto.WeightsList
	(*WeightsListInner)(nil), // 3: proto.WeightsListInner
	(*TrainResp)(nil),        // 4: proto.TrainResp
}
var file_train_proto_depIdxs = []int32{
	2, // 0: proto.Response.Weights:type_name -> proto.WeightsList
	3, // 1: proto.WeightsList.inner_list:type_name -> proto.WeightsListInner
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_train_proto_init() }
func file_train_proto_init() {
	if File_train_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_train_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainReq); i {
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
		file_train_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_train_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightsList); i {
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
		file_train_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeightsListInner); i {
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
		file_train_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrainResp); i {
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
			RawDescriptor: file_train_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_train_proto_goTypes,
		DependencyIndexes: file_train_proto_depIdxs,
		MessageInfos:      file_train_proto_msgTypes,
	}.Build()
	File_train_proto = out.File
	file_train_proto_rawDesc = nil
	file_train_proto_goTypes = nil
	file_train_proto_depIdxs = nil
}
