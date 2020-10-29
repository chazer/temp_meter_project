package router

import "net/http"

type Route struct {
	Method  string
	Route   string
	Handler func(http.ResponseWriter, *http.Request)
}

type Routes struct {
	Items map[string]*Route
}

func NewRoutes() *Routes {
	router := &Routes{}
	router.Items = make(map[string]*Route)
	return router
}

func (r *Routes) RouteHandler(method string, path string, handler http.Handler) {
	r.Route(method, path, func(resp http.ResponseWriter, req *http.Request) {
		handler.ServeHTTP(resp, req)
	})
}

func (r *Routes) Route(method string, route string, handler func(http.ResponseWriter, *http.Request)) {
	r.Items[hash(method, route)] = &Route{method, route, handler}
}

func (r *Routes) GET(path string, handler http.HandlerFunc) {
	r.Route("GET", path, handler)
}

func (r *Routes) POST(path string, handler http.HandlerFunc) {
	r.Route("POST", path, handler)
}

func (r *Routes) PUT(path string, handler http.HandlerFunc) {
	r.Route("PUT", path, handler)
}

func (r *Routes) DELETE(path string, handler http.HandlerFunc) {
	r.Route("DELETE", path, handler)
}
