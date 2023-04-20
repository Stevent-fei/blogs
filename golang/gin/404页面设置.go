package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "liuxin")
		c.String(http.StatusOK,fmt.Sprintf("hello %s",name))
	})
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound,"404 not found...")
	})
	r.Run()
}