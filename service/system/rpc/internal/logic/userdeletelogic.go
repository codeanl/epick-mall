package logic

import (
	"context"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *system.UserDeleteReq) (*system.UserDeleteResp, error) {
	err := l.svcCtx.UserModel.DeleteByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &system.UserDeleteResp{}, nil
}
