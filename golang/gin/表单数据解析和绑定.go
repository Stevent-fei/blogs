package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 定义接收数据的结构体
type Login struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	User    string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Pssword string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	r := gin.Default()
	//json 绑定
	r.POST("loginForm", func(c *gin.Context) {
		var form Login
		// Bind()默认解析并绑定form格式
		// 根据请求头中content-type自动推断
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		}
		//判断用户名密码是否正确

		if form.User != "root" || form.Pssword != "admin"{
			c.JSON(http.StatusBadRequest,gin.H{"status":"304"})
		}
		c.JSON(http.StatusOK,gin.H{"status":"200"})
	})
	r.Run(":8000")
}