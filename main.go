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
	r.GET("/hello/:name", func(c *gin.Context) {

		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")

}
