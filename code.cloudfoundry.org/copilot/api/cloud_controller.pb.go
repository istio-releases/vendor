// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cloud_controller.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	cloud_controller.proto
	istio.proto

It has these top-level messages:
	Route
	UpsertRouteRequest
	UpsertRouteResponse
	DeleteRouteRequest
	DeleteRouteResponse
	CapiProcess
	RouteMapping
	MapRouteRequest
	MapRouteResponse
	UnmapRouteRequest
	UnmapRouteResponse
	HealthRequest
	HealthResponse
	RoutesRequest
	RoutesResponse
	BackendSet
	Backend
*/
package api

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

type Route struct {
	Guid string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	Host string `protobuf:"bytes,2,opt,name=host" json:"host,omitempty"`
}

func (m *Route) Reset()                    { *m = Route{} }
func (m *Route) String() string            { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()               {}
func (*Route) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Route) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *Route) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type UpsertRouteRequest struct {
	Route *Route `protobuf:"bytes,1,opt,name=route" json:"route,omitempty"`
}

func (m *UpsertRouteRequest) Reset()                    { *m = UpsertRouteRequest{} }
func (m *UpsertRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*UpsertRouteRequest) ProtoMessage()               {}
func (*UpsertRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *UpsertRouteRequest) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

type UpsertRouteResponse struct {
}

