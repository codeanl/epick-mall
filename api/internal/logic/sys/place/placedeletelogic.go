package place

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceDeleteLogic {
	return &PlaceDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceDeleteLogic) PlaceDelete(req *types.DeletePlaceReq) (*types.DeletePlaceResp, error) {
	_, err := l.svcCtx.Sys.PlaceDelete(l.ctx, &sysclient.PlaceDeleteReq{
		Ids: req.Ids,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("删除用户失败")
	}
	return &types.DeletePlaceResp{
		Code:    "200",
		Message: "删除用户成功",
	}, nil
}
