package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"unicode"
)

const delimeters = " :,.;\n\t\r?!-()[]"

func printWords(w words) {
	for _, val := range w {
		fmt.Printf("%v - %v\n", val.word, val.count)
	}
}

func main() {
	alasysis := freqAnalysis(getTestText("warpeace.txt"))
	sortedList := sortByCount(alasysis)
	printWords(getFirstTenOfArray(sortedList))
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

func getFirstTenOfArray(w words) words {
	if len(w) > 10 {
		return w[0:10]
	}
	return w
}
