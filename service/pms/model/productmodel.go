package model

import (
	"epick-mall/service/pms/rpc/pms"
	"gorm.io/gorm"
)

type (
	ProductModel interface {
		AddProduct(product *Product) (err error)
		UpdateProduct(id int64, role *Product) error
		DeleteProductByIds(ids []int64) error
		GetProductList(in *pms.ProductListReq) ([]*Product, int64, error)
	}

	defaultProductModel struct {
		conn *gorm.DB
	}
	Product struct {
		gorm.Model
		BrandID                    int64   `json:"brand_id" gorm:"type:bigint;comment:品牌id;not null"`                          //品牌id
		ProductCategoryID          int64   `json:"product_category_id" gorm:"type:bigint;comment:商品分类id;not null"`             //商品分类id
		FreightTemplateID          int64   `json:"freight_template_id" gorm:"type:bigint;comment:货物模版id;not null"`             //货物模版id
		ProductAttributeCategoryID int64   `json:"product_attribute_category_id" gorm:"type:bigint;comment:产品属性类别id;not null"` //产品属性类别id
		Name                       string  `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`                        //商品名称
		Pic                        string  `json:"pic" gorm:"type:varchar(191);comment:图片;not null"`                           //图片
		ProductSn                  string  `json:"product_sn" gorm:"type:varchar(191);comment:货号;not null"`                    //货号
		DeleteStatus               string  `json:"delete_status" gorm:"type:varchar(191);comment:删除状态;not null"`               //删除状态：0->未删除；1->已删除'
		PublishStatus              string  `json:"publish_status" gorm:"type:varchar(191);comment:上架状态;not null"`              //上架状态：0->下架；1->上架',
		NewStatus                  string  `json:"new_status" gorm:"type:varchar(191);comment:新品状态;not null"`                  //新品状态:0->不是新品；1->新品',
		RecommendStatus            string  `json:"recommend_status" gorm:"type:varchar(191);comment:推荐状态;not null"`            //推荐状态；0->不推荐；1->推荐',
		Sort                       int64   `json:"sort" gorm:"type:bigint;comment:排序;not null"`                                //排序
		Sale                       int64   `json:"sale" gorm:"type:bigint;comment:销量;not null"`                                //销量
		Price                      float64 `json:"price" gorm:"type: decimal(10, 2);comment:价格;not null"`                      //价格
		PromotionPrice             float64 `json:"promotion_price" gorm:"type: decimal(10, 2);comment:促销价格;not null"`          //促销价格
		GiftGrowth                 int64   `json:"gift_growth" gorm:"type:bigint;comment:赠送的成长值;not null;default:0"`           //赠送的成长值
		GiftPoint                  int64   `json:"gift_point" gorm:"type:bigint;comment:赠送的积分;not null;default:0"`             //赠送的积分
		UsePointLimit              int64   `json:"use_point_limit" gorm:"type:bigint;comment:限制使用的积分数;not null;"`              //限制使用的积分数
		SubTitle                   string  `json:"sub_title" gorm:"type:varchar(191);comment:副标题;not null"`                    //副标题
		Description                string  `json:"description" gorm:"type:varchar(191);comment:商品描述;not null"`                 //商品描述
		OriginalPrice              float64 `json:"original_price" gorm:"type:decimal(10, 2);comment:市场价;not null"`             //市场价
		Stock                      int64   `json:"stock" gorm:"type:bigint;comment:库存;not null;default:0"`                     //库存
		LowStock                   int64   `json:"low_stock" gorm:"type:bigint;comment:库存预警值;not null;default:0"`              //库存预警值
		Unit                       string  `json:"unit" gorm:"type:varchar(191);comment:单位;not null"`                          //单位
		Weight                     float64 `json:"weight" gorm:"type:decimal(10, 2);comment:商品重量;not null"`                    //商品重量，默认为克
		PreviewStatus              string  `json:"preview_status" gorm:"type:varchar(191);comment:是否为预告商品;not null;default:0"` //是否为预告商品：0->不是；1->是
		ServiceIds                 string  `json:"service_ids" gorm:"type:varchar(191);comment:以逗号分割的产品服务;not null"`           //以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮',
		Keywords                   string  `json:"keywords" gorm:"type:varchar(191);comment:关键字;not null"`                     //关键字
		Note                       string  `json:"note" gorm:"type:varchar(191);comment:注释;not null"`                          //注释
		AlbumPics                  string  `json:"album_pics" gorm:"type:varchar(191);comment:画册图片;not null"`                  //画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle                string  `json:"detail_title" gorm:"type:varchar(191);comment:详情标题;not null"`                //详情标题
		DetailDesc                 string  `json:"detail_desc" gorm:"type:text;comment:详细说明;not null"`                         //详细说明
		DetailHtml                 string  `json:"detail_html" gorm:"type:text;comment:产品详情网页内容;not null"`                     //产品详情网页内容
		DetailMobileHtml           string  `json:"detail_mobile_html" gorm:"type:text;comment:移动端网页详情;not null"`               //移动端网页详情
		PromotionStartTime         string  `json:"promotion_start_time" gorm:"type:datetime;comment:促销开始时间;not null"`          //促销开始时间
		PromotionEndTime           string  `json:"promotion_end_time" gorm:"type:datetime;comment:促销结束时间;not null"`            //促销结束时间
		PromotionPerLimit          int64   `json:"promotion_per_limit" gorm:"type:bigint;comment:活动限购数量;not null;default:0"`   //活动限购数量
		PromotionType              string  `json:"promotion_type" gorm:"type:varchar(191);comment:促销类型;not null;default:0"`    //促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName                  string  `json:"brand_name" gorm:"type:text;comment:品牌名称;not null"`                          //品牌名称
		ProductCategoryName        string  `json:"product_category_name" gorm:"type:text;comment:商品分类名称;not null"`             //商品分类名称
	}
)

func NewProductModel(conn *gorm.DB) ProductModel {
	//如果没有表则自动构建表
	conn.AutoMigrate(&Product{})
	return &defaultProductModel{
		conn: conn,
	}
}
func (m *defaultProductModel) AddProduct(product *Product) (err error) {
	return m.conn.Model(&Product{}).Create(product).Error
}
func (m *defaultProductModel) UpdateProduct(id int64, role *Product) error {
	err := m.conn.Model(&Product{}).Where("id=?", id).Updates(role).Error
	return err
}

//DeleteRoleByIds 删除角色
func (m *defaultProductModel) DeleteProductByIds(ids []int64) error {
	id := map[string]interface{}{
		"id": ids,
	}
	err := m.conn.Where(id).Delete(&Product{}).Error
	return err
}

//GetUserList 获取用户列表
func (m *defaultProductModel) GetProductList(in *pms.ProductListReq) ([]*Product, int64, error) {
	var list []*Product
	db := m.conn.Model(&Product{}).Order("created_at DESC")

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
