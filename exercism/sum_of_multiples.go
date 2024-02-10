package summultiples

// SumMultiples takes a limit and a list of divisors, and returns the sum of all multiples of the divisors within the limit.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	// Loop through numbers up to the limit
	for i := 0; i < limit; i++ {
		// Loop through the divisors
		for _, v := range divisors {
			// If the divisor is positive and divides evenly into the current number, add it to the sum and break the loop
			if v > 0 && i%v == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
