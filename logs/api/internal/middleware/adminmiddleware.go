package middleware

import (
	"errors"
	"net/http"
	"server/common/response"
)

type AdminMiddleware struct {
}

func NewAdminMiddleware() *AdminMiddleware {
	return &AdminMiddleware{}
}

func (m *AdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 判断是否是管理员
		role := r.Header.Get("Role")
		if role != "2" { // 2 是管理员
			response.Response(w, nil, errors.New("角色鉴权失败"))
			return
		}

		next(w, r)
	}
}
