package logic

import (
	"context"
	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserListLogic) UserList(in *system.UserListReq) (*system.UserListResp, error) {
	all, total, err := l.svcCtx.UserModel.GetUserList(in)
	if err != nil {
		return nil, err
	}
	var list []*system.UserList
	for _, user := range all {
		list = append(list, &system.UserList{
			Id:         int64(user.ID),
			Username:   user.Username,
			Phone:      user.Phone,
			Nickname:   user.Nickname,
			Gender:     int64(user.Gender),
			Avatar:     user.Avatar,
			Email:      user.Email,
			Status:     int64(user.Status),
			CreatTime:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return &system.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
