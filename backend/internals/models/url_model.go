package models

import (
	"time"
)

type URL struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	OriginalURL string    `json:"original_url" gorm:"type:text;not null"`
	ShortCode   string    `json:"short_code" gorm:"type:varchar(20);uniqueIndex;not null"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"-"`
}

type Analytics struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	URLID     uint      `json:"url_id" gorm:"index:idx_url_clicked"`
	IPAddress string    `json:"ip_address" gorm:"type:varchar(50)"`
	UserAgent string    `json:"user_agent" gorm:"type:text"`
	ClickedAt time.Time `json:"clicked_at" gorm:"autoCreateTime"`
}