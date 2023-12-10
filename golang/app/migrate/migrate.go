package main

import (
	"app/db"
	"app/model"
	"fmt"
)

func main() {
	dbConnection := db.NewDB()
	defer fmt.Println("migration finished")
	defer db.CloseDB(dbConnection)
	dbConnection.AutoMigrate(&model.Admin{})
}
