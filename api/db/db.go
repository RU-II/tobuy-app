package db

import (
	"fmt"
	"time"
	"tobuy-app/api/env"
	"tobuy-app/api/utils"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zerologadapter"
	"github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	err error
)

func Init() *gorm.DB {
	// get logger
	zlogger := utils.CreateLogger(fmt.Sprintf("./db/logs/access.%s.log", time.Now().Local().Format("20060102")))

	// create gorm.DB instance for PostgreSQL service
	cfg, err := pgx.ParseConfig(env.PostgresDSN())
	if err != nil {
		zlogger.Error().Str("error", err.Error()).Msg("error in pgx.ParseConfig() method")
		return nil
	}
	cfg.Logger = zerologadapter.NewLogger(*zlogger)
	cfg.LogLevel = env.PgxlogLevel()

	// gormのlogger debug時はon
	gormLogger := logger.Discard
	if env.LogLevel() == env.LevelDebug {
		gormLogger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: stdlib.OpenDB(*cfg),
	}), &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic(err)
	}

	return db
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		panic(err)
	}
}
