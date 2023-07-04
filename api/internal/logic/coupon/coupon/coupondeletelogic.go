package coupon

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/coupon/rpc/couponclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDeleteLogic {
	return &CouponDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponDeleteLogic) CouponDelete(req *types.DeleteCouponReq) (*types.DeleteCouponResp, error) {
	_, err := l.svcCtx.Coupon.CouponDelete(l.ctx, &couponclient.CouponDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除优惠券失败")
	}
	return &types.DeleteCouponResp{
		Code:    "200",
		Message: "删除优惠券成功",
	}, nil
}
