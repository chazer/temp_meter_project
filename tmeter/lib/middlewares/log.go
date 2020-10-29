package middlewares

import (
	"bytes"
	"net/http"
)
import "tmeter/lib/debug"

type logWriter struct {
	http.ResponseWriter
	buf        *bytes.Buffer
	statusCode int
	length     int
}

func (s *logWriter) Write(b []byte) (int, error) {
	count, err := s.ResponseWriter.Write(b)
	s.length += count
	return count, err
}

func (s *logWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

func LogMiddleware(h http.HandlerFunc) http.HandlerFunc {
	debug.Printf("init")
	return func(resp http.ResponseWriter, req *http.Request) {
		debug.Printf("[LogMiddleware] >> %s %s, type=%s, length=%s ", req.Method, req.URL, req.Header.Get("Content-Type"), req.Header.Get("Content-Length"))

		wr := &logWriter{ResponseWriter: resp}
		h.ServeHTTP(wr, req)

		ct := resp.Header().Get("Content-Type")
		debug.Printf("[LogMiddleware] << %d, type=%s, length=%d", wr.statusCode, ct, wr.length)
	}
}
