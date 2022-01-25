// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	ragù          v0.2.3
// source: pkg/core/core.proto

package core

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

type BootstrapToken struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenID []byte `protobuf:"bytes,1,opt,name=tokenID,proto3" json:"tokenID,omitempty"`
	Secret  []byte `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	LeaseID int64  `protobuf:"varint,3,opt,name=leaseID,proto3" json:"leaseID,omitempty"`
	Ttl     int64  `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
}

func (x *BootstrapToken) Reset() {
	*x = BootstrapToken{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapToken) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapToken) ProtoMessage() {}

func (x *BootstrapToken) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapToken.ProtoReflect.Descriptor instead.
func (*BootstrapToken) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{0}
}

func (x *BootstrapToken) GetTokenID() []byte {
	if x != nil {
		return x.TokenID
	}
	return nil
}

func (x *BootstrapToken) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *BootstrapToken) GetLeaseID() int64 {
	if x != nil {
		return x.LeaseID
	}
	return 0
}

func (x *BootstrapToken) GetTtl() int64 {
	if x != nil {
		return x.Ttl
	}
	return 0
}

type BootstrapTokenList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*BootstrapToken `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *BootstrapTokenList) Reset() {
	*x = BootstrapTokenList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BootstrapTokenList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BootstrapTokenList) ProtoMessage() {}

func (x *BootstrapTokenList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BootstrapTokenList.ProtoReflect.Descriptor instead.
func (*BootstrapTokenList) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{1}
}

func (x *BootstrapTokenList) GetItems() []*BootstrapToken {
	if x != nil {
		return x.Items
	}
	return nil
}

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Labels map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{2}
}

func (x *Cluster) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cluster) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type ClusterList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Cluster `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ClusterList) Reset() {
	*x = ClusterList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClusterList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterList) ProtoMessage() {}

func (x *ClusterList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterList.ProtoReflect.Descriptor instead.
func (*ClusterList) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{3}
}

func (x *ClusterList) GetItems() []*Cluster {
	if x != nil {
		return x.Items
	}
	return nil
}

type LabelSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Labels map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *LabelSelector) Reset() {
	*x = LabelSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSelector) ProtoMessage() {}

func (x *LabelSelector) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSelector.ProtoReflect.Descriptor instead.
func (*LabelSelector) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{4}
}

func (x *LabelSelector) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClusterIDs  []string       `protobuf:"bytes,2,rep,name=clusterIDs,proto3" json:"clusterIDs,omitempty"`
	MatchLabels *LabelSelector `protobuf:"bytes,3,opt,name=matchLabels,proto3" json:"matchLabels,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{5}
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetClusterIDs() []string {
	if x != nil {
		return x.ClusterIDs
	}
	return nil
}

func (x *Role) GetMatchLabels() *LabelSelector {
	if x != nil {
		return x.MatchLabels
	}
	return nil
}

type RoleBinding struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	RoleName string   `protobuf:"bytes,2,opt,name=roleName,proto3" json:"roleName,omitempty"`
	Subjects []string `protobuf:"bytes,3,rep,name=subjects,proto3" json:"subjects,omitempty"`
}

func (x *RoleBinding) Reset() {
	*x = RoleBinding{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBinding) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBinding) ProtoMessage() {}

func (x *RoleBinding) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBinding.ProtoReflect.Descriptor instead.
func (*RoleBinding) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{6}
}

func (x *RoleBinding) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RoleBinding) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *RoleBinding) GetSubjects() []string {
	if x != nil {
		return x.Subjects
	}
	return nil
}

type RoleList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*Role `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleList) Reset() {
	*x = RoleList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleList) ProtoMessage() {}

func (x *RoleList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleList.ProtoReflect.Descriptor instead.
func (*RoleList) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{7}
}

func (x *RoleList) GetItems() []*Role {
	if x != nil {
		return x.Items
	}
	return nil
}

