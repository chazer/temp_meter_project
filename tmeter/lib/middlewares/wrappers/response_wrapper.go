package wrappers

import (
	"net/http"
	"tmeter/lib/middlewares"
)

// middlewares.ResponseWrapperInterface
type DummyWrapperWriter struct {
	http.ResponseWriter
	statusCode int
	length     int
	Name       string
}

func (s *DummyWrapperWriter) GetName() string {
	return s.Name
}

func (s *DummyWrapperWriter) Close() {}

func (s *DummyWrapperWriter) GetStatusCode() int {
	return s.statusCode
}
func (s *DummyWrapperWriter) GetLength() int {
	return s.length
}

func (s *DummyWrapperWriter) Write(p []byte) (int, error) {
	size, err := s.ResponseWriter.Write(p)
	s.length += size
	return size, err
}

func (s *DummyWrapperWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
	s.ResponseWriter.WriteHeader(statusCode)
}

type ResponseWrapperImplementationInterface interface {
	WrapResponse(resp http.ResponseWriter) middlewares.ResponseWrapperInterface
}

// ResponseWrapperImplementationInterface
type ResponseWrapperDefaultImplementation struct{}

func (ad ResponseWrapperDefaultImplementation) WrapResponse(resp http.ResponseWriter) middlewares.ResponseWrapperInterface {
	return &DummyWrapperWriter{Name: "Dummy", ResponseWriter: resp}
}

type ResponseWrapper struct {
	impl ResponseWrapperImplementationInterface
}

func (s *ResponseWrapper) AsMiddleware() middlewares.Middleware {
	return middlewares.ResponseWrapperMiddleware(s.impl.WrapResponse)
}

func NewDummyResponseWrapper() *ResponseWrapper {
	return &ResponseWrapper{&ResponseWrapperDefaultImplementation{}}
}
