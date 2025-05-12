package middlewares

import "net/http"

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO : complete this logic
		next.ServeHTTP(w, r)
	})
}
