package routes

import (
	"github.com/Reza-Rayan/internal/adapters"
	http "github.com/Reza-Rayan/internal/adapters/http/handlers"
	"github.com/Reza-Rayan/internal/applications"
	"github.com/Reza-Rayan/internal/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	userRepo := &adapters.UserRepository{Db: db.DB}

	// Usecase
	authUC := applications.NewAuthUsecase(userRepo)

	// Handler
	authHandler := http.NewAuthHandler(authUC)

	// Routes
	api := routes.Group("/api")
	{
		api.POST("/auth/signup", authHandler.Signup)
		api.POST("/auth/login", authHandler.Login)
	}

	return routes
}
