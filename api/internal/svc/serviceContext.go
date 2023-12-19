package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"pan-blitz-buy/api/internal/config"
	"pan-blitz-buy/api/rpcclients/userclient/userservice"
)

type ServiceContext struct {
	Config      config.Config
	UserService userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		UserService: userservice.NewUserService(zrpc.MustNewClient(c.User)),
	}
}
