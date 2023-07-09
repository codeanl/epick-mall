package model

import (
	"epick-mall/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	AttributeModel interface {
		AddAttribute(role *Attribute) (err error)
		UpdateAttribute(id int64, role *Attribute) error
		DeleteAttributeByIds(ids []int64) error
		GetAttributeList(in *pms.ProductAttributeListReq) ([]*Attribute, int64, error)
	}

	defaultAttributeModel struct {
		conn *gorm.DB
	}
	Attribute struct {
		gorm.Model
		Name                       string `json:"name" gorm:"type:varchar(191);comment:名称;not null"`                        //名称
		ProductAttributeCategoryId int64  `json:"product_attribute_category_id" gorm:"type:bigint;comment:属性分类id;not null"` //属性分类id
		Sort                       int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`                              //排序
		SelectType                 string `json:"select_type" gorm:"type:varchar(191);comment:属性选择类型;not null"`             //属性选择类型：0->唯一；1->单选；2->多选',：0->不是；1->是'
		InputType                  string `json:"input_type" gorm:"type:varchar(191);comment:属性录入方式;not null"`              //属性录入方式：0->手工录入；1->从列表中选取',
		InputList                  string `json:"input_list" gorm:"type:varchar(191);comment:可选值列表;not null"`               //可选值列表，以逗号隔开
		FilterType                 string `json:"filter_type" gorm:"type:varchar(191);comment:分类筛选样式;not null"`             //'分类筛选样式：1->普通；1->颜色',
		SearchType                 string `json:"search_type" gorm:"type:varchar(191);comment:检索类型;not null"`               //检索类型；0->不需要进行检索；1->关键字检索；2->范围检索',
		RelatedStatus              string `json:"related_status" gorm:"type:varchar(191);comment:相同属性产品是否关联;not null"`      //相同属性产品是否关联；0->不关联；1->关联',
		HandAddStatus              string `json:"hand_add_status" gorm:"type:varchar(191);comment:是否支持手动新增;not null"`       //相同属性产品是否关联；0->不关联；1->关联',
		Type                       string `json:"type" gorm:"type:varchar(191);comment:属性的类型;not null"`                     //属性的类型；0->规格；1->参数'
	}
)

func NewAttributeModel(conn *gorm.DB) AttributeModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Attribute{})
	return &defaultAttributeModel{
		conn: conn,
	}
}
func (m *defaultAttributeModel) AddAttribute(role *Attribute) (err error) {
	return m.conn.Model(&Attribute{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultAttributeModel) UpdateAttribute(id int64, role *Attribute) error {
	err := m.conn.Model(&Attribute{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultAttributeModel) DeleteAttributeByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Attribute{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultAttributeModel) GetAttributeList(in *pms.ProductAttributeListReq) ([]*Attribute, int64, error) {
	var list []*Attribute
	db := m.conn.Model(&Attribute{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Type != "" {
		db = db.Where("type LIKE ?", fmt.Sprintf("%%%s%%", in.Type))
	}
	if in.ProductAttributeCategoryId != 0 {
		db = db.Where("product_attribute_category_id = ?", fmt.Sprintf("%%%s%%", in.ProductAttributeCategoryId))
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
