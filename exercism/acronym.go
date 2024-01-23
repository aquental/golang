// Package acronym receives a phrase and retirn it's acronym
package acronym

import (
	"regexp"
	"strings"
)
func Abbreviate(s string) string {
	return strings.ToUpper(regexp.MustCompile(`([A-Za-z])(\w|')*(\W|_)*`).ReplaceAllString(s, "$1"))
}
