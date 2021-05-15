package database

import (
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQL struct {
	db *gorm.DB
}

func NewMySQL(dsn string) (*MySQL, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return nil, newDatabaseError(err)
	}

	if err := db.AutoMigrate(&ShortURL{}); err != nil {
		return nil, newDatabaseError(err)
	}

	return &MySQL{db: db}, nil
}

func (m *MySQL) InsertURL(url string, expireAt time.Time) (uint, error) {
	data := &ShortURL{
		URL:      url,
		ExpireAt: expireAt,
	}

	if err := m.db.Create(data).Error; err != nil {
		return 0, newDatabaseError(err)
	}

	return data.ID, nil
}

func (m *MySQL) GetURL(id uint) (string, error) {
	var data ShortURL
	if err := m.db.First(&data, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", err
		}
		return "", newDatabaseError(err)
	}

	return data.URL, nil
}

func (m *MySQL) DeleteURL(id uint) error {
	var condition ShortURL
	condition.ID = id
	if err := m.db.Delete(&condition).Error; err != nil {
		return newDatabaseError(err)
	}
	return nil
}
