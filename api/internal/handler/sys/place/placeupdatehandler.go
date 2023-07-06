package place

import (
	"net/http"

	"epick-mall/api/internal/logic/sys/place"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlaceUpdateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdatePlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := place.NewPlaceUpdateLogic(r.Context(), svcCtx)
		resp, err := l.PlaceUpdate(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
