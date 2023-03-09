package gyn

import (
	"net/http"
	"path"
)

type RouteGroup struct {
	handlers []HandlerFunc
	prefix   string
	parent   *RouteGroup
	engine   *Engine
}

func (group *RouteGroup) Group(prefix string) *RouteGroup {
	newGroup := &RouteGroup{
		prefix: path.Join(group.prefix, prefix),
		parent: group,
		engine: group.engine,
	}
	engine := group.engine
	engine.groups = append(engine.groups, newGroup)
	return newGroup
}

func (group *RouteGroup) Use(middlewares ...HandlerFunc) {
	group.handlers = append(group.handlers, middlewares...)
}

func (group *RouteGroup) createStaticHandler(relPath string, fs http.FileSystem) HandlerFunc {
	absPath := path.Join(group.prefix, relPath)
	fileServer := http.StripPrefix(absPath, http.FileServer(fs))
	return func(c *Context) {
		file := c.Param("filepath")
		if _, err := fs.Open(file); err != nil {
			c.Status(404)
			return
		}

		fileServer.ServeHTTP(c.Writer, c.Req)
	}
}

func (group *RouteGroup) Static(spath, fpath string) {
	handler := group.createStaticHandler(spath, http.Dir(fpath))
	pattern := path.Join(spath, "/*filepath")
	group.GET(pattern, handler)
}

func (group *RouteGroup) addRoute(method string, comp string, handler HandlerFunc) {
	pattern := group.prefix + comp
	group.engine.addRoute(method, pattern, handler)
}

func (group *RouteGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouteGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *RouteGroup) PUT(pattern string, handler HandlerFunc) {
	group.addRoute("PUT", pattern, handler)
}

func (group *RouteGroup) DELETE(pattern string, handler HandlerFunc) {
	group.addRoute("DELETE", pattern, handler)
}
