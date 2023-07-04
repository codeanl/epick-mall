package logic

import (
	"context"
	"epick-mall/service/system/model"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type LogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogAddLogic {
	return &LogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LogAddLogic) LogAdd(in *system.LogAddReq) (*system.LogAddResp, error) {
	log := &model.Log{
		Phone:     in.Phone,
		Operation: in.Operation,
		Method:    in.Method,
		Params:    in.Params,
		Time:      in.Time,
		IP:        in.Ip,
	}
	err := l.svcCtx.LogModel.AddLog(log)
	if err != nil {
		return nil, errors.New("添加日志失败")
	}
	return &system.LogAddResp{}, nil
}
