package user

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) (*types.ResetPasswordResp, error) {
	_, err := l.svcCtx.System.ResetPassword(l.ctx, &systemclient.ResetPasswordReq{
		Id:       req.ID,
		Password: "123456",
	})
	if err != nil {
		return nil, errorx.NewDefaultError("重置密码失败")
	}
	return &types.ResetPasswordResp{
		Code:    "200",
		Message: "重置密码成功",
	}, nil
}
