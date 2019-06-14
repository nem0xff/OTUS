package searchMax

type isLess func(interface{}, interface{}) bool

func searchMax(isLessValue isLess, values ...interface{}) interface{} {
	var maxVal interface{}

	for _, val := range values {
		if isLessValue(maxVal, val) {
			maxVal = val
		}
	}

	return maxVal
}
