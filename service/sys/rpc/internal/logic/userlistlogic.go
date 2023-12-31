package logic

import (
	"context"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *UserListLogic) UserList(in *sys.UserListReq) (*sys.UserListResp, error) {
	all, total, err := l.svcCtx.UserModel.GetUserList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.UserList
	for _, user := range all {
		roleinfo, _ := l.svcCtx.RoleModel.GetRoleInfoByUserID(int64(user.ID))
		list = append(list, &sys.UserList{
			Id:         int64(user.ID),
			Username:   user.Username,
			Phone:      user.Phone,
			Nickname:   user.Nickname,
			Gender:     user.Gender,
			Avatar:     user.Avatar,
			Email:      user.Email,
			Status:     user.Status,
			CreatTime:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateTime: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			RoleId:     int64(roleinfo.ID),
			RoleName:   roleinfo.Name,
		})
	}
	return &sys.UserListResp{
		Total: total,
		List:  list,
	}, nil
}
