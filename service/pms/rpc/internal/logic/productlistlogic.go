package logic

import (
	"context"
	"epick-mall/service/pms/rpc/internal/svc"
	"epick-mall/service/pms/rpc/pms"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductListLogic) ProductList(in *pms.ProductListReq) (*pms.ProductListResp, error) {
	all, total, err := l.svcCtx.ProductModel.GetProductList(in)
	if err != nil {
		return nil, err
	}
	var list []*pms.ProductListData
	for _, role := range all {
		list = append(list, &pms.ProductListData{
			Id:                         int64(role.ID),
			BrandId:                    role.BrandID,
			ProductCategoryId:          role.ProductCategoryID,
			FeightTemplateId:           role.FreightTemplateID,
			ProductAttributeCategoryId: role.ProductAttributeCategoryID,
			Name:                       role.Name,
			Pic:                        role.Pic,
			ProductSn:                  role.ProductSn,
			DeleteStatus:               role.DeleteStatus,
			PublishStatus:              role.PublishStatus,
			NewStatus:                  role.NewStatus,
			RecommandStatus:            role.RecommendStatus,
			Sort:                       role.Sort,
			Sale:                       role.Sale,
			Price:                      role.Price,
			PromotionPrice:             role.PromotionPrice,
			GiftGrowth:                 role.GiftGrowth,
			GiftPoint:                  role.GiftPoint,
			UsePointLimit:              role.UsePointLimit,
			SubTitle:                   role.SubTitle,
			Description:                role.Description,
			OriginalPrice:              role.OriginalPrice,
			Stock:                      int64(role.Stock),
			LowStock:                   int64(role.LowStock),
			Unit:                       role.Unit,
			Weight:                     role.Weight,
			PreviewStatus:              role.PreviewStatus,
			ServiceIds:                 role.ServiceIds,
			Keywords:                   role.Keywords,
			Note:                       role.Note,
			AlbumPics:                  role.AlbumPics,
			DetailTitle:                role.DetailTitle,
			DetailDesc:                 role.DetailDesc,
			DetailHtml:                 role.DetailHtml,
			DetailMobileHtml:           role.DetailMobileHtml,
			PromotionStartTime:         role.PromotionStartTime,
			PromotionEndTime:           role.PromotionEndTime,
			PromotionPerLimit:          role.PromotionPerLimit,
			PromotionType:              role.PromotionType,
			BrandName:                  role.BrandName,
			ProductCategoryName:        role.ProductCategoryName,
		})
	}
	return &pms.ProductListResp{
		Total: total,
		List:  list,
	}, nil
}
