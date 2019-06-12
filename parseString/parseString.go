package parseString

import (
	"strconv"
	"unicode/utf8"
)

var currentString string

func dePack(str string) string {
	var result []rune
	currentString = str
	if nextSymbIsNumber(str) {
		return ""
	} else {
		for len(currentString) > 0 {
			symb := getNextSymb()
			numberOfRep := 0
			if !nextSymbIsNumber(currentString) {
				numberOfRep = 1
			} else {
				numberOfRep = getNumber()
			}
			for numberOfRep > 0 {
				result = append(result, symb)
				numberOfRep--
			}
		}
	}

	return string(result)
}

func getNumber() int {
	var i int
	var strNum []rune
	for nextSymbIsNumber(currentString) {
		r, size := utf8.DecodeRuneInString(currentString)
		currentString = currentString[size:]
		strNum = append(strNum, r)
	}
	i, _ = strconv.Atoi(string(strNum))

	return i
}

func getNextSymb() rune {
	r, size := utf8.DecodeRuneInString(currentString)
	currentString = currentString[size:]
	if r == rune([]byte(`\`)[0]) {
		r, size = utf8.DecodeRuneInString(currentString)
		currentString = currentString[size:]
	}
	return r
}

func nextSymbIsNumber(str string) bool {
	firstLetter, _ := utf8.DecodeRuneInString(str)
	return rune([]byte("0")[0]) <= firstLetter && rune([]byte("9")[0]) >= firstLetter

}
