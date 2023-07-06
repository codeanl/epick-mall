package logic

import (
	"context"
	"epick-mall/service/sms/rpc/internal/svc"
	"epick-mall/service/sms/rpc/sms"

	"github.com/zeromicro/go-zero/core/logx"
)

type CouponListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCouponListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CouponListLogic {
	return &CouponListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CouponListLogic) CouponList(in *sms.CouponListReq) (*sms.CouponListResp, error) {
	all, total, err := l.svcCtx.CouponModel.GetCouponList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.CouponListData
	for _, role := range all {
		list = append(list, &sms.CouponListData{
			Id:           int64(role.ID),
			Type:         role.Type,
			Name:         role.Name,
			Platform:     role.Platform,
			Count:        role.Count,
			Amount:       role.Amount,
			PerLimit:     role.PerLimit,
			MinPoint:     role.MinPoint,
			StartTime:    role.StartTime,
			EndTime:      role.EndTime,
			UseType:      role.UseType,
			Note:         role.Note,
			UseCount:     role.UseCount,
			ReceiveCount: role.ReceiveCount,
			EnableTime:   role.EnableTime,
			Code:         role.Code,
		})
	}
	return &sms.CouponListResp{
		Total: total,
		List:  list,
	}, nil
}
