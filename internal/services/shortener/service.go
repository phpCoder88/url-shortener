package shortener

import (
	"database/sql"
	"errors"
	"net/url"
	"strconv"
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

func (s *Service) FindAll(limit, offset int64) ([]entities.ShortURL, error) {
	return s.repo.FindAll(limit, offset)
}

func (s *Service) CreateShortURL(urlStr string) (*entities.ShortURL, bool, error) {
	urlRecord, exists, err := s.IsURLExists(urlStr)
	if err != nil {
		return nil, false, err
	}

	if exists {
		return urlRecord, true, nil
	}

	token, err := s.shortURL(urlStr)
	if err != nil {
		return nil, false, err
	}

	urlRecord = &entities.ShortURL{
		LongURL: urlStr,
		Token:   token,
	}

	err = s.repo.Add(urlRecord)
	if err != nil {
		return nil, false, err
	}

	return urlRecord, false, nil
}

func (s *Service) IsURLExists(urlStr string) (*entities.ShortURL, bool, error) {
	urlRecord, err := s.repo.FindByURL(urlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, false, nil
		}

		return nil, false, err
	}

	return urlRecord, true, nil
}

func (s *Service) shortURL(urlStr string) (string, error) {
	hd := hashids.NewData()
	h, err := hashids.NewWithData(hd)
	if err != nil {
		return "", err
	}

	token, err := h.Encode([]int{int(time.Now().UnixNano()) + len(urlStr)})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *Service) GetFullURL(token string) (string, error) {
	urlRecord, err := s.repo.FindByToken(token)
	if err != nil {
		return "", err
	}

	return urlRecord.LongURL, nil
}

func (s *Service) VisitFullURL(token string) (string, error) {
	urlRecord, err := s.repo.FindByToken(token)
	if err != nil {
		return "", err
	}

	err = s.repo.IncURLVisits(urlRecord.ID)
	if err != nil {
		return "", err
	}

	return urlRecord.LongURL, nil
}

func (s *Service) ParseLimitOffsetQueryParams(query url.Values, param string, defaultVal int64) (int64, error) {
	var paramInt int64
	var err error

	if paramSlice, ok := query[param]; ok {
		if len(paramSlice) > 1 {
			return 0, errors.New("too many values for param: " + param)
		}

		paramInt, err = strconv.ParseInt(paramSlice[0], 10, 64)
		if err != nil {
			return 0, errors.New(param + " param value isn't correct number")
		}

		if paramInt < 0 {
			return 0, errors.New(param + " param value is negative")
		}

		return paramInt, nil
	}

	return defaultVal, nil
}
