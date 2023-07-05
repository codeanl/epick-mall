package model

import (
	"epick-mall/service/sys/rpc/sys"
	"fmt"
	"gorm.io/gorm"
)

type (
	LogModel interface {
		AddLog(log *Log) (err error)
		GetLogList(in *sys.SysLogListReq) ([]*Log, int64, error)
		DeleteLogByIds(ids []int64) error
	}

	defaultLogModel struct {
		conn *gorm.DB
	}
	Log struct {
		gorm.Model
		Username  string `json:"username" gorm:"type:varchar(191);comment:用户名;not null"`   //用户名
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
func (m *defaultLogModel) GetLogList(in *sys.SysLogListReq) ([]*Log, int64, error) {
	var list []*Log
	db := m.conn.Model(&Log{}).Order("created_at DESC")
	if in.Username != "" {
		db = db.Where("username LIKE ?", fmt.Sprintf("%%%s%%", in.Username))
	}
	if in.Method != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Method))
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
func (m *defaultLogModel) DeleteLogByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Log{}).Error
	return err
}
