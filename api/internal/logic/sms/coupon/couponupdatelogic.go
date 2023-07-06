package coupon

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/sms/rpc/smsclient"

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

func (l *CouponUpdateLogic) CouponUpdate(req *types.UpdateCouponReq) (*types.UpdateCouponResp, error) {
	_, err := l.svcCtx.Sms.CouponUpdate(l.ctx, &smsclient.CouponUpdateReq{
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
		logx.Error(err)
		return nil, errorx.NewDefaultError("更新失败")
	}
	return &types.UpdateCouponResp{
		Code:    "000000",
		Message: "更新成功",
	}, nil
}
