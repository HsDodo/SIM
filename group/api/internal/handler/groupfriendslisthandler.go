package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/common/response"
	"server/group/api/internal/logic"
	"server/group/api/internal/svc"
	"server/group/api/internal/types"
)

func groupfriendsListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GroupfriendsListRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewGroupfriendsListLogic(r.Context(), svcCtx)
		resp, err := l.GroupfriendsList(&req)
		response.Response(w, resp, err)

	}
}
