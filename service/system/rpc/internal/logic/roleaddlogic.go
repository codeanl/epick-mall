package logic

import (
	"context"
	"epick-mall/service/system/model"
	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleAddLogic {
	return &RoleAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleAddLogic) RoleAdd(in *system.RoleAddReq) (*system.RoleAddResp, error) {
	role := &model.Role{
		Name:     in.Name,
		Remark:   in.Remark,
		CreateBy: in.CreateBy,
	}
	err := l.svcCtx.RoleModel.AddRole(role)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	return &system.RoleAddResp{}, nil
}
