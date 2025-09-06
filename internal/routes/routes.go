package routes

import (
	http "github.com/Reza-Rayan/internal/adapters/http/handlers"
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/applications"
	"github.com/Reza-Rayan/internal/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	userRepo := &repositories.UserRepository{Db: db.DB}

	// Usecase
	authUC := applications.NewAuthUsecase(userRepo)
	friendUC := applications.NewFriendUsecase(userRepo)
	userUC := applications.NewUserUsecase(userRepo)

	// Handler
	authHandler := http.NewAuthHandler(authUC)
	friendHandler := http.NewFriendHandler(friendUC)
	userHandler := http.NewUserHandler(userUC)

	// Routes
	api := routes.Group("/api")
	{
		// Auth Routes
		api.POST("/auth/signup", authHandler.Signup)
		api.POST("/auth/login", authHandler.Login)
		// Friend Routes
		api.GET("/users/:id/friends", friendHandler.GetFriendsHandler)
		api.POST("/users/:id/friends/:friendId", friendHandler.AddFriendHandler)
		//	user Routes
		api.GET("/users/search", userHandler.SearchUsersHandler)
	}

	return routes
}
