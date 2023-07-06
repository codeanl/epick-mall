package svc

import (
	"epick-mall/common"
	"epick-mall/service/sms/model"
	"epick-mall/service/sms/rpc/internal/config"
)

type ServiceContext struct {
	Config             config.Config
	CouponModel        model.CouponModel
	HomeAdvertiseModel model.HomeAdvertiseModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := common.InitGorm(mysql.DataSource)
	return &ServiceContext{
		Config:             c,
		CouponModel:        model.NewCouponModel(conn),
		HomeAdvertiseModel: model.NewHomeAdvertiseModel(conn),
	}
}
