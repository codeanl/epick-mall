package logic

import (
	"context"
	"errors"

	"epick-mall/service/sms/rpc/internal/svc"
	"epick-mall/service/sms/rpc/sms"

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

func (l *CouponDeleteLogic) CouponDelete(in *sms.CouponDeleteReq) (*sms.CouponDeleteResp, error) {
	err := l.svcCtx.CouponModel.DeleteCouponByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}

	return &sms.CouponDeleteResp{}, nil
}
