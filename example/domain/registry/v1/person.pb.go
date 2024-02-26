// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: registry/v1/person.proto

package registryv1

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

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string          `protobuf:"bytes,1,opt,name=first_name,proto3" json:"first_name,omitempty"`
	Address *Person_Address `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// Types that are assignable to Type:
	//
	//	*Person_Admin_
	//	*Person_Manager_
	//	*Person_Client_
	Type isPerson_Type `protobuf_oneof:"type"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAddress() *Person_Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (m *Person) GetType() isPerson_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *Person) GetAdmin() *Person_Admin {
	if x, ok := x.GetType().(*Person_Admin_); ok {
		return x.Admin
	}
	return nil
}

func (x *Person) GetManager() *Person_Manager {
	if x, ok := x.GetType().(*Person_Manager_); ok {
		return x.Manager
	}
	return nil
}

func (x *Person) GetClient() *Person_Client {
	if x, ok := x.GetType().(*Person_Client_); ok {
		return x.Client
	}
	return nil
}

type isPerson_Type interface {
	isPerson_Type()
}

type Person_Admin_ struct {
	Admin *Person_Admin `protobuf:"bytes,3,opt,name=admin,proto3,oneof"`
}

type Person_Manager_ struct {
	Manager *Person_Manager `protobuf:"bytes,4,opt,name=manager,proto3,oneof"`
}

type Person_Client_ struct {
	Client *Person_Client `protobuf:"bytes,5,opt,name=client,proto3,oneof"`
}

func (*Person_Admin_) isPerson_Type() {}

func (*Person_Manager_) isPerson_Type() {}

func (*Person_Client_) isPerson_Type() {}

type Person_Admin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YearsInCharge int64 `protobuf:"varint,1,opt,name=years_in_charge,json=yearsInCharge,proto3" json:"years_in_charge,omitempty"`
	HasHolidays   bool  `protobuf:"varint,2,opt,name=has_holidays,json=hasHolidays,proto3" json:"has_holidays,omitempty"`
}

func (x *Person_Admin) Reset() {
	*x = Person_Admin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Admin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Admin) ProtoMessage() {}

func (x *Person_Admin) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Admin.ProtoReflect.Descriptor instead.
func (*Person_Admin) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Person_Admin) GetYearsInCharge() int64 {
	if x != nil {
		return x.YearsInCharge
	}
	return 0
}

func (x *Person_Admin) GetHasHolidays() bool {
	if x != nil {
		return x.HasHolidays
	}
	return false
}

type Person_Manager struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ManagesClients bool `protobuf:"varint,1,opt,name=manages_clients,json=managesClients,proto3" json:"manages_clients,omitempty"`
}

func (x *Person_Manager) Reset() {
	*x = Person_Manager{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Manager) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Manager) ProtoMessage() {}

func (x *Person_Manager) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Manager.ProtoReflect.Descriptor instead.
func (*Person_Manager) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Person_Manager) GetManagesClients() bool {
	if x != nil {
		return x.ManagesClients
	}
	return false
}

type Person_Client struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Subscription:
	//
	//	*Person_Client_Premium_
	//	*Person_Client_Gold_
	//	*Person_Client_Silver_
	Subscription isPerson_Client_Subscription `protobuf_oneof:"subscription"`
}

func (x *Person_Client) Reset() {
	*x = Person_Client{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Client) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Client) ProtoMessage() {}

func (x *Person_Client) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Client.ProtoReflect.Descriptor instead.
func (*Person_Client) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 2}
}

func (m *Person_Client) GetSubscription() isPerson_Client_Subscription {
	if m != nil {
		return m.Subscription
	}
	return nil
}

func (x *Person_Client) GetPremium() *Person_Client_Premium {
	if x, ok := x.GetSubscription().(*Person_Client_Premium_); ok {
		return x.Premium
	}
	return nil
}

func (x *Person_Client) GetGold() *Person_Client_Gold {
	if x, ok := x.GetSubscription().(*Person_Client_Gold_); ok {
		return x.Gold
	}
	return nil
}

func (x *Person_Client) GetSilver() *Person_Client_Silver {
	if x, ok := x.GetSubscription().(*Person_Client_Silver_); ok {
		return x.Silver
	}
	return nil
}

type isPerson_Client_Subscription interface {
	isPerson_Client_Subscription()
}

