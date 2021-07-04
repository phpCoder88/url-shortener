package shortener

import "github.com/phpCoder88/url-shortener/internal/entities"

type ShortURLRepository interface {
	FindAll() ([]entities.ShortURL, error)
	FindByURL(string) (*entities.ShortURL, error)
	FindByID(int64) (*entities.ShortURL, error)
	Add(*entities.ShortURL) error
	FindByToken(string) (*entities.ShortURL, error)
}
