package main

import (
	"log"

	"github.com/P47H4N/url_shortener/backend/cmd"
	"github.com/P47H4N/url_shortener/backend/internals/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// gin.SetMode("release")
	router := gin.Default()
	router.Use(cors.Default())
	cnf, err := cmd.LoadConfig()
	if err != nil {
		log.Fatalln("Unable to load config.")
	}
	rdb, err := database.NewRedis(cnf)
	db, err := database.NewDB(cnf)
	if err != nil {
		log.Fatalln("Database Connection Failed.", err)
		return
	}
	cmd.Start(router, db, rdb)
	router.Run(":" + cnf.Port)
}