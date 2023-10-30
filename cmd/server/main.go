package main

import (
	"log"

	"github.com/anirudh97/GollabEdit/internal/api"
	"github.com/anirudh97/GollabEdit/internal/config"
	"github.com/anirudh97/GollabEdit/internal/database"
	"github.com/gin-gonic/gin"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := database.InitDB(&conf.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	r := gin.Default()
	api.SetupRoutes(r)

	if err := r.Run(conf.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer database.DB.Close()
}