type Person_Client_Premium_ struct {
	Premium *Person_Client_Premium `protobuf:"bytes,1,opt,name=premium,proto3,oneof"`
}

type Person_Client_Gold_ struct {
	Gold *Person_Client_Gold `protobuf:"bytes,2,opt,name=gold,proto3,oneof"`
}

type Person_Client_Silver_ struct {
	Silver *Person_Client_Silver `protobuf:"bytes,3,opt,name=silver,proto3,oneof"`
}

func (*Person_Client_Premium_) isPerson_Client_Subscription() {}

func (*Person_Client_Gold_) isPerson_Client_Subscription() {}

func (*Person_Client_Silver_) isPerson_Client_Subscription() {}

type Person_Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street     string  `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	Number     string  `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
	Additional *string `protobuf:"bytes,3,opt,name=additional,proto3,oneof" json:"additional,omitempty"`
	Location   string  `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	Province   string  `protobuf:"bytes,5,opt,name=province,proto3" json:"province,omitempty"`
	PostalCode string  `protobuf:"bytes,6,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	Country    string  `protobuf:"bytes,7,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Person_Address) Reset() {
	*x = Person_Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Address) ProtoMessage() {}

func (x *Person_Address) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Address.ProtoReflect.Descriptor instead.
func (*Person_Address) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Person_Address) GetStreet() string {
	if x != nil {
		return x.Street
	}
	return ""
}

func (x *Person_Address) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *Person_Address) GetAdditional() string {
	if x != nil && x.Additional != nil {
		return *x.Additional
	}
	return ""
}

func (x *Person_Address) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Person_Address) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *Person_Address) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Person_Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type Person_Client_Premium struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	YearlyLimit int64 `protobuf:"varint,1,opt,name=yearly_limit,json=yearlyLimit,proto3" json:"yearly_limit,omitempty"`
}

func (x *Person_Client_Premium) Reset() {
	*x = Person_Client_Premium{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Client_Premium) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Client_Premium) ProtoMessage() {}

func (x *Person_Client_Premium) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Client_Premium.ProtoReflect.Descriptor instead.
func (*Person_Client_Premium) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 2, 0}
}

func (x *Person_Client_Premium) GetYearlyLimit() int64 {
	if x != nil {
		return x.YearlyLimit
	}
	return 0
}

type Person_Client_Gold struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MonthlyLimit int64 `protobuf:"varint,2,opt,name=monthly_limit,json=monthlyLimit,proto3" json:"monthly_limit,omitempty"`
}

func (x *Person_Client_Gold) Reset() {
	*x = Person_Client_Gold{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Client_Gold) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Client_Gold) ProtoMessage() {}

func (x *Person_Client_Gold) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Client_Gold.ProtoReflect.Descriptor instead.
func (*Person_Client_Gold) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 2, 1}
}

func (x *Person_Client_Gold) GetMonthlyLimit() int64 {
	if x != nil {
		return x.MonthlyLimit
	}
	return 0
}

type Person_Client_Silver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DailyLimit int64 `protobuf:"varint,3,opt,name=daily_limit,json=dailyLimit,proto3" json:"daily_limit,omitempty"`
}

func (x *Person_Client_Silver) Reset() {
	*x = Person_Client_Silver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_registry_v1_person_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person_Client_Silver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person_Client_Silver) ProtoMessage() {}

