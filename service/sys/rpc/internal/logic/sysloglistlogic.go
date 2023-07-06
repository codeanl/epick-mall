package logic

import (
	"context"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

	"github.com/zeromicro/go-zero/core/logx"
)

type SysLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSysLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SysLogListLogic {
	return &SysLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SysLogListLogic) SysLogList(in *sys.SysLogListReq) (*sys.SysLogListResp, error) {
	all, total, err := l.svcCtx.LogModel.GetLogList(in)
	if err != nil {
		return nil, err
	}
	var list []*sys.SysLogListData
	for _, log := range all {
		//userinfo, _, _ := l.svcCtx.UserModel.GetUserByUsername(log.Username)
		list = append(list, &sys.SysLogListData{
			Id:         int64(log.ID),
			Username:   log.Username,
			Method:     log.Method,
			Operation:  log.Operation,
			Params:     log.Params,
			Time:       log.Time,
			CreateTime: log.CreatedAt.Format("2006-01-02 15:04:05"),
			Ip:         log.IP,
			//Avatar:     userinfo.Avatar,
		})
	}
	return &sys.SysLogListResp{
		Total: total,
		List:  list,
	}, nil
}
