package service_test

import (
	"errors"
	"time"
	"urlshortener/internal/shorturl/database"
)

type fakeDatabase struct {
	isTimeout bool
	url       string
}

var _ database.Database = (*fakeDatabase)(nil)

func newFakeDatabase(isTimeout bool, url string) *fakeDatabase {
	return &fakeDatabase{
		isTimeout: isTimeout,
		url:       url,
	}
}

func (f *fakeDatabase) InsertURL(string, time.Time) (uint, error) {
	return 1, nil
}

func (f *fakeDatabase) GetURL(uint) (*database.ShortURL, error) {
	result := database.ShortURL{
		URL:      f.url,
		ExpireAt: time.Date(9999, time.December, 31, 23, 59, 59, 999, time.UTC),
	}
	if f.isTimeout {
		result.ExpireAt = time.Time{}
	}

	return &result, nil
}

func (f *fakeDatabase) DeleteURL(uint) error {
	return nil
}

type borkenDatabase struct {
}

var _ database.Database = (*borkenDatabase)(nil)

func newBrokenDatabase() *borkenDatabase {
	return &borkenDatabase{}
}

func (b *borkenDatabase) InsertURL(string, time.Time) (uint, error) {
	return 0, errors.New("")
}

func (b *borkenDatabase) GetURL(uint) (*database.ShortURL, error) {
	return nil, errors.New("")
}

func (b *borkenDatabase) DeleteURL(uint) error {
	return errors.New("")
}
