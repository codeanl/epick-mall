package logic

import (
	"context"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuListLogic {
	return &MenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuListLogic) MenuList(in *system.MenuListReq) (*system.MenuListResp, error) {
	all, total, err := l.svcCtx.MenuModel.GetMenuList(in)
	if err != nil {
		return nil, err
	}
	var list []*system.MenuListData
	for _, menu := range all {
		list = append(list, &system.MenuListData{
			Id:             int64(menu.ID),
			Name:           menu.Name,
			ParentId:       int64(menu.ParentID),
			Url:            menu.Url,
			Perms:          menu.Perms,
			Type:           int64(menu.Type),
			Icon:           menu.Icon,
			OrderNum:       int64(menu.OrderNum),
			CreateBy:       menu.CreateBy,
			CreateTime:     menu.CreatedAt.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   menu.UpdateBy,
			LastUpdateTime: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
			BackgroundUrl:  menu.BackgroundUrl,
			VuePath:        menu.VuePath,
			VueComponent:   menu.VueComponent,
			VueIcon:        menu.VueIcon,
			VueRedirect:    menu.VueRedirect,
		})
	}
	return &system.MenuListResp{
		Total: total,
		List:  list,
	}, nil
}
