package main

import (
	"gan"
	"net/http"
)

func main() {
	r := gan.New()
	r.GET("/index", func(c *gan.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gan.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gan</h1>")
		})

		v1.GET("/hello", func(c *gan.Context) {
			// expect /hello?name=lei
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gan.Context) {
			// expect /hello/lei
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gan.Context) {
			c.JSON(http.StatusOK, gan.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9000")
}
