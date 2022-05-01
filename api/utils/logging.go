package utils

import (
	"io"
	"os"
	"tobuy-app/api/env"

	"github.com/rs/zerolog"
)

func CreateLogger(logPath string) *zerolog.Logger {
	logger := zerolog.Nop()

	if env.ZerologLevel() == zerolog.NoLevel {
		return &logger
	}

	if env.EnableLogFile() {
		logfile, err := os.OpenFile(logPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			logger = zerolog.New(os.Stdout).Level(env.ZerologLevel()).With().Timestamp().Logger()
		} else {
			logger = zerolog.New(io.MultiWriter(
				logfile,
				zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
			)).Level(env.ZerologLevel()).With().Timestamp().Logger()
		}
		return &logger
	}

	logger = zerolog.New(os.Stdout).Level(env.ZerologLevel()).With().Timestamp().Logger()
	return &logger
}
