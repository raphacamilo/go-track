package hamming

import (
	"errors"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings don't have same length")
	}

	difference := 0
	for i, v := range a {
		if v != rune(b[i]) {
			difference++
		}
	}

	return difference, nil
}
