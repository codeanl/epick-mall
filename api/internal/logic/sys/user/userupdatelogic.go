package user

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserUpdateLogic) UserUpdate(req *types.UpdateUserReq) (resp *types.UpdateUserResp, err error) {
	_, err = l.svcCtx.Sys.UserUpdate(l.ctx, &sysclient.UserUpdateReq{
		Id:         req.ID,
		Username:   req.Username,
		Phone:      req.Phone,
		Nickname:   req.Nickname,
		Gender:     req.Gender,
		Avatar:     req.Avatar,
		Email:      req.Email,
		LastEditBy: l.ctx.Value("nickname").(string),
		RoleId:     req.RoleID,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("更新用户失败")
	}
	return &types.UpdateUserResp{
		Code:    "200",
		Message: "更新用户成功",
	}, nil
}
