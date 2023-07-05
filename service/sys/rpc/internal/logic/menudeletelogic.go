package logic

import (
	"context"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *MenuDeleteLogic) MenuDelete(in *sys.MenuDeleteReq) (*sys.MenuDeleteResp, error) {
	err := l.svcCtx.MenuModel.DeleteMenuByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}

	return &sys.MenuDeleteResp{}, nil
}
