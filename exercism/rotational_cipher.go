package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	var rotated = []rune(plain)
	for i, letter := range rotated {
		if letter >= 'A' && letter <= 'Z' {
			letter = (letter-'A'+int32(shiftKey))%26 + 'A'
		} else if letter >= 'a' && letter <= 'z' {
			letter = (letter-'a'+int32(shiftKey))%26 + 'a'
		}
		rotated[i] = letter
	}
	return string(rotated)
}
