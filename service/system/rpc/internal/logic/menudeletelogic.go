package logic

import (
	"context"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuDeleteLogic) MenuDelete(in *system.MenuDeleteReq) (*system.MenuDeleteResp, error) {
	err := l.svcCtx.MenuModel.DeleteMenuByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &system.MenuDeleteResp{}, nil
}
