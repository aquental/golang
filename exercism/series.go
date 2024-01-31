package series

// All returns a slice of all substrings of length n in the input string s.
func All(n int, s string) []string {
    var subs []string
    for i := 0; i+n <= len(s); i++ {
        subs = append(subs, s[i:i+n])
    }
    return subs
}

// UnsafeFirst returns the first n characters of the given string.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
