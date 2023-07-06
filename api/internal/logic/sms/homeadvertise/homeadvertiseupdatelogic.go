package homeadvertise

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/sms/rpc/smsclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type HomeAdvertiseUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomeAdvertiseUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomeAdvertiseUpdateLogic {
	return &HomeAdvertiseUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomeAdvertiseUpdateLogic) HomeAdvertiseUpdate(req *types.UpdateHomeAdvertiseReq) (*types.UpdateHomeAdvertiseResp, error) {
	_, err := l.svcCtx.Sms.HomeAdvertiseUpdate(l.ctx, &smsclient.HomeAdvertiseUpdateReq{
		Id:        req.Id,
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
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateHomeAdvertiseResp{
		Code:    "200",
		Message: "更新用户成功",
	}, nil
}
