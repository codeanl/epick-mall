package user

import (
	"context"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserAddLogic) UserAdd(req *types.AddUserReq) (resp *types.AddUserResp, err error) {
	_, err = l.svcCtx.Sys.UserAdd(l.ctx, &sysclient.UserAddReq{
		Phone:      req.Phone,
		Nickname:   req.Nickname,
		Password:   "123456",
		Gender:     req.Gender,
		Avatar:     "https://an23.co/upload/2023/04/avatar.jpg",
		Email:      req.Email,
		Status:     req.Status,
		LastEditBy: l.ctx.Value("nickname").(string),
		RoleId:     req.RoleID,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddUserResp{
		Code:    "200",
		Message: "添加用户成功",
	}, nil
}
