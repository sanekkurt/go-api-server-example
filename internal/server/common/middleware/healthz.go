package middleware

import (
	"net/http"
	"strings"
)

func HealthcheckMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && strings.HasSuffix(strings.ToLower(r.URL.Path), "/healthz") {
			w.Header().Set("Content-Type", "text/plain")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK")) //nolint
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

