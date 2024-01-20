// Package bob answers questions
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

const (
	whatever = "Whatever."
	sure     = "Sure."
	chill    = "Whoa, chill out!"
	calmDown = "Calm down, I know what I'm doing!"
	fine     = "Fine. Be that way!"
)

// Hey 
func Hey(remark string) string {
	if strings.HasSuffix(strings.TrimSpace(remark), "?") && strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		return calmDown
	} else if strings.ToUpper(remark) == remark && strings.ToLower(remark) != remark {
		return chill
	} else if strings.HasSuffix(strings.TrimSpace(remark), "?") {
		return sure
	} else if strings.TrimSpace(remark) == "" {
		return fine
	}
	return whatever
}
