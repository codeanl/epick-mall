package svc

import (
	"epick-mall/common"
	"epick-mall/service/system/model"
	"epick-mall/service/system/rpc/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     model.UserModel
	RoleModel     model.RoleModel
	UserRoleModel model.UserRoleModel
	MenuModel     model.MenuModel
	RoleMenuModel model.RoleMenuModel
	LoginLogModel model.LoginLogModel
	LogModel      model.LogModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	mysql := c.Mysql
	conn := common.InitGorm(mysql.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     model.NewUserModel(conn),
		RoleModel:     model.NewRoleModel(conn),
		UserRoleModel: model.NewUserRoleModel(conn),
		MenuModel:     model.NewMenuModel(conn),
		RoleMenuModel: model.NewRoleMenuModel(conn),
		LoginLogModel: model.NewLoginLogModel(conn),
		LogModel:      model.NewLogModel(conn),
	}
}
