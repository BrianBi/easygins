package router

import "github.com/gin-gonic/gin"

func InitApiRouter(Router *gin.RouterGroup)  {
	ApiRouter := Router.Group("api")
	{
		ApiRouter.POST("api", func(context *gin.Context) {
			context.JSON(200, "api")
		})
	}
}