type RoleBindingList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*RoleBinding `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *RoleBindingList) Reset() {
	*x = RoleBindingList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleBindingList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleBindingList) ProtoMessage() {}

func (x *RoleBindingList) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleBindingList.ProtoReflect.Descriptor instead.
func (*RoleBindingList) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{8}
}

func (x *RoleBindingList) GetItems() []*RoleBinding {
	if x != nil {
		return x.Items
	}
	return nil
}

type CertInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Issuer      string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Subject     string `protobuf:"bytes,2,opt,name=subject,proto3" json:"subject,omitempty"`
	IsCA        bool   `protobuf:"varint,3,opt,name=isCA,proto3" json:"isCA,omitempty"`
	NotBefore   string `protobuf:"bytes,4,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter    string `protobuf:"bytes,5,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Fingerprint string `protobuf:"bytes,6,opt,name=fingerprint,proto3" json:"fingerprint,omitempty"`
}

func (x *CertInfo) Reset() {
	*x = CertInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CertInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CertInfo) ProtoMessage() {}

func (x *CertInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CertInfo.ProtoReflect.Descriptor instead.
func (*CertInfo) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{9}
}

func (x *CertInfo) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *CertInfo) GetSubject() string {
	if x != nil {
		return x.Subject
	}
	return ""
}

func (x *CertInfo) GetIsCA() bool {
	if x != nil {
		return x.IsCA
	}
	return false
}

func (x *CertInfo) GetNotBefore() string {
	if x != nil {
		return x.NotBefore
	}
	return ""
}

func (x *CertInfo) GetNotAfter() string {
	if x != nil {
		return x.NotAfter
	}
	return ""
}

func (x *CertInfo) GetFingerprint() string {
	if x != nil {
		return x.Fingerprint
	}
	return ""
}

