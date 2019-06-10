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
	}
}
func TestDePack(t *testing.T) {
	for key, val := range testArray {
		if dePack(key) != val {
			t.Errorf("Error: key=%v , val=%v, dePack(key)=%v", key, val, dePack(key))
		} else {
			t.Logf("Success: key=%v, val=%v, dePack(key)=%v", key, val, dePack(key))
		}
	}
}
