package views

import (
	"tmeter/app/api"
	"tmeter/lib/api/views"
)

type structPageDoc struct {
	Items []interface{} `json:"items"`
	Total int           `json:"total"`
	Next  *string       `json:"next"`
}

type listApiScheme struct {
	views.StructTaggingScheme
	ApiView views.ApiViewInterface
}

func (w *listApiScheme) ToTaggedStruct(i interface{}) (interface{}, error) {
	d := i.(*api.PageDocument)

	iii := make([]interface{}, len(d.Items))

	for n, item := range d.Items {
		h, err := w.ApiView.GetTaggingScheme().ToTaggedStruct(item)
		if err != nil {
			return nil, err
		}
		iii[n] = h
	}
	return structPageDoc{
		Items: iii,
		Total: len(d.Items),
	}, nil
}

func NewPageApiView(v views.ApiViewInterface) views.ApiViewInterface {
	return &views.ApiView{
		Scheme: &listApiScheme{ApiView: v},
	}
}
