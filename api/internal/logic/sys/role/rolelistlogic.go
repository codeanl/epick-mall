package role

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RoleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRoleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RoleListLogic {
	return &RoleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RoleListLogic) RoleList(req *types.ListRoleReq) (*types.ListRoleResp, error) {
	resp, err := l.svcCtx.Sys.RoleList(l.ctx, &sysclient.RoleListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListRoleData
	for _, item := range resp.List {
		listUserData := types.ListRoleData{
			Id:             item.Id,
			Name:           item.Name,
			Remark:         item.Remark,
			CreateBy:       item.CreateBy,
			CreateTime:     item.CreateTime,
			LastUpdateBy:   item.LastUpdateBy,
			LastUpdateTime: item.LastUpdateTime,
		}
		list = append(list, &listUserData)
	}

	return &types.ListRoleResp{
		Code:     "000000",
		Message:  "查询角色列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
