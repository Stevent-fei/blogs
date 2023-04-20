package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
//type Login struct {
//	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
//	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
//	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
//}

func main() {
	r := gin.Default()

	r.POST("loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			return
		}
		if json.User != "root" || json.Pssword != "admin" {
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
			return
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	r.Run(":8080")
}
