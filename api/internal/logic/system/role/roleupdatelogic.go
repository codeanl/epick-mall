package role

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleUpdateLogic) RoleUpdate(req *types.UpdateRoleReq) (resp *types.UpdateRoleResp, err error) {
	_, err = l.svcCtx.System.RoleUpdate(l.ctx, &systemclient.RoleUpdateReq{
		Id:       req.Id,
		Name:     req.Name,
		Remark:   req.Remark,
		UpdateBy: l.ctx.Value("nickname").(string),
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateRoleResp{
		Code:    "200",
		Message: "更新用户成功",
	}, nil
}
