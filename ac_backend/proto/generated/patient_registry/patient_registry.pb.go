// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: patient_registry.proto

package patient_registry

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Gender int32

const (
	Gender_GENDER_NONE    Gender = 0
	Gender_GENDER_MALE    Gender = 1
	Gender_GENDER_FEMALE  Gender = 2
	Gender_GENDER_OTHER   Gender = 3
	Gender_GENDER_UNKNOWN Gender = 4
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_NONE",
		1: "GENDER_MALE",
		2: "GENDER_FEMALE",
		3: "GENDER_OTHER",
		4: "GENDER_UNKNOWN",
	}
	Gender_value = map[string]int32{
		"GENDER_NONE":    0,
		"GENDER_MALE":    1,
		"GENDER_FEMALE":  2,
		"GENDER_OTHER":   3,
		"GENDER_UNKNOWN": 4,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_patient_registry_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_patient_registry_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{0}
}

type EmergencyContactRelationship int32

const (
	EmergencyContactRelationship_RELATIONSHIP_FRIEND EmergencyContactRelationship = 0
	EmergencyContactRelationship_RELATIONSHIP_FAMILY EmergencyContactRelationship = 1
)

// Enum value maps for EmergencyContactRelationship.
var (
	EmergencyContactRelationship_name = map[int32]string{
		0: "RELATIONSHIP_FRIEND",
		1: "RELATIONSHIP_FAMILY",
	}
	EmergencyContactRelationship_value = map[string]int32{
		"RELATIONSHIP_FRIEND": 0,
		"RELATIONSHIP_FAMILY": 1,
	}
)

func (x EmergencyContactRelationship) Enum() *EmergencyContactRelationship {
	p := new(EmergencyContactRelationship)
	*p = x
	return p
}

func (x EmergencyContactRelationship) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmergencyContactRelationship) Descriptor() protoreflect.EnumDescriptor {
	return file_patient_registry_proto_enumTypes[1].Descriptor()
}

func (EmergencyContactRelationship) Type() protoreflect.EnumType {
	return &file_patient_registry_proto_enumTypes[1]
}

func (x EmergencyContactRelationship) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmergencyContactRelationship.Descriptor instead.
func (EmergencyContactRelationship) EnumDescriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{1}
}

type ReferredBy int32

const (
	ReferredBy_REFERRAL_SEARCH ReferredBy = 0
	ReferredBy_REFERRAL_FRIEND ReferredBy = 1
)

// Enum value maps for ReferredBy.
var (
	ReferredBy_name = map[int32]string{
		0: "REFERRAL_SEARCH",
		1: "REFERRAL_FRIEND",
	}
	ReferredBy_value = map[string]int32{
		"REFERRAL_SEARCH": 0,
		"REFERRAL_FRIEND": 1,
	}
)

func (x ReferredBy) Enum() *ReferredBy {
	p := new(ReferredBy)
	*p = x
	return p
}

func (x ReferredBy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReferredBy) Descriptor() protoreflect.EnumDescriptor {
	return file_patient_registry_proto_enumTypes[2].Descriptor()
}

func (ReferredBy) Type() protoreflect.EnumType {
	return &file_patient_registry_proto_enumTypes[2]
}

func (x ReferredBy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReferredBy.Descriptor instead.
func (ReferredBy) EnumDescriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{2}
}

type Status int32

const (
	Status_STATUS_ACTIVE            Status = 0
	Status_STATUS_DELETED           Status = 1
	Status_STATUS_PENDING_INSURANCE Status = 2
	Status_STATUS_HOME_HEALTH       Status = 3
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_ACTIVE",
		1: "STATUS_DELETED",
		2: "STATUS_PENDING_INSURANCE",
		3: "STATUS_HOME_HEALTH",
	}
	Status_value = map[string]int32{
		"STATUS_ACTIVE":            0,
		"STATUS_DELETED":           1,
		"STATUS_PENDING_INSURANCE": 2,
		"STATUS_HOME_HEALTH":       3,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_patient_registry_proto_enumTypes[3].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_patient_registry_proto_enumTypes[3]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{3}
}

type CreatPatientRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *CreatPatientRequest) Reset() {
	*x = CreatPatientRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_registry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatPatientRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatPatientRequest) ProtoMessage() {}

func (x *CreatPatientRequest) ProtoReflect() protoreflect.Message {
	mi := &file_patient_registry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatPatientRequest.ProtoReflect.Descriptor instead.
func (*CreatPatientRequest) Descriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{0}
}

func (x *CreatPatientRequest) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

