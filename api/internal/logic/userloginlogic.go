package logic

import (
	"context"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLoginLogic) UserLogin(req *types.LoginReq) (*types.LoginResp, error) {
	if len(req.Username) == 0 || len(req.Password) == 0 {
		return nil, errorx.NewCodeError(400, "用户名或密码不能为空")
	}
	resp, err := l.svcCtx.Sys.Login(l.ctx, &sysclient.LoginReq{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewCodeError(400, "查询用户异常")
	}
	return &types.LoginResp{
		Code:             "200",
		Message:          "登录成功",
		ID:               resp.Id,
		Username:         resp.Username,
		Phone:            resp.Phone,
		Nickname:         resp.Nickname,
		Gender:           resp.Gender,
		Avatar:           resp.Avatar,
		Email:            resp.Email,
		Status:           resp.Status,
		CurrentAuthority: resp.CurrentAuthority,
		AccessToken:      resp.AccessToken,
		AccessExpire:     resp.AccessExpire,
		RefreshAfter:     resp.RefreshAfter,
	}, nil
}
