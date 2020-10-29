package middlewares

import "net/http"

type ResponseWrapperInterface interface {
	http.ResponseWriter
	Close()
	GetStatusCode() int
	GetLength() int
	GetName() string
}
