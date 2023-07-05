package user

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/service/sys/rpc/sysclient"

	"github.com/zeromicro/go-zero/core/logx"
)

type SelectAllDataLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSelectAllDataLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SelectAllDataLogic {
	return &SelectAllDataLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SelectAllDataLogic) SelectAllData(req *types.SelectDataReq) (resp *types.SelectDataResp, err error) {
	roleList, roleErr := l.svcCtx.Sys.RoleList(l.ctx, &sysclient.RoleListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	})
	if roleErr != nil {
		return nil, roleErr
	}
	var roleData []*types.RoleAllResp
	for _, role := range roleList.List {
		roleData = append(roleData, &types.RoleAllResp{
			Id:     role.Id,
			Name:   role.Name,
			Remark: role.Remark,
		})
	}
	return &types.SelectDataResp{
		Code:    "000000",
		Message: "",
		RoleAll: roleData,
	}, nil
}
