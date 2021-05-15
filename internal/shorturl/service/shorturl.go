package service

import (
	"context"
	"encoding/base32"
	"errors"
	"fmt"
	"net/url"
	"time"
	"urlshorterner/internal/shorturl/cache"
	"urlshorterner/internal/shorturl/database"
	"urlshorterner/internal/util"
)

type ShortURL struct {
	db    database.Database
	cache cache.Cache
}

func NewShortURL(db database.Database, cache cache.Cache) *ShortURL {
	return &ShortURL{
		db:    db,
		cache: cache,
	}
}

func (s *ShortURL) Shorter(ctx context.Context, uri string, expireAt time.Time) (string, error) {
	if _, err := url.ParseRequestURI(uri); err != nil {
		return "", fmt.Errorf("invalid url %s", uri)
	}

	id, err := s.db.InsertURL(uri, expireAt)
	if err != nil {
		return "", err
	}

	if err := s.cache.Set(ctx, fmt.Sprint(id), uri, time.Since(expireAt)); err != nil {
		fmt.Println(err.Error())
	}

	urlID := util.PaddingLeadingZero(id)
	urlID = base32.StdEncoding.EncodeToString([]byte(urlID))
	return urlID, err
}

func (s *ShortURL) GetURL(ctx context.Context, urlID string) (string, error) {
	index, err := s.URLIDToIndex(urlID)
	if err != nil {
		return "", err
	}

	url, err := s.cache.Get(ctx, fmt.Sprint(index))
	if err == nil {
		return url, nil
	}

	shortURL, err := s.db.GetURL(index)
	if err != nil {
		return "", err
	}

	if shortURL.ExpireAt.Before(time.Now()) {
		return "", errors.New("url has been expired")
	}

	return shortURL.URL, nil
}

func (s *ShortURL) URLIDToIndex(urlID string) (uint, error) {
	data, err := base32.StdEncoding.DecodeString(urlID)
	if err != nil {
		return 0, nil
	}

	return util.RemoveLeadingZero(string(data))
}

func (s *ShortURL) DeleteURL(ctx context.Context, urlID string) error {
	index, err := s.URLIDToIndex(urlID)
	if err != nil {
		return err
	}

	if err := s.cache.Delete(ctx, fmt.Sprint(index)); err != nil {
		return err
	}

	return s.db.DeleteURL(index)
}
