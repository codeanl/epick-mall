package logic

import (
	"context"
	"epick-mall/service/product/model"
	"epick-mall/service/product/rpc/internal/svc"
	"epick-mall/service/product/rpc/product"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ProductAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductAddLogic {
	return &ProductAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductAddLogic) ProductAdd(in *product.ProductAddReq) (*product.ProductAddResp, error) {
	PromotionStartTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionStartTime)
	PromotionEndTime, _ := time.Parse("2006-01-02 15:04:05", in.PromotionEndTime)
	role := &model.Product{
		BrandID:                    in.BrandId,
		ProductCategoryID:          in.ProductCategoryId,
		FreightTemplateID:          in.FreightTemplateId,
		ProductAttributeCategoryID: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		Pic:                        in.Pic,
		ProductSn:                  in.ProductSn,
		DeleteStatus:               in.DeleteStatus,
		NewStatus:                  in.NewStatus,
		RecommendStatus:            in.RecommendStatus,
		Sort:                       in.Sort,
		Sale:                       in.Sale,
		Price:                      float64(in.Price),
		PromotionPrice:             float64(in.PromotionPrice),
		GiftGrowth:                 in.GiftGrowth,
		GiftPoint:                  in.GiftPoint,
		UsePointLimit:              in.UsePointLimit,
		SubTitle:                   in.SubTitle,
		Description:                in.Description,
		OriginalPrice:              float64(in.OriginalPrice),
		Stock:                      int(in.Stock),
		LowStock:                   int(in.LowStock),
		Unit:                       in.Unit,
		Weight:                     float64(in.Weight),
		PreviewStatus:              in.PreviewStatus,
		ServiceIds:                 in.ServiceIds,
		Keywords:                   in.Keywords,
		Note:                       in.Note,
		AlbumPics:                  in.AlbumPics,
		DetailTitle:                in.DetailTitle,
		DetailDesc:                 in.DetailDesc,
		DetailHtml:                 in.DetailHtml,
		DetailMobileHtml:           in.DetailMobileHtml,
		PromotionStartTime:         PromotionStartTime,
		PromotionEndTime:           PromotionEndTime,
		PromotionPerLimit:          in.PromotionPerLimit,
		PromotionType:              in.PromotionType,
		BrandName:                  in.BrandName,
		ProductCategoryName:        in.ProductCategoryName,
	}
	err := l.svcCtx.ProductModel.AddProduct(role)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}

	return &product.ProductAddResp{}, nil
}
