package lsproduct

import "errors"

// It returns the largest product and an error if there are any issues with the input.
func LargestSeriesProduct(s string, span int) (int64, error) {
	// Check for invalid input
	if len(s) < span {
		return -1, errors.New("span must be smaller than string length")
	} else if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	var product int64
	// Iterate through the input string to find the largest product
	for i := 0; i < len(s)-span+1; i++ {
		nowValue := int64(1)
		// Calculate the product of the current substring
		for j := 0; j < span; j++ {
			// Check if the current character is a digit
			if s[i+j] < '0' || s[i+j] > '9' {
				return -1, errors.New("digits input must only contain digits")
			}
			nowValue *= int64(s[i+j] - '0')
		}
		// Update the largest product if the current product is larger
		if nowValue > product {
			product = nowValue
		}
	}
	return product, nil
}
