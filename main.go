package main

import (
	"Railway-management-system/database"
	"Railway-management-system/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db := database.InitDB()

	r := gin.Default()
	routes.InitializeRoutes(r, db)
	r.Run(":8999")
}
