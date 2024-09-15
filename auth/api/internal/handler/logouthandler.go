package handler

import (
	"errors"
	"net/http"
	"server/auth/api/internal/logic"
	"server/auth/api/internal/svc"
	"server/common/response"
)

func logoutHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			response.Response(w, nil, errors.New("token不能为空"))
			return
		}
		//有token的话，注销登录
		l := logic.NewLogoutLogic(r.Context(), svcCtx)
		resp, err := l.Logout(token)
		response.Response(w, resp, err)

	}
}
