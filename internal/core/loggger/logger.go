package logger

import (
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
	"os"
	//"fmt"
)

var logger zerolog.Logger

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	//log = zerolog.New(os.Stdout).Level(zerolog.InfoLevel)
	//log = zerolog.TimeFormatUnix
	//zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func Log() zerolog.Logger {
	return logger.Log()
}