package logic

import (
	"context"
	"epick-mall/service/sys/model"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMenuRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMenuRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMenuRoleLogic {
	return &UpdateMenuRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateMenuRoleLogic) UpdateMenuRole(in *sys.UpdateMenuRoleReq) (*sys.UpdateMenuRoleResp, error) {
	_ = l.svcCtx.RoleMenuModel.DeleteByRoleId(in.RoleId)
	for _, menuId := range in.MenuIds {
		_ = l.svcCtx.RoleMenuModel.AddRoleMenu(&model.RoleMenu{
			RoleID:   int(in.RoleId),
			MenuID:   int(menuId),
			CreateBy: in.CreateBy,
		})
	}
	return &sys.UpdateMenuRoleResp{}, nil
}
