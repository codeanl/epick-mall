package logic

import (
	"context"
	"epick-mall/service/pms/model"
	"errors"

	"epick-mall/service/pms/rpc/internal/svc"
	"epick-mall/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductUpdateLogic {
	return &ProductUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc ProductList(ProductListReq) returns(ProductListResp);
func (l *ProductUpdateLogic) ProductUpdate(in *pms.ProductUpdateReq) (*pms.ProductUpdateResp, error) {
	err := l.svcCtx.ProductModel.UpdateProduct(in.Id, &model.Product{
		BrandID:                    in.BrandId,
		ProductCategoryID:          in.ProductCategoryId,
		FreightTemplateID:          in.FeightTemplateId,
		ProductAttributeCategoryID: in.ProductAttributeCategoryId,
		Name:                       in.Name,
		Pic:                        in.Pic,
		ProductSn:                  in.ProductSn,
		DeleteStatus:               in.DeleteStatus,
		PublishStatus:              in.PublishStatus,
		NewStatus:                  in.NewStatus,
		RecommendStatus:            in.RecommandStatus,
		Sort:                       in.Sort,
		Sale:                       in.Sale,
		Price:                      in.Price,
		PromotionPrice:             in.PromotionPrice,
		GiftGrowth:                 in.GiftGrowth,
		GiftPoint:                  in.GiftPoint,
		UsePointLimit:              in.UsePointLimit,
		SubTitle:                   in.SubTitle,
		Description:                in.Description,
		OriginalPrice:              in.OriginalPrice,
		Stock:                      int(in.Stock),
		LowStock:                   int(in.LowStock),
		Unit:                       in.Unit,
		Weight:                     in.Weight,
		PreviewStatus:              in.PreviewStatus,
		ServiceIds:                 in.ServiceIds,
		Keywords:                   in.Keywords,
		Note:                       in.Note,
		AlbumPics:                  in.AlbumPics,
		DetailTitle:                in.DetailTitle,
		DetailDesc:                 in.DetailDesc,
		DetailHtml:                 in.DetailHtml,
		DetailMobileHtml:           in.DetailMobileHtml,
		PromotionStartTime:         in.PromotionStartTime,
		PromotionEndTime:           in.PromotionEndTime,
		PromotionPerLimit:          in.PromotionPerLimit,
		PromotionType:              in.PromotionType,
		BrandName:                  in.BrandName,
		ProductCategoryName:        in.ProductCategoryName,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}

	return &pms.ProductUpdateResp{}, nil
}
