// Code generated by protoc-gen-go. DO NOT EDIT.
// source: connectioncontext.proto

package connectioncontext

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IpFamily_Family int32

const (
	IpFamily_IPV4 IpFamily_Family = 0
	IpFamily_IPV6 IpFamily_Family = 1
)

var IpFamily_Family_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}

var IpFamily_Family_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x IpFamily_Family) String() string {
	return proto.EnumName(IpFamily_Family_name, int32(x))
}

func (IpFamily_Family) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2, 0}
}

type IpNeighbor struct {
	Ip                   string   `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
	HardwareAddress      string   `protobuf:"bytes,2,opt,name=hardware_address,json=hardwareAddress,proto3" json:"hardware_address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IpNeighbor) Reset()         { *m = IpNeighbor{} }
func (m *IpNeighbor) String() string { return proto.CompactTextString(m) }
func (*IpNeighbor) ProtoMessage()    {}
func (*IpNeighbor) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{0}
}

func (m *IpNeighbor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpNeighbor.Unmarshal(m, b)
}
func (m *IpNeighbor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpNeighbor.Marshal(b, m, deterministic)
}
func (m *IpNeighbor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpNeighbor.Merge(m, src)
}
func (m *IpNeighbor) XXX_Size() int {
	return xxx_messageInfo_IpNeighbor.Size(m)
}
func (m *IpNeighbor) XXX_DiscardUnknown() {
	xxx_messageInfo_IpNeighbor.DiscardUnknown(m)
}

var xxx_messageInfo_IpNeighbor proto.InternalMessageInfo

func (m *IpNeighbor) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *IpNeighbor) GetHardwareAddress() string {
	if m != nil {
		return m.HardwareAddress
	}
	return ""
}

type Route struct {
	Prefix               string   `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{1}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type IpFamily struct {
	Family               IpFamily_Family `protobuf:"varint,1,opt,name=family,proto3,enum=connectioncontext.IpFamily_Family" json:"family,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IpFamily) Reset()         { *m = IpFamily{} }
func (m *IpFamily) String() string { return proto.CompactTextString(m) }
func (*IpFamily) ProtoMessage()    {}
func (*IpFamily) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{2}
}

func (m *IpFamily) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IpFamily.Unmarshal(m, b)
}
func (m *IpFamily) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IpFamily.Marshal(b, m, deterministic)
}
func (m *IpFamily) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IpFamily.Merge(m, src)
}
func (m *IpFamily) XXX_Size() int {
	return xxx_messageInfo_IpFamily.Size(m)
}
func (m *IpFamily) XXX_DiscardUnknown() {
	xxx_messageInfo_IpFamily.DiscardUnknown(m)
}

var xxx_messageInfo_IpFamily proto.InternalMessageInfo

func (m *IpFamily) GetFamily() IpFamily_Family {
	if m != nil {
		return m.Family
	}
	return IpFamily_IPV4
}

