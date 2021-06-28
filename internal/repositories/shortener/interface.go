package shortener

import "github.com/phpCoder88/url-shortener/internal/entities"

type ShortURLRepository interface {
	FindAll() ([]entities.ShortURL, error)
	FindByID(int64) (entities.ShortURL, error)
	Add(entities.ShortURL) (bool, error)
}
