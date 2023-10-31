package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, postController controller.PostController) {
	post := r.Group("/post")
	{
		post.POST("/", postController.Post)
		post.POST("/upload", postController.Upload)
	}
}
