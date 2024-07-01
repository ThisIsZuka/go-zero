package svc

import (
	"test.com/model"
	"test.com/rpc/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(c.Mongo.URI, c.Mongo.Database, "users"),
	}
}
