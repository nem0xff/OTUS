package freqAnalysis

import (
	"unicode/utf8"
)

const delimeters = " :,.;\n\t\r?!-"

func freqAnalysis(str string) map[string]int {
	result := map[string]int{}
	for len(str) > 0 {
		key := getNextWord(&str)
		if _, ok := result[key]; ok {
			result[key] = result[key] + 1
		} else {
			result[key] = 1
		}
	}
	return result
}

func getNextWord(str *string) string {

	var result []rune

	for (!isNextSymbDelimiter(str)) && (len(*str) > 0) {
		r, size := utf8.DecodeRuneInString(*str)
		(*str) = (*str)[size:]
		result = append(result, r)
	}
	deleteDelimeters(str)
	return string(result)

}

func isNextSymbDelimiter(str *string) bool {
	r, _ := utf8.DecodeRuneInString(*str)
	for _, val := range []byte(delimeters) {
		if rune(val) == r {
			return true
		}
	}

	return false
}

func deleteDelimeters(str *string) {
	for _, val := range []byte(delimeters) {
		r, _ := utf8.DecodeRuneInString(*str)
		if rune(val) == r {
			deleteNextSymb(str)
		}
	}
}

func deleteNextSymb(str *string) {

	if len(*str) > 0 {
		_, size := utf8.DecodeRuneInString(*str)
		*str = (*str)[size:]
	}
}
