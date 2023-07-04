package menu

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuListLogic) MenuList(req *types.ListMenuReq) (*types.ListMenuResp, error) {
	resp, err := l.svcCtx.System.MenuList(l.ctx, &systemclient.MenuListReq{
		Name: req.Name,
		Url:  req.Url,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListMenuData
	for _, menu := range resp.List {
		listUserData := types.ListMenuData{
			Id:             menu.Id,
			Name:           menu.Name,
			ParentId:       menu.ParentId,
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           menu.Type,
			Icon:           menu.Icon,
			OrderNum:       menu.OrderNum,
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreateTime,
			LastUpdateBy:   menu.LastUpdateBy,
			LastUpdateTime: menu.LastUpdateTime,
			BackgroundUrl:  menu.BackgroundUrl,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
		}
		list = append(list, &listUserData)
	}
	return &types.ListMenuResp{
		Code:    "000000",
		Message: "查询菜单列表成功",
		Total:   resp.Total,
		Data:    list,
	}, nil
}