func (m *UpsertRouteResponse) Reset()                    { *m = UpsertRouteResponse{} }
func (m *UpsertRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*UpsertRouteResponse) ProtoMessage()               {}
func (*UpsertRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type DeleteRouteRequest struct {
	Guid string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
}

func (m *DeleteRouteRequest) Reset()                    { *m = DeleteRouteRequest{} }
func (m *DeleteRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRouteRequest) ProtoMessage()               {}
func (*DeleteRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DeleteRouteRequest) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

type DeleteRouteResponse struct {
}

func (m *DeleteRouteResponse) Reset()                    { *m = DeleteRouteResponse{} }
func (m *DeleteRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteRouteResponse) ProtoMessage()               {}
func (*DeleteRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type CapiProcess struct {
	Guid             string `protobuf:"bytes,1,opt,name=guid" json:"guid,omitempty"`
	DiegoProcessGuid string `protobuf:"bytes,2,opt,name=diego_process_guid,json=diegoProcessGuid" json:"diego_process_guid,omitempty"`
}

func (m *CapiProcess) Reset()                    { *m = CapiProcess{} }
func (m *CapiProcess) String() string            { return proto.CompactTextString(m) }
func (*CapiProcess) ProtoMessage()               {}
func (*CapiProcess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CapiProcess) GetGuid() string {
	if m != nil {
		return m.Guid
	}
	return ""
}

func (m *CapiProcess) GetDiegoProcessGuid() string {
	if m != nil {
		return m.DiegoProcessGuid
	}
	return ""
}

type RouteMapping struct {
	CapiProcess *CapiProcess `protobuf:"bytes,1,opt,name=capi_process,json=capiProcess" json:"capi_process,omitempty"`
	RouteGuid   string       `protobuf:"bytes,2,opt,name=route_guid,json=routeGuid" json:"route_guid,omitempty"`
}

func (m *RouteMapping) Reset()                    { *m = RouteMapping{} }
func (m *RouteMapping) String() string            { return proto.CompactTextString(m) }
func (*RouteMapping) ProtoMessage()               {}
func (*RouteMapping) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RouteMapping) GetCapiProcess() *CapiProcess {
	if m != nil {
		return m.CapiProcess
	}
	return nil
}

func (m *RouteMapping) GetRouteGuid() string {
	if m != nil {
		return m.RouteGuid
	}
	return ""
}

type MapRouteRequest struct {
	RouteMapping *RouteMapping `protobuf:"bytes,1,opt,name=route_mapping,json=routeMapping" json:"route_mapping,omitempty"`
}

func (m *MapRouteRequest) Reset()                    { *m = MapRouteRequest{} }
func (m *MapRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*MapRouteRequest) ProtoMessage()               {}
func (*MapRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *MapRouteRequest) GetRouteMapping() *RouteMapping {
	if m != nil {
		return m.RouteMapping
	}
	return nil
}

type MapRouteResponse struct {
}

func (m *MapRouteResponse) Reset()                    { *m = MapRouteResponse{} }
func (m *MapRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*MapRouteResponse) ProtoMessage()               {}
func (*MapRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type UnmapRouteRequest struct {
	CapiProcessGuid string `protobuf:"bytes,1,opt,name=capi_process_guid,json=capiProcessGuid" json:"capi_process_guid,omitempty"`
	RouteGuid       string `protobuf:"bytes,2,opt,name=route_guid,json=routeGuid" json:"route_guid,omitempty"`
}

func (m *UnmapRouteRequest) Reset()                    { *m = UnmapRouteRequest{} }
func (m *UnmapRouteRequest) String() string            { return proto.CompactTextString(m) }
func (*UnmapRouteRequest) ProtoMessage()               {}
func (*UnmapRouteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UnmapRouteRequest) GetCapiProcessGuid() string {
	if m != nil {
		return m.CapiProcessGuid
	}
	return ""
}

func (m *UnmapRouteRequest) GetRouteGuid() string {
	if m != nil {
		return m.RouteGuid
	}
	return ""
}

type UnmapRouteResponse struct {
}

func (m *UnmapRouteResponse) Reset()                    { *m = UnmapRouteResponse{} }
func (m *UnmapRouteResponse) String() string            { return proto.CompactTextString(m) }
func (*UnmapRouteResponse) ProtoMessage()               {}
func (*UnmapRouteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Route)(nil), "api.Route")
	proto.RegisterType((*UpsertRouteRequest)(nil), "api.UpsertRouteRequest")
	proto.RegisterType((*UpsertRouteResponse)(nil), "api.UpsertRouteResponse")
	proto.RegisterType((*DeleteRouteRequest)(nil), "api.DeleteRouteRequest")
	proto.RegisterType((*DeleteRouteResponse)(nil), "api.DeleteRouteResponse")
	proto.RegisterType((*CapiProcess)(nil), "api.CapiProcess")
	proto.RegisterType((*RouteMapping)(nil), "api.RouteMapping")
	proto.RegisterType((*MapRouteRequest)(nil), "api.MapRouteRequest")
	proto.RegisterType((*MapRouteResponse)(nil), "api.MapRouteResponse")
	proto.RegisterType((*UnmapRouteRequest)(nil), "api.UnmapRouteRequest")
	proto.RegisterType((*UnmapRouteResponse)(nil), "api.UnmapRouteResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CloudControllerCopilot service

type CloudControllerCopilotClient interface {
	UpsertRoute(ctx context.Context, in *UpsertRouteRequest, opts ...grpc.CallOption) (*UpsertRouteResponse, error)
	DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error)
	MapRoute(ctx context.Context, in *MapRouteRequest, opts ...grpc.CallOption) (*MapRouteResponse, error)
	UnmapRoute(ctx context.Context, in *UnmapRouteRequest, opts ...grpc.CallOption) (*UnmapRouteResponse, error)
}

type cloudControllerCopilotClient struct {
	cc *grpc.ClientConn
}

func NewCloudControllerCopilotClient(cc *grpc.ClientConn) CloudControllerCopilotClient {
	return &cloudControllerCopilotClient{cc}
}

func (c *cloudControllerCopilotClient) UpsertRoute(ctx context.Context, in *UpsertRouteRequest, opts ...grpc.CallOption) (*UpsertRouteResponse, error) {
	out := new(UpsertRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/UpsertRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) DeleteRoute(ctx context.Context, in *DeleteRouteRequest, opts ...grpc.CallOption) (*DeleteRouteResponse, error) {
	out := new(DeleteRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/DeleteRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) MapRoute(ctx context.Context, in *MapRouteRequest, opts ...grpc.CallOption) (*MapRouteResponse, error) {
	out := new(MapRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/MapRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cloudControllerCopilotClient) UnmapRoute(ctx context.Context, in *UnmapRouteRequest, opts ...grpc.CallOption) (*UnmapRouteResponse, error) {
	out := new(UnmapRouteResponse)
	err := grpc.Invoke(ctx, "/api.CloudControllerCopilot/UnmapRoute", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CloudControllerCopilot service

type CloudControllerCopilotServer interface {
	UpsertRoute(context.Context, *UpsertRouteRequest) (*UpsertRouteResponse, error)
	DeleteRoute(context.Context, *DeleteRouteRequest) (*DeleteRouteResponse, error)
	MapRoute(context.Context, *MapRouteRequest) (*MapRouteResponse, error)
	UnmapRoute(context.Context, *UnmapRouteRequest) (*UnmapRouteResponse, error)
}

func RegisterCloudControllerCopilotServer(s *grpc.Server, srv CloudControllerCopilotServer) {
	s.RegisterService(&_CloudControllerCopilot_serviceDesc, srv)
}

func _CloudControllerCopilot_UpsertRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).UpsertRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/UpsertRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).UpsertRoute(ctx, req.(*UpsertRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_DeleteRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).DeleteRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/DeleteRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).DeleteRoute(ctx, req.(*DeleteRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_MapRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).MapRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/MapRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).MapRoute(ctx, req.(*MapRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CloudControllerCopilot_UnmapRoute_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmapRouteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudControllerCopilotServer).UnmapRoute(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.CloudControllerCopilot/UnmapRoute",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudControllerCopilotServer).UnmapRoute(ctx, req.(*UnmapRouteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CloudControllerCopilot_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.CloudControllerCopilot",
	HandlerType: (*CloudControllerCopilotServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertRoute",
			Handler:    _CloudControllerCopilot_UpsertRoute_Handler,
		},
		{
			MethodName: "DeleteRoute",
			Handler:    _CloudControllerCopilot_DeleteRoute_Handler,
		},
		{
			MethodName: "MapRoute",
			Handler:    _CloudControllerCopilot_MapRoute_Handler,
		},
		{
			MethodName: "UnmapRoute",
			Handler:    _CloudControllerCopilot_UnmapRoute_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cloud_controller.proto",
}

func init() { proto.RegisterFile("cloud_controller.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x5b, 0x6b, 0xea, 0x40,
	0x10, 0x56, 0xcf, 0xf1, 0x70, 0x9c, 0x58, 0xd4, 0xa9, 0x37, 0x84, 0x82, 0xe4, 0x49, 0x4a, 0xb1,
	0xa0, 0x20, 0xf4, 0xa9, 0xd0, 0x14, 0x4a, 0x1f, 0xa4, 0x25, 0xe0, 0x6b, 0x43, 0x8c, 0x8b, 0x5d,
	0x88, 0xd9, 0x6d, 0xb2, 0xf9, 0x37, 0xfd, 0xb1, 0x65, 0x2f, 0xea, 0x26, 0x11, 0xfa, 0x96, 0x9d,
	0x9d, 0xef, 0x36, 0x3b, 0x81, 0x61, 0x14, 0xb3, 0x7c, 0x17, 0x44, 0x2c, 0x11, 0x29, 0x8b, 0x63,
	0x92, 0xce, 0x79, 0xca, 0x04, 0xc3, 0x3f, 0x21, 0xa7, 0xee, 0x3d, 0x34, 0x7d, 0x96, 0x0b, 0x82,
	0x08, 0x7f, 0xf7, 0x39, 0xdd, 0x8d, 0xeb, 0xd3, 0xfa, 0xac, 0xe5, 0xab, 0x6f, 0x59, 0xfb, 0x64,
	0x99, 0x18, 0x37, 0x74, 0x4d, 0x7e, 0xbb, 0x2b, 0xc0, 0x0d, 0xcf, 0x48, 0x2a, 0x14, 0xcc, 0x27,
	0x5f, 0x39, 0xc9, 0x04, 0x4e, 0xa1, 0x99, 0xca, 0xb3, 0x82, 0x3b, 0x0b, 0x98, 0x87, 0x9c, 0xce,
	0x75, 0x87, 0xbe, 0x70, 0x07, 0x70, 0x5d, 0xc0, 0x65, 0x9c, 0x25, 0x19, 0x71, 0x67, 0x80, 0xcf,
	0x24, 0x26, 0x82, 0x14, 0xe8, 0x2e, 0x98, 0x91, 0x04, 0x85, 0x4e, 0x43, 0xf0, 0x06, 0x8e, 0x17,
	0x72, 0xfa, 0x9e, 0xb2, 0x88, 0x64, 0xd9, 0xc5, 0x18, 0x77, 0x80, 0x3b, 0x4a, 0xf6, 0x2c, 0xe0,
	0xba, 0x29, 0x50, 0x1d, 0x3a, 0x54, 0x57, 0xdd, 0x18, 0xf4, 0x8b, 0xd4, 0xd9, 0x42, 0x5b, 0x29,
	0xac, 0x43, 0xce, 0x69, 0xb2, 0xc7, 0x25, 0xb4, 0xa3, 0x90, 0xd3, 0x23, 0xd8, 0x24, 0xec, 0xaa,
	0x84, 0x96, 0xb2, 0xef, 0x44, 0x96, 0x8d, 0x1b, 0x00, 0x15, 0xdb, 0x96, 0x6a, 0xa9, 0x8a, 0xd2,
	0x78, 0x85, 0xce, 0x3a, 0xe4, 0x85, 0xc8, 0x2b, 0xb8, 0xd2, 0x88, 0x83, 0xd6, 0x35, 0x3a, 0xbd,
	0xf3, 0x24, 0x8d, 0x21, 0xbf, 0x9d, 0x5a, 0x27, 0x17, 0xa1, 0x7b, 0xa6, 0x32, 0x33, 0xf9, 0x80,
	0xde, 0x26, 0x39, 0x94, 0x04, 0x6e, 0xa1, 0x67, 0xe7, 0x08, 0xac, 0x31, 0x75, 0x2c, 0xeb, 0xd2,
	0xdf, 0x6f, 0xf6, 0xfb, 0x80, 0x36, 0xbf, 0x56, 0x5d, 0x7c, 0x37, 0x60, 0xe8, 0xc9, 0x55, 0xf3,
	0x4e, 0x9b, 0xe6, 0x31, 0x4e, 0x63, 0x26, 0xf0, 0x09, 0x1c, 0xeb, 0xf1, 0x71, 0xa4, 0x42, 0x55,
	0xd7, 0x68, 0x32, 0xae, 0x5e, 0x98, 0x48, 0x35, 0xc9, 0x61, 0xbd, 0xbf, 0xe1, 0xa8, 0xee, 0x8e,
	0xe1, 0xb8, 0xb4, 0x2a, 0x35, 0x7c, 0x80, 0xff, 0xc7, 0x61, 0x61, 0x5f, 0xf5, 0x95, 0x9e, 0x61,
	0x32, 0x28, 0x55, 0x4f, 0xd0, 0x47, 0x80, 0x73, 0x66, 0x1c, 0x6a, 0xa3, 0xe5, 0x21, 0x4f, 0x46,
	0x95, 0xfa, 0x91, 0x60, 0xfb, 0x4f, 0xfd, 0x75, 0xcb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x14,
	0x16, 0x0e, 0x5d, 0x8f, 0x03, 0x00, 0x00,
}
