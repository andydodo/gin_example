package main

import (
	"io"
	"log"
	"os"

	"github.com/LIYINGZHEN/ginexample/configs"
	"github.com/LIYINGZHEN/ginexample/internal/app/http"
	"github.com/LIYINGZHEN/ginexample/internal/app/postgres"
	"github.com/LIYINGZHEN/ginexample/internal/app/service/linkservice"
	"github.com/LIYINGZHEN/ginexample/internal/app/service/userservice"
)

func main() {
	c := configs.C
	postgresConfig := postgres.DBConfig{
		Host:     c.PSQL.Host,
		Port:     c.PSQL.Port,
		User:     c.PSQL.User,
		Password: c.PSQL.Password,
		Name:     c.PSQL.DBName,
	}

	repository := postgres.Initialize(postgresConfig)
	repository.AutoMigrate()

	var logDst io.Writer
	if c.LogFile == "" {
		logDst = os.Stdout
	} else {
		file, err := os.OpenFile(c.LogFile, os.O_APPEND|os.O_CREATE, 0600)
		if err != nil {
			log.Fatalf("error opening logfile %s: %v", c.LogFile, err)
		}
		defer file.Close()
		logDst = file
	}

	server := http.AppServer{
		Logger:      log.New(logDst, "", log.LstdFlags),
		UserService: userservice.New(repository.UserRepository),
		LinkService: linkservice.New(repository.LinkRepository),
	}
	server.Run()
}
