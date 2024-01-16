package isogram

import "strings"

func IsIsogram(word string) bool {
	letters:= map[rune]int{}
    s := strings.ToLower(word)
    for _, c := range s {
        if c == '-' || c == ' ' {
            continue
        }
        letters[c] = letters[c] + 1
        if letters[c] > 1 {
            return false
        }
    }
	return true
}
