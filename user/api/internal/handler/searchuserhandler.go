package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/user/api/internal/logic"
	"server/user/api/internal/svc"
	"server/user/api/internal/types"
)

func searchUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUserRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewSearchUserLogic(r.Context(), svcCtx)
		resp, err := l.SearchUser(&req)
		response.Response(w, resp, err)

	}
}
