package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(context.Context, string) (string, error)
	Set(context.Context, string, string, time.Duration) error
	Delete(context.Context, string) error
}
