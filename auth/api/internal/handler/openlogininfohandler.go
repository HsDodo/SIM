package handler

import (
	"net/http"
	"server/auth/api/internal/logic"
	"server/auth/api/internal/svc"
	"server/common/response"
)

func openLoginInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := logic.NewOpenLoginInfoLogic(r.Context(), svcCtx)
		resp, err := l.OpenLoginInfo()
		response.Response(w, resp, err)

	}
}
