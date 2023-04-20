package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	//1. 异步
	r.GET("/long_async", func(c *gin.Context) {
		//搞一个副本
		copyContext := c.Copy()
		//异步处理
		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步执行: " + copyContext.Request.URL.Path)
			c.String(http.StatusOK, copyContext.Request.URL.Path)
		}()
	})
	//2. 同步
	r.GET("/long_sync", func(c *gin.Context) {
		log.Println("同步执行: " + c.Request.URL.Path)
		c.String(http.StatusOK, c.Request.URL.Path)
	})
	r.Run(":8000")
}
