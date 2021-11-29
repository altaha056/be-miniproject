package main

import (
	"antonio/driver"
	"antonio/routes"
)

func main() {
	driver.InitDB()
	e:=routes.NewRoutes()
	e.Start("localhost:8000")
}