package main

import (
	"app/db"
	"app/repository"
	"app/usecase"
	"app/controller"
	"app/router"
)

func main() {
	db := db.NewDB()
	adminRepository := repository.NewAdminRepository(db)
	adminUseCase := usecase.NewAdminUsecase(adminRepository)
	adminController := controller.NewAdminController(adminUseCase)
	e := router.NewRouter(adminController)
	e.Logger.Fatal(e.Start(":8080"))
}
