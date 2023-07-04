package middleware

import (
	"epick-mall/service/system/rpc/system"
	"epick-mall/service/system/rpc/systemclient"
	"github.com/zeromicro/go-zero/rest/httpx"
	"io"
	"net/http"
	"time"
)

type AddLogMiddleware struct {
	System systemclient.System
}

func NewAddLogMiddleware(System systemclient.System) *AddLogMiddleware {
	return &AddLogMiddleware{System: System}
}

func (m *AddLogMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/api/system/user/login" {
			next(w, r)
			return
		}
		phone := r.Context().Value("phone").(string)
		all, _ := io.ReadAll(r.Body)
		startTime := time.Now()
		next(w, r)
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		// 添加操作日志
		_, _ = m.System.LogAdd(r.Context(), &system.LogAddReq{
			Phone:     phone,
			Operation: r.Method,
			Method:    r.RequestURI,
			Params:    string(all),
			Time:      duration.Microseconds(),
			Ip:        httpx.GetRemoteAddr(r),
		})
	}
}
