package agents

import (
	"context"
	"errors"
	"fmt"
	"log"
	"pocket-pal/configs"
	"pocket-pal/src/domain/entity"
	"strconv"

	"github.com/go-pg/pg/v9"
	vendor_postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	config configs.PostgresConfig
	DB     *gorm.DB
	pgDB   *pg.DB
}

func (p *Postgres) New() *gorm.DB {
	log.Println("Connecting to gorm postgresql")
	if p.config.Host == "" || p.config.Port == 0 {
		panic(errors.New("please set postgresql config"))
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		p.config.Host, p.config.Port, p.config.User, p.config.DbName, p.config.Pass, p.config.SSLMode)

	db, err := gorm.Open(vendor_postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Print("postgres is connected! \n")
	}

	if p.config.DebugMode {
		db = db.Debug()
	}

	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqldb.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqldb.SetConnMaxLifetime(0)

	log.Println("Checking connection to postgresql")
	err = sqldb.Ping()
	if err != nil {
		log.Println("Error connecting to postgresql")
		panic(err)
	}

	log.Println("Auto migrating tables")

	err = db.AutoMigrate(
		&entity.User{},
		&entity.Wallet{},
		&entity.Transactions{},
		&entity.Budgets{},
		&entity.Debts{},
		&entity.GlobalCategory{},
		&entity.CustomCategory{},
	)

	if err != nil {
		log.Fatal("Error migrating tables: ", err)
		panic(err)
	}

	log.Println("Auto migration has been completed")
	log.Println("Postgres gorm is connected")

	return db

}

type dbLogger struct{}

func (l dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (l dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) (err error) {
	query, err := q.FormattedQuery()
	fmt.Println(q.StartTime)
	fmt.Println(query)
	fmt.Println("---------------------------------------------------------------")
	return err
}

func (p *Postgres) NewPgConnection() *pg.DB {
	log.Println("Connecting to postgresql")
	if p.config.Host == "" || p.config.Port == 0 {
		panic(errors.New("please set postgresql config"))
	}
	db := pg.Connect(&pg.Options{
		User:     p.config.User,
		Password: p.config.Pass,
		Addr:     p.config.Host + ":" + strconv.FormatInt(int64(p.config.Port), 10),
		Database: p.config.DbName,
	})
	db.AddQueryHook(dbLogger{})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		panic(err)
	}

	return db
}

func NewPostgres(conf configs.PostgresConfig) *Postgres {
	var pg = &Postgres{config: conf}
	pg.DB = pg.New()
	pg.pgDB = pg.NewPgConnection()
	return pg
}
