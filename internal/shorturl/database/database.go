package database

import "time"

type Database interface {
	InsertURL(string, time.Time) (uint, error)
	GetURL(uint) (*ShortURL, error)
	DeleteURL(uint) error
}
