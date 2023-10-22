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
	corsInterceptor *interceptor.CorsInterceptor,
) *gin.Engine {

	r := gin.Default()

	// 允许跨域
	r.Use(corsInterceptor.Cors())

	//设置静态文件位置
	r.Static("/static", "/")

	SetupAccountRoutes(r, controller.AccountController)

	return r
}
