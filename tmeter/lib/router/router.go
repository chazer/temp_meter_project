package router

import (
	"net/http"
	"net/url"
	"path"
	"tmeter/lib/debug"
)

// TODO: use constant keys for add errors routes
const errNotFound = http.StatusNotFound

type HandlerKey struct {
	Method string
	Path   string
	Error  int
}

type Router struct {
	pathPrefix       string
	byMethodServeMux map[string]*http.ServeMux
}

func (r *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	debug.Print("ServeHTTP " + req.URL.Path)

	Handler, ok := r.byMethodServeMux[req.Method]
	if ok {
		Handler.ServeHTTP(resp, req)
	} else {
		resp.WriteHeader(404)
	}
}

func NewRouter() *Router {
	router := new(Router)
	router.pathPrefix = ""
	router.byMethodServeMux = make(map[string]*http.ServeMux)
	return router
}

func pathJoin(s ...string) string {
	p := path.Join(s...)
	if len(s) > 0 {
		if s[len(s)-1] == "/" {
			return p + "/"
		}
	}
	return p
}

// parse route definition
// format:  "/path/name?args"
func parseRoute(route *Route) string {
	u, err := url.Parse(route.Route)
	if err != nil {
		return route.Route
	}
	return u.Path
}

func (r *Router) AddRoute(path string, route *Route) {
	_, ok := r.byMethodServeMux[route.Method]
	if !ok {
		r.byMethodServeMux[route.Method] = http.NewServeMux()
	}

	routePath := pathJoin(path, parseRoute(route))
	debug.Printf("add handler: %s %s", route.Method, routePath)

	mux := r.byMethodServeMux[route.Method]
	mux.HandleFunc(routePath, func(resp http.ResponseWriter, req *http.Request) {
		// Only full path matched supported
		if req.URL.Path != routePath {
			resp.WriteHeader(404)
			return
		}
		route.Handler(resp, req)
	})
}

func (r *Router) AddRoutes(path string, routes *Routes) {
	for _, element := range routes.Items {
		r.AddRoute(path, element)
	}
}

// TODO: add route for errors
