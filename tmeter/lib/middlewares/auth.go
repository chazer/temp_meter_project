package middlewares

import "net/http"
import "tmeter/lib/debug"

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	debug.Print("init")
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		debug.Printf("[AuthMiddleware] >>")

		next.ServeHTTP(resp, req)

		debug.Printf("[AuthMiddleware] <<")
	})
}
