package role

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"
	"strconv"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMenuByRoleIdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMenuByRoleIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMenuByRoleIdLogic {
	return &QueryMenuByRoleIdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMenuByRoleIdLogic) QueryMenuByRoleId(req *types.RoleMenuReq) (*types.RoleMenuResp, error) {
	resp, _ := l.svcCtx.Sys.MenuList(l.ctx, &sysclient.MenuListReq{
		Name: "",
		Url:  "",
	})
	var list []*types.ListtMenuData
	var listIds []int64
	for _, menu := range resp.List {
		list = append(list, &types.ListtMenuData{
			Key:      strconv.FormatInt(menu.Id, 10),
			Title:    menu.Name,
			ParentId: menu.ParentId,
			Id:       menu.Id,
			Label:    menu.Name,
		})
		//admin账号全部权限
		listIds = append(listIds, menu.Id)
	}
	//如果角色不是admin则根据roleId查询菜单
	if req.Id != 1 {
		QueryMenu, err := l.svcCtx.Sys.QueryMenuByRoleId(l.ctx, &sysclient.QueryMenuByRoleIdReq{
			Id: req.Id,
		})
		listIds = QueryMenu.Ids
		if err != nil {
			return nil, errorx.NewDefaultError("查询非管理员的菜单id失败失败")
		}
	}
	return &types.RoleMenuResp{
		RoleData: listIds,
		AllData:  list,
		Code:     "000000",
		Message:  "根据角色id查询菜单成功",
	}, nil
}
