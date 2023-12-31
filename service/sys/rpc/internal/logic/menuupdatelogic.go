package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MenuUpdateLogic) MenuUpdate(in *sys.MenuUpdateReq) (*sys.MenuUpdateResp, error) {
	err := l.svcCtx.MenuModel.UpdateMenu(in.Id, &model.Menu{
		Name:          in.Name,
		ParentID:      int(in.ParentId),
		Url:           in.Url,
		Perms:         in.Perms,
		Type:          int(in.Type),
		Icon:          in.Icon,
		OrderNum:      int(in.OrderNum),
		UpdateBy:      in.LastUpdateBy,
		VuePath:       in.VuePath,
		VueComponent:  in.VueComponent,
		VueIcon:       in.VueIcon,
		VueRedirect:   in.VueRedirect,
		BackgroundUrl: in.BackgroundUrl,
	})
	if err != nil {
		return nil, errors.New("更新菜单失败")
	}

	return &sys.MenuUpdateResp{}, nil
}
