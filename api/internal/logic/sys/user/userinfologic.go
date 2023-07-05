package user

import (
	"context"
	"encoding/json"
	"epick-mall/common/errorx"
	"epick-mall/service/sys/rpc/sysclient"
	"strconv"
	"strings"

	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (*types.UserInfoResp, error) {
	userId, _ := l.ctx.Value("id").(json.Number).Int64()
	resp, err := l.svcCtx.Sys.UserInfo(l.ctx, &sysclient.InfoReq{
		UserId: userId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError("查询用户失败")
	}
	var MenuTree []*types.ListMenuTree
	//组装ant ui中的菜单
	for _, item := range resp.MenuListTree {
		MenuTree = append(MenuTree, &types.ListMenuTree{
			Id:       item.Id,
			Path:     item.Path,
			Name:     item.Name,
			ParentId: item.ParentId,
			Icon:     item.Icon,
		})
	}
	//组装element ui中的菜单
	var MenuTreeVue []*types.ListMenuTreeVue
	for _, item := range resp.MenuListTree {
		if len(strings.TrimSpace(item.VuePath)) != 0 {
			MenuTreeVue = append(MenuTreeVue, &types.ListMenuTreeVue{
				Id:           item.Id,
				ParentId:     item.ParentId,
				Title:        item.Name,
				Path:         item.VuePath,
				Name:         item.Name,
				Icon:         item.VueIcon,
				VueRedirect:  item.VueRedirect,
				VueComponent: item.VueComponent,
				Meta: types.MenuTreeMeta{
					Title: item.Name,
					Icon:  item.VueIcon,
				},
			})
		}
	}
	//把能访问的url存在在redis，在middleware中检验
	err = l.svcCtx.Redis.Set(strconv.FormatInt(userId, 10), strings.Join(resp.BackgroundUrls, ","))
	if err != nil {
		logx.Errorf("设置用户：%s,权限到redis异常: %+v", resp.Nickname, err)
	}
	return &types.UserInfoResp{
		Code:        "000000",
		Message:     "获取个人信息成功",
		Avatar:      resp.Avatar,
		Name:        resp.Nickname,
		MenuTree:    MenuTree,
		MenuTreeVue: MenuTreeVue,
	}, nil
}
