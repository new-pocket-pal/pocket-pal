package redis

import (
	"context"
	"pocket-pal/configs"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

func NewRedis(config *configs.Config, l *logrus.Logger) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         config.Redis.Addr,
		Password:     config.Redis.Password,
		DB:           config.Redis.DB,
		ReadTimeout:  50 * time.Second,
		WriteTimeout: 50 * time.Second,
		PoolSize:     10,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		l.WithError(err).Error("cannot ping to client redis")
		return nil, err
	}

	return client, nil
}
