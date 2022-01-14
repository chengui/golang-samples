package gee

import (
	"log"
)

type router struct {
	routes   map[string]*Trie
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		routes:   make(map[string]*Trie),
		handlers: make(map[string]HandlerFunc),
	}
}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	if _, ok := r.routes[method]; !ok {
		r.routes[method] = NewTrie()
	}
	r.routes[method].Insert(pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	var handler HandlerFunc
	if trie, ok := r.routes[c.Method]; ok {
		pattern, params := trie.Search(c.Path)
		if pattern != "" {
			c.Params = params
			key := c.Method + "-" + pattern
			if hdlr, ok := r.handlers[key]; ok {
				handler = hdlr
			} else {
				handler = func(c *Context) {
					c.String(404, "HANDLER NOT FOUND: %s\n", c.Path)
				}
			}
		} else {
			handler = func(c *Context) {
				c.String(404, "PATTERN NOT FOUND: %s\n", c.Path)
			}
		}
	} else {
		handler = func(c *Context) {
			c.String(400, "BAD REQUEST: %s\n", c.Method)
		}
	}
	c.handlers = append(c.handlers, handler)
	c.Next()
}
