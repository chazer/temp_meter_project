package api

import (
	"net/http"
	"tmeter/app/consts"
	"tmeter/lib/api/views"
	"tmeter/lib/env"
)

type APIProtocol struct{}

func (s *APIProtocol) json(view views.ApiViewInterface, data interface{}) ([]byte, error) {
	jsonIndent := env.GetBoolEnvOrDefault(consts.JsonIndentResponses, false)
	return view.RenderEntityAsJSON(data, jsonIndent)
}

func (s *APIProtocol) WriteEntityDocumentResponse(resp http.ResponseWriter, f *FormatterConfig, data interface{}) {
	// TODO: verify formatter value
	json, err := s.json(f.DocumentView, data)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(json)
	resp.Write([]byte("\n"))
}

func (s *APIProtocol) WritePageDocumentResponse(resp http.ResponseWriter, f *FormatterConfig, data PageDocument) {
	// TODO: verify formatter value
	json, err := s.json(f.PageApiView, &data)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)
	resp.Write(json)
	resp.Write([]byte("\n"))
}

func (s *APIProtocol) SendSlice(resp http.ResponseWriter, f *FormatterConfig, data []interface{}) {
	s.WritePageDocumentResponse(resp,
		f,
		PageDocument{
			Items: data,
			Total: len(data),
		},
	)
}