func (x *Person_Client_Silver) ProtoReflect() protoreflect.Message {
	mi := &file_registry_v1_person_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person_Client_Silver.ProtoReflect.Descriptor instead.
func (*Person_Client_Silver) Descriptor() ([]byte, []int) {
	return file_registry_v1_person_proto_rawDescGZIP(), []int{0, 2, 2}
}

func (x *Person_Client_Silver) GetDailyLimit() int64 {
	if x != nil {
		return x.DailyLimit
	}
	return 0
}

var File_registry_v1_person_proto protoreflect.FileDescriptor

var file_registry_v1_person_proto_rawDesc = []byte{
	0x0a, 0x18, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x22, 0xc3, 0x07, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x48, 0x00, 0x52,
	0x05, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x37, 0x0a, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12,
	0x34, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x06, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x1a, 0x52, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x26,
	0x0a, 0x0f, 0x79, 0x65, 0x61, 0x72, 0x73, 0x5f, 0x69, 0x6e, 0x5f, 0x63, 0x68, 0x61, 0x72, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x79, 0x65, 0x61, 0x72, 0x73, 0x49, 0x6e,
	0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x61, 0x73, 0x5f, 0x68, 0x6f,
	0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x68, 0x61,
	0x73, 0x48, 0x6f, 0x6c, 0x69, 0x64, 0x61, 0x79, 0x73, 0x1a, 0x32, 0x0a, 0x07, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x5f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x73, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73, 0x1a, 0xd2, 0x02,
	0x0a, 0x06, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x6d,
	0x69, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x48, 0x00, 0x52,
	0x07, 0x70, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x35, 0x0a, 0x04, 0x67, 0x6f, 0x6c, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x2e, 0x47, 0x6f, 0x6c, 0x64, 0x48, 0x00, 0x52, 0x04, 0x67, 0x6f, 0x6c, 0x64, 0x12,
	0x3b, 0x0a, 0x06, 0x73, 0x69, 0x6c, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x69, 0x6c, 0x76,
	0x65, 0x72, 0x48, 0x00, 0x52, 0x06, 0x73, 0x69, 0x6c, 0x76, 0x65, 0x72, 0x1a, 0x2c, 0x0a, 0x07,
	0x50, 0x72, 0x65, 0x6d, 0x69, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x79, 0x65, 0x61, 0x72, 0x6c,
	0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x79,
	0x65, 0x61, 0x72, 0x6c, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x2b, 0x0a, 0x04, 0x47, 0x6f,
	0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x6c, 0x79, 0x5f, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68,
	0x6c, 0x79, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x1a, 0x29, 0x0a, 0x06, 0x53, 0x69, 0x6c, 0x76, 0x65,
	0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x61, 0x69, 0x6c, 0x79, 0x4c, 0x69, 0x6d,
	0x69, 0x74, 0x42, 0x0e, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0xe0, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x23,
	0x0a, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x88, 0x01, 0x01, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70,
	0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x64, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0xa8, 0x01,
	0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x76,
	0x31, 0x42, 0x0b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6d, 0x61,
	0x75, 0x72, 0x79, 0x39, 0x35, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x69, 0x66, 0x79, 0x2f, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2f,
	0x76, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x76, 0x31, 0xa2, 0x02, 0x03,
	0x52, 0x58, 0x58, 0xaa, 0x02, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x2e, 0x56,
	0x31, 0xca, 0x02, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0xe2,
	0x02, 0x17, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0c, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_registry_v1_person_proto_rawDescOnce sync.Once
	file_registry_v1_person_proto_rawDescData = file_registry_v1_person_proto_rawDesc
)

func file_registry_v1_person_proto_rawDescGZIP() []byte {
	file_registry_v1_person_proto_rawDescOnce.Do(func() {
		file_registry_v1_person_proto_rawDescData = protoimpl.X.CompressGZIP(file_registry_v1_person_proto_rawDescData)
	})
	return file_registry_v1_person_proto_rawDescData
}

var file_registry_v1_person_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_registry_v1_person_proto_goTypes = []interface{}{
	(*Person)(nil),                // 0: registry.v1.Person
	(*Person_Admin)(nil),          // 1: registry.v1.Person.Admin
	(*Person_Manager)(nil),        // 2: registry.v1.Person.Manager
	(*Person_Client)(nil),         // 3: registry.v1.Person.Client
	(*Person_Address)(nil),        // 4: registry.v1.Person.Address
	(*Person_Client_Premium)(nil), // 5: registry.v1.Person.Client.Premium
	(*Person_Client_Gold)(nil),    // 6: registry.v1.Person.Client.Gold
	(*Person_Client_Silver)(nil),  // 7: registry.v1.Person.Client.Silver
}
var file_registry_v1_person_proto_depIdxs = []int32{
	4, // 0: registry.v1.Person.address:type_name -> registry.v1.Person.Address
	1, // 1: registry.v1.Person.admin:type_name -> registry.v1.Person.Admin
	2, // 2: registry.v1.Person.manager:type_name -> registry.v1.Person.Manager
	3, // 3: registry.v1.Person.client:type_name -> registry.v1.Person.Client
	5, // 4: registry.v1.Person.Client.premium:type_name -> registry.v1.Person.Client.Premium
	6, // 5: registry.v1.Person.Client.gold:type_name -> registry.v1.Person.Client.Gold
	7, // 6: registry.v1.Person.Client.silver:type_name -> registry.v1.Person.Client.Silver
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_registry_v1_person_proto_init() }
func file_registry_v1_person_proto_init() {
	if File_registry_v1_person_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_registry_v1_person_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_registry_v1_person_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Admin); i {
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
		file_registry_v1_person_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Manager); i {
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
		file_registry_v1_person_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Client); i {
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
		file_registry_v1_person_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Address); i {
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
		file_registry_v1_person_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Client_Premium); i {
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
		file_registry_v1_person_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Client_Gold); i {
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
		file_registry_v1_person_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person_Client_Silver); i {
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
	file_registry_v1_person_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Person_Admin_)(nil),
		(*Person_Manager_)(nil),
		(*Person_Client_)(nil),
	}
	file_registry_v1_person_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Person_Client_Premium_)(nil),
		(*Person_Client_Gold_)(nil),
		(*Person_Client_Silver_)(nil),
	}
	file_registry_v1_person_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_registry_v1_person_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_registry_v1_person_proto_goTypes,
		DependencyIndexes: file_registry_v1_person_proto_depIdxs,
		MessageInfos:      file_registry_v1_person_proto_msgTypes,
	}.Build()
	File_registry_v1_person_proto = out.File
	file_registry_v1_person_proto_rawDesc = nil
	file_registry_v1_person_proto_goTypes = nil
	file_registry_v1_person_proto_depIdxs = nil
}

