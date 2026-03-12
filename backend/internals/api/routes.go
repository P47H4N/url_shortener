package api

import "github.com/gin-gonic/gin"

func Routes(router *gin.Engine, con *APIControllers) {
	router.POST("/create", con.CreateShortUrl)
	router.GET("/stas/:code")
	router.GET("/:code", con.RedirectUrl)
}