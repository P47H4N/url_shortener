package cmd

import (
	"github.com/P47H4N/url_shortener/backend/internals/api"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Start(router *gin.Engine, db *gorm.DB, rdb *redis.Client) {
	apiService := api.NewServices(db, rdb)
	apiController := api.NewControllers(apiService)
	api.Routes(router, apiController)
}
