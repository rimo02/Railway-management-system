package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Railway-management-system/controllers"
	"Railway-management-system/middleware"
)

func InitializeRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api")

	api.POST("/register", func(c *gin.Context) { controllers.Register(c, db) })
	api.POST("/login", func(c *gin.Context) { controllers.Login(c, db) })

	admin := api.Group("/admin")
	admin.Use(middleware.APIKeyAuth())
	admin.POST("/train", func(c *gin.Context) { controllers.AddTrain(c, db) })

	user := api.Group("/")
	user.Use(middleware.JWTAuth())
	user.POST("/book", func(c *gin.Context) { controllers.BookSeat(c, db) })
}
