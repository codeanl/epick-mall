package coupon

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/coupon/rpc/couponclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponListLogic) CouponList(req *types.ListCouponReq) (*types.ListCouponResp, error) {
	resp, err := l.svcCtx.Coupon.CouponList(l.ctx, &couponclient.CouponListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListCouponData
	for _, item := range resp.List {
		listUserData := types.ListCouponData{
			Id:         item.Id,
			Type:       item.Type,
			Name:       item.Name,
			Platform:   item.Platform,
			Count:      item.Count,
			Amount:     item.Amount,
			PerLimit:   item.PerLimit,
			MinPoint:   item.MinPoint,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			UseType:    item.UseType,
			Note:       item.Note,
			EnableTime: item.EnableTime,
			Code:       item.Code,
		}
		list = append(list, &listUserData)
	}
	return &types.ListCouponResp{
		Code:     "000000",
		Message:  "查询优惠券列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
