package coupon

import (
	"net/http"

	"epick-mall/api/internal/logic/coupon/coupon"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CouponDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DeleteCouponReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := coupon.NewCouponDeleteLogic(r.Context(), svcCtx)
		resp, err := l.CouponDelete(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
