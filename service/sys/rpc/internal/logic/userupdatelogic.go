package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserUpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserUpdateLogic {
	return &UserUpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserUpdateLogic) UserUpdate(in *sys.UserUpdateReq) (*sys.UserUpdateResp, error) {
	err := l.svcCtx.UserModel.UpdateUser(in.Id, &model.User{
		Username: in.Username,
		Phone:    in.Phone,
		Nickname: in.Nickname,
		Gender:   in.Gender,
		Email:    in.Email,
		Status:   in.Status,
	})
	if err != nil {
		return nil, errors.New("更新用户失败")
	}
	if err != nil {
		return nil, errors.New("删除用户角色失败")
	}
	userrole, err := l.svcCtx.UserRoleModel.GetUserRoleByUserID(in.Id)
	if in.RoleId != int64(userrole.RoleID) {
		if in.RoleId == 0 {
			err = l.svcCtx.UserRoleModel.DeleteByUserID(in.Id)
		} else {
			err = l.svcCtx.UserRoleModel.DeleteByUserID(in.Id)
			err = l.svcCtx.UserRoleModel.AddUserRole(&model.UserRole{
				UserID:     int(in.Id),
				RoleID:     int(in.RoleId),
				LastEditBy: in.LastEditBy,
			})
		}
	}
	if err != nil {
		return nil, errors.New("修改用户角色失败")
	}

	return &sys.UserUpdateResp{}, nil
}
