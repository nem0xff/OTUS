package shortener

import (
	"errors"

	"github.com/asaskevich/govalidator"
)

// Shorten - main struct
type Shorten struct {
	repository IRepository
	//deduplication    bool
	enableValidation bool
	//storage          []string
	//hashToNum        map[string]string
}

// NewShorten - create instance with init
func NewShorten(deduplication bool, validation bool) (*Shorten, error) {
	var configRep RepositoryConfig
	var err error

	configRep.deduplication = true

	shorten := Shorten{}
	shorten.repository, err = NewRepository(configRep)
	if err != nil {
		return nil, err
	}

	// TODO: убрать лишнее
	//shorten.hashToNum = make(map[string]string)
	//shorten.deduplication = deduplication
	//shorten.enableValidation = validation

	return &shorten, nil
}

// Shorten - get short link of url
func (s *Shorten) Shorten(urlPath string) (string, error) {
	var key string
	var data Data

	if s.enableValidation {
		if !govalidator.IsURL(urlPath) {
			return "", errors.New("Не валидный Url. " + urlPath)
		}
	}

	data.url = urlPath
	id, err := s.repository.Create(data)
	key = base10ToNewBase(id)

	return key, err
}

// Resolve - resolv short link of url
func (s *Shorten) Resolve(url string) string {
	id := newBaseToBase10(url)
	result, err := s.repository.GetByID(id)
	if err != nil {
		return ""
	}
	return result.url
}
