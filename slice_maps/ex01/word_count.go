package word

import (
	"strings"
)

func WordFrequency(text string) map[string]int {
	wordMap := make(map[string]int)
	words := strings.Fields(text)

	for _, word := range words {
		if _, ok := wordMap[word]; ok {
			wordMap[word] = wordMap[word] + 1
			continue
		}
		wordMap[word] = 1
	}

	return wordMap
}
