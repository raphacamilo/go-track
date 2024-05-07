package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("invalid number")
	}

	if n == 1 {
		return 0, nil
	}

	if n%2 == 0 {
		v, err := CollatzConjecture(n / 2)
		return 1 + v, err
	} else {
		v, err := CollatzConjecture(3*n + 1)
		return 1 + v, err
	}
}