type Reference struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Id   string `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Reference) Reset() {
	*x = Reference{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_core_core_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Reference) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reference) ProtoMessage() {}

func (x *Reference) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_core_core_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reference.ProtoReflect.Descriptor instead.
func (*Reference) Descriptor() ([]byte, []int) {
	return file_pkg_core_core_proto_rawDescGZIP(), []int{10}
}

func (x *Reference) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Reference) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_pkg_core_core_proto protoreflect.FileDescriptor

var file_pkg_core_core_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x59, 0x0a, 0x0e, 0x42,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x11, 0x0a,
	0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x00,
	0x12, 0x10, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x42, 0x00, 0x12, 0x11, 0x0a, 0x07, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x00, 0x12, 0x0d, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x3d, 0x0a, 0x12, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74,
	0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x42, 0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x7f, 0x0a, 0x07, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x0c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x29,
	0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x3a, 0x00, 0x22, 0x2f, 0x0a, 0x0b, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x7d, 0x0a, 0x0d, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x2f, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x3a, 0x00, 0x22, 0x5a, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x14,
	0x0a, 0x0a, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x09, 0x42, 0x00, 0x12, 0x2a, 0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x42, 0x00,
	0x3a, 0x00, 0x22, 0x47, 0x0a, 0x0b, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x12, 0x0e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x00, 0x12, 0x12, 0x0a, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x12, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x29, 0x0a, 0x08, 0x52,
	0x6f, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x52, 0x6f,
	0x6c, 0x65, 0x42, 0x00, 0x3a, 0x00, 0x22, 0x37, 0x0a, 0x0f, 0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x42, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x00, 0x3a, 0x00, 0x22,
	0x81, 0x01, 0x0a, 0x08, 0x43, 0x65, 0x72, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x11,
	0x0a, 0x07, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x00, 0x12, 0x0e, 0x0a, 0x04, 0x69, 0x73, 0x43, 0x41, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x42,
	0x00, 0x12, 0x13, 0x0a, 0x09, 0x6e, 0x6f, 0x74, 0x42, 0x65, 0x66, 0x6f, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x12, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x41, 0x66, 0x74,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x12, 0x15, 0x0a, 0x0b, 0x66, 0x69,
	0x6e, 0x67, 0x65, 0x72, 0x70, 0x72, 0x69, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x00, 0x3a, 0x00, 0x22, 0x2b, 0x0a, 0x09, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x12, 0x0e, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00,
	0x12, 0x0c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x00, 0x3a, 0x00,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x72, 0x61, 0x6c, 0x69, 0x63, 0x6b, 0x79, 0x2f, 0x6f, 0x70, 0x6e, 0x69, 0x2d, 0x6d, 0x6f, 0x6e,
	0x69, 0x74, 0x6f, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_core_core_proto_rawDescOnce sync.Once
	file_pkg_core_core_proto_rawDescData = file_pkg_core_core_proto_rawDesc
)

func file_pkg_core_core_proto_rawDescGZIP() []byte {
	file_pkg_core_core_proto_rawDescOnce.Do(func() {
		file_pkg_core_core_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_core_core_proto_rawDescData)
	})
	return file_pkg_core_core_proto_rawDescData
}

var file_pkg_core_core_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_core_core_proto_goTypes = []interface{}{
	(*BootstrapToken)(nil),     // 0: core.BootstrapToken
	(*BootstrapTokenList)(nil), // 1: core.BootstrapTokenList
	(*Cluster)(nil),            // 2: core.Cluster
	(*ClusterList)(nil),        // 3: core.ClusterList
	(*LabelSelector)(nil),      // 4: core.LabelSelector
	(*Role)(nil),               // 5: core.Role
	(*RoleBinding)(nil),        // 6: core.RoleBinding
	(*RoleList)(nil),           // 7: core.RoleList
	(*RoleBindingList)(nil),    // 8: core.RoleBindingList
	(*CertInfo)(nil),           // 9: core.CertInfo
	(*Reference)(nil),          // 10: core.Reference
	nil,                        // 11: core.Cluster.LabelsEntry
	nil,                        // 12: core.LabelSelector.LabelsEntry
}
var file_pkg_core_core_proto_depIdxs = []int32{
	0,  // 0: core.BootstrapTokenList.items:type_name -> core.BootstrapToken
	11, // 1: core.Cluster.labels:type_name -> core.Cluster.LabelsEntry
	2,  // 2: core.ClusterList.items:type_name -> core.Cluster
	12, // 3: core.LabelSelector.labels:type_name -> core.LabelSelector.LabelsEntry
	4,  // 4: core.Role.matchLabels:type_name -> core.LabelSelector
	5,  // 5: core.RoleList.items:type_name -> core.Role
	6,  // 6: core.RoleBindingList.items:type_name -> core.RoleBinding
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_pkg_core_core_proto_init() }
func file_pkg_core_core_proto_init() {
	if File_pkg_core_core_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_core_core_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapToken); i {
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
		file_pkg_core_core_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BootstrapTokenList); i {
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
		file_pkg_core_core_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
		file_pkg_core_core_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClusterList); i {
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
		file_pkg_core_core_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSelector); i {
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
		file_pkg_core_core_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_pkg_core_core_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBinding); i {
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
		file_pkg_core_core_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleList); i {
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
		file_pkg_core_core_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleBindingList); i {
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
		file_pkg_core_core_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CertInfo); i {
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
		file_pkg_core_core_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Reference); i {
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
			RawDescriptor: file_pkg_core_core_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_core_core_proto_goTypes,
		DependencyIndexes: file_pkg_core_core_proto_depIdxs,
		MessageInfos:      file_pkg_core_core_proto_msgTypes,
	}.Build()
	File_pkg_core_core_proto = out.File
	file_pkg_core_core_proto_rawDesc = nil
	file_pkg_core_core_proto_goTypes = nil
	file_pkg_core_core_proto_depIdxs = nil
}
