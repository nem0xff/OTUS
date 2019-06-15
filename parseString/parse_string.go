package parseString

import (
	"log"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func dePack(str string) string {
	var result strings.Builder
	var symbol rune
	var numberOfRep int

	if !isRightStart(&str) {
		return ""
	}

	for len(str) > 0 {
		symbol = getNextSymb(&str)
		numberOfRep = 1 //дефолтное число повторений

		if symbol == rune([]byte(`\`)[0]) {
			symbol = getNextSymb(&str)
		}

		if isNextSymbNumber(&str) {
			numberOfRep = getNumber(&str)
		}

		for numberOfRep > 0 {
			result.WriteRune(symbol)
			numberOfRep--
		}
	}

	return result.String()
}

func getNumber(str *string) int {
	var strNum strings.Builder

	for isNextSymbNumber(str) {
		r, size := utf8.DecodeRuneInString(*str)
		*str = (*str)[size:]
		strNum.WriteRune(r)
	}

	i, err := strconv.Atoi(strNum.String())
	if err != nil {
		log.Fatal(err)
	}

	return i
}

func getNextSymb(str *string) rune {
	r, size := utf8.DecodeRuneInString(*str)
	*str = (*str)[size:]
	return r
}

func isNextSymbNumber(str *string) bool {
	nextRune, _ := utf8.DecodeRuneInString(*str)
	return unicode.IsDigit(nextRune)
}

// False если строка пустая или начинается с цифры
func isRightStart(str *string) bool {

	if len(*str) == 0 {
		return false
	}

	firstRune, _ := utf8.DecodeRuneInString(*str)
	return !unicode.IsDigit(firstRune)
}
