package router

import (
	"gin-course/controllers"
	"gin-course/models"
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func Router() *gin.Engine {
	ginServer := gin.Default()

	ginServer.Use(favicon.New("./favicon_io/favicon.ico"))

	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	userGroup := ginServer.Group("/user")
	{
		userController := &controllers.UserController{}

		// restful api
		userGroup.GET("/info/:userId", func(context *gin.Context) {
			userId := context.Param("userId")
			context.JSON(http.StatusOK, userController.UserInfo(userId))
		})

		// 前端给后端传json
		userGroup.POST("/add", func(context *gin.Context) {
			var user models.User

			_ = context.BindJSON(&user)

			result := userController.AddUser(user.Name)

			context.JSON(http.StatusOK, result)
		})
	}

	return ginServer
}
