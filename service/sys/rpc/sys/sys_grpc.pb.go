// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: sys.proto

package sys

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Sys_Login_FullMethodName             = "/sys.Sys/Login"
	Sys_UserInfo_FullMethodName          = "/sys.Sys/UserInfo"
	Sys_UserAdd_FullMethodName           = "/sys.Sys/UserAdd"
	Sys_UserUpdate_FullMethodName        = "/sys.Sys/UserUpdate"
	Sys_UserDelete_FullMethodName        = "/sys.Sys/UserDelete"
	Sys_UserList_FullMethodName          = "/sys.Sys/UserList"
	Sys_ResetPassword_FullMethodName     = "/sys.Sys/ResetPassword"
	Sys_RoleAdd_FullMethodName           = "/sys.Sys/RoleAdd"
	Sys_RoleUpdate_FullMethodName        = "/sys.Sys/RoleUpdate"
	Sys_RoleDelete_FullMethodName        = "/sys.Sys/RoleDelete"
	Sys_RoleList_FullMethodName          = "/sys.Sys/RoleList"
	Sys_QueryMenuByRoleId_FullMethodName = "/sys.Sys/QueryMenuByRoleId"
	Sys_UpdateMenuRole_FullMethodName    = "/sys.Sys/UpdateMenuRole"
	Sys_MenuAdd_FullMethodName           = "/sys.Sys/MenuAdd"
	Sys_MenuList_FullMethodName          = "/sys.Sys/MenuList"
	Sys_MenuUpdate_FullMethodName        = "/sys.Sys/MenuUpdate"
	Sys_MenuDelete_FullMethodName        = "/sys.Sys/MenuDelete"
	Sys_LoginLogAdd_FullMethodName       = "/sys.Sys/LoginLogAdd"
	Sys_LoginLogList_FullMethodName      = "/sys.Sys/LoginLogList"
	Sys_LoginLogDelete_FullMethodName    = "/sys.Sys/LoginLogDelete"
	Sys_LogAdd_FullMethodName            = "/sys.Sys/LogAdd"
	Sys_SysLogList_FullMethodName        = "/sys.Sys/SysLogList"
	Sys_SysLogDelete_FullMethodName      = "/sys.Sys/SysLogDelete"
	Sys_PlaceAdd_FullMethodName          = "/sys.Sys/PlaceAdd"
	Sys_PlaceList_FullMethodName         = "/sys.Sys/PlaceList"
	Sys_PlaceUpdate_FullMethodName       = "/sys.Sys/PlaceUpdate"
	Sys_PlaceDelete_FullMethodName       = "/sys.Sys/PlaceDelete"
)

// SysClient is the client API for Sys service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysClient interface {
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
	SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error)
	SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error)
	PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error)
	PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error)
	PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error)
	PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error)
}

type sysClient struct {
	cc grpc.ClientConnInterface
}

func NewSysClient(cc grpc.ClientConnInterface) SysClient {
	return &sysClient{cc}
}

