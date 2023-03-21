package gan

import (
	"net/http"
)

type router struct {
	root     *node
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{
		root: &node{
			children: make(map[string]*node),
		},
		handlers: make(map[string]HandlerFunc),
	}

}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	r.root.insert(method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, pattern string) (node *node, params map[string]string) {
	return r.root.find(method, pattern)
}

func (r *router) handle(c *Context) {
	node, params := r.getRoute(c.Method, c.Path)
	if node != nil {
		c.Params = params
		key := c.Method + "-" + node.pattern
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}
