/*
Package scrabble implements S]solution to the scrabble problem
*/
package scrabble

import (
	"strings"
)

// Score function
func Score(word string) int {
	scores := map[string]int{
		"AEIOULNRST": 1,
		"DG":         2,
		"BCMP":       3,
		"FHVWY":      4,
		"K":          5,
		"JX":         8,
		"QZ":         10,
	}
	word = strings.ToUpper(word)
	score := 0
	for idx := range word {
		for key := range scores {
			if strings.Contains(key, string(word[idx])) {
				score += scores[key]
			}
		}
	}

	return score
}
