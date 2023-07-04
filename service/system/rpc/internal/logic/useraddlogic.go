package logic

import (
	"context"
	"epick-mall/service/system/model"
	"errors"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserAddLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserAddLogic {
	return &UserAddLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserAddLogic) UserAdd(in *system.UserAddReq) (*system.UserAddResp, error) {
	_, exist, _ := l.svcCtx.UserModel.GetUserByPhone(in.Phone)
	if exist {
		return nil, errors.New("账户已存在")
	}
	user := &model.User{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Password: in.Password,
		Gender:   int(in.Gender),
		Avatar:   in.Avatar,
		Email:    in.Email,
		Status:   int(in.Status),
	}
	err := l.svcCtx.UserModel.AddUser(user)
	if err != nil {
		return nil, errors.New("添加用户失败")
	}
	if in.RoleId != 0 {
		err = l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
			UserID:     int(user.ID),
			RoleID:     int(in.RoleId),
			LastEditBy: in.LastEditBy,
		})
		if err != nil {
			return nil, errors.New("添加用户角色失败")
		}
	}
	return &system.UserAddResp{}, nil
}
