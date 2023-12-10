package main

import (
	"app/db"
)

// GO_ENV=dev go run seeder/seeder.go seeder/seeder_admins.go
func main() {
	dbConnection := db.NewDB()
	defer db.CloseDB(dbConnection)
	seedAdmins(dbConnection)
}
