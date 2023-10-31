package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupCommentRoutes(r *gin.Engine, commentController controller.CommentController) {
	comment := r.Group("/comment")
	{
		comment.GET("/:postId", commentController.GetPostComments)
		comment.POST("/add", commentController.AddComment)
		comment.POST("/delete", commentController.DeleteComment)
		comment.POST("/like", commentController.LikeComment)
		comment.POST("/dislike", commentController.DisLikeComment)
	}
}
