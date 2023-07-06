package logic

import (
	"context"
	"epick-mall/service/sms/model"
	"epick-mall/service/sms/rpc/sms"
	"errors"
	"time"

	"epick-mall/service/sms/rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(in *sms.HomeAdvertiseAddReq) (*sms.HomeAdvertiseAddResp, error) {
	StartTime, _ := time.Parse("2006-01-02 15:04:05", in.StartTime)
	EndTime, _ := time.Parse("2006-01-02 15:04:05", in.EndTime)
	role := &model.HomeAdvertise{
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
	}
	err := l.svcCtx.HomeAdvertiseModel.AddHomeAdvertise(role)
	if err != nil {
		return nil, errors.New("添加失败")
	}
	return &sms.HomeAdvertiseAddResp{}, nil
}
