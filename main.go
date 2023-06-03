package main

import (
	"app-api/configs"
	"app-api/routes"
)

func init() {
	configs.LoadEnv()
	configs.ConnectDatabase()
}

func main() {
	e := routes.Init()
	e.Start(":8000")
}
