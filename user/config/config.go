package config

import (
	"bytes"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Database Database `mapstructure:"db"`
	JWT      JWT      `mapstructure:"jwt"`
}

type Database struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	DBName   string `mapstructure:"dbname"`
	Password string `mapstructure:"password"`
	SSLmode  string `mapstructure:"sslmode"`
}

type JWT struct {
	Secret     string `mapstructure:"secret"`
	Expiration int    `mapstructure:"exp"`
}

func Read() Config {
	v := viper.New()

	v.AddConfigPath(".")
	v.SetConfigType("yml")

	if err := v.ReadConfig(bytes.NewBufferString(Default)); err != nil {
		log.Fatalf("err: %s", err)
	}

	v.SetConfigName("config")

	if err := v.MergeInConfig(); err != nil {
		log.Print("No config file found")
	}

	v.SetEnvPrefix("monitor")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	v.AutomaticEnv()

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("err: %s", err)
	}

	log.Printf("configuration %+v", cfg)

	return cfg
}

func (d Database) Cstring() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s  sslmode=%s",
		d.Host, d.Port, d.User, d.DBName, d.Password, d.SSLmode)
}
