package shortener

import (
	"math"
	"strings"
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
	deduplication bool
	storage       []string
	hashToNum     map[string]string
}

// Shorten - get short link of url
func (s *Shorten) Shorten(url string) string {
	s.storage = append(s.storage, url)
	lastPosition := len(s.storage) - 1
	return base10ToNewBase(lastPosition)
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
		//fmt.Printf("remain = %v, base = %v, remainbase = %v\n", remain, base, string(alphabet[remain%base]))
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
		//fmt.Printf("current symbol = %v, i = %v, s = %v, index = %v, result=%v, %v\n", string(s), i, s, index, result, math.Pow(float64(base), float64(i)))
	}
	return result
}
