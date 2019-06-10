package shortener

import (
	"math"
	"strings"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = len(alphabet)
)

type keyGenerator struct {
}

func NewKeyGenerator() (*keyGenerator, error) {
	keyGenerator := keyGenerator{}
	return &keyGenerator, nil
}

func (k *keyGenerator) GenerateKey(data int) (string, error) {
	var result []byte
	remain := data

	for remain > base-1 {
		result = append(result, alphabet[remain%base])
		remain = remain / base
	}
	result = append(result, alphabet[remain]) // добираем последний остаток

	//Реверсируем строку
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result), nil
}

func (k *keyGenerator) ResolvKey(key string) (int, error) {
	var result int
	lenght := len(key) - 1
	for i, s := range key {
		index := strings.Index(alphabet, string(s))
		result = result + index*int(math.Pow(float64(base), float64(lenght-i)))
	}
	return result, nil
}
