package gan

import (
	"fmt"
	"log"
	"net/http"
)

type Handfunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]Handfunc
}

func NewEngine() *Engine {
	return &Engine{
		router: make(map[string]Handfunc),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handler Handfunc) {
	key := method + "-" + pattern
	log.Printf("Router: %s", key)
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler Handfunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler Handfunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if hander, ok := engine.router[key]; ok {
		hander(w, r)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s", r.URL.Path)
	}
}

func (engine *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)

}
