package logic

import (
	"context"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *system.QueryMenuByRoleIdReq) (*system.QueryMenuByRoleIdResp, error) {
	RoleMenus, err := l.svcCtx.RoleMenuModel.GetMenuByRoleId(in.Id)
	if err != nil {
		return nil, errors.New("查询角色菜单失败")
	}
	var list []int64
	for _, user := range RoleMenus {
		list = append(list, int64(user.MenuID))
	}
	return &system.QueryMenuByRoleIdResp{
		Ids: list,
	}, nil
}
