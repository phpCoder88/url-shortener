package shortener

import (
	"github.com/phpCoder88/url-shortener/internal/entities"
)

type SliceRepository struct {
	storage []entities.ShortURL
}

func NewSliceRepository() *SliceRepository {
	return &SliceRepository{}
}

func (r *SliceRepository) FindAll(limit, offset int64) ([]entities.ShortURL, error) {
	if len(r.storage) == 0 || offset >= int64(len(r.storage)) || limit == 0 {
		return nil, nil
	}

	if offset+limit < int64(len(r.storage)) {
		return r.storage[offset : offset+limit], nil
	}

	return r.storage[offset:], nil
}

func (r *SliceRepository) Add(model *entities.ShortURL) error {
	r.storage = append(r.storage, entities.ShortURL{
		ID:      int64(len(r.storage)) + 1,
		LongURL: model.LongURL,
		Token:   model.Token,
	})

	return nil
}

func (r *SliceRepository) FindByURL(url string) (*entities.ShortURL, error) {
	for item := range r.storage {
		if r.storage[item].LongURL == url {
			return &r.storage[item], nil
		}
	}

	return nil, nil
}

func (r *SliceRepository) FindByToken(token string) (*entities.ShortURL, error) {
	for item := range r.storage {
		if r.storage[item].Token == token {
			return &r.storage[item], nil
		}
	}

	return nil, nil
}

func (r *SliceRepository) IncURLVisits(id int64) error {
	for item := range r.storage {
		if r.storage[item].ID == id {
			r.storage[item].Visits++
			return nil
		}
	}

	return nil
}
