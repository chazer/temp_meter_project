package helpers

import (
	"encoding/json"
	"strings"
)

type structApiErrorData struct {
	Message string `json:"message"`
}

type structApiErrorDoc struct {
	Error structApiErrorData `json:"error"`
}

func ErrorToJSON(msg string, indent bool) ([]byte, error) {
	doc := structApiErrorDoc{
		Error: structApiErrorData{
			Message: strings.TrimSpace(msg),
		},
	}
	if indent {
		return json.MarshalIndent(doc, "", "  ")
	}
	return json.Marshal(doc)
}
