// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: participation/participation.proto

package participation

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ParticipationServiceClient is the client API for ParticipationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParticipationServiceClient interface {
	InviteToParty(ctx context.Context, in *InviteToPartyRequest, opts ...grpc.CallOption) (*PartyInvite, error)
	DeclinePartyInvite(ctx context.Context, in *DeclinePartyInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AcceptPartyInvite(ctx context.Context, in *DeclinePartyInviteRequest, opts ...grpc.CallOption) (*PartyParticipant, error)
	GetUserInvites(ctx context.Context, in *GetUserInvitesRequest, opts ...grpc.CallOption) (*PagedPartyInvites, error)
	JoinParty(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*PartyParticipant, error)
	LeaveParty(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetPartyParticipant(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*PartyParticipant, error)
	GetPartyParticipants(ctx context.Context, in *GetPartyParticipantsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error)
	GetPartyRequests(ctx context.Context, in *GetPartyParticipantsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error)
	GetUserParticipations(ctx context.Context, in *GetUserParticipationsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error)
	GetManyUserParticipations(ctx context.Context, in *GetManyUserParticipationsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error)
}

type participationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParticipationServiceClient(cc grpc.ClientConnInterface) ParticipationServiceClient {
	return &participationServiceClient{cc}
}

func (c *participationServiceClient) InviteToParty(ctx context.Context, in *InviteToPartyRequest, opts ...grpc.CallOption) (*PartyInvite, error) {
	out := new(PartyInvite)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/InviteToParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) DeclinePartyInvite(ctx context.Context, in *DeclinePartyInviteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/DeclinePartyInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) AcceptPartyInvite(ctx context.Context, in *DeclinePartyInviteRequest, opts ...grpc.CallOption) (*PartyParticipant, error) {
	out := new(PartyParticipant)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/AcceptPartyInvite", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetUserInvites(ctx context.Context, in *GetUserInvitesRequest, opts ...grpc.CallOption) (*PagedPartyInvites, error) {
	out := new(PagedPartyInvites)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetUserInvites", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) JoinParty(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*PartyParticipant, error) {
	out := new(PartyParticipant)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/JoinParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) LeaveParty(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/LeaveParty", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetPartyParticipant(ctx context.Context, in *UserPartyRequest, opts ...grpc.CallOption) (*PartyParticipant, error) {
	out := new(PartyParticipant)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetPartyParticipant", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetPartyParticipants(ctx context.Context, in *GetPartyParticipantsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error) {
	out := new(PagedPartyParticipants)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetPartyParticipants", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetPartyRequests(ctx context.Context, in *GetPartyParticipantsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error) {
	out := new(PagedPartyParticipants)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetPartyRequests", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetUserParticipations(ctx context.Context, in *GetUserParticipationsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error) {
	out := new(PagedPartyParticipants)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetUserParticipations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *participationServiceClient) GetManyUserParticipations(ctx context.Context, in *GetManyUserParticipationsRequest, opts ...grpc.CallOption) (*PagedPartyParticipants, error) {
	out := new(PagedPartyParticipants)
	err := c.cc.Invoke(ctx, "/participation.ParticipationService/GetManyUserParticipations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParticipationServiceServer is the server API for ParticipationService service.
// All implementations must embed UnimplementedParticipationServiceServer
// for forward compatibility
type ParticipationServiceServer interface {
	InviteToParty(context.Context, *InviteToPartyRequest) (*PartyInvite, error)
	DeclinePartyInvite(context.Context, *DeclinePartyInviteRequest) (*emptypb.Empty, error)
	AcceptPartyInvite(context.Context, *DeclinePartyInviteRequest) (*PartyParticipant, error)
	GetUserInvites(context.Context, *GetUserInvitesRequest) (*PagedPartyInvites, error)
	JoinParty(context.Context, *UserPartyRequest) (*PartyParticipant, error)
	LeaveParty(context.Context, *UserPartyRequest) (*emptypb.Empty, error)
	GetPartyParticipant(context.Context, *UserPartyRequest) (*PartyParticipant, error)
	GetPartyParticipants(context.Context, *GetPartyParticipantsRequest) (*PagedPartyParticipants, error)
	GetPartyRequests(context.Context, *GetPartyParticipantsRequest) (*PagedPartyParticipants, error)
	GetUserParticipations(context.Context, *GetUserParticipationsRequest) (*PagedPartyParticipants, error)
	GetManyUserParticipations(context.Context, *GetManyUserParticipationsRequest) (*PagedPartyParticipants, error)
	mustEmbedUnimplementedParticipationServiceServer()
}

// UnimplementedParticipationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedParticipationServiceServer struct {
}

func (UnimplementedParticipationServiceServer) InviteToParty(context.Context, *InviteToPartyRequest) (*PartyInvite, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteToParty not implemented")
}
func (UnimplementedParticipationServiceServer) DeclinePartyInvite(context.Context, *DeclinePartyInviteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeclinePartyInvite not implemented")
}
func (UnimplementedParticipationServiceServer) AcceptPartyInvite(context.Context, *DeclinePartyInviteRequest) (*PartyParticipant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AcceptPartyInvite not implemented")
}
func (UnimplementedParticipationServiceServer) GetUserInvites(context.Context, *GetUserInvitesRequest) (*PagedPartyInvites, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInvites not implemented")
}
func (UnimplementedParticipationServiceServer) JoinParty(context.Context, *UserPartyRequest) (*PartyParticipant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinParty not implemented")
}
func (UnimplementedParticipationServiceServer) LeaveParty(context.Context, *UserPartyRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveParty not implemented")
}
func (UnimplementedParticipationServiceServer) GetPartyParticipant(context.Context, *UserPartyRequest) (*PartyParticipant, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyParticipant not implemented")
}
func (UnimplementedParticipationServiceServer) GetPartyParticipants(context.Context, *GetPartyParticipantsRequest) (*PagedPartyParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyParticipants not implemented")
}
func (UnimplementedParticipationServiceServer) GetPartyRequests(context.Context, *GetPartyParticipantsRequest) (*PagedPartyParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartyRequests not implemented")
}
func (UnimplementedParticipationServiceServer) GetUserParticipations(context.Context, *GetUserParticipationsRequest) (*PagedPartyParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserParticipations not implemented")
}
func (UnimplementedParticipationServiceServer) GetManyUserParticipations(context.Context, *GetManyUserParticipationsRequest) (*PagedPartyParticipants, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetManyUserParticipations not implemented")
}
func (UnimplementedParticipationServiceServer) mustEmbedUnimplementedParticipationServiceServer() {}

// UnsafeParticipationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParticipationServiceServer will
// result in compilation errors.
type UnsafeParticipationServiceServer interface {
	mustEmbedUnimplementedParticipationServiceServer()
}

func RegisterParticipationServiceServer(s grpc.ServiceRegistrar, srv ParticipationServiceServer) {
	s.RegisterService(&ParticipationService_ServiceDesc, srv)
}

func _ParticipationService_InviteToParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteToPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).InviteToParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/InviteToParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).InviteToParty(ctx, req.(*InviteToPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_DeclinePartyInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclinePartyInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).DeclinePartyInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/DeclinePartyInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).DeclinePartyInvite(ctx, req.(*DeclinePartyInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_AcceptPartyInvite_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeclinePartyInviteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).AcceptPartyInvite(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/AcceptPartyInvite",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).AcceptPartyInvite(ctx, req.(*DeclinePartyInviteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetUserInvites_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInvitesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetUserInvites(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetUserInvites",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetUserInvites(ctx, req.(*GetUserInvitesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_JoinParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).JoinParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/JoinParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).JoinParty(ctx, req.(*UserPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_LeaveParty_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).LeaveParty(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/LeaveParty",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).LeaveParty(ctx, req.(*UserPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetPartyParticipant_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserPartyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetPartyParticipant(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetPartyParticipant",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetPartyParticipant(ctx, req.(*UserPartyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetPartyParticipants_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyParticipantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetPartyParticipants(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetPartyParticipants",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetPartyParticipants(ctx, req.(*GetPartyParticipantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetPartyRequests_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartyParticipantsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetPartyRequests(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetPartyRequests",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetPartyRequests(ctx, req.(*GetPartyParticipantsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetUserParticipations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserParticipationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetUserParticipations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetUserParticipations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetUserParticipations(ctx, req.(*GetUserParticipationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParticipationService_GetManyUserParticipations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetManyUserParticipationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParticipationServiceServer).GetManyUserParticipations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/participation.ParticipationService/GetManyUserParticipations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParticipationServiceServer).GetManyUserParticipations(ctx, req.(*GetManyUserParticipationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ParticipationService_ServiceDesc is the grpc.ServiceDesc for ParticipationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParticipationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "participation.ParticipationService",
	HandlerType: (*ParticipationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "InviteToParty",
			Handler:    _ParticipationService_InviteToParty_Handler,
		},
		{
			MethodName: "DeclinePartyInvite",
			Handler:    _ParticipationService_DeclinePartyInvite_Handler,
		},
		{
			MethodName: "AcceptPartyInvite",
			Handler:    _ParticipationService_AcceptPartyInvite_Handler,
		},
		{
			MethodName: "GetUserInvites",
			Handler:    _ParticipationService_GetUserInvites_Handler,
		},
		{
			MethodName: "JoinParty",
			Handler:    _ParticipationService_JoinParty_Handler,
		},
		{
			MethodName: "LeaveParty",
			Handler:    _ParticipationService_LeaveParty_Handler,
		},
		{
			MethodName: "GetPartyParticipant",
			Handler:    _ParticipationService_GetPartyParticipant_Handler,
		},
		{
			MethodName: "GetPartyParticipants",
			Handler:    _ParticipationService_GetPartyParticipants_Handler,
		},
		{
			MethodName: "GetPartyRequests",
			Handler:    _ParticipationService_GetPartyRequests_Handler,
		},
		{
			MethodName: "GetUserParticipations",
			Handler:    _ParticipationService_GetUserParticipations_Handler,
		},
		{
			MethodName: "GetManyUserParticipations",
			Handler:    _ParticipationService_GetManyUserParticipations_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "participation/participation.proto",
}
