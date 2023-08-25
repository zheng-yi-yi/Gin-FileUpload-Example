package main

import (
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const (
	routeUpload = "/upload"
	fieldName   = "file"
)

func initRouter(router *gin.Engine) {
	// 设置静态文件路径，将根 URL 映射到 public 目录
	router.Static("/", "./public")

	// 设置文件上传的路由处理函数
	router.POST(routeUpload, func(c *gin.Context) {
		// 获取表单字段 "name" 的值
		name := c.PostForm("name")

		// 获取上传的文件
		file, err := c.FormFile(fieldName)
		if err != nil {
			c.String(http.StatusBadRequest, "获取文件失败: %s", err)
			return
		}

		// 获取上传文件的文件名
		filename := filepath.Base(file.Filename)

		// 视频保存路径
		dst := filepath.Join("./static/", filename)

		// 将上传的文件保存到服务器
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusInternalServerError, "上传文件失败: %s", err)
			return
		}

		// 返回成功上传的消息
		c.String(http.StatusOK, "%s 成功上传文件 %s.", name, file.Filename)
	})
}
