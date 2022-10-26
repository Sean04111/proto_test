// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: order.proto

package proto_test

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderSendClient is the client API for OrderSend service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderSendClient interface {
	OrderDely(ctx context.Context, opts ...grpc.CallOption) (OrderSend_OrderDelyClient, error)
}

type orderSendClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderSendClient(cc grpc.ClientConnInterface) OrderSendClient {
	return &orderSendClient{cc}
}

func (c *orderSendClient) OrderDely(ctx context.Context, opts ...grpc.CallOption) (OrderSend_OrderDelyClient, error) {
	stream, err := c.cc.NewStream(ctx, &OrderSend_ServiceDesc.Streams[0], "/Order.OrderSend/OrderDely", opts...)
	if err != nil {
		return nil, err
	}
	x := &orderSendOrderDelyClient{stream}
	return x, nil
}

type OrderSend_OrderDelyClient interface {
	Send(*OrderInfor) error
	Recv() (*Car, error)
	grpc.ClientStream
}

type orderSendOrderDelyClient struct {
	grpc.ClientStream
}

func (x *orderSendOrderDelyClient) Send(m *OrderInfor) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderSendOrderDelyClient) Recv() (*Car, error) {
	m := new(Car)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderSendServer is the server API for OrderSend service.
// All implementations must embed UnimplementedOrderSendServer
// for forward compatibility
type OrderSendServer interface {
	OrderDely(OrderSend_OrderDelyServer) error
	mustEmbedUnimplementedOrderSendServer()
}

// UnimplementedOrderSendServer must be embedded to have forward compatible implementations.
type UnimplementedOrderSendServer struct {
}

func (UnimplementedOrderSendServer) OrderDely(OrderSend_OrderDelyServer) error {
	return status.Errorf(codes.Unimplemented, "method OrderDely not implemented")
}
func (UnimplementedOrderSendServer) mustEmbedUnimplementedOrderSendServer() {}

// UnsafeOrderSendServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderSendServer will
// result in compilation errors.
type UnsafeOrderSendServer interface {
	mustEmbedUnimplementedOrderSendServer()
}

func RegisterOrderSendServer(s grpc.ServiceRegistrar, srv OrderSendServer) {
	s.RegisterService(&OrderSend_ServiceDesc, srv)
}

func _OrderSend_OrderDely_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderSendServer).OrderDely(&orderSendOrderDelyServer{stream})
}

type OrderSend_OrderDelyServer interface {
	Send(*Car) error
	Recv() (*OrderInfor, error)
	grpc.ServerStream
}

type orderSendOrderDelyServer struct {
	grpc.ServerStream
}

func (x *orderSendOrderDelyServer) Send(m *Car) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderSendOrderDelyServer) Recv() (*OrderInfor, error) {
	m := new(OrderInfor)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderSend_ServiceDesc is the grpc.ServiceDesc for OrderSend service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderSend_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Order.OrderSend",
	HandlerType: (*OrderSendServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OrderDely",
			Handler:       _OrderSend_OrderDely_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "order.proto",
}