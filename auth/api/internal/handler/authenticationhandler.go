package handler

import (
	"net/http"
	"server/auth/api/internal/logic"
	"server/auth/api/internal/svc"
	"server/common/response"
)

func authenticationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		l := logic.NewAuthenticationLogic(r.Context(), svcCtx)
		resp, err := l.Authentication(token)
		response.Response(w, resp, err)
	}
}
