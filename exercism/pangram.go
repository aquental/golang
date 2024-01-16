package pangram

import (
	"strings"
	"unicode"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	var letterTracker = make(map[rune]int)
	for _, letter := range input {
		if unicode.IsLetter(letter) {
			letterTracker[letter] = 1
		}
	}
	sum := 0
	for _, v := range letterTracker {
		sum += v
	}
	return sum == 26
}
