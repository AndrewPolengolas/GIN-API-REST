package main

import (
	"github.com/Polengolas/gin-api-rest/database"
	"github.com/Polengolas/gin-api-rest/routes"
)

func main() {
	database.ConectToDatabase()
	routes.HandleRequests()
}
