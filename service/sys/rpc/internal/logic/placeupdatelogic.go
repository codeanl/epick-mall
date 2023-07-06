package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceUpdateLogic {
	return &PlaceUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceUpdateLogic) PlaceUpdate(in *sys.PlaceUpdateReq) (*sys.PlaceUpdateResp, error) {
	err := l.svcCtx.PlaceModel.UpdatePlace(in.Id, &model.Place{
		Name:      in.Name,
		Place:     in.Place,
		Status:    in.Status,
		Pic:       in.Pic,
		Phone:     in.Phone,
		Principal: in.Principal,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	return &sys.PlaceUpdateResp{}, nil
}
