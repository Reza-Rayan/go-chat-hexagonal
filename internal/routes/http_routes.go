package routes

import (
	http "github.com/Reza-Rayan/internal/adapters/http/handlers"
	"github.com/Reza-Rayan/internal/adapters/http/middleware"
	"github.com/Reza-Rayan/internal/adapters/http/repositories"
	"github.com/Reza-Rayan/internal/adapters/ws"
	"github.com/Reza-Rayan/internal/applications"
	"github.com/Reza-Rayan/internal/db"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	// ----------------------------
	// HTTP Repositories & Usecases
	// ----------------------------
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

		protected := routes.Group("/api")
		protected.Use(middleware.AuthMiddleware())
		{
			// Friend Routes
			protected.GET("/users/:id/friends", friendHandler.GetFriendsHandler)
			protected.POST("/users/:id/friends/:friendId", friendHandler.AddFriendHandler)
			//	user Routes
			protected.GET("/users/search", userHandler.SearchUsersHandler)
		}
	}

	// ----------------------------
	// WebSocket Hub & Routes
	// ----------------------------
	messageRepo := &repositories.MessageRepository{Db: db.DB}
	messageUC := applications.NewMessageUsecase(messageRepo)

	hub := ws.NewHub()
	go hub.Run()

	wsHandler := ws.NewWSHandler(hub, messageUC)

	wsGroup := routes.Group("/ws")
	wsGroup.Use(middleware.AuthMiddleware())
	{
		wsGroup.GET("/chat", wsHandler.ServeWs())
	}

	return routes
}
