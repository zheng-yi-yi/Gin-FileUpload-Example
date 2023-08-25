package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认配置的 Gin 实例
	router := gin.Default()
	// 初始化路由
	initRouter(router)
	// 启动 HTTP 服务器并监听端口
	router.Run(":8080")
}
