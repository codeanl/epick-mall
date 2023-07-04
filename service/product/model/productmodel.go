package model

import (
	"gorm.io/gorm"
	"time"
)

type (
	ProductModel interface {
		AddProduct(product *Product) (err error)
	}

	defaultProductModel struct {
		conn *gorm.DB
	}
	Product struct {
		gorm.Model
		BrandID                    int64     `json:"brand_id" gorm:"type:int64;comment:品牌id;not null"`                        //品牌id
		ProductCategoryID          int64     `json:"product_category_id" gorm:"type:int64;comment:商品分类id;not null"`           //商品分类id
		FreightTemplateID          int64     `json:"freight_template_id" gorm:"type:int64;comment:货物模版id;not null"`           //货物模版id
		ProductAttributeCategoryID int64     `json:"product_attribute_category_id" gorm:"type:int;comment:产品属性类别id;not null"` //产品属性类别id
		Name                       string    `json:"name" gorm:"type:varchar(191);comment:商品名称;not null"`                     //商品名称
		Pic                        string    `json:"pic" gorm:"type:varchar(191);comment:图片;not null"`                        //图片
		ProductSn                  string    `json:"product_sn" gorm:"type:varchar(191);comment:货号;not null"`                 //货号
		DeleteStatus               string    `json:"delete_status" gorm:"type:int;comment:删除状态;not null"`                     //删除状态：0->未删除；1->已删除'
		NewStatus                  string    `json:"new_status" gorm:"type:int;comment:新品状态;not null"`                        //新品状态:0->不是新品；1->新品',
		RecommendStatus            string    `json:"recommend_status" gorm:"type:int;comment:推荐状态;not null"`                  //推荐状态；0->不推荐；1->推荐',
		Sort                       string    `json:"sort" gorm:"type:int;comment:排序;not null"`                                //排序
		Sale                       string    `json:"sale" gorm:"type:int;comment:销量;not null"`                                //销量
		Price                      float64   `json:"price" gorm:"type:float64;comment:价格;not null"`                           //价格
		PromotionPrice             float64   `json:"promotion_price" gorm:"type:float64;comment:促销价格;not null"`               //促销价格
		GiftGrowth                 int64     `json:"gift_growth" gorm:"type:int64;comment:赠送的成长值;not null;default:0"`         //赠送的成长值
		GiftPoint                  int64     `json:"gift_point" gorm:"type:int64;comment:赠送的积分;not null;default:0"`           //赠送的积分
		UsePointLimit              int64     `json:"use_point_limit" gorm:"type:int64;comment:限制使用的积分数;not null;"`            //限制使用的积分数
		SubTitle                   string    `json:"sub_title" gorm:"type:varchar(191);comment:副标题;not null"`                 //副标题
		Description                string    `json:"description" gorm:"type:varchar(191);comment:商品描述;not null"`              //商品描述
		OriginalPrice              float64   `json:"original_price" gorm:"type:float64;comment:市场价;not null"`                 //市场价
		Stock                      int       `json:"stock" gorm:"type:int;comment:库存;not null;default:0"`                     //库存
		LowStock                   int       `json:"low_stock" gorm:"type:int;comment:库存预警值;not null;default:0"`              //库存预警值
		Unit                       string    `json:"unit" gorm:"type:varchar(191);comment:单位;not null"`                       //单位
		Weight                     float64   `json:"weight" gorm:"type:float64;comment:商品重量;not null"`                        //商品重量，默认为克
		PreviewStatus              int64     `json:"preview_status" gorm:"type:int64;comment:是否为预告商品;not null;default:0"`     //是否为预告商品：0->不是；1->是
		ServiceIds                 string    `json:"service_ids" gorm:"type:varchar(191);comment:以逗号分割的产品服务;not null"`        //以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮',
		Keywords                   string    `json:"keywords" gorm:"type:varchar(191);comment:关键字;not null"`                  //关键字
		Note                       string    `json:"note" gorm:"type:varchar(191);comment:注释;not null"`                       //注释
		AlbumPics                  string    `json:"album_pics" gorm:"type:varchar(191);comment:画册图片;not null"`               //画册图片，连产品图片限制为5张，以逗号分割
		DetailTitle                string    `json:"detail_title" gorm:"type:varchar(191);comment:详情标题;not null"`             //详情标题
		DetailDesc                 string    `json:"detail_desc" gorm:"type:text;comment:详细说明;not null"`                      //详细说明
		DetailHtml                 string    `json:"detail_html" gorm:"type:text;comment:产品详情网页内容;not null"`                  //产品详情网页内容
		DetailMobileHtml           string    `json:"detail_mobile_html" gorm:"type:text;comment:移动端网页详情;not null"`            //移动端网页详情
		PromotionStartTime         time.Time `json:"promotion_start_time" gorm:"type:datetime;comment:促销开始时间;not null"`       //促销开始时间
		PromotionEndTime           time.Time `json:"promotion_end_time" gorm:"type:datetime;comment:促销结束时间;not null"`         //促销结束时间
		PromotionPerLimit          int64     `json:"promotion_per_limit" gorm:"type:int64;comment:活动限购数量;not null;default:0"` //活动限购数量
		PromotionType              int64     `json:"promotion_type" gorm:"type:int64;comment:促销类型;not null;default:0"`        //促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购
		BrandName                  string    `json:"brand_name" gorm:"type:text;comment:品牌名称;not null"`                       //品牌名称
		ProductCategoryName        string    `json:"product_category_name" gorm:"type:text;comment:商品分类名称;not null"`          //商品分类名称

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
