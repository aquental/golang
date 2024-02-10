package twelve

import "fmt"

var nthString = [12]string{"first", "second", "third", "fourth", "fifth", "sixth", "seventh", "eighth", "ninth", "tenth", "eleventh", "twelfth"}
var verseString = [12]string{"a Partridge in a Pear Tree", "two Turtle Doves, and ", "three French Hens, ", "four Calling Birds, ", "five Gold Rings, ", "six Geese-a-Laying, ",
	"seven Swans-a-Swimming, ", "eight Maids-a-Milking, ", "nine Ladies Dancing, ", "ten Lords-a-Leaping, ", "eleven Pipers Piping, ", "twelve Drummers Drumming, "}


// Verse returns a string representing a verse of the 12 Days of Christmas song.
func Verse(i int) string {
	// Create the initial verse string with the nth day of Christmas.
	str := fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", nthString[i-1])

	// Add the specific gift for each day of Christmas to the verse.
	for idx := i; idx > 0; idx-- {
		str = fmt.Sprintf("%s%s", str, verseString[idx-1])
	}

	// Add a period at the end of the verse and return the complete string.
	str = fmt.Sprintf("%s.", str)
	return str
}

// Song generates the lyrics for the entire song.
func Song() string {
    result := ""
    for n := 1; n <= 12; n++ {
        result += Verse(n)
        if n < 12 {
            result += "\n"
        }
    }
    return result
}
