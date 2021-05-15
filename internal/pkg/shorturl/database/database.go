package database

import "time"

type Database interface {
	InsertURL(string, time.Time) (uint, error)
	GetURL(uint) (string, error)
	DeleteURL(uint) error
}
