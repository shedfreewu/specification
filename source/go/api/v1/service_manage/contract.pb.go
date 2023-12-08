// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract.proto

package service_manage // import "github.com/polarismesh/specification/source/go/api/v1/service_manage"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type InterfaceDescriptor_Source int32

const (
	InterfaceDescriptor_UNKNOWN InterfaceDescriptor_Source = 0
	InterfaceDescriptor_Manual  InterfaceDescriptor_Source = 1
	InterfaceDescriptor_Client  InterfaceDescriptor_Source = 2
)

var InterfaceDescriptor_Source_name = map[int32]string{
	0: "UNKNOWN",
	1: "Manual",
	2: "Client",
}
var InterfaceDescriptor_Source_value = map[string]int32{
	"UNKNOWN": 0,
	"Manual":  1,
	"Client":  2,
}

func (x InterfaceDescriptor_Source) String() string {
	return proto.EnumName(InterfaceDescriptor_Source_name, int32(x))
}
func (InterfaceDescriptor_Source) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_contract_cc22240765a28406, []int{1, 0}
}

type ServiceContract struct {
	// 契约ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 契约名称
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// 所属命名空间
	Namespace string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	// 所属服务名称
	Service string `protobuf:"bytes,4,opt,name=service,proto3" json:"service,omitempty"`
	// 协议，http/grpc/dubbo/thrift
	Protocol string `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// 契约版本
	Version string `protobuf:"bytes,6,opt,name=version,proto3" json:"version,omitempty"`
	// 信息摘要
	Revision string `protobuf:"bytes,7,opt,name=revision,proto3" json:"revision,omitempty"`
	// 额外描述
	Content string `protobuf:"bytes,8,opt,name=content,proto3" json:"content,omitempty"`
	// 接口描述信息
	Interfaces []*InterfaceDescriptor `protobuf:"bytes,9,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	// 创建时间
	Ctime string `protobuf:"bytes,10,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 更新时间
	Mtime string `protobuf:"bytes,11,opt,name=mtime,proto3" json:"mtime,omitempty"`
	// 接口状态，Offline/Online
	Status               string   `protobuf:"bytes,12,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ServiceContract) Reset()         { *m = ServiceContract{} }
func (m *ServiceContract) String() string { return proto.CompactTextString(m) }
func (*ServiceContract) ProtoMessage()    {}
func (*ServiceContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_cc22240765a28406, []int{0}
}
func (m *ServiceContract) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ServiceContract.Unmarshal(m, b)
}
func (m *ServiceContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ServiceContract.Marshal(b, m, deterministic)
}
func (dst *ServiceContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServiceContract.Merge(dst, src)
}
func (m *ServiceContract) XXX_Size() int {
	return xxx_messageInfo_ServiceContract.Size(m)
}
func (m *ServiceContract) XXX_DiscardUnknown() {
	xxx_messageInfo_ServiceContract.DiscardUnknown(m)
}

var xxx_messageInfo_ServiceContract proto.InternalMessageInfo

func (m *ServiceContract) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServiceContract) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ServiceContract) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ServiceContract) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ServiceContract) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *ServiceContract) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ServiceContract) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *ServiceContract) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *ServiceContract) GetInterfaces() []*InterfaceDescriptor {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ServiceContract) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *ServiceContract) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func (m *ServiceContract) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

