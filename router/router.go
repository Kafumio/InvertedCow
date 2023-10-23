package router

import (
	c "InvertedCow/controller"
	"InvertedCow/interceptor"
	"github.com/gin-gonic/gin"
)

// SetupRouter
//
//	@Description: 启动路由
func SetupRouter(
	controller *c.Controller,
) *gin.Engine {

	r := gin.Default()

	// 允许跨域
	r.Use(interceptor.Cors())
	r.Use(interceptor.TokenAuthorize())
	//设置静态文件位置
	r.Static("/static", "/")

	//ping
	r.GET("/ping", c.Ping)
	SetupAccountRoutes(r, controller.AccountController)

	return r
}
