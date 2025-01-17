// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consignment/consignment.proto

package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description          string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight               int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers           []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId             string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Consignment) Reset()         { *m = Consignment{} }
func (m *Consignment) String() string { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()    {}
func (*Consignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_d51b1d6c06c6db39, []int{0}
}
func (m *Consignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consignment.Unmarshal(m, b)
}
func (m *Consignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consignment.Marshal(b, m, deterministic)
}
func (dst *Consignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consignment.Merge(dst, src)
}
func (m *Consignment) XXX_Size() int {
	return xxx_messageInfo_Consignment.Size(m)
}
func (m *Consignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Consignment.DiscardUnknown(m)
}

var xxx_messageInfo_Consignment proto.InternalMessageInfo

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId           string   `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin               string   `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId               string   `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Container) Reset()         { *m = Container{} }
func (m *Container) String() string { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()    {}
func (*Container) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_d51b1d6c06c6db39, []int{1}
}
func (m *Container) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Container.Unmarshal(m, b)
}
func (m *Container) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Container.Marshal(b, m, deterministic)
}
func (dst *Container) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Container.Merge(dst, src)
}
func (m *Container) XXX_Size() int {
	return xxx_messageInfo_Container.Size(m)
}
func (m *Container) XXX_DiscardUnknown() {
	xxx_messageInfo_Container.DiscardUnknown(m)
}

var xxx_messageInfo_Container proto.InternalMessageInfo

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRequest) Reset()         { *m = GetRequest{} }
func (m *GetRequest) String() string { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()    {}
func (*GetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_d51b1d6c06c6db39, []int{2}
}
func (m *GetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRequest.Unmarshal(m, b)
}
func (m *GetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRequest.Marshal(b, m, deterministic)
}
func (dst *GetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRequest.Merge(dst, src)
}
func (m *GetRequest) XXX_Size() int {
	return xxx_messageInfo_GetRequest.Size(m)
}
func (m *GetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetRequest proto.InternalMessageInfo

type Response struct {
	Created              bool           `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment          *Consignment   `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	Consignments         []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_consignment_d51b1d6c06c6db39, []int{3}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() {
	proto.RegisterFile("proto/consignment/consignment.proto", fileDescriptor_consignment_d51b1d6c06c6db39)
}

var fileDescriptor_consignment_d51b1d6c06c6db39 = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4e, 0xeb, 0x30,
	0x10, 0x87, 0x95, 0xfe, 0xcf, 0xa4, 0x7a, 0xd5, 0xf3, 0x02, 0x2c, 0x58, 0x10, 0xa5, 0x20, 0x75,
	0x15, 0xa4, 0x72, 0x84, 0x2c, 0xaa, 0x6c, 0xdd, 0x25, 0x12, 0xa8, 0xc4, 0xa3, 0xd4, 0x12, 0xb1,
	0x83, 0xed, 0x96, 0xab, 0xc1, 0x25, 0x38, 0x13, 0xaa, 0xd3, 0xb4, 0x46, 0xa8, 0xa8, 0x3b, 0xff,
	0x66, 0xfc, 0x4d, 0xbe, 0x8c, 0x0c, 0xd3, 0x5a, 0x2b, 0xab, 0xee, 0x0b, 0x25, 0x8d, 0x28, 0x65,
	0x85, 0xd2, 0xfa, 0xe7, 0xd4, 0x75, 0x09, 0x2d, 0x55, 0x5a, 0x89, 0x42, 0xab, 0xd4, 0xe8, 0x6d,
	0xea, 0xf5, 0x93, 0xcf, 0x00, 0xa2, 0xec, 0x98, 0xc9, 0x3f, 0xe8, 0x08, 0x4e, 0x83, 0x38, 0x98,
	0x85, 0xac, 0x23, 0x38, 0x89, 0x21, 0xe2, 0x68, 0x0a, 0x2d, 0x6a, 0x2b, 0x94, 0xa4, 0x1d, 0xd7,
	0xf0, 0x4b, 0xe4, 0x02, 0x06, 0xef, 0x28, 0xca, 0xb5, 0xa5, 0xdd, 0x38, 0x98, 0xf5, 0xd9, 0x3e,
	0x91, 0x0c, 0xa0, 0x50, 0xd2, 0xae, 0x84, 0x44, 0x6d, 0x68, 0x2f, 0xee, 0xce, 0xa2, 0xf9, 0x34,
	0x3d, 0x25, 0x92, 0x66, 0xed, 0x5d, 0xe6, 0x61, 0xe4, 0x1a, 0xc2, 0x2d, 0x1a, 0x83, 0xaf, 0xcf,
	0x82, 0xd3, 0xbe, 0xfb, 0xf8, 0xa8, 0x29, 0xe4, 0x3c, 0xa9, 0x20, 0x3c, 0x50, 0xbf, 0xc4, 0x6f,
	0x20, 0x2a, 0x36, 0xc6, 0xaa, 0x0a, 0xf5, 0x8e, 0x6d, 0xc4, 0xa1, 0x2d, 0xe5, 0x7c, 0xe7, 0xad,
	0xb4, 0x28, 0x85, 0x74, 0xde, 0x21, 0xdb, 0x27, 0x72, 0x09, 0xc3, 0x8d, 0x69, 0xa0, 0x5e, 0xd3,
	0xd8, 0xc5, 0x9c, 0x27, 0x63, 0x80, 0x05, 0x5a, 0x86, 0x6f, 0x1b, 0x34, 0x36, 0xf9, 0x08, 0x60,
	0xc4, 0xd0, 0xd4, 0x4a, 0x1a, 0x24, 0x14, 0x86, 0x85, 0xc6, 0x95, 0xc5, 0xc6, 0x60, 0xc4, 0xda,
	0x48, 0x16, 0x10, 0x79, 0x7f, 0xe9, 0x34, 0xa2, 0xf9, 0xdd, 0x9f, 0x6b, 0x68, 0xcf, 0xcc, 0x27,
	0x49, 0x0e, 0x63, 0x2f, 0x1a, 0xda, 0x75, 0x0b, 0x3d, 0x73, 0xd2, 0x0f, 0x74, 0xfe, 0x15, 0xc0,
	0x64, 0xb9, 0x16, 0x75, 0x2d, 0x64, 0xb9, 0x44, 0xbd, 0x15, 0x05, 0x92, 0x27, 0xf8, 0x9f, 0x39,
	0x65, 0xff, 0x31, 0x9c, 0x37, 0xfd, 0x2a, 0x39, 0x7d, 0xed, 0xb0, 0xa1, 0x47, 0x98, 0x2c, 0xd0,
	0x7a, 0x94, 0x21, 0xb7, 0xa7, 0xb1, 0xe3, 0x9e, 0xcf, 0x19, 0xfe, 0x32, 0x70, 0xaf, 0xfc, 0xe1,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x08, 0xc8, 0xb2, 0x38, 0x0c, 0x03, 0x00, 0x00,
}
