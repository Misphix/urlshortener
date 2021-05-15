package service

import (
	"encoding/base32"
	"fmt"
	"net/url"
	"time"
	"urlshorterner/internal/pkg/shorturl/database"
	"urlshorterner/internal/pkg/util"

	"github.com/go-redis/redis/v8"
)

type ShortURL struct {
	redisClient *redis.Client
	db          database.Database
}

func NewShortURL(db database.Database) *ShortURL {
	return &ShortURL{
		db: db,
	}
}

func (s *ShortURL) Shorter(uri string, expireAt time.Time) (string, error) {
	if _, err := url.ParseRequestURI(uri); err != nil {
		return "", fmt.Errorf("invalid url %s", uri)
	}

	id, err := s.db.InsertURL(uri, expireAt)
	if err != nil {
		return "", err
	}
	urlID := util.PaddingLeadingZero(id)
	urlID = base32.StdEncoding.EncodeToString([]byte(urlID))
	return urlID, err
}

func (s *ShortURL) GetURL(urlID string) (string, error) {
	index, err := s.URLIDToIndex(urlID)
	if err != nil {
		return "", err
	}

	return s.db.GetURL(index)
}

func (s *ShortURL) URLIDToIndex(urlID string) (uint, error) {
	data, err := base32.StdEncoding.DecodeString(urlID)
	if err != nil {
		return 0, nil
	}

	return util.RemoveLeadingZero(string(data))
}

func (s *ShortURL) DeleteURL(urlID string) error {
	index, err := s.URLIDToIndex(urlID)
	if err != nil {
		return err
	}

	return s.db.DeleteURL(index)
}
