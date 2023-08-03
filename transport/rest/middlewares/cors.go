package middlewares

import "net/http"

func (m *MiddlewareImpl) CORSHandle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")

		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token")

		w.Header().Set("Access-Control-Expose-Headers", "Authorization")

		next.ServeHTTP(w, r)
	})
}
