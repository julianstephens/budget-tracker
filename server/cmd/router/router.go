package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julianstephens/budget-tracker/cmd/app"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	app.Setup()
	urlPatterns := GetURLs()

	api := router.Group("/api/")
	{
		users := router.Group(urlPatterns.USER_PATH)
		{
			setupUserRoutes(users)
		}
	}

	return router
}
