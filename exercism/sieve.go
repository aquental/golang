package sieve

// Sieve returns a list of prime numbers up to the given limit
func Sieve(limit int) []int {
    // Initialize a boolean array to keep track of composite numbers
    isComposite := make([]bool, limit+1)
    // Initialize an empty list to store prime numbers
    primes := []int{}
    // Iterate through numbers starting from 2 up to the limit
    for num := 2; num <= limit; num++ {
        // If the number is composite, skip it
        if isComposite[num] {
            continue
        }
        // If the number is prime, add it to the list
        primes = append(primes, num)
        // Mark all multiples of the prime number as composite
        for multiple := num * 2; multiple <= limit; multiple += num {
            isComposite[multiple] = true
        }
    }
    // Return the list of prime numbers
    return primes
}
