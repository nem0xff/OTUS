package main

import (
	"fmt"
	"sort"
)

func sortByCount(freqWord map[string]int) words {
	result := make(words, len(freqWord))
	var i int

	for getWord, count := range freqWord {
		result[i] = word{getWord, count}
		i++
	}

	sort.Sort(sort.Reverse(result))
	return result
}

type word struct {
	word  string
	count int
}

type words []word

func (w words) Len() int {
	return len(w)
}

func (w words) Less(i, j int) bool {
	return w[i].count < w[j].count
}

func (w words) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w words) String() string {
	var result string
	for i := 0; i < len(w) && i < boundOfPrint; {
		result += fmt.Sprintf("%v - %v\n", w[i].word, w[i].count)
		i++
	}

	return result
}
