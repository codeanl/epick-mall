package menu

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MenuAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuAddLogic {
	return &MenuAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuAddLogic) MenuAdd(req *types.AddMenuReq) (*types.AddMenuResp, error) {
	_, err := l.svcCtx.Sys.MenuAdd(l.ctx, &sysclient.MenuAddReq{
		Name:         req.Name,
		ParentId:     req.ParentId,
		Url:          req.Url,
		Perms:        req.Perms,
		Type:         req.Type,
		Icon:         req.Icon,
		OrderNum:     req.OrderNum,
		CreateBy:     l.ctx.Value("nickname").(string),
		VuePath:      req.VuePath,
		VueComponent: req.VueComponent,
		VueIcon:      req.VueIcon,
		VueRedirect:  req.VueRedirect,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("添加菜单失败")
	}
	return &types.AddMenuResp{
		Code:    "000000",
		Message: "添加菜单成功!",
	}, nil
}
