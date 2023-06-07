package main

import (
	"github.com/azam-sh/books/initializers"
	"github.com/azam-sh/books/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	routes.StartRoutes()
}
