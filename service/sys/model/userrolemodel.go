package model

import (
	"gorm.io/gorm"
)

type (
	UserRoleModel interface {
		AddUserRole(userrole *UserRole) (err error)
		UpdateUserRole(userrole *UserRole) error
	}
	defaultUserRoleModel struct {
		conn *gorm.DB
	}
	UserRole struct {
		gorm.Model
		UserID     int    `json:"user_id" gorm:"type:int;comment:用户ID;not null"`                //用户id
		RoleID     int    `json:"role_id" gorm:"type:int;comment:角色ID;not null"`                //角色id
		LastEditBy string `json:"last_edit_by" gorm:"type:varchar(191);comment:谁最后修改;not null"` //谁最后修改
	}
)

func NewUserRoleModel(conn *gorm.DB) UserRoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&UserRole{})
	return &defaultUserRoleModel{
		conn: conn,
	}
}

//AddUserRole 添加用户角色
func (m *defaultUserRoleModel) AddUserRole(userrole *UserRole) (err error) {
	return m.conn.Model(&UserRole{}).Create(userrole).Error
}

//UpdateUserRole 更新用户角色
func (m *defaultUserRoleModel) UpdateUserRole(userrole *UserRole) error {
	err := m.conn.Model(&UserRole{}).Where("user_id = ?", userrole.UserID).Updates(userrole).Error
	return err
}
