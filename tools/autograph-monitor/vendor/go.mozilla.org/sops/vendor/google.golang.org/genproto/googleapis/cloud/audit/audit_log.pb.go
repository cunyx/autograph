// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/audit/audit_log.proto

package audit // import "google.golang.org/genproto/googleapis/cloud/audit"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import status "google.golang.org/genproto/googleapis/rpc/status"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Common audit log format for Google Cloud Platform API operations.
type AuditLog struct {
	// The name of the API service performing the operation. For example,
	// `"datastore.googleapis.com"`.
	ServiceName string `protobuf:"bytes,7,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The name of the service method or operation.
	// For API calls, this should be the name of the API method.
	// For example,
	//
	//     "google.datastore.v1.Datastore.RunQuery"
	//     "google.logging.v1.LoggingService.DeleteLog"
	MethodName string `protobuf:"bytes,8,opt,name=method_name,json=methodName,proto3" json:"method_name,omitempty"`
	// The resource or collection that is the target of the operation.
	// The name is a scheme-less URI, not including the API service name.
	// For example:
	//
	//     "shelves/SHELF_ID/books"
	//     "shelves/SHELF_ID/books/BOOK_ID"
	ResourceName string `protobuf:"bytes,11,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The number of items returned from a List or Query API method,
	// if applicable.
	NumResponseItems int64 `protobuf:"varint,12,opt,name=num_response_items,json=numResponseItems,proto3" json:"num_response_items,omitempty"`
	// The status of the overall operation.
	Status *status.Status `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	// Authentication information.
	AuthenticationInfo *AuthenticationInfo `protobuf:"bytes,3,opt,name=authentication_info,json=authenticationInfo,proto3" json:"authentication_info,omitempty"`
	// Authorization information. If there are multiple
	// resources or permissions involved, then there is
	// one AuthorizationInfo element for each {resource, permission} tuple.
	AuthorizationInfo []*AuthorizationInfo `protobuf:"bytes,9,rep,name=authorization_info,json=authorizationInfo,proto3" json:"authorization_info,omitempty"`
	// Metadata about the operation.
	RequestMetadata *RequestMetadata `protobuf:"bytes,4,opt,name=request_metadata,json=requestMetadata,proto3" json:"request_metadata,omitempty"`
	// The operation request. This may not include all request parameters,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Request *_struct.Struct `protobuf:"bytes,16,opt,name=request,proto3" json:"request,omitempty"`
	// The operation response. This may not include all response elements,
	// such as those that are too large, privacy-sensitive, or duplicated
	// elsewhere in the log record.
	// It should never include user-generated data, such as file contents.
	// When the JSON object represented here has a proto equivalent, the proto
	// name will be indicated in the `@type` property.
	Response *_struct.Struct `protobuf:"bytes,17,opt,name=response,proto3" json:"response,omitempty"`
	// Other service-specific data about the request, response, and other
	// activities.
	ServiceData          *any.Any `protobuf:"bytes,15,opt,name=service_data,json=serviceData,proto3" json:"service_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuditLog) Reset()         { *m = AuditLog{} }
func (m *AuditLog) String() string { return proto.CompactTextString(m) }
func (*AuditLog) ProtoMessage()    {}
func (*AuditLog) Descriptor() ([]byte, []int) {
	return fileDescriptor_197799a4334e4cbf, []int{0}
}
func (m *AuditLog) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuditLog.Unmarshal(m, b)
}
func (m *AuditLog) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuditLog.Marshal(b, m, deterministic)
}
func (m *AuditLog) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuditLog.Merge(m, src)
}
func (m *AuditLog) XXX_Size() int {
	return xxx_messageInfo_AuditLog.Size(m)
}
func (m *AuditLog) XXX_DiscardUnknown() {
	xxx_messageInfo_AuditLog.DiscardUnknown(m)
}

var xxx_messageInfo_AuditLog proto.InternalMessageInfo

func (m *AuditLog) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AuditLog) GetMethodName() string {
	if m != nil {
		return m.MethodName
	}
	return ""
}

func (m *AuditLog) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *AuditLog) GetNumResponseItems() int64 {
	if m != nil {
		return m.NumResponseItems
	}
	return 0
}

func (m *AuditLog) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *AuditLog) GetAuthenticationInfo() *AuthenticationInfo {
	if m != nil {
		return m.AuthenticationInfo
	}
	return nil
}

func (m *AuditLog) GetAuthorizationInfo() []*AuthorizationInfo {
	if m != nil {
		return m.AuthorizationInfo
	}
	return nil
}

