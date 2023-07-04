package logic

import (
	"context"
	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RoleListLogic) RoleList(in *system.RoleListReq) (*system.RoleListResp, error) {
	all, total, err := l.svcCtx.RoleModel.GetRoleList(in)
	if err != nil {
		return nil, err
	}
	var list []*system.RoleListData
	for _, role := range all {
		list = append(list, &system.RoleListData{
			Id:             int64(role.ID),
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreatedAt.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.UpdateBy,
			LastUpdateTime: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &system.RoleListResp{
		Total: total,
		List:  list,
	}, nil
}
