package shortener

import "github.com/phpCoder88/url-shortener/internal/entities"

type ShortURLRepository interface {
	FindAll(int64, int64) ([]entities.ShortURL, error)
	FindByURL(string) (*entities.ShortURL, error)
	Add(*entities.ShortURL) error
	FindByToken(string) (*entities.ShortURL, error)
	IncURLVisits(int64) error
}
