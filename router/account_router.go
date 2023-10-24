package router

import (
	"InvertedCow/controller"
	"github.com/gin-gonic/gin"
)

func SetupAccountRoutes(r *gin.Engine, accountController controller.AccountController) {
	account := r.Group("/account")
	{
		account.POST("/signIn", accountController.SignIn)
		account.POST("/signUp", accountController.SignUp)
		account.POST("/code/send", accountController.SendAuthCode)
		account.GET("/get/info", accountController.GetAccountInfo)
		account.PUT("/password", accountController.ChangePassword)
		account.PUT("", accountController.UpdateAccount)
	}
}
