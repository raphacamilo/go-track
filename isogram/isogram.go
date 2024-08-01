package isogram

import (
	"strings"
)

func IsIsogram(word string) bool {
	wls := map[rune]bool{}

	for _, letter := range strings.ToLower(word) {
		if letter == 45 || letter == 32 {
			continue
		}

		if _, exists := wls[letter]; exists {
			return false
		}

		wls[letter] = true
	}

	return true
}
