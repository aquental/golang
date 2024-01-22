package atbash

import (
	"strings"
)

// Atbash returns the Atbash cipher of the given string.
func Atbash(s string) string {
	var sb strings.Builder
    for _, r := range strings.ToLower(s) {
        switch {
            case r >= '0' && r <= '9': sb.WriteRune(r)
            case r >=  'A' && r <= 'Z': sb.WriteRune('Z' - r + 'A')
            case r >=  'a' && r <= 'z': sb.WriteRune('z' - r + 'a')
        }
        if sb.Len() % 6 == 5 {
            sb.WriteRune(' ')
        }
    }
    return strings.TrimSpace(sb.String())
}

