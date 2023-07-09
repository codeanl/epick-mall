package model

import (
	"epick-mall/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	AttributeCategoryModel interface {
		AddAttributeCategory(role *AttributeCategory) (err error)
		UpdateAttributeCategory(id int64, role *AttributeCategory) error
		DeleteAttributeCategoryByIds(ids []int64) error
		GetAttributeCategoryList(in *pms.ProductAttributeCategoryListReq) ([]*AttributeCategory, int64, error)
	}

	defaultAttributeCategoryModel struct {
		conn *gorm.DB
	}
	AttributeCategory struct {
		gorm.Model
		Name           string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`                  //名称
		AttributeCount int64  `json:"attribute_count" gorm:"type:bigint;default:0;comment:属性数量;not null"` //属性数量
		ParamCount     int64  `json:"param_count" gorm:"type:bigint;comment:参数数量;default:0;not null"`     //参数数量
	}
)

func NewAttributeCategoryModel(conn *gorm.DB) AttributeCategoryModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&AttributeCategory{})
	return &defaultAttributeCategoryModel{
		conn: conn,
	}
}
func (m *defaultAttributeCategoryModel) AddAttributeCategory(role *AttributeCategory) (err error) {
	return m.conn.Model(&AttributeCategory{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultAttributeCategoryModel) UpdateAttributeCategory(id int64, role *AttributeCategory) error {
	err := m.conn.Model(&AttributeCategory{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultAttributeCategoryModel) DeleteAttributeCategoryByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&AttributeCategory{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultAttributeCategoryModel) GetAttributeCategoryList(in *pms.ProductAttributeCategoryListReq) ([]*AttributeCategory, int64, error) {
	var list []*AttributeCategory
	db := m.conn.Model(&AttributeCategory{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return list, total, err
	}
	if in.Current > 0 && in.PageSize > 0 {
		err = db.Offset(int((in.Current - 1) * in.PageSize)).Limit(int(in.PageSize)).Find(&list).Error
	} else {
		err = db.Find(&list).Error
	}
	return list, total, err
}
