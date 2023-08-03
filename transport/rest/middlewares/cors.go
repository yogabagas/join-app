package middlewares

import "net/http"

func (m *MiddlewareImpl) CORSHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "*")

		w.Header().Set("Access-Control-Allow-Headers", "*")

		next.ServeHTTP(w, r)
	})
}
