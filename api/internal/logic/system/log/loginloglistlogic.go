package log

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogListLogic) LoginLogList(req *types.ListLoginLogReq) (*types.ListLoginLogResp, error) {
	resp, err := l.svcCtx.System.LoginLogList(l.ctx, &systemclient.LoginLogListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Phone:    req.Phone,
		Ip:       req.Ip,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListLoginLogData
	for _, item := range resp.List {
		listUserData := types.ListLoginLogData{
			Id:         item.Id,
			Phone:      item.Phone,
			Status:     item.Status,
			Ip:         item.Ip,
			CreateTime: item.CreateTime,
		}
		list = append(list, &listUserData)
	}

	return &types.ListLoginLogResp{
		Code:     "000000",
		Message:  "查询登录日志列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
