package main

import (
	"app-api/configs"
	"app-api/routes"
	"os"
)

func main() {
	configs.LoadEnv()
	configs.ConnectDatabase()
	e := routes.Init()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Start(":" + port)
}
