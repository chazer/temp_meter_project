package middlewares

import (
	"net/http"
	"tmeter/lib/debug"
)

type ResponseWrapperFunc = func(resp http.ResponseWriter) ResponseWrapperInterface

func ResponseWrapperMiddleware(wrap ResponseWrapperFunc) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		debug.Print("init")

		return func(resp http.ResponseWriter, req *http.Request) {
			writer := wrap(resp)

			debug.Printf("[%s] request", writer.GetName())

			next.ServeHTTP(writer, req)
			writer.Close()

			ct := writer.Header().Get("Content-Type")
			debug.Printf("[%s] << %d, type=%s, length=%d", writer.GetName(), writer.GetStatusCode(), ct, writer.GetLength())
		}
	}
}
