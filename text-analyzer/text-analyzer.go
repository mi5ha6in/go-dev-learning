package main

import (
	"fmt"
	"strings"
	"unicode"
)

func splitWords(text string) []string {
	runes := []rune(strings.ToLower(text))
	words := make([]string, 0)
	var current strings.Builder

	flush := func() {
		if current.Len() == 0 {
			return
		}

		words = append(words, current.String())
		current.Reset()
	}

	for i, r := range runes {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			current.WriteRune(r)
			continue
		}

		isConnector := r == '-' || r == '\'' || r == '’'
		hasNext := i+1 < len(runes)

		nextIsLetterOrNumber := hasNext &&
			(unicode.IsLetter(runes[i+1]) ||
				unicode.IsNumber(runes[i+1]))

		if isConnector &&
			current.Len() > 0 &&
			nextIsLetterOrNumber {
			current.WriteRune(r)
			continue
		}

		flush()
	}

	flush()

	return words
}

func countWords(text string) map[string]int {
	words := splitWords(text)
	counts := make(map[string]int, len(words))

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func mostFrequent(counts map[string]int) string {
	var result string
	maxCount := 0
	for key, value := range counts {
		if value > maxCount {
			maxCount = value
			result = key
		}
	}

	return result
}

func main() {
	text := "Go is simple, and Go is practical. Go! Кто-то сказал: don't stop."

	counts := countWords(text)

	fmt.Println("Words:", splitWords(text))
	fmt.Println("Counts:", counts)
	fmt.Println("Most frequent:", mostFrequent(counts))
}
