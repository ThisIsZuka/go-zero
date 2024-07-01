package config

import "github.com/zeromicro/go-zero/zrpc"

type MongoConfig struct {
	URI      string
	Database string
}

type Config struct {
	zrpc.RpcServerConf
	Mongo MongoConfig
}
