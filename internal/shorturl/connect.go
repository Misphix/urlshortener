package shorturl

import (
	"context"
	"urlshorterner/internal/configmanager"
	"urlshorterner/internal/shorturl/cache"
	"urlshorterner/internal/shorturl/database"

	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func newDatabase(config configmanager.DatabaseConfig) (database.Database, error) {
	mysql, err := newMySQL(config)
	if err != nil {
		return nil, err
	}

	db, err := database.NewGeneralDatabase(mysql)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func newMySQL(config configmanager.DatabaseConfig) (*gorm.DB, error) {
	return gorm.Open(mysql.New(mysql.Config{
		DSN:                       config.DSN,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})
}

func newCache(config configmanager.RedisConfig) (cache.Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: config.Address,
	})
	ctx, cancel := context.WithTimeout(context.Background(), config.DialTimeout)
	defer cancel()
	_, err := client.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	return cache.NewRedisCache(client, config.Expiration), nil
}
