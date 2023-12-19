package svc

import (
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"pan-blitz-buy/rpc/user/internal/config"
	"pan-blitz-buy/rpc/user/model"
)

type ServiceContext struct {
	Config    config.Config
	DB        *gorm.DB
	CodeCache *cache.Cache
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		DB:        model.InitMySQL(&c),
		CodeCache: model.InitCache(&c),
	}
}
