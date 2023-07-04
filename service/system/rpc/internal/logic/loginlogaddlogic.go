package logic

import (
	"context"
	"epick-mall/service/system/model"
	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogAddLogic {
	return &LoginLogAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogAddLogic) LoginLogAdd(in *system.LoginLogAddReq) (*system.LoginLogAddResp, error) {
	loginLog := &model.LoginLog{
		Phone:  in.Phone,
		Status: in.Status,
		IP:     in.Ip,
	}
	err := l.svcCtx.LoginLogModel.AddLoginLog(loginLog)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &system.LoginLogAddResp{}, nil
}
