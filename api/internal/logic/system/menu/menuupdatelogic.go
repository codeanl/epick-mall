package menu

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuUpdateLogic {
	return &MenuUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUpdateLogic) MenuUpdate(req *types.UpdateMenuReq) (resp *types.UpdateMenuResp, err error) {
	_, err = l.svcCtx.System.MenuUpdate(l.ctx, &systemclient.MenuUpdateReq{
		Id:            req.Id,
		Name:          req.Name,
		ParentId:      req.ParentId,
		Url:           req.Url,
		Perms:         req.Perms,
		Type:          req.Type,
		Icon:          req.Icon,
		OrderNum:      req.OrderNum,
		LastUpdateBy:  l.ctx.Value("nickname").(string),
		VuePath:       req.VuePath,
		VueComponent:  req.VueComponent,
		VueIcon:       req.VueIcon,
		VueRedirect:   req.VueRedirect,
		BackgroundUrl: req.BackgroundUrl,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateMenuResp{
		Code:    "200",
		Message: "更新菜单成功",
	}, nil
}
