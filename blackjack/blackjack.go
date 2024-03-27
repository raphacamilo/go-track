package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	parsedCard1 := ParseCard(card1)
	parsedCard2 := ParseCard(card2)
	parsedDealerCard := ParseCard(dealerCard)
	cardsSum := parsedCard1 + parsedCard2

	switch {
	// return split (P) if the two card are ace's.
	case parsedCard1 == 11 && parsedCard2 == 11:
		return "P"
	// return Automatically win (W) if the cards sum up to a blackjack and the dealer does not have a ten or higher.
	case cardsSum == 21 && parsedDealerCard < 10:
		return "W"
	// return stand (S) if the cards sum up to a blackjack and the dealer have a ten or higher.
	case cardsSum == 21 && parsedDealerCard >= 10:
		return "S"
	// return stand (S) if the sum up to a value within the range 17 to 20.
	case cardsSum >= 17 && cardsSum <= 20:
		return "S"
	// return stand (S) if the sum up to a value within the range 12 to 16 and the dealer card are below seven.
	case cardsSum >= 12 && cardsSum <= 16 && parsedDealerCard < 7:
		return "S"
	// return hit (H) if the above conditions are not met.
	default:
		return "H"
	}
}
