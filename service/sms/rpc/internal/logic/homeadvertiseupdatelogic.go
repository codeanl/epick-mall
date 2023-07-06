package logic

import (
	"context"
	"epick-mall/service/sms/model"
	"epick-mall/service/sms/rpc/internal/svc"
	"epick-mall/service/sms/rpc/sms"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type HomeAdvertiseUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(in *sms.HomeAdvertiseUpdateReq) (*sms.HomeAdvertiseUpdateResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	err := l.svcCtx.HomeAdvertiseModel.UpdateHomeAdvertise(in.Id, &model.HomeAdvertise{
		Name:       in.Name,
		Type:       in.Type,
		Pic:        in.Pic,
		StartTime:  StartTime,
		EndTime:    EndTime,
		Status:     in.Status,
		ClickCount: in.ClickCount,
		OrderCount: in.OrderCount,
		Url:        in.Url,
		Note:       in.Note,
		Sort:       in.Sort,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &sms.HomeAdvertiseUpdateResp{}, nil
}
