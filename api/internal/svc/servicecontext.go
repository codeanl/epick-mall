package svc

import (
	"epick-mall/api/internal/config"
	"epick-mall/api/middleware"
	"epick-mall/service/oms/rpc/omsclient"
	"epick-mall/service/pms/rpc/pmsclient"
	"epick-mall/service/sms/rpc/smsclient"
	"epick-mall/service/sys/rpc/sysclient"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	Sys    sysclient.Sys
	Sms    smsclient.Sms
	Pms    pmsclient.Pms
	Oms    omsclient.Oms
	Redis  *redis.Redis
	AddLog rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	newRedis := redis.New(c.Redis.Address, redisConfig(c))
	return &ServiceContext{
		Config: c,
		Sys:    sysclient.NewSys(zrpc.MustNewClient(c.SysRpc)),
		Sms:    smsclient.NewSms(zrpc.MustNewClient(c.SmsRpc)),
		Pms:    pmsclient.NewPms(zrpc.MustNewClient(c.PmsRpc)),
		Oms:    omsclient.NewOms(zrpc.MustNewClient(c.OmsRpc)),
		AddLog: middleware.NewAddLogMiddleware(sysclient.NewSys(zrpc.MustNewClient(c.SysRpc))).Handle,
		Redis:  newRedis,
	}
}

func redisConfig(c config.Config) redis.Option {
	return func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	}
}
