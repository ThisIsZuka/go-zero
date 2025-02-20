// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"test.com/rpc/internal/logic/userservice"
	"test.com/rpc/internal/svc"
	"test.com/rpc/pb/user"
)

type UserServiceServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServiceServer
}

func NewUserServiceServer(svcCtx *svc.ServiceContext) *UserServiceServer {
	return &UserServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServiceServer) Ping(ctx context.Context, in *user.Request) (*user.Response, error) {
	l := userservicelogic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}

func (s *UserServiceServer) GetUserById(ctx context.Context, in *user.GetUserByIdRequest) (*user.GetUserByIdResponse, error) {
	l := userservicelogic.NewGetUserByIdLogic(ctx, s.svcCtx)
	return l.GetUserById(in)
}

func (s *UserServiceServer) GetAllUsers(ctx context.Context, in *user.GetAllUserRequest) (*user.GetAllUsersResponse, error) {
	l := userservicelogic.NewGetAllUsersLogic(ctx, s.svcCtx)
	return l.GetAllUsers(in)
}

func (s *UserServiceServer) AddNewUser(ctx context.Context, in *user.AddNewUserRequest) (*user.AddNewUserResponse, error) {
	l := userservicelogic.NewAddNewUserLogic(ctx, s.svcCtx)
	return l.AddNewUser(in)
}

func (s *UserServiceServer) DeleteUserById(ctx context.Context, in *user.DeleteUserByIdRequest) (*user.DeleteUserByIdResponse, error) {
	l := userservicelogic.NewDeleteUserByIdLogic(ctx, s.svcCtx)
	return l.DeleteUserById(in)
}
