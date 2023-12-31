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
	// 设置静态文件位置
	r.Static("/static", "/")

	// ping
	r.GET("/ping", c.Ping)
	SetupAccountRoutes(r, controller.AccountController)
	// 动态相关
	SetupPostRoutes(r, controller.PostController)
	// 观看视频相关
	SetupViewRoutes(r, controller.ViewController)
	// follow相关
	SetupRelationRoutes(r, controller.RelationController)

	return r
}
