package place

import (
	"net/http"

	"epick-mall/api/internal/logic/sys/place"
	"epick-mall/api/internal/svc"
	"epick-mall/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PlaceListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ListPlaceReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := place.NewPlaceListLogic(r.Context(), svcCtx)
		resp, err := l.PlaceList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
