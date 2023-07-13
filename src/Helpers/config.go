package Helpers

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`

	MailAddress        string `mapstructure:"MAIL_ADDRESS"`
	MailPassword       string `mapstructure:"MAIL_PASSWORD"`
	UIURL              string `mapstructure:"UIURL"`
	MailServiceAddress string `mapstructure:"MAIL_SERVICE_ADDRESS"`
	MailServicePort    string `mapstructure:"MAIL_SERVICE_PORT"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("cannot start server", err)
	}

	err = viper.Unmarshal(&config)
	return
}
