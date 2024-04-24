package middleware

import (
	"net/http"
	"strings"
)

type CheckSpiderMiddleware struct {
}

func NewCheckSpiderMiddleware() *CheckSpiderMiddleware {
	return &CheckSpiderMiddleware{}
}

func (m *CheckSpiderMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 检查user-agent是否禁用
		if m.isUserAgentForbidden(r) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r)
	}
}

// 检测是否是禁用的User-Agent
func (m *CheckSpiderMiddleware) isUserAgentForbidden(r *http.Request) bool {
	ua := r.Header.Get("User-Agent")
	if ua == "" {
		return false
	}

	// 禁用的User-Agent列表
	spiderList := []string{
		"scrapy",
		"curl",
		"wget",
		"python-requests",
	}

	for _, spider := range spiderList {
		if strings.Contains(strings.ToLower(ua), strings.ToLower(spider)) {
			return true
		}
	}

	return false

}
