package model

import (
	"epick-mall/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	RoleModel interface {
		AddRole(role *Role) (err error)
		UpdateRole(id int64, role *Role) error
		DeleteRoleByIds(ids []int64) error
		GetRoleList(in *sys.RoleListReq) ([]*Role, int64, error)
		GetRoleInfoByUserID(userid int64) (*Role, error)
	}

	defaultRoleModel struct {
		conn *gorm.DB
	}
	Role struct {
		gorm.Model
		Name     string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"`     //角色名称
		Remark   string `json:"remark" gorm:"type:varchar(191);comment:备注;not null"`     //备注
		CreateBy string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"` //创建人
		UpdateBy string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"` //更新人
	}
)

func NewRoleModel(conn *gorm.DB) RoleModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Role{})
	return &defaultRoleModel{
		conn: conn,
	}
}
func (m *defaultRoleModel) AddRole(role *Role) (err error) {
	return m.conn.Model(&Role{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultRoleModel) UpdateRole(id int64, role *Role) error {
	err := m.conn.Model(&Role{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultRoleModel) DeleteRoleByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Role{}).Error
	return err
}
func (m *defaultRoleModel) GetRoleInfoByUserID(userid int64) (*Role, error) {
	var userrole UserRole
	if err := m.conn.Model(&UserRole{}).Where("user_id=?", userid).Find(&userrole).Error; err != nil {
		return nil, err
	}
	var roleinfo Role
	err := m.conn.Model(&Role{}).Where("id=?", userrole.RoleID).Find(&roleinfo).Error
	if err != nil {
		return nil, err
	}
	return &roleinfo, nil
}

//GetUserList 获取用户列表
func (m *defaultRoleModel) GetRoleList(in *sys.RoleListReq) ([]*Role, int64, error) {
	var list []*Role
	db := m.conn.Model(&Role{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if in.PageNum > 0 && in.PageSize > 0 {
		err = db.Offset(int((in.PageNum - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
