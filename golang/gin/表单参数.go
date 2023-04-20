package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		types := c.DefaultQuery("username", "post")
		username := c.PostForm("username")
		password := c.PostForm("userpassword")
		c.String(http.StatusOK,fmt.Sprintf("username:%s,password:%s,type:%s",username,password,types))
	})
	//r.GET("/user/:name/*action", func(context *gin.Context) {
	//	name := context.Param("name")
	//	action := context.Param("action")
	//	action = strings.Trim(action, "/")
	//	context.String(http.StatusOK,name+"is"+action)
	//})
	//r.Run(":8000")
	r.Run()
}