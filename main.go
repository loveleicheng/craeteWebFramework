package main

import (
	"fmt"
	"gan"
	"html/template"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gan.New()
	r.Use(gan.Logger())

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Lei", Age: 20}
	stu2 := &student{Name: "Rayer", Age: 22}

	r.GET("/", func(c *gan.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})

	r.GET("/students", func(c *gan.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gan.H{
			"title":  "gan",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gan.Context) {
		c.HTML(http.StatusOK, "func_example.tmpl", gan.H{
			"title": "gan",
			"now":   time.Date(2023, 3, 22, 16, 21, 0, 0, time.UTC),
		})
	})

	r.Run(":9000")
}
