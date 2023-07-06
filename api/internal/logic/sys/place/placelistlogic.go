package place

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PlaceListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPlaceListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PlaceListLogic {
	return &PlaceListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PlaceListLogic) PlaceList(req *types.ListPlaceReq) (*types.ListPlaceResp, error) {
	resp, err := l.svcCtx.Sys.PlaceList(l.ctx, &sysclient.PlaceListReq{
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Name:     req.Name,
		Place:    req.Place,
		Phone:    req.Phone,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListPlaceData
	for _, item := range resp.List {
		listUserData := types.ListPlaceData{
			Id:        item.Id,
			Name:      item.Name,
			Place:     item.Place,
			Status:    item.Status,
			Pic:       item.Pic,
			Phone:     item.Phone,
			Principal: item.Principal,
		}
		list = append(list, &listUserData)
	}

	return &types.ListPlaceResp{
		Code:     "000000",
		Message:  "查询角色列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
