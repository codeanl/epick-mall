package model

import (
	"epick-mall/service/sms/rpc/sms"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type (
	HomeAdvertiseModel interface {
		AddHomeAdvertise(coupon *HomeAdvertise) (err error)
		UpdateHomeAdvertise(id int64, coupon *HomeAdvertise) error
		DeleteHomeAdvertiseByIds(ids []int64) error
		GetHomeAdvertiseList(in *sms.HomeAdvertiseListReq) ([]*HomeAdvertise, int64, error)
	}
	defaultHomeAdvertiseModel struct {
		conn *gorm.DB
	}
	HomeAdvertise struct {
		gorm.Model
		Type       string    `json:"type" gorm:"type:varchar(191);comment:轮播位置;not null"` //轮播位置：0->PC首页轮播；1->app首页轮播
		Name       string    `json:"name" gorm:"type:varchar(191);comment:名称;not null"`   //名称
		Pic        string    `json:"pic" gorm:"type:varchar(191);comment:图片地址;not null"`  //图片地址
		StartTime  time.Time `json:"start_time" gorm:"type:datetime;comment:开始时间"`        //开始时间
		EndTime    time.Time `json:"end_time" gorm:"type:datetime;comment:结束时间"`          //结束时间
		Status     string    `json:"status" gorm:"type:varchar(191);comment:上下线状态"`       //上下线状态：0->下线；1->上线',
		ClickCount int64     `json:"click_count" gorm:"type:bigint;comment:点击数"`          //点击数
		OrderCount int64     `json:"order_count" gorm:"type:bigint;comment:下单数"`          //下单数
		Url        string    `json:"url" gorm:"type:varchar(191);comment:链接地址"`           //链接地址
		Note       string    `json:"note" gorm:"type:varchar(191);comment:备注"`            //备注
		Sort       int64     `json:"sort" gorm:"type:bigint;comment:排序"`                  //排序
	}
)

func NewHomeAdvertiseModel(conn *gorm.DB) HomeAdvertiseModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&HomeAdvertise{})
	return &defaultHomeAdvertiseModel{
		conn: conn,
	}
}

func (m *defaultHomeAdvertiseModel) AddHomeAdvertise(coupon *HomeAdvertise) (err error) {
	return m.conn.Model(&HomeAdvertise{}).Create(coupon).Error
}

func (m *defaultHomeAdvertiseModel) UpdateHomeAdvertise(id int64, coupon *HomeAdvertise) error {
	err := m.conn.Model(&HomeAdvertise{}).Where("id=?", id).Updates(coupon).Error
	return err
}

func (m *defaultHomeAdvertiseModel) DeleteHomeAdvertiseByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&HomeAdvertise{}).Error
	return err
}

func (m *defaultHomeAdvertiseModel) GetHomeAdvertiseList(in *sms.HomeAdvertiseListReq) ([]*HomeAdvertise, int64, error) {
	var list []*HomeAdvertise
	db := m.conn.Model(&list).Order("created_at DESC")
	if in.Type != "" {
		db = db.Where("type LIKE ?", fmt.Sprintf("%%%s%%", in.Type))
	}
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.StartTime != "" {
		db = db.Where("start_time = ?", fmt.Sprintf("%%%s%%", in.StartTime))
	}
	if in.EndTime != "" {
		db = db.Where("end_time = ?", fmt.Sprintf("%%%s%%", in.EndTime))
	}
	if in.Status != "" {
		db = db.Where("status = ?", fmt.Sprintf("%%%s%%", in.Status))
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
