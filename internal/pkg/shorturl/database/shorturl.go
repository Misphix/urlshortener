package database

import (
	"time"

	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	URL      string
	ExpireAt time.Time
}
