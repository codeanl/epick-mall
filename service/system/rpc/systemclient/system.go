// Code generated by goctl. DO NOT EDIT.
// Source: system.proto

package systemclient

import (
	"context"

	"epick-mall/service/system/rpc/system"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	InfoReq               = system.InfoReq
	InfoResp              = system.InfoResp
	LogAddReq             = system.LogAddReq
	LogAddResp            = system.LogAddResp
	LoginLogAddReq        = system.LoginLogAddReq
	LoginLogAddResp       = system.LoginLogAddResp
	LoginLogDeleteReq     = system.LoginLogDeleteReq
	LoginLogDeleteResp    = system.LoginLogDeleteResp
	LoginLogListData      = system.LoginLogListData
	LoginLogListReq       = system.LoginLogListReq
	LoginLogListResp      = system.LoginLogListResp
	LoginReq              = system.LoginReq
	LoginResp             = system.LoginResp
	MenuAddReq            = system.MenuAddReq
	MenuAddResp           = system.MenuAddResp
	MenuDeleteReq         = system.MenuDeleteReq
	MenuDeleteResp        = system.MenuDeleteResp
	MenuListData          = system.MenuListData
	MenuListReq           = system.MenuListReq
	MenuListResp          = system.MenuListResp
	MenuListTree          = system.MenuListTree
	MenuUpdateReq         = system.MenuUpdateReq
	MenuUpdateResp        = system.MenuUpdateResp
	QueryMenuByRoleIdReq  = system.QueryMenuByRoleIdReq
	QueryMenuByRoleIdResp = system.QueryMenuByRoleIdResp
	ResetPasswordReq      = system.ResetPasswordReq
	ResetPasswordResp     = system.ResetPasswordResp
	RoleAddReq            = system.RoleAddReq
	RoleAddResp           = system.RoleAddResp
	RoleDeleteReq         = system.RoleDeleteReq
	RoleDeleteResp        = system.RoleDeleteResp
	RoleListData          = system.RoleListData
	RoleListReq           = system.RoleListReq
	RoleListResp          = system.RoleListResp
	RoleUpdateReq         = system.RoleUpdateReq
	RoleUpdateResp        = system.RoleUpdateResp
	UpdateMenuRoleReq     = system.UpdateMenuRoleReq
	UpdateMenuRoleResp    = system.UpdateMenuRoleResp
	UserAddReq            = system.UserAddReq
	UserAddResp           = system.UserAddResp
	UserDeleteReq         = system.UserDeleteReq
	UserDeleteResp        = system.UserDeleteResp
	UserList              = system.UserList
	UserListReq           = system.UserListReq
	UserListResp          = system.UserListResp
	UserUpdateReq         = system.UserUpdateReq
	UserUpdateResp        = system.UserUpdateResp

	System interface {
		Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
		UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error)
		UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error)
		UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error)
		UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error)
		UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error)
		ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error)
		RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error)
		RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error)
		RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error)
		RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error)
		QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error)
		UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error)
		MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error)
		MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error)
		MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error)
		MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error)
		LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error)
		LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error)
		LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error)
		LogAdd(ctx context.Context, in *LogAddReq, opts ...grpc.CallOption) (*LogAddResp, error)
	}

	defaultSystem struct {
		cli zrpc.Client
	}
)

func NewSystem(cli zrpc.Client) System {
	return &defaultSystem{
		cli: cli,
	}
}

func (m *defaultSystem) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultSystem) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}

func (m *defaultSystem) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UserAdd(ctx, in, opts...)
}

func (m *defaultSystem) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UserUpdate(ctx, in, opts...)
}

func (m *defaultSystem) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UserDelete(ctx, in, opts...)
}

func (m *defaultSystem) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UserList(ctx, in, opts...)
}

func (m *defaultSystem) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.ResetPassword(ctx, in, opts...)
}

func (m *defaultSystem) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.RoleAdd(ctx, in, opts...)
}

func (m *defaultSystem) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.RoleUpdate(ctx, in, opts...)
}

func (m *defaultSystem) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.RoleDelete(ctx, in, opts...)
}

func (m *defaultSystem) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.RoleList(ctx, in, opts...)
}

func (m *defaultSystem) QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.QueryMenuByRoleId(ctx, in, opts...)
}

func (m *defaultSystem) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.UpdateMenuRole(ctx, in, opts...)
}

func (m *defaultSystem) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.MenuAdd(ctx, in, opts...)
}

func (m *defaultSystem) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.MenuList(ctx, in, opts...)
}

func (m *defaultSystem) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.MenuUpdate(ctx, in, opts...)
}

func (m *defaultSystem) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.MenuDelete(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.LoginLogAdd(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.LoginLogList(ctx, in, opts...)
}

func (m *defaultSystem) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.LoginLogDelete(ctx, in, opts...)
}

func (m *defaultSystem) LogAdd(ctx context.Context, in *LogAddReq, opts ...grpc.CallOption) (*LogAddResp, error) {
	client := system.NewSystemClient(m.cli.Conn())
	return client.LogAdd(ctx, in, opts...)
}