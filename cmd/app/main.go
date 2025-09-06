package main

import (
	"github.com/Reza-Rayan/internal/config"
	"github.com/Reza-Rayan/internal/db"
	"github.com/Reza-Rayan/internal/routes"
)

func main() {
	// Load Config
	cfg := config.LoadConfig()

	// Init Database
	db.InitDB()

	// Setup Router
	r := routes.SetupRouter()

	// Run Server
	r.Run(cfg.Server.Port)

}
