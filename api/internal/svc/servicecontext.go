package svc

import (
	"epick-mall/api/internal/config"
	"epick-mall/api/internal/middleware"
	"epick-mall/service/system/rpc/systemclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	System systemclient.System
	Redis  *redis.Redis
	AddLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	newSys := systemclient.NewSystem(zrpc.MustNewClient(c.SystemRpc))
	return &ServiceContext{
		Config: c,
		System: newSys,
		Redis:  newRedis,
		AddLog: middleware.NewAddLogMiddleware(newSys).Handle,
	}
}
func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
