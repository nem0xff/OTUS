package shortener

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/asaskevich/govalidator"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = len(alphabet)
)

//Shortener - interface of shortener
type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

// Shorten - main struct
type Shorten struct {
	deduplication    bool
	enableValidation bool
	storage          []string
	hashToNum        map[string]string
}

// NewShorten - create instance with init
func NewShorten(deduplication bool, validation bool) (*Shorten, error) {

	shorten := Shorten{}
	shorten.hashToNum = make(map[string]string)
	shorten.deduplication = deduplication
	shorten.enableValidation = validation

	return &shorten, nil
}

// Shorten - get short link of url
func (s *Shorten) Shorten(urlPath string) (string, error) {
	var key string

	if s.enableValidation {
		if !govalidator.IsURL(urlPath) {
			return "", errors.New("Не валидный Url. " + urlPath)
		}
	}

	s.storage = append(s.storage, urlPath)
	lastPosition := len(s.storage) - 1

	if s.deduplication {
		md5sum := fmt.Sprintf("%x", md5.Sum([]byte(urlPath)))
		if val, ok := s.hashToNum[md5sum]; ok {
			key = val
		} else {
			key = base10ToNewBase(lastPosition)
			s.hashToNum[md5sum] = key
		}
	} else {
		key = base10ToNewBase(lastPosition)
	}
	return key, nil
}

// Resolve - resolv short link of url
func (s *Shorten) Resolve(url string) string {
	key := newBaseToBase10(url)
	if key >= len(s.storage) {
		return ""
	}
	return s.storage[key]
}

// SetDeduplicationStatus - set status deduplication link
func (s *Shorten) SetDeduplicationStatus(status bool) error {
	s.deduplication = status
	return nil
}

// GetDeduplicationStatus - get deduplication status
func (s *Shorten) GetDeduplicationStatus() (bool, error) {
	return s.deduplication, nil
}

func base10ToNewBase(i int) string {
	var result []byte
	remain := i

	for remain > base-1 {
		result = append(result, alphabet[remain%base])
		remain = remain / base
	}
	result = append(result, alphabet[remain]) // добираем последний остаток

	//Реверсируем строку
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

func newBaseToBase10(number string) int {
	var result int
	lenght := len(number) - 1
	for i, s := range number {
		index := strings.Index(alphabet, string(s))
		result = result + index*int(math.Pow(float64(base), float64(lenght-i)))
	}
	return result
}
