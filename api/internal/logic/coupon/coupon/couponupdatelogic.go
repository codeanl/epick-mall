package coupon

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/coupon/rpc/couponclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponUpdateLogic {
	return &CouponUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponUpdateLogic) CouponUpdate(req *types.UpdateCouponReq) (resp *types.UpdateCouponResp, err error) {
	_, err = l.svcCtx.Coupon.CouponUpdate(l.ctx, &couponclient.CouponUpdateReq{
		Id:         req.Id,
		Type:       req.Type,
		Name:       req.Name,
		Platform:   req.Platform,
		Count:      req.Count,
		Amount:     req.Amount,
		PerLimit:   req.PerLimit,
		MinPoint:   req.MinPoint,
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		UseType:    req.UseType,
		Note:       req.Note,
		EnableTime: req.EnableTime,
		Code:       req.Code,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateCouponResp{
		Code:    "200",
		Message: "更新优惠码成功",
	}, nil
}
