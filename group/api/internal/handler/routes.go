// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"server/group/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/group/friends",
				Handler: groupfriendsListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/group/group",
				Handler: groupCreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/group",
				Handler: groupUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/group/:groupID",
				Handler: groupInfoHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/group/:groupID",
				Handler: groupRemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/history/:groupID",
				Handler: groupHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/history/:groupID",
				Handler: groupHistoryDeleteHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/member",
				Handler: groupMemberHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/group/member",
				Handler: groupMemberRemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/group/member",
				Handler: groupMemberAddHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/member/nickname",
				Handler: groupMemberNicknameUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/member/prohibition",
				Handler: groupProhibitionUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/member/role",
				Handler: groupMemberRoleUpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/my",
				Handler: groupMyHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/search",
				Handler: groupSearchHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/session",
				Handler: groupSessionHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/top",
				Handler: groupTopHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/group/valid",
				Handler: groupValidAddHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/valid",
				Handler: groupValidListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/valid/:groupID",
				Handler: groupValidHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/group/valid/status",
				Handler: groupValidStatusHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/group/ws/chat",
				Handler: groupChatHandler(serverCtx),
			},
		},
	)
}
