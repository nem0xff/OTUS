package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

const boundOfPrint int = 10

func main() {
	alasysis := freqAnalysis(getTestText("warpeace.txt"))
	sortedList := sortByCount(alasysis)
	fmt.Print(sortedList)
}

func getTestText(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func freqAnalysis(str string) map[string]int {

	result := map[string]int{}

	isDelimeter := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	strs := strings.FieldsFunc(str, isDelimeter)

	for _, val := range strs {
		result[val]++
	}
	return result
}
