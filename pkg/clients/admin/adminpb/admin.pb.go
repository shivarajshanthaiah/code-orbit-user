// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: admin.proto

package __

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

type SubscriptionID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *SubscriptionID) Reset() {
	*x = SubscriptionID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionID) ProtoMessage() {}

func (x *SubscriptionID) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionID.ProtoReflect.Descriptor instead.
func (*SubscriptionID) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{0}
}

func (x *SubscriptionID) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

type AdNoParam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AdNoParam) Reset() {
	*x = AdNoParam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdNoParam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdNoParam) ProtoMessage() {}

func (x *AdNoParam) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdNoParam.ProtoReflect.Descriptor instead.
func (*AdNoParam) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{1}
}

type AdPlanList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Plans []*AdSubscription `protobuf:"bytes,1,rep,name=plans,proto3" json:"plans,omitempty"`
}

func (x *AdPlanList) Reset() {
	*x = AdPlanList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdPlanList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdPlanList) ProtoMessage() {}

func (x *AdPlanList) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdPlanList.ProtoReflect.Descriptor instead.
func (*AdPlanList) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{2}
}

func (x *AdPlanList) GetPlans() []*AdSubscription {
	if x != nil {
		return x.Plans
	}
	return nil
}

type AdSubscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         uint32  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Plan       string  `protobuf:"bytes,2,opt,name=plan,proto3" json:"plan,omitempty"`
	Duration   string  `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Price      float64 `protobuf:"fixed64,4,opt,name=price,proto3" json:"price,omitempty"`
	Gst        float64 `protobuf:"fixed64,5,opt,name=gst,proto3" json:"gst,omitempty"`
	TotalPrice float64 `protobuf:"fixed64,6,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *AdSubscription) Reset() {
	*x = AdSubscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdSubscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdSubscription) ProtoMessage() {}

func (x *AdSubscription) ProtoReflect() protoreflect.Message {
	mi := &file_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdSubscription.ProtoReflect.Descriptor instead.
func (*AdSubscription) Descriptor() ([]byte, []int) {
	return file_admin_proto_rawDescGZIP(), []int{3}
}

func (x *AdSubscription) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AdSubscription) GetPlan() string {
	if x != nil {
		return x.Plan
	}
	return ""
}

func (x *AdSubscription) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *AdSubscription) GetPrice() float64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *AdSubscription) GetGst() float64 {
	if x != nil {
		return x.Gst
	}
	return 0
}

func (x *AdSubscription) GetTotalPrice() float64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

var File_admin_proto protoreflect.FileDescriptor

var file_admin_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x20, 0x0a, 0x0e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x02, 0x49, 0x44, 0x22, 0x0b, 0x0a, 0x09, 0x41, 0x64, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x22, 0x36, 0x0a, 0x0a, 0x41, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28,
	0x0a, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x70, 0x6c, 0x61, 0x6e, 0x73, 0x22, 0x99, 0x01, 0x0a, 0x0e, 0x41, 0x64, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6c, 0x61, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x73, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x03,
	0x67, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x32, 0x7b, 0x0a, 0x0c, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6c,
	0x61, 0x6e, 0x73, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x4e, 0x6f, 0x50, 0x61, 0x72,
	0x61, 0x6d, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x64, 0x50, 0x6c, 0x61, 0x6e, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x3d, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x79, 0x49, 0x44, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x1a, 0x12, 0x2e,
	0x70, 0x62, 0x2e, 0x41, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_admin_proto_rawDescOnce sync.Once
	file_admin_proto_rawDescData = file_admin_proto_rawDesc
)

func file_admin_proto_rawDescGZIP() []byte {
	file_admin_proto_rawDescOnce.Do(func() {
		file_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_admin_proto_rawDescData)
	})
	return file_admin_proto_rawDescData
}

var file_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_admin_proto_goTypes = []interface{}{
	(*SubscriptionID)(nil), // 0: pb.SubscriptionID
	(*AdNoParam)(nil),      // 1: pb.AdNoParam
	(*AdPlanList)(nil),     // 2: pb.AdPlanList
	(*AdSubscription)(nil), // 3: pb.AdSubscription
}
var file_admin_proto_depIdxs = []int32{
	3, // 0: pb.AdPlanList.plans:type_name -> pb.AdSubscription
	1, // 1: pb.AdminService.GetAllPlans:input_type -> pb.AdNoParam
	0, // 2: pb.AdminService.GetSubscriptionByID:input_type -> pb.SubscriptionID
	2, // 3: pb.AdminService.GetAllPlans:output_type -> pb.AdPlanList
	3, // 4: pb.AdminService.GetSubscriptionByID:output_type -> pb.AdSubscription
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_admin_proto_init() }
func file_admin_proto_init() {
	if File_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_admin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionID); i {
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
		file_admin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdNoParam); i {
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
		file_admin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdPlanList); i {
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
		file_admin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdSubscription); i {
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
			RawDescriptor: file_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_admin_proto_goTypes,
		DependencyIndexes: file_admin_proto_depIdxs,
		MessageInfos:      file_admin_proto_msgTypes,
	}.Build()
	File_admin_proto = out.File
	file_admin_proto_rawDesc = nil
	file_admin_proto_goTypes = nil
	file_admin_proto_depIdxs = nil
}
