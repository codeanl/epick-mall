package log

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogListLogic) SysLogList(req *types.ListSysLogReq) (*types.ListSysLogResp, error) {
	resp, err := l.svcCtx.Sys.SysLogList(l.ctx, &sysclient.SysLogListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Username: req.Username,
		Method:   req.Method,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListSysLogData
	for _, item := range resp.List {
		listUserData := types.ListSysLogData{
			Id:        item.Id,
			Username:  item.Username,
			Operation: item.Operation,
			Method:    item.Method,
			Params:    item.Params,
			Time:      item.Time,
			Ip:        item.Ip,
			Avatar:    item.Avatar,
		}
		list = append(list, &listUserData)
	}
	return &types.ListSysLogResp{
		Code:     "000000",
		Message:  "查询登录日志列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
