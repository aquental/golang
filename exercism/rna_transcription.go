package strand

import "bytes"

var nukes = map[rune]rune{
	'C': 'G',
	'G': 'C',
	'T': 'A',
	'A': 'U',
}

func ToRNA(dna string) string {
	var buf bytes.Buffer
	for _, each := range dna {
		buf.WriteRune(nukes[each])
	}
	return buf.String()
}
