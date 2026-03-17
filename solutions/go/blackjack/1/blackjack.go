package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    result := 0
    switch {
        case card == "ace":	
            result = 11
        case card == "two":	
            result = 2
        case card == "three": 
            result = 3
        case card == "four": 
            result = 4
        case card == "five": 
            result = 5
        case card == "six":	
            result = 6
        case card == "seven": 
            result = 7
        case card == "eight": 
            result = 8
        case card == "nine": 
            result = 9
        case card == "ten": 
            result = 10
        case card == "jack": 
            result = 10
        case card == "queen": 
            result = 10
        case card == "king": 
            result = 10
        default:
        	result = 0
    }

    return result
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var strategy string = ""

    var sum = ParseCard(card1) + ParseCard(card2)
    var parsedDealerCard = ParseCard(dealerCard)
    switch {
        // pair of aces then split 
        case sum == 22:
        	strategy = "P"
        
        // if sum == 21 {
    	// 		if dealer smaller than 10 then win
    	//		if dealer more than eq 10 then stand
        case sum == 21:
        	if parsedDealerCard < 10 { 
                strategy = "W" 
            } else { strategy = "S" } 
    	
        // if sum >= 17 then stand
        case sum >= 17:
        	strategy = "S"
        
        // if sum between 12 and 16 {
        //		if dealer less than 7 then stand
    	//		if dealer more than 7 then hit
        case sum >= 12 && sum <= 16:
        	if parsedDealerCard < 7 { strategy = "S" 
                                    } else { strategy = "H" } 
    	
        // if sum <= 11 then hit
        case sum <= 11:
        	strategy = "H"
    } 
    
    return strategy
}
