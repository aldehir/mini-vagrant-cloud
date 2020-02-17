package server

import (
	"context"
	"net/http"
	"regexp"
)

type Router struct {
	rules []*routerRule
	NotFoundHandler http.Handler
}

type routerRule struct {
	pattern *regexp.Regexp
	handler http.Handler
}

func NewRouter() *Router {
	return &Router{
		NotFoundHandler: http.NotFoundHandler(),
		rules: []*routerRule{},
	}
}

func newRouterRule(pattern string, handler http.Handler) (*routerRule, error) {
	compiled, err := regexp.Compile(pattern)

	if err != nil {
		return nil, err
	}

	return &routerRule{compiled, handler}, nil
}

func (r *Router) Handle(pattern string, handler http.Handler) {
	rule, err := newRouterRule(pattern, handler)
	if err != nil {
		panic(err)
	}

	r.rules = append(r.rules, rule)
}

func (r *Router) HandleFunc(pattern string, f func(http.ResponseWriter, *http.Request)) {
	r.Handle(pattern, http.HandlerFunc(f))
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, rule := range r.rules {
		if match := rule.pattern.FindStringSubmatch(req.URL.Path); len(match) > 0 {
			params := composeParams(rule.pattern.SubexpNames(), match)
			ctx := context.WithValue(req.Context(), "params", params)
			req = req.WithContext(ctx)

			rule.handler.ServeHTTP(w, req)
			return
		}
	}

	r.NotFoundHandler.ServeHTTP(w, req)
}

func composeParams(names []string, matches []string) map[string]string {
	names = names[1:]
	matches = matches[1:]
	params := make(map[string]string)

	for idx, name := range names {
		params[name] = matches[idx]
	}

	return params
}
