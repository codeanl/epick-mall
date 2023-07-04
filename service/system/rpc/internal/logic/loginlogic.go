package logic

import (
	"context"
	"epick-mall/common"
	"errors"
	"time"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *system.LoginReq) (*system.LoginResp, error) {
	userInfo, exist, _ := l.svcCtx.UserModel.GetUserByPhone(in.Phone)
	if !exist {
		return nil, errors.New("用户不存在")
	}
	if userInfo.Password != in.Password {
		return nil, errors.New("用户密码不正确")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	AccessToken, err := common.GetJwtToken(accessSecret, now, accessExpire, int64(userInfo.ID), userInfo.Phone, userInfo.Nickname)
	if err != nil {
		return nil, err
	}

	resp := &system.LoginResp{
		Id:           int64(userInfo.ID),
		Nickname:     userInfo.Nickname,
		Username:     userInfo.Username,
		Email:        userInfo.Email,
		Avatar:       userInfo.Avatar,
		Phone:        userInfo.Phone,
		Gender:       int64(userInfo.Gender),
		Status:       int64(userInfo.Status),
		AccessToken:  AccessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}
	return resp, nil
}
