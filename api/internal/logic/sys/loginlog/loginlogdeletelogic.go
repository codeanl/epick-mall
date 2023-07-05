package loginlog

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogDeleteLogic {
	return &LoginLogDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogDeleteLogic) LoginLogDelete(req *types.DeleteLoginLogReq) (*types.DeleteLoginLogResp, error) {
	_, err := l.svcCtx.Sys.LoginLogDelete(l.ctx, &sysclient.LoginLogDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除登录日志失败")
	}
	return &types.DeleteLoginLogResp{
		Code:    "200",
		Message: "删除登录日志成功",
	}, nil

}
