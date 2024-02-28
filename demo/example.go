package main

import (
	"fmt"
	"net/http"

	"github.com/cpf2021-gif/gow"
)

func main() {
	r := gow.Default()

	r.GET("/", func(c *gow.Context) {
		// c.HTML(http.StatusOK, "<h1>Hello gow</h1>")
		c.Render(200, page(), gow.H{})
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
			// test panic
			nums := []int{1, 2, 3}
			fmt.Println(nums[100])
			c.HTML(http.StatusOK, "<h1>Hello gow, v1<h1>")
		})
	}

	r.Run(":9999")
}
