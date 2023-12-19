// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"pan-blitz-buy/api/rpcclients/userclient/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BasicUserInfo   = user.BasicUserInfo
	GeneralReq      = user.GeneralReq
	GeneralResp     = user.GeneralResp
	GetUserInfoReq  = user.GetUserInfoReq
	GetUserInfoResp = user.GetUserInfoResp
	RegisterUserReq = user.RegisterUserReq
	UserInfo        = user.UserInfo

	UserService interface {
		Register(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*GeneralResp, error)
		Login(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error)
		AdminLogin(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error)
		SendEmail(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error)
		GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) Register(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUserService) Login(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUserService) AdminLogin(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.AdminLogin(ctx, in, opts...)
}

func (m *defaultUserService) SendEmail(ctx context.Context, in *GeneralReq, opts ...grpc.CallOption) (*GeneralResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.SendEmail(ctx, in, opts...)
}

func (m *defaultUserService) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	client := user.NewUserServiceClient(m.cli.Conn())
	return client.GetUserInfo(ctx, in, opts...)
}