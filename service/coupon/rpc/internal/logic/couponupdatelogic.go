package logic

import (
	"context"
	"epick-mall/service/coupon/model"
	"errors"
	"time"

	"epick-mall/service/coupon/rpc/coupon"
	"epick-mall/service/coupon/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponUpdateLogic {
	return &CouponUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc CouponList(CouponListReq) returns(CouponListResp);
func (l *CouponUpdateLogic) CouponUpdate(in *coupon.CouponUpdateReq) (*coupon.CouponUpdateResp, error) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	enableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)
	err := l.svcCtx.CouponModel.UpdateCoupon(in.Id, &model.Coupon{
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Count:        in.Count,
		Amount:       in.Amount,
		PerLimit:     in.PerLimit,
		MinPoint:     in.MinPoint,
		StartTime:    startTime,
		EndTime:      endTime,
		UseType:      in.UseType,
		Note:         in.Note,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   enableTime,
		Code:         in.Code,
	})
	if err != nil {
		return nil, errors.New("更新优惠券失败")
	}
	return &coupon.CouponUpdateResp{}, nil
}
