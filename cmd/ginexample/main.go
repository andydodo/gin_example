package main

import (
	"flag"
	"io"
	"log"
	"os"

	"github.com/LIYINGZHEN/ginexample/configs"
	"github.com/LIYINGZHEN/ginexample/internal/agent/cron"
	"github.com/LIYINGZHEN/ginexample/internal/app/http"
	"github.com/LIYINGZHEN/ginexample/internal/app/postgres"
	"github.com/LIYINGZHEN/ginexample/internal/app/service/linkservice"
	"github.com/LIYINGZHEN/ginexample/internal/app/service/userservice"
	"github.com/LIYINGZHEN/ginexample/pkg/jwt"
	"github.com/spf13/viper"
)

var (
	cfg = flag.String("config", "development", "config file path")
)

func main() {
	flag.Parse()

	// init config
	if err := configs.Init(*cfg); err != nil {
		panic(err)
	}

	postgresConfig := postgres.DBConfig{
		Host:     viper.GetString("psql.host"),
		Port:     viper.GetString("psql.port"),
		User:     viper.GetString("psql.user"),
		Password: viper.GetString("psql.Password"),
		Name:     viper.GetString("psql.db_name"),
	}

	repository := postgres.Initialize(postgresConfig)
	repository.AutoMigrate()

	var logDst io.Writer
	if viper.GetString("log_file") == "" {
		logDst = os.Stdout
	} else {
		file, err := os.OpenFile(viper.GetString("log_file"), os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("error opening logfile %s: %v", viper.GetString("log_file"), err)
		}
		defer file.Close()
		logDst = file
	}

	logger := log.New(logDst, "", log.LstdFlags)
	server := http.AppServer{
		Logger:      logger,
		UserService: userservice.New(repository.UserRepository),
		LinkService: linkservice.New(repository.LinkRepository),

		JWT: jwt.NewJWT(viper.GetString("jwt.private_key"), viper.GetString("jwt.public_key")),
	}

	agent := cron.New(repository, logger)
	go agent.StartCheck()

	server.Run()
}
