package gan

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router       *router
	*RouterGroup // Engine 作为顶层核心，需要由其调用router group能力
}

type RouterGroup struct {
	prefix      string
	middlewares []HandlerFunc // 支持中间件
	engine      *Engine       // router group 需要有操作路由的能力
}

func New() *Engine {
	engine := &Engine{
		router: newRouter(),
	}
	engine.RouterGroup = &RouterGroup{
		middlewares: make([]HandlerFunc, 1),
		engine:      engine,
	}
	return engine
}

func (group *RouterGroup) Group(prefix string) *RouterGroup {
	newGroup := &RouterGroup{
		prefix:      group.prefix + prefix,
		middlewares: make([]HandlerFunc, 1),
		engine:      group.engine,
	}
	return newGroup
}

func (group *RouterGroup) addRoute(method string, pattern string, handler HandlerFunc) {
	pattern = group.prefix + pattern
	group.engine.router.addRoute(method, pattern, handler)
}

func (group *RouterGroup) GET(pattern string, handler HandlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *RouterGroup) POST(pattern string, handler HandlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	engine.router.handle(c)
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)

}
