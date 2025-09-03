package main

import (
	"github.com/Reza-Rayan/internal/config"
	"github.com/Reza-Rayan/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load Configs
	cfg := config.LoadConfig()
	// Load Database
	db.InitDB()

	server := gin.Default()

	server.Run(cfg.Server.Port)

}
