package main

import (
	"inventory-app/app"
	"inventory-app/internal/config"
	"inventory-app/internal/handler"
	"inventory-app/internal/repository"
	"inventory-app/internal/service"

)

func main() {
	db := config.DBConnection()
	
	accRepo := repository.NewAccRepository(db)
	accServ := service.NewAccService(accRepo)
	accHand := handler.NewAccHandler(accServ)

	r := app.SetRouter(accHand)
	r.Run(":8080")
}