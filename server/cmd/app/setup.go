package app

import (
	"github.com/julianstephens/budget-tracker/cmd/controllers"
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"github.com/julianstephens/budget-tracker/cmd/repositories"
	"github.com/julianstephens/budget-tracker/cmd/services"
)

var CtrlSet *ControllerSet

func Setup() {
	postgresdb.ConnectDB()
	userRepo := repositories.CreateUserRepo(postgresdb.DB)
	userSVC := services.CreateUserService(userRepo)
	userCtrl := controllers.CreateUserController(userSVC)
	CtrlSet = newControllerSet(userCtrl)
}
