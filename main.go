package main

import (
	"fmt"
	"gan"
	"net/http"
)

func main() {
	g := gan.NewEngine()

	g.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you visited: %q", r.URL.Path)
	})

	g.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello: %q", r.URL.Path)
	})

	g.Run(":9090")

}
