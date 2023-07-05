package logic

import (
	"context"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(in *sys.QueryMenuByRoleIdReq) (*sys.QueryMenuByRoleIdResp, error) {
	RoleMenus, err := l.svcCtx.RoleMenuModel.GetMenuByRoleId(in.Id)
	if err != nil {
		return nil, errors.New("查询角色菜单失败")
	}
	var list []int64
	for _, user := range RoleMenus {
		list = append(list, int64(user.MenuID))
	}
	return &sys.QueryMenuByRoleIdResp{
		Ids: list,
	}, nil
}
