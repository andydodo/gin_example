package configs

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	URL  string
	Port string
}

type PSQL struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type Config struct {
	Server  Server
	PSQL    PSQL
	LogFile string
}

var C Config

func init() {
	env, ok := os.LookupEnv("ENV")
	if !ok {
		env = "development"
	}
	viper.SetConfigFile("configs/" + env + ".yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err = viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
}
