package logic

import (
	"context"
	"epick-mall/service/system/model"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleUpdateLogic {
	return &RoleUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}
func (l *RoleUpdateLogic) RoleUpdate(in *system.RoleUpdateReq) (*system.RoleUpdateResp, error) {
	err := l.svcCtx.RoleModel.UpdateRole(in.Id, &model.Role{
		Name:     in.Name,
		Remark:   in.Remark,
		UpdateBy: in.UpdateBy,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &system.RoleUpdateResp{}, nil
}
