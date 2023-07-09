package coupon

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/service/sms/rpc/smsclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponAddLogic {
	return &CouponAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponAddLogic) CouponAdd(req *types.AddCouponReq) (resp *types.AddCouponResp, err error) {
	_, err = l.svcCtx.Sms.CouponAdd(l.ctx, &smsclient.CouponAddReq{
		Type:       req.Type,
		Name:       req.Name,
		Platform:   req.Platform,
		Count:      req.Count,
		Amount:     float32(req.Amount),
		PerLimit:   req.PerLimit,
		MinPoint:   float32(req.MinPoint),
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		UseType:    req.UseType,
		Note:       req.Note,
		EnableTime: req.EnableTime,
		Code:       req.Code,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddCouponResp{
		Code:    "200",
		Message: "添加成功",
	}, nil
	return
}
