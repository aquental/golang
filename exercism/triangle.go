// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT = 0// not a triangle
    Equ = 1// equilateral
    Iso = 2// isosceles
    Sca = 3// scalene
)

// KindFromSides test which kind of triangle it is
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || a + b <= c || b + c <= a || c + a <= b {
        return NaT
    }
    if a == b && a == c {      
    	return Equ // triangle with equal sides
    }
	if a == b || a == c || b == c {     
    	return Iso  // at least 2 sides same
    }
	if a != b && a !=c || b !=c &&  b != a  {       
    	return Sca  //all sides are different
    }
	return NaT
}

