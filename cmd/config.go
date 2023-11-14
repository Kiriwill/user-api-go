package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Database struct {
		Driver             string `default:"mysql"`
		User               string `default:"devwill"`
		Passwd             string `default:"supersecret"`
		Net                string `default:"tcp"`
		Addr               string `default:"localhost:3306"`
		DBName             string `default:"verifymychallenge"`
		DSN                string
		IdleTime           int    `default:"3"`
		MaxConnections     int    `default:"10"`
		MaxIdleConnections int    `default:"10"`
		SecretKey          string `default:""`
		JWTDurationHours   int    `default:"72"`
		AuthTokenSecret    string `default:"secret"`
	}
}

var config Config

func init() {
	if err := envconfig.Process("VERIFYMY", &config); err != nil {
		panic(err)
	}

	cfg := mysql.Config{
		User:   config.Database.User,
		Passwd: config.Database.Passwd,
		Net:    config.Database.Net,
		Addr:   config.Database.Addr,
		DBName: config.Database.DBName,
	}

	config.Database.DSN = cfg.FormatDSN()
}
