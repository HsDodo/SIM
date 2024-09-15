package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/logs/api/internal/logic"
	"server/logs/api/internal/svc"
	"server/logs/api/internal/types"
)

func logReadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LogReadRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewLogReadLogic(r.Context(), svcCtx)
		resp, err := l.LogRead(&req)
		response.Response(w, resp, err)

	}
}
