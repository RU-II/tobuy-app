package env

import (
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

const (
	ServiceName = "elephantsql"
)

func init() {
	//load .env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}

func PostgresDSN() string {
	return os.Getenv("DATABASE_URL")
}

func LogLevel() LoggerLevel {
	return getLogLevel(os.Getenv("LOGLEVEL"))
}

func ZerologLevel() zerolog.Level {
	return LogLevel().ZerlogLevel()
}

func PgxlogLevel() pgx.LogLevel {
	return LogLevel().PgxLogLevel()
}

func EnableLogFile() bool {
	return strings.EqualFold(os.Getenv("ENABLE_LOGFILE"), "true")
}
