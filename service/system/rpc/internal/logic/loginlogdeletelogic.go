package logic

import (
	"context"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogDeleteLogic {
	return &LoginLogDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(in *system.LoginLogDeleteReq) (*system.LoginLogDeleteResp, error) {
	err := l.svcCtx.LoginLogModel.DeleteLoginLogByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &system.LoginLogDeleteResp{}, nil
}
