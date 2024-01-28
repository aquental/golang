package armstrong

import (
	"math"
	"strconv"
)
func IsNumber(n int) bool {
	str := strconv.Itoa(n)
    exp := float64(len(str))
	sum := 0
	x := n
	for x > 0 {
        digit := x % 10
		sum += int(math.Pow(float64(digit), exp))
		x /= 10
	}
	return sum == n
}
