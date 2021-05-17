package service_test

import (
	"context"
	"encoding/base32"
	"testing"
	"time"
	"urlshortener/internal/shorturl/service"

	"github.com/stretchr/testify/assert"
)

func TestURLShorter_Normal(t *testing.T) {
	url := "https://google.com"
	expectedID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, url)
	cache := newFakeCache(false, url)

	s := service.NewShortURL(db, cache)
	id, err := s.Shorter(context.Background(), url, time.Time{})
	assert.Nil(t, err)
	assert.Equal(t, expectedID, id)
}

func TestURLShorter_InvalidURL(t *testing.T) {
	url := "a+123-!@#$%^&*("
	db := newFakeDatabase(false, url)
	cache := newFakeCache(false, url)

	s := service.NewShortURL(db, cache)
	id, err := s.Shorter(context.Background(), url, time.Time{})
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func TestURLShorter_CacheBroken(t *testing.T) {
	url := "https://google.com"
	expectedID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, url)
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	id, err := s.Shorter(context.Background(), url, time.Time{})
	assert.Nil(t, err)
	assert.Equal(t, expectedID, id)
}

func TestURLShorter_DatabaseBroken(t *testing.T) {
	url := "https://google.com"
	db := newBrokenDatabase()
	cache := newFakeCache(false, url)

	s := service.NewShortURL(db, cache)
	id, err := s.Shorter(context.Background(), url, time.Time{})
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func TestURLShorter_BothBroken(t *testing.T) {
	url := "https://google.com"
	db := newBrokenDatabase()
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	id, err := s.Shorter(context.Background(), url, time.Time{})
	assert.NotNil(t, err)
	assert.Empty(t, id)
}

func TestGetURL_Cache(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, "database")
	cache := newFakeCache(false, "cache")

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.Nil(t, err)
	assert.Equal(t, "cache", url)
}

func TestGetURL_CacheMiss(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, "database")
	cache := newFakeCache(true, "cache")

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.Nil(t, err)
	assert.Equal(t, "database", url)
}

func TestGetURL_DatabaseExpired(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(true, "database")
	cache := newFakeCache(true, "cache")

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.NotNil(t, err)
	assert.Empty(t, url)
}

func TestGetURL_CacheBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, "database")
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.Nil(t, err)
	assert.Equal(t, "database", url)
}

func TestGetURL_DatabaseBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newBrokenDatabase()
	cache := newFakeCache(false, "cache")

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.Nil(t, err)
	assert.Equal(t, "cache", url)
}

func TestGetURL_BothBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newBrokenDatabase()
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.NotNil(t, err)
	assert.Empty(t, url)
}

func TestGetURL_ErrorURLID(t *testing.T) {
	urlID := "avasd+324"
	db := newFakeDatabase(false, "database")
	cache := newFakeCache(true, "")

	s := service.NewShortURL(db, cache)
	url, err := s.GetURL(context.Background(), urlID)
	assert.NotNil(t, err)
	assert.Empty(t, url)
}

func TestDeleteURL_Normal(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, "database")
	cache := newFakeCache(false, "cache")

	s := service.NewShortURL(db, cache)
	assert.Nil(t, s.DeleteURL(context.Background(), urlID))
}

func TestDeleteURL_CacheBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newFakeDatabase(false, "database")
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	assert.NotNil(t, s.DeleteURL(context.Background(), urlID))
}

func TestDeleteURL_DatabaseBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newBrokenDatabase()
	cache := newFakeCache(false, "cache")

	s := service.NewShortURL(db, cache)
	assert.NotNil(t, s.DeleteURL(context.Background(), urlID))
}

func TestDeleteURL_BothBroken(t *testing.T) {
	urlID := base32.StdEncoding.EncodeToString([]byte("0000000001"))
	db := newBrokenDatabase()
	cache := newBrokenCache()

	s := service.NewShortURL(db, cache)
	assert.NotNil(t, s.DeleteURL(context.Background(), urlID))
}

func TestDeleteURL_InvalidURLID(t *testing.T) {
	urlID := "1qaz!@#$%^&*()"
	db := newFakeDatabase(false, "database")
	cache := newFakeCache(false, "cache")

	s := service.NewShortURL(db, cache)
	assert.NotNil(t, s.DeleteURL(context.Background(), urlID))
}
