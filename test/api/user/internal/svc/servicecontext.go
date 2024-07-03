package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"test.com/api/user/internal/config"
	"test.com/rpc/client/userservice"
)

type ServiceContext struct {
	Config     config.Config
	UserClient userservice.UserService
}

func NewServiceContext(c config.Config) *ServiceContext {
	userClient := userservice.NewUserService(zrpc.MustNewClient(c.RpcClientConf))
	return &ServiceContext{
		Config:     c,
		UserClient: userClient,
	}
}
