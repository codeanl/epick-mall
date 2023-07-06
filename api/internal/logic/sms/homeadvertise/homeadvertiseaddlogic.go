package homeadvertise

import (
	"context"
	"epick-mall/service/sms/rpc/smsclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseAddLogic {
	return &HomeAdvertiseAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseAddLogic) HomeAdvertiseAdd(req *types.AddHomeAdvertiseReq) (*types.AddHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseAdd(l.ctx, &smsclient.HomeAdvertiseAddReq{
		Name:      req.Name,
		Type:      req.Type,
		Pic:       "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20181113/movie_ad.jpg", //暂时没有上传,用这个当默认
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    req.Status,
		Url:       req.Url,
		Note:      req.Note,
		Sort:      req.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddHomeAdvertiseResp{
		Code:    "200",
		Message: "添加成功",
	}, nil
}
