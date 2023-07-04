package logic

import (
	"context"
	"epick-mall/service/system/model"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ResetPasswordLogic) ResetPassword(in *system.ResetPasswordReq) (*system.ResetPasswordResp, error) {
	err := l.svcCtx.UserModel.UpdateUser(in.Id, &model.User{
		Password: in.Password,
	})
	if err != nil {
		return nil, errors.New("重置密码失败")
	}
	return &system.ResetPasswordResp{}, nil
}
