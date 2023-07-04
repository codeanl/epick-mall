package logic

import (
	"context"
	"epick-mall/service/coupon/rpc/coupon"
	"epick-mall/service/coupon/rpc/internal/svc"

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

func (l *CouponListLogic) CouponList(in *coupon.CouponListReq) (*coupon.CouponListResp, error) {
	all, total, err := l.svcCtx.CouponModel.GetCouponList(in)
	if err != nil {
		return nil, err
	}
	var list []*coupon.CouponListData
	for _, role := range all {
		list = append(list, &coupon.CouponListData{
			Id:           int64(role.ID),
			Type:         role.Type,
			Name:         role.Name,
			Platform:     role.Platform,
			Count:        role.Count,
			Amount:       role.Amount,
			PerLimit:     role.PerLimit,
			MinPoint:     role.MinPoint,
			StartTime:    role.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:      role.EndTime.Format("2006-01-02 15:04:05"),
			UseType:      role.UseType,
			Note:         role.Note,
			UseCount:     role.UseCount,
			ReceiveCount: role.ReceiveCount,
			EnableTime:   role.EnableTime.Format("2006-01-02 15:04:05"),
			Code:         role.Code,
		})
	}
	return &coupon.CouponListResp{
		Total: total,
		List:  list,
	}, nil

}
