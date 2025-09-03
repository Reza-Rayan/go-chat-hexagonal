package main

import (
	"github.com/Reza-Rayan/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	server := gin.Default()

	server.Run(cfg.Server.Port)

}
