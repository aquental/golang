package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	sum := 0
	if len(isbn)!=10{
		return false
	}
	for i, ch := range isbn {
		sum += (int(ch) - 48) * (10 - i)
		if ch == 'X' {
			sum -= 30
		}
	}
	return sum%11 == 0
}
