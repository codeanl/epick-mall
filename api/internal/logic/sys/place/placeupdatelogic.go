package place

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceUpdateLogic) PlaceUpdate(req *types.UpdatePlaceReq) (resp *types.UpdatePlaceResp, err error) {
	_, err = l.svcCtx.Sys.PlaceUpdate(l.ctx, &sysclient.PlaceUpdateReq{
		Id:        req.Id,
		Name:      req.Name,
		Place:     req.Place,
		Status:    req.Status,
		Pic:       req.Pic,
		Phone:     req.Phone,
		Principal: req.Principal,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdatePlaceResp{
		Code:    "200",
		Message: "更新用户成功",
	}, nil
}
