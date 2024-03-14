package database

import (
	"context"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/anirudh97/GollabEdit/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *sqlx.DB
var MongoDB *mongo.Client

func InitDB(c *config.DatabaseConfig) error {
	log.Println("Initializing SQL.......")
	connDetails := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.DBName)

	var err error
	DB, err = sqlx.Connect("mysql", connDetails)
	if err != nil {
		return err
	}

	return nil
}

func InitMongo() error {
	log.Println("Initializing MongoDB.......")
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var mongoErr error

	MongoDB, mongoErr = mongo.Connect(ctx, clientOptions)
	if mongoErr != nil {
		return mongoErr
	}

	err := MongoDB.Ping(ctx, nil)
	if err != nil {
		return err
	}

	return nil
}
