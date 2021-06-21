package postgres

import (
	"fmt"
	"time"

	_ "github.com/jackc/pgx/stdlib" // pgx driver
	"github.com/jmoiron/sqlx"
)

const (
	maxOpenConnections = 60
	connMaxLifetime    = 120
	maxIdleConnections = 30
	connMaxIdleTime    = 20
)

func NewPgConnection(host string, port uint16, dbName, user, password string) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s",
		host,
		port,
		user,
		dbName,
		password,
	)

	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnections)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConnections)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
