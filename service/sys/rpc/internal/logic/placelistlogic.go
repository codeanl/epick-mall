package logic

import (
	"context"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PlaceListLogic) PlaceList(in *sys.PlaceListReq) (*sys.PlaceListResp, error) {
	all, total, err := l.svcCtx.PlaceModel.GetPlaceList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.PlaceListData
	for _, role := range all {
		list = append(list, &sys.PlaceListData{
			Id:        int64(role.ID),
			Name:      role.Name,
			Place:     role.Place,
			Status:    role.Status,
			Pic:       role.Pic,
			Phone:     role.Phone,
			Principal: role.Principal,
		})
	}
	return &sys.PlaceListResp{
		Total: total,
		List:  list,
	}, nil
}
