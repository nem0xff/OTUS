package parseString

import (
	"testing"
)

var testArray map[string]string

func init() {
	testArray = map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
		`qwe\4\5`:  "qwe45",
		`qwe\45`:   `qwe44444`,
		`qwe\\5`:   `qwe\\\\\`,
		`якцупцоп`: `якцупцоп`,
		`я4ба8qwe`: `яяяябааааааааqwe`,
		`a11ф12`:   `aaaaaaaaaaaфффффффффффф`,
	}
}
func TestDePack(t *testing.T) {
	for key, val := range testArray {
		if dePack(key) != val {
			t.Errorf("Error: key='%v' , val='%v', dePack(key)='%v'", key, val, dePack(key))
		} else {
			t.Logf("Success: key='%v', val='%v', dePack(key)='%v'", key, val, dePack(key))
		}
	}
}

func TestNextSymbIsNumber(t *testing.T) {
	testData := map[string]bool{
		"asd":    false,
		"0asd":   true,
		"9asd":   true,
		"4asd":   true,
		" asd":   false,
		"_asd":   false,
		"4ющц":   true,
		"ющу":    false,
		"ю7щ":    false,
		"d5dfdf": false,
	}
	for testStr, val := range testData {
		if nextSymbIsNumber(testStr) != val {
			t.Errorf("Функциция вернула не верное значение firstLetterIsNumber('%v')=%v, ожидаемое значение %v", testStr, nextSymbIsNumber(testStr), val)
		} else {
			t.Logf(" firstLetterIsNumber('%v')=%v, ожидаемое значение %v", testStr, nextSymbIsNumber(testStr), val)
		}
	}
}
