package model

import (
	"epick-mall/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	MenuModel interface {
		AddMenu(menu *Menu) (err error)
		GetMenuList(in *sys.MenuListReq) ([]*Menu, int64, error)
		UpdateMenu(id int64, menu *Menu) error
		DeleteMenuByIds(ids []int64) error
		GetMenusByUserID(userID int64) ([]sys.MenuListData, error)
	}

	defaultMenuModel struct {
		conn *gorm.DB
	}
	Menu struct {
		gorm.Model
		Name          string `json:"name" gorm:"type:varchar(191);comment:菜单名称;not null"`              //菜单名称
		ParentID      int    `json:"parent_id" gorm:"type:int;comment:父菜单ID;not null;default:0"`       //父菜单ID
		Url           string `json:"url" gorm:"type:varchar(191);comment:路径;not null"`                 //路径
		Perms         string `json:"perms" gorm:"type:varchar(191);comment:授权;not null"`               //授权(多个用逗号分隔，如：sys:user:add,sys:user:edit)
		Type          int    `json:"type" gorm:"type:int;comment:类型;not null"`                         //类型  0：目录   1：菜单   2：按钮',
		Icon          string `json:"icon" gorm:"type:varchar(191);comment:菜单图标;not null"`              //菜单图标
		OrderNum      int    `json:"order_num" gorm:"type:int;comment:排序;not null"`                    //排序
		CreateBy      string `json:"create_by" gorm:"type:varchar(191);comment:创建人;not null"`          //创建人
		UpdateBy      string `json:"update_by" gorm:"type:varchar(191);comment:更新人;not null"`          //更新人
		BackgroundUrl string `json:"background_url" gorm:"type:varchar(191);comment:接口地址;not null"`    //接口地址
		VuePath       string `json:"vue_path" gorm:"type:varchar(191);comment:vue系统的path;not null"`    //vue系统的path
		VueComponent  string `json:"vue_component" gorm:"type:varchar(191);comment:vue的页面;not null"`   //vue的页面
		VueIcon       string `json:"vue_icon" gorm:"type:varchar(191);comment:vue的图标;not null"`        //vue的图标
		VueRedirect   string `json:"vue_redirect" gorm:"type:varchar(191);comment:vue的路由重定向;not null"` //vue的路由重定向
	}
)

func NewMenuModel(conn *gorm.DB) MenuModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Menu{})
	return &defaultMenuModel{
		conn: conn,
	}
}
func (m *defaultMenuModel) AddMenu(menu *Menu) (err error) {
	return m.conn.Model(&Menu{}).Create(menu).Error
}
func (m *defaultMenuModel) GetMenuList(in *sys.MenuListReq) ([]*Menu, int64, error) {
	var list []*Menu
	db := m.conn.Model(&Menu{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Url != "" {
		db = db.Where("url LIKE ?", fmt.Sprintf("%%%s%%", in.Url))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	err = db.Find(&list).Error
	return list, total, err
}
func (m *defaultMenuModel) UpdateMenu(id int64, menu *Menu) error {
	err := m.conn.Model(&Menu{}).Where("id=?", id).Updates(menu).Error
	return err
}
func (m *defaultMenuModel) DeleteMenuByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Menu{}).Error
	return err
}
func (m *defaultMenuModel) GetMenusByUserID(userID int64) ([]sys.MenuListData, error) {
	var userrole UserRole
	if err := m.conn.Model(&UserRole{}).Where("user_id=?", userID).First(&userrole).Error; err != nil {
		return nil, err
	}
	var role Role
	if err := m.conn.Model(&Role{}).First(&role, userrole.RoleID).Error; err != nil {
		return nil, err
	}
	var roleMenus []RoleMenu
	if err := m.conn.Model(&RoleMenu{}).Where("role_id = ?", role.ID).Find(&roleMenus).Error; err != nil {
		return nil, err
	}
	var menuIDs []uint
	for _, rm := range roleMenus {
		menuIDs = append(menuIDs, uint(rm.MenuID))
	}
	var menus []sys.MenuListData
	if err := m.conn.Model(&Menu{}).Where("id IN (?)", menuIDs).Find(&menus).Error; err != nil {
		return nil, err
	}
	return menus, nil
}
