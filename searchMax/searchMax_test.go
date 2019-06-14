package searchMax

import (
	"fmt"
	"testing"
)

func TestSearchMax(t *testing.T) {
	type myval struct {
		username string
		age      int
	}

	users := []myval{
		{"nick", 12},
		{"mike", 30},
		{"misha", 22},
		{"john", 23},
		{"alex", 19},
	}

	trueResult := myval{"mike", 30}

	vals := make([]interface{}, len(users))

	for i, user := range users {
		vals[i] = user
	}

	var compare isLess = func(one interface{}, two interface{}) bool {

		if one == nil {
			return true
		}

		if two == nil {
			return false
		}

		return one.(myval).age < two.(myval).age
	}

	result := searchMax(compare, vals...)
	t.Logf("Исходные данные:\n%v\n", users)
	t.Logf("Максимальный элемент:\n%v\n", result)

	if result != trueResult {
		t.Error("Вернулось не верное значение")
	}
	fmt.Println(result)

}
