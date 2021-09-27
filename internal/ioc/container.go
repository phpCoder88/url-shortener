package ioc

import (
	"time"

	"github.com/jmoiron/sqlx"

	shortenerRepo "github.com/phpCoder88/url-shortener/internal/repositories/shortener"
	"github.com/phpCoder88/url-shortener/internal/services/shortener"
)

type Container struct {
	ShortenerService *shortener.Service
}

func NewContainer(db *sqlx.DB, queryTimeout time.Duration) *Container {
	return &Container{
		ShortenerService: shortener.NewService(shortenerRepo.NewPgRepository(db, queryTimeout)),
	}
}
