package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/neocxf/go-exercises/web/gin/json"
)

func main() {
	r := gin.Default()

	fmt.Println(json.Hello)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "hello %s", name)
	})

	r.Run()
}
