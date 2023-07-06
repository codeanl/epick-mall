package homeadvertise

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sms/rpc/smsclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseDeleteLogic {
	return &HomeAdvertiseDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseDeleteLogic) HomeAdvertiseDelete(req *types.DeleteHomeAdvertiseReq) (*types.DeleteHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseDelete(l.ctx, &smsclient.HomeAdvertiseDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeleteHomeAdvertiseResp{
		Code:    "200",
		Message: "删除成功",
	}, nil
}
