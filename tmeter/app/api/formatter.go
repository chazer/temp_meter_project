package api

import "tmeter/lib/api/views"

type FormatterConfig struct {
	DocumentView views.ApiViewInterface
	PageApiView  views.ApiViewInterface
}

type Formatter struct {
	config FormatterConfig
}
