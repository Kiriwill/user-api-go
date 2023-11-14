package mysql

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

func NewConn(connStr string, idleTime int, maxConnections int, maxIdleConnections int) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, fmt.Errorf("could not open connection to database")
	}

	db.SetConnMaxIdleTime(time.Minute * time.Duration(idleTime))
	db.SetMaxOpenConns(maxConnections)
	db.SetMaxIdleConns(maxIdleConnections)

	return &MysqlRepository{db: db, connectionString: connStr}, nil
}

//go:embed migrations/*.sql
var migrationsFolder embed.FS

func (repo *MysqlRepository) migrate() error {
	driver, err := iofs.New(migrationsFolder, "migrations")
	if err != nil {
		return fmt.Errorf("could not create io driver: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", driver, "mysql://"+repo.connectionString)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w -", err)
	}
	v, _, err := m.Version()
	if v == 0 {
		if err = m.Up(); err != nil {
			return fmt.Errorf("could not migrate database: %w - ", err)
		}
	}

	return nil
}

func NewWithMigrate(connStr string, idleTime int, maxConnections int, maxIdleConnections int) (*MysqlRepository, error) {
	repo, err := NewConn(connStr, idleTime, maxConnections, maxIdleConnections)
	if err != nil {
		return nil, err
	}

	err = repo.migrate()
	if err != nil {
		return nil, err
	}

	return repo, nil
}

type MysqlRepository struct {
	db               *sql.DB
	connectionString string
}
