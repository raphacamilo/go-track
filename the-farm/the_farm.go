package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(fodder FodderCalculator, cows int) (float64, error) {
	amount, err := fodder.FodderAmount(cows)
	factor, factorErr := fodder.FatteningFactor()

	if err != nil {
		return 0, err
	} else if factorErr != nil {
		return 0, factorErr
	}

	return ((factor * amount) / float64(cows)), nil
}

func ValidateInputAndDivideFood(fodder FodderCalculator, cows int) (float64, error) {
	if cows > 0 {
		return DivideFood(fodder, cows)
	}

	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	cows          int
	customMessage string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.cows, err.customMessage)
}

func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:          cows,
			customMessage: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			cows:          cows,
			customMessage: "no cows don't need food",
		}
	}

	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
