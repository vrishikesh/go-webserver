package router

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	pattern *regexp.Regexp
	handler map[string]http.Handler
}

type RegexRouter struct {
	routes map[string]*route
}

func NewRegexRouter() *RegexRouter {
	return &RegexRouter{
		routes: make(map[string]*route, 0),
	}
}

func (rr *RegexRouter) Handler(pattern *regexp.Regexp, method string, handler http.Handler) {
	path := pattern.String()
	path = strings.TrimRight(path, "/")

	var r *route
	if temp, ok := rr.routes[path]; ok {
		r = temp
	}

	if r == nil {
		r = &route{
			pattern: pattern,
			handler: make(map[string]http.Handler),
		}
		rr.routes[path] = r
	}
	if _, ok := r.handler[method]; ok {
		log.Panicf("handler already exists for pattern [%s] and method [%s]", pattern, method)
	}
	r.handler[method] = handler
}

func (rr *RegexRouter) HandlerFunc(pattern *regexp.Regexp, method string, handler func(http.ResponseWriter, *http.Request)) {
	rr.Handler(pattern, method, http.HandlerFunc(handler))
}

func (rr *RegexRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range rr.routes {
		route := route
		if route.pattern.MatchString(r.URL.Path) {
			if handle, ok := route.handler[r.Method]; ok {
				handle.ServeHTTP(w, r)
				return
			}
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
			return
		}
	}

	http.NotFound(w, r)
}
