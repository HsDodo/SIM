package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"server/chat/api/internal/logic"
	"server/chat/api/internal/svc"
	"server/chat/api/internal/types"
	"server/common/response"
)

func chatDeleteHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatDeleteRequest
		if err := httpx.Parse(r, &req); err != nil {
			response.Response(w, nil, err)
			return
		}

		l := logic.NewChatDeleteLogic(r.Context(), svcCtx)
		resp, err := l.ChatDelete(&req)
		response.Response(w, resp, err)

	}
}