type ExtraPrefixRequest struct {
	AddrFamily           *IpFamily `protobuf:"bytes,1,opt,name=addr_family,json=addrFamily,proto3" json:"addr_family,omitempty"`
	PrefixLen            uint32    `protobuf:"varint,2,opt,name=prefix_len,json=prefixLen,proto3" json:"prefix_len,omitempty"`
	RequiredNumber       uint32    `protobuf:"varint,3,opt,name=required_number,json=requiredNumber,proto3" json:"required_number,omitempty"`
	RequestedNumber      uint32    `protobuf:"varint,4,opt,name=requested_number,json=requestedNumber,proto3" json:"requested_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ExtraPrefixRequest) Reset()         { *m = ExtraPrefixRequest{} }
func (m *ExtraPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*ExtraPrefixRequest) ProtoMessage()    {}
func (*ExtraPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{3}
}

func (m *ExtraPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraPrefixRequest.Unmarshal(m, b)
}
func (m *ExtraPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraPrefixRequest.Marshal(b, m, deterministic)
}
func (m *ExtraPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraPrefixRequest.Merge(m, src)
}
func (m *ExtraPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_ExtraPrefixRequest.Size(m)
}
func (m *ExtraPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraPrefixRequest proto.InternalMessageInfo

func (m *ExtraPrefixRequest) GetAddrFamily() *IpFamily {
	if m != nil {
		return m.AddrFamily
	}
	return nil
}

func (m *ExtraPrefixRequest) GetPrefixLen() uint32 {
	if m != nil {
		return m.PrefixLen
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequiredNumber() uint32 {
	if m != nil {
		return m.RequiredNumber
	}
	return 0
}

func (m *ExtraPrefixRequest) GetRequestedNumber() uint32 {
	if m != nil {
		return m.RequestedNumber
	}
	return 0
}

type IPContext struct {
	SrcIpAddr            string                `protobuf:"bytes,1,opt,name=src_ip_addr,json=srcIpAddr,proto3" json:"src_ip_addr,omitempty"`
	DstIpAddr            string                `protobuf:"bytes,2,opt,name=dst_ip_addr,json=dstIpAddr,proto3" json:"dst_ip_addr,omitempty"`
	SrcIpRequired        bool                  `protobuf:"varint,3,opt,name=src_ip_required,json=srcIpRequired,proto3" json:"src_ip_required,omitempty"`
	DstIpRequired        bool                  `protobuf:"varint,4,opt,name=dst_ip_required,json=dstIpRequired,proto3" json:"dst_ip_required,omitempty"`
	SrcRoutes            []*Route              `protobuf:"bytes,5,rep,name=src_routes,json=srcRoutes,proto3" json:"src_routes,omitempty"`
	DstRoutes            []*Route              `protobuf:"bytes,6,rep,name=dst_routes,json=dstRoutes,proto3" json:"dst_routes,omitempty"`
	ExcludedPrefixes     []string              `protobuf:"bytes,7,rep,name=excluded_prefixes,json=excludedPrefixes,proto3" json:"excluded_prefixes,omitempty"`
	IpNeighbors          []*IpNeighbor         `protobuf:"bytes,8,rep,name=ip_neighbors,json=ipNeighbors,proto3" json:"ip_neighbors,omitempty"`
	ExtraPrefixRequest   []*ExtraPrefixRequest `protobuf:"bytes,9,rep,name=extra_prefix_request,json=extraPrefixRequest,proto3" json:"extra_prefix_request,omitempty"`
	ExtraPrefixes        []string              `protobuf:"bytes,10,rep,name=extra_prefixes,json=extraPrefixes,proto3" json:"extra_prefixes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *IPContext) Reset()         { *m = IPContext{} }
func (m *IPContext) String() string { return proto.CompactTextString(m) }
func (*IPContext) ProtoMessage()    {}
func (*IPContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{4}
}

func (m *IPContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IPContext.Unmarshal(m, b)
}
func (m *IPContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IPContext.Marshal(b, m, deterministic)
}
func (m *IPContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IPContext.Merge(m, src)
}
func (m *IPContext) XXX_Size() int {
	return xxx_messageInfo_IPContext.Size(m)
}
func (m *IPContext) XXX_DiscardUnknown() {
	xxx_messageInfo_IPContext.DiscardUnknown(m)
}

var xxx_messageInfo_IPContext proto.InternalMessageInfo

func (m *IPContext) GetSrcIpAddr() string {
	if m != nil {
		return m.SrcIpAddr
	}
	return ""
}

func (m *IPContext) GetDstIpAddr() string {
	if m != nil {
		return m.DstIpAddr
	}
	return ""
}

func (m *IPContext) GetSrcIpRequired() bool {
	if m != nil {
		return m.SrcIpRequired
	}
	return false
}

func (m *IPContext) GetDstIpRequired() bool {
	if m != nil {
		return m.DstIpRequired
	}
	return false
}

func (m *IPContext) GetSrcRoutes() []*Route {
	if m != nil {
		return m.SrcRoutes
	}
	return nil
}

func (m *IPContext) GetDstRoutes() []*Route {
	if m != nil {
		return m.DstRoutes
	}
	return nil
}

func (m *IPContext) GetExcludedPrefixes() []string {
	if m != nil {
		return m.ExcludedPrefixes
	}
	return nil
}

func (m *IPContext) GetIpNeighbors() []*IpNeighbor {
	if m != nil {
		return m.IpNeighbors
	}
	return nil
}

func (m *IPContext) GetExtraPrefixRequest() []*ExtraPrefixRequest {
	if m != nil {
		return m.ExtraPrefixRequest
	}
	return nil
}

func (m *IPContext) GetExtraPrefixes() []string {
	if m != nil {
		return m.ExtraPrefixes
	}
	return nil
}

type DNSConfig struct {
	// ips of DNS Servers for this DNSConfig.  Any given IP may be IPv4 or IPv6
	DnsServerIps []string `protobuf:"bytes,1,rep,name=dns_server_ips,json=dnsServerIps,proto3" json:"dns_server_ips,omitempty"`
	// domains for which this DNSConfig provides resolution.  If empty, all domains.
	SearchDomains        []string `protobuf:"bytes,2,rep,name=search_domains,json=searchDomains,proto3" json:"search_domains,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DNSConfig) Reset()         { *m = DNSConfig{} }
func (m *DNSConfig) String() string { return proto.CompactTextString(m) }
func (*DNSConfig) ProtoMessage()    {}
func (*DNSConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{5}
}

func (m *DNSConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSConfig.Unmarshal(m, b)
}
func (m *DNSConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSConfig.Marshal(b, m, deterministic)
}
func (m *DNSConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSConfig.Merge(m, src)
}
func (m *DNSConfig) XXX_Size() int {
	return xxx_messageInfo_DNSConfig.Size(m)
}
func (m *DNSConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DNSConfig proto.InternalMessageInfo

func (m *DNSConfig) GetDnsServerIps() []string {
	if m != nil {
		return m.DnsServerIps
	}
	return nil
}

func (m *DNSConfig) GetSearchDomains() []string {
	if m != nil {
		return m.SearchDomains
	}
	return nil
}

type EthernetContext struct {
	SrcMac               string   `protobuf:"bytes,1,opt,name=src_mac,json=srcMac,proto3" json:"src_mac,omitempty"`
	DstMac               string   `protobuf:"bytes,2,opt,name=dst_mac,json=dstMac,proto3" json:"dst_mac,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthernetContext) Reset()         { *m = EthernetContext{} }
func (m *EthernetContext) String() string { return proto.CompactTextString(m) }
func (*EthernetContext) ProtoMessage()    {}
func (*EthernetContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{6}
}

func (m *EthernetContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthernetContext.Unmarshal(m, b)
}
func (m *EthernetContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthernetContext.Marshal(b, m, deterministic)
}
func (m *EthernetContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthernetContext.Merge(m, src)
}
func (m *EthernetContext) XXX_Size() int {
	return xxx_messageInfo_EthernetContext.Size(m)
}
func (m *EthernetContext) XXX_DiscardUnknown() {
	xxx_messageInfo_EthernetContext.DiscardUnknown(m)
}

var xxx_messageInfo_EthernetContext proto.InternalMessageInfo

func (m *EthernetContext) GetSrcMac() string {
	if m != nil {
		return m.SrcMac
	}
	return ""
}

func (m *EthernetContext) GetDstMac() string {
	if m != nil {
		return m.DstMac
	}
	return ""
}

type DNSContext struct {
	Configs              []*DNSConfig `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *DNSContext) Reset()         { *m = DNSContext{} }
func (m *DNSContext) String() string { return proto.CompactTextString(m) }
func (*DNSContext) ProtoMessage()    {}
func (*DNSContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{7}
}

func (m *DNSContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DNSContext.Unmarshal(m, b)
}
func (m *DNSContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DNSContext.Marshal(b, m, deterministic)
}
func (m *DNSContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DNSContext.Merge(m, src)
}
func (m *DNSContext) XXX_Size() int {
	return xxx_messageInfo_DNSContext.Size(m)
}
func (m *DNSContext) XXX_DiscardUnknown() {
	xxx_messageInfo_DNSContext.DiscardUnknown(m)
}

var xxx_messageInfo_DNSContext proto.InternalMessageInfo

func (m *DNSContext) GetConfigs() []*DNSConfig {
	if m != nil {
		return m.Configs
	}
	return nil
}

type ConnectionContext struct {
	IpContext            *IPContext       `protobuf:"bytes,1,opt,name=ip_context,json=ipContext,proto3" json:"ip_context,omitempty"`
	DnsContext           *DNSContext      `protobuf:"bytes,2,opt,name=dns_context,json=dnsContext,proto3" json:"dns_context,omitempty"`
	EthernetContext      *EthernetContext `protobuf:"bytes,3,opt,name=ethernet_context,json=ethernetContext,proto3" json:"ethernet_context,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ConnectionContext) Reset()         { *m = ConnectionContext{} }
func (m *ConnectionContext) String() string { return proto.CompactTextString(m) }
func (*ConnectionContext) ProtoMessage()    {}
func (*ConnectionContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_c30b3f1555e8b686, []int{8}
}

func (m *ConnectionContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectionContext.Unmarshal(m, b)
}
func (m *ConnectionContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectionContext.Marshal(b, m, deterministic)
}
func (m *ConnectionContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectionContext.Merge(m, src)
}
func (m *ConnectionContext) XXX_Size() int {
	return xxx_messageInfo_ConnectionContext.Size(m)
}
func (m *ConnectionContext) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectionContext.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectionContext proto.InternalMessageInfo

func (m *ConnectionContext) GetIpContext() *IPContext {
	if m != nil {
		return m.IpContext
	}
	return nil
}

func (m *ConnectionContext) GetDnsContext() *DNSContext {
	if m != nil {
		return m.DnsContext
	}
	return nil
}

func (m *ConnectionContext) GetEthernetContext() *EthernetContext {
	if m != nil {
		return m.EthernetContext
	}
	return nil
}

func init() {
	proto.RegisterEnum("connectioncontext.IpFamily_Family", IpFamily_Family_name, IpFamily_Family_value)
	proto.RegisterType((*IpNeighbor)(nil), "connectioncontext.IpNeighbor")
	proto.RegisterType((*Route)(nil), "connectioncontext.Route")
	proto.RegisterType((*IpFamily)(nil), "connectioncontext.IpFamily")
	proto.RegisterType((*ExtraPrefixRequest)(nil), "connectioncontext.ExtraPrefixRequest")
	proto.RegisterType((*IPContext)(nil), "connectioncontext.IPContext")
	proto.RegisterType((*DNSConfig)(nil), "connectioncontext.DNSConfig")
	proto.RegisterType((*EthernetContext)(nil), "connectioncontext.EthernetContext")
	proto.RegisterType((*DNSContext)(nil), "connectioncontext.DNSContext")
	proto.RegisterType((*ConnectionContext)(nil), "connectioncontext.ConnectionContext")
}

func init() { proto.RegisterFile("connectioncontext.proto", fileDescriptor_c30b3f1555e8b686) }

var fileDescriptor_c30b3f1555e8b686 = []byte{
	// 668 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x51, 0x4f, 0x13, 0x41,
	0x10, 0xb6, 0x2d, 0x94, 0xde, 0x14, 0xda, 0x72, 0x1a, 0xb9, 0x44, 0x50, 0x72, 0x11, 0xc5, 0x98,
	0xf0, 0x50, 0x0d, 0x26, 0x6a, 0x8c, 0x5a, 0xd0, 0x34, 0x11, 0xd2, 0x2c, 0x89, 0x1a, 0x5f, 0x2e,
	0xc7, 0xed, 0x40, 0x37, 0xa1, 0x7b, 0xc7, 0xee, 0x55, 0xeb, 0x0f, 0xf4, 0x77, 0xf8, 0x47, 0x7c,
	0x30, 0x3b, 0xbb, 0x7b, 0xa2, 0xad, 0xfa, 0xd4, 0xed, 0x37, 0xdf, 0xf7, 0xcd, 0xdc, 0xcc, 0xec,
	0xc2, 0x46, 0x96, 0x4b, 0x89, 0x59, 0x29, 0x72, 0x99, 0xe5, 0xb2, 0xc4, 0x59, 0xb9, 0x57, 0xa8,
	0xbc, 0xcc, 0xc3, 0xf5, 0xb9, 0x40, 0xfc, 0x16, 0x60, 0x58, 0x1c, 0xa3, 0x38, 0x1f, 0x9f, 0xe6,
	0x2a, 0xec, 0x40, 0x5d, 0x14, 0x51, 0x6d, 0xbb, 0xb6, 0x1b, 0xb0, 0xba, 0x28, 0xc2, 0x07, 0xd0,
	0x1b, 0xa7, 0x8a, 0x7f, 0x49, 0x15, 0x26, 0x29, 0xe7, 0x0a, 0xb5, 0x8e, 0xea, 0x14, 0xed, 0x7a,
	0xfc, 0x95, 0x85, 0xe3, 0x3b, 0xb0, 0xcc, 0xf2, 0x69, 0x89, 0xe1, 0x4d, 0x68, 0x16, 0x0a, 0xcf,
	0xc4, 0xcc, 0xf9, 0xb8, 0x7f, 0x31, 0x87, 0xd6, 0xb0, 0x78, 0x93, 0x4e, 0xc4, 0xc5, 0xd7, 0xf0,
	0x29, 0x34, 0xcf, 0xe8, 0x44, 0x9c, 0x4e, 0x3f, 0xde, 0x9b, 0x2f, 0xd9, 0x93, 0xf7, 0xec, 0x0f,
	0x73, 0x8a, 0x78, 0x13, 0x9a, 0xce, 0xa5, 0x05, 0x4b, 0xc3, 0xd1, 0xfb, 0xc7, 0xbd, 0x6b, 0xee,
	0xb4, 0xdf, 0xab, 0xc5, 0xdf, 0x6a, 0x10, 0x1e, 0xce, 0x4a, 0x95, 0x8e, 0x28, 0x2b, 0xc3, 0xcb,
	0x29, 0xea, 0x32, 0x7c, 0x0e, 0x6d, 0x53, 0x7f, 0x72, 0x25, 0x6b, 0xbb, 0x7f, 0xeb, 0x1f, 0x59,
	0x19, 0x18, 0xbe, 0x4b, 0xb4, 0x05, 0x60, 0x3f, 0x22, 0xb9, 0x40, 0x49, 0x0d, 0x58, 0x63, 0x81,
	0x45, 0xde, 0xa1, 0x0c, 0xef, 0x43, 0x57, 0xe1, 0xe5, 0x54, 0x28, 0xe4, 0x89, 0x9c, 0x4e, 0x4e,
	0x51, 0x45, 0x0d, 0xe2, 0x74, 0x3c, 0x7c, 0x4c, 0xa8, 0x69, 0xa7, 0xb2, 0x05, 0xfd, 0x62, 0x2e,
	0x11, 0xb3, 0x5b, 0xe1, 0x96, 0x1a, 0xff, 0x68, 0x40, 0x30, 0x1c, 0x0d, 0x6c, 0x55, 0xe1, 0x6d,
	0x68, 0x6b, 0x95, 0x25, 0xa2, 0xa0, 0x29, 0xb8, 0xc6, 0x06, 0x5a, 0x65, 0xc3, 0xc2, 0xf4, 0xdf,
	0xc4, 0xb9, 0x2e, 0xab, 0xb8, 0x1d, 0x51, 0xc0, 0x75, 0xe9, 0xe2, 0xf7, 0xa0, 0xeb, 0xf4, 0xbe,
	0x22, 0xaa, 0xb0, 0xc5, 0xd6, 0xc8, 0x83, 0x39, 0xd0, 0xf0, 0x9c, 0x4f, 0xc5, 0x5b, 0xb2, 0x3c,
	0xf2, 0xaa, 0x78, 0x4f, 0x00, 0x8c, 0x9f, 0x32, 0x03, 0xd7, 0xd1, 0xf2, 0x76, 0x63, 0xb7, 0xdd,
	0x8f, 0x16, 0x74, 0x93, 0x36, 0x82, 0x0a, 0xa5, 0x93, 0x36, 0x42, 0x93, 0xc0, 0x09, 0x9b, 0xff,
	0x13, 0x72, 0x5d, 0x3a, 0xe1, 0x43, 0x58, 0xc7, 0x59, 0x76, 0x31, 0xe5, 0xc8, 0x13, 0xdb, 0x79,
	0xd4, 0xd1, 0xca, 0x76, 0x63, 0x37, 0x60, 0x3d, 0x1f, 0x18, 0x39, 0x3c, 0x7c, 0x09, 0xab, 0xa2,
	0x48, 0xa4, 0xdb, 0x6a, 0x1d, 0xb5, 0x28, 0xcf, 0xd6, 0xc2, 0x71, 0xfb, 0xdd, 0x67, 0x6d, 0x51,
	0x9d, 0x75, 0xf8, 0x01, 0x6e, 0xa0, 0xd9, 0x22, 0x97, 0x2b, 0x71, 0xe3, 0x89, 0x02, 0x72, 0xda,
	0x59, 0xe0, 0x34, 0xbf, 0x74, 0x2c, 0xc4, 0xf9, 0x45, 0xdc, 0x81, 0xce, 0x55, 0x63, 0xd4, 0x11,
	0xd0, 0x47, 0xac, 0x5d, 0xe1, 0xa2, 0x8e, 0x3f, 0x42, 0x70, 0x70, 0x7c, 0x32, 0xc8, 0xe5, 0x99,
	0x38, 0x0f, 0xef, 0x42, 0x87, 0x4b, 0x9d, 0x68, 0x54, 0x9f, 0x51, 0x25, 0xa2, 0xd0, 0x51, 0x8d,
	0x34, 0xab, 0x5c, 0xea, 0x13, 0x02, 0x87, 0x85, 0x36, 0xce, 0x1a, 0x53, 0x95, 0x8d, 0x13, 0x9e,
	0x4f, 0x52, 0x21, 0xcd, 0x4d, 0x25, 0x67, 0x8b, 0x1e, 0x58, 0x30, 0x1e, 0x40, 0xf7, 0xb0, 0x1c,
	0xa3, 0x92, 0x58, 0xfa, 0xed, 0xda, 0x80, 0x15, 0x33, 0xcd, 0x49, 0x9a, 0xf9, 0x2b, 0xab, 0x55,
	0x76, 0x94, 0x66, 0x26, 0x60, 0xa6, 0x65, 0x02, 0x76, 0xa5, 0x9a, 0x5c, 0x97, 0x47, 0x69, 0x16,
	0x1f, 0x00, 0xd8, 0xf2, 0x48, 0xbf, 0x0f, 0x2b, 0x19, 0x55, 0x6a, 0x0b, 0x6b, 0xf7, 0x37, 0x17,
	0xf4, 0xa7, 0xfa, 0x1c, 0xe6, 0xc9, 0xf1, 0xf7, 0x1a, 0xac, 0x0f, 0x2a, 0xa2, 0x77, 0x7b, 0x06,
	0x20, 0x8a, 0xc4, 0xc9, 0xdc, 0x4d, 0x5d, 0x64, 0x58, 0xdd, 0x0e, 0x16, 0x88, 0xc2, 0x8b, 0x5f,
	0x40, 0xdb, 0xb4, 0xca, 0xab, 0xeb, 0xa4, 0xde, 0xfa, 0x6b, 0x39, 0x24, 0x07, 0x2e, 0xb5, 0xd7,
	0x1f, 0x41, 0x0f, 0x5d, 0x77, 0x2a, 0x93, 0x06, 0x99, 0x2c, 0x7a, 0xa2, 0xfe, 0x68, 0x24, 0xeb,
	0xe2, 0xef, 0xc0, 0xeb, 0xeb, 0x9f, 0xe6, 0x9f, 0xdc, 0xd3, 0x26, 0x3d, 0xc6, 0x8f, 0x7e, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x24, 0xbd, 0x94, 0xc2, 0xa7, 0x05, 0x00, 0x00,
}
