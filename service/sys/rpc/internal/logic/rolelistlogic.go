package logic

import (
	"context"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *RoleListLogic) RoleList(in *sys.RoleListReq) (*sys.RoleListResp, error) {
	all, total, err := l.svcCtx.RoleModel.GetRoleList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.RoleListData
	for _, role := range all {
		list = append(list, &sys.RoleListData{
			Id:             int64(role.ID),
			Name:           role.Name,
			Remark:         role.Remark,
			CreateBy:       role.CreateBy,
			CreateTime:     role.CreatedAt.Format("2006-01-02 15:04:05"),
			LastUpdateBy:   role.UpdateBy,
			LastUpdateTime: role.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &sys.RoleListResp{
		Total: total,
		List:  list,
	}, nil

	return &sys.RoleListResp{}, nil
}
