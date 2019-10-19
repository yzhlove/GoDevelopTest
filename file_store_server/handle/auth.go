package handle

import (
	"WorkSpace/DevelopTest/file_store_server/util"
	"net/http"
)

//拦截器
func HttpInterceptor(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm()
		userName := r.Form.Get("username")
		token := r.Form.Get("token")

		if len(userName) < 3 || !isTokenValid(userName, token) {
			w.WriteHeader(http.StatusForbidden)
			_, _ = w.Write(util.NewResp(-1, "sign invalid", nil).JSONBytes())
			return
		}
		handler(w, r)
	}
}
