package provider

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Provider struct {
	conf *viper.Viper
}

func NewProvider(conf *viper.Viper) *Provider {
	return &Provider{
		conf: conf,
	}
}

func Start(p *Provider) {
	log.Info().Msg("Provider started !")
}

var Module = fx.Options(
	fx.Provide(
		NewProvider,
	),
)
