package lasagna

import "math"

func PreparationTime(layers []string, avaragePreparationTime int) int {
	if avaragePreparationTime == 0 {
		avaragePreparationTime = 2
	}

	return len(layers) * avaragePreparationTime
}

func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}

	return
}

func AddSecretIngredient(friendList, myList []string) {
	myList[len(myList)-1] = friendList[len(friendList)-1]
}

func ScaleRecipe(quantities []float64, portions int) []float64 {
	newQuantities := []float64{}

	for _, amount := range quantities {
		calculatedAmount := amount * (float64(portions) / 2)
		newQuantities = append(newQuantities, math.Round(calculatedAmount*100)/100)
	}

	return newQuantities
}
