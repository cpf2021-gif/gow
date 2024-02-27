package main

import (
	"net/http"

	"github.com/cpf2021-gif/gow"
)

func main() {
	r := gow.New()

	r.GET("/", func(c *gow.Context) {
		c.HTML(http.StatusOK, "<h1>Hello gow</h1>")
	})

	r.GET("/hello", func(c *gow.Context) {
		c.String(http.StatusOK, "hello, Have a nice day!\n")
	})

	r.POST("/login", func(c *gow.Context) {
		c.JSON(http.StatusOK, gow.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.GET("people/:name/info", func(c *gow.Context) {
		c.JSON(http.StatusOK, gow.H{"name": c.Param("name")})
	})

	r.GET("/file/*filepath", func(c *gow.Context) {
		c.JSON(http.StatusOK, gow.H{"filepath": c.Param("filepath")})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gow.Context) {
			c.HTML(http.StatusOK, "<h1>Hello gow, v1<h1>")
		})
	}

	r.Run(":9999")
}
