package main

import (
	"mpbe/config"
	"mpbe/migrate"
	"mpbe/routes"
)

func main() {
	config.InitDB()
	migrate.AutoMigrate()
	e:=routes.New()
	e.Start(":9090")

}