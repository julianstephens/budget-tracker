package router

import (
	"github.com/gin-gonic/gin"
	"github.com/julianstephens/budget-tracker/cmd/app"
)

func setupUserRoutes(r *gin.RouterGroup) {
	r.GET("", func(c *gin.Context) {
		app.CtrlSet.UserCtrl.GetUsers(c)
	})
	r.GET(":id", func(c *gin.Context) {
		app.CtrlSet.UserCtrl.GetUser(c)
	})
	r.POST("", func(c *gin.Context) {
		app.CtrlSet.UserCtrl.AddUser(c)
	})
	r.PUT("", func(c *gin.Context) {
		app.CtrlSet.UserCtrl.EditUser(c)
	})
	r.DELETE("", func(c *gin.Context) {
		app.CtrlSet.UserCtrl.DeleteUser(c)
	})
}
