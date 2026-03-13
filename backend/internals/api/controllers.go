package api

import (
	"net/http"
	"strings"

	"github.com/P47H4N/url_shortener/backend/internals/helpers"
	"github.com/P47H4N/url_shortener/backend/internals/models"
	"github.com/gin-gonic/gin"
)

type APIControllers struct {
	srv *APIServices
}

type Body struct {
	LongURL string `json:"long_url" binding:"required" example:"https://supathan.com/"`
}

func NewControllers(srv *APIServices) *APIControllers {
	return &APIControllers{
		srv: srv,
	}
}

func (con *APIControllers) CreateShortUrl(c *gin.Context) {
	var body Body
	if err := c.ShouldBindBodyWithJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Status: "failed",
			Error:  "Invalid request.",
		})
		return
	}
	body.LongURL = strings.TrimSpace(body.LongURL)
	if !helpers.IsValidURL(body.LongURL) {
		c.JSON(http.StatusBadRequest, models.Response{
			Status: "failed",
			Error:  "The URL provided is not valid.",
		})
		return
	}
	short, err := con.srv.CreateShortUrl(c, body.LongURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status: "failed",
			Error:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "URL short successfully.",
		Data: gin.H{
			"long_url":  body.LongURL,
			"short_url": short,
		},
	})
}

func (con *APIControllers) RedirectUrl(c *gin.Context) {
	code := c.Param("code")
	if len(code) != helpers.Length {
		c.JSON(http.StatusBadRequest, models.Response{
			Status: "failed",
			Error:  "Invalid shorturl. Short URL must be in 6 characters.",
		})
		return
	}
	long, err := con.srv.RedirectUrl(c, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Status: "failed",
			Error:  err.Error(),
		})
		return
	}
	c.Redirect(http.StatusMovedPermanently, long)
}
