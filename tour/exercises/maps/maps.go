package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	countWordsMap := make(map[string]int)

	for _, word := range strings.Fields(s) {
		countWordsMap[word]++
	}

	return countWordsMap
}

func main() {
	wc.Test(WordCount)
}
