package main

import (
	"github.com/ConsenSys/fc-retrieval-provider/internal/config"
	"github.com/ConsenSys/fc-retrieval-provider/internal/logger"
	"github.com/ConsenSys/fc-retrieval-provider/internal/services/provider"
	"github.com/rs/zerolog/log"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var conf *viper.Viper

func main() {
	conf = config.NewConfig()
	logger.InitLogger(conf)

	log.Debug().Msg("Running app ...")
	app := fx.New(
		config.Module,
		provider.Module,
		fx.Invoke(provider.Start),
	)
	app.Run()
}
