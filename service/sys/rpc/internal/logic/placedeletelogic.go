package logic

import (
	"context"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceDeleteLogic {
	return &PlaceDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceDeleteLogic) PlaceDelete(in *sys.PlaceDeleteReq) (*sys.PlaceDeleteResp, error) {
	err := l.svcCtx.PlaceModel.DeletePlaceByIds(in.Ids)
	if err != nil {
		return nil, errors.New("删除失败")
	}
	return &sys.PlaceDeleteResp{}, nil
}
