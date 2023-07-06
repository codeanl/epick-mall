package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *LogAddLogic) LogAdd(in *sys.LogAddReq) (*sys.LogAddResp, error) {
	log := &model.Log{
		Username:  in.Username,
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
	return &sys.LogAddResp{}, nil
}
