package main

import (
	"app/db"
	"app/repository"
	"app/usecase"
	"fmt"
)

func main() {
	db := db.NewDB()
	adminRepository := repository.NewAdminRepository(db)
	adminUseCase := usecase.NewAdminUsecase(adminRepository)
	admin, err := adminUseCase.GetAdminById(1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(admin)
}
