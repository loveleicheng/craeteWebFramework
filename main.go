package main

import (
	"gan"
	"net/http"
)

func main() {
	g := gan.NewEngine()

	g.GET("/hello", func(c *gan.Context) {
		c.String(http.StatusOK, "visited: %q\n", c.Path)
	})

	g.POST("/login", func(c *gan.Context) {
		c.JSON(http.StatusOK, gan.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	g.Run(":9090")

}
