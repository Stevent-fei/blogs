package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//必须拿到这个接口
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("file")
		if err != nil {
			c.String(500, "上传图片出错")
		}
		c.SaveUploadedFile(file, file.Filename)
		c.String(http.StatusOK, file.Filename)
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