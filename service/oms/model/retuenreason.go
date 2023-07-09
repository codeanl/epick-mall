package model

import (
	"epick-mall/service/oms/rpc/oms"
	"fmt"
	"gorm.io/gorm"
)

type (
	ReturnReasonModel interface {
		AddReturnReason(info *ReturnReason) (err error)
		UpdateReturnReason(id int64, info *ReturnReason) error
		DeleteReturnReasonByIds(ids []int64) error
		GetReturnReasonList(in *oms.OrderReturnReasonListReq) ([]*ReturnReason, int64, error)
	}

	defaultReturnReasonModel struct {
		conn *gorm.DB
	}
	ReturnReason struct {
		gorm.Model
		Name string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"` //角色名称
		Sort int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`         //排序
	}
)

func NewReturnReasonModel(conn *gorm.DB) ReturnReasonModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&ReturnReason{})
	return &defaultReturnReasonModel{
		conn: conn,
	}
}
func (m *defaultReturnReasonModel) AddReturnReason(info *ReturnReason) (err error) {
	return m.conn.Model(&ReturnReason{}).Create(info).Error
}

//UpdateReturnReason 修改角色
func (m *defaultReturnReasonModel) UpdateReturnReason(id int64, info *ReturnReason) error {
	err := m.conn.Model(&ReturnReason{}).Where("id=?", id).Updates(info).Error
	return err
}

//DeleteReturnReasonByIds 删除角色
func (m *defaultReturnReasonModel) DeleteReturnReasonByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&ReturnReason{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultReturnReasonModel) GetReturnReasonList(in *oms.OrderReturnReasonListReq) ([]*ReturnReason, int64, error) {
	var list []*ReturnReason
	db := m.conn.Model(&ReturnReason{}).Order("created_at DESC")
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
