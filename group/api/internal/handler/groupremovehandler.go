package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/group/api/internal/logic"
	"server/group/api/internal/svc"
	"server/group/api/internal/types"
)

func groupRemoveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupRemoveRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewGroupRemoveLogic(r.Context(), svcCtx)
		resp, err := l.GroupRemove(&req)
		response.Response(w, resp, err)

	}
}
