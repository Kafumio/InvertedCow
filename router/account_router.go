package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(r *gin.Engine, accountController controller.AccountController) {
	account := r.Group("/account")
	{
		account.POST("/signUp", accountController.SignUp)
	}
}
