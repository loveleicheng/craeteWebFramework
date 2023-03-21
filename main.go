package main

import (
	"gan"
	"log"
	"net/http"
	"time"
)

func onlyForV2() gan.HandlerFunc {
	return func(c *gan.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := gan.New()
	r.Use(gan.Logger()) // global midlleware
	r.GET("/", func(c *gan.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gan</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *gan.Context) {
			// expect /hello/lei
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9000")
}
