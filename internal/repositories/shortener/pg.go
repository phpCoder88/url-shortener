package shortener

import (
	"github.com/jmoiron/sqlx"
	"github.com/phpCoder88/url-shortener/internal/entities"
)

type PgRepository struct {
	db *sqlx.DB
}

func NewPgRepository(db *sqlx.DB) *PgRepository {
	return &PgRepository{
		db: db,
	}
}

func (r *PgRepository) FindAll() ([]entities.ShortURL, error) {
	return nil, nil
}

func (r *PgRepository) FindByID(id int64) (entities.ShortURL, error) {
	return entities.ShortURL{}, nil
}

func (r *PgRepository) Add(model entities.ShortURL) (bool, error) {
	return false, nil
}
