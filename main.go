package main

import (
	"books/initializers"
	"books/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	routes.StartRoutes()
}
