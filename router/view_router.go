package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupViewRoutes(r *gin.Engine, viewController controller.ViewController) {
	view := r.Group("/post/view")
	{
		view.GET("/:postId", viewController.GetPostById)
		view.GET("/next", viewController.NextPost)
		view.GET("/pre", viewController.PrePost)
		view.POST("/like", viewController.LikePost)
	}
}
