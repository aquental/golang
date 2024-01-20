package grains
import "fmt"
const Squares int = 64
func Square(number int) (uint64, error) {
    if number < 1 || number > Squares {
		return 0, fmt.Errorf("square must be between 1 and 64")
	}
	return uint64(1 << (number - 1)), nil
}
func Total() uint64 {
	return 1 << Squares - 1
}
