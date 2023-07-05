package loginlog

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

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
	resp, err := l.svcCtx.Sys.LoginLogList(l.ctx, &sysclient.LoginLogListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Username: req.Username,
		Ip:       req.Ip,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListLoginLogData
	for _, item := range resp.List {
		listUserData := types.ListLoginLogData{
			Id:         item.Id,
			Username:   item.Username,
			Status:     "1",
			Ip:         item.Ip,
			CreateTime: item.CreateTime,
			Avatar:     item.Avatar,
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
