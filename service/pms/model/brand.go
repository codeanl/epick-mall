package model

import (
	"epick-mall/service/pms/rpc/pms"
	"fmt"
	"gorm.io/gorm"
)

type (
	BrandModel interface {
		AddBrand(role *Brand) (err error)
		UpdateBrand(id int64, role *Brand) error
		DeleteBrandByIds(ids []int64) error
		GetBrandList(in *pms.BrandListReq) ([]*Brand, int64, error)
		GetBrandByIdsList(in *pms.BrandListByIdsReq) ([]*Brand, int64, error)
	}

	defaultBrandModel struct {
		conn *gorm.DB
	}
	Brand struct {
		gorm.Model
		Name                string `json:"name" gorm:"type:varchar(191);comment:角色名称;not null"`               //角色名称
		FirstLetter         string `json:"first_letter" gorm:"type:varchar(191);comment:首字母;not null"`        //首字母
		Sort                int64  `json:"sort" gorm:"type:bigint;comment:排序;not null"`                       //排序
		FactoryStatus       string `json:"factory_status" gorm:"type:varchar(191);comment:是否为品牌制造商;not null"` //是否为品牌制造商：0->不是；1->是'
		ShowStatus          string `json:"show_status" gorm:"type:varchar(191);comment:展示状态;not null"`        //展示状态0->不是；1->是'
		ProductCount        int64  `json:"product_count" gorm:"type:bigint;comment:产品数量;not null"`            //产品数量
		ProductCommentCount int64  `json:"product_comment_count" gorm:"type:bigint;comment:产品评论数量;not null"`  //产品评论数量
		Logo                string `json:"logo" gorm:"type:varchar(191);comment:logo;not null"`               //logo
		BigPic              string `json:"big_pic" gorm:"type:varchar(191);comment:专区大图;not null"`            //专区大图
		BrandStory          string `json:"brand_story" gorm:"type:varchar(191);comment:品牌故事;not null"`        //品牌故事
	}
)

func NewBrandModel(conn *gorm.DB) BrandModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Brand{})
	return &defaultBrandModel{
		conn: conn,
	}
}
func (m *defaultBrandModel) AddBrand(role *Brand) (err error) {
	return m.conn.Model(&Brand{}).Create(role).Error
}

//UpdateRole 修改角色
func (m *defaultBrandModel) UpdateBrand(id int64, role *Brand) error {
	err := m.conn.Model(&Brand{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultBrandModel) DeleteBrandByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Brand{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultBrandModel) GetBrandList(in *pms.BrandListReq) ([]*Brand, int64, error) {
	var list []*Brand
	db := m.conn.Model(&Brand{}).Order("created_at DESC")
	if in.Name != "" {
		db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", in.Name))
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
func (m *defaultBrandModel) GetBrandByIdsList(in *pms.BrandListByIdsReq) ([]*Brand, int64, error) {
	var list []*Brand
	err := m.conn.Model(&Brand{}).Where("id in (in)", in).Find(&list).Order("created_at DESC").Error
	var total int64
	err = m.conn.Model(&Brand{}).Count(&total).Error
	return list, total, err
}
