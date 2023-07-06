package model

import (
	"epick-mall/service/sms/rpc/sms"
	"fmt"
	"gorm.io/gorm"
)

type (
	CouponModel interface {
		////GetUserByPhone(phone string) (user *User, exist bool, err error)
		AddCoupon(coupon *Coupon) (err error)
		UpdateCoupon(id int64, coupon *Coupon) error
		DeleteCouponByIds(ids []int64) error
		GetCouponList(in *sms.CouponListReq) ([]*Coupon, int64, error)
		////GetUserByID(id int64) (user *User, err error)
	}
	defaultCouponModel struct {
		conn *gorm.DB
	}
	Coupon struct {
		gorm.Model
		Type         string  `json:"type" gorm:"type:varchar(191);comment:优惠券类型;not null"`    //优惠券类型；0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券'
		Name         string  `json:"name" gorm:"type:varchar(191);comment:名称;not null"`       //名称
		Platform     string  `json:"platform" gorm:"type:varchar(191);comment:使用平台;not null"` //使用平台 '使用平台：0->全部；1->移动；2->PC',
		Count        int64   `json:"count" gorm:"type:bigint;comment:数量;not null"`            //数量
		Amount       float64 `json:"amount" gorm:"type:decimal(10,2);comment:金额;not null"`    //金额
		PerLimit     int64   `json:"per_limit" gorm:"type:bigint;comment:每人限领张数"`             //每人限领张数
		MinPoint     float64 `json:"min_point" gorm:"type:decimal(10,2);comment:使用w门槛"`       //使用门槛；0表示无门槛
		StartTime    string  `json:"start_time" gorm:"type:datetime;comment:开始时间"`            //开始时间
		EndTime      string  `json:"end_time" gorm:"type:datetime;comment:结束时间"`              //结束时间
		UseType      string  `json:"use_type" gorm:"type:varchar(191);comment:使用类型"`          //使用类型：0->全场通用；1->指定分类；2->指定商品
		Note         string  `json:"note" gorm:"type:varchar(191);comment:备注"`                //备注
		UseCount     int64   `json:"use_count" gorm:"type:bigint;comment:已使用数量"`              //已使用数量
		ReceiveCount int64   `json:"receive_count" gorm:"type:bigint;comment:领取数量"`           //领取数量
		EnableTime   string  `json:"enable_time" gorm:"type:datetime;comment:可以领取的日期"`        //可以领取的日期
		Code         string  `json:"code" gorm:"type:varchar(191);comment:优惠码"`               //优惠码
	}
)

func NewCouponModel(conn *gorm.DB) CouponModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Coupon{})
	return &defaultCouponModel{
		conn: conn,
	}
}

//func (m *defaultUserModel) GetUserByID(id int64) (user *User, err error) {
//	err = m.conn.Model(&User{}).Where("id=?", id).First(&user).Error
//	return user, err
//}
//
////GetUserByPhone 查询用户是或存在
//func (m *defaultUserModel) GetUserByPhone(phone string) (user *User, exist bool, err error) {
//	var count int64
//	err = m.conn.Model(&User{}).Where("phone=?", phone).Count(&count).Error
//	if count == 0 {
//		return nil, false, err
//	}
//	err = m.conn.Model(&User{}).Where("phone=?", phone).First(&user).Error
//	if err != nil {
//		return nil, false, err
//	}
//	return user, true, nil
//}
//
func (m *defaultCouponModel) AddCoupon(coupon *Coupon) (err error) {
	return m.conn.Model(&Coupon{}).Create(coupon).Error
}

//
func (m *defaultCouponModel) UpdateCoupon(id int64, coupon *Coupon) error {
	err := m.conn.Model(&Coupon{}).Where("id=?", id).Updates(coupon).Error
	return err
}

func (m *defaultCouponModel) DeleteCouponByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Coupon{}).Error
	return err
}

//
////GetUserList 获取用户列表
func (m *defaultCouponModel) GetCouponList(in *sms.CouponListReq) ([]*Coupon, int64, error) {
	var list []*Coupon
	db := m.conn.Model(&list).Order("created_at DESC")
	if in.Type != "" {
		db = db.Where("type LIKE ?", fmt.Sprintf("%%%s%%", in.Type))
	}
	if in.Name != "" {
		db = db.Where("name LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
	}
	if in.Platform != "" {
		db = db.Where("platform LIKE ?", fmt.Sprintf("%%%s%%", in.Platform))
	}
	if in.StartTime != "" {
		db = db.Where("start_time = ?", fmt.Sprintf("%%%s%%", in.StartTime))
	}
	if in.EndTime != "" {
		db = db.Where("end_time = ?", fmt.Sprintf("%%%s%%", in.EndTime))
	}
	if in.UseType != "" {
		db = db.Where("use_type = ?", fmt.Sprintf("%%%s%%", in.UseType))
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
