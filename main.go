package main

import (
	"antonio/db"
	"antonio/migrate"
	"antonio/routes"
)
func main() {
	db.InitDB()
	migrate.AutoMigrate()
	e := routes.New()
	e.Start(":8000")
}