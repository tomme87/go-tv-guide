package config

import (
	"github.com/spf13/viper"
	"github.com/tomme87/go-tv-guide/internal/pkg/storage"
)

type config struct {
	MongoDB storage.MongoDB `mapstructure:"mongodb"`
}

// C global config reference
var C config

func (c *config) Init() error {
	viper.SetConfigName("go_tv_guide")
	viper.AddConfigPath("/etc/gotvguide/")
	viper.AddConfigPath("$HOME/.gotvgude")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return err
	}

	return nil
}
