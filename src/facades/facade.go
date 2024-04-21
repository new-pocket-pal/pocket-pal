package facades

import (
	"pocket-pal/configs"
	"pocket-pal/src/domain/repository"
	"pocket-pal/src/infras/persistence"
	"pocket-pal/src/utils/agents"
	"pocket-pal/src/utils/loggers"
	"pocket-pal/src/utils/redis"

	"github.com/sirupsen/logrus"
)

// Facades is a struct to hold all facades
type Facades struct {
	Logger    *logrus.Logger
	Postgres  *agents.Postgres
	UserRepo  repository.UserRepository
	RedisRepo repository.RedisRepository
	Config    *configs.Config
}

// NewFacades is a constructor to create Facades instance
func NewFacades(config *configs.Config) *Facades {
	log := loggers.NewLogger()

	db := agents.NewPostgres(config.Postgres)
	redisClient, err := redis.NewRedis(config, log)
	if err != nil {
		log.WithError(err).Error("cannot create redis client")
	}

	return &Facades{
		Config:    config,
		Logger:    log,
		Postgres:  db,
		RedisRepo: persistence.NewRedisPersistence(redisClient, log),
		UserRepo:  persistence.NewUserPersistence(db, log),
	}
}
