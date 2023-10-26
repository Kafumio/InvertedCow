package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupRelationRoutes(r *gin.Engine, relationController controller.RelationController) {
	account := r.Group("/relation")
	{
		account.POST("/addFollow", relationController.AddFollow)
		account.POST("/cancelFollow", relationController.CancelFollow)
		account.GET("/followList", relationController.GetFollowerList)
		account.GET("/followerList", relationController.GetFollowerList)
	}
}
