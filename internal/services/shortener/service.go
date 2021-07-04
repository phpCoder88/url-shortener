package shortener

import (
	"database/sql"
	"time"

	"github.com/speps/go-hashids/v2"

	"github.com/phpCoder88/url-shortener/internal/entities"
	"github.com/phpCoder88/url-shortener/internal/repositories/shortener"
)

type Service struct {
	repo shortener.ShortURLRepository
}

func NewService(repo shortener.ShortURLRepository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) FindAll() ([]entities.ShortURL, error) {
	return s.repo.FindAll()
}

func (s *Service) CreateShortURL(url string) (*entities.ShortURL, bool, error) {
	urlRecord, exists, err := s.IsURLExists(url)
	if err != nil {
		return nil, false, err
	}

	if exists {
		return urlRecord, true, nil
	}

	token, err := s.shortURL(url)
	if err != nil {
		return nil, false, err
	}

	urlRecord = &entities.ShortURL{
		LongURL: url,
		Token:   token,
	}

	err = s.repo.Add(urlRecord)
	if err != nil {
		return nil, false, err
	}

	return urlRecord, false, nil
}

func (s *Service) IsURLExists(url string) (*entities.ShortURL, bool, error) {
	urlRecord, err := s.repo.FindByURL(url)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	return urlRecord, true, nil
}

func (s *Service) shortURL(url string) (string, error) {
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	token, err := h.Encode([]int{int(time.Now().UnixNano()) + len(url)})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetFullURL(token string) (string, error) {
	url, err := s.repo.FindByToken(token)
	if err != nil {
		return "", err
	}

	return url.LongURL, nil
}
