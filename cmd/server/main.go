package main

import (
	"log"

	"github.com/anirudh97/GollabEdit/internal/api"
	"github.com/anirudh97/GollabEdit/internal/config"
	"github.com/anirudh97/GollabEdit/internal/database"
	"github.com/gin-gonic/gin"

	"github.com/anirudh97/GollabEdit/internal/awsutils"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	if err := database.InitDB(&conf.Database); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := database.InitMongo(); err != nil {
		log.Fatalf("Falied to initialize MongoDB: %v", err)
	}
	_, awsInstErr := awsutils.GetInstance()
	if awsInstErr != nil {
		log.Fatalf("Failed to initialize aws client: %v", awsInstErr)
	}

	r := gin.Default()
	api.SetupRoutes(r)

	if err := r.Run(conf.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	defer database.DB.Close()
}
