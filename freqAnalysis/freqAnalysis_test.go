package freqAnalysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFreqAnalysis(t *testing.T) {

	alasysis := freqAnalysis(getTestText("test.txt"))
	sortedList := sortByCount(alasysis)
	fmt.Println(getFirstTenOfArray(sortedList))
}

func printMap(analysis *map[string]int) {
	for val, amount := range *analysis {
		fmt.Printf("%v - %v\n", val, amount)
	}
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
