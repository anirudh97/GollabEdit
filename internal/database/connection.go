package database

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/anirudh97/GollabEdit/internal/config"
)

var DB *sqlx.DB

func InitDB(c *config.DatabaseConfig) error {

	connDetails := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username, c.Password, c.Host, c.Port, c.DBName)

	var err error
	DB, err = sqlx.Connect("mysql", connDetails)
	if err != nil {
		return err
	}

	return nil
}
