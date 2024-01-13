package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    var r int
	switch card {
        case "ace": r = 11
    	case "two": r = 2
    	case "three": r = 3
    	case "four": r = 4
    	case "five": r = 5
    	case "six": r = 6
    	case "seven": r = 7
  		case "eight": r = 8
    	case "nine": r = 9
  	  	case "ten": r = 10
    	case "jack": r = 10
    	case "queen": r = 10
    	case "king": r = 10
    	default : r = 0
    }
	return r
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    var one, two, dealer int
    one = ParseCard(card1)
    two = ParseCard(card2)
    dealer = ParseCard(dealerCard)

    if (one == 11 && two == 11) { return "P"}
    sum := one + two
    if (sum == 21 && dealer < 10) {return "W"}
    if (sum >= 17 && sum <= 20) {return "S"}
    if (sum >= 12 && sum <= 16) { if (dealer >= 7) {return "H"} else {return "S"}}
    if sum <= 11 {return "H"}
    return "S"
}