func (m *AuditLog) GetRequestMetadata() *RequestMetadata {
	if m != nil {
		return m.RequestMetadata
	}
	return nil
}

func (m *AuditLog) GetRequest() *_struct.Struct {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *AuditLog) GetResponse() *_struct.Struct {
	if m != nil {
		return m.Response
	}
	return nil
}

func (m *AuditLog) GetServiceData() *any.Any {
	if m != nil {
		return m.ServiceData
	}
	return nil
}

// Authentication information for the operation.
type AuthenticationInfo struct {
	// The email address of the authenticated user making the request.
	PrincipalEmail       string   `protobuf:"bytes,1,opt,name=principal_email,json=principalEmail,proto3" json:"principal_email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthenticationInfo) Reset()         { *m = AuthenticationInfo{} }
func (m *AuthenticationInfo) String() string { return proto.CompactTextString(m) }
func (*AuthenticationInfo) ProtoMessage()    {}
func (*AuthenticationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_197799a4334e4cbf, []int{1}
}
func (m *AuthenticationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthenticationInfo.Unmarshal(m, b)
}
func (m *AuthenticationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthenticationInfo.Marshal(b, m, deterministic)
}
func (m *AuthenticationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticationInfo.Merge(m, src)
}
func (m *AuthenticationInfo) XXX_Size() int {
	return xxx_messageInfo_AuthenticationInfo.Size(m)
}
func (m *AuthenticationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticationInfo proto.InternalMessageInfo

func (m *AuthenticationInfo) GetPrincipalEmail() string {
	if m != nil {
		return m.PrincipalEmail
	}
	return ""
}

// Authorization information for the operation.
type AuthorizationInfo struct {
	// The resource being accessed, as a REST-style string. For example:
	//
	//     bigquery.googlapis.com/projects/PROJECTID/datasets/DATASETID
	Resource string `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// The required IAM permission.
	Permission string `protobuf:"bytes,2,opt,name=permission,proto3" json:"permission,omitempty"`
	// Whether or not authorization for `resource` and `permission`
	// was granted.
	Granted              bool     `protobuf:"varint,3,opt,name=granted,proto3" json:"granted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationInfo) Reset()         { *m = AuthorizationInfo{} }
func (m *AuthorizationInfo) String() string { return proto.CompactTextString(m) }
func (*AuthorizationInfo) ProtoMessage()    {}
func (*AuthorizationInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_197799a4334e4cbf, []int{2}
}
func (m *AuthorizationInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationInfo.Unmarshal(m, b)
}
func (m *AuthorizationInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationInfo.Marshal(b, m, deterministic)
}
func (m *AuthorizationInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationInfo.Merge(m, src)
}
func (m *AuthorizationInfo) XXX_Size() int {
	return xxx_messageInfo_AuthorizationInfo.Size(m)
}
func (m *AuthorizationInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationInfo.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationInfo proto.InternalMessageInfo

func (m *AuthorizationInfo) GetResource() string {
	if m != nil {
		return m.Resource
	}
	return ""
}

func (m *AuthorizationInfo) GetPermission() string {
	if m != nil {
		return m.Permission
	}
	return ""
}

func (m *AuthorizationInfo) GetGranted() bool {
	if m != nil {
		return m.Granted
	}
	return false
}

// Metadata about the request.
type RequestMetadata struct {
	// The IP address of the caller.
	CallerIp string `protobuf:"bytes,1,opt,name=caller_ip,json=callerIp,proto3" json:"caller_ip,omitempty"`
	// The user agent of the caller.
	// This information is not authenticated and should be treated accordingly.
	// For example:
	//
	// +   `google-api-python-client/1.4.0`:
	//     The request was made by the Google API client for Python.
	// +   `Cloud SDK Command Line Tool apitools-client/1.0 gcloud/0.9.62`:
	//     The request was made by the Google Cloud SDK CLI (gcloud).
	// +   `AppEngine-Google; (+http://code.google.com/appengine; appid: s~my-project`:
	//     The request was made from the `my-project` App Engine app.
	CallerSuppliedUserAgent string   `protobuf:"bytes,2,opt,name=caller_supplied_user_agent,json=callerSuppliedUserAgent,proto3" json:"caller_supplied_user_agent,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *RequestMetadata) Reset()         { *m = RequestMetadata{} }
func (m *RequestMetadata) String() string { return proto.CompactTextString(m) }
func (*RequestMetadata) ProtoMessage()    {}
func (*RequestMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_197799a4334e4cbf, []int{3}
}
func (m *RequestMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMetadata.Unmarshal(m, b)
}
func (m *RequestMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMetadata.Marshal(b, m, deterministic)
}
func (m *RequestMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMetadata.Merge(m, src)
}
func (m *RequestMetadata) XXX_Size() int {
	return xxx_messageInfo_RequestMetadata.Size(m)
}
func (m *RequestMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMetadata proto.InternalMessageInfo

func (m *RequestMetadata) GetCallerIp() string {
	if m != nil {
		return m.CallerIp
	}
	return ""
}

func (m *RequestMetadata) GetCallerSuppliedUserAgent() string {
	if m != nil {
		return m.CallerSuppliedUserAgent
	}
	return ""
}

func init() {
	proto.RegisterType((*AuditLog)(nil), "google.cloud.audit.AuditLog")
	proto.RegisterType((*AuthenticationInfo)(nil), "google.cloud.audit.AuthenticationInfo")
	proto.RegisterType((*AuthorizationInfo)(nil), "google.cloud.audit.AuthorizationInfo")
	proto.RegisterType((*RequestMetadata)(nil), "google.cloud.audit.RequestMetadata")
}

func init() { proto.RegisterFile("google/cloud/audit/audit_log.proto", fileDescriptor_197799a4334e4cbf) }

var fileDescriptor_197799a4334e4cbf = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0x5f, 0x6f, 0xd3, 0x30,
	0x14, 0xc5, 0x55, 0x36, 0x6d, 0xad, 0xbb, 0xd1, 0xd6, 0x20, 0x1a, 0xca, 0x04, 0xa5, 0x13, 0x50,
	0x21, 0x94, 0x88, 0xed, 0x61, 0x0f, 0x13, 0x0f, 0x9d, 0xe0, 0xa1, 0x12, 0x4c, 0x53, 0x0a, 0x42,
	0xe2, 0x25, 0x72, 0x93, 0xdb, 0xcc, 0x22, 0xb1, 0x8d, 0xff, 0x20, 0x8d, 0xef, 0xcc, 0x77, 0x40,
	0xbd, 0x71, 0x4a, 0xd7, 0x0e, 0x5e, 0x2c, 0xf9, 0x9c, 0xdf, 0xbd, 0x76, 0xaf, 0x4f, 0x43, 0x46,
	0xb9, 0x94, 0x79, 0x01, 0x51, 0x5a, 0x48, 0x97, 0x45, 0xcc, 0x65, 0xdc, 0x56, 0x6b, 0x52, 0xc8,
	0x3c, 0x54, 0x5a, 0x5a, 0x49, 0x69, 0xc5, 0x84, 0xc8, 0x84, 0xe8, 0x0e, 0x8e, 0x7c, 0x1d, 0x53,
	0x3c, 0x62, 0x42, 0x48, 0xcb, 0x2c, 0x97, 0xc2, 0x54, 0x15, 0x83, 0xc7, 0xde, 0xc5, 0xdd, 0xdc,
	0x2d, 0x22, 0x26, 0x6e, 0xbc, 0x75, 0xb4, 0x69, 0x19, 0xab, 0x5d, 0x6a, 0xbd, 0xdb, 0xf7, 0xae,
	0x56, 0x69, 0x64, 0x2c, 0xb3, 0xce, 0x77, 0x1c, 0xfd, 0xde, 0x25, 0xcd, 0xc9, 0xf2, 0xe4, 0x8f,
	0x32, 0xa7, 0xcf, 0xc9, 0x81, 0x01, 0xfd, 0x93, 0xa7, 0x90, 0x08, 0x56, 0x42, 0xb0, 0x3f, 0x6c,
	0x8c, 0x5b, 0x71, 0xdb, 0x6b, 0x97, 0xac, 0x04, 0xfa, 0x8c, 0xb4, 0x4b, 0xb0, 0xd7, 0x32, 0xab,
	0x88, 0x26, 0x12, 0xa4, 0x92, 0x10, 0x38, 0x26, 0x87, 0x1a, 0x8c, 0x74, 0xba, 0x6e, 0xd2, 0x46,
	0xe4, 0xa0, 0x16, 0x11, 0x7a, 0x43, 0xa8, 0x70, 0x65, 0xa2, 0xc1, 0x28, 0x29, 0x0c, 0x24, 0xdc,
	0x42, 0x69, 0x82, 0x83, 0x61, 0x63, 0xbc, 0x13, 0x77, 0x85, 0x2b, 0x63, 0x6f, 0x4c, 0x97, 0x3a,
	0x7d, 0x4d, 0xf6, 0xaa, 0x3b, 0x07, 0xf7, 0x86, 0x8d, 0x71, 0xfb, 0x84, 0x86, 0x7e, 0x70, 0x5a,
	0xa5, 0xe1, 0x0c, 0x9d, 0xd8, 0x13, 0xf4, 0x2b, 0x79, 0xc0, 0x9c, 0xbd, 0x06, 0x61, 0x79, 0x8a,
	0xa3, 0x4b, 0xb8, 0x58, 0xc8, 0x60, 0x07, 0x0b, 0x5f, 0x86, 0xdb, 0x13, 0x0f, 0x27, 0xb7, 0xf0,
	0xa9, 0x58, 0xc8, 0x98, 0xb2, 0x2d, 0x8d, 0x7e, 0x26, 0xa8, 0x4a, 0xcd, 0x7f, 0xad, 0xf5, 0x6d,
	0x0d, 0x77, 0xc6, 0xed, 0x93, 0x17, 0xff, 0xea, 0xbb, 0xa2, 0xb1, 0x6d, 0x8f, 0x6d, 0x4a, 0xf4,
	0x92, 0x74, 0x35, 0xfc, 0x70, 0x60, 0x6c, 0x52, 0x82, 0x65, 0x19, 0xb3, 0x2c, 0xd8, 0xc5, 0xbb,
	0x1e, 0xdf, 0xd5, 0x33, 0xae, 0xd8, 0x4f, 0x1e, 0x8d, 0x3b, 0xfa, 0xb6, 0x40, 0xdf, 0x92, 0x7d,
	0x2f, 0x05, 0x5d, 0x6c, 0xd3, 0xaf, 0xdb, 0xd4, 0xb9, 0x08, 0x67, 0x98, 0x8b, 0xb8, 0xe6, 0xe8,
	0x29, 0x69, 0xd6, 0xef, 0x10, 0xf4, 0xfe, 0x5f, 0xb3, 0x02, 0xe9, 0xd9, 0xdf, 0xa4, 0xe0, 0x9d,
	0x3b, 0x58, 0xf8, 0x70, 0xab, 0x70, 0x22, 0x6e, 0x56, 0xf9, 0x79, 0xcf, 0x2c, 0x1b, 0xbd, 0x23,
	0x74, 0x7b, 0xe0, 0xf4, 0x15, 0xe9, 0x28, 0xcd, 0x45, 0xca, 0x15, 0x2b, 0x12, 0x28, 0x19, 0x2f,
	0x82, 0x06, 0xc6, 0xe6, 0xfe, 0x4a, 0xfe, 0xb0, 0x54, 0x47, 0x9c, 0xf4, 0xb6, 0xe6, 0x4a, 0x07,
	0xf8, 0x0b, 0x30, 0x5d, 0xbe, 0x6c, 0xb5, 0xa7, 0x4f, 0x09, 0x51, 0xa0, 0x4b, 0x6e, 0x0c, 0x97,
	0x02, 0xf3, 0xd3, 0x8a, 0xd7, 0x14, 0x1a, 0x90, 0xfd, 0x5c, 0x33, 0x61, 0x21, 0xc3, 0x8c, 0x34,
	0xe3, 0x7a, 0x3b, 0xfa, 0x4e, 0x3a, 0x1b, 0xe3, 0xa6, 0x4f, 0x48, 0x2b, 0x65, 0x45, 0x01, 0x3a,
	0xe1, 0xaa, 0x3e, 0xa9, 0x12, 0xa6, 0x8a, 0x9e, 0x93, 0x81, 0x37, 0x8d, 0x53, 0xaa, 0xe0, 0x90,
	0x25, 0xce, 0x80, 0x4e, 0x58, 0x0e, 0xc2, 0xfa, 0x93, 0xfb, 0x15, 0x31, 0xf3, 0xc0, 0x17, 0x03,
	0x7a, 0xb2, 0xb4, 0x2f, 0xe6, 0xe4, 0x51, 0x2a, 0xcb, 0x3b, 0x9e, 0xfc, 0xe2, 0xb0, 0xfe, 0x77,
	0x5e, 0x2d, 0x67, 0x7a, 0xd5, 0xf8, 0x76, 0xe6, 0xa1, 0x5c, 0x16, 0x4c, 0xe4, 0xa1, 0xd4, 0x79,
	0x94, 0x83, 0xc0, 0x89, 0x47, 0x95, 0xc5, 0x14, 0x37, 0xeb, 0x1f, 0x9e, 0x73, 0x5c, 0xe7, 0x7b,
	0xc8, 0x9c, 0xfe, 0x09, 0x00, 0x00, 0xff, 0xff, 0x90, 0xe4, 0x37, 0xbf, 0x9b, 0x04, 0x00, 0x00,
}
