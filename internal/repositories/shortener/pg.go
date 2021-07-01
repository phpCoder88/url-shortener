package shortener

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/phpCoder88/url-shortener/internal/entities"
)

type PgRepository struct {
	db      *sqlx.DB
	timeout time.Duration
}

func NewPgRepository(db *sqlx.DB, timeout time.Duration) *PgRepository {
	return &PgRepository{
		db:      db,
		timeout: timeout,
	}
}

func (r *PgRepository) FindAll() ([]entities.ShortURL, error) {
	return nil, nil
}

func (r *PgRepository) FindByID(id int64) (*entities.ShortURL, error) {
	return nil, nil
}

func (r *PgRepository) Add(model *entities.ShortURL) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	sql := "INSERT INTO short_urls (long_url, token) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, sql, model.LongURL, model.Token)
	if err != nil {
		return err
	}

	return nil
}

func (r *PgRepository) FindByURL(url string) (*entities.ShortURL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	urlRecord := new(entities.ShortURL)
	err := r.db.GetContext(ctx, urlRecord, "SELECT * FROM short_urls WHERE long_url = $1", url)
	if err != nil {
		return nil, err
	}

	return urlRecord, nil
}
