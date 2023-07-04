package role

import (
	"context"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleAddLogic) RoleAdd(req *types.AddRoleReq) (resp *types.AddRoleResp, err error) {
	_, err = l.svcCtx.System.RoleAdd(l.ctx, &systemclient.RoleAddReq{
		Name:     req.Name,
		Remark:   req.Remark,
		CreateBy: l.ctx.Value("nickname").(string),
	})
	if err != nil {
		return nil, err
	}
	return &types.AddRoleResp{
		Code:    "200",
		Message: "添加角色成功",
	}, nil
	return
}