type PatientEmergencyContact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName    string                       `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName     string                       `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	PhoneNumber  string                       `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Relationship EmergencyContactRelationship `protobuf:"varint,4,opt,name=relationship,proto3,enum=api.v1.EmergencyContactRelationship" json:"relationship,omitempty"`
}

func (x *PatientEmergencyContact) Reset() {
	*x = PatientEmergencyContact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_registry_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PatientEmergencyContact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PatientEmergencyContact) ProtoMessage() {}

func (x *PatientEmergencyContact) ProtoReflect() protoreflect.Message {
	mi := &file_patient_registry_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PatientEmergencyContact.ProtoReflect.Descriptor instead.
func (*PatientEmergencyContact) Descriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{1}
}

func (x *PatientEmergencyContact) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *PatientEmergencyContact) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *PatientEmergencyContact) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *PatientEmergencyContact) GetRelationship() EmergencyContactRelationship {
	if x != nil {
		return x.Relationship
	}
	return EmergencyContactRelationship_RELATIONSHIP_FRIEND
}

type Referral struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReferredBy   ReferredBy `protobuf:"varint,1,opt,name=referredBy,proto3,enum=api.v1.ReferredBy" json:"referredBy,omitempty"`
	ReferralDate string     `protobuf:"bytes,2,opt,name=referralDate,proto3" json:"referralDate,omitempty"`
	Comment      string     `protobuf:"bytes,3,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *Referral) Reset() {
	*x = Referral{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_registry_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Referral) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Referral) ProtoMessage() {}

func (x *Referral) ProtoReflect() protoreflect.Message {
	mi := &file_patient_registry_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Referral.ProtoReflect.Descriptor instead.
func (*Referral) Descriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{2}
}

func (x *Referral) GetReferredBy() ReferredBy {
	if x != nil {
		return x.ReferredBy
	}
	return ReferredBy_REFERRAL_SEARCH
}

func (x *Referral) GetReferralDate() string {
	if x != nil {
		return x.ReferralDate
	}
	return ""
}

func (x *Referral) GetComment() string {
	if x != nil {
		return x.Comment
	}
	return ""
}

