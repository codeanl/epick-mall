package logic

import (
	"context"
	"epick-mall/service/coupon/model"
	"epick-mall/service/coupon/rpc/coupon"
	"epick-mall/service/coupon/rpc/internal/svc"
	"errors"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponAddLogic) CouponAdd(in *coupon.CouponAddReq) (*coupon.CouponAddResp, error) {
	startTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	endTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	enableTime, _ := time.Parse("2006-01-02 15:04:05", in.EnableTime)
	info := &model.Coupon{
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Count:        in.Count,
		Amount:       float64(in.Amount),
		PerLimit:     in.PerLimit,
		MinPoint:     float64(in.MinPoint),
		StartTime:    startTime,
		EndTime:      endTime,
		UseType:      in.UseType,
		Note:         in.Note,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   enableTime,
		Code:         in.Code,
	}
	err := l.svcCtx.CouponModel.AddCoupon(info)
	if err != nil {
		return nil, errors.New("添加优惠券失败")
	}

	return &coupon.CouponAddResp{}, nil
}
