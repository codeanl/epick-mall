package middleware

import (
	"epick-mall/service/sys/rpc/sys"
	"epick-mall/service/sys/rpc/sysclient"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"time"
)

type AddLogMiddleware struct {
	Sys sysclient.Sys
}

func NewAddLogMiddleware(Sys sysclient.Sys) *AddLogMiddleware {
	return &AddLogMiddleware{Sys: Sys}
}

func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/system/user/login" {
			next(w, r)
			return
		}
		//username := r.Context().Value("username").(string)
		all, _ := io.ReadAll(r.Body)
		startTime := time.Now()
		next(w, r)
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		// 添加操作日志
		_, _ = m.Sys.LogAdd(r.Context(), &sys.LogAddReq{
			Username: "username",
			//Username:  username,
			Operation: r.Method,
			Method:    r.RequestURI,
			Params:    string(all),
			Time:      duration.Microseconds(),
			Ip:        httpx.GetRemoteAddr(r),
		})
	}
}
