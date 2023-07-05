package logic

import (
	"context"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogListLogic {
	return &LoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogListLogic) LoginLogList(in *sys.LoginLogListReq) (*sys.LoginLogListResp, error) {
	all, total, err := l.svcCtx.LoginLogModel.GetLoginLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.LoginLogListData
	for _, loginlog := range all {
		userinfo, _, _ := l.svcCtx.UserModel.GetUserByUsername(loginlog.Username)
		list = append(list, &sys.LoginLogListData{
			Id:         int64(loginlog.ID),
			Username:   loginlog.Username,
			CreateTime: loginlog.CreatedAt.Format("2006-01-02 15:04:05"),
			Ip:         loginlog.IP,
			Status:     loginlog.Status,
			Avatar:     userinfo.Avatar,
		})
	}
	return &sys.LoginLogListResp{
		Total: total,
		List:  list,
	}, nil
}
