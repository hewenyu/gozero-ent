package svc

import (
	"github.com/hewenyu/ent/ent"
	"github.com/hewenyu/gozero-ent/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config config.Config
	DB     *ent.Client
}

func NewServiceContext(c *config.Config) *ServiceContext {

	db := ent.NewClient(
		ent.Log(logx.Info),               // logger
		ent.Driver(GetCacheTTLDriver(c)), // 缓存 这个走lru
		ent.Debug(),                      // debug mode
	)

	logx.Info("Initialize database connection successfully")

	return &ServiceContext{
		Config: *c,
		DB:     db,
	}
}
