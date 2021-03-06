// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/company/company.proto

/*
Package company is a generated protocol buffer package.

It is generated from these files:
	api/company/company.proto

It has these top-level messages:
	CompanyData
	CompanyListRrequest
	CompanyListResponse
	CompanyDetailRequest
	CompanyDetailResponse
*/
package company

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type CompanyData struct {
	Id          int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Descrpition string `protobuf:"bytes,3,opt,name=descrpition" json:"descrpition,omitempty"`
	LogoUrl     string `protobuf:"bytes,4,opt,name=logo_url,json=logoUrl" json:"logo_url,omitempty"`
}

func (m *CompanyData) Reset()                    { *m = CompanyData{} }
func (m *CompanyData) String() string            { return proto.CompactTextString(m) }
func (*CompanyData) ProtoMessage()               {}
func (*CompanyData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CompanyData) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CompanyData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CompanyData) GetDescrpition() string {
	if m != nil {
		return m.Descrpition
	}
	return ""
}

func (m *CompanyData) GetLogoUrl() string {
	if m != nil {
		return m.LogoUrl
	}
	return ""
}

type CompanyListRrequest struct {
}

func (m *CompanyListRrequest) Reset()                    { *m = CompanyListRrequest{} }
func (m *CompanyListRrequest) String() string            { return proto.CompactTextString(m) }
func (*CompanyListRrequest) ProtoMessage()               {}
func (*CompanyListRrequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CompanyListResponse struct {
	Companies []*CompanyData `protobuf:"bytes,1,rep,name=companies" json:"companies,omitempty"`
}

func (m *CompanyListResponse) Reset()                    { *m = CompanyListResponse{} }
func (m *CompanyListResponse) String() string            { return proto.CompactTextString(m) }
func (*CompanyListResponse) ProtoMessage()               {}
func (*CompanyListResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CompanyListResponse) GetCompanies() []*CompanyData {
	if m != nil {
		return m.Companies
	}
	return nil
}

type CompanyDetailRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *CompanyDetailRequest) Reset()                    { *m = CompanyDetailRequest{} }
func (m *CompanyDetailRequest) String() string            { return proto.CompactTextString(m) }
func (*CompanyDetailRequest) ProtoMessage()               {}
func (*CompanyDetailRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CompanyDetailRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type CompanyDetailResponse struct {
	Company *CompanyData `protobuf:"bytes,1,opt,name=company" json:"company,omitempty"`
}

func (m *CompanyDetailResponse) Reset()                    { *m = CompanyDetailResponse{} }
func (m *CompanyDetailResponse) String() string            { return proto.CompactTextString(m) }
func (*CompanyDetailResponse) ProtoMessage()               {}
func (*CompanyDetailResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CompanyDetailResponse) GetCompany() *CompanyData {
	if m != nil {
		return m.Company
	}
	return nil
}

func init() {
	proto.RegisterType((*CompanyData)(nil), "company.CompanyData")
	proto.RegisterType((*CompanyListRrequest)(nil), "company.CompanyListRrequest")
	proto.RegisterType((*CompanyListResponse)(nil), "company.CompanyListResponse")
	proto.RegisterType((*CompanyDetailRequest)(nil), "company.CompanyDetailRequest")
	proto.RegisterType((*CompanyDetailResponse)(nil), "company.CompanyDetailResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Company service

type CompanyClient interface {
	List(ctx context.Context, in *CompanyListRrequest, opts ...grpc.CallOption) (*CompanyListResponse, error)
	GetDetail(ctx context.Context, in *CompanyDetailRequest, opts ...grpc.CallOption) (*CompanyDetailResponse, error)
}

type companyClient struct {
	cc *grpc.ClientConn
}

func NewCompanyClient(cc *grpc.ClientConn) CompanyClient {
	return &companyClient{cc}
}

func (c *companyClient) List(ctx context.Context, in *CompanyListRrequest, opts ...grpc.CallOption) (*CompanyListResponse, error) {
	out := new(CompanyListResponse)
	err := grpc.Invoke(ctx, "/company.Company/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *companyClient) GetDetail(ctx context.Context, in *CompanyDetailRequest, opts ...grpc.CallOption) (*CompanyDetailResponse, error) {
	out := new(CompanyDetailResponse)
	err := grpc.Invoke(ctx, "/company.Company/GetDetail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Company service

type CompanyServer interface {
	List(context.Context, *CompanyListRrequest) (*CompanyListResponse, error)
	GetDetail(context.Context, *CompanyDetailRequest) (*CompanyDetailResponse, error)
}

func RegisterCompanyServer(s *grpc.Server, srv CompanyServer) {
	s.RegisterService(&_Company_serviceDesc, srv)
}

func _Company_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyListRrequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.Company/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).List(ctx, req.(*CompanyListRrequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Company_GetDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompanyDetailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompanyServer).GetDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/company.Company/GetDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompanyServer).GetDetail(ctx, req.(*CompanyDetailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Company_serviceDesc = grpc.ServiceDesc{
	ServiceName: "company.Company",
	HandlerType: (*CompanyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Company_List_Handler,
		},
		{
			MethodName: "GetDetail",
			Handler:    _Company_GetDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/company/company.proto",
}

func init() { proto.RegisterFile("api/company/company.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 272 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0xb7, 0x5a, 0x3b, 0x05, 0x0f, 0xe3, 0x2e, 0xc4, 0x45, 0xa5, 0xf4, 0x20, 0x3d,
	0x55, 0xa8, 0x6f, 0xa0, 0xc2, 0xa2, 0x78, 0x2a, 0x78, 0x96, 0xec, 0x36, 0x48, 0xa0, 0x9b, 0xc4,
	0x24, 0x7b, 0xd8, 0x47, 0xf1, 0x6d, 0x65, 0x93, 0x06, 0x6b, 0x75, 0x4f, 0x9d, 0x99, 0xff, 0x6f,
	0xbe, 0xf9, 0x13, 0x20, 0x6a, 0x7d, 0xb7, 0x91, 0x5b, 0x25, 0xe8, 0xde, 0x7d, 0xa9, 0xd8, 0xd7,
	0x4a, 0x4b, 0x2b, 0x31, 0x1d, 0xda, 0x52, 0x40, 0xfe, 0xe8, 0xcb, 0x27, 0x6a, 0x29, 0x9e, 0x43,
	0xcc, 0x3b, 0x12, 0x15, 0x51, 0x75, 0xd2, 0xc6, 0xbc, 0x43, 0x84, 0x44, 0xd0, 0x2d, 0x23, 0x71,
	0x11, 0x55, 0x59, 0xeb, 0x6a, 0x2c, 0x20, 0xef, 0x98, 0xd9, 0x68, 0xc5, 0x2d, 0x97, 0x82, 0xcc,
	0x9c, 0x34, 0x1e, 0xe1, 0x25, 0x9c, 0xf5, 0xf2, 0x43, 0xbe, 0xef, 0x74, 0x4f, 0x12, 0x27, 0xa7,
	0x87, 0xfe, 0x4d, 0xf7, 0xe5, 0x02, 0x2e, 0x06, 0xde, 0x2b, 0x37, 0xb6, 0xd5, 0xec, 0x73, 0xc7,
	0x8c, 0x2d, 0x9f, 0x7f, 0x8f, 0x99, 0x51, 0x52, 0x18, 0x86, 0x0d, 0x64, 0x7e, 0x51, 0xce, 0x0c,
	0x89, 0x8a, 0x59, 0x95, 0x37, 0xf3, 0x3a, 0x24, 0x19, 0xed, 0xdd, 0xfe, 0xd8, 0xca, 0x5b, 0x98,
	0x07, 0x85, 0x59, 0xca, 0xfb, 0xd6, 0x23, 0xa6, 0xd1, 0xca, 0x15, 0x2c, 0x26, 0xbe, 0x01, 0x5a,
	0x43, 0xb8, 0x1d, 0xe7, 0x3e, 0x86, 0x0c, 0xa6, 0xe6, 0x2b, 0x82, 0x74, 0x10, 0xf0, 0x01, 0x92,
	0x43, 0x00, 0xbc, 0x9a, 0xfe, 0x32, 0x4e, 0xbb, 0xfc, 0x5f, 0x0d, 0xfc, 0x17, 0xc8, 0x56, 0xcc,
	0xfa, 0xa5, 0xf0, 0xfa, 0x0f, 0x7b, 0x1c, 0x6a, 0x79, 0x73, 0x4c, 0xf6, 0x67, 0xad, 0x4f, 0xdd,
	0x73, 0xdf, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9f, 0x9b, 0xb8, 0xc5, 0x0a, 0x02, 0x00, 0x00,
}
