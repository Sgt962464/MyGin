package main

import (
	"fmt"
	"gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "<h1>hello mgin<h1>")
	})
	r.GET("/hello", func(c *gin.Context) {
		fmt.Println(c.Path)
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

}