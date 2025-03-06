package main

import (
	"log"
	"user-management-app/config"
	"user-management-app/routes"
)

func main() {
	config.ConnectDB()

	r := routes.SetupRouter()

	log.Println("Server running on http://localhost:8080, press CTRL+C to stop")
	r.Run(":8080")
}
