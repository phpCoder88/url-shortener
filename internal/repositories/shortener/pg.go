package shortener

import (
	"context"
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/phpCoder88/url-shortener/internal/dto"
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

func (r *PgRepository) FindAll(limit, offset int64) ([]dto.ShortURLReportDto, error) {
	var rows []dto.ShortURLReportDto
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := `SELECT su.*, COUNT(uv.id) AS visits
				FROM short_urls su
				    LEFT OUTER JOIN url_visits uv ON su.id = uv.url_id
				GROUP BY su.id
				LIMIT $1 OFFSET $2`
	err := r.db.SelectContext(ctx, &rows, query, limit, offset)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (r *PgRepository) Add(model *entities.ShortURL) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := "INSERT INTO short_urls (long_url, token) VALUES ($1, $2)"
	_, err := r.db.ExecContext(ctx, query, model.LongURL, model.Token)
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
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return urlRecord, nil
}

func (r *PgRepository) FindByToken(token string) (*entities.ShortURL, error) {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	urlRecord := new(entities.ShortURL)
	err := r.db.GetContext(ctx, urlRecord, "SELECT * FROM short_urls WHERE token = $1", token)
	if err != nil {
		return nil, err
	}

	return urlRecord, nil
}

func (r *PgRepository) AddURLVisit(urlID int64, userIP string) error {
	ctx, cancel := context.WithTimeout(context.Background(), r.timeout)
	defer cancel()

	query := "INSERT INTO url_visits (url_id, ip, created_at) VALUES ($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, urlID, userIP, time.Now())
	if err != nil {
		return err
	}

	return nil
}
