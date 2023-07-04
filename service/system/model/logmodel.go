package model

import (
	"gorm.io/gorm"
)

type (
	LogModel interface {
		AddLog(log *Log) (err error)
	}

	defaultLogModel struct {
		conn *gorm.DB
	}
	Log struct {
		gorm.Model
		Phone     string `json:"phone" gorm:"type:varchar(191);comment:手机号;not null"`      //手机号
		Operation string `json:"operation" gorm:"type:varchar(191);comment:用户操作;not null"` //用户操作
		Method    string `json:"method" gorm:"type:varchar(191);comment:请求方法;not null"`    //请求方法
		Params    string `json:"params" gorm:"type:varchar(191);comment:请求参数;not null"`    //请求参数
		Time      int64  `json:"time" gorm:"type:varchar(191);comment:执行时长(毫秒);not null"`  //执行时长(毫秒)
		IP        string `json:"ip" gorm:"type:varchar(191);comment:IP地址;not null"`        //IP地址
	}
)

func NewLogModel(conn *gorm.DB) LogModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Log{})
	return &defaultLogModel{
		conn: conn,
	}
}
func (m *defaultLogModel) AddLog(log *Log) (err error) {
	return m.conn.Model(&Log{}).Create(log).Error
}
