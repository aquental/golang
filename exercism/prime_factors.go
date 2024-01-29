package prime

// Factors returns the prime factors of a given number.
func Factors(n int64) []int64 {
	var result []int64
	for i := int64(2); i <= n; i++ {
		for n%i == 0 {
			result = append(result, i)
			n /= i
		}
	}
	return result
}
