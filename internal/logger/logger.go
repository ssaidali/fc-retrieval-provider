package logger

import (
	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

func InitLogger(conf *viper.Viper) {
	setTimeFormat()
	setLogLevel(conf)
}

func setTimeFormat() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
}

func setLogLevel(conf *viper.Viper) {
	logLevel := conf.GetString("LOGGING_LEVEL")
	if logLevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else if logLevel == "trace" {
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	} else if logLevel == "warn" {
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	} else if logLevel == "error" {
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	} else if logLevel == "fatal" {
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	} else if logLevel == "panic" {
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}