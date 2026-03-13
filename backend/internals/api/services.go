package api

import (
	"errors"
	"time"

	"github.com/P47H4N/url_shortener/backend/internals/helpers"
	"github.com/P47H4N/url_shortener/backend/internals/models"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type APIServices struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewServices(db *gorm.DB, rdb *redis.Client) *APIServices {
	return &APIServices{
		db:  db,
		rdb: rdb,
	}
}

func (srv *APIServices) CreateShortUrl(c *gin.Context, long string) (string, error) {
	var URL models.URL
	if err := srv.db.Where("original_url = ?", long).First(&URL).Error; err == nil {
		return URL.ShortCode, nil
	}
	var short string
	for {
		short = helpers.GenerateShortCode()
		var temp models.URL
		if err := srv.db.Where("short_code = ?", short).First(&temp).Error; err != nil {
			break
		}
	}
	newUrl := &models.URL{
		OriginalURL: long,
		ShortCode:   short,
	}
	if err := srv.db.Create(&newUrl).Error; err != nil {
		return "", errors.New("Url can not be stored.")
	}
	srv.rdb.Set(c, short, long, 24*time.Hour)
	return short, nil
}

func (srv *APIServices) RedirectUrl(c *gin.Context, short string) (string, error) {
	stored, err := srv.rdb.Get(c, short).Result()
	if err == nil {
		return stored, nil
	}
	var URL models.URL
	if err := srv.db.Where("short_code = ?", short).First(&URL).Error; err == nil {
		return URL.OriginalURL, nil
	}
	return "", errors.New("URL not found.")
}
