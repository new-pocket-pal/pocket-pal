package facades

import (
	"pocket-pal/configs"
	"pocket-pal/src/utils/agents"
	"pocket-pal/src/utils/loggers"

	"github.com/sirupsen/logrus"
)

// Facades is a struct to hold all facades
type Facades struct {
	Logger   *logrus.Logger
	Postgres *agents.Postgres
}

// NewFacades is a constructor to create Facades instance
func NewFacades(config *configs.Config) *Facades {
	log := loggers.NewLogger()

	db := agents.NewPostgres(config.Postgres)

	return &Facades{
		Logger:   log,
		Postgres: db,
	}
}
