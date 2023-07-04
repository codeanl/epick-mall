package logic

import (
	"context"
	"epick-mall/service/system/model"
	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuAddLogic) MenuAdd(in *system.MenuAddReq) (*system.MenuAddResp, error) {
	menu := &model.Menu{
		Name:          in.Name,
		ParentID:      int(in.ParentId),
		Url:           in.Url,
		Perms:         in.Perms,
		Type:          int(in.Type),
		Icon:          in.Icon,
		OrderNum:      int(in.OrderNum),
		CreateBy:      in.CreateBy,
		BackgroundUrl: in.BackgroundUrl,
		VuePath:       in.VuePath,
		VueComponent:  in.VueComponent,
		VueIcon:       in.VueIcon,
		VueRedirect:   in.VueRedirect,
	}
	err := l.svcCtx.MenuModel.AddMenu(menu)
	if err != nil {
		return nil, errors.New("添加菜单失败")
	}
	return &system.MenuAddResp{}, nil
}
