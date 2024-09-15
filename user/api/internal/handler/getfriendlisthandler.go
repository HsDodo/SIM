package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/user/api/internal/logic"
	"server/user/api/internal/svc"
	"server/user/api/internal/types"
)

func getFriendListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetFriendListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewGetFriendListLogic(r.Context(), svcCtx)
		resp, err := l.GetFriendList(&req)
		response.Response(w, resp, err)

	}
}
