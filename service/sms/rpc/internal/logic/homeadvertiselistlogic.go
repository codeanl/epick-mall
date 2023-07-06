package logic

import (
	"context"
	"epick-mall/service/sms/rpc/internal/svc"
	"epick-mall/service/sms/rpc/sms"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewHomeAdvertiseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseListLogic {
	return &HomeAdvertiseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *HomeAdvertiseListLogic) HomeAdvertiseList(in *sms.HomeAdvertiseListReq) (*sms.HomeAdvertiseListResp, error) {
	all, total, err := l.svcCtx.HomeAdvertiseModel.GetHomeAdvertiseList(in)
	if err != nil {
		return nil, err
	}
	var list []*sms.HomeAdvertiseListData
	for _, role := range all {
		list = append(list, &sms.HomeAdvertiseListData{
			Id:         int64(role.ID),
			Name:       role.Name,
			Type:       role.Type,
			Pic:        role.Pic,
			StartTime:  role.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:    role.EndTime.Format("2006-01-02 15:04:05"),
			Status:     role.Status,
			ClickCount: role.ClickCount,
			OrderCount: role.OrderCount,
			Url:        role.Url,
			Note:       role.Note,
			Sort:       role.Sort,
		})
	}
	print(list)
	return &sms.HomeAdvertiseListResp{
		Total: total,
		List:  list,
	}, nil

}
