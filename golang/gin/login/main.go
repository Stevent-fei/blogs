package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 主页路由处理函数
func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录",
	})
}

// 登录表单提交处理函数
func loginHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	// TODO: 处理用户登录逻辑
	if username == "" && password == "" {
		c.JSON(http.StatusBadRequest,gin.H{"error": "Login name and password cannot be empty"})
	}
	//跳转页面
	c.Redirect(http.StatusFound, "http://www.baidu.com")
}

func main() {
	// 创建gin实例
	r := gin.Default()

	// 注册静态文件服务
	r.Static("static/css", "static/*")



	// 注册模板引擎
	r.LoadHTMLGlob("/Users/wangfei/go/src/Blogs/golang/gin/login/templates/*")

	// 注册路由
	r.GET("/", homeHandler)
	r.POST("/login", loginHandler)

	// 启动Web服务器
	r.Run(":8080")
}