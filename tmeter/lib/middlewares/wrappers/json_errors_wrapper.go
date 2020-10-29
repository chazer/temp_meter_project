package wrappers

import (
	"bytes"
	"net/http"
	"strings"
	config2 "tmeter/app/consts"
	"tmeter/lib/api/helpers"
	"tmeter/lib/debug"
	"tmeter/lib/env"
	"tmeter/lib/middlewares"
)

// middlewares.ResponseWrapperInterface
type jsonErrorsWrapperWriter struct {
	http.ResponseWriter
	name                      string
	buf                       *bytes.Buffer
	convertErrorMessageToJson bool
	jsonIndent                bool
	statusCode                int
	length                    int
}

func (s *jsonErrorsWrapperWriter) GetStatusCode() int {
	return s.statusCode
}

func (s *jsonErrorsWrapperWriter) GetLength() int {
	return s.length
}

func (s *jsonErrorsWrapperWriter) GetName() string {
	return s.name
}

func (s *jsonErrorsWrapperWriter) Write(p []byte) (int, error) {
	size, err := s.buf.Write(p)
	s.length += size
	return size, err
}

func (s *jsonErrorsWrapperWriter) Header() http.Header {
	return s.ResponseWriter.Header()
}

func (s *jsonErrorsWrapperWriter) WriteHeader(statusCode int) {
	s.statusCode = statusCode
	s.convertErrorMessageToJson = false

	if statusCode >= 400 && statusCode < 500 {
		s.Header().Del("Content-Type")
		s.Header().Del("Content-Length")
		s.Header().Set("Content-Type", "application/json; charset=utf-8")
		s.convertErrorMessageToJson = true
	}

	s.ResponseWriter.WriteHeader(statusCode)
}

func (s *jsonErrorsWrapperWriter) Close() {
	if s.convertErrorMessageToJson {
		errorMessage := strings.TrimSpace(s.buf.String())
		debug.Printf("<<< Error %d %s", s.GetStatusCode(), errorMessage)
		if doc, err := helpers.ErrorToJSON(errorMessage, s.jsonIndent); err == nil {
			s.ResponseWriter.Write(doc)
			return
		}
	}
	s.ResponseWriter.Write(s.buf.Bytes())
}

type JsonErrorsWrapperImplementation struct {
	indent bool
}

func (s *JsonErrorsWrapperImplementation) WrapResponse(resp http.ResponseWriter) middlewares.ResponseWrapperInterface {
	return &jsonErrorsWrapperWriter{
		ResponseWriter: resp,
		name:           "JSON-Errors",
		buf:            &bytes.Buffer{},
		jsonIndent:     env.GetBoolEnvOrDefault(config2.JsonIndentResponses, false),
	}
}

func NewJsonErrorsWrapper() *ResponseWrapper {
	return &ResponseWrapper{&JsonErrorsWrapperImplementation{}}
}
