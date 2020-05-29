package isogram

import (
	"strings"
)

// IsIsogram implements the solution
func IsIsogram(word string) bool {

	// Preprocess the words
	word = strings.ToLower(word)
	word = strings.Replace(word, " ", "", -1)
	word = strings.Replace(word, "-", "", -1)

	charCounts := make(map[byte]int)
	for idx := range word {
		charCounts[word[idx]]++
	}

	for key := range charCounts {
		if charCounts[key] > 1 {
			return false
		}
	}
	return true
}
