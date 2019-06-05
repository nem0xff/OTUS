package shortener

import (
	"math"
	"strings"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base     = len(alphabet)
)

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
