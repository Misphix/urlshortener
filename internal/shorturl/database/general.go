package database

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type General struct {
	db *gorm.DB
}

func NewGeneralDatabase(db *gorm.DB) (*General, error) {
	if err := db.AutoMigrate(&ShortURL{}); err != nil {
		return nil, newDatabaseError(err)
	}

	return &General{db: db}, nil
}

func (m *General) InsertURL(url string, expireAt time.Time) (uint, error) {
	data := &ShortURL{
		URL:      url,
		ExpireAt: expireAt,
	}

	if err := m.db.Create(data).Error; err != nil {
		return 0, newDatabaseError(err)
	}

	return data.ID, nil
}

func (m *General) GetURL(id uint) (*ShortURL, error) {
	var data ShortURL
	if err := m.db.First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, newDatabaseError(err)
	}

	return &data, nil
}

func (m *General) DeleteURL(id uint) error {
	var condition ShortURL
	condition.ID = id
	if err := m.db.Delete(&condition).Error; err != nil {
		return newDatabaseError(err)
	}
	return nil
}
