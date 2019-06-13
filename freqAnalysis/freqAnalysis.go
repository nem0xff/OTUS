package freqAnalysis

import (
	"sort"
	"unicode/utf8"
)

const delimeters = " :,.;\n\t\r?!-"

func freqAnalysis(str string) map[string]int {
	result := map[string]int{}
	for len(str) > 0 {
		key := getNextWord(&str)
		if key != "" { // последние знаки препинания могут вернуть пустую строку
			if _, ok := result[key]; ok {
				result[key] = result[key] + 1
			} else {
				result[key] = 1
			}
		}
	}
	return result
}

func getNextWord(str *string) string {

	var result []rune

	for isNextSymbDelimiter(str) {
		deleteDelimeters(str)
	}

	for (!isNextSymbDelimiter(str)) && (len(*str) > 0) {
		r, size := utf8.DecodeRuneInString(*str)
		(*str) = (*str)[size:]
		result = append(result, r)
	}

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

func sortByCount(freqWord map[string]int) words {
	result := make(words, len(freqWord))
	var i int

	for getWord, count := range freqWord {
		result[i] = word{getWord, count}
		i++
	}

	sort.Sort(sort.Reverse(result))
	return result
}

type word struct {
	word  string
	count int
}

type words []word

func (w words) Len() int {
	return len(w)
}

func (w words) Less(i, j int) bool {
	return w[i].count < w[j].count
}

func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func getFirstTenOfArray(w words) words {
	return w[0:10]
}
