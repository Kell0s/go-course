package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	mapWords := make(map[string]int)

	for _, word := range words {
		mapWords[word]++
	}

	return mapWords
}

func main() {
	wc.Test(WordCount)
}
