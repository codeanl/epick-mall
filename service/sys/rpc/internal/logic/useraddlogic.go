package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *UserAddLogic) UserAdd(in *sys.UserAddReq) (*sys.UserAddResp, error) {
	_, exist, _ := l.svcCtx.UserModel.GetUserByUsername(in.Username)
	if exist {
		return nil, errors.New("账户已存在")
	}
	user := &model.User{
		Username: in.Phone,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Password: in.Password,
		Gender:   in.Gender,
		Avatar:   in.Avatar,
		Email:    in.Email,
		Status:   in.Status,
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
	return &sys.UserAddResp{}, nil
}
