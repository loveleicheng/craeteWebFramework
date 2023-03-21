package main

import (
	"gan"
	"net/http"
)

func main() {
	r := gan.NewEngine()
	r.GET("/", func(c *gan.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gan</h1>")
	})

	r.GET("/hello", func(c *gan.Context) {
		// expect /hello?name=lei
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gan.Context) {
		// expect /hello/lei
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gan.Context) {
		c.JSON(http.StatusOK, gan.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9000")
}
