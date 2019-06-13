package freqAnalysis

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

func TestFreqAnalysis(t *testing.T) {
	//alasysis := freqAnalysis("tst tst one one two three four five six seven seven seven seven nine nine nine ten ten")

	alasysis := freqAnalysis(getTestText("test.txt"))
	//t.Logf("%v", alasysis)
	printMap(&alasysis)
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
