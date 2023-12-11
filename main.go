package main

import (
	"user-backend/app"
	"user-backend/database"
)

func main() {
	database.StartDbEngine()
	app.StartRoute()
}