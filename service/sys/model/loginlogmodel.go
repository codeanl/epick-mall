package model

import (
	"epick-mall/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	LoginLogModel interface {
		AddLoginLog(loginLog *LoginLog) (err error)
		GetLoginLogList(in *sys.LoginLogListReq) ([]*LoginLog, int64, error)
		DeleteLoginLogByIds(ids []int64) error
	}

	defaultLoginLogModel struct {
		conn *gorm.DB
	}
	LoginLog struct {
		gorm.Model
		Username string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"` //用户名
		Status   string `json:"status" gorm:"type:varchar(191);comment:登录状态;not null"`  //登录状态（online:在线，登录初始状态，方便统计在线人数；login:退出登录后将online置为login；logout:退出登录）
		IP       string `json:"ip" gorm:"type:varchar(191);comment:IP地址;not null"`      //IP地址
	}
)

func NewLoginLogModel(conn *gorm.DB) LoginLogModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&LoginLog{})
	return &defaultLoginLogModel{
		conn: conn,
	}
}
func (m *defaultLoginLogModel) AddLoginLog(loginLog *LoginLog) (err error) {
	return m.conn.Model(&LoginLog{}).Create(loginLog).Error
}
func (m *defaultLoginLogModel) GetLoginLogList(in *sys.LoginLogListReq) ([]*LoginLog, int64, error) {
	var list []*LoginLog
	db := m.conn.Model(&LoginLog{}).Order("created_at DESC")
	if in.Username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Ip != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Ip))
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
func (m *defaultLoginLogModel) DeleteLoginLogByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&LoginLog{}).Error
	return err
}
