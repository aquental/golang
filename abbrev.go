package main

import (
	"regexp"
	"strings"
)
func Abbreviate(s string) string {
	return strings.ToUpper(regexp.MustCompile(`([A-Za-z])(\w|')*(\W|_)*`).ReplaceAllString(s, "$1"))
}

func main() {
   jargon := "HyperText Markup Language--World Wide Web's"
   abbreviation := Abbreviate(jargon)
   //println("Jargon:", jargon)
   println("Abbreviation:", abbreviation) // Output: Abbreviation: HTMLWW
}

