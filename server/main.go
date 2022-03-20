package main

import (
	"github.com/julianstephens/budget-tracker/cmd/domain"
	"github.com/julianstephens/budget-tracker/cmd/postgresdb"
	"github.com/julianstephens/budget-tracker/cmd/router"
)

func main() {
	postgresdb.ConnectDB()
	postgresdb.DB.AutoMigrate(&domain.User{})

	r := router.SetupRouter()
	r.Run()
}
