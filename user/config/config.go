package config

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