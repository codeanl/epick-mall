package user

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"

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

func (l *UserLoginLogic) UserLogin(req *types.LoginReq, ip string) (*types.LoginResp, error) {
	if len(req.Phone) != 11 {
		return nil, errorx.NewCodeError(400, "手机号不为11位数")
	}
	if len(req.Password) < 6 {
		return nil, errorx.NewCodeError(400, "密码字数不够")
	}
	if len(req.Phone) == 0 || len(req.Password) == 0 {
		return nil, errorx.NewCodeError(400, "用户名或密码不能为空")
	}
	resp, err := l.svcCtx.System.Login(l.ctx, &systemclient.LoginReq{
		Phone:    req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, errorx.NewCodeError(400, "查询用户异常")
	}
	//保存登录日志
	_, _ = l.svcCtx.System.LoginLogAdd(l.ctx, &systemclient.LoginLogAddReq{
		Phone:  resp.Phone,
		Status: "login",
		Ip:     ip,
	})
	return &types.LoginResp{
		Code:         "200",
		Message:      "登录成功",
		ID:           resp.Id,
		Username:     resp.Username,
		Phone:        resp.Phone,
		Nickname:     resp.Nickname,
		Gender:       resp.Gender,
		Avatar:       resp.Avatar,
		Email:        resp.Email,
		Status:       resp.Status,
		AccessToken:  resp.AccessToken,
		AccessExpire: resp.AccessExpire,
		RefreshAfter: resp.RefreshAfter,
	}, nil
}
