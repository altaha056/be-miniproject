package main

import (
	"antonio/driver"
	"antonio/routes"
)
func main() {
	driver.InitDB()
	conn:=routes.New()
	conn.Start("localhost:8000")
}