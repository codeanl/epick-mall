package logic

import (
	"context"
	"epick-mall/service/sys/model"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"epick-mall/service/sys/rpc/internal/svc"
	"epick-mall/service/sys/rpc/sys"

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

func (l *UserInfoLogic) UserInfo(in *sys.InfoReq) (*sys.InfoResp, error) {
	userInfo, err := l.svcCtx.UserModel.GetUserByID(in.UserId)
	exist, err := l.svcCtx.UserRoleModel.IsSuperAdmin(in.UserId)

	if err != nil {
		return nil, errors.New(fmt.Sprintf("用户不存在userId: %s", strconv.FormatInt(in.UserId, 10)))
	}
	var list []*sys.MenuListTree
	var listUrls []string
	if exist {
		menulist, _, _ := l.svcCtx.MenuModel.GetMenuList(&sys.MenuListReq{Name: "", Url: ""})
		list, listUrls = listTrees(menulist, list, listUrls)
	} else {
		menus, _ := l.svcCtx.MenuModel.GetMenusByUserID(in.UserId)
		list, listUrls = listTrees(menus, list, listUrls)
	}
	return &sys.InfoResp{
		Avatar:         userInfo.Avatar,
		Nickname:       userInfo.Nickname,
		Username:       userInfo.Username,
		MenuListTree:   list,
		BackgroundUrls: listUrls,
	}, nil
}

func listTrees(menus []model.Menu, list []*sys.MenuListTree, listUrls []string) ([]*sys.MenuListTree, []string) {
	for _, menu := range menus {
		if menu.Type == 1 || menu.Type == 0 {
			list = append(list, &sys.MenuListTree{
				Id:           int64(menu.ID),
				Name:         menu.Name,
				Icon:         menu.Icon,
				ParentId:     int64(menu.ParentID),
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
