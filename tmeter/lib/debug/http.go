package debug

import (
	"net/http"
)

func DumpHeaders(h http.Header) {
	for k, v := range h {
		for _, vv := range v {
			Printf("Header %s = %s", k, vv)
		}
	}
}
