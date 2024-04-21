package repository

import "time"

type RedisRepository interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Del(key string) error
	GetWithStruct(key string, structData interface{}) error
	SetWithStruct(key string, value interface{}, expiration time.Duration) error
	SetNX(keylock, value string) (bool, error)
}
