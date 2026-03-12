package database

import (
	"context"

	"github.com/P47H4N/url_shortener/backend/cmd"
	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()

func NewRedis(cnf *cmd.Config) (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr: cnf.RedisAddress,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	return rdb, nil
}