package middlewares

import "net/http"

// TODO: do not work, fix it
// Merge middlewares sequence
func FoldR(handlers ...Middleware) Middleware {
	if len(handlers) == 0 {
		panic("FoldR: elements expected")
	}
	var result Middleware = nil
	for i := len(handlers) - 1; i >= 0; i-- {
		if result == nil {
			result = handlers[i]
		} else {
			fn := handlers[i]
			result = func(h http.HandlerFunc) http.HandlerFunc { return fn(h) }
		}
	}
	return result
}
