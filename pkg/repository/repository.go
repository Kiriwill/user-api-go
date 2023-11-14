package repository

import (
	"fmt"

	"github.com/kiriwill/desafio-verifymy/pkg/repository/mysql"
	"github.com/kiriwill/desafio-verifymy/pkg/service"
)

var _ service.Repository = &mysql.MysqlRepository{}

func New(connStr string, driver string, idleTime int, maxConnections int, maxIdleConnections int) (service.Repository, error) {
	switch driver {
	case "mysql", "mysqlx":
		return mysql.NewWithMigrate(connStr, idleTime, maxConnections, maxIdleConnections)
	}
	return nil, fmt.Errorf("could not find driver name: %s", driver)
}
