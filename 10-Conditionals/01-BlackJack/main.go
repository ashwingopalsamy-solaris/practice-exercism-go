package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	const ace = "ace"
	card1Value := ParseCard(card1)
	card2Value := ParseCard(card2)
	cardsValue := card1Value + card2Value
	dealerCardValue := ParseCard(dealerCard)
	if card1 == ace && card2 == ace {
		return "P"
	} else if cardsValue == 21 {
		if dealerCard != "ace" || dealerCard != "figure" || dealerCard != "ten" {
			return "W"
		}
		return "S"
	} else if cardsValue >= 17 && cardsValue <= 20 {
		return "S"
	} else if cardsValue >= 12 && cardsValue <= 16 {
		if dealerCardValue >= 7 {
			return "H"
		}
		return "S"
	} else if cardsValue == 11 {
		return "H"
	}
	return "H"
}