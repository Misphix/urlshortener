package service_test

import (
	"context"
	"errors"
	"time"
	"urlshortener/internal/shorturl/cache"
)

type fakeCache struct {
	isMiss bool
	url    string
}

var _ cache.Cache = (*fakeCache)(nil)

func newFakeCache(isMiss bool, url string) *fakeCache {
	return &fakeCache{
		isMiss: isMiss,
		url:    url,
	}
}

func (f *fakeCache) Get(context.Context, string) (string, error) {
	if f.isMiss {
		return "", errors.New("")
	}
	return f.url, nil
}

func (f *fakeCache) Set(context.Context, string, string, time.Duration) error {
	return nil
}

func (f *fakeCache) Delete(context.Context, string) error {
	return nil
}

type borkenCache struct {
}

var _ cache.Cache = (*borkenCache)(nil)

func newBrokenCache() *borkenCache {
	return &borkenCache{}
}

func (b *borkenCache) Get(context.Context, string) (string, error) {
	return "", errors.New("")
}

func (b *borkenCache) Set(context.Context, string, string, time.Duration) error {
	return errors.New("")
}

func (b *borkenCache) Delete(context.Context, string) error {
	return errors.New("")
}
