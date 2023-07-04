package svc

import (
	"epick-mall/common"
	"epick-mall/service/product/model"
	"epick-mall/service/product/rpc/internal/config"
)

type ServiceContext struct {
	Config       config.Config
	ProductModel model.ProductModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := common.InitGorm(mysql.DataSource)

	return &ServiceContext{
		Config:       c,
		ProductModel: model.NewProductModel(conn),
	}
}