type InterfaceDescriptor struct {
	// 接口ID
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// 方法名称，对应 http method/ dubbo interface func/grpc service func
	Method string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	// 接口名称，http path/dubbo interface/grpc service
	Path string `protobuf:"bytes,3,opt,name=path,proto3" json:"path,omitempty"`
	// 接口描述信息
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	// 创建来源
	Source InterfaceDescriptor_Source `protobuf:"varint,5,opt,name=source,proto3,enum=v1.InterfaceDescriptor_Source" json:"source,omitempty"`
	// 接口信息摘要
	Revision string `protobuf:"bytes,6,opt,name=revision,proto3" json:"revision,omitempty"`
	// 创建时间
	Ctime string `protobuf:"bytes,7,opt,name=ctime,proto3" json:"ctime,omitempty"`
	// 更新时间
	Mtime                string   `protobuf:"bytes,8,opt,name=mtime,proto3" json:"mtime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceDescriptor) Reset()         { *m = InterfaceDescriptor{} }
func (m *InterfaceDescriptor) String() string { return proto.CompactTextString(m) }
func (*InterfaceDescriptor) ProtoMessage()    {}
func (*InterfaceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_contract_cc22240765a28406, []int{1}
}
func (m *InterfaceDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceDescriptor.Unmarshal(m, b)
}
func (m *InterfaceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceDescriptor.Marshal(b, m, deterministic)
}
func (dst *InterfaceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceDescriptor.Merge(dst, src)
}
func (m *InterfaceDescriptor) XXX_Size() int {
	return xxx_messageInfo_InterfaceDescriptor.Size(m)
}
func (m *InterfaceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceDescriptor proto.InternalMessageInfo

func (m *InterfaceDescriptor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *InterfaceDescriptor) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *InterfaceDescriptor) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *InterfaceDescriptor) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *InterfaceDescriptor) GetSource() InterfaceDescriptor_Source {
	if m != nil {
		return m.Source
	}
	return InterfaceDescriptor_UNKNOWN
}

func (m *InterfaceDescriptor) GetRevision() string {
	if m != nil {
		return m.Revision
	}
	return ""
}

func (m *InterfaceDescriptor) GetCtime() string {
	if m != nil {
		return m.Ctime
	}
	return ""
}

func (m *InterfaceDescriptor) GetMtime() string {
	if m != nil {
		return m.Mtime
	}
	return ""
}

func init() {
	proto.RegisterType((*ServiceContract)(nil), "v1.ServiceContract")
	proto.RegisterType((*InterfaceDescriptor)(nil), "v1.InterfaceDescriptor")
	proto.RegisterEnum("v1.InterfaceDescriptor_Source", InterfaceDescriptor_Source_name, InterfaceDescriptor_Source_value)
}

func init() { proto.RegisterFile("contract.proto", fileDescriptor_contract_cc22240765a28406) }

var fileDescriptor_contract_cc22240765a28406 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x5d, 0x6b, 0x13, 0x41,
	0x14, 0x35, 0xdb, 0x74, 0x93, 0xdc, 0x48, 0x0c, 0x63, 0xa9, 0x83, 0x88, 0x94, 0x3c, 0xf5, 0xc5,
	0x59, 0x52, 0xc1, 0xbe, 0xb7, 0x7d, 0x11, 0x31, 0x4a, 0x8a, 0x08, 0xbe, 0xc8, 0x74, 0x72, 0x9b,
	0x0c, 0x64, 0x67, 0x96, 0x99, 0x9b, 0xfd, 0x27, 0xbe, 0xf8, 0xc3, 0xfc, 0x3d, 0x32, 0x1f, 0xa9,
	0x49, 0x4d, 0x9f, 0x76, 0xce, 0xb9, 0xf7, 0x9c, 0xe5, 0x1c, 0x2e, 0x8c, 0x94, 0x35, 0xe4, 0xa4,
	0x22, 0xd1, 0x38, 0x4b, 0x96, 0x15, 0xed, 0x74, 0xf2, 0xa7, 0x80, 0x17, 0xb7, 0xe8, 0x5a, 0xad,
	0xf0, 0x3a, 0x4f, 0xd9, 0x08, 0x0a, 0xbd, 0xe0, 0x9d, 0xb3, 0xce, 0xf9, 0x60, 0x5e, 0xe8, 0x05,
	0x63, 0xd0, 0x35, 0xb2, 0x46, 0x5e, 0x44, 0x26, 0xbe, 0xd9, 0x1b, 0x18, 0x84, 0xaf, 0x6f, 0xa4,
	0x42, 0x7e, 0x14, 0x07, 0xff, 0x08, 0xc6, 0xa1, 0xe7, 0x93, 0x29, 0xef, 0xc6, 0xd9, 0x16, 0xb2,
	0xd7, 0xd0, 0x8f, 0x3f, 0x57, 0x76, 0xcd, 0x8f, 0xe3, 0xe8, 0x01, 0x07, 0x55, 0x8b, 0xce, 0x6b,
	0x6b, 0x78, 0x99, 0x54, 0x19, 0x06, 0x95, 0xc3, 0x56, 0xc7, 0x51, 0x2f, 0xa9, 0xb6, 0x38, 0xa8,
	0x42, 0x2e, 0x34, 0xc4, 0xfb, 0x49, 0x95, 0x21, 0xbb, 0x04, 0xd0, 0x86, 0xd0, 0xdd, 0x4b, 0x85,
	0x9e, 0x0f, 0xce, 0x8e, 0xce, 0x87, 0x17, 0xaf, 0x44, 0x3b, 0x15, 0x1f, 0xb7, 0xec, 0x0d, 0x7a,
	0xe5, 0x74, 0x43, 0xd6, 0xcd, 0x77, 0x56, 0xd9, 0x09, 0x1c, 0x2b, 0xd2, 0x35, 0x72, 0x88, 0x86,
	0x09, 0x04, 0xb6, 0x8e, 0xec, 0x30, 0xb1, 0x11, 0xb0, 0x53, 0x28, 0x3d, 0x49, 0xda, 0x78, 0xfe,
	0x3c, 0xd2, 0x19, 0x4d, 0x7e, 0x17, 0xf0, 0xf2, 0xc0, 0x7f, 0xfe, 0x2b, 0xf7, 0x14, 0xca, 0x1a,
	0x69, 0x65, 0x17, 0xb9, 0xde, 0x8c, 0x42, 0xe9, 0x8d, 0xa4, 0x55, 0xee, 0x36, 0xbe, 0x77, 0xa3,
	0x76, 0xf7, 0xa3, 0x7e, 0x80, 0xd2, 0xdb, 0x8d, 0x53, 0x18, 0x4b, 0x1d, 0x5d, 0xbc, 0x7d, 0x22,
	0xa6, 0xb8, 0x8d, 0x5b, 0xf3, 0xbc, 0xbd, 0x57, 0x6c, 0xf9, 0xa8, 0xd8, 0x87, 0x16, 0x7a, 0x07,
	0x5b, 0xe8, 0xef, 0xb4, 0x30, 0x79, 0x07, 0x65, 0x72, 0x66, 0x43, 0xe8, 0x7d, 0x9b, 0x7d, 0x9a,
	0x7d, 0xf9, 0x3e, 0x1b, 0x3f, 0x63, 0x00, 0xe5, 0x67, 0x69, 0x36, 0x72, 0x3d, 0xee, 0x84, 0xf7,
	0xf5, 0x5a, 0xa3, 0xa1, 0x71, 0x71, 0xf5, 0xab, 0x03, 0x97, 0xca, 0xd6, 0x82, 0xd0, 0x28, 0x34,
	0x24, 0x1a, 0xbb, 0x96, 0x4e, 0x7b, 0xe1, 0x1b, 0x54, 0xfa, 0x5e, 0x2b, 0x49, 0xda, 0x1a, 0x21,
	0x1b, 0x1d, 0x62, 0xe4, 0xd3, 0x11, 0xb5, 0x34, 0x72, 0x89, 0x57, 0x27, 0x8f, 0xce, 0xf5, 0x6b,
	0x38, 0x9f, 0x1f, 0x37, 0x4b, 0x4d, 0xab, 0xcd, 0x9d, 0x50, 0xb6, 0xae, 0xb2, 0x5b, 0x8d, 0x7e,
	0x55, 0xed, 0x39, 0x56, 0x29, 0x75, 0xb5, 0xb4, 0x95, 0x6c, 0x74, 0xd5, 0x4e, 0xab, 0xec, 0xfd,
	0x33, 0x79, 0xdf, 0x95, 0xf1, 0x16, 0xdf, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x53, 0x77, 0x70,
	0xdc, 0x2a, 0x03, 0x00, 0x00,
}