package gyn

import (
	"html/template"
	"net/http"
	"strings"
)

type HandlerFunc func(*Context)

type Engine struct {
	*RouteGroup
	router        *router
	groups        []*RouteGroup
	htmlTemplates *template.Template
}

func Default() *Engine {
	engine := New()
	engine.Use(Logger(), Recovery())
	return engine
}

func New() *Engine {
	engine := &Engine{}
	engine.RouteGroup = &RouteGroup{engine: engine}
	engine.router = newRouter()
	engine.groups = []*RouteGroup{engine.RouteGroup}
	return engine
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func (engine *Engine) LoadHTMLTemplates(pattern string) {
	engine.htmlTemplates = template.Must(template.ParseGlob(pattern))
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) PUT(pattern string, handler HandlerFunc) {
	engine.addRoute("PUT", pattern, handler)
}

func (engine *Engine) DELETE(pattern string, handler HandlerFunc) {
	engine.addRoute("DELETE", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var handlers []HandlerFunc
	for _, group := range engine.groups {
		if strings.HasPrefix(req.URL.Path, group.prefix) {
			handlers = append(handlers, group.handlers...)
		}
	}
	c := newContext(w, req)
	c.engine = engine
	c.handlers = handlers
	engine.router.handle(c)
}
