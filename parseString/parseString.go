package parseString

import (
	"unicode/utf8"
)

var result []rune

func dePack(str string) string {
	if firstLetterIsNumber(str) {
		return ""
	}

	return string(result)
}

func getNextSymb() rune {
	return int32(0)
}

func firstLetterIsNumber(str string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(str)
	return rune([]byte("0")[0]) <= firstLetter && rune([]byte("9")[0]) >= firstLetter

}
