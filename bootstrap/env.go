package bootstrap

import (
	"log"

	"github.com/spf13/viper"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	PostgresURL   string `mapstructure:"POSTGRES_URL"`
}

func NewEnv() *Env {
	env := Env{}

	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	if err := viper.Unmarshal(&env); err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}
