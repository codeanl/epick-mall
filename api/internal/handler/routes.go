// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	systemlog "epick-mall/api/internal/handler/system/log"
	systemmenu "epick-mall/api/internal/handler/system/menu"
	systemrole "epick-mall/api/internal/handler/system/role"
	systemuser "epick-mall/api/internal/handler/system/user"
	"epick-mall/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/system/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/currentUser",
				Handler: systemuser.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: systemuser.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: systemuser.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: systemuser.UserDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: systemuser.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/resetpassword",
				Handler: systemuser.ResetPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/user"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: systemrole.RoleAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: systemrole.RoleUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: systemrole.RoleDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: systemrole.RoleListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/queryMenuByRoleId",
				Handler: systemrole.QueryMenuByRoleIdHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/updateRoleMenu",
				Handler: systemrole.UpdateRoleMenuHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/role"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: systemmenu.MenuAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/list",
				Handler: systemmenu.MenuListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: systemmenu.MenuUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: systemmenu.MenuDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/menu"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: systemlog.LoginLogListHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/delete",
				Handler: systemlog.LoginLogDeleteHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/system/loginLog"),
	)
}
