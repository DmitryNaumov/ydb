// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.0
// source: ydb/library/yql/providers/generic/connector/api/common/data_source.proto

package common

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

// EDataSourceKind enumerates the external data sources
// supported in the federated query system
type EDataSourceKind int32

const (
	EDataSourceKind_DATA_SOURCE_KIND_UNSPECIFIED EDataSourceKind = 0
	EDataSourceKind_CLICKHOUSE                   EDataSourceKind = 1
	EDataSourceKind_POSTGRESQL                   EDataSourceKind = 2
)

// Enum value maps for EDataSourceKind.
var (
	EDataSourceKind_name = map[int32]string{
		0: "DATA_SOURCE_KIND_UNSPECIFIED",
		1: "CLICKHOUSE",
		2: "POSTGRESQL",
	}
	EDataSourceKind_value = map[string]int32{
		"DATA_SOURCE_KIND_UNSPECIFIED": 0,
		"CLICKHOUSE":                   1,
		"POSTGRESQL":                   2,
	}
)

func (x EDataSourceKind) Enum() *EDataSourceKind {
	p := new(EDataSourceKind)
	*p = x
	return p
}

func (x EDataSourceKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EDataSourceKind) Descriptor() protoreflect.EnumDescriptor {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes[0].Descriptor()
}

func (EDataSourceKind) Type() protoreflect.EnumType {
	return &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes[0]
}

func (x EDataSourceKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EDataSourceKind.Descriptor instead.
func (EDataSourceKind) EnumDescriptor() ([]byte, []int) {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP(), []int{0}
}

// EProtocol generalizes various kinds of network protocols supported by different databases.
type EProtocol int32

const (
	EProtocol_PROTOCOL_UNSPECIFIED EProtocol = 0
	EProtocol_NATIVE               EProtocol = 1 // CLICKHOUSE, POSTGRESQL
	EProtocol_HTTP                 EProtocol = 2 // CLICKHOUSE
)

// Enum value maps for EProtocol.
var (
	EProtocol_name = map[int32]string{
		0: "PROTOCOL_UNSPECIFIED",
		1: "NATIVE",
		2: "HTTP",
	}
	EProtocol_value = map[string]int32{
		"PROTOCOL_UNSPECIFIED": 0,
		"NATIVE":               1,
		"HTTP":                 2,
	}
)

func (x EProtocol) Enum() *EProtocol {
	p := new(EProtocol)
	*p = x
	return p
}

func (x EProtocol) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EProtocol) Descriptor() protoreflect.EnumDescriptor {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes[1].Descriptor()
}

func (EProtocol) Type() protoreflect.EnumType {
	return &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes[1]
}

func (x EProtocol) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EProtocol.Descriptor instead.
func (EProtocol) EnumDescriptor() ([]byte, []int) {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP(), []int{1}
}

// TCredentials represents various ways of user authentication in the data source instance
type TCredentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Payload:
	//	*TCredentials_Basic
	Payload isTCredentials_Payload `protobuf_oneof:"payload"`
}

func (x *TCredentials) Reset() {
	*x = TCredentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCredentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCredentials) ProtoMessage() {}

