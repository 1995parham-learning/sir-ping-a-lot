package config

import (
	"bytes"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Nats     Nats     `mapstructure:"nats"`
}

type Nats struct {
	Host  string `mapstructure:"host"`
	Topic string `mapstructure:"topic"`
	Queue string `mapstructure:"queue"`
}

func Read() Config {
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")

	if err := viper.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	viper.SetConfigName("config")

	if err := viper.MergeInConfig(); err != nil {
		log.Print("No config file found")
	}

	viper.SetEnvPrefix("monitor")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	return cfg
}