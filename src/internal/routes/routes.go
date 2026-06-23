package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/whotterre/odysseus/src/internal/handlers"
	"github.com/whotterre/odysseus/src/internal/repositories"
	"github.com/whotterre/odysseus/src/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(app *gin.Engine, db *gorm.DB){
	userRepo := repositories.NewUserRepository(db)
	authService := services.NewAuthService(userRepo)
	authHandler := handlers.NewAuthHandler(authService)

	authGroup := app.Group("/auth")
	authGroup.POST("/login", authHandler.LoginUser)
	authGroup.POST("/signup", authHandler.SignupUser)
}