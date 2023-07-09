package logic

import (
	"context"
	"epick-mall/service/sms/model"
	"epick-mall/service/sms/rpc/internal/svc"
	"epick-mall/service/sms/rpc/sms"
	"errors"
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

func (l *CouponUpdateLogic) CouponUpdate(in *sms.CouponUpdateReq) (*sms.CouponUpdateResp, error) {
	err := l.svcCtx.CouponModel.UpdateCoupon(in.Id, &model.Coupon{
		Type:         in.Type,
		Name:         in.Name,
		Platform:     in.Platform,
		Count:        in.Count,
		Amount:       in.Amount,
		PerLimit:     in.PerLimit,
		MinPoint:     in.MinPoint,
		StartTime:    in.StartTime,
		EndTime:      in.EndTime,
		UseType:      in.UseType,
		Note:         in.Note,
		UseCount:     in.UseCount,
		ReceiveCount: in.ReceiveCount,
		EnableTime:   in.EnableTime,
		Code:         in.Code,
	})
	if err != nil {
		return nil, errors.New("更新失败")
	}

	return &sms.CouponUpdateResp{}, nil
}
