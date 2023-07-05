package role

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleDeleteLogic {
	return &RoleDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleDeleteLogic) RoleDelete(req *types.DeleteRoleReq) (*types.DeleteRoleResp, error) {
	_, err := l.svcCtx.Sys.RoleDelete(l.ctx, &sysclient.RoleDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeleteRoleResp{
		Code:    "200",
		Message: "删除用户成功",
	}, nil
}