// LoadMap loads map values into given struct.
func (e *Person) LoadMap(m map[string]interface{}) {
	e.Name = m["first_name"].(string)
	e.Address = m["address"].(*Person_Address)
	if _, ok := m["Type"].(map[string]interface{}); ok {
		if _, ok := option["Admin"].(map[string]interface{}); ok {
			// YearsInCharge
			// HasHolidays
			e.Type = &Person_Admin_{}
		}
		if _, ok := option["Manager"].(map[string]interface{}); ok {
			// ManagesClients
			e.Type = &Person_Manager_{}
		}
		if _, ok := option["Client"].(map[string]interface{}); ok {
			// Premium
			// Gold
			// Silver
			e.Type = &Person_Client_{}
		}
	}
}

// LoadMap loads map values into given struct.
func (e *Person_Admin) LoadMap(m map[string]interface{}) {
	e.YearsInCharge = m["years_in_charge"].(int64)
	e.HasHolidays = m["has_holidays"].(bool)
}

// LoadMap loads map values into given struct.
func (e *Person_Manager) LoadMap(m map[string]interface{}) {
	e.ManagesClients = m["manages_clients"].(bool)
}

// LoadMap loads map values into given struct.
func (e *Person_Client) LoadMap(m map[string]interface{}) {
	if _, ok := m["Subscription"].(map[string]interface{}); ok {
		if _, ok := option["Premium"].(map[string]interface{}); ok {
			// YearlyLimit
			e.Type = &Person_Client_Premium_{}
		}
		if _, ok := option["Gold"].(map[string]interface{}); ok {
			// MonthlyLimit
			e.Type = &Person_Client_Gold_{}
		}
		if _, ok := option["Silver"].(map[string]interface{}); ok {
			// DailyLimit
			e.Type = &Person_Client_Silver_{}
		}
	}
}

// LoadMap loads map values into given struct.
func (e *Person_Client_Premium) LoadMap(m map[string]interface{}) {
	e.YearlyLimit = m["yearly_limit"].(int64)
}

// LoadMap loads map values into given struct.
func (e *Person_Client_Gold) LoadMap(m map[string]interface{}) {
	e.MonthlyLimit = m["monthly_limit"].(int64)
}

// LoadMap loads map values into given struct.
func (e *Person_Client_Silver) LoadMap(m map[string]interface{}) {
	e.DailyLimit = m["daily_limit"].(int64)
}

// LoadMap loads map values into given struct.
func (e *Person_Address) LoadMap(m map[string]interface{}) {
	e.Street = m["street"].(string)
	e.Number = m["number"].(string)
	e.Location = m["location"].(string)
	e.Province = m["province"].(string)
	e.PostalCode = m["postal_code"].(string)
	e.Country = m["country"].(string)
}