type Patient struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                      string                   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName               string                   `protobuf:"bytes,2,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName                string                   `protobuf:"bytes,3,opt,name=lastName,proto3" json:"lastName,omitempty"`
	SSN                     string                   `protobuf:"bytes,4,opt,name=SSN,proto3" json:"SSN,omitempty"`
	Gender                  Gender                   `protobuf:"varint,5,opt,name=gender,proto3,enum=api.v1.Gender" json:"gender,omitempty"`
	YearOfBirth             int32                    `protobuf:"varint,6,opt,name=yearOfBirth,proto3" json:"yearOfBirth,omitempty"`
	MonthOfBirth            int32                    `protobuf:"varint,7,opt,name=monthOfBirth,proto3" json:"monthOfBirth,omitempty"`
	DayOfBirth              int32                    `protobuf:"varint,8,opt,name=dayOfBirth,proto3" json:"dayOfBirth,omitempty"`
	Status                  Status                   `protobuf:"varint,9,opt,name=status,proto3,enum=api.v1.Status" json:"status,omitempty"`
	PrimaryCarePhysician    string                   `protobuf:"bytes,10,opt,name=primaryCarePhysician,proto3" json:"primaryCarePhysician,omitempty"`
	PhoneWork               string                   `protobuf:"bytes,11,opt,name=phoneWork,proto3" json:"phoneWork,omitempty"`
	PhoneHome               string                   `protobuf:"bytes,12,opt,name=phoneHome,proto3" json:"phoneHome,omitempty"`
	PhoneMobile             string                   `protobuf:"bytes,13,opt,name=phoneMobile,proto3" json:"phoneMobile,omitempty"`
	AddressStreet           string                   `protobuf:"bytes,14,opt,name=addressStreet,proto3" json:"addressStreet,omitempty"`
	AddressCity             string                   `protobuf:"bytes,15,opt,name=addressCity,proto3" json:"addressCity,omitempty"`
	AddressState            string                   `protobuf:"bytes,16,opt,name=addressState,proto3" json:"addressState,omitempty"`
	AddressZip              string                   `protobuf:"bytes,17,opt,name=addressZip,proto3" json:"addressZip,omitempty"`
	WeightInKg              float64                  `protobuf:"fixed64,18,opt,name=weightInKg,proto3" json:"weightInKg,omitempty"` // in kg
	HeightInCm              float64                  `protobuf:"fixed64,19,opt,name=heightInCm,proto3" json:"heightInCm,omitempty"` // in centimeter
	Ethnicity               string                   `protobuf:"bytes,20,opt,name=ethnicity,proto3" json:"ethnicity,omitempty"`
	Email                   string                   `protobuf:"bytes,21,opt,name=email,proto3" json:"email,omitempty"`
	PatientEmergencyContact *PatientEmergencyContact `protobuf:"bytes,22,opt,name=patientEmergencyContact,proto3" json:"patientEmergencyContact,omitempty"`
	Referral                *Referral                `protobuf:"bytes,23,opt,name=referral,proto3" json:"referral,omitempty"`
	CreatedAt               int64                    `protobuf:"varint,24,opt,name=createdAt,proto3" json:"createdAt,omitempty"` // in millisecond
	UpdatedAt               int64                    `protobuf:"varint,25,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"` // in millisecond
}

func (x *Patient) Reset() {
	*x = Patient{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_registry_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Patient) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Patient) ProtoMessage() {}

func (x *Patient) ProtoReflect() protoreflect.Message {
	mi := &file_patient_registry_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Patient.ProtoReflect.Descriptor instead.
func (*Patient) Descriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{3}
}

func (x *Patient) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Patient) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Patient) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Patient) GetSSN() string {
	if x != nil {
		return x.SSN
	}
	return ""
}

func (x *Patient) GetGender() Gender {
	if x != nil {
		return x.Gender
	}
	return Gender_GENDER_NONE
}

func (x *Patient) GetYearOfBirth() int32 {
	if x != nil {
		return x.YearOfBirth
	}
	return 0
}

func (x *Patient) GetMonthOfBirth() int32 {
	if x != nil {
		return x.MonthOfBirth
	}
	return 0
}

func (x *Patient) GetDayOfBirth() int32 {
	if x != nil {
		return x.DayOfBirth
	}
	return 0
}

func (x *Patient) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_ACTIVE
}

func (x *Patient) GetPrimaryCarePhysician() string {
	if x != nil {
		return x.PrimaryCarePhysician
	}
	return ""
}

func (x *Patient) GetPhoneWork() string {
	if x != nil {
		return x.PhoneWork
	}
	return ""
}

func (x *Patient) GetPhoneHome() string {
	if x != nil {
		return x.PhoneHome
	}
	return ""
}

func (x *Patient) GetPhoneMobile() string {
	if x != nil {
		return x.PhoneMobile
	}
	return ""
}

func (x *Patient) GetAddressStreet() string {
	if x != nil {
		return x.AddressStreet
	}
	return ""
}

func (x *Patient) GetAddressCity() string {
	if x != nil {
		return x.AddressCity
	}
	return ""
}

func (x *Patient) GetAddressState() string {
	if x != nil {
		return x.AddressState
	}
	return ""
}

func (x *Patient) GetAddressZip() string {
	if x != nil {
		return x.AddressZip
	}
	return ""
}

func (x *Patient) GetWeightInKg() float64 {
	if x != nil {
		return x.WeightInKg
	}
	return 0
}

func (x *Patient) GetHeightInCm() float64 {
	if x != nil {
		return x.HeightInCm
	}
	return 0
}

func (x *Patient) GetEthnicity() string {
	if x != nil {
		return x.Ethnicity
	}
	return ""
}

func (x *Patient) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Patient) GetPatientEmergencyContact() *PatientEmergencyContact {
	if x != nil {
		return x.PatientEmergencyContact
	}
	return nil
}

func (x *Patient) GetReferral() *Referral {
	if x != nil {
		return x.Referral
	}
	return nil
}

func (x *Patient) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Patient) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type CreatPatientResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Patient *Patient `protobuf:"bytes,1,opt,name=patient,proto3" json:"patient,omitempty"`
}

func (x *CreatPatientResponse) Reset() {
	*x = CreatPatientResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_patient_registry_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatPatientResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatPatientResponse) ProtoMessage() {}

func (x *CreatPatientResponse) ProtoReflect() protoreflect.Message {
	mi := &file_patient_registry_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatPatientResponse.ProtoReflect.Descriptor instead.
func (*CreatPatientResponse) Descriptor() ([]byte, []int) {
	return file_patient_registry_proto_rawDescGZIP(), []int{4}
}

func (x *CreatPatientResponse) GetPatient() *Patient {
	if x != nil {
		return x.Patient
	}
	return nil
}

var File_patient_registry_proto protoreflect.FileDescriptor

var file_patient_registry_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x22, 0x40, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65,
	0x6e, 0x74, 0x22, 0xbf, 0x01, 0x0a, 0x17, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d,
	0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x48, 0x0a, 0x0c, 0x72, 0x65,
	0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x24, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65,
	0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x52, 0x0c, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x68, 0x69, 0x70, 0x22, 0x7c, 0x0a, 0x08, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c,
	0x12, 0x32, 0x0a, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x52, 0x0a, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72,
	0x65, 0x64, 0x42, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x72, 0x61, 0x6c, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x22, 0xf2, 0x06, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x53, 0x4e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x53, 0x4e, 0x12, 0x26, 0x0a, 0x06, 0x67, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x79, 0x65, 0x61, 0x72, 0x4f, 0x66, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x4f, 0x66, 0x42,
	0x69, 0x72, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6d, 0x6f, 0x6e, 0x74,
	0x68, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x61, 0x79, 0x4f,
	0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x64, 0x61,
	0x79, 0x4f, 0x66, 0x42, 0x69, 0x72, 0x74, 0x68, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x32, 0x0a, 0x14, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x61, 0x72, 0x65, 0x50,
	0x68, 0x79, 0x73, 0x69, 0x63, 0x69, 0x61, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14,
	0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x43, 0x61, 0x72, 0x65, 0x50, 0x68, 0x79, 0x73, 0x69,
	0x63, 0x69, 0x61, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x57, 0x6f,
	0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x48, 0x6f, 0x6d, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x48, 0x6f, 0x6d, 0x65,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4d, 0x6f, 0x62, 0x69,
	0x6c, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x43, 0x69, 0x74, 0x79, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x43, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5a, 0x69, 0x70, 0x18, 0x11, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x5a, 0x69, 0x70, 0x12, 0x1e,
	0x0a, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x4b, 0x67, 0x18, 0x12, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x4b, 0x67, 0x12, 0x1e,
	0x0a, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x43, 0x6d, 0x18, 0x13, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x0a, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x49, 0x6e, 0x43, 0x6d, 0x12, 0x1c,
	0x0a, 0x09, 0x65, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x18, 0x14, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x65, 0x74, 0x68, 0x6e, 0x69, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x59, 0x0a, 0x17, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x74,
	0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e,
	0x74, 0x61, 0x63, 0x74, 0x52, 0x17, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x45, 0x6d, 0x65,
	0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2c, 0x0a,
	0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x18, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61,
	0x6c, 0x52, 0x08, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x61, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x18, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x19, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x41, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x07, 0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x2a, 0x63, 0x0a, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4e,
	0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f,
	0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52,
	0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x47, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x03, 0x12, 0x12, 0x0a, 0x0e, 0x47,
	0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x04, 0x2a,
	0x50, 0x0a, 0x1c, 0x45, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x6e, 0x63, 0x79, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x68, 0x69, 0x70, 0x12,
	0x17, 0x0a, 0x13, 0x52, 0x45, 0x4c, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f,
	0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45, 0x4c, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x53, 0x48, 0x49, 0x50, 0x5f, 0x46, 0x41, 0x4d, 0x49, 0x4c, 0x59, 0x10,
	0x01, 0x2a, 0x36, 0x0a, 0x0a, 0x52, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x42, 0x79, 0x12,
	0x13, 0x0a, 0x0f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x41, 0x52,
	0x43, 0x48, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x45, 0x46, 0x45, 0x52, 0x52, 0x41, 0x4c,
	0x5f, 0x46, 0x52, 0x49, 0x45, 0x4e, 0x44, 0x10, 0x01, 0x2a, 0x65, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
	0x5f, 0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x49, 0x4e, 0x53,
	0x55, 0x52, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54,
	0x55, 0x53, 0x5f, 0x48, 0x4f, 0x4d, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x4c, 0x54, 0x48, 0x10, 0x03,
	0x32, 0x64, 0x0a, 0x16, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x50, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x58, 0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x63, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x61, 0x74, 0x69,
	0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x42, 0x14, 0x50, 0x61,
	0x74, 0x69, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f,
	0x70, 0x61, 0x74, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_patient_registry_proto_rawDescOnce sync.Once
	file_patient_registry_proto_rawDescData = file_patient_registry_proto_rawDesc
)

func file_patient_registry_proto_rawDescGZIP() []byte {
	file_patient_registry_proto_rawDescOnce.Do(func() {
		file_patient_registry_proto_rawDescData = protoimpl.X.CompressGZIP(file_patient_registry_proto_rawDescData)
	})
	return file_patient_registry_proto_rawDescData
}

var file_patient_registry_proto_enumTypes = make([]protoimpl.EnumInfo, 4)
var file_patient_registry_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_patient_registry_proto_goTypes = []interface{}{
	(Gender)(0),                       // 0: api.v1.Gender
	(EmergencyContactRelationship)(0), // 1: api.v1.EmergencyContactRelationship
	(ReferredBy)(0),                   // 2: api.v1.ReferredBy
	(Status)(0),                       // 3: api.v1.Status
	(*CreatPatientRequest)(nil),       // 4: api.v1.CreatPatientRequest
	(*PatientEmergencyContact)(nil),   // 5: api.v1.PatientEmergencyContact
	(*Referral)(nil),                  // 6: api.v1.Referral
	(*Patient)(nil),                   // 7: api.v1.patient
	(*CreatPatientResponse)(nil),      // 8: api.v1.CreatPatientResponse
}
var file_patient_registry_proto_depIdxs = []int32{
	7, // 0: api.v1.CreatPatientRequest.patient:type_name -> api.v1.patient
	1, // 1: api.v1.PatientEmergencyContact.relationship:type_name -> api.v1.EmergencyContactRelationship
	2, // 2: api.v1.Referral.referredBy:type_name -> api.v1.ReferredBy
	0, // 3: api.v1.patient.gender:type_name -> api.v1.Gender
	3, // 4: api.v1.patient.status:type_name -> api.v1.Status
	5, // 5: api.v1.patient.patientEmergencyContact:type_name -> api.v1.PatientEmergencyContact
	6, // 6: api.v1.patient.referral:type_name -> api.v1.Referral
	7, // 7: api.v1.CreatPatientResponse.patient:type_name -> api.v1.patient
	4, // 8: api.v1.PatientRegistryService.CreatePatient:input_type -> api.v1.CreatPatientRequest
	8, // 9: api.v1.PatientRegistryService.CreatePatient:output_type -> api.v1.CreatPatientResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_patient_registry_proto_init() }
func file_patient_registry_proto_init() {
	if File_patient_registry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_patient_registry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatPatientRequest); i {
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
		file_patient_registry_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PatientEmergencyContact); i {
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
		file_patient_registry_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Referral); i {
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
		file_patient_registry_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Patient); i {
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
		file_patient_registry_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatPatientResponse); i {
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
			RawDescriptor: file_patient_registry_proto_rawDesc,
			NumEnums:      4,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_patient_registry_proto_goTypes,
		DependencyIndexes: file_patient_registry_proto_depIdxs,
		EnumInfos:         file_patient_registry_proto_enumTypes,
		MessageInfos:      file_patient_registry_proto_msgTypes,
	}.Build()
	File_patient_registry_proto = out.File
	file_patient_registry_proto_rawDesc = nil
	file_patient_registry_proto_goTypes = nil
	file_patient_registry_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PatientRegistryServiceClient is the client API for PatientRegistryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PatientRegistryServiceClient interface {
	CreatePatient(ctx context.Context, in *CreatPatientRequest, opts ...grpc.CallOption) (*CreatPatientResponse, error)
}

type patientRegistryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientRegistryServiceClient(cc grpc.ClientConnInterface) PatientRegistryServiceClient {
	return &patientRegistryServiceClient{cc}
}

func (c *patientRegistryServiceClient) CreatePatient(ctx context.Context, in *CreatPatientRequest, opts ...grpc.CallOption) (*CreatPatientResponse, error) {
	out := new(CreatPatientResponse)
	err := c.cc.Invoke(ctx, "/api.v1.PatientRegistryService/CreatePatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientRegistryServiceServer is the server API for PatientRegistryService service.
type PatientRegistryServiceServer interface {
	CreatePatient(context.Context, *CreatPatientRequest) (*CreatPatientResponse, error)
}

// UnimplementedPatientRegistryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPatientRegistryServiceServer struct {
}

func (*UnimplementedPatientRegistryServiceServer) CreatePatient(context.Context, *CreatPatientRequest) (*CreatPatientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePatient not implemented")
}

func RegisterPatientRegistryServiceServer(s *grpc.Server, srv PatientRegistryServiceServer) {
	s.RegisterService(&_PatientRegistryService_serviceDesc, srv)
}

func _PatientRegistryService_CreatePatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientRegistryServiceServer).CreatePatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.v1.PatientRegistryService/CreatePatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientRegistryServiceServer).CreatePatient(ctx, req.(*CreatPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PatientRegistryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.PatientRegistryService",
	HandlerType: (*PatientRegistryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePatient",
			Handler:    _PatientRegistryService_CreatePatient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "patient_registry.proto",
}
