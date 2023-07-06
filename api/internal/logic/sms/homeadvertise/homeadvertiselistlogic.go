package homeadvertise

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/sms/rpc/smsclient"
	"github.com/zeromicro/go-zero/core/logx"
	"log"
)

type HomeAdvertiseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(req *types.ListHomeAdvertiseReq) (*types.ListHomeAdvertiseResp, error) {
	resp, err := l.svcCtx.Sms.HomeAdvertiseList(l.ctx, &smsclient.HomeAdvertiseListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
		Name:     req.Name,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.ListHomeAdvertiseData
	for _, item := range resp.List {
		listUserData := types.ListHomeAdvertiseData{
			Id:        item.Id,
			Name:      item.Name,
			Type:      item.Type,
			Pic:       item.Pic,
			StartTime: item.StartTime,
			EndTime:   item.EndTime,
			Status:    item.Status,
			Url:       item.Url,
			Note:      item.Note,
			Sort:      item.Sort,
		}
		list = append(list, &listUserData)
	}
	log.Print(list)
	return &types.ListHomeAdvertiseResp{
		Code:     "000000",
		Message:  "查询角色列表成功",
		Total:    resp.Total,
		Data:     list,
		Current:  req.Current,
		PageSize: req.PageSize,
	}, nil
}
