package main

import (
	"fmt"
	"time"

	"github.com/kiriwill/desafio-verifymy/pkg/http/restecho"
	"github.com/kiriwill/desafio-verifymy/pkg/repository"
)

func main() {
	repo, err := repository.New(
		config.Database.DSN,
		config.Database.Driver,
		config.Database.IdleTime,
		config.Database.MaxConnections,
		config.Database.MaxIdleConnections,
	)

	if err != nil {
		fmt.Printf(err.Error())
		panic("failed to iniatialize repository ")
	}

	router := restecho.New(
		repo, config.Database.AuthTokenSecret, time.Duration(config.Database.JWTDurationHours))
	router.Start(":8080")
}