func (c *sysClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, Sys_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserInfo(ctx context.Context, in *InfoReq, opts ...grpc.CallOption) (*InfoResp, error) {
	out := new(InfoResp)
	err := c.cc.Invoke(ctx, Sys_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserAdd(ctx context.Context, in *UserAddReq, opts ...grpc.CallOption) (*UserAddResp, error) {
	out := new(UserAddResp)
	err := c.cc.Invoke(ctx, Sys_UserAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserUpdate(ctx context.Context, in *UserUpdateReq, opts ...grpc.CallOption) (*UserUpdateResp, error) {
	out := new(UserUpdateResp)
	err := c.cc.Invoke(ctx, Sys_UserUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserDelete(ctx context.Context, in *UserDeleteReq, opts ...grpc.CallOption) (*UserDeleteResp, error) {
	out := new(UserDeleteResp)
	err := c.cc.Invoke(ctx, Sys_UserDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UserList(ctx context.Context, in *UserListReq, opts ...grpc.CallOption) (*UserListResp, error) {
	out := new(UserListResp)
	err := c.cc.Invoke(ctx, Sys_UserList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) ResetPassword(ctx context.Context, in *ResetPasswordReq, opts ...grpc.CallOption) (*ResetPasswordResp, error) {
	out := new(ResetPasswordResp)
	err := c.cc.Invoke(ctx, Sys_ResetPassword_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RoleAdd(ctx context.Context, in *RoleAddReq, opts ...grpc.CallOption) (*RoleAddResp, error) {
	out := new(RoleAddResp)
	err := c.cc.Invoke(ctx, Sys_RoleAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RoleUpdate(ctx context.Context, in *RoleUpdateReq, opts ...grpc.CallOption) (*RoleUpdateResp, error) {
	out := new(RoleUpdateResp)
	err := c.cc.Invoke(ctx, Sys_RoleUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RoleDelete(ctx context.Context, in *RoleDeleteReq, opts ...grpc.CallOption) (*RoleDeleteResp, error) {
	out := new(RoleDeleteResp)
	err := c.cc.Invoke(ctx, Sys_RoleDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) RoleList(ctx context.Context, in *RoleListReq, opts ...grpc.CallOption) (*RoleListResp, error) {
	out := new(RoleListResp)
	err := c.cc.Invoke(ctx, Sys_RoleList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) QueryMenuByRoleId(ctx context.Context, in *QueryMenuByRoleIdReq, opts ...grpc.CallOption) (*QueryMenuByRoleIdResp, error) {
	out := new(QueryMenuByRoleIdResp)
	err := c.cc.Invoke(ctx, Sys_QueryMenuByRoleId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) UpdateMenuRole(ctx context.Context, in *UpdateMenuRoleReq, opts ...grpc.CallOption) (*UpdateMenuRoleResp, error) {
	out := new(UpdateMenuRoleResp)
	err := c.cc.Invoke(ctx, Sys_UpdateMenuRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuAdd(ctx context.Context, in *MenuAddReq, opts ...grpc.CallOption) (*MenuAddResp, error) {
	out := new(MenuAddResp)
	err := c.cc.Invoke(ctx, Sys_MenuAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuList(ctx context.Context, in *MenuListReq, opts ...grpc.CallOption) (*MenuListResp, error) {
	out := new(MenuListResp)
	err := c.cc.Invoke(ctx, Sys_MenuList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuUpdate(ctx context.Context, in *MenuUpdateReq, opts ...grpc.CallOption) (*MenuUpdateResp, error) {
	out := new(MenuUpdateResp)
	err := c.cc.Invoke(ctx, Sys_MenuUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) MenuDelete(ctx context.Context, in *MenuDeleteReq, opts ...grpc.CallOption) (*MenuDeleteResp, error) {
	out := new(MenuDeleteResp)
	err := c.cc.Invoke(ctx, Sys_MenuDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogAdd(ctx context.Context, in *LoginLogAddReq, opts ...grpc.CallOption) (*LoginLogAddResp, error) {
	out := new(LoginLogAddResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogList(ctx context.Context, in *LoginLogListReq, opts ...grpc.CallOption) (*LoginLogListResp, error) {
	out := new(LoginLogListResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LoginLogDelete(ctx context.Context, in *LoginLogDeleteReq, opts ...grpc.CallOption) (*LoginLogDeleteResp, error) {
	out := new(LoginLogDeleteResp)
	err := c.cc.Invoke(ctx, Sys_LoginLogDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) LogAdd(ctx context.Context, in *LogAddReq, opts ...grpc.CallOption) (*LogAddResp, error) {
	out := new(LogAddResp)
	err := c.cc.Invoke(ctx, Sys_LogAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) SysLogList(ctx context.Context, in *SysLogListReq, opts ...grpc.CallOption) (*SysLogListResp, error) {
	out := new(SysLogListResp)
	err := c.cc.Invoke(ctx, Sys_SysLogList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) SysLogDelete(ctx context.Context, in *SysLogDeleteReq, opts ...grpc.CallOption) (*SysLogDeleteResp, error) {
	out := new(SysLogDeleteResp)
	err := c.cc.Invoke(ctx, Sys_SysLogDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PlaceAdd(ctx context.Context, in *PlaceAddReq, opts ...grpc.CallOption) (*PlaceAddResp, error) {
	out := new(PlaceAddResp)
	err := c.cc.Invoke(ctx, Sys_PlaceAdd_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PlaceList(ctx context.Context, in *PlaceListReq, opts ...grpc.CallOption) (*PlaceListResp, error) {
	out := new(PlaceListResp)
	err := c.cc.Invoke(ctx, Sys_PlaceList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PlaceUpdate(ctx context.Context, in *PlaceUpdateReq, opts ...grpc.CallOption) (*PlaceUpdateResp, error) {
	out := new(PlaceUpdateResp)
	err := c.cc.Invoke(ctx, Sys_PlaceUpdate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysClient) PlaceDelete(ctx context.Context, in *PlaceDeleteReq, opts ...grpc.CallOption) (*PlaceDeleteResp, error) {
	out := new(PlaceDeleteResp)
	err := c.cc.Invoke(ctx, Sys_PlaceDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysServer is the server API for Sys service.
// All implementations must embed UnimplementedSysServer
// for forward compatibility
type SysServer interface {
	Login(context.Context, *LoginReq) (*LoginResp, error)
	UserInfo(context.Context, *InfoReq) (*InfoResp, error)
	UserAdd(context.Context, *UserAddReq) (*UserAddResp, error)
	UserUpdate(context.Context, *UserUpdateReq) (*UserUpdateResp, error)
	UserDelete(context.Context, *UserDeleteReq) (*UserDeleteResp, error)
	UserList(context.Context, *UserListReq) (*UserListResp, error)
	ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordResp, error)
	RoleAdd(context.Context, *RoleAddReq) (*RoleAddResp, error)
	RoleUpdate(context.Context, *RoleUpdateReq) (*RoleUpdateResp, error)
	RoleDelete(context.Context, *RoleDeleteReq) (*RoleDeleteResp, error)
	RoleList(context.Context, *RoleListReq) (*RoleListResp, error)
	QueryMenuByRoleId(context.Context, *QueryMenuByRoleIdReq) (*QueryMenuByRoleIdResp, error)
	UpdateMenuRole(context.Context, *UpdateMenuRoleReq) (*UpdateMenuRoleResp, error)
	MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error)
	MenuList(context.Context, *MenuListReq) (*MenuListResp, error)
	MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error)
	MenuDelete(context.Context, *MenuDeleteReq) (*MenuDeleteResp, error)
	LoginLogAdd(context.Context, *LoginLogAddReq) (*LoginLogAddResp, error)
	LoginLogList(context.Context, *LoginLogListReq) (*LoginLogListResp, error)
	LoginLogDelete(context.Context, *LoginLogDeleteReq) (*LoginLogDeleteResp, error)
	LogAdd(context.Context, *LogAddReq) (*LogAddResp, error)
	SysLogList(context.Context, *SysLogListReq) (*SysLogListResp, error)
	SysLogDelete(context.Context, *SysLogDeleteReq) (*SysLogDeleteResp, error)
	PlaceAdd(context.Context, *PlaceAddReq) (*PlaceAddResp, error)
	PlaceList(context.Context, *PlaceListReq) (*PlaceListResp, error)
	PlaceUpdate(context.Context, *PlaceUpdateReq) (*PlaceUpdateResp, error)
	PlaceDelete(context.Context, *PlaceDeleteReq) (*PlaceDeleteResp, error)
	mustEmbedUnimplementedSysServer()
}

// UnimplementedSysServer must be embedded to have forward compatible implementations.
type UnimplementedSysServer struct {
}

func (UnimplementedSysServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedSysServer) UserInfo(context.Context, *InfoReq) (*InfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedSysServer) UserAdd(context.Context, *UserAddReq) (*UserAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserAdd not implemented")
}
func (UnimplementedSysServer) UserUpdate(context.Context, *UserUpdateReq) (*UserUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserUpdate not implemented")
}
func (UnimplementedSysServer) UserDelete(context.Context, *UserDeleteReq) (*UserDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
func (UnimplementedSysServer) UserList(context.Context, *UserListReq) (*UserListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserList not implemented")
}
func (UnimplementedSysServer) ResetPassword(context.Context, *ResetPasswordReq) (*ResetPasswordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ResetPassword not implemented")
}
func (UnimplementedSysServer) RoleAdd(context.Context, *RoleAddReq) (*RoleAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleAdd not implemented")
}
func (UnimplementedSysServer) RoleUpdate(context.Context, *RoleUpdateReq) (*RoleUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleUpdate not implemented")
}
func (UnimplementedSysServer) RoleDelete(context.Context, *RoleDeleteReq) (*RoleDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}
func (UnimplementedSysServer) RoleList(context.Context, *RoleListReq) (*RoleListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleList not implemented")
}
func (UnimplementedSysServer) QueryMenuByRoleId(context.Context, *QueryMenuByRoleIdReq) (*QueryMenuByRoleIdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryMenuByRoleId not implemented")
}
func (UnimplementedSysServer) UpdateMenuRole(context.Context, *UpdateMenuRoleReq) (*UpdateMenuRoleResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMenuRole not implemented")
}
func (UnimplementedSysServer) MenuAdd(context.Context, *MenuAddReq) (*MenuAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuAdd not implemented")
}
func (UnimplementedSysServer) MenuList(context.Context, *MenuListReq) (*MenuListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuList not implemented")
}
func (UnimplementedSysServer) MenuUpdate(context.Context, *MenuUpdateReq) (*MenuUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuUpdate not implemented")
}
func (UnimplementedSysServer) MenuDelete(context.Context, *MenuDeleteReq) (*MenuDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MenuDelete not implemented")
}
func (UnimplementedSysServer) LoginLogAdd(context.Context, *LoginLogAddReq) (*LoginLogAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogAdd not implemented")
}
func (UnimplementedSysServer) LoginLogList(context.Context, *LoginLogListReq) (*LoginLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogList not implemented")
}
func (UnimplementedSysServer) LoginLogDelete(context.Context, *LoginLogDeleteReq) (*LoginLogDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginLogDelete not implemented")
}
func (UnimplementedSysServer) LogAdd(context.Context, *LogAddReq) (*LogAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LogAdd not implemented")
}
func (UnimplementedSysServer) SysLogList(context.Context, *SysLogListReq) (*SysLogListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLogList not implemented")
}
func (UnimplementedSysServer) SysLogDelete(context.Context, *SysLogDeleteReq) (*SysLogDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysLogDelete not implemented")
}
func (UnimplementedSysServer) PlaceAdd(context.Context, *PlaceAddReq) (*PlaceAddResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceAdd not implemented")
}
func (UnimplementedSysServer) PlaceList(context.Context, *PlaceListReq) (*PlaceListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceList not implemented")
}
func (UnimplementedSysServer) PlaceUpdate(context.Context, *PlaceUpdateReq) (*PlaceUpdateResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceUpdate not implemented")
}
func (UnimplementedSysServer) PlaceDelete(context.Context, *PlaceDeleteReq) (*PlaceDeleteResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceDelete not implemented")
}
func (UnimplementedSysServer) mustEmbedUnimplementedSysServer() {}

// UnsafeSysServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysServer will
// result in compilation errors.
type UnsafeSysServer interface {
	mustEmbedUnimplementedSysServer()
}

func RegisterSysServer(s grpc.ServiceRegistrar, srv SysServer) {
	s.RegisterService(&Sys_ServiceDesc, srv)
}

func _Sys_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserInfo(ctx, req.(*InfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserAdd(ctx, req.(*UserAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserUpdate(ctx, req.(*UserUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserDelete(ctx, req.(*UserDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UserList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UserList(ctx, req.(*UserListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_ResetPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetPasswordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).ResetPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_ResetPassword_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).ResetPassword(ctx, req.(*ResetPasswordReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RoleAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RoleAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RoleAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RoleAdd(ctx, req.(*RoleAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RoleUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RoleUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RoleUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RoleUpdate(ctx, req.(*RoleUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RoleDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RoleDelete(ctx, req.(*RoleDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_RoleList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).RoleList(ctx, req.(*RoleListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_QueryMenuByRoleId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryMenuByRoleIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).QueryMenuByRoleId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_QueryMenuByRoleId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).QueryMenuByRoleId(ctx, req.(*QueryMenuByRoleIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_UpdateMenuRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateMenuRoleReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).UpdateMenuRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_UpdateMenuRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).UpdateMenuRole(ctx, req.(*UpdateMenuRoleReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuAdd(ctx, req.(*MenuAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuList(ctx, req.(*MenuListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuUpdate(ctx, req.(*MenuUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_MenuDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MenuDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).MenuDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_MenuDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).MenuDelete(ctx, req.(*MenuDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogAdd(ctx, req.(*LoginLogAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogList(ctx, req.(*LoginLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LoginLogDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LoginLogDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LoginLogDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LoginLogDelete(ctx, req.(*LoginLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_LogAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).LogAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_LogAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).LogAdd(ctx, req.(*LogAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_SysLogList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLogListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).SysLogList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_SysLogList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).SysLogList(ctx, req.(*SysLogListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_SysLogDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysLogDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).SysLogDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_SysLogDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).SysLogDelete(ctx, req.(*SysLogDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PlaceAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PlaceAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_PlaceAdd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PlaceAdd(ctx, req.(*PlaceAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PlaceList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PlaceList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_PlaceList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PlaceList(ctx, req.(*PlaceListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PlaceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PlaceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_PlaceUpdate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PlaceUpdate(ctx, req.(*PlaceUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sys_PlaceDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysServer).PlaceDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sys_PlaceDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysServer).PlaceDelete(ctx, req.(*PlaceDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Sys_ServiceDesc is the grpc.ServiceDesc for Sys service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sys_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sys.Sys",
	HandlerType: (*SysServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Sys_Login_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _Sys_UserInfo_Handler,
		},
		{
			MethodName: "UserAdd",
			Handler:    _Sys_UserAdd_Handler,
		},
		{
			MethodName: "UserUpdate",
			Handler:    _Sys_UserUpdate_Handler,
		},
		{
			MethodName: "UserDelete",
			Handler:    _Sys_UserDelete_Handler,
		},
		{
			MethodName: "UserList",
			Handler:    _Sys_UserList_Handler,
		},
		{
			MethodName: "ResetPassword",
			Handler:    _Sys_ResetPassword_Handler,
		},
		{
			MethodName: "RoleAdd",
			Handler:    _Sys_RoleAdd_Handler,
		},
		{
			MethodName: "RoleUpdate",
			Handler:    _Sys_RoleUpdate_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _Sys_RoleDelete_Handler,
		},
		{
			MethodName: "RoleList",
			Handler:    _Sys_RoleList_Handler,
		},
		{
			MethodName: "QueryMenuByRoleId",
			Handler:    _Sys_QueryMenuByRoleId_Handler,
		},
		{
			MethodName: "UpdateMenuRole",
			Handler:    _Sys_UpdateMenuRole_Handler,
		},
		{
			MethodName: "MenuAdd",
			Handler:    _Sys_MenuAdd_Handler,
		},
		{
			MethodName: "MenuList",
			Handler:    _Sys_MenuList_Handler,
		},
		{
			MethodName: "MenuUpdate",
			Handler:    _Sys_MenuUpdate_Handler,
		},
		{
			MethodName: "MenuDelete",
			Handler:    _Sys_MenuDelete_Handler,
		},
		{
			MethodName: "LoginLogAdd",
			Handler:    _Sys_LoginLogAdd_Handler,
		},
		{
			MethodName: "LoginLogList",
			Handler:    _Sys_LoginLogList_Handler,
		},
		{
			MethodName: "LoginLogDelete",
			Handler:    _Sys_LoginLogDelete_Handler,
		},
		{
			MethodName: "LogAdd",
			Handler:    _Sys_LogAdd_Handler,
		},
		{
			MethodName: "SysLogList",
			Handler:    _Sys_SysLogList_Handler,
		},
		{
			MethodName: "SysLogDelete",
			Handler:    _Sys_SysLogDelete_Handler,
		},
		{
			MethodName: "PlaceAdd",
			Handler:    _Sys_PlaceAdd_Handler,
		},
		{
			MethodName: "PlaceList",
			Handler:    _Sys_PlaceList_Handler,
		},
		{
			MethodName: "PlaceUpdate",
			Handler:    _Sys_PlaceUpdate_Handler,
		},
		{
			MethodName: "PlaceDelete",
			Handler:    _Sys_PlaceDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sys.proto",
}