package log

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSysLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogDeleteLogic {
	return &SysLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SysLogDeleteLogic) SysLogDelete(req *types.DeleteSysLogReq) (*types.DeleteSysLogResp, error) {
	_, err := l.svcCtx.Sys.SysLogDelete(l.ctx, &sysclient.SysLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除登录日志失败")
	}
	return &types.DeleteSysLogResp{
		Code:    "200",
		Message: "删除登录日志成功",
	}, nil
}
