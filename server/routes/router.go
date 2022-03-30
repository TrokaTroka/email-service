package routes

import (
	"mail-sender/email"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/")
	{
		mail := main.Group("email")
		{
			mail.POST("send", email.Post)
		}
	}

	return router
}
