package main

import (
	"log"

	"github.com/P47H4N/url_shortener/backend/cmd"
	_ "github.com/P47H4N/url_shortener/backend/docs"
	"github.com/P47H4N/url_shortener/backend/internals/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			URL Shortener API
// @version			1.0
// @description		A professional URL Shortener service built with Go, Gin, GORM and Redis.
// @termsOfService	http://swagger.io/terms/

// @contact.name	API Support
// @contact.url		https://github.com/P47H4N
// @contact.email	shariatullahpathan02@gmail.com

// @host			localhost:8080
// @BasePath		/
func main() {
	// gin.SetMode("release")
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
