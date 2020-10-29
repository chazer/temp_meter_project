package views

import (
	"encoding/json"
)

type StructTaggingScheme interface {
	ToTaggedStruct(data interface{}) (interface{}, error)
}

type ApiViewInterface interface {
	GetTaggingScheme() StructTaggingScheme
	RenderEntityAsJSON(data interface{}, jsonIndent bool) ([]byte, error)
}

//	ApiViewInterface
type ApiView struct {
	Scheme StructTaggingScheme
}

func (s *ApiView) GetTaggingScheme() StructTaggingScheme {
	return s.Scheme
}

func (s *ApiView) RenderEntityAsJSON(data interface{}, jsonIndent bool) ([]byte, error) {
	st, err := s.Scheme.ToTaggedStruct(data)
	if err != nil {
		return nil, err
	}
	if jsonIndent {
		return json.MarshalIndent(st, "", "  ")
	}
	return json.Marshal(st)
}
