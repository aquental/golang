package cipher
import (
	"unicode"
	"strings"
	"bytes"
)
// Define the shift and vigenere types here.
type vigenere []rune
type shift int
// Both types should satisfy the Cipher interface.
func NewCaesar() Cipher {
    return NewShift(3)
}
func NewShift(distance int) Cipher {
	if distance < -25 || distance == 0 || distance > 25 {
		return nil
	}
	return vigenere{rune(distance + 'a')}
}
func NewVigenere(key string) Cipher {
	if strings.ReplaceAll(key, "a", "") == "" {
        return nil
    }
	for _, ch := range key {
        if !unicode.IsLower(ch) {
            return nil
        }
    }
	return vigenere(key)
}
func translate(a, b rune, dir int) (rune) {
	b = b - 'a'
    var ret rune
	if dir == 1 { ret = a + b } else { ret = a - b }
	if ret > 'z' { ret -= 26 }
	if ret < 'a' { ret += 26 }
	return ret
}
func compute(v vigenere, input string, dir int) string {
	runes := bytes.NewBuffer(make([]byte, 0, len(input)))
	i:= 0
	for _, ch := range strings.ToLower(input) {
		if unicode.IsLower(ch) {
			char := translate(ch, v[i], dir)
			runes.WriteRune(char)
			i++
		}
		if i == len(v) { i = 0 }
	}
	return runes.String()
}
func (v vigenere) Encode(input string) string {
	return compute(v, input, 1)
}
func (v vigenere) Decode(input string) string {
	return compute(v, input, -1)
}
