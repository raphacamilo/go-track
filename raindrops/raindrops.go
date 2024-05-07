package raindrops

import "strconv"

func Convert(number int) string {
	label := ""
	if number%3 == 0 {
		label += "Pling"
	}
	if number%5 == 0 {
		label += "Plang"
	}
	if number%7 == 0 {
		label += "Plong"
	}

	if label == "" {
		return strconv.Itoa(number)
	}

	return label
}
