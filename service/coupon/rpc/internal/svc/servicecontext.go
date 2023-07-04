package svc

import (
	"epick-mall/common"
	"epick-mall/service/coupon/model"
	"epick-mall/service/coupon/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	CouponModel model.CouponModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := common.InitGorm(mysql.DataSource)
	return &ServiceContext{
		Config:      c,
		CouponModel: model.NewCouponModel(conn),
	}
}
