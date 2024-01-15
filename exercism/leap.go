// Package leap calculates if a year is a leap year
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear returns true if year is leap
func IsLeapYear(year int) bool {
    if year % 4 == 0 && year % 100 != 0 || year % 400 == 0 {
        return true
    }
	return false
}
