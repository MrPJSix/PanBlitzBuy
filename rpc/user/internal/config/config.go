package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	Mysql struct {
		Host     string
		User     string
		Password string
		Schema   string
	}
	Email struct {
		Host     string
		Port     int
		Username string
		Password string
	}
	CodeCache struct {
		ExpireTime int
		CleanTime  int
	}
}
