package user

import (
	"context"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"epick-mall/common/errorx"
	"epick-mall/service/system/rpc/systemclient"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.ListUserReq) (*types.ListUserResp, error) {
	resp, err := l.svcCtx.System.UserList(l.ctx, &systemclient.UserListReq{
		Nickname: req.Nickname,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
		Gender:   req.Gander,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询失败")
	}
	var list []*types.User
	for _, item := range resp.List {
		listUserData := types.User{
			ID:         item.Id,
			Username:   item.Username,
			NickName:   item.Nickname,
			Phone:      item.Phone,
			Gander:     item.Gender,
			Avatar:     item.Avatar,
			Email:      item.Email,
			Status:     item.Status,
			CreateTime: item.CreatTime,
			UpdateTime: item.UpdateTime,
		}
		list = append(list, &listUserData)
	}
	return &types.ListUserResp{
		Code:     "000000",
		Message:  "查询用户列表成功",
		Total:    resp.Total,
		Data:     list,
		PageNum:  req.PageNum,
		PageSize: req.PageSize,
	}, nil
}
