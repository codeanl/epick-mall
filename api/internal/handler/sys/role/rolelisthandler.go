package role

import (
	"net/http"

	"epick-mall/api/internal/logic/sys/role"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RoleListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListRoleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := role.NewRoleListLogic(r.Context(), svcCtx)
		resp, err := l.RoleList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
