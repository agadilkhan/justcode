package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"lecture10/models"
	"time"
)

const BookCacheTimeout = 10 * time.Minute

type book interface {
	Get(ctx context.Context, key string) (*models.Book, error)
	Set(ctx context.Context, key string, val *models.Book) error
}

type BookCache struct {
	Expiration time.Duration
	RedisCli   *redis.Client
}

func NewBookCache(exp time.Duration, cli *redis.Client) BookCache {
	return BookCache{
		exp,
		cli,
	}
}

func (bc *BookCache) Get(ctx context.Context, key string) (*models.Book, error) {
	value := bc.RedisCli.Get(ctx, key).Val()

	if value == "" {
		return nil, fmt.Errorf("not value from cache")
	}

	var book models.Book
	err := json.Unmarshal([]byte(value), &book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (bc *BookCache) Set(ctx context.Context, key string, value *models.Book) error {
	bookJson, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return bc.RedisCli.Set(ctx, key, string(bookJson), bc.Expiration).Err()
}
