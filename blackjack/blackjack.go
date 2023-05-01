package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "eight":
		return 8
	case "two":
		return 2
	case "nine":
		return 9
	case "three":
		return 3
	case "ten":
		return 10
	case "four":
		return 4
	case "jack":
		return 10
	case "five":
		return 5
	case "queen":
		return 10
	case "six":
		return 6
	case "king":
		return 10
	case "seven":
		return 7
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	/*
		Stand (S)
		Hit (H)
		Split (P)
		Automatically win (W)
	*/
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	cardSum := card1Value + card2Value
	dealerCardValue := ParseCard(dealerCard)

	switch {

	case card1 == "ace" && card2 == "ace": // If you have a pair of aces you must always split them.
		return "P"

	case cardSum == 21: // If you have a Blackjack (two cards that sum up to a value of 21)

		if dealerCard != "ace" && dealerCardValue != 10 { // and the dealer does not have an ace, a figure or a ten then you automatically win.
			return "W"
		} else { // If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
			return "S"
		}

	case cardSum >= 17 && cardSum <= 20: // If your cards sum up to a value within the range [17, 20] you should always stand.
		return "S"

	case cardSum >= 12 && cardSum <= 16: // If your cards sum up to a value within the range [12, 16] you should always stand
		// unless the dealer has a 7 or higher, in which case you should always hit.
		if dealerCardValue >= 7 {
			return "H"
		} else {
			return "S"
		}

	case cardSum <= 11: // If your cards sum up to 11 or lower you should always hit.
		return "H"
	}

	return ""

}