func (x *TCredentials) ProtoReflect() protoreflect.Message {
	mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCredentials.ProtoReflect.Descriptor instead.
func (*TCredentials) Descriptor() ([]byte, []int) {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP(), []int{0}
}

func (m *TCredentials) GetPayload() isTCredentials_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (x *TCredentials) GetBasic() *TCredentials_TBasic {
	if x, ok := x.GetPayload().(*TCredentials_Basic); ok {
		return x.Basic
	}
	return nil
}

type isTCredentials_Payload interface {
	isTCredentials_Payload()
}

type TCredentials_Basic struct {
	Basic *TCredentials_TBasic `protobuf:"bytes,1,opt,name=basic,proto3,oneof"`
}

func (*TCredentials_Basic) isTCredentials_Payload() {}

// TDataSourceInstance helps to identify the instance of a data source to redirect request to.
type TDataSourceInstance struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Data source kind
	Kind EDataSourceKind `protobuf:"varint,1,opt,name=kind,proto3,enum=NYql.NConnector.NApi.EDataSourceKind" json:"kind,omitempty"`
	// Network address
	Endpoint *TEndpoint `protobuf:"bytes,2,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	// Database name
	Database string `protobuf:"bytes,3,opt,name=database,proto3" json:"database,omitempty"`
	// Credentials to access database
	Credentials *TCredentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// If true, Connector server will use secure connections to access remote data sources.
	// Certificates will be obtained from the standard system paths.
	UseTls bool `protobuf:"varint,5,opt,name=use_tls,json=useTls,proto3" json:"use_tls,omitempty"`
	// Allows to specify network protocol that should be used between
	// during the connection between Connector and the remote data source
	Protocol EProtocol `protobuf:"varint,6,opt,name=protocol,proto3,enum=NYql.NConnector.NApi.EProtocol" json:"protocol,omitempty"`
}

func (x *TDataSourceInstance) Reset() {
	*x = TDataSourceInstance{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TDataSourceInstance) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TDataSourceInstance) ProtoMessage() {}

func (x *TDataSourceInstance) ProtoReflect() protoreflect.Message {
	mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TDataSourceInstance.ProtoReflect.Descriptor instead.
func (*TDataSourceInstance) Descriptor() ([]byte, []int) {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *TDataSourceInstance) GetKind() EDataSourceKind {
	if x != nil {
		return x.Kind
	}
	return EDataSourceKind_DATA_SOURCE_KIND_UNSPECIFIED
}

func (x *TDataSourceInstance) GetEndpoint() *TEndpoint {
	if x != nil {
		return x.Endpoint
	}
	return nil
}

func (x *TDataSourceInstance) GetDatabase() string {
	if x != nil {
		return x.Database
	}
	return ""
}

func (x *TDataSourceInstance) GetCredentials() *TCredentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *TDataSourceInstance) GetUseTls() bool {
	if x != nil {
		return x.UseTls
	}
	return false
}

func (x *TDataSourceInstance) GetProtocol() EProtocol {
	if x != nil {
		return x.Protocol
	}
	return EProtocol_PROTOCOL_UNSPECIFIED
}

type TCredentials_TBasic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *TCredentials_TBasic) Reset() {
	*x = TCredentials_TBasic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TCredentials_TBasic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TCredentials_TBasic) ProtoMessage() {}

func (x *TCredentials_TBasic) ProtoReflect() protoreflect.Message {
	mi := &file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TCredentials_TBasic.ProtoReflect.Descriptor instead.
func (*TCredentials_TBasic) Descriptor() ([]byte, []int) {
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TCredentials_TBasic) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TCredentials_TBasic) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_ydb_library_yql_providers_generic_connector_api_common_data_source_proto protoreflect.FileDescriptor

var file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDesc = []byte{
	0x0a, 0x48, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x4e, 0x59, 0x71, 0x6c,
	0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69,
	0x1a, 0x45, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79, 0x2f, 0x79, 0x71,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x0c, 0x54, 0x43, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x41, 0x0a, 0x05, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x2e, 0x54, 0x42, 0x61, 0x73,
	0x69, 0x63, 0x48, 0x00, 0x52, 0x05, 0x62, 0x61, 0x73, 0x69, 0x63, 0x1a, 0x40, 0x0a, 0x06, 0x54,
	0x42, 0x61, 0x73, 0x69, 0x63, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x42, 0x09, 0x0a,
	0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xc5, 0x02, 0x0a, 0x13, 0x54, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65,
	0x12, 0x39, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25,
	0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x3b, 0x0a, 0x08, 0x65,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x4e, 0x41, 0x70, 0x69, 0x2e, 0x54, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08,
	0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x4e, 0x59, 0x71, 0x6c,
	0x2e, 0x4e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69,
	0x2e, 0x54, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x5f, 0x74, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x54, 0x6c, 0x73, 0x12, 0x3b, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x4e, 0x59, 0x71, 0x6c, 0x2e, 0x4e, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x4e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2a, 0x53, 0x0a, 0x0f, 0x45, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4b,
	0x69, 0x6e, 0x64, 0x12, 0x20, 0x0a, 0x1c, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x43, 0x4c, 0x49, 0x43, 0x4b, 0x48, 0x4f,
	0x55, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x50, 0x4f, 0x53, 0x54, 0x47, 0x52, 0x45,
	0x53, 0x51, 0x4c, 0x10, 0x02, 0x2a, 0x3b, 0x0a, 0x09, 0x45, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x18, 0x0a, 0x14, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4e, 0x41, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x54, 0x54, 0x50,
	0x10, 0x02, 0x42, 0x49, 0x5a, 0x47, 0x61, 0x2e, 0x79, 0x61, 0x6e, 0x64, 0x65, 0x78, 0x2d, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x72, 0x75, 0x2f, 0x79, 0x64, 0x62, 0x2f, 0x6c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x79, 0x2f, 0x79, 0x71, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x73,
	0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescOnce sync.Once
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescData = file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDesc
)

func file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescGZIP() []byte {
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescOnce.Do(func() {
		file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescData)
	})
	return file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDescData
}

var file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_goTypes = []interface{}{
	(EDataSourceKind)(0),        // 0: NYql.NConnector.NApi.EDataSourceKind
	(EProtocol)(0),              // 1: NYql.NConnector.NApi.EProtocol
	(*TCredentials)(nil),        // 2: NYql.NConnector.NApi.TCredentials
	(*TDataSourceInstance)(nil), // 3: NYql.NConnector.NApi.TDataSourceInstance
	(*TCredentials_TBasic)(nil), // 4: NYql.NConnector.NApi.TCredentials.TBasic
	(*TEndpoint)(nil),           // 5: NYql.NConnector.NApi.TEndpoint
}
var file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_depIdxs = []int32{
	4, // 0: NYql.NConnector.NApi.TCredentials.basic:type_name -> NYql.NConnector.NApi.TCredentials.TBasic
	0, // 1: NYql.NConnector.NApi.TDataSourceInstance.kind:type_name -> NYql.NConnector.NApi.EDataSourceKind
	5, // 2: NYql.NConnector.NApi.TDataSourceInstance.endpoint:type_name -> NYql.NConnector.NApi.TEndpoint
	2, // 3: NYql.NConnector.NApi.TDataSourceInstance.credentials:type_name -> NYql.NConnector.NApi.TCredentials
	1, // 4: NYql.NConnector.NApi.TDataSourceInstance.protocol:type_name -> NYql.NConnector.NApi.EProtocol
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_init() }
func file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_init() {
	if File_ydb_library_yql_providers_generic_connector_api_common_data_source_proto != nil {
		return
	}
	file_ydb_library_yql_providers_generic_connector_api_common_endpoint_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCredentials); i {
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
		file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TDataSourceInstance); i {
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
		file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TCredentials_TBasic); i {
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
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*TCredentials_Basic)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_goTypes,
		DependencyIndexes: file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_depIdxs,
		EnumInfos:         file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_enumTypes,
		MessageInfos:      file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_msgTypes,
	}.Build()
	File_ydb_library_yql_providers_generic_connector_api_common_data_source_proto = out.File
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_rawDesc = nil
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_goTypes = nil
	file_ydb_library_yql_providers_generic_connector_api_common_data_source_proto_depIdxs = nil
}
