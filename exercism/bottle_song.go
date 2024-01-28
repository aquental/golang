package bottlesong

import (
	"strings"
)

// Recite returns the verses of the song "99 Bottles of Beer"
var cnt = []string{"No", "One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten"}

// plural returns the plural suffix for a given number
// If the number is not equal to 1, it returns "s", otherwise it returns an empty string.
func plural(n int) string {
	if n != 1 {
		return "s"
	}
	return ""
}

// Recite generates the verses of the "99 Bottles of Beer" song starting from a given number of bottles.
func Recite(start, verses int) (out []string) {
	// Iterate from start to start-verses and generate the verses
	for i := start; i > start-verses; i-- {
		// Construct the bottle string
		bottle := cnt[i] + " green bottle" + plural(i)
		// Generate the verse and append it to the output
		out = append(out, "", 
			bottle + " hanging on the wall,", 
			bottle + " hanging on the wall,", 
			"And if one green bottle should accidentally fall,",
			"There'll be " + strings.ToLower(cnt[i-1]) + " green bottle" + plural(i-1) + " hanging on the wall.")
	}
	// Return the verses, excluding the first empty string
	return out[1:]
}
