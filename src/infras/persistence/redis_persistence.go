package persistence

import (
	"pocket-pal/src/domain/repository"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type RedisPersistence struct {
	client *redis.Client
	log    *logrus.Logger
}

func (rs *RedisPersistence) Set(key string, value interface{}, expiration time.Duration) error {
	panic("implement me")

}

func (rs *RedisPersistence) Get(key string) (string, error) {
	panic("implement me")
	return "", nil
}

func (rs *RedisPersistence) Del(key string) error {
	panic("implement me")

}

func (rs *RedisPersistence) GetWithStruct(key string, structData interface{}) error {
	panic("implement me")

}

func (rs *RedisPersistence) SetWithStruct(key string, value interface{}, expiration time.Duration) error {
	panic("implement me")

}

func (rs *RedisPersistence) SetNX(keylock, value string) (bool, error) {
	panic("implement me")

}

func NewRedisPersistence(client *redis.Client, log *logrus.Logger) repository.RedisRepository {
	return &RedisPersistence{
		client: client,
		log:    log,
	}
}
