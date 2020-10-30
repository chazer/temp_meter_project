package views

import (
	"tmeter/lib/api/views"
)

type structTokenDoc struct {
	Token string `json:"token"`
}

type tokenScheme struct {
	views.StructTaggingScheme
}

func (w *tokenScheme) ToTaggedStruct(data interface{}) (interface{}, error) {
	token := data.(string)
	return structTokenDoc{
		Token: token,
	}, nil
}

func NewAuthTokenResponseView() views.ApiViewInterface {
	return &views.ApiView{
		Scheme: &tokenScheme{},
	}
}
