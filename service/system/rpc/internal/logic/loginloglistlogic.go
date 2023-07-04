package logic

import (
	"context"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

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

func (l *LoginLogListLogic) LoginLogList(in *system.LoginLogListReq) (*system.LoginLogListResp, error) {
	all, total, err := l.svcCtx.LoginLogModel.GetLoginLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*system.LoginLogListData
	for _, loginlog := range all {
		list = append(list, &system.LoginLogListData{
			Id:         int64(loginlog.ID),
			Phone:      loginlog.Phone,
			Status:     loginlog.Status,
			CreateTime: loginlog.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return &system.LoginLogListResp{
		Total: total,
		List:  list,
	}, nil
}
