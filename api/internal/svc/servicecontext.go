package svc

import (
	"epick-mall/api/internal/config"
	"epick-mall/api/internal/middleware"
	"epick-mall/service/coupon/rpc/couponclient"
	"epick-mall/service/system/rpc/systemclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	System systemclient.System
	Coupon couponclient.Coupon
	Redis  *redis.Redis
	AddLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,
		System: systemclient.NewSystem(zrpc.MustNewClient(c.SystemRpc)),
		Coupon: couponclient.NewCoupon(zrpc.MustNewClient(c.CouponRpc)),
		Redis:  newRedis,
		AddLog: middleware.NewAddLogMiddleware(systemclient.NewSystem(zrpc.MustNewClient(c.SystemRpc))).Handle,
	}
}
func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
