package views

import (
	"tmeter/lib/api/views"
)

type structTokenDoc struct {
	Token string `json:"token"`
}

type userTokenScheme struct {
	views.StructTaggingScheme
}

func (w *userTokenScheme) ToTaggedStruct(data interface{}) (interface{}, error) {
	token := data.(string)
	return structTokenDoc{
		Token: token,
	}, nil
}

func NewUserTokenResponseView() views.ApiViewInterface {
	return &views.ApiView{
		Scheme: &userTokenScheme{},
	}
}
