package logic

import (
	"context"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleDeleteLogic) RoleDelete(in *system.RoleDeleteReq) (*system.RoleDeleteResp, error) {
	err := l.svcCtx.RoleModel.DeleteRoleByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除用户失败")
	}
	return &system.RoleDeleteResp{}, nil
}
