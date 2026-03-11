package main

import (
	"fmt"
	"gin-web-api-go/config"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func DB(config *config.Config) *sqlx.DB {
	user := config.Database.User
	pw := config.Database.Password
	host := config.Database.Host
	port := config.Database.Port
	dbName := config.Database.Name
	sslMode := config.Database.SSLMode

	db, err := sqlx.Connect("postgres",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", user, pw, host, port, dbName, sslMode))

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	r := gin.Default()
	ginConfig := config.LoadConfig()
	r.Run(fmt.Sprintf(":%d", ginConfig.Server.Port))
	db = DB(ginConfig)
}
