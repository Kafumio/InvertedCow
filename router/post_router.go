package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupPostRoutes(r *gin.Engine, postController controller.PostController) {
	account := r.Group("/post")
	{
		account.POST("/", postController.Post)
		account.POST("/upload", postController.Upload)
	}
}
