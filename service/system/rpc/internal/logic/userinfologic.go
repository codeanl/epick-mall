package logic

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"epick-mall/service/system/rpc/internal/svc"
	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *system.InfoReq) (*system.InfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.GetUserByID(in.UserId)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(in.UserId, 10)))
	}
	var list []*system.MenuListTree
	var listUrls []string
	if in.UserId == 1 {
		menulist, _, _ := l.svcCtx.MenuModel.GetMenuList(&system.MenuListReq{Name: "", Url: ""})
		var menus []system.MenuListData
		for _, menu := range menulist {
			menus = append(menus, system.MenuListData{
				Id:             int64(menu.ID),
				Name:           menu.Name,
				ParentId:       int64(menu.ParentID),
				Url:            menu.Url,
				Perms:          menu.Perms,
				Type:           int64(menu.Type),
				Icon:           menu.Icon,
				OrderNum:       int64(menu.OrderNum),
				CreateBy:       menu.CreateBy,
				CreateTime:     menu.CreatedAt.Format("2006-01-02 15:04:05"),
				LastUpdateBy:   menu.UpdateBy,
				LastUpdateTime: menu.UpdatedAt.Format("2006-01-02 15:04:05"),
				BackgroundUrl:  menu.BackgroundUrl,
				VuePath:        menu.VuePath,
				VueComponent:   menu.VueComponent,
				VueIcon:        menu.VueIcon,
				VueRedirect:    menu.VueRedirect,
			})
		}
		list, listUrls = listTrees(&menus, list, listUrls)
	} else {
		menus, _ := l.svcCtx.MenuModel.GetMenusByUserID(in.UserId)
		list, listUrls = listTrees(&menus, list, listUrls)
	}

	return &system.InfoResp{
		Avatar:         userInfo.Avatar,
		Nickname:       userInfo.Nickname,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}, nil
}

func listTrees(menus *[]system.MenuListData, list []*system.MenuListTree, listUrls []string) ([]*system.MenuListTree, []string) {
	for _, menu := range *menus {
		if menu.Type == 1 || menu.Type == 0 {
			list = append(list, &system.MenuListTree{
				Id:           menu.Id,
				Name:         menu.Name,
				Icon:         menu.Icon,
				ParentId:     menu.ParentId,
				Path:         menu.Url,
				VuePath:      menu.VuePath,
				VueComponent: menu.VueComponent,
				VueIcon:      menu.VueIcon,
				VueRedirect:  menu.VueRedirect,
			})
		}
		if len(strings.TrimSpace(menu.BackgroundUrl)) != 0 {
			listUrls = append(listUrls, menu.BackgroundUrl)
		}
	}
	return list, listUrls
}
