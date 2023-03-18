package gan

import (
	"log"
	"net/http"
)

type router struct {
	handlers map[string]Handfunc
}

func newRouter() *router {
	return &router{
		handlers: make(map[string]Handfunc),
	}

}

func (r *router) addRoute(method string, pattern string, handler Handfunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	log.Printf("Router: %s", key)
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
