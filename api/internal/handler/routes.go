// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	sysuser "epick-mall/api/internal/handler/sys/user"
	"epick-mall/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/sys/user/login",
				Handler: UserLoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/currentUser",
				Handler: sysuser.UserInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/add",
				Handler: sysuser.UserAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/update",
				Handler: sysuser.UserUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/delete",
				Handler: sysuser.UserDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/list",
				Handler: sysuser.UserListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/resetpassword",
				Handler: sysuser.ResetPasswordHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/sys/user"),
	)
}
