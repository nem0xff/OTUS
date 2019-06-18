package searchMax

import (
	"reflect"
)

type isLess func(interface{}, interface{}) bool

func searchMax(isLessValue isLess, values interface{}) interface{} {
	var maxVal interface{}

	if reflect.TypeOf(values).Kind() == reflect.Slice {

		vals := reflect.ValueOf(values)

		for i := 0; i < vals.Len(); i++ {
			val := vals.Index(i).Interface()
			if isLessValue(maxVal, val) {
				maxVal = val
			}
		}
	}

	return maxVal
}
