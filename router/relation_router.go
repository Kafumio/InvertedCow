package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupRelationRoutes(r *gin.Engine, relationController controller.RelationController) {
	account := r.Group("/relation")
	{
		account.GET("/addFollow", relationController.AddFollow)
		account.POST("/cancelFollow", relationController.CancelFollow)
		account.POST("/followList", relationController.GetFollowerList)
		account.GET("/followerList", relationController.GetFollowerList)
	}
}
