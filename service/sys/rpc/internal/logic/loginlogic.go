package logic

import (
	"context"
	"epick-mall/common"
	"errors"
	"time"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *LoginLogic) Login(in *sys.LoginReq) (*sys.LoginResp, error) {
	userInfo, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if !exist {
		return nil, errors.New("用户不存在")
	}
	if userInfo.Password != in.Password {
		return nil, errors.New("用户密码不正确")
	}
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	accessSecret := l.svcCtx.Config.JWT.AccessSecret
	AccessToken, err := common.GetJwtToken(accessSecret, now, accessExpire, int64(userInfo.ID), userInfo.Username, userInfo.Nickname)
	if err != nil {
		return nil, err
	}
	resp := &sys.LoginResp{
		Id:               int64(userInfo.ID),
		Nickname:         userInfo.Nickname,
		Username:         userInfo.Username,
		Email:            userInfo.Email,
		Avatar:           userInfo.Avatar,
		Phone:            userInfo.Phone,
		CurrentAuthority: "admin",
		Gender:           userInfo.Gender,
		Status:           "ok",
		AccessToken:      AccessToken,
		AccessExpire:     now + accessExpire,
		RefreshAfter:     now + accessExpire/2,
	}
	return resp, nil
}
