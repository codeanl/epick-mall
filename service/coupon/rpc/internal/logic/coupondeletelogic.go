package logic

import (
	"context"
	"epick-mall/service/coupon/rpc/coupon"
	"epick-mall/service/coupon/rpc/internal/svc"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponDeleteLogic {
	return &CouponDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponDeleteLogic) CouponDelete(in *coupon.CouponDeleteReq) (*coupon.CouponDeleteResp, error) {
	err := l.svcCtx.CouponModel.DeleteCouponByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &coupon.CouponDeleteResp{}, nil
}
